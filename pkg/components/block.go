package components

import (
	"fmt"
	"math/rand"
	"revdriller/assets"
	"revdriller/pkg/collision"

	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type BlockData struct {
	Durability    int
	MaxDurability int
	BlockType     BlockType
}

type BlockType int

const (
	_ BlockType = iota

	BlockType1 // 32x32 block
	BlockType2 // 64x32 block
)

var Block = donburi.NewComponentType[BlockData]()

func (b *BlockData) Damage(d int) {
	b.Durability -= d
	if b.Durability < 0 {
		b.Durability = 0
	}
}

func (b *BlockData) IsBreakable() bool {
	return b.Durability > 0
}

func (b *BlockData) prefix() string {
	return fmt.Sprintf("block%d", b.BlockType)
}

func (b *BlockData) Width() float64 {
	switch b.BlockType {
	case BlockType1:
		return 32
	case BlockType2:
		return 64
	}
	panic("invalid block type")
}

func (b *BlockData) Height() float64 {
	return 32
}

func (b *BlockData) Hitboxes() []collision.Hitbox {
	return assets.GetHitboxes(b.prefix())
}

func (b *BlockData) Animation() *ganim8.Animation {
	r := float64(b.Durability) / float64(b.MaxDurability)
	switch {
	case r > 0.75:
		return assets.GetAnimation(b.prefix() + "_1")
	case r > 0.5:
		return assets.GetAnimation(b.prefix() + "_2")
	case r > 0.25:
		return assets.GetAnimation(b.prefix() + "_3")
	default:
		return assets.GetAnimation(b.prefix() + "_4")
	}
}

// RandomBlockType returns random block type.
func RandomBlockType() BlockType {
	return BlockType(rand.Intn(2) + 1)
}
