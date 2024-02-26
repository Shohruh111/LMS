package helper

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel struct {
	Headers []string        `json:"headers"`
	Datas   [][]interface{} `json:"datas"`
}

func GenerateExcelFile(req *Excel) (*excelize.File, error) {
	file := excelize.NewFile()

	// Add a sheet
	sheetName := "Sheet1"
	index := file.NewSheet(sheetName)

	// Add headers
	headers := req.Headers
	for colIdx, header := range headers {
		cell := excelize.ToAlphaString(colIdx+1) + "1"
		file.SetCellValue(sheetName, cell, header)
	}

	for rowIdx, rowData := range req.Datas {
		for colIdx, cellValue := range rowData {
			cell := excelize.ToAlphaString(colIdx+1) + fmt.Sprint(rowIdx+2)
			file.SetCellValue(sheetName, cell, cellValue)
		}
	}

	// Set active sheet
	file.SetActiveSheet(index)

	return file, nil
}
