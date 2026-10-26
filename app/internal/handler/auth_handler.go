package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/db"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/handler_model"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	redismodels "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/redis_models"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/os_config"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authDTO struct {
	username                vo.Username
	password                vo.Password
	bodyContentError        error
	user                    postgresql_entity.User
	authUserError           error
	redisIp                 redismodels.AuthenticationIP
	authIpError             error
	IP                      string
	authenticationUserGroup postgresql_entity.AutherticationAccountGroupsUser
	authenticationGroup     postgresql_entity.AuthenticationAccountGroups
	multipleDeviceError     error
	config                  *os_config.EnvironmentConfig
	redisDevice             redismodels.AuthenticationDevice
	deviceCount             int
	differentDevicesCount   int
	isDifferentDevice       bool
}

// GetUser godoc
// @Summary
// @Description  smoke test
// @Tags         test
// @Produce      json
// @Success      200 {object} string
// @Failure      401 {object} string
// @Router       /api/v1/Authentication [post]
func Authentication(c *gin.Context) {
	db := db.GetDbConnection()
	cache := cache.NewRedisCache()

	dto := &authDTO{}
	dto.config = os_config.NewEnvironmentConfig()

	dto = ensureBodyContentIsValid(c, dto)
	dto = authenticateUser(db, dto)
	dto = dataComparingStepValidarGruposUsuario(cache, c, dto, db)
	dto = dataComparingStepMultipleDevices(cache, c, dto, db)

}

func ensureBodyContentIsValid(c *gin.Context, dto *authDTO) *authDTO {
	var body handler_model.Authentication

	if err := c.ShouldBindJSON(&body); err != nil {
		dto.bodyContentError = fmt.Errorf("no body content was detected")
		return dto
	}

	username, err := vo.NewUsername(body.Username)
	if err != nil {
		dto.bodyContentError = fmt.Errorf("invalid username format")
		return dto
	}
	password, err := vo.NewPassword(body.Password)
	if err != nil {
		dto.bodyContentError = fmt.Errorf("invalid password format")
		return dto
	}
	dto.username = username
	dto.password = password

	return dto
}
func authenticateUser(db *gorm.DB, dto *authDTO) *authDTO {
	var existingUser *postgresql_entity.User

	err := db.Where(&postgresql_entity.User{
		Username: dto.username,
		Password: dto.password,
	}).First(&existingUser).Error

	dto.authUserError = err
	if err == nil && existingUser != nil {
		dto.user = *existingUser
	}

	return dto
}

func dataComparingStep(c *gin.Context, dto *authDTO) *authDTO {
	_ = db.GetDbConnection()
	_ = cache.NewRedisCache()

	if dto.authUserError != nil && dto.authIpError != nil {

	}

	// switch {
	// 	case dto.authUserError != nil && dto.redisIp.Score > 0:
	// 	case dto.authUserError != nil && dto.redisIp.Score > 0:
	// 	case
	// }
	return dto
}

func dataComparingStepSpam() {

}
func spamDetectionStep(dto *authDTO) {

}
func dataComparingStepMultipleDevices(c *cache.RedisCache,
	ctx *gin.Context,
	dto *authDTO,
	db *gorm.DB) *authDTO {
	_ = db

	if dto == nil {
		return nil
	}

	if c == nil {
		c = cache.NewRedisCache()
	}

	var currentDevice string
	if ctx != nil {
		currentDevice = ctx.GetHeader("X-Device-ID")
	}

	userKey := fmt.Sprintf("%d", dto.user.ID)
	legacyRuneKey := string(rune(dto.user.ID))

	// Verificar se existe um usuário armazenado no banco de dados do redis
	cachedDevice, err := c.GetBytes(ctx, userKey)
	if err != nil && legacyRuneKey != userKey {
		if b, errFallback := c.GetBytes(ctx, legacyRuneKey); errFallback == nil {
			cachedDevice = b
			err = nil
			userKey = legacyRuneKey
		}
	}

	if err != nil {
		// Usuário não existe no cache do Redis: registrar o dispositivo atual
		devices := make([]string, 0)
		if currentDevice != "" {
			devices = append(devices, currentDevice)
		}

		authDevice := redismodels.AuthenticationDevice{
			Device:            currentDevice,
			UserId:            dto.user.ID,
			LastDevice:        currentDevice,
			Devices:           devices,
			DeviceCount:       len(devices),
			DifferentDevices:  len(devices),
			IsDifferentDevice: false,
		}

		jsonb, errMarshal := json.Marshal(authDevice)
		if errMarshal == nil {
			_ = c.SetBytes(ctx, userKey, jsonb, 0)
			if legacyRuneKey != userKey {
				_ = c.SetBytes(ctx, legacyRuneKey, jsonb, 0)
			}
		}

		dto.redisDevice = authDevice
		dto.deviceCount = authDevice.DeviceCount
		dto.differentDevicesCount = authDevice.DifferentDevices
		dto.isDifferentDevice = false
		return dto
	}

	// Usuário existe no Redis: desserializar os dados armazenados
	var ad redismodels.AuthenticationDevice
	if errUnmarshal := json.Unmarshal(cachedDevice, &ad); errUnmarshal != nil {
		dto.multipleDeviceError = errUnmarshal
		return dto
	}

	// Identificar o último dispositivo utilizado
	lastDevice := ad.Device
	if lastDevice == "" && ad.LastDevice != "" {
		lastDevice = ad.LastDevice
	}

	// Identificar se o último dispositivo é diferente do atual
	isDifferent := false
	if lastDevice != "" && currentDevice != "" && lastDevice != currentDevice {
		isDifferent = true
	}

	// Garantir que a lista de dispositivos contenha o dispositivo anterior
	if len(ad.Devices) == 0 && lastDevice != "" {
		ad.Devices = []string{lastDevice}
	}

	// Validar e adicionar o dispositivo atual à lista se ainda não existir
	deviceExists := false
	for _, d := range ad.Devices {
		if d == currentDevice {
			deviceExists = true
			break
		}
	}
	if !deviceExists && currentDevice != "" {
		ad.Devices = append(ad.Devices, currentDevice)
	}

	// Validar quantos dispositivos diferentes foram utilizados
	uniqueDevices := make(map[string]struct{})
	for _, d := range ad.Devices {
		if d != "" {
			uniqueDevices[d] = struct{}{}
		}
	}
	if currentDevice != "" {
		uniqueDevices[currentDevice] = struct{}{}
	}
	differentDevicesCount := len(uniqueDevices)

	// Atualizar informações do dispositivo no Redis
	ad.LastDevice = lastDevice
	ad.Device = currentDevice
	ad.DeviceCount = len(ad.Devices)
	ad.DifferentDevices = differentDevicesCount
	ad.IsDifferentDevice = isDifferent

	updatedBytes, errMarshal := json.Marshal(ad)
	if errMarshal == nil {
		_ = c.SetBytes(ctx, userKey, updatedBytes, 0)
		if legacyRuneKey != userKey {
			_ = c.SetBytes(ctx, legacyRuneKey, updatedBytes, 0)
		}
	}

	// Identificar se múltiplos dispositivos ou dispositivo diferente foram utilizados
	if isDifferent {
		dto.multipleDeviceError = fmt.Errorf("different device detected: last device was '%s', current is '%s'", lastDevice, currentDevice)
	} else if differentDevicesCount > 1 {
		dto.multipleDeviceError = fmt.Errorf("multiple devices detected: user has %d different devices", differentDevicesCount)
	}

	dto.redisDevice = ad
	dto.deviceCount = ad.DeviceCount
	dto.differentDevicesCount = ad.DifferentDevices
	dto.isDifferentDevice = isDifferent

	return dto
}
func dataComparingStepValidarGruposUsuario(c *cache.RedisCache,
	ctx *gin.Context,
	dto *authDTO,
	db *gorm.DB) *authDTO {
	//Encontrar IP e validar score
	entity := redismodels.NewAuthenticationIp()
	dto.IP = ctx.ClientIP()

	jsonb, err := json.Marshal(entity)
	if err != nil {
		dto.authIpError = err
	}
	dto.redisIp = *entity
	if entity.Score <= dto.config.AuthenticationSpamScore {
		spamDetectionStep(dto)
	}
	if entity.Score <= 0 {
		entity.Score -= dto.config.AuthenticationScoreFailedAttempt
		dto.authIpError = fmt.Errorf("Authentication failed due to score")
		return dto
	}

	//Validar grupos de usuário
	var bl postgresql_entity.AuthenticationIpBlacklist
	validation := db.Where("IP = ?", dto.IP).First(&bl).Error
	if validation != nil {
		dto.authIpError = fmt.Errorf("ip is blacklisted")
		return dto
	}
	var validationGroup postgresql_entity.AuthenticationAccountGroups

	vg := db.Where("IP = ?", dto.IP).First(&validationGroup)

	if vg != nil {
		dto.authenticationGroup = validationGroup

		var validationGroupUser postgresql_entity.AutherticationAccountGroupsUser
		vgu := db.Where("userId = ?", dto.user.ID).First(&validationGroupUser)

		if vgu.Error != nil && vgu != nil {
			dto.authIpError = fmt.Errorf("User is not allowed to authenticate in this IP")
			return dto
		}

		dto.authenticationUserGroup = validationGroupUser
	}

	cachedIp, err := cache.NewRedisCache().GetBytes(ctx, dto.IP)

	//Validar autenticação de usuário
	if dto.authUserError != nil {
		entity.Score -= dto.config.AuthenticationScoreFailedAttempt
	}

	if err != nil {
		cache.NewRedisCache().
			SetBytes(ctx, ctx.ClientIP(), jsonb, 360)
	} else {
		json.Unmarshal(cachedIp, &entity)
	}
	//

	return dto
}
