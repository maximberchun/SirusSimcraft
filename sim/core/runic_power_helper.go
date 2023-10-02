package core

import (
	"fmt"
	"time"
)

// RuneCost's bit layout is: <rrrr_rrrr_dduu_ffbb>. Each part is just a count now (0..3 for runes).
type RuneCost uint16

func NewRuneCost(rp, blood, frost, unholy, death int8) RuneCost {
	return RuneCost(rp)<<8 | RuneCost((death&0b11)<<6|(unholy&0b11)<<4|(frost&0b11)<<2|blood&0b11)
}

func (rc RuneCost) String() string {
	return fmt.Sprintf("RP: %d, Blood: %d, Frost: %d, Unholy: %d, Death: %d", rc.RunicPower(), rc.Blood(), rc.Frost(), rc.Unholy(), rc.Death())
}

// HasRune returns if this cost includes a rune portion.
func (rc RuneCost) HasRune() bool {
	return rc&0b1111_1111 > 0
}

func (rc RuneCost) RunicPower() int8 {
	return int8(rc >> 8)
}

func (rc RuneCost) Blood() int8 {
	return int8(rc & 0b11)
}

func (rc RuneCost) Frost() int8 {
	return int8((rc >> 2) & 0b11)
}

func (rc RuneCost) Unholy() int8 {
	return int8((rc >> 4) & 0b11)
}

func (rc RuneCost) Death() int8 {
	return int8((rc >> 6) & 0b11)
}

type Predictor struct {
	rp         *runicPowerBar
	runeStates int16
	runeMeta   [6]RuneMeta
}

func (p *Predictor) SpendRuneCost(sim *Simulation, cost RuneCost) {
	if !cost.HasRune() {
		return
	}

	for i := int8(0); i < cost.Blood(); i++ {
		p.spendRune(sim, 0)
	}
	for i := int8(0); i < cost.Frost(); i++ {
		p.spendRune(sim, 2)
	}
	for i := int8(0); i < cost.Unholy(); i++ {
		p.spendRune(sim, 4)
	}
	for i := int8(0); i < cost.Death(); i++ {
		p.spendDeathRune(sim)
	}
}

func (p *Predictor) spendRune(sim *Simulation, firstSlot int8) {
	slot := p.findReadyRune(firstSlot)
	p.runeStates |= isSpents[slot]
	p.launchRuneRegen(sim, slot)
}

func (p *Predictor) findReadyRune(slot int8) int8 {
	if p.runeStates&isSpentDeath[slot] == 0 {
		return slot
	}
	if p.runeStates&isSpentDeath[slot+1] == 0 {
		return slot + 1
	}
	panic(fmt.Sprintf("findReadyRune(%d) - no slot found (runeStates = %12b)", slot, p.runeStates))
}

func (p *Predictor) spendDeathRune(sim *Simulation) {
	slot := p.findReadyDeathRune()
	if p.rp.btSlot != slot {
		p.runeStates ^= isDeaths[slot] // clear death bit to revert.
	}

	// mark spent bit to spend
	p.runeStates |= isSpents[slot]
	p.launchRuneRegen(sim, slot)
}

func (p *Predictor) findReadyDeathRune() int8 {
	for _, slot := range []int8{4, 5, 2, 3, 0, 1} { // Death runes are spent in the order Unholy -> Frost -> Blood in-game...
		if p.runeStates&isSpentDeath[slot] == isDeaths[slot] {
			return slot
		}
	}
	panic(fmt.Sprintf("findReadyDeathRune() - no slot found (runeStates = %12b)", p.runeStates))
}
func (p *Predictor) launchRuneRegen(sim *Simulation, slot int8) {
	runeGracePeriod := p.runeGraceAt(slot, sim.CurrentTime)
	p.runeMeta[slot].regenAt = sim.CurrentTime + (p.rp.runeCD - runeGracePeriod)
}

func (p *Predictor) runeGraceAt(slot int8, at time.Duration) time.Duration {
	lastRegenTime := p.runeMeta[slot].lastRegenTime
	// pre-pull casts should not get rune-grace
	if at <= 0 || lastRegenTime <= 0 {
		return 0
	}
	return MinDuration(time.Millisecond*2500, at-lastRegenTime)
}

func (p *Predictor) CurrentBloodRunes() int8 {
	return rs2c[(p.runeStates>>0)&0b1111]
}

func (p *Predictor) CurrentFrostRunes() int8 {
	return rs2c[(p.runeStates>>4)&0b1111]
}

func (p *Predictor) CurrentUnholyRunes() int8 {
	return rs2c[(p.runeStates>>8)&0b1111]
}

func (p *Predictor) BloodRuneReadyAt(sim *Simulation) time.Duration {
	if p.runeStates&anyBloodSpent != anyBloodSpent { // if any are not spent
		return sim.CurrentTime
	}
	return MinDuration(p.runeMeta[0].regenAt, p.runeMeta[1].regenAt)
}

func (p *Predictor) FrostRuneReadyAt(sim *Simulation) time.Duration {
	if p.runeStates&anyFrostSpent != anyFrostSpent { // if any are not spent
		return sim.CurrentTime
	}
	return MinDuration(p.runeMeta[2].regenAt, p.runeMeta[3].regenAt)
}

func (p *Predictor) UnholyRuneReadyAt(sim *Simulation) time.Duration {
	if p.runeStates&anyUnholySpent != anyUnholySpent { // if any are not spent
		return sim.CurrentTime
	}
	return MinDuration(p.runeMeta[4].regenAt, p.runeMeta[5].regenAt)
}
