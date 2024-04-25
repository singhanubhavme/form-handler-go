package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/singhanubhavme/form-handler-go/configs"
	"github.com/singhanubhavme/form-handler-go/models"
	"github.com/singhanubhavme/form-handler-go/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

var validate = validator.New()

// user
func LoginUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User

	defer cancel()
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GeneralResponses{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	err := configs.GetMongoDbClient().Collection("users").FindOne(ctx, bson.M{
		"email": user.Email, "password": user.Password,
	}).Decode(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(responses.GeneralResponses{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"token": t, "email": user.Email}})

}

// user
func RegisterUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GeneralResponses{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GeneralResponses{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	result, err := configs.GetMongoDbClient().Collection("users").InsertOne(ctx, newUser)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.GeneralResponses{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
