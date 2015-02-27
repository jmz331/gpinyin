package gpinyin

import (
	"regexp"
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

var allUnmarkedVowel map[string]string
var allUnmarkedVowelRegex map[string]*regexp.Regexp

func init() {
	allUnmarkedVowelRegex = map[string]*regexp.Regexp{
		"a": regexp.MustCompile("[āáǎà]"),
		"e": regexp.MustCompile("[ēéěè]"),
		"i": regexp.MustCompile("[īíǐì]"),
		"o": regexp.MustCompile("[ōóǒò]"),
		"u": regexp.MustCompile("[ūúǔù]"),
		"v": regexp.MustCompile("[ǖǘǚǜ]"),
	}
}

func ConvertToPinyinString(source string, separator string, pinyinFormat int) string {
	result := ""
	sourceRuneArray := []rune(source)
	runeLength := len(sourceRuneArray)
	for index := 0; index < runeLength; index++ {
		runeValue := sourceRuneArray[index]
		char := string(runeValue)
		if isChinese(char) || char == "〇" {
			rightIndex := index + rightMove
			foundFlag := false
			if rightIndex > runeLength {
				rightIndex = runeLength - 1
			}
			for ; rightIndex > index; rightIndex-- {
				wordRuneArray := sourceRuneArray[index:(rightIndex + 1)]
				pinyinString := toMultiPinyin(string(wordRuneArray))
				if len(pinyinString) > 0 {
					pinyinStringArray := formatPinyin(pinyinString, pinyinFormat)
					result += strings.Join(pinyinStringArray, separator)
					index = rightIndex
					foundFlag = true
					break
				}
			}
			if !foundFlag {
				char = string(sourceRuneArray[index])
				pinyinStringArray := convertPinyinArray(char, pinyinFormat)
				if pinyinStringArray != nil {
					result += pinyinStringArray[0]
				} else {
					result += char
				}
			}
			if index < (runeLength - 1) {
				result += separator
			}
		} else {
			result += char
			if ((index + 1) < runeLength) && isChinese(string(sourceRuneArray[(index+1)])) {
				result += separator
			}
		}
	}
	return result
}

func formatPinyin(pinyinString string, pinyinFormat int) []string {
	switch pinyinFormat {
	case PINYIN_WITH_TONE_MARK:
		return strings.Split(pinyinString, pinyin_separator)
	case PINYIN_WITH_TONE_NUMBER:
		//todo: 包含数字音标的拼音实现
		return nil
	case PINYIN_WITHOUT_TONE:
		return convertWithOutTone(pinyinString)
	}
	return nil
}

/**
 * pinyinArrayString 带声调格式的拼音
 */
func convertWithOutTone(pinyinString string) []string {
	for k, v := range allUnmarkedVowelRegex {
		pinyinString = v.ReplaceAllString(pinyinString, k)
	}
	return strings.Split(pinyinString, pinyin_separator)
}

func convertPinyinArray(char string, pinyinFormat int) []string {
	pinyinString := toPinyin(char)
	if len(pinyinString) > 0 {
		return formatPinyin(pinyinString, pinyinFormat)
	}
	return nil
}
