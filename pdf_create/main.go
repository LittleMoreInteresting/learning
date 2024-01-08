package main

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	// 创建一个新的 PDF 文件
	outfile := "D:\\code\\learning\\pdf_create\\test.pdf"
	json := "D:\\code\\learning\\pdf_create\\pdf.json"
	conf := api.LoadConfiguration()
	fonts := []string{
		"D:\\go_stu\\pkg\\mod\\github.com\\pdfcpu\\pdfcpu@v0.6.0\\pkg\\testdata\\fonts\\unifont-13.0.03.ttf",
	}
	_ = api.InstallFonts(fonts)
	err2 := api.CreateFile("", json, outfile, conf)
	fmt.Println(err2)
}
