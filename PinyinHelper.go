package gpinyin

import (
	"fmt"
)

const (
	//默认处理多音字时，依次从右侧添加3，2，1来判断
	rightMove = 3
)

func ConvertToPinyinString(source string, separator string) {
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
				wordArray := sourceRuneArray[index:rightIndex]
				fmt.Println(wordArray)
			}
		} else {
			result += char
		}
	}
}
