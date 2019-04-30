package parser

import (
	"crawler/engine"
	"regexp"
)

// 将所有的正则表达式编译放到函数外面提升程序性能
var userUrlRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">[^<]+</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)`)

// 解析每个城市对应的用户信息概览界面，获取用户的 url
// 如 http://www.zhenai.com/zhenghun/beijing ，ParseCity 会获取该 url （北京市）下面的所有用户的 url
func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	userUrlMatches := userUrlRe.FindAllSubmatch(contents, -1)
	cityUrlMatches := cityUrlRe.FindAllSubmatch(contents, -1)

	for _, v := range userUrlMatches {
		userUrl := string(v[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: userUrl,
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, userUrl)
			},
		})
	}

	for _, v := range cityUrlMatches {
		cityUrl := string(v[1])

		result.Requests = append(result.Requests, engine.Request{
			Url:        cityUrl,
			ParserFunc: ParseCity,
		})
	}

	return result
}

//// 将所有的正则表达式编译放到函数外面提升程序性能
//
//// 解析每个城市对应的用户信息概览界面，获取用户的 url
//// 如 http://www.zhenai.com/zhenghun/beijing ，ParseCity 会获取该 url （北京市）下面的所有用户的 url
//func ParseCity(contents []byte) engine.ParseResult {
//	urlAndNameMatches := userUrlAndNameRe.FindAllSubmatch(contents, -1)
//
//	result := engine.ParseResult{}
//
//	for _, v := range urlAndNameMatches {
//		//result.Items = append(result.Items, "User : "+string(v[2]))
//		result.Requests = append(result.Requests, engine.Request{
//			Url:        string(v[1]),
//			ParserFunc: ParseProfile,
//		})
//	}
//
//	return result
//}

// 如果希望在 ParseCity 函数中传递部分参数到 ParseProfile 函数中，可以使用闭包的函数思想
// 如下所示：此处是将用户名当做参数传递给 ParseProfile 函数
// （ParseProfile函数的定义已经改变，为func ParseProfile(contents []byte, name string) engine.ParseResult{}）
//
//var userUrlAndNameRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)
//
//func ParseCity(contents []byte) engine.ParseResult{
//	urlAndNameMatches := userUrlAndNameRe.FindAllSubmatch(contents, -1)
//
//	result := engine.ParseResult{}
//
//	for _, v := range urlAndNameMatches {
//		name := string(v[2])
//		result.Items = append(result.Items, name)
//		result.Requests = append(result.Requests, engine.Request{
//			Url:        string(v[1]),
//			// 此处是故意使用闭包将参数传递到 ParseProfile 函数中
//			// 之所以不直接传入 name 而不是 string(v[2])是因为此处只是声明了调用哪个函数，
//			// 具体函数的执行会等到该 for 循环结束，从而导致传入的 string(v[2]) 全是一样的值
//			// 使用 name 存储 string(v[2]) 值则可以避免这一情况
//			ParserFunc: func(contents []byte) engine.ParseResult {
//				return ParseProfile(contents, name)
//			},
//		})
//	}
//
//	return result
//}
