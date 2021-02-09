package yuri

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ttacon/libphonenumber"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)




func IsEmpty(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true

	}
	return false



}
func JengaTime(t time.Time) string {
	layout := "2006-01-02"
	///t := time.Now()
	formatted := t.Format(layout)

	return formatted


}
func MakePin(min, max, n int) string {
	var pin []string
	for i := 1; i <= n; i++ {

		pin = append(pin, strconv.Itoa(random(min, max)))

	}

	return strings.Join(pin, "")

}

func MakeIntPin(min, max, n int) int {

	data,_:=strconv.Atoi(MakePin(min,max,n))
	return data

}




func CheckTokenId(r *http.Request, JwtSecretKey []byte) (int, error) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {

		return 0, errors.New("authorized")

	}
	splitToken := strings.Split(authHeader, "Bearer")
	authHeader = splitToken[1]
	sentToken := strings.TrimSpace(authHeader)


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
	if err!=nil {
		return false
	}
	return true
}


func PrintStruct(data interface{}){
	fmt.Printf("%+v\n", data)



}
func MakeTimestamp() int64 {
	t := time.Now()
	tUnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
	return tUnixMilli
}


func ToString(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}


func ToJson(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b,nil
}

///TODO::make sure numbers dont pass 9
func checkKenyaInternationalPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`(\+254)\d{9}$`)
	return re.MatchString(phone)
}


func FormatNumberToInternationalFormat(phoneNumber, region string) (string, error) {
	num, err := libphonenumber.Parse(phoneNumber, region)
	if err != nil {
		return "", err

	}
	formatted := libphonenumber.Format(num, libphonenumber.E164)
	return formatted,nil
}

func FormatNumberToNationalFormat(phoneNumber, region string) (string, error) {
	num, err := libphonenumber.Parse(phoneNumber, region)
	if err != nil {
		return "", err

	}
	formatted := libphonenumber.Format (num, libphonenumber.NATIONAL)
	trimmed := strings.Replace(formatted, " ", "", -1)
	return trimmed,nil
}

func CreateGid() string {

	u2 := uuid.NewV4()

	return u2.String()


}

//07