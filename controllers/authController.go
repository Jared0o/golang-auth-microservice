package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jared0o/auth-microservice/initializers"
	"github.com/jared0o/auth-microservice/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	//get email and password
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "dupazbita",
		})
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "Failed to hash password",
		})
	}
	//create the user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "Failed to create the user",
		})
	}

	//respond
	return c.SendStatus(http.StatusCreated)
}

func Login(c *fiber.Ctx) error {
	//get email and password
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "dupazbita",
		})

	}
	//look up the user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "invalid email or password",
		})

	}
	//campare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "invalid email or password",
		})
	}
	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	fmt.Println(os.Getenv("SECRET"))
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "failed created token " + err.Error(),
		})

	}

	// send token
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})

}

func Validate(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user":      c.Get("user"),
		"czyplacki": "lubie placki",
	})
}
