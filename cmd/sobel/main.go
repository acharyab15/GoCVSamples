package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tsobel [imgfile]")
		return
	}

	filename := os.Args[1]

	window := gocv.NewWindow("Sobel Window")
	defer window.Close()

	img := gocv.IMRead(filename, gocv.IMReadUnchanged)
	if img.Empty() {
		fmt.Println("Error")
		return
	}
	defer img.Close()

	gray := gocv.NewMat()
	defer gray.Close()
	gradX := gocv.NewMat()
	defer gradX.Close()
	gradY := gocv.NewMat()
	defer gradY.Close()
	absGradX := gocv.NewMat()
	defer absGradX.Close()
	absGradY := gocv.NewMat()
	defer absGradY.Close()
	grad := gocv.NewMat()
	defer grad.Close()

	gocv.GaussianBlur(img, &img, image.Pt(3, 3), 0, 0, gocv.BorderDefault)
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	gocv.Sobel(gray, &gradX, gocv.MatTypeCV16S, 1, 0, 3, 1, 0, gocv.BorderDefault)

	gocv.Sobel(gray, &gradY, gocv.MatTypeCV16S, 0, 1, 3, 1, 0, gocv.BorderDefault)

	gocv.ConvertScaleAbs(gradX, &absGradX, 0.5, 0.5)
	gocv.ConvertScaleAbs(gradY, &absGradY, 0.5, 0.5)

	gocv.AddWeighted(absGradX, 0.5, absGradY, 0.5, 0, &grad)
	ok := gocv.IMWrite("img.png", grad)
	if !ok {
		fmt.Println("Error")
		return
	}

	for {
		window.IMShow(grad)

		if window.WaitKey(1) == 27 {
			break
		}
	}
}
