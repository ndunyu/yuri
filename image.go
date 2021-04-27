package yuri

import (
	"image"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func DownloadFile(URL, dir, prefix string) (string, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if !(response.StatusCode >= 200 && response.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(response.Body)

		return "", &RequestError{Url: URL, Message: string(b), StatusCode: response.StatusCode}

	}

	file, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return "", err
		////log.Fatal(err)
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}

func GetFileContentTypeWithExtension(paths string) (string, error) {
	fileType := mime.TypeByExtension(path.Ext(paths))
	if fileType == "" {
		f, err := os.Open(paths)
		if err != nil {
			return "", err
		}
		defer f.Close()
		fileType, err = GetFileContentType(f)

	}
	return fileType, nil

}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

///do all function involving images like resizing them

// GetMultiPartFileContentType  returns the content type of a file e.g
///image/png
func GetMultiPartFileContentType(file multipart.File) (string, error) {

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
func ResizeImage(images string, width, height int, dir, prefix string) (string, error) {
	reader, err := os.Open(images)
	if err != nil {
		return "", err

	}
	var dst image.Image
	defer reader.Close()
	src, _, err := image.Decode(reader)
	if err != nil {
		return "", err

	}
	//log.Println("name is ",name)
	dst = imaging.Resize(src, width, height, imaging.Lanczos)
	file, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return "", err

	}
	defer file.Close()
	err = imaging.Encode(file, dst, imaging.JPEG)
	if err != nil {
		return "", err

	}
	return file.Name(), nil

}
