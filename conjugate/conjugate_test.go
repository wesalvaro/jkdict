package conjugate

import (
	"reflect"
	"testing"
)

var conjugationTests = []struct {
	desc     string
	word     string
	pos      string
	expected Conjugations
}{
	{
		"I-Adjective",
		"美味しい",
		"adj-i",
		Conjugations{
			NonPast:          []conjVariant{{Plain: "美味しい", Formal: "美味しいです", PlainNegative: "美味しくない", FormalNegative: "美味しくないです"}},
			Past:             []conjVariant{{Plain: "美味しかった", Formal: "美味しかったです", PlainNegative: "美味しくなかった", FormalNegative: "美味しくなかったです"}},
			Conjunctive:      []conjVariant{{Plain: "美味しくて", Formal: "", PlainNegative: "美味しくなくて", FormalNegative: ""}},
			Provisional:      []conjVariant{{Plain: "美味しければ", Formal: "", PlainNegative: "美味しくなければ", FormalNegative: ""}},
			Potential:        []conjVariant{},
			Passive:          []conjVariant{},
			Causative:        []conjVariant{{Plain: "美味しくさせる", Formal: "", PlainNegative: "", FormalNegative: ""}},
			CausativePassive: []conjVariant{},
			Volitional:       []conjVariant{{Plain: "美味しかろう", Formal: "美味しいでしょう", PlainNegative: "", FormalNegative: ""}},
			Imperative:       []conjVariant{},
			Conditional:      []conjVariant{{Plain: "美味しかったら", Formal: "", PlainNegative: "美味しくなかったら", FormalNegative: ""}},
			Alternative:      []conjVariant{{Plain: "美味しかったり", Formal: "", PlainNegative: "", FormalNegative: ""}},
			Continuative:     []conjVariant{},
		},
	},
	{
		"Ichidan Verb",
		"食べる",
		"v1",
		Conjugations{},
	},
	{
		"Godan verb -bu",
		"遊ぶ",
		"v5b",
		Conjugations{},
	},
	{
		"Godan verb -gu",
		"游ぐ",
		"v5g",
		Conjugations{},
	},
}

func TestConjugations(t *testing.T) {
	for _, tt := range conjugationTests {
		if !reflect.DeepEqual(Conjugate(tt.word, tt.pos), &tt.expected) {
			t.Errorf("%s --\n%+v", tt.desc, Conjugate(tt.word, tt.pos))
		}
	}
}
