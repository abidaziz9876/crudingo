package config

import (
	"context"
	"fmt"
	"gomongo/models"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var MongoDB *mongo.Client

func ConnectToMongoDB() {
	mongoURI := os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal("cannot connect to mongodb")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	MongoDB = client
}


func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	collection:=client.Database("Employee").Collection(collectionName)
	return collection
}

var PostGresDB *gorm.DB
func ConnecToPostresDB(){
	host :=os.Getenv("HOST")
	port:=os.Getenv("PORT")
	user:=os.Getenv("USER")
	dbname:=os.Getenv("DB_NAME")
	password:=os.Getenv("PASSWORD")

	connectionString:=fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	host, port, user, password, dbname, "disable")

	client,err:=gorm.Open(postgres.Open(connectionString),&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgresDB")
	PostGresDB = client

}


func SetupPostgresSchemaAndTables() {
	var PostgresModel models.MongoPostgres
	
	PostGresDB.Exec("CREATE SCHEMA IF NOT EXISTS initpractice")

	err := PostGresDB.Table("initpractice.postgresmodel").AutoMigrate(&PostgresModel)
	if err != nil {
		log.Fatal(err.Error())
	}
}
