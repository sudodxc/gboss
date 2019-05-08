package persist

import (
	//"encoding/json"
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Itrm saver:got item"+"%d:%v", itemCount, item)
			itemCount++
			// json.Unmarshal(data, v)
			_, err := save(item)
			if err != nil {
				log.Printf("Item Error:Error"+"Saving item %v:%v", item, err)
			}

		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		//must run at docker
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().
		Index("data").
		Type("admin").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
