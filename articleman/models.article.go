// models.article.go

package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//"strings"
	"time"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
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
	db := client.Database(dbConfig.dbName).Collection(dbConfig.articleColName)

	cursor, err := db.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var articleList []article
	if err := cursor.All(ctx, &articleList); err != nil {
		log.Fatal(err)
	}
	fmt.Println(articleList)
	return articleList
}

// Fetch an article based on the ID supplied
func getArticleByID(id int) (*article, error) {
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
		var articleOne article
		if err := cursor.Decode(&articleOne); err != nil {
			log.Fatal(err)
		}
		if articleOne.ID == id {
			return &articleOne, nil
		}
	}
	return nil, errors.New("Article Can't be found")
}

// Create a new article with the title and content provided
func createNewArticle(title, content string) (*article, error) {
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
	db := client.Database(dbConfig.dbName).Collection(dbConfig.articleColName)
	// Set the ID of a new article to one more than the number of articles
	a := article{ID: len(getAllArticles()) + 1, Title: title, Content: content}

	// Add the article to the list of articles
	_, err := db.InsertOne(ctx, a)
	if err != nil {
		log.Fatal(err)
	}

	return &a, nil
}
