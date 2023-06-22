package components

import (
	"revdriller/assets"

	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type BlockData struct {
	Durability    int
	MaxDurability int
}

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

func (b *BlockData) Animation() *ganim8.Animation {
	r := float64(b.Durability) / float64(b.MaxDurability)
	switch {
	case r > 0.75:
		return assets.GetAnimation("block1_1")
	case r > 0.5:
		return assets.GetAnimation("block1_2")
	case r > 0.25:
		return assets.GetAnimation("block1_3")
	default:
		return assets.GetAnimation("block1_4")
	}
}
