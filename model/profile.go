package model

type Profile struct {
	Name            string //姓名（可能为用户自定义的昵称）
	Gender          string //性别
	Birth           string //生日信息，此处通过星座来代替
	Age             int    //年纪
	Marriage        string //婚姻状况
	MarriageWilling string //结婚意愿
	WorkLocation    string //工作地点
	LivingLocation  string //当前居住地
	NativePlace     string //籍贯
	BodyShape       string //体型
	Nation          string //民族
	Income          string //收入
	Education       string //教育情况
	Height          int    //身高(cm)
	Weight          int    //体重(kg)
	Work            string //工作内容
	HasChild        string //是否有子女
	WantChild       string //是否想要子女
	Car             string //购车状况
	House           string //购房状况
	Cigarette       string //抽烟状况
	Wine            string //喝酒状况
}
