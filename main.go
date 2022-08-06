package main

import (
	"fmt"
	"os"

	"github.com/kanguki/imgconv/fmtconv"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		printError("error: not enough args: %v", args)
		printGuide()
		return
	}
	outFile, err := fmtconv.Convert(args[1], fmtconv.Format(args[2]))
	if err != nil {
		printError("error converting: %v", err)
		printGuide()
		return
	}
	printInfo("convert to file: %v", outFile)
}
func printInfo(str string, vals ...interface{}) {
	print(fgBlue, str, vals...)
}
func printError(str string, vals ...interface{}) {
	print(fgRed, str, vals...)
}
func print(color color, str string, vals ...interface{}) {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", color, fmt.Sprintf(str, vals...))
}

type color int

const (
	fgBlack color = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

func printGuide() {
	fmt.Print(`
NAME
	imgconv - convert image from one format to another
SYNOPSIS:
	go run main.go [in] [out_exprected]
DESCRIPTION:
	[in]: first argument. Must be an absolute path. Supported formats are: jpg | jpeg | png
	[out_expected]: out format expected. Supported values are: jpg | jpeg | png
OUTPUT SUMMARY:
	an image with same name and expected format in the same dir as the input image 
`)
}
