package motherboardinfo

import (
	"computer-specs-viewer/utils"
	"context"
	"fmt"

	"github.com/yusufpapurcu/wmi"
)

type Win32_BaseBoard struct {
	Manufacturer          string
	Status                string
	Product               string
	SerialNumber          string
	Version               string
	Tag                   string
	HostingBoard          bool
	HotSwappable          bool
	Removable             bool
	Replaceable           bool
	RequiresDaughterBoard bool
}

type MotherboardInformation struct {
	Manufacturer          string
	Status                string
	ProductName           string
	SerialNumber          string
	ProductFullName       string
	Version               string
	Tag                   string
	IsHostingBoard        bool
	IsHotSwappable        bool
	IsRemovable           bool
	IsReplaceable         bool
	RequiresDaughterBoard bool
}

func GetMotherboardsInfo() []MotherboardInformation {
	boards, err := GetMotherboardsInfoWithErr()

	if err != nil {
		return []MotherboardInformation{}
	}

	return boards
}

func GetMotherboardsInfoWithErr() ([]MotherboardInformation, error) {
	var boards []MotherboardInformation
	var dst []Win32_BaseBoard

	ctx := context.Background()
	query := wmi.CreateQuery(&dst, "")

	if err := utils.WMIQueryWithContext(ctx, query, &dst); err != nil {
		return boards, err
	}

	for _, board := range dst {
		boardInfo := MotherboardInformation{
			Manufacturer:          board.Manufacturer,
			Status:                board.Status,
			ProductName:           board.Product,
			SerialNumber:          board.SerialNumber,
			ProductFullName:       fmt.Sprintf("%v %v", board.Product, board.SerialNumber),
			Version:               board.Version,
			Tag:                   board.Tag,
			IsHostingBoard:        board.HostingBoard,
			IsHotSwappable:        board.HotSwappable,
			IsRemovable:           board.Removable,
			IsReplaceable:         board.Replaceable,
			RequiresDaughterBoard: board.RequiresDaughterBoard,
		}

		boards = append(boards, boardInfo)
	}

	return boards, nil
}

func PrintMotherboardsInfo() {
	boards := GetMotherboardsInfo()

	utils.PrintSectionTitle("Motherboard")
	utils.PrintStartBlock()

	for i := 0; i < len(boards); i++ {
		printSingleMotherboardInfo(boards[i])

		if i < len(boards)-1 {
			utils.PrintInfoDelim()
		}
	}

	utils.PrintEndBlock()
}

func printSingleMotherboardInfo(board MotherboardInformation) {
	fmt.Printf("Manufacturer: %v\n", board.Manufacturer)
	fmt.Printf("Status: %v\n", board.Status)
	fmt.Printf("Product name: %v\n", board.ProductFullName)
	fmt.Printf("Version: %v\n", board.Version)
	fmt.Printf("Tag: %v\n", board.Tag)
	fmt.Printf("Is hosting board (is a motherboard/baseboard in a chassis): %v\n", board.IsHostingBoard)
	fmt.Printf("Is hot-swappable: %v\n", board.IsHotSwappable)
	fmt.Printf("Is removable: %v\n", board.IsRemovable)
	fmt.Printf("Is replaceable: %v\n", board.IsReplaceable)
	fmt.Printf("Requires daughter board (to function properly): %v\n", board.RequiresDaughterBoard)
}
