package services

import (
	"context"
	"test/models/db"
	"test/models/entity"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection *mongo.Collection

// Create User with route POST /user
func Create(c *gin.Context) {
	if userCollection == nil {
		userCollection = db.GetUserCollection()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userInfo entity.User
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	id, err := userCollection.InsertOne(ctx, userInfo)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

// GetUser with route GET /user/:id
func GetUser(c *gin.Context) {
	if userCollection == nil {
		userCollection = db.GetUserCollection()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userInfo entity.User
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	err := userCollection.FindOne(ctx, entity.User{ID: id}).Decode(&userInfo)
	if err != nil {
		c.AbortWithStatusJSON(501, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"User": &userInfo})
}

// GetAllUser with route GET /user
func GetAllUser(c *gin.Context) {
	if userCollection == nil {
		userCollection = db.GetUserCollection()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userInfo []entity.User

	result, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		c.AbortWithStatusJSON(501, gin.H{"Error": err.Error()})
		return
	}
	defer result.Close(ctx)

	if err := result.All(ctx, &userInfo); err != nil {
		c.AbortWithStatusJSON(502, gin.H{"Error": err.Error()})
		return
	}

	for result.Next(ctx) {
		var user entity.User
		if err = result.Decode(&user); err != nil {
			c.AbortWithStatusJSON(503, gin.H{"Error": err.Error()})
		}
		userInfo = append(userInfo, user)
	}
	c.JSON(200, gin.H{"Result": &userInfo})
}

//UpdateUser with route PUT /user/:id
func UpdateUser(c *gin.Context) {
	if userCollection == nil {
		userCollection = db.GetUserCollection()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var userInfo entity.User
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	result, err := userCollection.UpdateOne(ctx, entity.User{ID: id}, bson.M{"$set": userInfo})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Result": result})
}

//DeleteUser with route DELETE /user
func DeleteUser(c *gin.Context) {
	if userCollection == nil {
		userCollection = db.GetUserCollection()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	result, err := userCollection.DeleteOne(ctx, entity.User{ID: id})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Result": result})
}
