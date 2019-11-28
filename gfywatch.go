package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/bmhatfield/gfywatch/files"
	"github.com/bmhatfield/gfywatch/gfycat"
	"github.com/bmhatfield/gfywatch/notifications"
	"github.com/bmhatfield/gfywatch/overwatch"

	"github.com/fsnotify/fsnotify"
)

func metadataFromFilename(filepath string) *gfycat.UploadFile {
	filename := path.Base(filepath)

	fileparts := strings.Split(filename, "_")
	title := strings.Title(fileparts[0])

	return &gfycat.UploadFile{
		Title:       title,
		Description: "Overwatch Play! Automatically uploaded by GfyWatch (github.com/bmhatfield/gfywatch)",
		Tags:        overwatch.TagsFromTitle(title),
		NoMd5:       true,
	}
}

func handleNewUpload(grant *gfycat.GFYClientGrant, filepath string) {
	waitForWriteComplete(filepath)

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

func waitForWriteComplete(filepath string) {
	watcher := files.Watch(path.Dir(filepath))
	defer watcher.Close()

	timer := time.NewTimer(300 * time.Second)

	log.Println("Waiting 5 minutes for rendering to complete...")
	for {
		select {
		case event := <-watcher.Events:
			if event.Op == fsnotify.Write && event.Name == filepath {
				log.Println("Detected final write of file, reducing wait...")
				timer.Reset(10 * time.Second)
			}

		case err := <-watcher.Errors:
			if err != nil {
				log.Println("Error watching for events:", err)
				return
			}

		case <-timer.C:
			log.Println("Finished waiting for writes")
			return
		}
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

		if status.Task == "encoding" && status.Progress != 0 {
			log.Printf("Encoding is %d%% complete", int(status.Progress*100))
		}
	}

	notifications.ProcessingComplete(current.GfyName)
}

func main() {
	grant, err := gfycat.NewClientGrantFromFile("grant.json")

	if err != nil {
		log.Printf("Unable to load credentials: %s\n", err)
		os.Exit(1)
	}

	watcher := files.Watch(".")
	defer watcher.Close()

	log.Println("Watching for new files...")

	tracker := files.NewTracker(5)

	for {
		select {
		case event := <-watcher.Events:
			if event.Op == fsnotify.Create && path.Ext(event.Name) == ".mp4" {
				if !tracker.In(event.Name) {
					tracker.Add(event.Name)
					log.Println("New file detected:", event.Name)
					go handleNewUpload(grant, event.Name)
				}
			}
		case err := <-watcher.Errors:
			if err != nil {
				log.Println("Error watching for events:", err)
			}
		}
	}
}
