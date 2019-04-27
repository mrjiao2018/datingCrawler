package tools

import "regexp"

//根据正则表达式从 contents 中找到匹配的内容
func ExtractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
