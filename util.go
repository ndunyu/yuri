package yuri

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
)

func MakePin(min, max, n int) string {
	var pin []string
	for i := 1; i <= n; i++ {

		pin = append(pin, strconv.Itoa(random(min, max)))

	}

	return strings.Join(pin, "")

}

func random(min, max int) int {
	//rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

func makeHashId() {

}

func ExtractIntUrlParams(field string, r *http.Request) int {
	i, _ := strconv.Atoi(chi.URLParam(r, field))
	return i

}

func ExtractStringUrlParams(field string, r *http.Request) string {
	i := chi.URLParam(r, field)
	return i

}
