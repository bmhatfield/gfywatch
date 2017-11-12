package notifications

import (
	"fmt"
	"log"

	"github.com/0xAX/notificator"
)

func init() {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon.png",
		AppName:     "GfyWatch",
	})
}

var notify *notificator.Notificator

// UploadStarted notifies that an upload has begun
func UploadStarted(filename string) {
	message := fmt.Sprintf("File '%s' detected. Uploading to Gfycat...", filename)
	notify.Push("Upload started!", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)
}

// UploadComplete notifies that an upload has completed
func UploadComplete(filename string, gfycatName string) {
	message := fmt.Sprintf("Completed uploading '%s' to Gfycat! It will be processed as '%s'", filename, gfycatName)
	notify.Push("Upload Complete", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)
}

// UploadError notifies that an upload has failed
func UploadError(err error) {
	message := fmt.Sprintf("Failed uploading to Gfycat: %s", err)
	notify.Push("Upload Complete", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)
}

// ProcessingComplete notifies that an upload has completed processing
func ProcessingComplete(gfycatName string) {
	message := fmt.Sprintf("Gfycat has completed processing '%s'", gfycatName)
	notify.Push("Processing Complete", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)
}
