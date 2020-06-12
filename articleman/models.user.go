// models.user.go

package main

import (
	"context"
	"errors"
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
// var userList = []user{
// 	{Username: "user1", Password: "pass1"},
// 	{Username: "user2", Password: "pass2"},
// 	{Username: "user3", Password: "pass3"},
// }

// Check if the username and password combination is valid
func isUserValid(username, password string) bool {
	client, clierr := mongo.NewClient(options.Client().ApplyURI(dbConfig.dburi))
	if clierr != nil {
		log.Fatal(clierr)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	connerr := client.Connect(ctx)
	if connerr != nil {
		log.Fatal(connerr)
	}
	defer client.Disconnect(ctx)
	db := client.Database(dbConfig.dbName).Collection(dbConfig.usersColName)

	cursor, err := db.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var userOne user
		if err := cursor.Decode(&userOne); err != nil {
			log.Fatal(err)
		}
		if userOne.Username == username && userOne.Password == password {
			return true
		}
	}
	return false
}

// Register a new user with the given username and password
// NOTE: For this demo, we
func registerNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := user{Username: username, Password: password}

	client, err := mongo.NewClient(options.Client().ApplyURI(dbConfig.dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database(dbConfig.dbName).Collection(dbConfig.usersColName)
	_, err = db.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	client, clierr := mongo.NewClient(options.Client().ApplyURI(dbConfig.dburi))
	if clierr != nil {
		log.Fatal(clierr)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	connerr := client.Connect(ctx)
	if connerr != nil {
		log.Fatal(connerr)
	}
	defer client.Disconnect(ctx)
	db := client.Database(dbConfig.dbName).Collection(dbConfig.usersColName)

	cursor, err := db.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var userOne user
		if err := cursor.Decode(&userOne); err != nil {
			log.Fatal(err)
		}
		if userOne.Username == username {
			return false
		}
	}
	return true
}
