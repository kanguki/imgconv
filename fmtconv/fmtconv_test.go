package fmtconv

import (
	"testing"
)

func TestToOutFileName(t *testing.T) {
	cases := []struct {
		inPath, outFmt, expectOutPath string
	}{
		{inPath: "x/y/z.png", outFmt: "jpg", expectOutPath: "x/y/z.jpg"},
		{inPath: "x/y/z", outFmt: "jpg", expectOutPath: "out.jpg"},
	}
	for _, c := range cases {
		if toOutFileName(c.inPath, Format(c.outFmt)) != c.expectOutPath {
			t.Fatalf("unxpected results for toOutFileName: %v", c)
		}
	}
}

func TestSameFormat(t *testing.T) {
	cases := []struct {
		inPath, outFmt string
		expect         bool
	}{
		{inPath: "x/y/z.png", outFmt: "png", expect: true},
		{inPath: "x/y/z.jpg", outFmt: "png", expect: false},
	}
	for _, c := range cases {
		if sameFormat(c.inPath, Format(c.outFmt)) != c.expect {
			t.Fatalf("unxpected results for inAndOutHaveTheSameFormat: %v", c)
		}
	}
}

func TestConvert(t *testing.T) {
	_, err := Convert("assets/test.jpeg", PNG)
	if err != nil {
		t.Fatal(err)
	}
}
