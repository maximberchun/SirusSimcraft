package mage

import (
	"testing"

	_ "github.com/wowsims/wotlk/sim/common"
	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/core/proto"
)

func init() {
	RegisterMage()
}

func TestArcane(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator(core.CharacterSuiteConfig{
		Class: proto.Class_ClassMage,
		Race:  proto.Race_RaceTroll,

		GearSet:     core.GearSetCombo{Label: "P3Arcane", GearSet: P3ArcaneGear},
		Talents:     ArcaneTalents,
		Glyphs:      ArcaneGlyphs,
		Consumes:    FullArcaneConsumes,
		SpecOptions: core.SpecOptionsCombo{Label: "Arcane", SpecOptions: PlayerOptionsArcane},
		Rotation:    core.RotationCombo{Label: "Arcane", Rotation: ArcaneRotation},
		OtherRotations: []core.RotationCombo{
			{Label: "AOE", Rotation: ArcaneAOERotation},
		},

		ItemFilter: core.ItemFilter{
			WeaponTypes: []proto.WeaponType{
				proto.WeaponType_WeaponTypeDagger,
				proto.WeaponType_WeaponTypeSword,
				proto.WeaponType_WeaponTypeOffHand,
				proto.WeaponType_WeaponTypeStaff,
			},
			ArmorType: proto.ArmorType_ArmorTypeCloth,
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeWand,
			},
		},
	}))
}

func TestFire(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator(core.CharacterSuiteConfig{
		Class: proto.Class_ClassMage,
		Race:  proto.Race_RaceTroll,

		GearSet:     core.GearSetCombo{Label: "P3Fire", GearSet: P3FireGear},
		Talents:     FireTalents,
		Glyphs:      FireGlyphs,
		Consumes:    FullFireConsumes,
		SpecOptions: core.SpecOptionsCombo{Label: "Fire", SpecOptions: PlayerOptionsFire},
		Rotation:    core.RotationCombo{Label: "Fire", Rotation: FireRotation},
		OtherRotations: []core.RotationCombo{
			{Label: "AOE", Rotation: FireAOERotation},
		},

		ItemFilter: core.ItemFilter{
			WeaponTypes: []proto.WeaponType{
				proto.WeaponType_WeaponTypeDagger,
				proto.WeaponType_WeaponTypeSword,
				proto.WeaponType_WeaponTypeOffHand,
				proto.WeaponType_WeaponTypeStaff,
			},
			ArmorType: proto.ArmorType_ArmorTypeCloth,
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeWand,
			},
		},
	}))
}

func TestFrostFire(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator(core.CharacterSuiteConfig{
		Class: proto.Class_ClassMage,
		Race:  proto.Race_RaceTroll,

		GearSet:     core.GearSetCombo{Label: "P3FrostFire", GearSet: P3FireGear},
		Talents:     FrostFireTalents,
		Glyphs:      FrostFireGlyphs,
		Consumes:    FullFireConsumes,
		SpecOptions: core.SpecOptionsCombo{Label: "Fire", SpecOptions: PlayerOptionsFire},
		Rotation:    core.RotationCombo{Label: "Frostfire", Rotation: FrostfireRotation},

		ItemFilter: core.ItemFilter{
			WeaponTypes: []proto.WeaponType{
				proto.WeaponType_WeaponTypeDagger,
				proto.WeaponType_WeaponTypeSword,
				proto.WeaponType_WeaponTypeOffHand,
				proto.WeaponType_WeaponTypeStaff,
			},
			ArmorType: proto.ArmorType_ArmorTypeCloth,
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeWand,
			},
		},
	}))
}

func TestFrost(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator(core.CharacterSuiteConfig{
		Class: proto.Class_ClassMage,
		Race:  proto.Race_RaceTroll,

		GearSet:     core.GearSetCombo{Label: "P3Frost", GearSet: P3FrostGear},
		Talents:     FrostTalents,
		Glyphs:      FrostGlyphs,
		Consumes:    FullFrostConsumes,
		SpecOptions: core.SpecOptionsCombo{Label: "Frost", SpecOptions: PlayerOptionsFrost},
		Rotation:    core.RotationCombo{Label: "Frost", Rotation: FrostRotation},
		OtherRotations: []core.RotationCombo{
			{Label: "AOE", Rotation: FrostAOERotation},
		},

		ItemFilter: core.ItemFilter{
			WeaponTypes: []proto.WeaponType{
				proto.WeaponType_WeaponTypeDagger,
				proto.WeaponType_WeaponTypeSword,
				proto.WeaponType_WeaponTypeOffHand,
				proto.WeaponType_WeaponTypeStaff,
			},
			ArmorType: proto.ArmorType_ArmorTypeCloth,
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeWand,
			},
		},
	}))
}

var ArcaneTalents = "23000513310033015032310250532-03-023303001"
var FireTalents = "23000503110003-0055030012303331053120301351"
var FrostFireTalents = "23000503110003-0055030012303331053120301351"
var FrostTalents = "23000503110003--0533030310233100030152231351"
var ArcaneGlyphs = &proto.Glyphs{
	Major1: int32(proto.MageMajorGlyph_GlyphOfArcaneBlast),
	Major2: int32(proto.MageMajorGlyph_GlyphOfArcaneMissiles),
	Major3: int32(proto.MageMajorGlyph_GlyphOfMoltenArmor),
}
var FireGlyphs = &proto.Glyphs{
	Major1: int32(proto.MageMajorGlyph_GlyphOfFireball),
	Major2: int32(proto.MageMajorGlyph_GlyphOfMoltenArmor),
	Major3: int32(proto.MageMajorGlyph_GlyphOfLivingBomb),
}
var FrostFireGlyphs = &proto.Glyphs{
	Major1: int32(proto.MageMajorGlyph_GlyphOfFrostfire),
	Major2: int32(proto.MageMajorGlyph_GlyphOfMoltenArmor),
	Major3: int32(proto.MageMajorGlyph_GlyphOfLivingBomb),
}
var FrostGlyphs = &proto.Glyphs{
	Major1: int32(proto.MageMajorGlyph_GlyphOfFrostbolt),
	Major3: int32(proto.MageMajorGlyph_GlyphOfEternalWater),
	Major2: int32(proto.MageMajorGlyph_GlyphOfMoltenArmor),
}

var PlayerOptionsFire = &proto.Player_Mage{
	Mage: &proto.Mage{
		Options: &proto.Mage_Options{
			Armor:          proto.Mage_Options_MoltenArmor,
			IgniteMunching: true,
		},
		Rotation: &proto.Mage_Rotation{},
	},
}

var PlayerOptionsFrost = &proto.Player_Mage{
	Mage: &proto.Mage{
		Options: &proto.Mage_Options{
			Armor: proto.Mage_Options_MageArmor,
		},
		Rotation: &proto.Mage_Rotation{},
	},
}

var PlayerOptionsArcane = &proto.Player_Mage{
	Mage: &proto.Mage{
		Options: &proto.Mage_Options{
			Armor: proto.Mage_Options_MoltenArmor,
		},
		Rotation: &proto.Mage_Rotation{},
	},
}

var FullFireConsumes = &proto.Consumes{
	Flask:         proto.Flask_FlaskOfTheFrostWyrm,
	Food:          proto.Food_FoodFirecrackerSalmon,
	DefaultPotion: proto.Potions_PotionOfSpeed,
}
var FullFrostConsumes = FullFireConsumes

var FullArcaneConsumes = FullFireConsumes

var ArcaneRotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [
		{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}}
	],
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"condition":{"not":{"val":{"auraIsActive":{"auraId":{"spellId":12472}}}}},"castSpell":{"spellId":{"spellId":26297}}}},
		{"action":{"condition":{"not":{"val":{"auraIsActive":{"auraId":{"spellId":12472}}}}},"castSpell":{"spellId":{"spellId":54758}}}},
		{"action":{"condition":{"not":{"val":{"auraIsActive":{"auraId":{"spellId":12472}}}}},"castSpell":{"spellId":{"itemId":40211}}}},
		{"action":{"condition":{"cmp":{"op":"OpLt","lhs":{"auraNumStacks":{"auraId":{"spellId":36032}}},"rhs":{"const":{"val":"4"}}}},"castSpell":{"spellId":{"spellId":42897}}}},
		{"action":{"condition":{"auraIsActiveWithReactionTime":{"auraId":{"spellId":44401}}},"castSpell":{"spellId":{"spellId":42846}}}},
		{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"currentManaPercent":{}},"rhs":{"const":{"val":"25%"}}}},"castSpell":{"spellId":{"spellId":12051}}}},
		{"action":{"condition":{"cmp":{"op":"OpGt","lhs":{"currentManaPercent":{}},"rhs":{"const":{"val":"25%"}}}},"castSpell":{"spellId":{"spellId":42897}}}},
		{"action":{"castSpell":{"spellId":{"spellId":42846}}}}
	]
}`)

var ArcaneAOERotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [
		{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}}
	],
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"castSpell":{"spellId":{"spellId":42921}}}}
	]
}`)

var FireRotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [
		{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}}
	],
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"condition":{"auraShouldRefresh":{"auraId":{"spellId":12873},"maxOverlap":{"const":{"val":"4s"}}}},"castSpell":{"spellId":{"spellId":42859}}}},
		{"action":{"condition":{"auraIsActiveWithReactionTime":{"auraId":{"spellId":44448}}},"castSpell":{"spellId":{"spellId":42891}}}},
		{"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpGt","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"12s"}}}}]}},"multidot":{"spellId":{"spellId":55360},"maxDots":10,"maxOverlap":{"const":{"val":"0ms"}}}}},
		{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"remainingTime":{}},"rhs":{"spellCastTime":{"spellId":{"spellId":42859}}}}},"castSpell":{"spellId":{"spellId":42873}}}},
		{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"4s"}}}},"castSpell":{"spellId":{"spellId":42859}}}},
		{"action":{"castSpell":{"spellId":{"spellId":42833}}}}
	]
}`)

var FireAOERotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [],
	"priorityList": [
	   {"action":{"autocastOtherCooldowns":{}}},
	   {"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"12s"}}}},"multidot":{"spellId":{"spellId":55360},"maxDots":10,"maxOverlap":{"const":{"val":"0ms"}}}}},
	   {"action":{"condition":{"and":{"vals":[{"auraIsActive":{"auraId":{"spellId":54741}}},{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42926,"tag":9}}}}}]}},"castSpell":{"spellId":{"spellId":42926,"tag":9}}}},
	   {"action":{"condition":{"and":{"vals":[{"auraIsActive":{"auraId":{"spellId":54741}}},{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42925,"tag":8}}}}}]}},"castSpell":{"spellId":{"spellId":42925,"tag":8}}}},
	   {"action":{"condition":{"or":{"vals":[{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42926,"tag":9}}}}},{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42925,"tag":8}}}}}]}},"castSpell":{"spellId":{"spellId":42950}}}},
	   {"action":{"condition":{"or":{"vals":[{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42926,"tag":9}}}}},{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42925,"tag":8}}}}}]}},"castSpell":{"spellId":{"spellId":42945}}}},
	   {"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42926,"tag":9}}}}},"castSpell":{"spellId":{"spellId":42926,"tag":9}}}},
	   {"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":42925,"tag":8}}}}},"castSpell":{"spellId":{"spellId":42925,"tag":8}}}},
	   {"action":{"condition":{"auraIsActiveWithReactionTime":{"auraId":{"spellId":44448}}},"castSpell":{"spellId":{"spellId":42891}}}},
	   {"action":{"castSpell":{"spellId":{"spellId":42921}}}}
	]
}`)

var FrostfireRotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [
		{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}}
	],
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"condition":{"auraShouldRefresh":{"auraId":{"spellId":12873},"maxOverlap":{"const":{"val":"4s"}}}},"castSpell":{"spellId":{"spellId":42859}}}},
		{"action":{"condition":{"auraIsActiveWithReactionTime":{"auraId":{"spellId":44448}}},"castSpell":{"spellId":{"spellId":42891}}}},
		{"action":{"condition":{"and":{"vals":[{"not":{"val":{"dotIsActive":{"spellId":{"spellId":55360}}}}},{"cmp":{"op":"OpGt","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"12s"}}}}]}},"castSpell":{"spellId":{"spellId":55360}}}},
		{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"remainingTime":{}},"rhs":{"spellCastTime":{"spellId":{"spellId":42859}}}}},"castSpell":{"spellId":{"spellId":42873}}}},
		{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"3.5s"}}}},"castSpell":{"spellId":{"spellId":42859}}}},
		{"action":{"castSpell":{"spellId":{"spellId":47610}}}}
	]
}`)

var FrostRotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [
		{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}}
	],
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"condition":{"not":{"val":{"auraIsActive":{"auraId":{"spellId":12472}}}}},"castSpell":{"spellId":{"spellId":26297}}}},
		{"action":{"condition":{"not":{"val":{"auraIsActive":{"auraId":{"spellId":12472}}}}},"castSpell":{"spellId":{"spellId":54758}}}},
		{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"currentManaPercent":{}},"rhs":{"const":{"val":"25%"}}}},"castSpell":{"spellId":{"spellId":12051}}}},
		{"action":{"condition":{"auraIsActive":{"auraId":{"spellId":44545}}},"castSpell":{"spellId":{"spellId":44572}}}},
		{"action":{"condition":{"auraIsActiveWithReactionTime":{"auraId":{"spellId":44549}}},"castSpell":{"spellId":{"spellId":47610}}}},
		{"action":{"castSpell":{"spellId":{"spellId":42842}}}}
	]
}`)

var FrostAOERotation = core.APLRotationFromJsonString(`{
	"type": "TypeAPL",
	"prepullActions": [
		{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}}
	],
	"priorityList": [
		{"action":{"autocastOtherCooldowns":{}}},
		{"action":{"castSpell":{"spellId":{"spellId":42939}}}}
	]
}`)

var P3FireGear = core.EquipmentSpecFromJsonString(`{"items": [
		{"id":51281,"enchant":3820,"gems":[41285,40133]},
        {"id":50724,"gems":[40133]},
        {"id":51284,"enchant":3810,"gems":[40153]},
        {"id":50628,"enchant":3722,"gems":[40153]},
        {"id":51283,"enchant":3832,"gems":[40113,40133]},
        {"id":54582,"enchant":2332,"gems":[40155,0]},
        {"id":50722,"enchant":3604,"gems":[40153,40133,0]},
        {"id":50613,"enchant":3601,"gems":[40133,40113,40113]},
        {"id":51282,"enchant":3872,"gems":[40133,40153]},
        {"id":50699,"enchant":3606,"gems":[40133,40113]},
        {"id":50664,"gems":[40133]},
        {"id":50398,"gems":[40153]},
        {"id":47188},
        {"id":50348},
        {"id":50732,"enchant":3834,"gems":[40113]},
        {"id":50719},
        {"id":50684,"gems":[40153]}
]}`)
var P3FrostGear = P3FireGear
var P3ArcaneGear = P3FireGear
