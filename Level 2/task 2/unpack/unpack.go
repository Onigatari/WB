package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func isEscape(tmp rune) bool {
	return tmp == rune('\\')
}

func UnpackingString(str string) (string, error) {
	var buf string
	var isEscaped bool
	var res strings.Builder

	for _, r := range str {
		s := string(r)

		if unicode.IsDigit(r) {
			if buf == "" {
				return "", errors.New("invalid input string")
			}

			// Если число равно счетчику, то сохранить буферизованный символ в результат
			if buf != "" && !isEscaped {
				n, _ := strconv.Atoi(s)
				res.WriteString(strings.Repeat(buf, n))
				buf = ""
				continue
			}
		}

		if buf != "" {
			if isEscape(r) && !isEscaped {
				isEscaped = true
				continue
			}

			isEscaped = false
			res.WriteString(buf)
			buf = s
			continue
		}

		buf = s
	}

	// Отправка оставшегося результата в результат
	res.WriteString(buf)

	return res.String(), nil
}
