package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/singhanubhavme/form-handler-go/configs"
	"github.com/singhanubhavme/form-handler-go/models"
	"github.com/singhanubhavme/form-handler-go/responses"
	"github.com/singhanubhavme/form-handler-go/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateForm(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var Form models.Form
	defer cancel()
	if err := c.BodyParser(&Form); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GeneralResponses{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	var formId string = utils.RandomString(16)

	newForm := models.Form{
		Email:       Form.Email,
		FormTitle:   Form.FormTitle,
		FormId:      formId,
		RedirectUrl: Form.RedirectUrl,
	}
	_, err := configs.GetMongoDbClient().Collection("forms").InsertOne(ctx, newForm)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{
		"$push": bson.M{
			"form_ids": formId,
		},
	}

	filter := bson.M{"email": Form.Email}

	_, err_user := configs.GetMongoDbClient().Collection("users").UpdateOne(ctx, filter, update)

	if err_user != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err_user.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.GeneralResponses{Status: http.StatusCreated, Message: "success"})
}
