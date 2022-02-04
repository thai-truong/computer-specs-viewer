package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func SpaceOutFieldNames(field string) string {
	spacedOutField := strings.Builder{}
	spacedOutField.WriteByte(field[0])

	// Exclude fields that are all caps
	if IsAllCharsUpper(field) {
		return field
	}

	for _, c := range field[1:] {
		if unicode.IsUpper(c) {
			spacedOutField.WriteString(fmt.Sprintf(" %v", string(c)))
		} else {
			spacedOutField.WriteRune(c)
		}
	}

	return spacedOutField.String()
}

func GetStrWithOrder(str string, order string) string {
	return fmt.Sprintf("%s %s\n", str, order)
}

func IsAllCharsUpper(str string) bool {
	for _, c := range str {
		if !unicode.IsUpper(c) {
			return false
		}
	}

	return true
}

func StrListToPrettyStr(strs []string) string {
	res := "[\n"

	for _, str := range strs {
		res += fmt.Sprintf("\t%s,\n", str)
	}

	res += "]"
	return res
}
