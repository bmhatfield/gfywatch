package notifications

import (
	"fmt"

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
}

// UploadComplete notifies that an upload has completed
func UploadComplete(filename string) {
	message := fmt.Sprintf("Completed uploading '%s' to Gfycat!", filename)
	notify.Push("Upload Complete", message, "icon.png", notificator.UR_NORMAL)
}

// ProcessingComplete notifies that an upload has completed processing
func ProcessingComplete(filename string) {
	message := fmt.Sprintf("Gfycat has completed processing '%s'", filename)
	notify.Push("Processing Complete", message, "icon.png", notificator.UR_NORMAL)
}
