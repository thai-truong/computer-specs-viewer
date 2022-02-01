package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func SpaceOutFieldNames(field string) string {
	spacedOutField := strings.Builder{}
	spacedOutField.WriteByte(field[0])

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
