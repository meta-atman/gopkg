package mapping

import (
	"io"
)

// UnmarshalTomlBytes unmarshals TOML bytes into the given v.
func UnmarshalTomlBytes(content []byte, v any, opts ...UnmarshalOption) error {
	b, err := TomlToJson(content)
	if err != nil {
		return err
	}

	return UnmarshalJsonBytes(b, v, opts...)
}

// UnmarshalTomlReader unmarshals TOML from the given io.Reader into the given v.
func UnmarshalTomlReader(r io.Reader, v any, opts ...UnmarshalOption) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	return UnmarshalTomlBytes(b, v, opts...)
}
