import { Consumes } from '../core/proto/common.js';
import { EquipmentSpec } from '../core/proto/common.js';
import { Flask } from '../core/proto/common.js';
import { Food } from '../core/proto/common.js';
import { Glyphs } from '../core/proto/common.js';
import { Potions } from '../core/proto/common.js';
import { RaidBuffs } from '../core/proto/common.js';
import { IndividualBuffs } from '../core/proto/common.js';
import { Debuffs } from '../core/proto/common.js';
import { TristateEffect } from '../core/proto/common.js';
import { SavedRotation, SavedTalents } from '../core/proto/ui.js';
import { Player } from '../core/player.js';
import { APLRotation } from '../core/proto/apl.js';

import {
	ShadowPriest_Rotation as Rotation,
	ShadowPriest_Options as Options,
	ShadowPriest_Rotation_RotationType,
	PriestMajorGlyph as MajorGlyph,
	PriestMinorGlyph as MinorGlyph,
} from '../core/proto/priest.js';


import * as Tooltips from '../core/constants/tooltips.js';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Standard',
	data: SavedTalents.create({
		talentsString: '05032031--325023051223010323151301351',
		glyphs: Glyphs.create({
			major1: MajorGlyph.GlyphOfShadow,
			major2: MajorGlyph.GlyphOfMindFlay,
			major3: MajorGlyph.GlyphOfDispersion,
			minor1: MinorGlyph.GlyphOfFortitude,
			minor2: MinorGlyph.GlyphOfShadowProtection,
			minor3: MinorGlyph.GlyphOfShadowfiend,
		}),
	}),
};

export const DefaultRotation = Rotation.create({
	rotationType: ShadowPriest_Rotation_RotationType.Ideal,
});

export const DefaultOptions = Options.create({
	useShadowfiend: true,
	useMindBlast: true,
	useShadowWordDeath: true,
	latency: 100,
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskOfTheFrostWyrm,
	food: Food.FoodFishFeast,
	defaultPotion: Potions.PotionOfSpeed,
	prepopPotion: Potions.PotionOfWildMagic,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	arcaneBrilliance: true,
	divineSpirit: true,
	trueshotAura: true,
	leaderOfThePack: TristateEffect.TristateEffectImproved,
	icyTalons: true,
	totemOfWrath: true,
	moonkinAura: TristateEffect.TristateEffectImproved,
	wrathOfAirTotem: true,
	sanctifiedRetribution: true,
	bloodlust: true,
	demonicPact: 500,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	vampiricTouch: true,
});

export const DefaultDebuffs = Debuffs.create({
	sunderArmor: true,
	faerieFire: TristateEffect.TristateEffectImproved,
	bloodFrenzy: true,
	ebonPlaguebringer: true,
	heartOfTheCrusader: true,
	judgementOfWisdom: true,
});

export const PreBis_PRESET = {
	name: 'PreBis Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{"items": [
		{"id":42553,"enchant":3820,"gems":[41285,40049]},
		{"id":40680},
		{"id":34210,"enchant":3810,"gems":[39998,40026]},
		{"id":41610,"enchant":3722},
		{"id":43792,"enchant":1144,"gems":[39998,40051]},
		{"id":37361,"enchant":2332,"gems":[0]},
		{"id":39530,"enchant":3604,"gems":[40049,0]},
		{"id":40696,"gems":[40049,39998]},
		{"id":37854,"enchant":3719},
		{"id":44202,"enchant":3826,"gems":[40026]},
		{"id":40585},
		{"id":37694},
		{"id":37835},
		{"id":37873},
		{"id":41384,"enchant":3834},
		{"id":40698},
		{"id":37177}
  ]}`),
};
export const P1_PRESET = {
	name: 'P1 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(` {"items": [
		{"id":40562,"enchant":3820,"gems":[41285,39998]},
		{"id":44661,"gems":[40026]},
		{"id":40459,"enchant":3810,"gems":[39998]},
		{"id":44005,"enchant":3722,"gems":[40026]},
		{"id":44002,"enchant":1144,"gems":[39998,39998]},
		{"id":44008,"enchant":2332,"gems":[39998,0]},
		{"id":40454,"enchant":3604,"gems":[40049,0]},
		{"id":40561,"gems":[39998]},
		{"id":40560,"enchant":3719},
		{"id":40558,"enchant":3606},
		{"id":40719},
		{"id":40399},
		{"id":40255},
		{"id":40432},
		{"id":40395,"enchant":3834},
		{"id":40273},
		{"id":39712}
  ]}`),
};
export const P2_PRESET = {
	name: 'P2 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{ "items": [
        {"id":46172,"enchant":3820,"gems":[41285,45883]},
        {"id":45243,"gems":[39998]},
        {"id":46165,"enchant":3810,"gems":[39998]},
        {"id":45242,"enchant":3722,"gems":[40049]},
        {"id":46168,"enchant":1144,"gems":[39998,39998]},
        {"id":45446,"enchant":2332,"gems":[39998,0]},
        {"id":45665,"enchant":3604,"gems":[39998,39998,0]},
        {"id":45619,"enchant":3601,"gems":[39998,39998,39998]},
        {"id":46170,"enchant":3719,"gems":[39998,40049]},
        {"id":45135,"enchant":3606,"gems":[39998,40049]},
        {"id":45495,"gems":[40026]},
        {"id":46046,"gems":[39998]},
        {"id":45518},
        {"id":45466},
        {"id":45620,"enchant":3834,"gems":[40026]},
        {"id":45617},
        {"id":45294,"gems":[39998]}
      ]
    }`),
};
export const P3_PRESET = {
	name: 'P3 Preset',
	tooltip: Tooltips.BASIC_BIS_DISCLAIMER,
	gear: EquipmentSpec.fromJsonString(`{"items": [
        {"id":48088,"enchant":3820,"gems":[41285,40133]},
        {"id":47468,"gems":[40155]},
        {"id":48091,"enchant":3810,"gems":[40155]},
        {"id":47551,"enchant":3722,"gems":[40113]},
        {"id":48090,"enchant":1144,"gems":[40113,40133]},
        {"id":47467,"enchant":2332,"gems":[40155,0]},
        {"id":45665,"enchant":3604,"gems":[40113,40113,0]},
        {"id":47419,"enchant":3601,"gems":[40133,40113,40113]},
        {"id":48089,"enchant":3719,"gems":[40113,40133]},
        {"id":47454,"enchant":3606,"gems":[40133,40113]},
        {"id":47489,"gems":[40155]},
        {"id":45495,"gems":[40113]},
        {"id":45518},
        {"id":47477},
        {"id":47483,"enchant":3834},
        {"id":47437},
        {"id":47995}
      ]
    }`),
};
export const ROTATION_PRESET_DEFAULT = {
	name: 'Default',
	rotation: SavedRotation.create({
		specRotationOptionsJson: Rotation.toJsonString(Rotation.create()),
		rotation: APLRotation.fromJsonString(`{
			"type": "TypeAPL",
			"prepullActions": [
				{"action":{"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}},"doAtValue":{"const":{"val":"-1s"}}},
				{"action":{"castSpell":{"spellId":{"spellId":48160}}},"doAtValue":{"const":{"val":"-0.97s"}}}
			],
			"priorityList": [
				{"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"0s"}}}},"castSpell":{"spellId":{"spellId":34433}}}},
				{"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"currentTime":{}},"rhs":{"const":{"val":"1s"}}}},"autocastOtherCooldowns":{}}},
				{"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"currentTime":{}},"rhs":{"const":{"val":"61s"}}}},"castSpell":{"spellId":{"otherId":"OtherActionPotion"}}}},
				{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"1.75s"}}}},"castSpell":{"spellId":{"spellId":48300}}}},
				{"action":{"condition":{"and":{"vals":[{"not":{"val":{"dotIsActive":{"spellId":{"spellId":48125}}}}},{"or":{"vals":[{"and":{"vals":[{"cmp":{"op":"OpEq","lhs":{"const":{"val":"5"}},"rhs":{"auraNumStacks":{"auraId":{"spellId":15258}}}}},{"cmp":{"op":"OpGe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"75s"}}}}]}},{"and":{"vals":[{"cmp":{"op":"OpLe","lhs":{"const":{"val":"3"}},"rhs":{"auraNumStacks":{"auraId":{"spellId":15258}}}}},{"cmp":{"op":"OpLt","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"75s"}}}}]}}]}}]}},"castSpell":{"spellId":{"spellId":48125}}}},
				{"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpLe","lhs":{"dotRemainingTime":{"spellId":{"spellId":48160}}},"rhs":{"spellCastTime":{"spellId":{"spellId":48160}}}}},{"cmp":{"op":"OpGe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"3s"}}}}]}},"castSpell":{"spellId":{"spellId":48160}}}},
				{"action":{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":48300}}}}},"castSpell":{"spellId":{"spellId":48300}}}},
				{"hide":true,"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpGe","lhs":{"spellCastTime":{"spellId":{"spellId":48127}}},"rhs":{"const":{"val":"750ms"}}}},{"cmp":{"op":"OpLe","lhs":{"auraRemainingTime":{"auraId":{"spellId":57669}}},"rhs":{"const":{"val":"5s"}}}}]}},"castSpell":{"spellId":{"spellId":48127}}}},
				{"hide":true,"action":{"condition":{"cmp":{"op":"OpGe","lhs":{"spellCastTime":{"spellId":{"spellId":48127}}},"rhs":{"const":{"val":"750ms"}}}},"castSpell":{"spellId":{"spellId":48127}}}},
				{"action":{"condition":{"cmp":{"op":"OpEq","lhs":{"auraNumStacks":{"auraId":{"spellId":15258}}},"rhs":{"const":{"val":"5"}}}},"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":48156}}}]}}},
				{"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpLe","lhs":{"dotRemainingTime":{"spellId":{"spellId":48300}}},"rhs":{"const":{"val":"200ms"}}}}]}},"wait":{"duration":{"dotRemainingTime":{"spellId":{"spellId":48300}}}}}},
				{"action":{"condition":{"and":{"vals":[{"cmp":{"op":"OpLe","lhs":{"dotRemainingTime":{"spellId":{"spellId":48160}}},"rhs":{"math":{"op":"OpAdd","lhs":{"spellCastTime":{"spellId":{"spellId":48160}}},"rhs":{"channelClipDelay":{}}}}}},{"dotIsActive":{"spellId":{"spellId":48160}}},{"spellIsChanneling":{"spellId":{"spellId":48156}}}]}},"wait":{"duration":{"math":{"op":"OpSub","lhs":{"dotRemainingTime":{"spellId":{"spellId":48160}}},"rhs":{"spellCastTime":{"spellId":{"spellId":48160}}}}}}}},
				{"action":{"channelSpell":{"spellId":{"spellId":48156},"interruptIf":{"cmp":{"op":"OpLe","lhs":{"gcdTimeToReady":{}},"rhs":{"channelClipDelay":{}}}}}}},
				{"action":{"castSpell":{"spellId":{"spellId":47585}}}}
			]
		}`),
	}),
};