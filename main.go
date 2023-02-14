package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	fmt.Println("Application Barcode Started !")

	//Create Barcode
	qrCode, _ := qr.Encode("www.facebook.com", qr.M, qr.Auto)

	//Scale the barcode here
	qrCode, _ = barcode.Scale(qrCode, 300, 300)

	file, _ := os.Create("file/qrcodeFacebook.png")

	// encode the barcode to png
	png.Encode(file, qrCode)

}
