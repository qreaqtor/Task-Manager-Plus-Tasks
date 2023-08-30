package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	getCurrentUserURL string
	client            *http.Client
)

func init() {
	authServiceURL := os.Getenv("Auth_service_adress")
	getCurrentUserURL = authServiceURL + "users/get/me"
	client = &http.Client{}
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")
		req, err := http.NewRequest("GET", getCurrentUserURL, nil)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		req.Header.Add("Authorization", bearerToken)

		resp, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		defer resp.Body.Close()

		respUserId := struct {
			UserId primitive.ObjectID `json:"id"`
		}{
			UserId: primitive.NilObjectID,
		}

		err = json.NewDecoder(resp.Body).Decode(&respUserId)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userId", respUserId.UserId)
		ctx.Next()
	}
}
