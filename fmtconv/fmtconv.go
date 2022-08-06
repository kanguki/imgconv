package fmtconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

type Format string

const (
	PNG  Format = "png"
	JPEG Format = "jpeg"
	JPG  Format = "jpg"
)

func Convert(inPath string, outFmtExpected Format) (outPath string, err error) {
	if sameFormat(inPath, outFmtExpected) {
		return "", fmt.Errorf("file is already in expected format")
	}
	in, err := os.Open(inPath)
	if err != nil {
		return "", fmt.Errorf("open file %v %v", in, err)
	}
	defer in.Close()
	rawImg, _, err := image.Decode(in)
	if err != nil {
		return "", fmt.Errorf("image.Decode %v", err)
	}
	outFile := toOutFileName(inPath, outFmtExpected)
	f, err := os.Create(outFile)
	if err != nil {
		return "", fmt.Errorf("creating file %v %v", outFile, err)
	}
	defer f.Close()
	switch outFmtExpected {
	case PNG:
		enc := png.Encoder{
			CompressionLevel: png.BestSpeed,
		}
		err = enc.Encode(f, rawImg)
		if err != nil {
			return "", fmt.Errorf("png encode %v", err)
		}
		return outFile, nil
	case JPEG, JPG:
		opt := jpeg.Options{
			Quality: 100,
		}
		err = jpeg.Encode(f, rawImg, &opt)
		if err != nil {
			return "", fmt.Errorf("jpeg encode %v", err)
		}
		return outFile, nil
	default:
		return "", fmt.Errorf("%v is not supported", outFmtExpected)
	}
}

//sameFormat check if input file and expected output format has the same file extension already
func sameFormat(inPath string, outExpectedFmt Format) bool {
	var inFmt Format
	for i := len(inPath) - 1; i >= 0; i-- {
		if string(inPath[i]) == "." {
			inFmt = Format(inPath[i+1:])
			break
		}
	}
	switch Format(inFmt) {
	case JPEG, JPG:
		return outExpectedFmt == JPEG || outExpectedFmt == JPG
	default:
		return inFmt == outExpectedFmt
	}
}

//toOutFileName converts /abs/input/file.actualFmt to /abs/input/file.expectedFmt
func toOutFileName(absoluteInFilePath string, expectedFmt Format) string {
	for i := len(absoluteInFilePath) - 1; i >= 0; i-- {
		if string(absoluteInFilePath[i]) == "." {
			return fmt.Sprintf("%v.%v", string(absoluteInFilePath[:i]), expectedFmt)
		}
	}
	return "out." + string(expectedFmt)
}
