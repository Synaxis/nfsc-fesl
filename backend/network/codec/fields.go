package codec

import (
	"fmt"
	"strconv"
	"strings"
)

// Fields contains payload of FESL response converted as a go-lang friendly map
type Fields map[string]string

// Get returns string value
// Note: String() in somehow reserved in golang (like toString() in C#)
func (m Fields) Get(key string) string {
	return m[key]
}

// Exists checks if specified key was defined
func (m Fields) Exists(key string) bool {
	_, ok := m[key]
	return ok
}

// IntVal tries to cast specified value as an integer
func (m Fields) IntVal(key string) (int, error) {
	return strconv.Atoi(m.Get(key))
}

// FloatVal tries to cast specified value as a float
func (m Fields) FloatVal(key string) (float64, error) {
	return strconv.ParseFloat(m.Get(key), 32)
}

// FloatAsInt attempts to guess precision and cast float as an integer.
// It allows to take care of floating point math inaccuracy (0.1+0.2 != 0.3)
func (m Fields) FloatAsInt(key string) (int, int, error) {
	val := m.Get(key)

	fidx := strings.IndexByte(val, '.')
	if fidx == -1 {
		// TODO: error might be helpful here
		return 0, 0, nil
	}

	norm := strings.Replace(val, ".", "", 1)
	i, err := strconv.Atoi(norm)
	if err != nil {
		return 0, 0, err
	}

	prec := len(val) - fidx
	if prec != 0 {
		prec--
	}

	return i, prec, nil
}

// IntArr explodes an encoded array which is separated by commas
func (m Fields) IntArr(key, sep string) []int {
	strVals := strings.Split(m.Get(key), sep)
	intVals := []int{}
	for _, s := range strVals {
		// TODO: how about triming whitespaces?
		i, _ := strconv.Atoi(strings.Trim(s, `"`))
		intVals = append(intVals, i)
	}
	return intVals
}

// StrArr explodes an encoded array which is separated by semicolons
func (m Fields) StrArr(key, sep string) []string {
	// TODO: how about triming whitespaces?
	return strings.Split(m.Get(key), sep)
}

// ArrayStrings scraps all items in specified key as a array
func (m Fields) ArrayStrings(prefix string) []string {
	count, _ := strconv.Atoi(m.Get(prefix + ".[]"))
	arr := make([]string, count)
	for i := 0; i < count; i++ {
		arr[i] = m.Get(fmt.Sprintf("%s.%d", prefix, i))
	}
	return arr
}
