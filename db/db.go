package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ES *elasticsearch.Client
var MongoDB *mongo.Client

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	if MongoDB, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URL"))); err != nil {
		panic(err)
	}
	// ping health check
	if err := MongoDB.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MongoDB.")
}

func GetCollection(collection string) *mongo.Collection {
	return (*MongoDB).Database("volunteery-db").Collection(collection)
}

func DisconnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := MongoDB.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func InitES() {
	cfg := elasticsearch.Config{
		Addresses: []string{os.Getenv("ES_URL")},
		Username:  os.Getenv("ES_USER"),
		Password:  os.Getenv("ES_PASSWORD"),
		Logger: &estransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		ES = es
		log.Println(es.Info())
	}
}

func Search(query string, index string) ([]map[string]interface{}, error) {
	byteQuery := []byte(fmt.Sprintf(`
		{
			"query": {
				"wildcard": {
					"name": {
						"value": "%s*"
					}
				}
			}
		}
	`, query))
	response, err := ES.Search(ES.Search.WithIndex(index), ES.Search.WithBody(bytes.NewReader(byteQuery)))
	if err != nil {
		log.Fatalf("Error searching for all entries in index %s.", index)
		return nil, err
	}
	var result map[string]map[string][]map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	allMatches := result["hits"]["hits"]
	if len(allMatches) == 0 {
		return nil, errors.New("No entries found.")
	}
	return allMatches, nil
}
