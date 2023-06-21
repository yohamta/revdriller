package assets

import (
	"embed"
	"encoding/json"
)

//go:embed *
var fs embed.FS

func MustRead(name string) []byte {
	b, err := fs.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return b
}

func MustReadJSON(name string, v interface{}) {
	b := MustRead(name)
	if err := json.Unmarshal(b, v); err != nil {
		panic(err)
	}
}
