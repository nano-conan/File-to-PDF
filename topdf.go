package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jung-kurt/gofpdf"

)

func main() {
	args := os.Args[1:]
	numOfArgs := len(args)

	for i := 0; i < numOfArgs; i++ {

		file := args[i]

		content, err := ioutil.ReadFile(file)

		if err != nil {
			log.Println(fmt.Sprintf("%s file not found", file))
			continue
		}

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 14)

		pdf.MultiCell(190, 5, string(content), "0", "0", false)

		fileName := excludeFileExtension(args[i])
		_ = pdf.OutputFileAndClose(fmt.Sprintf("%s.pdf", fileName))

		fmt.Println("PDF Created")

	}

}

func excludeFileExtension(file string) string {
	return strings.Split(file, ".")[0]
}
