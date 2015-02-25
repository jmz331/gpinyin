package gpinyin

import (
	"regexp"
)

var traditionalChinese map[string]string
var simplifiedChinese map[string]string
var chineseRegex *regexp.Regexp

func init() {
	traditionalChinese = make(map[string]string)
	err := loadResource("chinese.db", traditionalChinese, true)
	if err != nil {
		panic(err)
	}
	simplifiedChinese = make(map[string]string)
	err = loadResource("chinese.db", simplifiedChinese, false)
	if err != nil {
		panic(err)
	}
	chineseRegex = regexp.MustCompile("[\u4e00-\u9fa5]")
}

func ConvertToSimplifiedChinese(source string) string {
	result := ""
	for _, runeValue := range source {
		result += toSimplifiedChinese(string(runeValue))
	}
	return result
}

func ConvertToTraditionalChinese(source string) string {
	result := ""
	for _, runeValue := range source {
		result += toTraditionalChinese(string(runeValue))
	}
	return result
}

func toSimplifiedChinese(source string) string {
	v := simplifiedChinese[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func toTraditionalChinese(source string) string {
	v := traditionalChinese[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func isChinese(char string) bool {
	return chineseRegex.MatchString(char)
}
