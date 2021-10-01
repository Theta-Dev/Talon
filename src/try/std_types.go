package try

// gotry auto-generated type definitions. DO NOT EDIT.

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

// Bool is a helper method to handle errors of func() (bool, error) functions
func Bool(v bool, err error) (bool) {
	Check(err)
	return v
}

// Bools is a helper method to handle errors of func() ([]bool, error) functions
func Bools(v []bool, err error) ([]bool) {
	Check(err)
	return v
}

// Byte is a helper method to handle errors of func() (byte, error) functions
func Byte(v byte, err error) (byte) {
	Check(err)
	return v
}

// Bytes is a helper method to handle errors of func() ([]byte, error) functions
func Bytes(v []byte, err error) ([]byte) {
	Check(err)
	return v
}

// Rune is a helper method to handle errors of func() (rune, error) functions
func Rune(v rune, err error) (rune) {
	Check(err)
	return v
}

// Runes is a helper method to handle errors of func() ([]rune, error) functions
func Runes(v []rune, err error) ([]rune) {
	Check(err)
	return v
}

// String is a helper method to handle errors of func() (string, error) functions
func String(v string, err error) (string) {
	Check(err)
	return v
}

// Strings is a helper method to handle errors of func() ([]string, error) functions
func Strings(v []string, err error) ([]string) {
	Check(err)
	return v
}

// Int is a helper method to handle errors of func() (int, error) functions
func Int(v int, err error) (int) {
	Check(err)
	return v
}

// Ints is a helper method to handle errors of func() ([]int, error) functions
func Ints(v []int, err error) ([]int) {
	Check(err)
	return v
}

// Int8 is a helper method to handle errors of func() (int8, error) functions
func Int8(v int8, err error) (int8) {
	Check(err)
	return v
}

// Int8s is a helper method to handle errors of func() ([]int8, error) functions
func Int8s(v []int8, err error) ([]int8) {
	Check(err)
	return v
}

// Int16 is a helper method to handle errors of func() (int16, error) functions
func Int16(v int16, err error) (int16) {
	Check(err)
	return v
}

// Int16s is a helper method to handle errors of func() ([]int16, error) functions
func Int16s(v []int16, err error) ([]int16) {
	Check(err)
	return v
}

// Int32 is a helper method to handle errors of func() (int32, error) functions
func Int32(v int32, err error) (int32) {
	Check(err)
	return v
}

// Int32s is a helper method to handle errors of func() ([]int32, error) functions
func Int32s(v []int32, err error) ([]int32) {
	Check(err)
	return v
}

// Int64 is a helper method to handle errors of func() (int64, error) functions
func Int64(v int64, err error) (int64) {
	Check(err)
	return v
}

// Int64s is a helper method to handle errors of func() ([]int64, error) functions
func Int64s(v []int64, err error) ([]int64) {
	Check(err)
	return v
}

// Uint is a helper method to handle errors of func() (uint, error) functions
func Uint(v uint, err error) (uint) {
	Check(err)
	return v
}

// Uints is a helper method to handle errors of func() ([]uint, error) functions
func Uints(v []uint, err error) ([]uint) {
	Check(err)
	return v
}

// Uint8 is a helper method to handle errors of func() (uint8, error) functions
func Uint8(v uint8, err error) (uint8) {
	Check(err)
	return v
}

// Uint8s is a helper method to handle errors of func() ([]uint8, error) functions
func Uint8s(v []uint8, err error) ([]uint8) {
	Check(err)
	return v
}

// Uint16 is a helper method to handle errors of func() (uint16, error) functions
func Uint16(v uint16, err error) (uint16) {
	Check(err)
	return v
}

// Uint16s is a helper method to handle errors of func() ([]uint16, error) functions
func Uint16s(v []uint16, err error) ([]uint16) {
	Check(err)
	return v
}

// Uint32 is a helper method to handle errors of func() (uint32, error) functions
func Uint32(v uint32, err error) (uint32) {
	Check(err)
	return v
}

// Uint32s is a helper method to handle errors of func() ([]uint32, error) functions
func Uint32s(v []uint32, err error) ([]uint32) {
	Check(err)
	return v
}

// Uint64 is a helper method to handle errors of func() (uint64, error) functions
func Uint64(v uint64, err error) (uint64) {
	Check(err)
	return v
}

// Uint64s is a helper method to handle errors of func() ([]uint64, error) functions
func Uint64s(v []uint64, err error) ([]uint64) {
	Check(err)
	return v
}

// Float32 is a helper method to handle errors of func() (float32, error) functions
func Float32(v float32, err error) (float32) {
	Check(err)
	return v
}

// Float32s is a helper method to handle errors of func() ([]float32, error) functions
func Float32s(v []float32, err error) ([]float32) {
	Check(err)
	return v
}

// Float64 is a helper method to handle errors of func() (float64, error) functions
func Float64(v float64, err error) (float64) {
	Check(err)
	return v
}

// Float64s is a helper method to handle errors of func() ([]float64, error) functions
func Float64s(v []float64, err error) ([]float64) {
	Check(err)
	return v
}

// StrStr is a helper method to handle errors of func() (string, string, error) functions
func StrStr(v string, v1 string, err error) (string, string) {
	Check(err)
	return v, v1
}

// X is a helper method to handle errors of func() (interface{}, error) functions
func X(v interface{}, err error) (interface{}) {
	Check(err)
	return v
}

// File is a helper method to handle errors of func() (*os.File, error) functions
func File(v *os.File, err error) (*os.File) {
	Check(err)
	return v
}

// Reader is a helper method to handle errors of func() (io.Reader, error) functions
func Reader(v io.Reader, err error) (io.Reader) {
	Check(err)
	return v
}

// Writer is a helper method to handle errors of func() (io.Writer, error) functions
func Writer(v io.Writer, err error) (io.Writer) {
	Check(err)
	return v
}

// Request is a helper method to handle errors of func() (*http.Request, error) functions
func Request(v *http.Request, err error) (*http.Request) {
	Check(err)
	return v
}

// Response is a helper method to handle errors of func() (*http.Response, error) functions
func Response(v *http.Response, err error) (*http.Response) {
	Check(err)
	return v
}

// Url is a helper method to handle errors of func() (*url.URL, error) functions
func Url(v *url.URL, err error) (*url.URL) {
	Check(err)
	return v
}
