package main

import (
	"context"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type db struct {
	dburi          string
	dbName         string
	usersColName   string
	articleColName string
}

var dbConfig db

//InitConfig is used to initliaze some constants
func InitConfig() {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		os.Exit(1)
	}
	err := viper.Unmarshal(&dbConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}

func pingDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConfig.dburi))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
}
