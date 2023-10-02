package core

import (
	"fmt"
	"time"

	"github.com/wowsims/wotlk/sim/core/proto"
)

type APLValueSpellCanCast struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellCanCast(config *proto.APLValueSpellCanCast) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellCanCast{
		spell: spell,
	}
}
func (value *APLValueSpellCanCast) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueSpellCanCast) GetBool(sim *Simulation) bool {
	return value.spell.CanCast(sim, value.spell.Unit.CurrentTarget)
}
func (value *APLValueSpellCanCast) String() string {
	return fmt.Sprintf("Can Cast(%s)", value.spell.ActionID)
}

type APLValueSpellIsReady struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellIsReady(config *proto.APLValueSpellIsReady) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellIsReady{
		spell: spell,
	}
}
func (value *APLValueSpellIsReady) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueSpellIsReady) GetBool(sim *Simulation) bool {
	return value.spell.IsReady(sim)
}
func (value *APLValueSpellIsReady) String() string {
	return fmt.Sprintf("Is Ready(%s)", value.spell.ActionID)
}

type APLValueSpellTimeToReady struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellTimeToReady(config *proto.APLValueSpellTimeToReady) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellTimeToReady{
		spell: spell,
	}
}
func (value *APLValueSpellTimeToReady) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueSpellTimeToReady) GetDuration(sim *Simulation) time.Duration {
	return value.spell.TimeToReady(sim)
}
func (value *APLValueSpellTimeToReady) String() string {
	return fmt.Sprintf("Time To Ready(%s)", value.spell.ActionID)
}

type APLValueSpellCastTime struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellCastTime(config *proto.APLValueSpellCastTime) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellCastTime{
		spell: spell,
	}
}
func (value *APLValueSpellCastTime) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueSpellCastTime) GetDuration(sim *Simulation) time.Duration {
	return value.spell.Unit.ApplyCastSpeedForSpell(value.spell.DefaultCast.CastTime, value.spell)
}
func (value *APLValueSpellCastTime) String() string {
	return fmt.Sprintf("Cast Time(%s)", value.spell.ActionID)
}

type APLValueSpellChannelTime struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellChannelTime(config *proto.APLValueSpellChannelTime) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellChannelTime{
		spell: spell,
	}
}
func (value *APLValueSpellChannelTime) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueSpellChannelTime) GetDuration(sim *Simulation) time.Duration {
	return value.spell.Unit.ApplyCastSpeedForSpell(value.spell.DefaultCast.ChannelTime, value.spell)
}
func (value *APLValueSpellChannelTime) String() string {
	return fmt.Sprintf("Channel Time(%s)", value.spell.ActionID)
}

type APLValueSpellTravelTime struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellTravelTime(config *proto.APLValueSpellTravelTime) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellTravelTime{
		spell: spell,
	}
}
func (value *APLValueSpellTravelTime) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueSpellTravelTime) GetDuration(sim *Simulation) time.Duration {
	return time.Duration(float64(time.Second) * value.spell.Unit.DistanceFromTarget / value.spell.MissileSpeed)
}
func (value *APLValueSpellTravelTime) String() string {
	return fmt.Sprintf("Travel Time(%s)", value.spell.ActionID)
}

type APLValueSpellCPM struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellCPM(config *proto.APLValueSpellCPM) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellCPM{
		spell: spell,
	}
}
func (value *APLValueSpellCPM) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeFloat
}
func (value *APLValueSpellCPM) GetFloat(sim *Simulation) float64 {
	return value.spell.CurCPM(sim)
}
func (value *APLValueSpellCPM) String() string {
	return fmt.Sprintf("CPM(%s)", value.spell.ActionID)
}

type APLValueSpellIsChanneling struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellIsChanneling(config *proto.APLValueSpellIsChanneling) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellIsChanneling{
		spell: spell,
	}
}
func (value *APLValueSpellIsChanneling) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueSpellIsChanneling) GetBool(sim *Simulation) bool {
	return value.spell.Unit.ChanneledDot != nil && value.spell.Unit.ChanneledDot.Spell == value.spell
}
func (value *APLValueSpellIsChanneling) String() string {
	return fmt.Sprintf("IsChanneling(%s)", value.spell.ActionID)
}

type APLValueSpellChanneledTicks struct {
	DefaultAPLValueImpl
	spell *Spell
}

func (rot *APLRotation) newValueSpellChanneledTicks(config *proto.APLValueSpellChanneledTicks) APLValue {
	spell := rot.GetAPLSpell(config.SpellId)
	if spell == nil {
		return nil
	}
	return &APLValueSpellChanneledTicks{
		spell: spell,
	}
}
func (value *APLValueSpellChanneledTicks) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeInt
}
func (value *APLValueSpellChanneledTicks) GetInt(sim *Simulation) int32 {
	channeledDot := value.spell.Unit.ChanneledDot
	if channeledDot == nil {
		return 0
	} else {
		return channeledDot.TickCount
	}
}
func (value *APLValueSpellChanneledTicks) String() string {
	return fmt.Sprintf("ChanneledTicks(%s)", value.spell.ActionID)
}
