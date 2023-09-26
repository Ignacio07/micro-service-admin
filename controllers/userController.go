package controllers

import (
	db "github.com/Ignacio07/micro-service-admin/config"
	"github.com/Ignacio07/micro-service-admin/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Select("id, store").Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var user models.User
	db.DB.Select("id, first_name, last_name, email").Where("id = ?", userId).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request", "error": err})
	}

	db.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var user models.User
	db.DB.Select("id, store, password").Where("id = ?", userId).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error", "error": err})
	}

	db.DB.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("taskId")
	var user models.User
	db.DB.Select("id, store, password").Where("id = ?", userId).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found"})
	}

	db.DB.Delete(&user) //Change delete field to true and assign deleted_at
	//db.DB.Unscoped().Delete(&task) //Unscoped() is used to delete permanently

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
	})
}
