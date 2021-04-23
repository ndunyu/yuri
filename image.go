package yuri

import (
	"image"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

///do all function involving images like resizing them

// GetFileContentType returns the content type of a file e.g
///image/png
func GetFileContentType(file multipart.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.

	fileHeader := make([]byte, 512)
	if _, err := file.Read(fileHeader); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(fileHeader)

	return contentType, nil
}

func GetFileExtension(name string) string {
	extension := filepath.Ext(name)

	return extension
}

func CreateATempFile(name string, file io.Reader) (*os.File, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), name)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(tempFile, file)
	return tempFile, err

}

// ResizeImage pass width to either height or width to maintain aspect ratio
func ResizeImage(images *os.File, width, height int) (*image.Image, error) {
	reader, err := os.Open(images.Name())
	var dst image.Image
	defer reader.Close()
	src, _, err := image.Decode(reader)
	if err != nil {
		return &dst, err

	}
	//log.Println("name is ",name)
	dst = imaging.Resize(src, width, height, imaging.Lanczos)
	return &dst, nil

}
