package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
	"unicode"
)

var (
	posTypes = map[string]string{
		"adj-kari":  "adj. かり (archaic)",
		"adj-ku":    "adj. く (archaic)",
		"adj-nari":  "adj. なり (archaic/formal)",
		"adj-shiku": "adj. しく (archaic)",
		"adj-f":     "n. or v. (prenominal)",
		"adj-i":     "adj.",                      //【形容詞】
		"adj-ix":    "adj. (良い・いい)",              //【形容詞】
		"adj-na":    "adj-n. or quasi-adjective", //【形容動詞】
		"adj-no":    "may take possessive の",
		"adj-pn":    "adj-pre-n.", //【連体詞】
		"adj-t":     "adj. 〜たる",
		"adv":       "adv.", //【副詞】
		"adv-to":    "adv. taking the と particle",
		"aux":       "aux.",
		"aux-adj":   "aux. adj.",
		"aux-v":     "aux. v.",
		"conj":      "conj.",
		"cop-da":    "copula",
		"ctr":       "counter",
		"int":       "interjection",   //【感動詞】
		"n":         "n.",             //【普通名詞】
		"n-adv":     "n. (adverbial)", //【副詞的名詞】
		"n-pr":      "n. (proper)",
		"n-pref":    "n. (as prefix)",
		"n-suf":     "n. (as suffix)",
		"n-t":       "n. (temporal)", //【時相名詞】
		"pn":        "pronoun",
		"prt":       "particle",
		"pref":      "prefix",
		"suf":       "suffix",
		// Verbs:
		"iv":       "v. (irr.)",
		"vn":       "v. (irr.)",
		"vr":       "v. (irr.), plain form ends with 〜り",
		"vk":       "v. くる (special)",
		"vs-s":     "v. する (special)",
		"vs-i":     "v. する (irr.)",
		"vs":       "takes する",
		"vs-c":     "precursor to modern する",
		"vi":       "intransitive",
		"vt":       "transitive",
		"v-unspec": "unspecified",
		// Ichidan: 一段
		"v1":   "v1.",
		"v1-s": "v1. くれる (special)",
		"vz":   "v1. ずる (alternative form of -jiru verbs)",
		// Nidan: 二段
		"v2a-s": "v2d. (archaic)",
		"v2b-k": "v2u. (archaic)",
		"v2b-s": "v2d. (archaic)",
		"v2d-k": "v2u. (archaic)",
		"v2d-s": "v2d. (archaic)",
		"v2g-k": "v2u. (archaic)",
		"v2g-s": "v2d. (archaic)",
		"v2h-k": "v2u. (archaic)",
		"v2h-s": "v2d. (archaic)",
		"v2k-k": "v2u. (archaic)",
		"v2k-s": "v2d. (archaic)",
		"v2m-k": "v2u. (archaic)",
		"v2m-s": "v2d. (archaic)",
		"v2n-s": "v2d. (archaic)",
		"v2r-k": "v2u. (archaic)",
		"v2r-s": "v2d. (archaic)",
		"v2s-s": "v2d. (archaic)",
		"v2t-k": "v2u. (archaic)",
		"v2t-s": "v2d. (archaic)",
		"v2w-s": "v2d. (archaic) (ゑ conjugation)",
		"v2y-k": "v2u. (archaic)",
		"v2y-s": "v2d. (archaic)",
		"v2z-s": "v2d. (archaic)",
		// Yodan: 四段
		"v4b": "v4. (archaic)",
		"v4g": "v4. (archaic)",
		"v4h": "v4. (archaic)",
		"v4k": "v4. (archaic)",
		"v4m": "v4. (archaic)",
		"v4n": "v4. (archaic)",
		"v4r": "v4. (archaic)",
		"v4s": "v4. (archaic)",
		"v4t": "v4. (archaic)",
		// Godan: 五段
		"v5aru": "v5. (special)",
		"v5b":   "v5.",
		"v5g":   "v5.",
		"v5k":   "v5.",
		"v5k-s": "v5. Iku/Yuku (special)",
		"v5m":   "v5.",
		"v5n":   "v5.",
		"v5r":   "v5.",
		"v5r-i": "v5. (irr.)",
		"v5s":   "v5.",
		"v5t":   "v5.",
		"v5u":   "v5.",
		"v5u-s": "v5. (special)",
		"v5uru": "v5. Uru (old, Eru)",
		// Terminology:
		"anat":    "anatomy",
		"archit":  "architecture",
		"astron":  "astronomy",
		"baseb":   "baseball",
		"biol":    "biology",
		"bot":     "botany",
		"bus":     "business",
		"econ":    "economics",
		"engr":    "engineering",
		"finc":    "finance",
		"geol":    "geology",
		"joc":     "jocular",
		"law":     "law",
		"mahj":    "mahjong",
		"med":     "medical",
		"music":   "music",
		"Shinto":  "Shinto",
		"shogi":   "shogi",
		"sports":  "sports",
		"sumo":    "sumo",
		"zool":    "zoology",
		"Buddh":   "Buddhist",
		"chem":    "chemistry",
		"chn":     "childish",
		"comp":    "technology",
		"MA":      "martial arts",
		"ling":    "linguistics",
		"proverb": "proverb",
		"physics": "physics",
		"derog":   "derogatory",
		"math":    "math",
		"mil":     "military",
		"food":    "food",
		"geom":    "geometry",
		"poet":    "poetic",
		"num":     "numeric",
		// Class:
		"fem":     "female",
		"male":    "male",
		"sl":      "slang",
		"m-sl":    "manga slang",
		"male-sl": "male slang",
		// Politeness:
		"pol":  "polite",     //【丁寧語】
		"hon":  "respectful", //【尊敬語】
		"hum":  "humble",     //【謙譲語】
		"vulg": "vulgar",
		"X":    "rude or X-rated",
		// Display:
		"eK":    "kanji excl.",
		"uK":    "kanji usu.",
		"ek":    "kana excl.",
		"uk":    "kana usu.",
		"iK":    "irr. kanji",
		"ik":    "irr. kana",
		"io":    "irr. okurigana",
		"ok":    "out-dated kana",
		"oK":    "out-dated kanji",
		"oik":   "old or irr. kana",
		"ateji": "phonetic",        //【当て字】
		"gikun": "special reading", //【義訓】【熟字訓】
		// Dialects:
		"kyb":  "Kyoto",    // 京都弁
		"osb":  "Osaka",    // 大坂弁
		"ksb":  "Kansai",   // 関西弁
		"ktb":  "Kanto",    // 関東弁
		"tsb":  "Tosa",     // 土佐弁
		"thb":  "Tohoku",   // 東北弁
		"tsug": "Tsugaru",  // 津軽弁
		"kyu":  "Kyushu",   // 九州弁
		"rkb":  "Ryukyu",   // 琉球弁
		"nab":  "Nagano",   // 長野弁
		"hob":  "Hokkaido", // 北海道弁
		// Style:
		"abbr":   "abbr.",
		"quote":  "quotation",
		"sens":   "sensitive",
		"on-mim": "onomatopoeia",
		"exp":    "expression",
		"fam":    "familiar",
		"col":    "colloq.",
		"id":     "idiomatic",
		"obs":    "obsolete",
		"obsc":   "obscure",
		"arch":   "archaism",
		"rare":   "rare",
		"yoji":   "4-character phrase", // 四字熟語
		"unc":    "unclassified",
	}
	sectionRange = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x3041, 0x3096, 1},
			{0x30a1, 0x30fa, 1},
		},
		R32: []unicode.Range32{
			{Hi: 0x1b000, Lo: 0x1b000, Stride: 1},
			{Hi: 0x1b001, Lo: 0x1b11e, Stride: 1},
			{Hi: 0x1f200, Lo: 0x1f200, Stride: 1},
		},
	}
)

type section struct {
	ID      rune
	Entries []entry
}

type metaData struct {
	UID       string
	Title     string
	Language  string
	Creator   string
	Copyright string
	InLang    string
	OutLang   string
}

type outputData struct {
	Date     time.Time
	Sections []*section
	MetaData metaData
}

type byID []*section

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func posConvert(pos string) string {
	return posTypes[pos]
}

func createSections(dict *jmDict) []*section {
	sectionMap := make(map[rune]*section)
	for _, e := range dict.Entry {
		var first rune
		for _, c := range e.Reading[0].Reading {
			first = c
			break
		}
		if !unicode.In(first, sectionRange) {
			continue
		}
		s, ok := sectionMap[first]
		if ok {
			s.Entries = append(s.Entries, e)
		} else {
			sectionMap[first] = &section{
				ID:      first,
				Entries: []entry{e},
			}
		}
	}
	sections := make([]*section, 0, len(sectionMap))
	for _, val := range sectionMap {
		sort.Sort(byReading(val.Entries))
		sections = append(sections, val)
	}
	sort.Sort(byID(sections))
	return sections
}

func outputOpf(dict *jmDict) {
	sections := createSections(dict)
	pattern := filepath.Join("tpl", "*.tpl.*")
	funcs := template.FuncMap{
		"posConvert": posConvert,
	}
	tmpl := template.Must(template.New("jkdict").Funcs(funcs).ParseGlob(pattern))

	data := outputData{
		Date: time.Now().Local(),
		MetaData: metaData{
			UID:       "ef5dd3dc-1625-4fd5-860f-77b4c1fa528a",
			Title:     "JKDict",
			Creator:   "Wes Alvaro",
			Copyright: "Wes Alvaro",
			Language:  "ja",
			InLang:    "ja",
			OutLang:   "en",
		},
		Sections: sections,
	}

	var wg sync.WaitGroup
	wg.Add(len(sections) + 1)
	go func() {
		f, err := os.Create("./out/jkdict.opf")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		err = tmpl.ExecuteTemplate(f, "opf.tpl.xml", data)
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
		wg.Done()
	}()
	for _, s := range sections {
		go func(s *section) {
			f, err := os.Create(fmt.Sprintf("./out/section-%v.html", s.ID))
			if err != nil {
				log.Fatal("section file creation: ", err)
			}
			defer f.Close()
			err = tmpl.ExecuteTemplate(f, "section.tpl.html", s)
			if err != nil {
				log.Fatalf("template execution: %s", err)
			}
			wg.Done()
		}(s)
	}
	wg.Wait()
}
