/*this package will do all the operations regarding

the http request methods*/
package yuri

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Pagination struct {
	Page   int
	Size   int
	Offset int
	Max    int
}

type Filters struct {
	Column string
	Value  interface{}
}

func FormFileNoName(r *http.Request, filename, path, url string) (string, *ErrResponse) {
	_, image, err := ReadRequestFileNoName(r, filename, path, url)
	if err != nil {
		log.Println(err)
		return "", ErrInvalidRequest
	}

	return image, nil

}
func DynamicData(r *http.Request, dataname string) ([]byte, error) {

	data := r.FormValue(dataname)

	return []byte(data), nil
}

func ReadRequestFileNoName(r *http.Request, filename string, storagePath string, BaseUrl string) (string, string, error) {
	_ = r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(filename)

	//ex, err := os.Executable()
	if err != nil {
		log.Println(err)
		return "", "", err

	}
	defer file.Close()
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		_ = os.MkdirAll(storagePath, os.ModePerm)
	}

	////imageName := uuid.NewV4().String() + filepath.Ext(handler.Filename)
	imageName := handler.Filename

	imagePath := storagePath + imageName

	f, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", "", err

	}

	result := strings.Replace(imagePath, storagePath, "", -1)

	imageUrl := BaseUrl + result
	//image_name =   handler.Filename
	defer f.Close()
	_, _ = io.Copy(f, file)
	return imageName, imageUrl, nil

}

////this will get any image sent with
///sent in the body of a request given
///the imazge key
func FormFile(r *http.Request, filename, path, url string) (string, *ErrResponse) {
	_, image, err := ReadRequestFile(r, filename, path, url)
	if err != nil {
		log.Println(err)
		return "", ErrInvalidRequest
	}

	return image, nil

}

///func UploadAndResize(r *http.Request, filename,path, url string)(string, *ErrResponse){

//}

///////get the file from the data
/////filename is the key the formfile was sent with
////path is where you want to store the file
////base url is the base http you want to be accesssing the file with
////example of baseUrl="https://ndunyu.co.ke/images"
func ReadRequestFile(r *http.Request, filename string, storagePath string, BaseUrl string) (string, string, error) {
	_ = r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(filename)

	//ex, err := os.Executable()
	if err != nil {
		log.Println(err)
		return "", "", err

	}
	defer file.Close()
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		_ = os.MkdirAll(storagePath, os.ModePerm)
	}

	imageName := uuid.NewV4().String() + filepath.Ext(handler.Filename)

	imagePath := storagePath + imageName

	f, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", "", err

	}

	result := strings.Replace(imagePath, storagePath, "", -1)

	imageUrl := BaseUrl + result
	//image_name =   handler.Filename
	defer f.Close()
	_, _ = io.Copy(f, file)
	return imageName, imageUrl, nil

}

//takes in a pointer and reads to it the request body sent
func RequestBody(r *http.Request, item interface{}) *ErrResponse {
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		return ErrInvalidRequest

	}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		///sentry.CaptureException(err)
		return ErrInvalidRequest
	}
	defer r.Body.Close()
	return nil

}

//RequestBodyToFromJson is to replace [RequestBody to return ]
//a go error instead of http error
func RequestBodyToFromJson(r *http.Request, item interface{}) error {
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		return errors.New("A pointer is required")

	}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		///sentry.CaptureException(err)
		return err
	}
	defer r.Body.Close()

	return nil
}

//this will get the size
//and the page from request if they exist
//else use the default
func (p *Pagination) GetPagination(r *http.Request) {
	if r.URL.Query().Get("page") != "" {
		p.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		//p.Page = page

	}

	if r.URL.Query().Get("size") != "" {
		p.Size, _ = strconv.Atoi(r.URL.Query().Get("size"))
		///p.Size = size

	}

	p.ProcessPagination()


}

func (p *Pagination) ProcessPagination(){
	if p.Max == 0 {
		p.Max = 40

	}
	if p.Size > p.Max {
		p.Size = p.Max
	}
	if p.Page > 0 {
		p.Offset = (p.Page - 1) * p.Size

	} else {
		p.Offset = 0
	}
}

func GetQueryParameter(key string, r *http.Request) string {

	return r.URL.Query().Get(key)

}
