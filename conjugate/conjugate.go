package conjugate

import (
	"log"
	"unicode"
)

// Conjugations of a word.
type Conjugations struct {
	Causative                                string
	Conditional                              string
	Imperative                               string
	Negative                                 string
	NegativeNominal                          string
	NegativeParticiple                       string
	NegativePast                             string
	NegativePolite                           string
	NegativeProvisionalConditional           string
	NegativeProvisionalConditionalColloquial string
	Nominal                                  string
	Participle                               string
	Passive                                  string
	Past                                     string
	PastPolite                               string
	Polite                                   string
	Potential                                string
	ProvisionalConditional                   string
	Volitional                               string
	VolitionalPolite                         string
	Wish                                     string
	WishNominal                              string
	WishPast                                 string
}

type conjugator struct {
	expectedEnd []string
	conjugate   func([]rune) *Conjugations
}

var conjugators = map[string]conjugator{
	"adj-i": conjugator{[]string{"い"}, iAdjConjugator},
	"v1":    conjugator{[]string{"る"}, fillConjugator(v1Conjugator)},

	"v5b": conjugator{[]string{"ぶ"}, fillConjugator(v5Conjugator(v5bConjugator))},
	"v5g": conjugator{[]string{"ぐ"}, fillConjugator(v5Conjugator(v5gConjugator))},
	"v5k": conjugator{[]string{"く"}, fillConjugator(v5Conjugator(v5kConjugator))},
	// TODO: v5k-s
	"v5m": conjugator{[]string{"む"}, fillConjugator(v5Conjugator(v5mConjugator))},
	"v5n": conjugator{[]string{"ぬ"}, fillConjugator(v5Conjugator(v5nConjugator))},
	"v5r": conjugator{[]string{"る"}, fillConjugator(v5Conjugator(v5rConjugator))},
	// TODO: v5r-i
	"v5s": conjugator{[]string{"す"}, fillConjugator(v5Conjugator(v5sConjugator))},
	"v5t": conjugator{[]string{"つ"}, fillConjugator(v5Conjugator(v5tConjugator))},
	"v5u": conjugator{[]string{"う"}, fillConjugator(v5Conjugator(v5uConjugator))},
	// TODO: v5u-s
	// TODO: v5z
	"vk":   conjugator{[]string{"くる", "来る", "來る"}, fillConjugator(vkConjugator)},
	"vs-i": conjugator{[]string{"する", "為る", "す"}, fillConjugator(vsiConjugator)},
}

// Conjugate a word given its part-of-speech.
func Conjugate(word string, pos string) *Conjugations {
	runeWord := []rune(word)
	c, ok := conjugators[pos]
	if !ok {
		return nil
	}
	if !unicode.In(runeWord[len(runeWord)-1], unicode.Hiragana) {
		return nil
	}
	defer func() {
		if r := recover(); r != nil {
			if r == "checkSuffix" {
				log.Printf("Expected %s %s to end with %s", pos, string(word), c.expectedEnd)
			} else {
				log.Fatal(r)
			}
		}
	}()
	checkSuffix(runeWord, c.expectedEnd)
	return c.conjugate(runeWord)
}

func checkSuffix(word []rune, expected []string) {
	for _, e := range expected {
		suffix := string(word[len(word)-len([]rune(e)):])
		if suffix == e {
			return
		}
	}
	panic("checkSuffix")
}

func cutOff(word []rune, n int) string {
	return string(word[:len(word)-n])
}

func cutOffOne(word []rune) string {
	return cutOff(word, 1)
}

func fillConjugator(inner func([]rune) *Conjugations) func([]rune) *Conjugations {
	return func(word []rune) *Conjugations {
		c := inner(word)
		naiForm := []rune(c.Negative)
		checkSuffix(naiForm, []string{"ない"})
		c.Conditional = c.Past + "ら"
		c.NegativeNominal = cutOff(naiForm, 2) + "なく"
		c.NegativeParticiple = cutOff(naiForm, 2) + "ないで"
		c.NegativePast = cutOff(naiForm, 2) + "なかった"
		c.NegativePolite = c.Nominal + "ません"
		c.NegativeProvisionalConditional = cutOff(naiForm, 2) + "なければ"
		c.NegativeProvisionalConditionalColloquial = cutOff(naiForm, 2) + "なきゃ"
		c.PastPolite = c.Nominal + "ました"
		c.Polite = c.Nominal + "ます"
		c.Wish = c.Nominal + "たい"
		c.WishNominal = c.Nominal + "たく"
		c.WishPast = c.Nominal + "たかった"
		c.VolitionalPolite = c.Nominal + "ましょう"
		return c
	}
}

func iAdjConjugator(word []rune) *Conjugations {
	radical := cutOffOne(word)
	Nominal := radical + "く"
	return &Conjugations{
		Negative:       Nominal + "ない",
		NegativePolite: Nominal + "ありません",
		Nominal:        Nominal,
		Participle:     Nominal + "て",
		Past:           radical + "かった",
		ProvisionalConditional: radical + "ければ",
		Volitional:             radical + "かろう",
	}
}

func v1Conjugator(word []rune) *Conjugations {
	Nominal := cutOffOne(word)
	return &Conjugations{
		Causative:              Nominal + "させる",
		Imperative:             Nominal + "ろ",
		Negative:               Nominal + "ない",
		Nominal:                Nominal,
		Participle:             Nominal + "て",
		Passive:                Nominal + "られる",
		Past:                   Nominal + "た",
		Potential:              Nominal + "れる",
		ProvisionalConditional: Nominal + "れば",
		Volitional:             Nominal + "よう",
	}
}

func v5Conjugator(inner func([]rune) Conjugations) func([]rune) *Conjugations {
	return func(word []rune) *Conjugations {
		c := inner(word)

		Negative := []rune(c.Negative)
		checkSuffix(Negative, []string{"ない"})
		c.Causative = cutOff(Negative, 2) + "せる"
		c.Passive = cutOff(Negative, 2) + "れる"

		Potential := []rune(c.Potential)
		checkSuffix(Potential, []string{"る"})
		c.Imperative = cutOffOne(Potential)
		c.ProvisionalConditional = cutOffOne(Potential) + "ば"
		return &c
	}
}

func v5bConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "ばない",
		Nominal:    root + "び",
		Participle: root + "んで",
		Past:       root + "んだ",
		Potential:  root + "べる",
		Volitional: root + "ぼう",
	}
}

func v5gConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "がない",
		Nominal:    root + "ぎ",
		Participle: root + "いで",
		Past:       root + "いた",
		Potential:  root + "げる",
		Volitional: root + "ごう",
	}
}

func v5kConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "かない",
		Nominal:    root + "き",
		Participle: root + "いて",
		Past:       root + "いた",
		Potential:  root + "ける",
		Volitional: root + "こう",
	}
}

func v5mConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "まない",
		Nominal:    root + "み",
		Participle: root + "んで",
		Past:       root + "んだ",
		Potential:  root + "める",
		Volitional: root + "もう",
	}
}

func v5nConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "なない",
		Nominal:    root + "に",
		Participle: root + "んで",
		Past:       root + "んだ",
		Potential:  root + "ねる",
		Volitional: root + "のう",
	}
}

func v5rConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "らない",
		Nominal:    root + "り",
		Participle: root + "って",
		Past:       root + "った",
		Potential:  root + "れる",
		Volitional: root + "ろう",
	}
}

func v5sConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	Nominal := root + "し"
	return Conjugations{
		Negative:   root + "さない",
		Nominal:    Nominal,
		Participle: Nominal + "て",
		Past:       Nominal + "た",
		Potential:  root + "せる",
		Volitional: root + "ぞう",
	}
}

func v5tConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "たない",
		Nominal:    root + "ち",
		Participle: root + "って",
		Past:       root + "った",
		Potential:  root + "てる",
		Volitional: root + "とう",
	}
}

func v5uConjugator(word []rune) Conjugations {
	root := cutOffOne(word)
	return Conjugations{
		Negative:   root + "わない",
		Nominal:    root + "い",
		Participle: root + "って",
		Past:       root + "った",
		Potential:  root + "える",
		Volitional: root + "おう",
	}
}

func vkConjugator(word []rune) *Conjugations {
	uForm := cutOffOne(word)
	var iForm, oForm string
	if word[len(word)-2] == 'く' {
		iForm = cutOff(word, 2) + "き"
		oForm = cutOff(word, 2) + "こ"
	} else {
		iForm = cutOffOne(word)
		oForm = cutOffOne(word)
	}
	return &Conjugations{
		Causative:              oForm + "させる",
		Imperative:             oForm + "い",
		Negative:               oForm + "ない",
		Nominal:                iForm,
		Participle:             iForm + "て",
		Passive:                oForm + "られる",
		Past:                   iForm + "た",
		Potential:              oForm + "れる",
		ProvisionalConditional: uForm + "れば",
		Volitional:             oForm + "よう",
	}
}

func vsiConjugator(word []rune) *Conjugations {
	var c *Conjugations
	suConj := func(root string) *Conjugations {
		return &Conjugations{
			Causative:              root + "させる",
			Imperative:             root + "しろ",
			Nominal:                root + "し",
			Passive:                root + "される",
			Potential:              root + "できる",
			ProvisionalConditional: root + "すれば",
			Volitional:             root + "しよう",
		}
	}
	if word[len(word)-2] == 'す' {
		c = suConj(cutOff(word, 2))
	} else if word[len(word)-1] == 'す' {
		c = suConj(cutOffOne(word))
	} else {
		c = &Conjugations{Nominal: cutOffOne(word)}
	}
	c.Negative = c.Nominal + "ない"
	c.Participle = c.Nominal + "て"
	c.Past = c.Nominal + "た"
	return c
}
