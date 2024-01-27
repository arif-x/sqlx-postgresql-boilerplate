package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/auth"
	repo "github.com/arif-x/sqlx-gofiber-boilerplate/app/repository/auth"
	"github.com/arif-x/sqlx-gofiber-boilerplate/config"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	hash "github.com/arif-x/sqlx-gofiber-boilerplate/pkg/hash"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/response"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	JWTTokenAuthed "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"golang.org/x/crypto/bcrypt"
)

// Register method for new user registration.
// @Description new user registration.
// @Summary new user registration.
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name"
// @Param username formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password" format(password)
// @Failure 400,401,403,500 {object} response.ErrorResponse "Error"
// @Success 200 {object} response.AuthResponse
// @Router /api/v1/auth/register [post]
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
	user, permission, err := repository.Register(register)

	if err != nil {
		return response.InternalServerError(c, err)
	}

	token, err := GenerateNewAccessToken(user.UUID, user.Username, user.Email, user.Name, user.IsActive, user.EmailVerifiedAt, user.RoleUUID, permission)
	if err != nil {
		return response.InternalServerError(c, errors.New("Internal Error"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Token will be expired within %d minutes", config.AppCfg().JWTSecretExpireMinutesCount),
		"data":    token,
	})
}

// Register method for user login.
// @Description user login.
// @Summary user login.
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username Or Email"
// @Param password formData string true "Password" format(password)
// @Failure 400,401,403,500 {object} response.ErrorResponse "Error"
// @Success 200 {object} response.AuthResponse
// @Router /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	login := &model.Login{}

	if err := c.BodyParser(login); err != nil {
		return response.BadRequest(c, err)
	}

	repository := repo.NewAuthRepo(database.GetDB())
	user, permission, err := repository.Login(login.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.InvalidCredential(c, errors.New("No credential"))
		} else {
			return response.InvalidCredential(c, errors.New("No credential"))
		}
	}

	isValid := IsValidPassword([]byte(user.Password), []byte(login.Password))
	if !isValid {
		return response.InvalidCredential(c, errors.New("Incorrect password"))
	}

	token, err := GenerateNewAccessToken(user.UUID, user.Username, user.Email, user.Name, user.IsActive, user.EmailVerifiedAt, user.RoleUUID, permission)
	if err != nil {
		return response.InternalServerError(c, errors.New("Internal Error"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Token will be expired within %d minutes", config.AppCfg().JWTSecretExpireMinutesCount),
		"data":    token,
	})
}

// Register method for user send verificationemail.
// @Description user send verificationemail.
// @Summary user send verificationemail.
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Failure 400,401,403,500 {object} response.ErrorResponse "Error"
// @Success 200 {object} response.AuthResponse
// @Router /api/v1/auth/send-email [post]
func SendEmail(c *fiber.Ctx) error {
	user := c.Locals("user").(*JWTTokenAuthed.Token)
	claims := user.Claims.(JWTTokenAuthed.MapClaims)
	email_token := claims["email_token"]
	user_email := claims["email"]
	email_verified_at := claims["email_verified_at"]

	if email_verified_at != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Your email has already verified!",
			"data":    nil,
		})
	}

	email, ok := user_email.(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to send email!",
			"data":    nil,
		})
	}

	token, ok := email_token.(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to send email!",
			"data":    nil,
		})
	}

	err := SendVerificationEmail(email, token)
	if err != nil {
		log.Println("Error sending verification email:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to send email!",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Email has been sent!",
		"data":    "OK",
	})
}

// Register method for user verify email.
// @Description user verify email.
// @Summary user verify email.
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param token path string true "Email Token" default(f72cb686-2fc3-4147-8183-f93684780765)
// @Security ApiKeyAuth
// @Failure 400,401,403,500 {object} response.ErrorResponse "Error"
// @Success 200 {object} response.AuthResponse
// @Router /api/v1/auth/verify/{token} [get]
func Verify(c *fiber.Ctx) error {
	token := c.Params("token")

	user := c.Locals("user").(*JWTTokenAuthed.Token)
	claims := user.Claims.(JWTTokenAuthed.MapClaims)
	username := claims["username"]
	email_token := claims["email_token"]

	if token != email_token {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid email code!",
			"data":    "OK",
		})
	}

	repository := repo.NewAuthRepo(database.GetDB())
	user_data, permission, err := repository.Verify(username.(string))

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  false,
			"message": "Can't verify your email!",
			"data":    "OK",
		})
	}

	token_data, err := GenerateNewAccessToken(user_data.UUID, user_data.Username, user_data.Email, user_data.Name, user_data.IsActive, user_data.EmailVerifiedAt, user_data.RoleUUID, permission)
	if err != nil {
		return response.InternalServerError(c, errors.New("Internal Error"))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Token has been regenerated and will be expired within %d minutes", config.AppCfg().JWTSecretExpireMinutesCount),
		"data":    token_data,
	})
}

func GenerateNewAccessToken(UserID uuid.UUID, Username string, Email string, Name string, IsActive bool, EmailVerifiedAt *time.Time, RoleUUID string, Permission []string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = UserID.String()
	claims["username"] = Username
	claims["email"] = Email
	claims["name"] = Name
	claims["is_active"] = IsActive
	claims["email_verified_at"] = EmailVerifiedAt
	claims["email_token"] = uuid.New().String()
	claims["role_uuid"] = RoleUUID
	claims["permission"] = Permission
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AppCfg().JWTSecretExpireMinutesCount)).Unix()

	t, err := token.SignedString([]byte(config.AppCfg().JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func SendVerificationEmail(emailAddress, token string) error {
	emailConfig := &email.Email{
		From:    "sudo.ariffudin@gmail.com",
		To:      []string{emailAddress},
		Subject: "Email Verification",
		Text:    []byte(fmt.Sprintf("Click the following link to verify your email: http://localhost:8080/verify/%s", token)),
	}

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_ADDRESS"))
	return emailConfig.Send(os.Getenv("SMTP_ADDRESS")+":"+fmt.Sprint(os.Getenv("SMTP_PORT")), auth)
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
