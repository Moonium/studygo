package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

var (
	client *elastic.Client
	ch     chan *LogData
)

// LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// Init ...
func Init(address string, chanSize, nums int) (err error) {
	if strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return err
	}
	fmt.Println("connect to es success")

	ch = make(chan *LogData, chanSize)

	for i := 0; i < nums; i++ {
		go SendToES()
	}

	return
}

// SendToESChan ...
func SendToESChan(msg *LogData) {
	ch <- msg

}

// SendToES ...
func SendToES() {
	for {
		select {
		case msg := <-ch:
			// fmt.Println(indexStr, data)
			// b, err := json.Marshal(data)
			// if err != nil {
			// 	// Handle error
			// 	return err
			// }
			// fmt.Println(b)
			put1, err := client.Index().
				Index(msg.Topic).
				Type("xxx").
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				// Handle error
				// return err
				fmt.Println(err)
				continue
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
			// return err
		default:
			time.Sleep(time.Second)
		}
	}

}
