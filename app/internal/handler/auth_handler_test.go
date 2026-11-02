package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/auth"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	redismodels "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/redis_models"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestRedis(t *testing.T) (*cache.RedisCache, *miniredis.Miniredis) {
	t.Helper()
	mr := miniredis.RunT(t)

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	c := cache.NewRedisCacheWithClient(rdb)
	return c, mr
}

func createTestGinContext(deviceID string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("POST", "/api/v1/Authentication", nil)
	if deviceID != "" {
		req.Header.Set("X-Device-ID", deviceID)
	}
	ctx.Request = req
	return ctx, w
}

func TestDataComparingStepMultipleDevices_UserNotInRedis(t *testing.T) {
	c, _ := setupTestRedis(t)
	ctx, _ := createTestGinContext("device-phone-1")

	dto := &authDTO{
		user: postgresql_entity.User{
			ID: 10,
		},
	}

	resultDto := dataComparingStepMultipleDevices(c, ctx, dto, nil)

	require.NotNil(t, resultDto)
	assert.NoError(t, resultDto.multipleDeviceError)
	assert.Equal(t, 1, resultDto.deviceCount)
	assert.Equal(t, 1, resultDto.differentDevicesCount)
	assert.False(t, resultDto.isDifferentDevice)
	assert.Equal(t, "device-phone-1", resultDto.redisDevice.Device)
	assert.Equal(t, "device-phone-1", resultDto.redisDevice.LastDevice)
	assert.Equal(t, []string{"device-phone-1"}, resultDto.redisDevice.Devices)

	// Verificar se foi persistido no redis
	val, err := c.GetBytes(ctx, fmt.Sprintf("%d", dto.user.ID))
	require.NoError(t, err)

	var saved redismodels.AuthenticationDevice
	err = json.Unmarshal(val, &saved)
	require.NoError(t, err)
	assert.Equal(t, "device-phone-1", saved.Device)
	assert.Equal(t, 1, saved.DeviceCount)
	assert.Equal(t, 1, saved.DifferentDevices)
}

func TestDataComparingStepMultipleDevices_SameDeviceAgain(t *testing.T) {
	c, _ := setupTestRedis(t)

	// Primeiro login
	ctx1, _ := createTestGinContext("device-phone-1")
	dto := &authDTO{
		user: postgresql_entity.User{
			ID: 20,
		},
	}
	dto = dataComparingStepMultipleDevices(c, ctx1, dto, nil)
	require.NoError(t, dto.multipleDeviceError)

	// Segundo login com o MESMO dispositivo
	ctx2, _ := createTestGinContext("device-phone-1")
	dto2 := &authDTO{
		user: postgresql_entity.User{
			ID: 20,
		},
	}
	resultDto := dataComparingStepMultipleDevices(c, ctx2, dto2, nil)

	require.NotNil(t, resultDto)
	assert.NoError(t, resultDto.multipleDeviceError)
	assert.False(t, resultDto.isDifferentDevice)
	assert.Equal(t, 1, resultDto.deviceCount)
	assert.Equal(t, 1, resultDto.differentDevicesCount)
	assert.Equal(t, "device-phone-1", resultDto.redisDevice.Device)
	assert.Equal(t, "device-phone-1", resultDto.redisDevice.LastDevice)
}

func TestDataComparingStepMultipleDevices_DifferentDevice(t *testing.T) {
	c, _ := setupTestRedis(t)

	// Primeiro login com dispositivo 1
	ctx1, _ := createTestGinContext("device-phone-1")
	dto := &authDTO{
		user: postgresql_entity.User{
			ID: 30,
		},
	}
	dto = dataComparingStepMultipleDevices(c, ctx1, dto, nil)
	require.NoError(t, dto.multipleDeviceError)

	// Segundo login com dispositivo DIFERENTE
	ctx2, _ := createTestGinContext("device-laptop-2")
	dto2 := &authDTO{
		user: postgresql_entity.User{
			ID: 30,
		},
	}
	resultDto := dataComparingStepMultipleDevices(c, ctx2, dto2, nil)

	require.NotNil(t, resultDto)
	assert.Error(t, resultDto.multipleDeviceError)
	assert.True(t, resultDto.isDifferentDevice)
	assert.Equal(t, 2, resultDto.deviceCount)
	assert.Equal(t, 2, resultDto.differentDevicesCount)
	assert.Equal(t, "device-laptop-2", resultDto.redisDevice.Device)
	assert.Equal(t, "device-phone-1", resultDto.redisDevice.LastDevice)
	assert.ElementsMatch(t, []string{"device-phone-1", "device-laptop-2"}, resultDto.redisDevice.Devices)

	// Terceiro login com um terceiro dispositivo DIFERENTE
	ctx3, _ := createTestGinContext("device-tablet-3")
	dto3 := &authDTO{
		user: postgresql_entity.User{
			ID: 30,
		},
	}
	resultDto3 := dataComparingStepMultipleDevices(c, ctx3, dto3, nil)

	require.NotNil(t, resultDto3)
	assert.Error(t, resultDto3.multipleDeviceError)
	assert.True(t, resultDto3.isDifferentDevice)
	assert.Equal(t, 3, resultDto3.deviceCount)
	assert.Equal(t, 3, resultDto3.differentDevicesCount)
	assert.Equal(t, "device-tablet-3", resultDto3.redisDevice.Device)
	assert.Equal(t, "device-laptop-2", resultDto3.redisDevice.LastDevice)
	assert.ElementsMatch(t, []string{"device-phone-1", "device-laptop-2", "device-tablet-3"}, resultDto3.redisDevice.Devices)
}

func TestDataComparingStepMultipleDevices_SwitchBackToKnownDevice(t *testing.T) {
	c, _ := setupTestRedis(t)

	// Login 1: phone
	ctx1, _ := createTestGinContext("device-phone")
	dto := &authDTO{user: postgresql_entity.User{ID: 40}}
	dataComparingStepMultipleDevices(c, ctx1, dto, nil)

	// Login 2: laptop
	ctx2, _ := createTestGinContext("device-laptop")
	dto2 := &authDTO{user: postgresql_entity.User{ID: 40}}
	dataComparingStepMultipleDevices(c, ctx2, dto2, nil)

	// Login 3: switch back to phone (already known, but different from last which was laptop)
	ctx3, _ := createTestGinContext("device-phone")
	dto3 := &authDTO{user: postgresql_entity.User{ID: 40}}
	resultDto := dataComparingStepMultipleDevices(c, ctx3, dto3, nil)

	require.NotNil(t, resultDto)
	assert.Error(t, resultDto.multipleDeviceError)
	assert.True(t, resultDto.isDifferentDevice)
	assert.Equal(t, 2, resultDto.deviceCount)            // Continua sendo 2 dispositivos que o usuário possui
	assert.Equal(t, 2, resultDto.differentDevicesCount)   // 2 dispositivos diferentes utilizados
	assert.Equal(t, "device-phone", resultDto.redisDevice.Device)
	assert.Equal(t, "device-laptop", resultDto.redisDevice.LastDevice)
}

func TestDataComparingStepMultipleDevices_LegacyRedisFormat(t *testing.T) {
	c, mr := setupTestRedis(t)

	// Simular dado legado armazenado no Redis com apenas {"device":"old-dev","userid":50}
	legacyJSON := `{"device":"old-dev","userid":50}`
	mr.Set("50", legacyJSON)

	ctx, _ := createTestGinContext("new-dev")
	dto := &authDTO{user: postgresql_entity.User{ID: 50}}
	resultDto := dataComparingStepMultipleDevices(c, ctx, dto, nil)

	require.NotNil(t, resultDto)
	assert.True(t, resultDto.isDifferentDevice)
	assert.Equal(t, "old-dev", resultDto.redisDevice.LastDevice)
	assert.Equal(t, "new-dev", resultDto.redisDevice.Device)
	assert.Equal(t, 2, resultDto.deviceCount)
	assert.Equal(t, 2, resultDto.differentDevicesCount)
	assert.ElementsMatch(t, []string{"old-dev", "new-dev"}, resultDto.redisDevice.Devices)
}

func TestDataComparingStep_NilDTO(t *testing.T) {
	ctx, w := createTestGinContext("")
	resp := dataComparingStep(ctx, nil)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotEmpty(t, resp.Error)
	assert.Empty(t, resp.Token)
}

func TestDataComparingStep_BodyContentError(t *testing.T) {
	ctx, w := createTestGinContext("")
	dto := &authDTO{
		bodyContentError: errors.New("no body content was detected"),
	}
	resp := dataComparingStep(ctx, dto)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "no body content was detected", resp.Error)
	assert.Empty(t, resp.Token)
}

func TestDataComparingStep_AuthUserError(t *testing.T) {
	ctx, w := createTestGinContext("")
	dto := &authDTO{
		authUserError: errors.New("record not found"),
	}
	resp := dataComparingStep(ctx, dto)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "record not found", resp.Error)
	assert.Empty(t, resp.Token)
}

func TestDataComparingStep_UserZeroID(t *testing.T) {
	ctx, w := createTestGinContext("")
	dto := &authDTO{
		user: postgresql_entity.User{
			ID: 0,
		},
	}
	resp := dataComparingStep(ctx, dto)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "invalid credentials", resp.Error)
	assert.Empty(t, resp.Token)
}

func TestDataComparingStep_AuthIpError(t *testing.T) {
	ctx, w := createTestGinContext("")
	dto := &authDTO{
		user:        postgresql_entity.User{ID: 10},
		authIpError: errors.New("ip is blacklisted"),
	}
	resp := dataComparingStep(ctx, dto)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "ip is blacklisted", resp.Error)
	assert.Empty(t, resp.Token)
}

func TestDataComparingStep_ScoreZero(t *testing.T) {
	ctx, w := createTestGinContext("")
	dto := &authDTO{
		user: postgresql_entity.User{ID: 10},
		redisIp: redismodels.AuthenticationIP{
			Score:     0,
			CreatedAt: time.Now(),
		},
	}
	resp := dataComparingStep(ctx, dto)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "Authentication failed due to score", resp.Error)
	assert.Empty(t, resp.Token)
}

func TestDataComparingStep_MultipleDeviceError(t *testing.T) {
	ctx, w := createTestGinContext("")
	dto := &authDTO{
		user:                postgresql_entity.User{ID: 10},
		multipleDeviceError: errors.New("different device detected: last device was 'phone', current is 'laptop'"),
	}
	resp := dataComparingStep(ctx, dto)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, resp.Error, "different device detected")
	assert.Empty(t, resp.Token)
}

func TestDataComparingStep_Success_GeneratesJWT(t *testing.T) {
	ctx, w := createTestGinContext("device-1")
	email, err := vo.NewEmail("walter@example.com")
	require.NoError(t, err)

	dto := &authDTO{
		user: postgresql_entity.User{
			ID:    42,
			Email: email,
		},
		redisIp: redismodels.AuthenticationIP{
			Score:     100,
			CreatedAt: time.Now(),
		},
		authenticationGroup: postgresql_entity.AuthenticationAccountGroups{
			Name: "admin",
		},
	}

	resp := dataComparingStep(ctx, dto)

	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Empty(t, resp.Error)
	assert.Equal(t, "authentication successful", resp.Message)
	assert.NotEmpty(t, resp.Token)

	// Validar que o token gerado é um JWT válido com as claims corretas
	claims, err := auth.ValidateToken(resp.Token)
	require.NoError(t, err)
	assert.Equal(t, uint(42), claims.UserID)
	assert.Equal(t, "walter@example.com", claims.Email)
	assert.Equal(t, "admin", claims.Role)

	// Validar que a resposta JSON no context bate com o DTO
	var jsonResp authResponseDTO
	err = json.Unmarshal(w.Body.Bytes(), &jsonResp)
	require.NoError(t, err)
	assert.Equal(t, resp.Token, jsonResp.Token)
	assert.Equal(t, resp.Message, jsonResp.Message)
}
