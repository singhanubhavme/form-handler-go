package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/singhanubhavme/form-handler-go/configs"
	"github.com/singhanubhavme/form-handler-go/helpers"
	"github.com/singhanubhavme/form-handler-go/models"
	"github.com/singhanubhavme/form-handler-go/responses"
	"github.com/singhanubhavme/form-handler-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func SubmitForm(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_id := c.Params("user")
	formId := c.Params("formid")

	bodyBytes := c.Body()

	var jsonData map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &jsonData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON body: " + err.Error(),
		})
	}

	objectId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusInternalServerError, Message: "error"})
	}
	var User models.User
	err_user := configs.GetMongoDbClient().Collection("users").FindOne(ctx, bson.M{
		"_id": objectId,
	}).Decode(&User)

	if err_user != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusInternalServerError, Message: "error"})
	}

	containsFormId := slices.Contains(User.FormIds, formId)

	if !containsFormId {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusInternalServerError, Message: "error"})
	}

	var Form models.Form
	err_form := configs.GetMongoDbClient().Collection("forms").FindOne(ctx, bson.M{
		"formid": formId,
	}).Decode(&Form)

	if err_form != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GeneralResponses{Status: http.StatusInternalServerError, Message: "error"})
	}

	dataString := helpers.MapToString(jsonData)

	htmlBody := "<div>" + dataString + "</div>"
	subject := "Form Submission for " + Form.Email + " on " + Form.FormId
	utils.SendMail(Form.Email, subject, htmlBody)

	return c.Redirect(Form.RedirectUrl)
}
