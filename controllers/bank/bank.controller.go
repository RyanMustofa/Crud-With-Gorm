package controllers

import (
	"strconv"
	"time"

	connection_db "com.server/luis/config/connection"
	"com.server/luis/config/helpers"
	model_bank "com.server/luis/model"
	"github.com/gofiber/fiber/v2"
)

func GetBank(c *fiber.Ctx) error {
		page := c.Query("page")
		_limit := c.Query("limit")

		var offset int = 0
		var limit int = 10
		if _limit != "" {
			limit, _ = strconv.Atoi(_limit)
		}
		if page != "" {
			offset, _ = strconv.Atoi(page)
			offset = (offset - 1) * limit
		}
		var banks []model_bank.Bank
		results := connection_db.DB.Table("banks").Limit(limit).Offset(offset).Find(&banks)

		if results.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
				"message": "failed",
				"data": results.Error.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": true,
			"message": "Hallo im luis",
			"data": banks,
		})
}

func FindBank(c *fiber.Ctx) error {
	bankid := c.Params("bank_id")

	var bank model_bank.Bank
	results := connection_db.DB.Table("banks").Where(map[string]interface{}{"id": bankid}).First(&bank)
	if results.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": results.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": bank,
	})
}

func CreateBank(c *fiber.Ctx) error {
	var bank_payload model_bank.CreateBank

	if err := c.BodyParser(&bank_payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error": err.Error(),
		})
	}

	err := helpers.Validation(bank_payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error": err,
		})
	}

	payloadAdd := model_bank.Bank{
		Name: bank_payload.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	results := connection_db.DB.Table("banks").Create(&payloadAdd)

	if results.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error": results.Error.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": bank_payload,
	})
}

func UpdateBank(c *fiber.Ctx) error {
	bankid := c.Params("bank_id")
	var id int 
	if bankid != "" {
		id, _ = strconv.Atoi(bankid)
	}

	var payload_bank model_bank.CreateBank
	if err := c.BodyParser(&payload_bank); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"error": err.Error(),
		})
	}

	var bank model_bank.Bank

	err := connection_db.DB.Table("banks").Where(&model_bank.Bank{ID: id }).First(&bank)
	
	if err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"error": err.Error.Error(),
		})
	}


	bank.Name = payload_bank.Name
	bank.UpdatedAt = time.Now()

	results := connection_db.DB.Table("banks").Save(&bank)

	if results.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error": results.Error.Error(),
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": bank,
	})
}

func DeleteBank(c *fiber.Ctx) error {
	bankid := c.Params("bank_id")
	var id int
	if bankid != "" {
		id, _ = strconv.Atoi(bankid)
	}
	var bank model_bank.Bank

	results := connection_db.DB.Where(model_bank.Bank{
		ID: id,
	}).Delete(&bank)
	if results.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"error": results.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": id,
	})
}
