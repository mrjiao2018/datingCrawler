package persist

import (
	"context"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name:     "test",
		Gender:   "male",
		Birth:    "双鱼座",
		Age:      20,
		Marriage: "未婚",
	}

	id, err := save(profile)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", result.Source)

	var actual model.Profile
	err = json.Unmarshal(result.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != profile {
		t.Logf("expected %+v, got %+v", profile, actual)
	}
}
