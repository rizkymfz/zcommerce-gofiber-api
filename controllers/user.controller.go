package controllers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizkymfz/zcommerce-gofiber-api/database"
	"github.com/rizkymfz/zcommerce-gofiber-api/models"
	"github.com/rizkymfz/zcommerce-gofiber-api/utils"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []*models.User
	database.DB.Debug().Find(&users)

	return utils.SuccessResponse(c, users, "success")
}

func CreateUsers(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return utils.FailedResponse(c, nil, "failed to validate", []string{errValidate.Error()})
	}

	newUser := models.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "status internal server error",
		})
	}

	newUser.Password = hashPassword

	database.DB.Debug().Create(&newUser)

	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	// 	"message": "Success create users",
	// 	"data":    newUser,
	// })
	return utils.SuccessResponse(c, newUser, "success", fiber.StatusCreated)
}

func GetUserById(c *fiber.Ctx) error {
	var user []*models.User

	result := database.DB.Debug().First(&user, c.Params("id"))

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"user": user,
	// })
	return utils.SuccessResponse(c, user, "success")

}

func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Debug().Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
	})

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"message": "succes update user",
	// })
	return utils.SuccessResponse(c, user, "success")
}

func DeleteUser(c *fiber.Ctx) error {
	user := new(models.User)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Debug().Where("id = ?", id).Delete(&user)

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"message": "delete user successfully",
	// })
	return utils.SuccessResponse(c, user, "delete user successfully")
}
