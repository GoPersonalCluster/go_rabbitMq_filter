package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/auth"
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
	plainPassword           string
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
type authResponseDTO struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// GetUser godoc
// @Summary
// @Description  smoke test
// @Tags         test
// @Produce      json
// @Param user body handler_model.Authentication true "User data"
// @Success      200 {object} string
// @Failure      401 {object} string
// @Router       /api/v1/authentication [post]
func Authentication(c *gin.Context) {
	db := db.GetDbConnection()
	cache := cache.NewRedisCache()

	dto := &authDTO{}
	dto.config = os_config.NewEnvironmentConfig()

	dto = ensureBodyContentIsValid(c, dto)
	if dto.bodyContentError != nil {
		dataComparingStep(c, dto)
		return
	}

	dto = authenticateUser(db, dto)
	dto = dataComparingStepValidarGruposUsuario(cache, c, dto, db)
	if dto.authUserError == nil && dto.user.ID != 0 {
		dto = dataComparingStepMultipleDevices(cache, c, dto, db)
	}

	dataComparingStep(c, dto)
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
	dto.plainPassword = body.Password

	return dto
}
func authenticateUser(db *gorm.DB, dto *authDTO) *authDTO {
	var existingUser postgresql_entity.User

	err := db.Where("username = ?", dto.username).First(&existingUser).Error
	if err != nil {
		dto.authUserError = err
		return dto
	}

	if !existingUser.VerifyPassword(dto.plainPassword) {
		dto.authUserError = fmt.Errorf("invalid credentials")
		return dto
	}

	dto.user = existingUser
	return dto
}

func dataComparingStep(c *gin.Context, dto *authDTO) *authResponseDTO {
	if dto == nil {
		resp := &authResponseDTO{
			Error: "dto cannot be nil",
		}
		if c != nil {
			c.JSON(http.StatusInternalServerError, resp)
		}
		return resp
	}

	if dto.bodyContentError != nil {
		resp := &authResponseDTO{
			Error: dto.bodyContentError.Error(),
		}
		if c != nil {
			c.JSON(http.StatusBadRequest, resp)
		}
		return resp
	}

	if dto.authUserError != nil {
		resp := &authResponseDTO{
			Error: dto.authUserError.Error(),
		}
		if c != nil {
			c.JSON(http.StatusUnauthorized, resp)
		}
		return resp
	}

	if dto.user.ID == 0 {
		resp := &authResponseDTO{
			Error: "invalid credentials",
		}
		if c != nil {
			c.JSON(http.StatusUnauthorized, resp)
		}
		return resp
	}

	if dto.authIpError != nil {
		resp := &authResponseDTO{
			Error: dto.authIpError.Error(),
		}
		if c != nil {
			c.JSON(http.StatusUnauthorized, resp)
		}
		return resp
	}

	if dto.redisIp.Score < 0 || (!dto.redisIp.CreatedAt.IsZero() && dto.redisIp.Score <= 0) {
		resp := &authResponseDTO{
			Error: "Authentication failed due to score",
		}
		if c != nil {
			c.JSON(http.StatusUnauthorized, resp)
		}
		return resp
	}

	if dto.multipleDeviceError != nil {
		resp := &authResponseDTO{
			Error: dto.multipleDeviceError.Error(),
		}
		if c != nil {
			c.JSON(http.StatusUnauthorized, resp)
		}
		return resp
	}

	emailStr := dto.user.Email.String()
	if emailStr == "" {
		if val, err := dto.user.Email.Value(); err == nil && val != nil {
			emailStr = fmt.Sprintf("%v", val)
		}
	}

	role := "user"
	if dto.authenticationGroup.Name != "" {
		role = dto.authenticationGroup.Name
	}

	token, err := auth.GenerateToken(dto.user.ID, emailStr, role)
	if err != nil {
		resp := &authResponseDTO{
			Error: fmt.Sprintf("failed to generate token: %s", err.Error()),
		}
		if c != nil {
			c.JSON(http.StatusInternalServerError, resp)
		}
		return resp
	}

	resp := &authResponseDTO{
		Token:   token,
		Message: "authentication successful",
	}
	if c != nil {
		c.JSON(http.StatusOK, resp)
	}
	return resp
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
	if c == nil {
		c = cache.NewRedisCache()
	}

	//Encontrar IP e validar score
	entity := redismodels.NewAuthenticationIp()
	dto.IP = ctx.ClientIP()

	cachedIp, err := c.GetBytes(ctx, dto.IP)
	if err == nil {
		_ = json.Unmarshal(cachedIp, entity)
	}

	//Validar autenticação de usuário
	if dto.authUserError != nil {
		entity.Score -= dto.config.AuthenticationScoreFailedAttempt
	}

	dto.redisIp = *entity
	if entity.Score <= dto.config.AuthenticationSpamScore {
		spamDetectionStep(dto)
	}
	if entity.Score <= 0 {
		entity.Score -= dto.config.AuthenticationScoreFailedAttempt
		dto.redisIp = *entity
		dto.authIpError = fmt.Errorf("Authentication failed due to score")
		if jsonb, errMarshal := json.Marshal(entity); errMarshal == nil {
			_ = c.SetBytes(ctx, dto.IP, jsonb, 360)
		}
		return dto
	}

	if jsonb, errMarshal := json.Marshal(entity); errMarshal == nil {
		_ = c.SetBytes(ctx, dto.IP, jsonb, 360)
	} else {
		dto.authIpError = errMarshal
	}

	//Validar grupos de usuário
	var bl postgresql_entity.AuthenticationIpBlacklist
	validation := db.Where("IP = ?", dto.IP).First(&bl).Error
	if validation == nil && bl.IP != "" {
		dto.authIpError = fmt.Errorf("ip is blacklisted")
		return dto
	}

	var validationGroup postgresql_entity.AuthenticationAccountGroups
	vg := db.Where("IP = ?", dto.IP).First(&validationGroup)

	if vg.Error == nil && validationGroup.IP != "" {
		dto.authenticationGroup = validationGroup

		var validationGroupUser postgresql_entity.AutherticationAccountGroupsUser
		vgu := db.Where("userId = ? AND groupId = ?", dto.user.ID, validationGroup.IP).First(&validationGroupUser)

		if vgu.Error != nil {
			dto.authIpError = fmt.Errorf("User is not allowed to authenticate in this IP")
			return dto
		}

		dto.authenticationUserGroup = validationGroupUser
	}

	return dto
}
