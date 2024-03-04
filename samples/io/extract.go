package io

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/io"
)

func SampleIoUnzip() {
	destFolder := "unzipped"
	zipFileName := "archive.zip"
	err := io.Unzip(zipFileName, destFolder)
	if err != nil {
		fmt.Println("Error unzipping files:", err)
	}

	println("Files unzipped successfully to:", destFolder)
}
