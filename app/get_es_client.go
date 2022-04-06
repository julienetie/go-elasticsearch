package app

import (
	"fmt"
	elastic "github.com/olivere/elastic/v7"
)

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("https://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialised...")

	return client, err
}
