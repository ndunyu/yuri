/*this will send  back the payload to the user*/

package yuri

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

type ResponseData struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"total_items"`
	//Pages      int `json:"pages"`
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
func JsonGzipBytesResponder(w http.ResponseWriter, r *http.Request, b []byte, err *ErrResponse) {
	t:=time.Now();
	defer  ExecutionTime(t,"gzip for bytes took about")
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gz.Write(b)
	///gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	///gzr.Write(b)
}

func JsonGzipResponder(w http.ResponseWriter, r *http.Request, item interface{}, err *ErrResponse) {
	// create header
	if err != nil {

		http.Error(w, err.StatusText, err.HTTPStatusCode)
		return
	}
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	// Gzip data
	gz := gzip.NewWriter(w)
	json.NewEncoder(gz).Encode(item)
	gz.Close()
}



func JsonResponder(w http.ResponseWriter, r *http.Request, item interface{}, err *ErrResponse) {
	if err != nil {

		http.Error(w, err.StatusText, err.HTTPStatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(item)

}

func Error404Or500(err error, w http.ResponseWriter, r *http.Request, text string) {
	if err != nil && err == pg.ErrNoRows {
		if text != "" {
			JsonResponder(w, r, nil, &ErrResponse{
				HTTPStatusCode: 404,
				Message:        text,
				StatusText:     text,
			})
			return

		}
		JsonResponder(w, r, nil, ErrNotFound)
		return
	}
	if err != nil {
		JsonResponder(w, r, nil, ErrInternalServerError)

		return

	}

}

func ValidationError(w http.ResponseWriter, r *http.Request, err []Field) {
	var errorResponse ErrorsResponse
	errorResponse.Errors = err
	data, _ := ToJson(errorResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)

	/*JsonResponder(w,r,data,&ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		Message:       "Invalid Request Body" ,
		StatusText:     data,
	})
	*/

}

type ErrorsResponse struct {
	Errors interface{} `json:"errors"`
}
