package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf/v2"
	"github.com/jung-kurt/gofpdfcontrib/gofpdi"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")

	imp := gofpdi.NewImporter()
	tpl := imp.ImportPage(pdf, "demobill.pdf", 1, "/MediaBox")
	pdf.AddPage()
	pdf.SetFillColor(0, 0, 0)
	pdf.SetFontSize(14)
	imp.UseImportedTemplate(pdf, tpl, 0, 0, 200, 0)
	err := pdf.OutputFileAndClose("demo.pdf")
	if err != nil {
		fmt.Println(err)
	}
	// 	gc := draw2dpdf.NewGraphicContext(pdf)
	// 	gc.SetFillColor(color.Black)
	// 	gc.SetStrokeColor(color.Black)
	// 	gc.SetLineWidth(5)
	// 	gc.SetFontSize(14)
	// 	gc.FillStringAt("Hello world", 10, 10)
	// 	gc.Close()
	// 	// err := draw2dpdf.SaveToPdfFile("example.pdf", pdf)
	// 	// if err != nil {
	// 	// 	fmt.Println(err.Error())
	// 	// }
}
