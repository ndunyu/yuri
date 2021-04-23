package yuri

import (
	"mime/multipart"
	"net/http"
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




