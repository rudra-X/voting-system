package services

import (
	"log"
	"net/http"
	"os"
	"time"
	"voiting-system/internal/dtos"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Login(c *gin.Context) {
	var loginRequestDto dtos.LoginRequestDto

	if err := c.BindJSON(&loginRequestDto); err != nil {
		//TODO::throw exception 400 as incoming request body was not able to map
		return
	}

	storedUser, exists := Users[loginRequestDto.Email]

	if !exists {
		c.IndentedJSON(http.StatusNotFound, dtos.ErrorResponse{
			BaseResponseDto: dtos.BaseResponseDto{
				Message: "User not found, please register",
			},
		})
		return
	}

	if !storedUser.ComparePassword(loginRequestDto.Password) {
		c.IndentedJSON(http.StatusUnauthorized, dtos.ErrorResponse{
			BaseResponseDto: dtos.BaseResponseDto{
				Message: "Invalid email or password",
			},
			ErrorCode: 401,
		})
		return
	}

	log.Printf("User %s has successfully logged in", loginRequestDto.Email)

	token, err := createJWT(storedUser.UserId)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, dtos.LoginResponseDto{
		Token: token,
	})
}

func createJWT(userId string) (string, error) {

	claims := &dtos.JwtClaimsDto{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
