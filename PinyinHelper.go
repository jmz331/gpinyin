package gpinyin

import (
	"fmt"
	"strings"
)

const (
	//默认处理多音字时，依次从右侧添加3，2，1来判断
	rightMove = 3
	//配置文件中的拼音分割字符
	pinyin_separator = ","
	//转换类型
	PINYIN_WITH_TONE_MARK, PINYIN_WITH_TONE_NUMBER, PINYIN_WITHOUT_TONE = 1, 2, 3
	//无音标拼音字母
	all_unmarked_vowel = "aeiouv"
	//带音标的拼音字母
	all_marked_vowel = "āáǎàēéěèīíǐìōóǒòūúǔùǖǘǚǜ"
)

func ConvertToPinyinString(source string, separator string) string {
	result := ""
	sourceRuneArray := []rune(source)
	runeLength := len(sourceRuneArray)
	for index, runeValue := range sourceRuneArray {
		char := string(runeValue)
		if isChinese(char) || char == "〇" {
			rightIndex := index + rightMove
			if rightIndex > runeLength {
				rightIndex = runeLength - 1
			}
			for ; rightIndex > index; rightIndex-- {
				wordRuneArray := sourceRuneArray[index:rightIndex]

				fmt.Println(wordArray, ":", string(wordArray))
			}
		} else {
			result += char
		}
	}
	return result
}

func formatPinyin(pinyinString string, pinyinFormat int) []string {
	switch pinyinFormat {
	case PINYIN_WITH_TONE_MARK:
		return strings.Split(pinyinString, pinyin_separator)
	case PINYIN_WITH_TONE_NUMBER:
		break
	case PINYIN_WITHOUT_TONE:
		return convertWithOutTone(pinyinString)
	}
}

func convertWithOutTone(pinyinArrayString string) []string {

}
