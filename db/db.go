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

var db *elasticsearch.Client

func InitMongoDB() {
	uri := "mongodb://localhost:27017/volunteery-db"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MongoDB.")
}

func InitDB() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://volunteery-deployment.es.ap-southeast-1.aws.found.io:9243", // TODO: put in env file (https://github.com/aoyinke/lianjiaEngine/blob/f51e8a446349e054d5cd851d3e2f80b2857825d6/lianjia_svr.go#L11)
		},
		Username: "elastic", // TODO: put in env file
		Password: os.Getenv("ES_PASSWORD"),
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
		db = es
		log.Println(es.Info())
	}
}

func GetDB() *elasticsearch.Client {
	return db
}

// TODO: decide on return type. []map[string]interface{} is fine if we leave unmarshalling of id and body to the model.
func GetAllByField(index string, field map[string]string) ([]map[string]interface{}, error) {
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"term": field,
				},
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	response, err := db.Search(db.Search.WithIndex(index), db.Search.WithBody(&buffer))
	if err != nil {
		log.Fatalf("Error searching for %s in index %s.", index, field)
		return nil, err
	}
	var result map[string]map[string][]map[string]interface{} // absolutely disgusting. no btr way here?
	json.NewDecoder(response.Body).Decode(&result)
	allMatches := result["hits"]["hits"]
	if len(allMatches) == 0 {
		return nil, errors.New("No entries found.")
	}
	return allMatches, nil
}

func GetOneByField(index string, field map[string]string) (map[string]interface{}, error) {
	allMatches, err := GetAllByField(index, field)
	if err != nil {
		return nil, err
	}
	return allMatches[0], nil
}

func GetAll(index string) ([]map[string]interface{}, error) {
	byteQuery := []byte(`
		{
			"query": {
				"match_all": {}
			}
		}
	`)
	response, err := db.Search(db.Search.WithIndex(index), db.Search.WithBody(bytes.NewReader(byteQuery)))
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
	response, err := db.Search(db.Search.WithIndex(index), db.Search.WithBody(bytes.NewReader(byteQuery)))
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
