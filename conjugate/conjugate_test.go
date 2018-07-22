package conjugate

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
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
			NonPast: []Variant{
				{Plain: "美味しい", Formal: "美味しいです", PlainNegative: "美味しくない", FormalNegative: "美味しくないです"}},
			Past: []Variant{
				{Plain: "美味しかった", Formal: "美味しかったです", PlainNegative: "美味しくなかった", FormalNegative: "美味しくなかったです"}},
			Conjunctive: []Variant{
				{Plain: "美味しくて", PlainNegative: "美味しくなくて"}},
			Provisional: []Variant{
				{Plain: "美味しければ", PlainNegative: "美味しくなければ"}},
			Causative: []Variant{
				{Plain: "美味しくさせる"}},
			CausativePassive: []Variant{},
			Volitional: []Variant{
				{Plain: "美味しかろう", Formal: "美味しいでしょう"}},
			Conditional: []Variant{
				{Plain: "美味しかったら", PlainNegative: "美味しくなかったら"}},
			Alternative: []Variant{
				{Plain: "美味しかったり"}},
		},
	},
	{
		"Ichidan Verb",
		"食べる",
		"v1",
		Conjugations{
			NonPast: []Variant{
				{Plain: "食べる", Formal: "食べます", PlainNegative: "食べない", FormalNegative: "食べました"}},
			Past: []Variant{
				{Plain: "食べた", Formal: "食べました", PlainNegative: "食べなかった", FormalNegative: "食べませんでした"}},
			Conjunctive: []Variant{
				{Plain: "食べて", Formal: "食べまして", PlainNegative: "食べなくて", FormalNegative: "食べませんで"},
				{PlainNegative: "食べないで"}},
			Provisional: []Variant{
				{Plain: "食べれば", Formal: "食べますなら", PlainNegative: "食べなければ", FormalNegative: "食べませんなら"}},
			Potential: []Variant{
				{Plain: "食べられる", Formal: "食べられます", PlainNegative: "食べられない", FormalNegative: "食べられません"},
				{Plain: "食べれる", Formal: "食べれます", PlainNegative: "食べれない", FormalNegative: "食べれません"}},
			Passive: []Variant{
				{Plain: "食べられる", Formal: "食べられます", PlainNegative: "食べられない", FormalNegative: "食べられません"}},
			Causative: []Variant{
				{Plain: "食べさせる", Formal: "食べさせます", PlainNegative: "食べさせない", FormalNegative: "食べさせません"},
				{Plain: "食べさす", Formal: "食べさします", PlainNegative: "食べささない", FormalNegative: "食べさしません"}},
			CausativePassive: []Variant{
				{Plain: "食べさせられる", Formal: "食べさせられます", PlainNegative: "食べさせられない", FormalNegative: "食べさせられません"}},
			Volitional: []Variant{
				{Plain: "食べよう", Formal: "食べましょう", PlainNegative: "食べまい", FormalNegative: "食べますまい"}},
			Imperative: []Variant{
				{Plain: "食べろ", Formal: "食べなさい", PlainNegative: "食べるな", FormalNegative: "食べなさるな"}},
			Conditional: []Variant{
				{Plain: "食べたら", Formal: "食べましたら", PlainNegative: "食べなかったら", FormalNegative: "食べませんでしたら"}},
			Alternative: []Variant{
				{Plain: "食べたり", Formal: "食べましたり", PlainNegative: "食べなかったり", FormalNegative: "食べませんでしたり"}},
			Continuative: []Variant{
				{Plain: "食べ"}}},
	},
	{
		"Godan verb -bu",
		"遊ぶ",
		"v5b",
		Conjugations{
			NonPast: []Variant{
				{Plain: "遊ぶ", Formal: "遊びます", PlainNegative: "遊ばない", FormalNegative: "遊びません"}},
			Past: []Variant{
				{Plain: "遊んだ", Formal: "遊びました", PlainNegative: "遊ばなかった", FormalNegative: "遊びませんでした"}},
			Conjunctive: []Variant{
				{Plain: "遊んで", Formal: "遊びまして", PlainNegative: "遊ばなくて", FormalNegative: "遊びませんで"},
				{PlainNegative: "遊ばないで"}},
			Provisional: []Variant{
				{Plain: "遊べば", Formal: "遊びますなら", PlainNegative: "遊ばなければ", FormalNegative: "遊びませんなら"}},
			Potential: []Variant{
				{Plain: "遊べる", Formal: "遊べます", PlainNegative: "遊べない", FormalNegative: "遊べません"}},
			Passive: []Variant{
				{Plain: "遊ばれる", Formal: "遊ばれます", PlainNegative: "遊ばれない", FormalNegative: "遊ばれません"}},
			Causative: []Variant{
				{Plain: "遊ばせる", Formal: "遊ばせます", PlainNegative: "遊ばせない", FormalNegative: "遊ばせません"},
				{Plain: "遊ばす", Formal: "遊ばします", PlainNegative: "遊ばさない", FormalNegative: "遊ばしません"}},
			CausativePassive: []Variant{
				{Plain: "遊ばせられる", Formal: "遊ばせられます", PlainNegative: "遊ばせられない", FormalNegative: "遊ばせられません"},
				{Plain: "遊ばされる", Formal: "遊ばされます", PlainNegative: "遊ばされない", FormalNegative: "遊ばされません"}},
			Volitional: []Variant{
				{Plain: "遊ぼう", Formal: "遊びましょう", PlainNegative: "遊ぶまい", FormalNegative: "遊びませんまい"}},
			Imperative: []Variant{
				{Plain: "遊べ", Formal: "遊びなさい", PlainNegative: "遊ぶな", FormalNegative: "遊びなさるな"}},
			Conditional: []Variant{
				{Plain: "遊んだら", Formal: "遊びましたら", PlainNegative: "遊ばなかったら", FormalNegative: "遊びませんでしたら"}},
			Alternative: []Variant{
				{Plain: "遊んだり", Formal: "遊びましたり", PlainNegative: "遊ばなかったり", FormalNegative: "遊びませんでしたり"}},
			Continuative: []Variant{
				{Plain: "遊び"}}},
	},
	{
		"Godan verb -gu",
		"游ぐ",
		"v5g",
		Conjugations{
			NonPast: []Variant{
				{Plain: "游ぐ", Formal: "游ぎます", PlainNegative: "游がない", FormalNegative: "游ぎません"}},
			Past: []Variant{
				{Plain: "游いだ", Formal: "游ぎました", PlainNegative: "游がなかった", FormalNegative: "游ぎませんでした"}},
			Conjunctive: []Variant{
				{Plain: "游いで", Formal: "游ぎまして", PlainNegative: "游がなくて", FormalNegative: "游ぎませんで"},
				{PlainNegative: "游がないで"}},
			Provisional: []Variant{
				{Plain: "游げば", Formal: "游ぎますなら", PlainNegative: "游がなければ", FormalNegative: "游ぎませんなら"}},
			Potential: []Variant{
				{Plain: "游げる", Formal: "游げます", PlainNegative: "游げない", FormalNegative: "游げません"}},
			Passive: []Variant{
				{Plain: "游がれる", Formal: "游がれます", PlainNegative: "游がれない", FormalNegative: "游がれません"}},
			Causative: []Variant{
				{Plain: "游がせる", Formal: "游がせます", PlainNegative: "游がせない", FormalNegative: "游がせません"},
				{Plain: "游がす", Formal: "游がします", PlainNegative: "游がさない", FormalNegative: "游がしません"}},
			CausativePassive: []Variant{
				{Plain: "游がせられる", Formal: "游がせられます", PlainNegative: "游がせられない", FormalNegative: "游がせられません"},
				{Plain: "游がされる", Formal: "游がされます", PlainNegative: "游がされない", FormalNegative: "游がされません"}},
			Volitional: []Variant{
				{Plain: "游ごう", Formal: "游ぎましょう", PlainNegative: "游ぐまい", FormalNegative: "游ぎませんまい"}},
			Imperative: []Variant{
				{Plain: "游げ", Formal: "游ぎなさい", PlainNegative: "游ぐな", FormalNegative: "游ぎなさるな"}},
			Conditional: []Variant{
				{Plain: "游いだら", Formal: "游ぎましたら", PlainNegative: "游がなかったら", FormalNegative: "游ぎませんでしたら"}},
			Alternative: []Variant{
				{Plain: "游いだり", Formal: "游ぎましたり", PlainNegative: "游がなかったり", FormalNegative: "游ぎませんでしたり"}},
			Continuative: []Variant{
				{Plain: "游ぎ"}}},
	},
}

func TestConjugations(t *testing.T) {
	for _, tt := range conjugationTests {
		if diff := pretty.Compare(Conjugate(tt.word, tt.pos), &tt.expected); diff != "" {
			t.Errorf("%s --\n%s", tt.desc, diff)
		}
	}
}
