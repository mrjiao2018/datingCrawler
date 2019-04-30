package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver got item %d : %v", itemCount, item)
			itemCount++
			err := save(index, client, item)
			if err != nil {
				log.Printf("Item Saver: error saving item : %+v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func save(index string, client *elastic.Client, item engine.Item) error {

	if item.Type == "" && item.Id != "" {
		return errors.New("must supply type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
