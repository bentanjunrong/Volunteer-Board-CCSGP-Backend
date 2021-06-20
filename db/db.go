package db

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
)

var db *elasticsearch.Client

func InitDB() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://volunteery-deployment.es.ap-southeast-1.aws.found.io:9243", // TODO: put in env file (https://github.com/aoyinke/lianjiaEngine/blob/f51e8a446349e054d5cd851d3e2f80b2857825d6/lianjia_svr.go#L11)
		},
		Username: "elastic", // TODO: put in env file
		Password: "eC0f0x3Tl1fYo8IKcrEmbwXR",
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
