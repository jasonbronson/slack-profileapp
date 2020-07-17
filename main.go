package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

func main() {

	token := flag.String("token", "", "token api for slack")
	path := flag.String("path", "./pics", "Path for pictures to pull from")
	flag.Parse()
	log.Print("profileapp started")
	changeProfilePic(*token, *path)

}

func changeProfilePic(token string, path string) {
	api := slack.New(token, slack.OptionDebug(true))
	fmt.Println("Changed pic")
	userSetPhotoParams := slack.UserSetPhotoParams{}
	filename := getPics(path)
	err := api.SetUserPhoto(filename, userSetPhotoParams)
	if err != nil {
		log.Print(err)
	}
}

func getPics(root string) string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	//for _, file := range files {
	//fmt.Println("*****", shuffle(files))
	//}
	return shuffle(files)
}

func shuffle(src []string) string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	randomPic := strings.Join(final[:1], " ")
	fmt.Println(randomPic)
	return string(randomPic)
}
