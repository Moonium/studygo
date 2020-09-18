package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

// Student ...
type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func (s *Student) run() *Student {
	fmt.Printf("%s在跑...\n", s.Name)
	return s
}

func (s *Student) wang() {
	fmt.Printf("%s在汪汪汪...\n", s.Name)
}

func main() {
	// zhangsan := Student{
	// 	Name:    "zhangsan",
	// 	Age:     22,
	// 	Married: false,
	// }
	// zhangsan.run()
	// zhangsan.wang()

	// zhangsan.run().wang()

	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")

	p1 := Student{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().
		Index("student").
		Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
