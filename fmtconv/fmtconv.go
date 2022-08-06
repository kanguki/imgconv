package fmtconv

import (
	"fmt"
	"golang.org/x/image/webp"
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
	WEBP Format = "webp"
)

func Convert(inPath string, outFmtExpected Format) (outPath string, err error) {
	if sameFormat(inPath, outFmtExpected) {
		return "", fmt.Errorf("file is already in expected format")
	}
	rawImg, err := getRawImage(inPath)
	if err != nil {
		return "", fmt.Errorf("getRawImage %v %v", inPath, err)
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
		return "", fmt.Errorf("output %v is not supported", outFmtExpected)
	}
}

func getRawImage(inPath string) (image.Image, error) {
	in, err := os.Open(inPath)
	if err != nil {
		return nil, fmt.Errorf("open file %v %v", in, err)
	}
	defer in.Close()
	inFmt := extractImgFormat(inPath)
	switch inFmt {
	case WEBP:
		return webp.Decode(in)
	case JPEG, JPG, PNG:
		img, _, err := image.Decode(in)
		return img, err
	default:
		return nil, fmt.Errorf("not a supported file format")
	}
}

func extractImgFormat(imageFilePath string) Format {
	for i := len(imageFilePath) - 1; i >= 0; i-- {
		if string(imageFilePath[i]) == "." {
			return Format(imageFilePath[i+1:])
		}
	}
	return ""
}

//sameFormat check if input file and expected output format has the same file extension already
func sameFormat(inPath string, outExpectedFmt Format) bool {
	inFmt := extractImgFormat(inPath)
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
