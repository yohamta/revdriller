package components

import (
	"fmt"
	"math/rand"
	"revdriller/assets"
	"revdriller/pkg/collision"
	"revdriller/pkg/consts"

	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type BlockData struct {
	Type          BlockType
	Durability    int
	MaxDurability int
	Invalidated   bool
	Score         int
}

var Block = donburi.NewComponentType[BlockData]()

type BlockType string

const (
	BlockTypeNormal     BlockType = "block1"
	BlockTypeNormal2    BlockType = "block2"
	BlockTypeObstacle1  BlockType = "obstacle1"
	BlockTypeObstackle2 BlockType = "obstacle2"
	BlockTypeReverse    BlockType = "reverse"
	BlockTypeBomb       BlockType = "bomb"
	BlockTypeNeedle     BlockType = "needle"
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

func (b BlockType) Shorten() BlockType {
	switch b {
	case BlockTypeNormal2:
		return BlockTypeNormal
	case BlockTypeObstackle2:
		return BlockTypeObstacle1
	}
	return b
}

func (b BlockType) IsLong() bool {
	return b == BlockTypeNormal2 || b == BlockTypeObstackle2
}

func (b BlockType) String() string {
	return string(b)
}

func (b *BlockData) Init(bt BlockType) {
	b.Type = bt
	switch bt {
	case BlockTypeNormal, BlockTypeNormal2:
		b.MaxDurability = 20
	case BlockTypeObstacle1, BlockTypeObstackle2, BlockTypeNeedle:
		b.MaxDurability = -1
	default:
		b.MaxDurability = 10
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

func (b *BlockData) ForceBreak() {
	b.Durability = 0
	b.MaxDurability = 0
	b.Invalidated = true
}

func (b *BlockData) IsBroken() bool {
	return b.IsBreakable() && b.Durability == 0
}

func (b *BlockData) IsBreakable() bool {
	return b.MaxDurability >= 0
}

func (b *BlockData) IsReversible() bool {
	return b.Type == BlockTypeNormal || b.Type == BlockTypeNormal2 ||
		b.Type == BlockTypeObstacle1 || b.Type == BlockTypeObstackle2
}

func (b *BlockData) IsItem() bool {
	return b.Type == BlockTypeReverse || b.Type == BlockTypeBomb
}

func (b *BlockData) Width() float64 {
	switch b.Type {
	case BlockTypeNormal2, BlockTypeObstackle2:
		return consts.BlockWidth * 2
	}
	return consts.BlockWidth
}

func (b *BlockData) Height() float64 {
	return consts.BlockHeight
}

func (b *BlockData) Hitboxes() []collision.Hitbox {
	return assets.GetHitboxes(b.Type.String())
}

func (b *BlockData) Animation() *ganim8.Animation {
	r := float64(b.Durability) / float64(b.MaxDurability)

	switch b.Type {
	case BlockTypeReverse, BlockTypeBomb, BlockTypeNeedle:
		return assets.GetAnimation(b.Type.String())
	}

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
	case r < 0.05:
		return BlockTypeReverse
	case r < 0.2:
		return BlockTypeBomb
	case r < 0.3:
		return BlockType("needle")
	case r < 0.7:
		return BlockType(fmt.Sprintf("obstacle%d", rand.Intn(2)+1))
	}
	return BlockType(fmt.Sprintf("block%d", rand.Intn(2)+1))
}
