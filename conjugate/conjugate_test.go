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
			Causative:                                "",
			Conditional:                              "",
			Imperative:                               "",
			Negative:                                 "美味しくない",
			NegativeNominal:                          "",
			NegativeParticiple:                       "",
			NegativePast:                             "",
			NegativePolite:                           "美味しくありません",
			NegativeProvisionalConditional:           "",
			NegativeProvisionalConditionalColloquial: "",
			Nominal:                "美味しく",
			Participle:             "美味しくて",
			Passive:                "",
			Past:                   "美味しかった",
			PastPolite:             "",
			Polite:                 "",
			Potential:              "",
			ProvisionalConditional: "美味しければ",
			Volitional:             "美味しかろう",
			VolitionalPolite:       "",
			Wish:                   "",
			WishNominal:            "",
			WishPast:               "",
		},
	},
	{
		"Ichidan Verb",
		"食べる",
		"v1",
		Conjugations{
			Causative:                                "食べさせる",
			Conditional:                              "食べたら",
			Imperative:                               "食べろ",
			Negative:                                 "食べない",
			NegativeNominal:                          "食べなく",
			NegativeParticiple:                       "食べないで",
			NegativePast:                             "食べなかった",
			NegativePolite:                           "食べません",
			NegativeProvisionalConditional:           "食べなければ",
			NegativeProvisionalConditionalColloquial: "食べなきゃ",
			Nominal:                "食べ",
			Participle:             "食べて",
			Passive:                "食べられる",
			Past:                   "食べた",
			PastPolite:             "食べました",
			Polite:                 "食べます",
			Potential:              "食べれる",
			ProvisionalConditional: "食べれば",
			Volitional:             "食べよう",
			VolitionalPolite:       "食べましょう",
			Wish:                   "食べたい",
			WishNominal:            "食べたく",
			WishPast:               "食べたかった",
		},
	},
	{
		"Godan verb -bu",
		"遊ぶ",
		"v5b",
		Conjugations{
			Causative:                                "遊ばせる",
			Conditional:                              "遊んだら",
			Imperative:                               "遊べ",
			Negative:                                 "遊ばない",
			NegativeNominal:                          "遊ばなく",
			NegativeParticiple:                       "遊ばないで",
			NegativePast:                             "遊ばなかった",
			NegativePolite:                           "遊びません",
			NegativeProvisionalConditional:           "遊ばなければ",
			NegativeProvisionalConditionalColloquial: "遊ばなきゃ",
			Nominal:                "遊び",
			Participle:             "遊んで",
			Passive:                "遊ばれる",
			Past:                   "遊んだ",
			PastPolite:             "遊びました",
			Polite:                 "遊びます",
			Potential:              "遊べる",
			ProvisionalConditional: "遊べば",
			Volitional:             "遊ぼう",
			VolitionalPolite:       "遊びましょう",
			Wish:                   "遊びたい",
			WishNominal:            "遊びたく",
			WishPast:               "遊びたかった",
		},
	},
	{
		"Godan verb -gu",
		"游ぐ",
		"v5g",
		Conjugations{
			Causative:                                "游がせる",
			Conditional:                              "游いたら",
			Imperative:                               "游げ",
			Negative:                                 "游がない",
			NegativeNominal:                          "游がなく",
			NegativeParticiple:                       "游がないで",
			NegativePast:                             "游がなかった",
			NegativePolite:                           "游ぎません",
			NegativeProvisionalConditional:           "游がなければ",
			NegativeProvisionalConditionalColloquial: "游がなきゃ",
			Nominal:                "游ぎ",
			Participle:             "游いで",
			Passive:                "游がれる",
			Past:                   "游いた",
			PastPolite:             "游ぎました",
			Polite:                 "游ぎます",
			Potential:              "游げる",
			ProvisionalConditional: "游げば",
			Volitional:             "游ごう",
			VolitionalPolite:       "游ぎましょう",
			Wish:                   "游ぎたい",
			WishNominal:            "游ぎたく",
			WishPast:               "游ぎたかった",
		},
	},
}

func TestConjugations(t *testing.T) {
	for _, tt := range conjugationTests {
		if !reflect.DeepEqual(Conjugate(tt.word, tt.pos), &tt.expected) {
			t.Errorf("%s --\n%+v", tt.desc, Conjugate(tt.word, tt.pos))
		}
	}
}
