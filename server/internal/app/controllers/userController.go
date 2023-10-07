package controllers

import (
	"os"

	random "github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joomcold/go-next-docker/internal/app/models"
	"github.com/joomcold/go-next-docker/internal/initializers"
	"golang.org/x/crypto/bcrypt"
)

func validateToken(c *fiber.Ctx) (*jwt.RegisteredClaims, error) {
	cookie := c.Cookies("jwt_token")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	// Valid
	return token.Claims.(*jwt.RegisteredClaims), nil
}

func findUserByToken(c *fiber.Ctx, user *models.User) error {
	claims, err := validateToken(c)
	if err != nil {
		return err
	}

	initializers.DB.Find(user, "id = ?", claims.Issuer)

	return nil
}

func Register(c *fiber.Ctx) error {
	type formData struct {
		Email                string `json:"email"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"passwordConfirmation"`
	}

	var form formData

	// Parse body into form
	err := c.BodyParser(&form)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Invalid input",
		})
	}

	// Check password matching
	if form.Password != form.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Password and Password Confirmation do not match",
		})
	}

	// Encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Failed to hash password",
		})
	}

	// Create user
	user := models.User{
		ID:       uuid.New(),
		Name:     random.Name(),
		Email:    form.Email,
		Password: string(hashedPassword),
		Address:  random.Address().Address,
	}

	err = initializers.DB.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": "Failed to register", "data": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "successful", "message": "Registration successfully", "data": user,
	})
}

func Profile(c *fiber.Ctx) error {
	var user models.User

	err := findUserByToken(c, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	type formData struct {
		Email    string
		Name     string
		Password string
	}

	var (
		user models.User
		form formData
	)

	// Parse body into form
	err := c.BodyParser(&form)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Invalid input",
		})
	}

	err = findUserByToken(c, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	// Check user email
	if user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error", "message": "User not found",
		})
	}

	// Detect changed password
	if form.Password != "" {
		newPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error", "message": "Failed to hash password",
			})
		}

		form.Password = string(newPassword)
	}

	// Update user info
	initializers.DB.Model(&user).Updates(form)

	return c.JSON(user)
}

func CancelUser(c *fiber.Ctx) error {
	var user models.User

	err := findUserByToken(c, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error", "message": err.Error(),
		})
	}

	initializers.DB.Delete(&user)

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status": "success", "message": "Cancellation successfully",
	})
}
