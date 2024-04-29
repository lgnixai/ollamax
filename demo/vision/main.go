package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/lgnixai/ollamax"
	"github.com/ollama/ollama/api"
)

var (
	model = flag.String("model", "llava", "model to use")
)

//go:embed 2.jpg
var imageRaw []byte

func main() {
	flag.Parse()

	if err := ollamax.Init(); err != nil {
		log.Fatalln(err)
	}
	defer ollamax.Cleanup()

	o, err := ollamax.NewWithAutoDownload(*model)
	if err != nil {
		panic(err)
	}
	defer o.Close()

	resp, err := o.ChatStream(context.Background(),
		[]api.Message{{"user", "请用中文描述图片内容", []api.ImageData{imageRaw}}})
	if err != nil {
		panic(err)
	}
	for response := range resp {
		if response.Err != nil {
			fmt.Println(response.Err)
			return
		}
		fmt.Print(response.Result.Content)
	}
}
