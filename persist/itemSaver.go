package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver got item %d : %v", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item : %+v: %v", item, err)
			}
		}
	}()

	return nil
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	response, err := client.Index().Index("dating_profile").Type("zhenai").
		BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}

	return response.Id, nil
}