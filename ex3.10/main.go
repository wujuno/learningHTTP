package main

import (
	"io"
	"mime/multipart"
	"os"
)

func main() {
	var buffer bytes.buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name:", "Micheal Jackson")

	fileWriter, err := writer.CreateFormField("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}

	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()
}