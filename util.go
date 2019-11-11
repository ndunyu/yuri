package yuri

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

func MakePin(min, max, n int) string {
	var pin []string
	for i := 1; i <= n; i++ {

		pin = append(pin, strconv.Itoa(random(min, max)))

	}

	return strings.Join(pin, "")

}

func CheckTokenId(r *http.Request, JwtSecretKey []byte) (int, error) {

	authHeader := r.Header.Get("Authorization")
	log.Println("below")
	if authHeader == "" {

		return 0, errors.New("authorized")

	}
	splitToken := strings.Split(authHeader, "Bearer")
	authHeader = splitToken[1]
	sentToken := strings.TrimSpace(authHeader)
	log.Println(sentToken)

	if sentToken == "" {

		return 0, errors.New("authorized")

	}

	userId, err := ValidateToken(sentToken, JwtSecretKey)
	if err != nil {
		return 0, errors.New("unauthorized")

	}
	return userId, nil

}

func ValidateToken(tok string, JwtSecretKey []byte) (int, error) {
	var userId int

	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return JwtSecretKey, nil

	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := int(claims["user_id"].(float64)) //.(uint)
		return userId, err

	} else {
		err = new("Wrong User Credentials sent")

		return userId, err

	}

}

func CreateToken(claims jwt.MapClaims, JwtSecretKey []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {

	}

	return tokenString

}

type WrongTokenCredential struct {
	s string
}

func (e *WrongTokenCredential) Error() string {
	return e.s
}

func new(text string) error {
	return &WrongTokenCredential{text}
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
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
