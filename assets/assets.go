package assets

import (
	"embed"
	"encoding/json"
)

// Load loads all assets.
func Load() error {

	// Load sprites and animations.
	sprites := &spriteConfig{}
	mustReadJSON("config/sprites.json", sprites)

	loadSprites(sprites)
	loadAnimations(sprites)

	// load hitboxes
	hitboxes := &hitboxConfig{}
	mustReadJSON("config/hitboxes.json", hitboxes)
	loadHitboxes(hitboxes)

	return nil
}

//go:embed img/*.png config/*.json
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
