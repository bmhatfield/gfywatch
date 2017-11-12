package gfycat

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// FileDrop represents an inflight upload
type FileDrop struct {
	GfyName string `json:"gfyname"`
	Secret  string `json:"secret"`
}

// UploadStatus represents an inflight upload status response
type UploadStatus struct {
	// While encoding or not found
	Task string `json:"task"`
	Time int    `json:"time"`

	// When complete
	GfyName string `json:"gfyname"`
}

// Caption represents in-gif captions
type Caption struct {
	Text               string  `json:"text"`
	StartSeconds       int     `json:"startSeconds"`
	Duration           int     `json:"duration"`
	FontHeight         int     `json:"fontHeight"`
	X                  int     `json:"x"`
	Y                  int     `json:"y"`
	FontHeightRelative float32 `json:"fontHeightRelative"`
	XRelatice          float32 `json:"xRelative"`
	YRelative          float32 `json:"yRelative"`
}

// Cut represents the total duration of the processed video, and where to start trimming from
type Cut struct {
	Duration int `json:"duration"`
	Start    int `json:"start"`
}

// Crop represents the dimensions of the reduced GFYCat
type Crop struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// UploadFile represents the structure of an UploadFile request
type UploadFile struct {
	FetchURL     string    `json:"fetchUrl,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	Tags         []string  `json:"tags,omitempty"`
	NoMd5        bool      `json:"noMd5,omitempty"`
	Private      int       `json:"private,omitempty"`
	Nsfw         int       `json:"nsfw,omitempty"`
	FetchSeconds int       `json:"fetchSeconds,omitempty"`
	FetchMinutes int       `json:"fetchMinutes,omitempty"`
	FetchHours   int       `json:"fetchHours,omitempty"`
	Captions     []Caption `json:"captions,omitempty"`
	Cut          []Cut     `json:"cut,omitempty"`
	Crop         []Crop    `json:"crop,omitempty"`
}

// LocalUpload uploads a local file to GFYCat
func (gfy *GFYClient) LocalUpload(Filepath string, Upload *UploadFile) (*FileDrop, error) {
	body, err := ToJSON(Upload)
	if err != nil {
		return nil, err
	}

	gfycats := fmt.Sprintf("%s/%s", BaseURI, "gfycats")
	resp, err := gfy.request("POST", gfycats, body)
	if err != nil {
		return nil, err
	}

	activeUpload := &FileDrop{}
	apiresp := APIResponse(*resp)
	err = apiresp.ToType(activeUpload)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	err = writer.WriteField("key", activeUpload.GfyName)
	if err != nil {
		return nil, err
	}

	part, err := writer.CreateFormFile("file", activeUpload.GfyName)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(Filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	written, err := io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("wrote %v bytes but hit error: %s", written, err)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://filedrop.gfycat.com", buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err = gfy.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("upload completed with unexpected response code (%d): %s", resp.StatusCode, responseBody)
	}

	return activeUpload, nil
}

// UploadStatus provides updates on the processing status of a file
func (gfy *GFYClient) UploadStatus(activeUpload *FileDrop) (*UploadStatus, error) {
	path := fmt.Sprintf("%s/%s/%s", BaseURI, "gfycats/fetch/status", activeUpload.GfyName)
	resp, err := gfy.request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	status := &UploadStatus{}
	apiresp := APIResponse(*resp)
	err = apiresp.ToType(status)
	if err != nil {
		return nil, err
	}

	return status, nil
}
