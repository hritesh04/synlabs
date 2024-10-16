package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/helper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Secret []byte
}

func New(secret string) *AuthService {
	return &AuthService{
		Secret: []byte(secret),
	}
}

func (a *AuthService) Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("syn")
		if err != nil {
			helper.ReturnFailed(ctx, http.StatusBadRequest, err)
			ctx.Abort()
			return
		}
		claims, err := a.Validate(tokenString)
		if err != nil {
			helper.ReturnFailed(ctx, http.StatusBadRequest, err)
			ctx.Abort()
			return
		}
		if userId, ok := claims["userID"].(string); ok {
			ctx.Request.Header.Set("userID", userId)
		} else {
			helper.ReturnFailed(ctx, http.StatusBadRequest, fmt.Errorf("invalid token: no user ID"))
			ctx.Abort()
			return
		}
		if role, ok := claims["role"].(string); ok {
			ctx.Request.Header.Set("role", role)
		}
		ctx.Next()
	}
}

func (a *AuthService) Validate(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return a.Secret, nil
	})
	if err != nil {
		return jwt.MapClaims{}, nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return jwt.MapClaims{}, fmt.Errorf("token is invalid")
}

func (a *AuthService) AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.Request.Header.Get("role")
		if role != domain.Admin {
			ctx.Next()
			return
		}
		helper.ReturnFailed(ctx, http.StatusUnauthorized, fmt.Errorf("user is not an admin"))
		ctx.Abort()
	}
}

func (a *AuthService) GenerateToken() {

}

func (a *AuthService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password : %v", err)
	}
	return string(hash), nil
}

func (a *AuthService) ComparePassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
