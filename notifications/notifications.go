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
	message := fmt.Sprintf("File '%s' ready. Uploading to Gfycat...", filename)
	err := notify.Push("Upload started!", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)

	if err != nil {
		log.Println("Error publishing notification: ", err)
	}
}

// UploadComplete notifies that an upload has completed
func UploadComplete(filename string, gfycatName string) {
	message := fmt.Sprintf("Completed uploading '%s' to Gfycat! It will be processed as '%s'", filename, gfycatName)
	err := notify.Push("Upload Complete", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)

	if err != nil {
		log.Println("Error publishing notification: ", err)
	}
}

// UploadError notifies that an upload has failed
func UploadError(err error) {
	message := fmt.Sprintf("Failed uploading to Gfycat: %s", err)
	perr := notify.Push("Upload Complete", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)

	if perr != nil {
		log.Println("Error publishing notification: ", perr)
	}
}

// ProcessingComplete notifies that an upload has completed processing
func ProcessingComplete(gfycatName string) {
	message := fmt.Sprintf("Gfycat has completed processing '%s'", gfycatName)
	err := notify.Push("Processing Complete", message, "icon.png", notificator.UR_NORMAL)
	log.Println(message)

	if err != nil {
		log.Println("Error publishing notification: ", err)
	}
}
