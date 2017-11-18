package files

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// Watch opens a new FSNotify watch for the given path
func Watch(path string) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = watcher.Add(path)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return watcher
}
