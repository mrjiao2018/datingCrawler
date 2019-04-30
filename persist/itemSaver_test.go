package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Id:   "1234567890",
		Type: "zhenai",
		Url:  "www.test.com",
		PayLoad: model.Profile{
			Name:     "test",
			Gender:   "male",
			Birth:    "双鱼座",
			Age:      20,
			Marriage: "未婚",
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = save("dating_profile_test", client, expected)
	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", result.Source)

	var actual engine.Item
	err = json.Unmarshal(result.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.PayLoad)
	actual.PayLoad = actualProfile

	if actual != expected {
		t.Logf("expected %+v, got %+v", expected, actual)
	}
}
