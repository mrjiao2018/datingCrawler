package parser

import (
	"crawler/engine"
	"crawler/model"
	"crawler/tools"
	"regexp"
	"strconv"
)

var nameRe = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)
var genderRe = regexp.MustCompile(`(男士|女士)征婚`)
var birthRe = regexp.MustCompile(`(白羊座|金牛座|双子座|巨蟹座|狮子座|处女座|天秤座|天蝎座|射手座|魔羯座|水瓶座|双鱼座)`)
var ageRe = regexp.MustCompile(`([0-9]+)岁`)
var marriageRe = regexp.MustCompile(`(未婚|离异|丧偶)`)
var marriageWillingRe = regexp.MustCompile(`何时结婚:([^<]+)</div>`)
var workLocationRe = regexp.MustCompile(`工作地:([^<]+)</div>`)
var livingLocationRe = regexp.MustCompile(`<div class="des f-cl" [^>]*>([^ ]+) `)
var nativePlaceRe = regexp.MustCompile(`籍贯:([^<]+)</div>`)
var bodyShapeRe = regexp.MustCompile(`体型:([^<]+)</div>`)
var nationRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^族]+族)</div>`)
var incomeRe = regexp.MustCompile(`月收入:([^<]+)</div>`)
var educationRe = regexp.MustCompile(`(高中及以下|中专|大专|大学本科|硕士|博士)`)
var heightRe = regexp.MustCompile(`([0-9]+)cm`)
var weightRe = regexp.MustCompile(`([0-9]+)kg`)
var workRe = regexp.MustCompile(`(销售|客户服务|计算机/互联网|通信/电子|生产/制造|物流/仓储|商贸/采购|人事/行政|
									高级管理|广告/市场|传媒/艺术|生物/制药|医疗/护理|金融/银行/保险|建筑/房地产|咨询/顾问|法律
									|财会/省计|教育/科研|服务业|交通运输|政府机构|军人/警察|农林牧渔|自由职业|在校学生|待业|其他行业)`)
var hasChildRe = regexp.MustCompile(`(没有小孩|有孩子且住一起|有孩子且偶尔会一起住|有孩子但不在身边)`)
var wantChildRe = regexp.MustCompile(`是否想要孩子:([^<]+)</div>`)
var carRe = regexp.MustCompile(`(已买车|未买车)`)
var houseRe = regexp.MustCompile(`(和家人同住|已购房|租房|打算婚后购房|住在单位宿舍)`)
var cigaretteRe = regexp.MustCompile(`(不吸烟|稍微抽一点烟|烟抽得很多|社交场合会抽烟)`)
var wineRe = regexp.MustCompile(`(不喝酒|稍微喝一点酒|酒喝得很多|社交场合会喝酒)`)

// 解析每个用户的详细信息
// 如 http://album.zhenai.com/u/102088914 即对应着一个具体的用户信息
func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	//以下前置添加属性字段名是为了方便测试
	profile.Name = "Name: " + tools.ExtractString(contents, nameRe)
	profile.LivingLocation = "LivingLocation: " + tools.ExtractString(contents, livingLocationRe)
	profile.Birth = "Birth: " + tools.ExtractString(contents, birthRe)
	profile.Education = "Education: " + tools.ExtractString(contents, educationRe)
	profile.Marriage = "Marriage: " + tools.ExtractString(contents, marriageRe)
	profile.Income = "Income: " + tools.ExtractString(contents, incomeRe)
	profile.WorkLocation = "WorkLocation: " + tools.ExtractString(contents, workLocationRe)
	profile.Work = "Work: " + tools.ExtractString(contents, workRe)
	profile.Nation = "Nation: " + tools.ExtractString(contents, nationRe)
	profile.NativePlace = "NativePlace: " + tools.ExtractString(contents, nativePlaceRe)
	profile.BodyShape = "BodyShape: " + tools.ExtractString(contents, bodyShapeRe)
	profile.Cigarette = "Cigarette: " + tools.ExtractString(contents, cigaretteRe)
	profile.Wine = "Wine: " + tools.ExtractString(contents, wineRe)
	profile.House = "House: " + tools.ExtractString(contents, houseRe)
	profile.HasChild = "HasChild: " + tools.ExtractString(contents, hasChildRe)
	profile.WantChild = "WantChild: " + tools.ExtractString(contents, wantChildRe)
	profile.Gender = "Gender: " + tools.ExtractString(contents, genderRe)
	profile.Car = "Car: " + tools.ExtractString(contents, carRe)
	profile.MarriageWilling = "MarriageWilling: " + tools.ExtractString(contents, marriageWillingRe)

	//profile.Name = tools.ExtractString(contents, nameRe)
	//profile.LivingLocation = tools.ExtractString(contents, livingLocationRe)
	//profile.Birth = tools.ExtractString(contents, birthRe)
	//profile.Education = tools.ExtractString(contents, educationRe)
	//profile.Marriage = tools.ExtractString(contents, marriageRe)
	//profile.Income = tools.ExtractString(contents, incomeRe)
	//profile.WorkLocation = tools.ExtractString(contents, workLocationRe)
	//profile.Work = tools.ExtractString(contents, workRe)
	//profile.Nation = tools.ExtractString(contents, nationRe)
	//profile.NativePlace = tools.ExtractString(contents, nativePlaceRe)
	//profile.BodyShape = tools.ExtractString(contents, bodyShapeRe)
	//profile.Cigarette = tools.ExtractString(contents, cigaretteRe)
	//profile.Wine = tools.ExtractString(contents, wineRe)
	//profile.House = tools.ExtractString(contents, houseRe)
	//profile.HasChild = tools.ExtractString(contents, hasChildRe)
	//profile.WantChild = tools.ExtractString(contents, wantChildRe)
	//profile.Gender = tools.ExtractString(contents, genderRe)
	//profile.Car = tools.ExtractString(contents, carRe)
	//profile.MarriageWilling = tools.ExtractString(contents, marriageWillingRe)

	age, ageErr := strconv.Atoi(tools.ExtractString(contents, ageRe))
	if ageErr == nil {
		profile.Age = age
	}

	height, heightErr := strconv.Atoi(tools.ExtractString(contents, heightRe))
	if heightErr == nil {
		profile.Height = height
	}

	weight, weightErr := strconv.Atoi(tools.ExtractString(contents, weightRe))
	if weightErr == nil {
		profile.Weight = weight
	}

	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}
