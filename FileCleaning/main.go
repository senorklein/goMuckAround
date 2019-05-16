package main

import "fmt"
import "os"
import "strings"

type Conversion struct {
	from string
	to   string
}

func main() {
	fmt.Println("Cleaning your filenames....")

	conversionArray := [3]Conversion{
		Conversion{"-", ""},
		Conversion{"_", ""},
		Conversion{" ", ""},
	}

	fileList := os.Args[1:]

	for i, filename := range fileList {
		fmt.Printf("Cleaning filename %d: %s\n", i, filename)

		newFilename := filename
		for _, conv := range conversionArray {
			newFilename = strings.Replace(newFilename, conv.from, conv.to, -1)
		}
		if filename != newFilename {
			fmt.Printf("converting %s to %s\n", filename, newFilename)
		}

		// check if it's there
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println("file does not exist")
		}

		os.Rename(filename, newFilename)
	}

}
