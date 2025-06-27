package mapping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const jsonTagKey = "json"

var jsonUnmarshaler = NewUnmarshaler(jsonTagKey)

func UnmarshalJson(data []byte, v any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(string(data), err)
	}

	return nil
}

// UnmarshalJsonBytes unmarshals content into v.
func UnmarshalJsonBytes(content []byte, v any, opts ...UnmarshalOption) error {
	return unmarshalJsonBytes(content, v, getJsonUnmarshaler(opts...))
}

// UnmarshalJsonMap unmarshals content from m into v.
func UnmarshalJsonMap(m map[string]any, v any, opts ...UnmarshalOption) error {
	return getJsonUnmarshaler(opts...).Unmarshal(m, v)
}

// UnmarshalJsonReader unmarshals content from reader into v.
func UnmarshalJsonReader(reader io.Reader, v any, opts ...UnmarshalOption) error {
	return unmarshalJsonReader(reader, v, getJsonUnmarshaler(opts...))
}

func UnmarshalJsonFromString(str string, v any) error {
	decoder := json.NewDecoder(strings.NewReader(str))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(str, err)
	}

	return nil
}

func UnmarshalJsonFromReader(reader io.Reader, v any) error {
	var buf strings.Builder
	teeReader := io.TeeReader(reader, &buf)
	decoder := json.NewDecoder(teeReader)
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(buf.String(), err)
	}

	return nil
}

func getJsonUnmarshaler(opts ...UnmarshalOption) *Unmarshaler {
	if len(opts) > 0 {
		return NewUnmarshaler(jsonTagKey, opts...)
	}

	return jsonUnmarshaler
}

func unmarshalJsonBytes(content []byte, v any, unmarshaler *Unmarshaler) error {
	var m any
	if err := json.Unmarshal(content, &m); err != nil {
		return err
	}

	return unmarshaler.Unmarshal(m, v)
}

func unmarshalJsonReader(reader io.Reader, v any, unmarshaler *Unmarshaler) error {
	var m any
	if err := UnmarshalJsonFromReader(reader, &m); err != nil {
		return err
	}

	return unmarshaler.Unmarshal(m, v)
}

func unmarshalUseNumber(decoder *json.Decoder, v any) error {
	decoder.UseNumber()
	return decoder.Decode(v)
}

func formatError(v string, err error) error {
	return fmt.Errorf("string: `%s`, error: `%w`", v, err)
}
