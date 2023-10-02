package protection

import (
	"testing"

	_ "github.com/wowsims/wotlk/sim/common" // imported to get item effects included.
	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/core/proto"
)

func init() {
	RegisterProtectionWarrior()
}

func TestProtectionWarrior(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator(core.CharacterSuiteConfig{
		Class:      proto.Class_ClassWarrior,
		Race:       proto.Race_RaceOrc,
		OtherRaces: []proto.Race{proto.Race_RaceHuman},

		GearSet:     core.GearSetCombo{Label: "P1", GearSet: P1Gear},
		Talents:     DefaultTalents,
		Glyphs:      DefaultGlyphs,
		Consumes:    FullConsumes,
		SpecOptions: core.SpecOptionsCombo{Label: "Basic", SpecOptions: PlayerOptionsBasic},
		Rotation:    core.RotationCombo{Label: "Default", Rotation: DefaultRotation},

		IsTank:          true,
		InFrontOfTarget: true,

		ItemFilter: core.ItemFilter{
			ArmorType: proto.ArmorType_ArmorTypePlate,

			WeaponTypes: []proto.WeaponType{
				proto.WeaponType_WeaponTypeAxe,
				proto.WeaponType_WeaponTypeSword,
				proto.WeaponType_WeaponTypeMace,
				proto.WeaponType_WeaponTypeDagger,
				proto.WeaponType_WeaponTypeFist,
				proto.WeaponType_WeaponTypeShield,
			},
		},

		EPReferenceStat: proto.Stat_StatAttackPower,
		StatsToWeigh: []proto.Stat{
			proto.Stat_StatStrength,
			proto.Stat_StatAttackPower,
			proto.Stat_StatArmor,
			proto.Stat_StatDodge,
			proto.Stat_StatBlockValue,
		},
	}))
}

func BenchmarkSimulate(b *testing.B) {
	rsr := &proto.RaidSimRequest{
		Raid: core.SinglePlayerRaidProto(
			&proto.Player{
				Race:          proto.Race_RaceOrc,
				Class:         proto.Class_ClassWarrior,
				Equipment:     P1Gear,
				Consumes:      FullConsumes,
				Spec:          PlayerOptionsBasic,
				Buffs:         core.FullIndividualBuffs,
				TalentsString: DefaultTalents,
				Glyphs:        DefaultGlyphs,

				InFrontOfTarget: true,
			},
			core.FullPartyBuffs,
			core.FullRaidBuffs,
			core.FullDebuffs),
		Encounter: &proto.Encounter{
			Duration: 300,
			Targets: []*proto.Target{
				core.NewDefaultTarget(),
			},
		},
		SimOptions: core.AverageDefaultSimTestOptions,
	}

	core.RaidBenchmark(b, rsr)
}

var DefaultTalents = "2500030023-302-053351225000012521030113321"
var DefaultGlyphs = &proto.Glyphs{
	Major1: int32(proto.WarriorMajorGlyph_GlyphOfBlocking),
	Major2: int32(proto.WarriorMajorGlyph_GlyphOfDevastate),
	Major3: int32(proto.WarriorMajorGlyph_GlyphOfVigilance),
}

var PlayerOptionsBasic = &proto.Player_ProtectionWarrior{
	ProtectionWarrior: &proto.ProtectionWarrior{
		Options:  warriorOptions,
		Rotation: &proto.ProtectionWarrior_Rotation{},
	},
}

var warriorOptions = &proto.ProtectionWarrior_Options{
	Shout:        proto.WarriorShout_WarriorShoutCommanding,
	StartingRage: 0,
}

var FullConsumes = &proto.Consumes{
	BattleElixir:   proto.BattleElixir_ElixirOfMastery,
	GuardianElixir: proto.GuardianElixir_GiftOfArthas,
}

var DefaultRotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [
		{"action":{"castSpell":{"spellId":{"spellId":47440}}},"doAtValue":{"const":{"val":"-10s"}}},
		{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}}
	],
	"priorityList": [
		{"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"currentRage":{}},"rhs":{"const":{"val":"30"}}}},"castSpell":{"spellId":{"tag":1,"spellId":47450}}}},
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"castSpell":{"spellId":{"spellId":47488}}}},
		{"action":{"castSpell":{"spellId":{"spellId":57823}}}},
		{"action":{"condition":{"auraShouldRefresh":{"sourceUnit":{"type":"Self"},"auraId":{"spellId":47440},"maxOverlap":{"const":{"val":"3s"}}}},"castSpell":{"spellId":{"spellId":47440}}}},
		{"action":{"condition":{"auraShouldRefresh":{"auraId":{"spellId":47502},"maxOverlap":{"const":{"val":"2s"}}}},"castSpell":{"spellId":{"spellId":47502}}}},
		{"action":{"condition":{"auraShouldRefresh":{"auraId":{"spellId":47437},"maxOverlap":{"const":{"val":"2s"}}}},"castSpell":{"spellId":{"spellId":25203}}}},
		{"action":{"castSpell":{"spellId":{"spellId":47498}}}}
	]
}`)

var P1Gear = core.EquipmentSpecFromJsonString(`{"items": [
	{"id":40546,"enchant":3818,"gems":[41380,40034]},
	{"id":40387},
	{"id":39704,"enchant":3852,"gems":[40034]},
	{"id":40722,"enchant":3605},
	{"id":44000,"enchant":3832,"gems":[40034,40015]},
	{"id":39764,"enchant":3850,"gems":[0]},
	{"id":40545,"enchant":3860,"gems":[40034,0]},
	{"id":39759,"enchant":3601,"gems":[40008,36767]},
	{"id":40589,"enchant":3822},
	{"id":39717,"enchant":3232,"gems":[40089]},
	{"id":40370},
	{"id":40718},
	{"id":40257},
	{"id":44063,"gems":[36767,40089]},
	{"id":40402,"enchant":3788},
	{"id":40400,"enchant":3849},
	{"id":41168,"gems":[36767]}
]}`)
