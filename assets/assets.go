package assets

import (
	"embed"
	"encoding/json"
)

// Load loads all assets.
func Load() error {
	sprites := &SpriteConfig{}
	mustReadJSON("config/sprites.json", sprites)

	// Load sprites and animations.
	for _, fn := range []func(*SpriteConfig){
		loadSprites,
		loadAnimations,
	} {
		fn(sprites)
	}

	return nil
}

//go:embed *
var fs embed.FS

func mustRead(name string) []byte {
	b, err := fs.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return b
}

func mustReadJSON(name string, v interface{}) {
	b := mustRead(name)
	if err := json.Unmarshal(b, v); err != nil {
		panic(err)
	}
}
