package src

import "fmt"

const StartBlock = "================================"
const EndBlock = StartBlock
const InfoDelim = "--------------------------------"

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
