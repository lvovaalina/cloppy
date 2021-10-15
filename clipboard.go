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

var clipboardHistory [][]byte

func watchImg(wg sync.WaitGroup) {
	defer wg.Done()
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := clipboard.Watch(context, clipboard.FmtImage)
	for data := range ch {
		setClipboardValue(data)
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
		setClipboardValue(data)
		fmt.Println(string(data))
	}
}

func setClipboardValue(value []byte) {
	clipboardHistory = append(clipboardHistory, value)
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
	clipboardHistory = make([][]byte, 0)
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
	result := make([]string, 0)
	for _, bytes := range clipboardHistory {
		result = append(result, string(bytes))
	}
	return result
}
