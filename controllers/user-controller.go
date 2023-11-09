package controllers

import (
	"net/http"
	"test/common"
	"test/db"
	"test/entity"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(ctx *gin.Context) {
	var loginCredentials struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := ctx.ShouldBind(&loginCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCollection := db.GetDBCollection("users")

	var user entity.User

	if err := userCollection.FindOne(ctx, gin.H{"email": loginCredentials.Email}).Decode(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding user"})
		return
	}

	if !common.CheckPasswordHash(loginCredentials.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := common.GenerateToken(loginCredentials.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func Register(ctx *gin.Context) {

	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := common.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	user.Password = hashedPassword

	userCollection := db.GetDBCollection("users")

	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A user with this email already exists."})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting the user into the database."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User inserted successfully",
		"user":    result,
	})
}
