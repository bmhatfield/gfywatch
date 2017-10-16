package files

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// WatchForNew tracks a directory for new files
func WatchForNew(path string) *fsnotify.Watcher {
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
