package parser

import (
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	contents, err := ioutil.ReadFile("cityListData.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	const resultSize = 470
	if resultSize != len(result.Requests) {
		log.Error("size of result requests should be %d rather %d", resultSize, len(result.Requests))
	}
	if resultSize != len(result.Items) {
		log.Error("size of result items should be %d rather %d", resultSize, len(result.Items))
	}
}
