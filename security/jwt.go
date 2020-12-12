package security

import (
	"footcer-backend/model"
	"footcer-backend/security/pro"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		TokenNotify:   user.TokenNotify,
		UserName: user.DisplayName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3600).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(pro.JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
