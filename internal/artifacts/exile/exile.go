package exile

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/artifact"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func init() {
	core.RegisterSetFunc(keys.TheExile, NewSet)
}

type Set struct {
	Index int
}

func (s *Set) SetIndex(idx int) { s.Index = idx }
func (s *Set) Init() error      { return nil }

// 2-Piece Bonus: Energy Recharge +20%.
// 4-Piece Bonus: Using an Elemental Burst regenerates 2 Energy for all party members (excluding the wearer) every 2s for 6s. This effect cannot stack.
func NewSet(c *core.Core, char *character.CharWrapper, count int, param map[string]int) (artifact.Set, error) {
	s := Set{}

	if count >= 2 {
		m := make([]float64, attributes.EndStatType)
		m[attributes.ER] = 0.20
		char.AddStatMod(character.StatMod{
			Base:         modifier.NewBase("exile-2pc", -1),
			AffectedStat: attributes.ER,
			Amount: func() ([]float64, bool) {
				return m, true
			},
		})
	}
	if count >= 4 {
		// TODO: does multiple exile holders extend the duration?
		c.Events.Subscribe(event.OnBurst, func(args ...interface{}) bool {
			if c.Player.Active() != char.Index {
				return false
			}
			if c.Status.Duration("exile") > 0 {
				return false
			}
			c.Status.Add("exile", 6*60)

			for _, x := range c.Player.Chars() {
				this := x
				if char.Index == this.Index {
					continue
				}
				// 3 ticks
				for i := 120; i <= 360; i += 120 {
					//TODO: should this delay be affected by wearer's hitlag?
					c.Tasks.Add(func() {
						this.AddEnergy("exile-4pc", 2)
					}, i)
				}
			}

			return false
		}, fmt.Sprintf("exile-4pc-%v", char.Base.Key.String()))
	}

	return &s, nil
}
