package middleware

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAdminMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("ключ")},
		ContextKey: "user",
		SuccessHandler: func(c *fiber.Ctx) error {
			fmt.Println("Что0то")
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			if role, ok := claims["role"].(string); !ok || role != "admin" {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "Access denied: admin role required",
				})
			}
			return c.Next()
		},
	})
}

// fmt.Println("Меня вызывли!")
// // _id, ok := utilts.GetLoginFromJWT(c, "sub")
// // id, _ := strconv.Atoi(_id)
// id := 4
// fmt.Println(id)
// // if !ok {
// // 	return c.SendStatus(fiber.StatusUnauthorized)
// // }
// role, ok := utilts.GetLoginFromJWT(c, "role")
// fmt.Println(role)
// fmt.Println("Меня вызывли!")
// if !ok {
// 	return c.SendStatus(fiber.StatusUnauthorized)
// }

// if role != "admin" {
// 	fmt.Println("Не админ - не лезь")
// 	return c.SendStatus(fiber.StatusBadRequest)
// }

// _, exists := h.storage.FindUserById(strconv.Itoa(id))
// if !exists {
// 	return c.SendStatus(fiber.StatusUnauthorized)
// }
