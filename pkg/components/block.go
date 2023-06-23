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

var Block = donburi.NewComponentType[BlockData]()

type BlockType string

const (
	Block1    = "block1"
	Block2    = "block2"
	Obstacle1 = "obstacle1"
	Reverse   = "reverse"
)

func (b BlockType) String() string {
	return string(b)
}

func (b *BlockData) Init(bt BlockType) {
	b.BlockType = bt
	switch bt {
	case Block1, Block2:
		b.MaxDurability = 10
	case Obstacle1:
		b.MaxDurability = -1
	case Reverse:
		b.MaxDurability = 1
	}
	b.Durability = b.MaxDurability
}

func (b *BlockData) Damage(d int) {
	if b.IsBreakable() {
		b.Durability -= d
		if b.Durability < 0 {
			b.Durability = 0
		}
	}
}

func (b *BlockData) IsBroken() bool {
	return b.IsBreakable() && b.Durability == 0
}

func (b *BlockData) IsBreakable() bool {
	return b.MaxDurability > 0
}

func (b *BlockData) Width() float64 {
	switch b.BlockType {
	case Block2:
		return 64
	}
	return 32
}

func (b *BlockData) Height() float64 {
	return 32
}

func (b *BlockData) Hitboxes() []collision.Hitbox {
	return assets.GetHitboxes(b.BlockType.String())
}

func (b *BlockData) Animation() *ganim8.Animation {
	r := float64(b.Durability) / float64(b.MaxDurability)
	prefix := b.BlockType.String()
	switch {
	case r > 0.75:
		return assets.GetAnimation(prefix + "_1")
	case r > 0.5:
		return assets.GetAnimation(prefix + "_2")
	case r > 0.25:
		return assets.GetAnimation(prefix + "_3")
	default:
		return assets.GetAnimation(prefix + "_4")
	}
}

// RandomBlockType returns random block type.
func RandomBlockType() BlockType {
	r := rand.Float64()
	switch {
	case r < 0.05:
		return Reverse
	case r < 0.1:
		return Obstacle1
	}
	return BlockType(fmt.Sprintf("block%d", rand.Intn(2)+1))
}
