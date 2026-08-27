package service 

import (
	"github.com/golang-jwt/jwt/v5",
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/internal/config",
)


func setJwtCookies(sub string   ) err {
	claims := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	conf := config.NewEnvironmentConfig()
	
	signed, err := token.SignedString([]byte(conf.JwtSecretKey))

	if err != nil {
		http.Error(w, "could not create token", http.StatusInternalServerError)
		return err 
	}

	tokenExpiration := strconv.ParseInt(conf.JwtTokenExpiration, 10, 64)  
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    signed,
		Expires:  time.Now().Add(tokenExpiration * time.Hour),
		HttpOnly: true,                    // JS can't read it — mitigates XSS token theft
		Secure:   true,                    // only sent over HTTPS
		SameSite: http.SameSiteStrictMode, // mitigates CSRF
		Path:     "/",
	})
}