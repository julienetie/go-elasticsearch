package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienetie/go-elasticsearch/app"

	elastic "github.com/olivere/elastic/v7"
)

func main() {

	ctx := context.Background()
	esclient, err := app.GetESClient()

	if err != nil {
		fmt.Println("Error initializing :", err)
		panic("Client fail")
	}

	var students []app.Student

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", "Doe"))

	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}

	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := esclient.Search().Index("students").SearchSource(searchSource)
	searchResult, err := searchService.Do(ctx)

	if err != nil {
		fmt.Println("[ProductsES][GetIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var student app.Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			fmt.Println("[Getting Strudents][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, s := range students {
			fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
		}
	}

}
