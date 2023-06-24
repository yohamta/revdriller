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
	Type          BlockType
}

var Block = donburi.NewComponentType[BlockData]()

type BlockType string

const (
	BlockTypeNormal     = "block1"
	BlockTypeNormal2    = "block2"
	BlockTypeObstacle1  = "obstacle1"
	BlockTypeObstackle2 = "obstacle2"
	BlockTypeReverse    = "reverse"
)

func (b BlockType) Reverse() BlockType {
	switch b {
	case BlockTypeNormal:
		return BlockTypeObstacle1
	case BlockTypeNormal2:
		return BlockTypeObstackle2
	case BlockTypeObstacle1:
		return BlockTypeNormal
	case BlockTypeObstackle2:
		return BlockTypeNormal2
	}
	return b // other blocks are not reversed
}

func (b BlockType) String() string {
	return string(b)
}

func (b *BlockData) Init(bt BlockType) {
	b.Type = bt
	switch bt {
	case BlockTypeNormal, BlockTypeNormal2:
		b.MaxDurability = 10
	case BlockTypeObstacle1, BlockTypeObstackle2:
		b.MaxDurability = -1
	case BlockTypeReverse:
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
	switch b.Type {
	case BlockTypeNormal2, BlockTypeObstackle2:
		return 64
	}
	return 32
}

func (b *BlockData) Height() float64 {
	return 32
}

func (b *BlockData) Hitboxes() []collision.Hitbox {
	return assets.GetHitboxes(b.Type.String())
}

func (b *BlockData) Animation() *ganim8.Animation {
	r := float64(b.Durability) / float64(b.MaxDurability)
	prefix := b.Type.String()
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
	case r < 0.1:
		return BlockTypeReverse
	case r < 0.5:
		return BlockType(fmt.Sprintf("obstacle%d", rand.Intn(2)+1))
	}
	return BlockType(fmt.Sprintf("block%d", rand.Intn(2)+1))
}
