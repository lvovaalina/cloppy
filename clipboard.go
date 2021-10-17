package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"sync"

	"golang.design/x/clipboard"
)

func watchImg(wg sync.WaitGroup) {
	defer wg.Done()
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := clipboard.Watch(context, clipboard.FmtImage)
	for data := range ch {
		serveFrames(data)
		fmt.Println(string(data))
	}
}

func watchText(wg sync.WaitGroup) {
	defer wg.Done()
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := clipboard.Watch(context, clipboard.FmtText)
	for data := range ch {
		add(data)
		fmt.Println("Add value: " + string(data))
	}
}

func serveFrames(imgByte []byte) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./img.png")
	defer out.Close()

	err = png.Encode(out, img)

	if err != nil {
		log.Println(err)
	}

}

func clipboardManagerInit() {
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("Main: Starting worker")
	go watchText(wg)

	wg.Add(1)
	go watchImg(wg)

	wg.Wait()

	fmt.Println("Main: Waiting for worker to finish")
	fmt.Println("Main: Completed")
}

func getClipboardHistory() []string {
	result := getValues(10)
	for _, bytes := range result {
		log.Println(bytes)
	}
	return result
}
