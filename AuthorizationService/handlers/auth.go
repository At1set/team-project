package handlers

import (
	"authorization-service/models"
	"authorization-service/storage"
	"authorization-service/utilts"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
	storage storage.UserStorage
}

func NewAuthHandler(storage storage.UserStorage) *AuthHandler {
	return &AuthHandler{storage: storage}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {

	fmt.Println("Меня вызывли!")
	id, ok := utilts.GetIdFromJWT(c)

	if !ok {
		return errors.New("Что-то не так с токеном")
	}
	_, exists := h.storage.FindUserById(id)
	if exists {
		return errors.New("Такой пользователь уже существует!")
	}
	var regReq models.RegisterRequest

	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	if _, exists := h.storage.FindUserByLogin(regReq.Login); exists {
		return errors.New("пользователь с таким email уже зарегистрирован")
	}

	h.storage.CreateUser(models.User{
		Id:       int(uuid.New().ID()),
		Login:    regReq.Login,
		Name:     regReq.Name,
		Password: regReq.Password,
		Role:     "user",
	})
	c.SendStatus(fiber.StatusCreated)
	return nil
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var logReq models.LoginRequest

	if err := c.BodyParser(&logReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	user, exists := h.storage.FindUserByLogin(logReq.Login)
	if !exists || user.Password != logReq.Password {
		return errors.New("email or password is incorrect")
	}

	token, err := utilts.GenerateJWT(user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(models.LoginResponse{AccessToken: token})
}

func (h *AuthHandler) GetUsers(c *fiber.Ctx) error {
	s, ok := h.storage.GetAllUsersJSON()
	if !ok {
		fmt.Println("Не удалось всех спарсить!")
	}
	c.Set("Content-Type", "application/json")
	return c.SendString(s)
}

func (h *AuthHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Print(id)
	user, ok := h.storage.FindUserById(id)

	if !ok {
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return errors.New("Ошибка при преобразовании в JSON:")
	}
	c.Set("Content-Type", "application/json")
	return c.SendString(string(jsonData))
}

func (h *AuthHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	fmt.Println(c.Params("id"))
	if err != nil {
		fmt.Println(err)
	}
	h.storage.DeleteUser(id)
	fmt.Println(id)
	return c.SendStatus(fiber.StatusAccepted)
}
