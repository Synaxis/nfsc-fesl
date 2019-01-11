package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

const (
	// charNull is a "null" character from ASCII table, it ends a packet
	charNull    byte = 0x00
	charEqual   byte = '='
	charNewLine byte = '\n'
)

const (
	tagName = "fesl"
)

type Encoder struct {
	wr EncWriter
}

type EncWriter interface {
	WriteString(string)
	WriteByte(byte)
	Len() int
	Bytes() []byte
}

type BufWriter struct {
	buf []byte
}

func (e *BufWriter) WriteString(s string) {
	e.buf = append(e.buf, []byte(s)...)
}

func (e *BufWriter) WriteByte(b byte) {
	e.buf = append(e.buf, b)
}

func (e *BufWriter) Len() int {
	return len(e.buf)
}

func (e *BufWriter) Bytes() []byte {
	return e.buf
}

func NewEncoder() *Encoder {
	return &Encoder{
		wr: &BufWriter{[]byte{}},
	}
}

func (e *Encoder) EncodePacket(packet *Answer) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	// Append type
	if _, err := buf.Write([]byte(packet.Type)); err != nil { // 4 bytes
		return nil, err
	}

	// Append status
	t := make([]byte, 4)
	binary.BigEndian.PutUint32(t, packet.PacketNumber)
	if _, err := buf.Write(t); err != nil {
		return nil, err
	}

	// Encode payload
	if err := e.Encode(packet.Payload); err != nil {
		return nil, err
	}

	// Append packet length
	c := make([]byte, 4)
	binary.BigEndian.PutUint32(c, uint32(e.wr.Len()+12))
	if _, err := buf.Write(c); err != nil {
		return nil, err
	}

	// Append payload
	if _, err := buf.Write(e.wr.Bytes()); err != nil {
		return nil, err
	}
	return buf, nil
}

func (e *Encoder) Encode(v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			if s, ok := r.(string); ok {
				panic(s)
			}
			err = r.(error)
		}
	}()
	e.enc("", reflect.ValueOf(v), nil)
	e.wr.WriteByte(charNull)
	return err
}

func (e *Encoder) set(k, v string) {
	if k == "" {
		return
	}

	e.wr.WriteString(k)
	e.wr.WriteByte(charEqual)
	e.wr.WriteString(v)
	e.wr.WriteByte(charNewLine)
}

func (e *Encoder) enc(key string, v reflect.Value, opt *EncOptions) {
	switch v.Kind() {
	case reflect.String:
		e.encString(key, v, opt)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		e.encInt(key, v, nil)
	case reflect.Bool:
		e.encBool(key, v, nil)
	case reflect.Float32, reflect.Float64:
		e.encFloat(key, v, nil)
	case reflect.Array, reflect.Slice:
		e.encArray(key, v, nil)
	case reflect.Map:
		e.encMap(key, v, nil)
	case reflect.Struct:
		e.encStruct(key, v, nil)
	case reflect.Interface:
		e.encInterface(key, v, nil)
	default:
		panic(fmt.Sprintf("codec: Not implemented type of %s value", v.Kind()))
	}
}

func (e *Encoder) encMap(key string, v reflect.Value, opt *EncOptions) {
	keys := v.MapKeys()
	e.set(key+".{}", strconv.Itoa(len(keys)))
	for _, k := range keys {
		if k.Type().Kind() != reflect.String {
			panic("codec: only maps with keys as strings are supported")
		}
		e.enc(
			fmt.Sprintf("%s.{%s}", key, k.String()),
			v.MapIndex(k),
			nil,
		)
	}
}

func (e *Encoder) encInterface(key string, v reflect.Value, opt *EncOptions) {
	if v.IsNil() {
		// Omit nil values
		return
	}
	e.enc(key, v.Elem(), nil)
}

func (e *Encoder) encArray(key string, v reflect.Value, opt *EncOptions) {
	e.set(key+".[]", strconv.Itoa(v.Len()))
	length := v.Len()
	for i := 0; i < length; i++ {
		e.enc(fmt.Sprintf("%s.%d", key, i), v.Index(i), nil)
	}
}

func (e *Encoder) encString(key string, v reflect.Value, opt *EncOptions) {
	val := v.String()
	if val == "" && opt != nil && opt.OmitEmpty {
		return
	}
	e.set(key, val)
}

func (e *Encoder) encBool(key string, v reflect.Value, opt *EncOptions) {
	if v.Bool() {
		e.set(key, "1")
	} else {
		e.set(key, "0")
	}
}

func (e *Encoder) encInt(key string, v reflect.Value, opt *EncOptions) {
	e.set(key, strconv.FormatInt(v.Int(), 10))
}

func (e *Encoder) encFloat(key string, v reflect.Value, opt *EncOptions) {
	e.set(key, fmt.Sprintf("%g", v.Float()))
}

func (e *Encoder) encStruct(key string, v reflect.Value, opt *EncOptions) {
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		switch value.Kind() {
		case reflect.Struct:
			e.encStruct(key, value, nil)
		default:
			e.encStructField(key, v.Type().Field(i), value)
		}
	}
}

func (e *Encoder) encStructField(prefix string, f reflect.StructField, vs reflect.Value) {
	tag := f.Tag.Get(tagName)

	switch true {
	case tag == "", tag == "-":
		// Not defined / Ignored
		return
	case unicode.IsLower(rune(f.Name[0])):
		// Unexported
		return
	}

	key := prefix
	if prefix != "" {
		key += "."
	}

	tagValue := strings.Split(tag, ",")
	key += tagValue[0]

	opt := EncOptions{}
	count := len(tagValue)
	for index := 1; index < count; index++ {
		switch tagValue[index] {
		case "omitempty":
			opt.OmitEmpty = true
		}
	}

	e.enc(key, vs, &opt)
}

type EncOptions struct {
	OmitEmpty bool
}
