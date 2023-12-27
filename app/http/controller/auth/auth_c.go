package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/auth"
	repo "github.com/arif-x/sqlx-gofiber-boilerplate/app/repository/auth"
	"github.com/arif-x/sqlx-gofiber-boilerplate/config"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	hash "github.com/arif-x/sqlx-gofiber-boilerplate/pkg/hash"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/response"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	register := &model.Register{}

	if err := c.BodyParser(register); err != nil {
		return response.BadRequest(c, err)
	}

	password, err := hash.Hash([]byte("password"))
	if err != nil {
		return response.InternalServerError(c, err)
	}

	register.Password = password

	repository := repo.NewAuthRepo(database.GetDB())
	user, err := repository.Register(register)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	token, err := GenerateNewAccessToken(user.UUID, user.Username, user.Email, user.Name)
	if err != nil {
		return response.InternalServerError(c, errors.New("Internal Error"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Token will be expired within %d minutes", config.AppCfg().JWTSecretExpireMinutesCount),
		"data":    token,
	})
}

func Login(c *fiber.Ctx) error {
	login := &model.Login{}

	if err := c.BodyParser(login); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewAuthRepo(database.GetDB())
	user, err := repository.Login(login.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.NotFound(c, err)
		} else {
			return response.InternalServerError(c, errors.New("No credential"))
		}
	}

	isValid := IsValidPassword([]byte(user.Password), []byte(login.Password))
	if !isValid {
		return response.InternalServerError(c, errors.New("Incorrect password"))
	}

	token, err := GenerateNewAccessToken(user.UUID, user.Username, user.Email, user.Name)
	if err != nil {
		return response.InternalServerError(c, errors.New("Internal Error"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Token will be expired within %d minutes", config.AppCfg().JWTSecretExpireMinutesCount),
		"data":    token,
	})

}

func GenerateNewAccessToken(UserID uuid.UUID, Username string, Email string, Name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = UserID.String()
	claims["username"] = Username
	claims["email"] = Email
	claims["name"] = Name
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AppCfg().JWTSecretExpireMinutesCount)).Unix()

	t, err := token.SignedString([]byte(config.AppCfg().JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GeneratePasswordHash(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func IsValidPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return false
	}

	return true
}
