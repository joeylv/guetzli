package gueztli

import (
	"bytes"
	"github.com/chai2010/guetzli-go"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func init() {

}

//func main() {
//	m0 := loadImage("E:/leveldb.jpg")
//
//	data1 := jpegEncode(m0, 95)
//	data2 := guetzliEncode(m0, 95)
//
//	fmt.Println("jpeg encoded size:", len(data1))
//	fmt.Println("guetzli encoded size:", len(data2))
//
//	if err := ioutil.WriteFile("a.out.jpeg", data1, 0666); err != nil {
//		log.Println(err)
//	}
//	if err := ioutil.WriteFile("a.out.guetzli.jpeg", data2, 0666); err != nil {
//		log.Println(err)
//	}
//
//	fmt.Println("Done")
//}

func LoadImage(name string) image.Image {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func jpegEncode(m image.Image, quality int) []byte {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, m, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func GuetzliEncode(m image.Image, quality int) []byte {
	var buf bytes.Buffer
	err := guetzli.Encode(&buf, m, &guetzli.Options{Quality: quality})
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
