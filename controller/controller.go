package controller

import (
	"context"
	"gomongo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func GetAllDataHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := GetAllData(context.TODO(), collection)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})

	}
}

func GetAllData(ctx context.Context, collection *mongo.Collection) ([]models.DataModels, error) {
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []models.DataModels
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
func GetOneDataHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		data, err := GetOneData(context.TODO(), collection, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func GetOneData(ctx context.Context, collection *mongo.Collection, id string) (*models.DataModels, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var result models.DataModels
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func InsertData(collection *mongo.Collection, ctx context.Context, data *models.DataModels) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
	}

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func InsertDataHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.DataModels
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := InsertData(collection, context.TODO(), &data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})

	}
}

func DeleteDataByIdHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := DeleteData(context.TODO(), collection, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"messageg": "data is deleted"})
	}
}

func DeleteData(ctx context.Context, collection *mongo.Collection, id string) error {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": ObjectID}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDataHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var newData models.DataModels
		if err := ctx.ShouldBindJSON(&newData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		UpdatedData, err := UpdateData(collection, context.TODO(), id, &newData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"updated": UpdatedData})
	}
}

func UpdateData(collection *mongo.Collection, ctx context.Context, id string, newData *models.DataModels) (*models.DataModels, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": bson.M{
		"name":       newData.Name,
		"age":        newData.Age,
		"salary":     newData.Salary,
		"department": newData.Department,
	}}
	var updatedData models.DataModels
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedData)
	if err != nil {
		return nil, err
	}

	return &updatedData, nil
}






////////from here postgres start////////








func GetFromPostgresHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := GetFromPostgres(db)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func GetFromPostgres(db *gorm.DB) ([]models.MongoPostgres, error) {
	var result []models.MongoPostgres
	sql := "SELECT * FROM initpractice.postgresmodel"
	err := db.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func PostInPostgresHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newData models.MongoPostgres
		if err := ctx.ShouldBindJSON(&newData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := InsertIntoPostgres(db, newData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})

	}
}

func InsertIntoPostgres(db *gorm.DB, newData models.MongoPostgres) error {

	sql := "INSERT INTO initpractice.postgresmodel (name, age, salary, department , created_at) VALUES (?, ?, ?, ?, ?)"
	result := db.Exec(sql, newData.Name, newData.Age, newData.Salary, newData.Department, time.Now())
	if result.Error != nil {
		return result.Error
	}

	return nil
}
