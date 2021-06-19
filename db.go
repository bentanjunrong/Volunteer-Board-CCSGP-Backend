package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

var db *elasticsearch.Client

func InitDB() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://volunteery-deployment.es.ap-southeast-1.aws.found.io:9243",
		},
		Username: "elastic",
		Password: "nicetry",
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
