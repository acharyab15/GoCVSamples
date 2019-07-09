package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("cmd/sobel/coins.jpg", gocv.IMReadUnchanged)
	if img.Empty() {
		fmt.Println("error")
	}
	defer img.Close()

	window := gocv.NewWindow("Sobel Window")
	defer window.Close()

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

	for {
		window.IMShow(grad)

		if window.WaitKey(1) == 27 {
			break
		}
	}
}
