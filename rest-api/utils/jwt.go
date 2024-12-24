package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Claims struct {
	Foo int64 `json:"foo"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId int64) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	fmt.Println(ss)

	return ss, nil
}

func VerifyJWT(c *gin.Context) {
	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")

	// Check if the header is present and properly formatted
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided or invalid"})
		c.Abort()
		return
	}

	// Extract the token by trimming the "Bearer " prefix
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//	Proceed with the request
		c.Set("foo", claims["foo"])
		c.Next()
	}
}
