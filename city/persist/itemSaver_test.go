package persist

import (
	"context"
	"model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {

	profile := model.Profile{
		CityYear:   "上海静安", //所属城市 所需的工作经验
		PostSalary: "开发",   //岗位 薪资
		Status:     "以招聘",  //招聘状态

	}
	id, err := save(profile)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").
		Type("boss").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", resp)
}
