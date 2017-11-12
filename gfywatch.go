package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/bmhatfield/gfywatch/files"
	"github.com/bmhatfield/gfywatch/gfycat"
	"github.com/bmhatfield/gfywatch/notifications"
	"github.com/fsnotify/fsnotify"
)

func loadClientConfig(path string) (*gfycat.GFYClientGrant, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	grant := &gfycat.GFYClientGrant{}

	err = json.Unmarshal(file, grant)
	if err != nil {
		return nil, err
	}

	return grant, nil
}

func tagsFromTitle(title string) []string {
	tags := strings.Split(title, " ")

	return tags
}

func metadataFromFilename(filepath string) *gfycat.UploadFile {
	filename := "lucio double boop potg_17-11-08_23-38-28.mp4"

	fileparts := strings.Split(filename, "_")
	title := strings.Title(fileparts[0])

	return &gfycat.UploadFile{
		Title:       title,
		Description: "Overwatch POTG",
		Tags:        tagsFromTitle(title),
		NoMd5:       true,
	}
}

func watchUploadStatus(client *gfycat.GFYClient, current *gfycat.FileDrop) {
	status, err := client.UploadStatus(current)
	if err != nil {
		fmt.Println("Error checking for upload status: ", err)
		return
	}

	for status.Time != 0 {
		time.Sleep(time.Duration(status.Time) * time.Second)
		status, err = client.UploadStatus(current)
		if err != nil {
			fmt.Println("Error checking for upload status: ", err)
			return
		}
	}

	notifications.ProcessingComplete(current.GfyName)
}

func handleNewUpload(grant *gfycat.GFYClientGrant, filepath string) {
	client, err := gfycat.NewGFYClient(grant)

	if err != nil {
		fmt.Println(err)
	}

	metadata := metadataFromFilename(filepath)

	notifications.UploadStarted(path.Base(filepath))
	current, err := client.LocalUpload(filepath, metadata)

	if err != nil {
		fmt.Println(err)
		notifications.UploadError(err)
	}

	notifications.UploadComplete(path.Base(filepath), current.GfyName)

	go watchUploadStatus(client, current)
}

func main() {
	grant, err := loadClientConfig("grant.json")

	if err != nil {
		log.Printf("Unable to load credentials: %s\n", err)
		os.Exit(1)
	}

	watcher := files.WatchForNew("/Users/bhatfield/Desktop")
	defer watcher.Close()

	fmt.Println("Watching for new files...")
	for {
		select {
		case event := <-watcher.Events:
			if event.Op == fsnotify.Create {
				log.Println(fmt.Sprintf("New file found: %s", event.Name))

				handleNewUpload(grant, event.Name)
			}
		case err := <-watcher.Errors:
			if err != nil {
				log.Println("Error watching for events:", err)
			}
		}
	}
}
