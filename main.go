package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: pdftojpg <pdf file> <output> <number of page>")
		os.Exit(1)
	}

	pdfName := os.Args[1]
	imageName := os.Args[2]
	pagesStr := os.Args[3]

	fmt.Printf("Converting %s to %s\n", pdfName, imageName)

	if pdfName == "" || imageName == "" || pagesStr == "" {
		log.Fatal("No pdf file or output file or number of pages specified")
		os.Exit(1)
	}

	pages, err := strconv.ParseInt(pagesStr, 10, 32)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := ConvertPdfToJpg(pdfName, imageName, pages); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// ConvertPdfToJpg will take a filename of a pdf file and convert the file into an
// image which will be saved back to the same location. It will save the image as a
// high resolution jpg file with minimal compression.
func ConvertPdfToJpg(pdfName string, imageName string, pages int64) error {

	// Setup
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Must be *before* ReadImageFile
	// Make sure our image is high quality
	if err := mw.SetResolution(100, 100); err != nil {
		return err
	}

	// Load the image file into imagick
	if err := mw.ReadImage(pdfName); err != nil {
		return err
	}

	// Must be *after* ReadImageFile
	// Flatten image and remove alpha channel, to prevent alpha turning black in jpg
	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE); err != nil {
		return err
	}

	// Set any compression (100 = max quality)
	if err := mw.SetCompressionQuality(50); err != nil {
		return err
	}

	for i := int64(0); i < pages; i++ {
		// Select only first page of pdf
		mw.SetIteratorIndex(int(i))

		// Convert into JPG
		if err := mw.SetFormat("jpg"); err != nil {
			return err
		}

		// Save File
		err := mw.WriteImage(fmt.Sprintf("%03d.jpg", i + 1))
		if err != nil {
			log.Fatal(fmt.Sprintf("error at page: %d", i+1))
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
	}

	return nil
}
