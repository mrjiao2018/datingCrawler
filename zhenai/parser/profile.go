package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
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
func ParseProfile(contents []byte, url string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = extractString(contents, nameRe)
	profile.LivingLocation = extractString(contents, livingLocationRe)
	profile.Birth = extractString(contents, birthRe)
	profile.Education = extractString(contents, educationRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Income = extractString(contents, incomeRe)
	profile.WorkLocation = extractString(contents, workLocationRe)
	profile.Work = extractString(contents, workRe)
	profile.Nation = extractString(contents, nationRe)
	profile.NativePlace = extractString(contents, nativePlaceRe)
	profile.BodyShape = extractString(contents, bodyShapeRe)
	profile.Cigarette = extractString(contents, cigaretteRe)
	profile.Wine = extractString(contents, wineRe)
	profile.House = extractString(contents, houseRe)
	profile.HasChild = extractString(contents, hasChildRe)
	profile.WantChild = extractString(contents, wantChildRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.MarriageWilling = extractString(contents, marriageWillingRe)
	setIntValue(&profile.Age, contents, ageRe)
	setIntValue(&profile.Height, contents, heightRe)
	setIntValue(&profile.Weight, contents, weightRe)
	//age, ageErr := strconv.Atoi(extractString(contents, ageRe))
	//if ageErr == nil {
	//	profile.Age = age
	//}
	//height, heightErr := strconv.Atoi(extractString(contents, heightRe))
	//if heightErr == nil {
	//	profile.Height = height
	//}
	//weight, weightErr := strconv.Atoi(extractString(contents, weightRe))
	//if weightErr == nil {
	//	profile.Weight = weight
	//}
	return engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idRe),
				PayLoad: profile,
			},
		},
	}
}

// 设置 profile 中的 Int 类型的属性
func setIntValue(profileAttribute *int, contents []byte, reg *regexp.Regexp) {
	value, err := strconv.Atoi(extractString(contents, reg))
	if err == nil {
		*profileAttribute = value
	}
}

//根据正则表达式从 contents 中找到匹配的内容
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
