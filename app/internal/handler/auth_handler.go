package handlers

import (
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary
// @Description  smoke test
// @Tags         test
// @Produce      json
// @Success      200 {object} string
// @Failure      401 {object} string
// @Router       /api/v1/Authentication [post]
func Authentication(c *gin.Context) {
	// var body handler_model.Authentication

	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// // cache := cache.NewRedisCache()
	// // authIp, err := cache.Get(c, c.ClientIP())

	// db := db.GetDbConnection()
	// username, err := vo.NewUsername(body.Username)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid username or password",
	// 	})
	// 	return
	// }
	// password, err := vo.NewPassword(body.Password)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid username or password",
	// 	})
	// 	return
	// }

	// cache := cache.NewRedisCache()

	// var existingUser postgresql_entity.User

	// validation := db.Where(&postgresql_entity.User{
	// 	Username: username,
	// 	Password: password,
	// }).First(&existingUser).Error

}
func AuthenticateIp(c *cache.RedisCache, ctx *gin.Context) {
	authIp, err := c.Get(
		ctx,
		ctx.ClientIP(),
	)

}
