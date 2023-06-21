package sprites

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"revdriller/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

type Config struct {
	Sprites    []Sprite    `json:"sprites"`
	Animations []Animation `json:"animations"`
}

type Sprite struct {
	File string `json:"file"`
	W    int    `json:"w"`
	H    int    `json:"h"`
}

type Animation struct {
	File   string        `json:"file"`
	Name   string        `json:"name"`
	Frames []interface{} `json:"frame"`
	Flip   bool          `json:"flip"`
}

// Load loads all sprites.
func Load() error {
	cfg := &Config{}
	assets.MustReadJSON("config/sprites.json", cfg)

	// load sprites and animations
	for _, fn := range []func(*Config){
		loadSprites,
		loadAnimations,
	} {
		fn(cfg)
	}

	return nil
}

var (
	// grids is a map of grids by name.
	grids = map[string]*ganim8.Grid{}
	// iamge is a map of images by name.
	images = map[string]*ebiten.Image{}
	// sprites is a map of sprites by name.
	sprites = map[string]*ganim8.Sprite{}
	// animations is a map of animations by name.
	animations = map[string]*ganim8.Animation{}
)

// GetSprite returns a sprite by name.
func GetSprite(name string) *ganim8.Sprite {
	if _, ok := sprites[name]; !ok {
		panic(fmt.Sprintf("sprite not found: %s", name))
	}
	return sprites[name]
}

// GetAnimation returns an animation by name.
func GetAnimation(name string) *ganim8.Animation {
	if _, ok := animations[name]; !ok {
		panic(fmt.Sprintf("animation not found: %s", name))
	}
	return animations[name]
}

// loadSprites loads all sprites.
func loadSprites(cfg *Config) {
	for _, s := range cfg.Sprites {
		// load image from file
		b := assets.MustRead(s.File)
		// convert to ebiten.Image
		img := ebiten.NewImageFromImage(*decodeImage(&b))
		// add image to the map
		images[s.File] = img
		// get image size
		size := img.Bounds().Size()
		// create grid for the sprite
		g := ganim8.NewGrid(s.W, s.H, size.X, size.Y)
		// add grid to the map
		grids[s.File] = g
		// create sprite with the grid
		spr := ganim8.NewSprite(img, g.Frames())
		// add sprite to the map
		sprites[s.File] = spr
	}
}

// loadAnimations loads all animations.
func loadAnimations(cfg *Config) {
	for _, a := range cfg.Animations {
		g, ok := grids[a.File]
		if !ok {
			panic(fmt.Sprintf("grid not found: %s", a.File))
		}

		img, ok := images[a.File]
		if !ok {
			panic(fmt.Sprintf("image not found: %s", a.File))
		}

		// create sprite for the specified frames
		spr := ganim8.NewSprite(img, g.GetFrames(a.Frames...))

		// flip sprite if needed
		if a.Flip {
			spr.FlipH()
		}

		// create the animation
		anim := ganim8.NewAnimation(spr, time.Millisecond*20)
		animations[a.Name] = anim
	}
}

func decodeImage(rawImage *[]byte) *image.Image {
	img, _, _ := image.Decode(bytes.NewReader(*rawImage))
	return &img
}
