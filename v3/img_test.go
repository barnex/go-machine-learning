package vs

func ExampleImgFromSlice() {
	list := []float64{1, 2, 3, 4, 5, 6}
	img := ImgFromSlice(list, 2, 3)
	img.Print()

	//Output:
	//1.00 2.00 3.00
	//4.00 5.00 6.00
}
