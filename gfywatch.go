package main

import (
	"fmt"
	"log"
	"path"

	"github.com/bmhatfield/gfywatch/files"
	"github.com/bmhatfield/gfywatch/notifications"
	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher := files.WatchForNew("/Users/bhatfield/Documents/go/src/github.com/bmhatfield/gfywatch")
	defer watcher.Close()

	for {
		select {
		case event := <-watcher.Events:
			if event.Op == fsnotify.Create {
				log.Println(fmt.Sprintf("event: %+v", event))
				notifications.UploadStarted(path.Base(event.Name))
			}
		case err := <-watcher.Errors:
			if err != nil {
				log.Println("error:", err)
			}
		}
	}
}
