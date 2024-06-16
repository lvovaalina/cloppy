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

type ClipboardService struct {
	repo Repo
}

func NewClipboardService(repo Repo) ClipboardService {
	return ClipboardService{
		repo: repo,
	}
}

func (s *ClipboardService) StartClipboardService(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("Main: Starting worker")
	go s.watchText(ctx, &wg)

	wg.Add(1)
	go s.watchImg(ctx, &wg)

	wg.Wait()

	fmt.Println("Main: Waiting for worker to finish")
	fmt.Println("Main: Completed")
}

func (s *ClipboardService) GetClipboardHistory() []string {
	result := s.repo.GetValues(10)
	for _, bytes := range result {
		log.Println(bytes)
	}
	return result
}

func (s *ClipboardService) SetValueToClipboard(value string) {
	changed := clipboard.Write(clipboard.FmtText, []byte(value))
	log.Println(<-changed)
}

func (s *ClipboardService) watchImg(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ch := clipboard.Watch(ctx, clipboard.FmtImage)
	for data := range ch {
		s.serveFrames(data)
		fmt.Println(string(data))
	}
}

func (s *ClipboardService) watchText(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ch := clipboard.Watch(ctx, clipboard.FmtText)
	for data := range ch {
		s.repo.Save(data)
		fmt.Println("Add value: " + string(data))
	}
}

func (s *ClipboardService) serveFrames(imgByte []byte) {

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
