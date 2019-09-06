package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/corona10/goimagehash"
)

func main() {
	file1, _ := os.Open("cat.jpg")
	file2, _ := os.Open("cat-compress.jpg")
	file3, _ := os.Open("cat-rotation.jpg")
	file4, _ := os.Open("cat-watermark.jpg")
	file5, _ := os.Open("cat-watermark-compress.jpg")
	file6, _ := os.Open("14m.gif")
	file7, _ := os.Open("24m.gif")
	file8, _ := os.Open("png.png")
	// file9, _ := os.Open("error2.jpg") // 存在图片编码异常，无法解码的情况
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()
	defer file4.Close()
	defer file5.Close()
	defer file6.Close()
	defer file7.Close()
	defer file8.Close()
	// defer file9.Close()

	img1, m, _ := image.Decode(file1)
	fmt.Println(m)
	img2, m, _ := image.Decode(file2)
	img3, m, _ := image.Decode(file3)
	img4, m, _ := image.Decode(file4)
	img5, m, _ := image.Decode(file5)
	img6, m, _ := image.Decode(file6)
	fmt.Println(m)

	img7, m, _ := image.Decode(file7)
	img8, m, _ := image.Decode(file8)
	fmt.Println(m)

	// img9, err := jpeg.Decode(file9)
	// if err != nil {
	// 	fmt.Println("=====----====:", err)
	// }

	hash1, err := goimagehash.PerceptionHash(img1)
	hash2, _ := goimagehash.PerceptionHash(img2)
	hash3, _ := goimagehash.PerceptionHash(img3)
	hash4, _ := goimagehash.PerceptionHash(img4)
	hash5, _ := goimagehash.PerceptionHash(img5)
	hash6, _ := goimagehash.PerceptionHash(img6)
	hash7, _ := goimagehash.PerceptionHash(img7)
	hash8, _ := goimagehash.PerceptionHash(img8)
	// hash9, err := goimagehash.PerceptionHash(img9)
	if err != nil {
		fmt.Println("=========:", err)
	}

	distance, _ := hash1.Distance(hash6)
	fmt.Printf("Distance between images: %v\n", distance)

	fmt.Println("cat:                   ", hash1.ToString())
	fmt.Println("cat-compress:          ", hash2.ToString())
	fmt.Println("cat-rotation:          ", hash3.ToString())
	fmt.Println("cat-watermark:         ", hash4.ToString())
	fmt.Println("cat-watermark-compress:", hash5.ToString())
	fmt.Println("14m gif:               ", hash6.ToString())
	fmt.Println("20m gif:               ", hash7.ToString())
	fmt.Println("png png:               ", hash8.ToString())
	// fmt.Println("error2 :               ", hash9.ToString())
	fmt.Println(hash1.Bits())
	fmt.Println(hash2.Bits())
}
