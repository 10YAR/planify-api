package controllers

import (
	"api/database"
	"api/repositories"
	"api/types"
	"api/utils"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	auth := new(types.Auth)

	err := c.BodyParser(auth)
	if err != nil {
		return c.JSON(utils.E400("Bad Request", err))
	}

	// Valide les données envoyées
	errors := utils.ValidateStruct(*auth)
	if errors != "" {
		return c.JSON(utils.E400("Bad request : "+errors, nil))
	}

	// Vérification de l'utilisateur
	//res, err := database.DoQuery("SELECT * FROM users WHERE email = ?", auth.Email)
	//if err != nil {
	//	return c.JSON(utils.E503("Internal Server Error", err))
	//}
	//
	//user := new(types.User)
	//for res.Next() {
	//	err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	//	if err != nil {
	//		return c.JSON(utils.E503("Internal Server Error", err))
	//	}
	//
	//	if !utils.CheckPasswordHash(auth.Password, user.Password) {
	//		return c.JSON(utils.E401("Unauthorized", nil))
	//	}
	//}

	db := utils.GetLocal[*sql.DB](c, "db")
	user, err := repositories.Login(db, auth.Email)

	if !utils.CheckPasswordHash(auth.Password, user.Password) {
		return c.JSON(utils.E401("Unauthorized", nil))
	}

	if user.ID == 0 {
		return c.JSON(utils.E401("Unauthorized", nil))
	}

	// Données du token
	claims := jwt.MapClaims{
		"id":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}

	// Génère un token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.JSON(utils.E503("Internal Server Error", err))
	}

	successMessage := fmt.Sprintf("Bienvenue %s %s vous êtes maintenant connecté", user.FirstName, user.LastName)
	return c.JSON(types.HttpResponse{Status: 1, Message: successMessage, HttpCode: 200, Token: jwtToken})
}

func Register(c *fiber.Ctx) error {
	user := new(types.User)
	db := utils.GetLocal[*sql.DB](c, "db")

	err := c.BodyParser(user)
	if err != nil {
		return c.JSON(utils.E400("Bad Request", err))
	}

	// Valide les données envoyées
	errors := utils.ValidateStruct(*user)
	if errors != "" {
		return c.JSON(utils.E400("Bad request : "+errors, nil))
	}

	// Vérification de l'utilisateur
	res, err := database.DoQuery(db, "SELECT * FROM users WHERE email = ?", user.Email)
	if err != nil {
		return c.JSON(utils.E503("Internal Server Error", err))
	}

	for res.Next() {
		return c.JSON(types.HttpResponse{Status: 0, Message: "Email already exists", HttpCode: 400})
	}

	user.Password, _ = utils.HashPassword(user.Password)

	// Insertion de l'utilisateur
	_, errDb := repositories.Register(db, user)
	if errDb != nil {
		return c.JSON(utils.E503("Internal Server Error", errDb))
	}

	return c.JSON(types.HttpResponse{Status: 1, Message: "User created successfully", HttpCode: 200})
}
