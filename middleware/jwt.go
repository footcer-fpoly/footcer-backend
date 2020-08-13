package middleware

import (
	"footcer-backend/model"
	"footcer-backend/security/pro"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(pro.JWT_KEY),
	}

	return middleware.JWTWithConfig(config)
}
