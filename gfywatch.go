package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/bmhatfield/gfywatch/overwatch"

	"github.com/bmhatfield/gfywatch/files"
	"github.com/bmhatfield/gfywatch/gfycat"
	"github.com/bmhatfield/gfywatch/notifications"
	"github.com/fsnotify/fsnotify"
)

func contains(keywords []string, keyword string) bool {
	for _, word := range keywords {
		if word == keyword {
			return true
		}
	}

	return false
}

func tagsFromTitle(title string) []string {
	keywords := strings.Split(strings.Title(title), " ")
	tags := []string{"Overwatch", "Awesome"}

	heroTags := overwatch.TagsForHero(keywords)
	tags = append(tags, heroTags...)

	for index, word := range keywords {
		switch word {
		case "Solo", "Single", "Double", "Quadruple", "Quintuple", "Sextuple":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]))

		case "Triple", "3x":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]), "Three's a Crowd")

		case "1x", "2x", "4x", "5x", "6x":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]))

		case "Potg":
			tags = append(tags, "POTG")

		case "On", "The", "Go", "No":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]))

		case "Boop":
			tags = append(tags, "Boop", "Satisfying", "See You Next Fall")
		}
	}

	if !contains(keywords, "Potg") && !contains(keywords, "Highlight") {
		tags = append(tags, "Highlight")
	}

	return tags
}

func metadataFromFilename(filepath string) *gfycat.UploadFile {
	filename := path.Base(filepath)

	fileparts := strings.Split(filename, "_")
	title := strings.Title(fileparts[0])

	return &gfycat.UploadFile{
		Title:       title,
		Description: "Overwatch Play! Automatically uploaded by GfyWatch (github.com/bmhatfield/gfywatch)",
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
	watcher := files.Watch(filepath)
	defer watcher.Close()

	timeLastWritten := time.Now()
	waitDelay := time.Duration(5 * time.Second)

	log.Println("Waiting for writes to complete...")
	for timeLastWritten.Add(waitDelay).After(time.Now()) {
		select {
		case event := <-watcher.Events:
			if event.Op == fsnotify.Write {
				log.Println("Detected Write, extending timeout...")
				timeLastWritten = time.Now()
			}
		case err := <-watcher.Errors:
			if err != nil {
				log.Println("Error watching for events:", err)
			}
		}

		// Slow loop down to avoid hot-looping
		time.Sleep(100 * time.Millisecond)
	}

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
