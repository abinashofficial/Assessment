package utils

import (
	"Assessment/consts"
	"Assessment/loged"
	"Assessment/tapcontext"
	"errors"
	"fmt"
	"os"
	"reflect"
)

var supportedLang = []string{"en", "fr"}
var ErrorCodes = map[string]map[string]string{}

func ValidateUserEmail(email string) error {
	if email == "" {
		return fmt.Errorf("%s %q", consts.ErrUserEmailNotInHeader, "email")
	}
	return nil
}

func CheckKeyInSlice(strArray []string, key string) bool {
	if strArray == nil {
		return false
	}
	for _, val := range strArray {
		if val == key {
			return true
		}
	}
	return false
}

func GetError(msg string, lang string) string {
	if !CheckKeyInSlice(supportedLang, lang) {
		lang = "en"
	}
	return ErrorCodes[lang][msg]
}

func GetEnv(ctx tapcontext.TContext, key string, logValue ...bool) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		loged.GenericError(ctx, errors.New("key not found"), loged.FieldsMap{"key": key})
	}
	if len(logValue) > 0 {
		if logValue[0] {
			loged.GenericInfo(ctx, "Value Found", loged.FieldsMap{"key": key, "value": value})
		}
	}
	return value
}

// ContainsString tells whether a contains x.
func ContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// IsSubset tells whether array a is subset of array b or not
func IsSubset(a interface{}, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	s := reflect.ValueOf(b)
	m := make(map[interface{}]int)
	for i := 0; i < s.Len(); i++ {
		m[s.Index(i).String()]++
	}

	s = reflect.ValueOf(a)
	for i := 0; i < s.Len(); i++ {
		if m[s.Index(i).String()] > 0 {
			m[s.Index(i).String()]--
		} else {
			return false
		}
	}
	return true
}

func SplitNumericAndNonNumeric(input string) (numericPart, nonNumericPart string) {
	var numericChars []rune
	var nonNumericChars []rune

	for _, char := range input {
		if char >= '0' && char <= '9' {
			numericChars = append(numericChars, char)
		} else {
			nonNumericChars = append(nonNumericChars, char)
		}
	}

	numericPart = string(numericChars)
	nonNumericPart = string(nonNumericChars)
	return numericPart, nonNumericPart
}
