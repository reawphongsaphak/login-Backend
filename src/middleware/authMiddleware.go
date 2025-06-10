package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"main/src/database"
	"main/src/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RequireAuth(c *gin.Context) {
	// get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Decode / validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// find the user with token sub
		var user models.User
		client := database.ConnectDB()
		coll := client.Database("test01").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		email, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err := coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// attach to req
		c.Set("user", user)

		// continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}