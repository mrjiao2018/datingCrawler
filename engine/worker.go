package engine

import (
	"crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching url : %s", r.Url)
	bytes, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url : %s, %v", r.Url, err)
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(bytes)

	return parseResult, nil

}
