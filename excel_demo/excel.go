package excel_demo

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Read() {
	f, err := excelize.OpenFile("230419161420.xlsx")
	if err != nil {
		return
	}
	defer f.Close()
	index := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(index)
	rows, err := f.GetRows(sheet, excelize.Options{RawCellValue: true})
	if err != nil {
		return
	}
	for i, row := range rows {
		fmt.Println(i, row)
	}
}
