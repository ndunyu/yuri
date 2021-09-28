package yuri

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"github.com/ttacon/libphonenumber"
	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// IntStringOrFloatColumn /User may send data to a struct property
///as a string int or float,this will convert them to a string
///and save them to the database as a string (Var char)
//It will also make sure json does not complain
///that number int or float is not of type  string
type IntStringOrFloatColumn struct {
	sql.NullString
}

func (ni *IntStringOrFloatColumn) MarshalJSON() ([]byte, error) {

	if !ni.Valid {
		return json.Marshal(ni.String)
	}
	return json.Marshal(ni.String)
}

func (ni *IntStringOrFloatColumn) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		ni.Valid = false
		return err
	}
	switch v := item.(type) {
	case int:

		ni.String = IntToString(v)
		// v is an int here, so e.g. v + 1 is possible.

	case float64:
		ni.String = fmt.Sprintf("%.0f", v)

	case string:

		ni.String = v

	default:
		//fmt.Printf("%T\n", item)
		///log.Println("unknwn type  of is int", v)
		// And here I'm feeling dumb. ;)
	}
	///err := json.Unmarshal(b, &ni.String)
	///ni.Valid = err == nil
	ni.Valid = true
	return nil
}

const millisInSecond = 1000
const nsInSecond = 1000000

// MilliToTime FromUnixMilli Converts Unix Epoch from milliseconds to time.Time
func MilliToTime(ms int64) time.Time {
	return time.Unix(ms/int64(millisInSecond), (ms%int64(millisInSecond))*int64(nsInSecond))
}

//check if a string can be converted into an integer
func StringIsInt(s string) (*int, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, false
	}
	return &i, true
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
func ExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
func IsEmpty(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true

	}
	return false

}
func IntToString(item int) string {
	return strconv.Itoa(item)

}

func EqualIgnoreCase(a, b string) bool {
	return strings.EqualFold(a, b)

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

	data, _ := strconv.Atoi(MakePin(min, max, n))
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
	if err != nil {
		return false
	}
	return true
}

func PrintStruct(data interface{}) {
	fmt.Printf("%+v\n", data)

}
func MakeTimestamp() int64 {
	t := time.Now()
	tUnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
	return tUnixMilli
}

func InterfaceToStruct(input interface{}, outputStruct interface{}) error {

	if reflect.ValueOf(outputStruct).Kind() != reflect.Ptr {
		return errors.New("output must be a pointer")

	}
	toJson, err := ToJson(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(toJson, outputStruct)
	return err

}

func InterfaceToString(input interface{}) string {
	str := fmt.Sprintf("%v", input)
	return str

}
func StructToMap(input interface{}) (map[string]interface{}, error) {
	if reflect.ValueOf(input).Kind() != reflect.Ptr {
		return nil, errors.New("input must be a pointer")

	}
	var myInterface map[string]interface{}
	toJson, err := ToJson(input)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(toJson, &myInterface)

	return myInterface, err

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
	return b, nil
}

///TODO::make sure numbers dont pass 9
func checkKenyaInternationalPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`(^\+254)\d{9}$`)
	return re.MatchString(phone)
}

func FormatNumberToInternationalFormat(phoneNumber, region string) (string, error) {
	num, err := libphonenumber.Parse(phoneNumber, region)
	if err != nil {
		return "", err

	}
	formatted := libphonenumber.Format(num, libphonenumber.E164)
	return formatted, nil
}

func FormatNumberToNationalFormat(phoneNumber, region string) (string, error) {
	num, err := libphonenumber.Parse(phoneNumber, region)
	if err != nil {
		return "", err

	}
	formatted := libphonenumber.Format(num, libphonenumber.NATIONAL)
	trimmed := strings.Replace(formatted, " ", "", -1)
	return trimmed, nil
}

func CreateGid() string {

	u2 := uuid.NewV4()

	return u2.String()

}

func ContainsInt(items []int, item int) bool {
	for _, number := range items {
		if number == item {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CreateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
