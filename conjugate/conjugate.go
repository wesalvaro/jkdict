package conjugate

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// Conjugations of a word.
type Conjugations struct {
	NonPast          []Variant
	Past             []Variant
	Conjunctive      []Variant
	Provisional      []Variant
	Potential        []Variant
	Passive          []Variant
	Causative        []Variant
	CausativePassive []Variant
	Volitional       []Variant
	Imperative       []Variant
	Conditional      []Variant
	Alternative      []Variant
	Continuative     []Variant
}

// Variant shows the variants for a conjugation.
type Variant struct {
	Plain, Formal, PlainNegative, FormalNegative string
}

func loadCsv(reader io.Reader) [][]string {
	r := csv.NewReader(reader)
	r.Comma = '\t'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

type posID int

func loadPos() map[string]posID {
	f, err := os.Open("../jconj/data/kwpos.csv")
	if err != nil {
		log.Fatal(err)
	}
	result := make(map[string]posID)
	for _, record := range loadCsv(f) {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		result[record[1]] = posID(id)
	}
	delete(result, "n")
	delete(result, "adj-na")
	delete(result, "vs")
	return result
}

type conjoData struct {
	stem                int
	okuri, euphr, euphk string
	pos2                int
}
type conjoKey struct {
	pos      posID
	conj     conjID
	neg, fml bool
	onum     int
}

func loadConjo() map[conjoKey]conjoData {
	f, err := os.Open("../jconj/data/conjo.csv")
	if err != nil {
		log.Fatal(err)
	}
	result := make(map[conjoKey]conjoData)
	checkAtoi := func(a string) int {
		i, err := strconv.Atoi(a)
		if err != nil {
			return 0
		}
		return i
	}
	checkParseBool := func(a string) bool {
		b, err := strconv.ParseBool(a)
		if err != nil {
			log.Fatal(err)
		}
		return b
	}
	for i, record := range loadCsv(f) {
		if i == 0 {
			continue
		}
		result[conjoKey{
			posID(checkAtoi(record[0])),
			conjID(checkAtoi(record[1])),
			checkParseBool(record[2]),
			checkParseBool(record[3]),
			checkAtoi(record[4]),
		}] = conjoData{
			checkAtoi(record[5]),
			record[6], record[7], record[8],
			checkAtoi(record[9]),
		}
	}
	return result
}

var (
	conjos = loadConjo()
	posIDs = loadPos()
)

type conjID int

const (
	nonPast conjID = 1 + iota
	past
	conjunctive
	provisional
	potential
	passive
	causative
	causativePassive
	volitional
	imperative
	conditional
	alternative
	continuative
)

var conjs = []conjID{
	nonPast,
	past,
	conjunctive,
	provisional,
	potential,
	passive,
	causative,
	causativePassive,
	volitional,
	imperative,
	conditional,
	alternative,
	continuative,
}

// Conjugate a word given its part-of-speech.
func Conjugate(word string, pos string) *Conjugations {
	runeWord := []rune(word)
	if len(runeWord) < 2 {
		return nil
	}
	posID := posIDs[pos]
	return &Conjugations{
		NonPast:          conjugate(nonPast, runeWord, posID),
		Past:             conjugate(past, runeWord, posID),
		Conjunctive:      conjugate(conjunctive, runeWord, posID),
		Provisional:      conjugate(provisional, runeWord, posID),
		Potential:        conjugate(potential, runeWord, posID),
		Passive:          conjugate(passive, runeWord, posID),
		Causative:        conjugate(causative, runeWord, posID),
		CausativePassive: conjugate(causativePassive, runeWord, posID),
		Volitional:       conjugate(volitional, runeWord, posID),
		Imperative:       conjugate(imperative, runeWord, posID),
		Conditional:      conjugate(conditional, runeWord, posID),
		Alternative:      conjugate(alternative, runeWord, posID),
		Continuative:     conjugate(continuative, runeWord, posID),
	}
}

func conjugate(conj conjID, runeWord []rune, pos posID) []Variant {
	var variants []Variant
	for onum := 1; onum < 10; onum++ {
		variant := Variant{}
		set := false
		for neg := 0; neg <= 1; neg++ {
			for fml := 0; fml <= 1; fml++ {
				key := conjoKey{pos, conj, neg == 1, fml == 1, onum}
				c, ok := conjos[key]
				if !ok {
					break
				}
				set = true
				result := construct(runeWord, c)
				if neg == 1 {
					if fml == 1 {
						variant.FormalNegative = result
					} else {
						variant.PlainNegative = result
					}
				} else {
					if fml == 1 {
						variant.Formal = result
					} else {
						variant.Plain = result
					}
				}
			}
		}
		if set {
			variants = append(variants, variant)
		}
	}
	return variants
}

func construct(runeWord []rune, c conjoData) string {
	iskana := runeWord[len(runeWord)-2] > 'あ' && runeWord[len(runeWord)-2] <= 'ん'
	if iskana && c.euphr != "" || !iskana && c.euphk != "" {
		c.stem++
	}
	var result string
	if iskana {
		result = string(runeWord[:len(runeWord)-c.stem]) + c.euphr + c.okuri
	} else {
		result = string(runeWord[:len(runeWord)-c.stem]) + c.euphk + c.okuri
	}
	return result
}
