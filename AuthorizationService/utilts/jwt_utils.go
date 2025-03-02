package utilts

import (
	"authorization-service/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("ключ")

func GenerateJWT(user models.User) (string, error) {
	payload := jwt.MapClaims{
		"sub":  strconv.Itoa(user.Id),
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"role": user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(jwtSecretKey)
}

func GetIdFromJWT(c *fiber.Ctx) (string, bool) {
	user := c.Locals("user").(*jwt.Token)
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}
	res, ok := claims["sub"].(string)
	return res, ok
}
