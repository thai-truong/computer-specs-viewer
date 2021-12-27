package utils

import (
	"fmt"
	"math"
	"strings"
)

const StartBlock = "================================"
const EndBlock = StartBlock
const InfoDelim = "--------------------------------"
const ItemDelim = "++++++++++++++++++++++++++++++++"

func PrintStartBlock() {
	fmt.Println(StartBlock)
}

func PrintEndBlock() {
	fmt.Println(EndBlock)
}

func PrintSectionTitle(sectionName string) {
	fmt.Printf("\n\n%v Section:\n\n", sectionName)
}

func PrintInfoDelim() {
	fmt.Printf("\n%v\n\n", InfoDelim)
}

func PrintItemDelim() {
	fmt.Printf("%v\n", ItemDelim)
}

func PrintStrWithOrder(str string, order int) {
	fmt.Printf("%v #%v\n", str, order)
}

func GetStrListAsStr(strList []string) string {
	var contentStr string

	for i := 0; i < len(strList); i++ {
		contentStr += strList[i]

		if i < len(strList)-1 {
			contentStr += ", "
		}
	}

	return fmt.Sprintf("[%v]", contentStr)
}

func GetPercentString(percentFloat float64) string {
	percent := percentFloat * 100

	return fmt.Sprintf("%.2f%%", percent)
}

func GetSpaceString(spaceBytes uint64, unit string) string {
	spaceUnitExponent := map[string]int{
		"B":  0,
		"KB": 3,
		"MB": 6,
		"GB": 9,
	}

	unitExp, found := spaceUnitExponent[strings.ToUpper(unit)]

	if !found {
		return fmt.Sprintf("%v B", spaceBytes)
	}

	spaceAmtInUnit := float64(spaceBytes) / math.Pow(10, float64(unitExp))
	return fmt.Sprintf("%.2f %v", spaceAmtInUnit, unit)
}
