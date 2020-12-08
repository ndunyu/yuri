/*this will send  back the payload to the user*/

package yuri

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseData struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"total_items"`
	//Pages      int `json:"pages"`
}

func JsonResponder(w http.ResponseWriter, r *http.Request, item interface{}, err *ErrResponse) {
	if err != nil {
		log.Println("here")
		http.Error(w, err.StatusText, err.HTTPStatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(item)

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
