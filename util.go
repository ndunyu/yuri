package yuri

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
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
