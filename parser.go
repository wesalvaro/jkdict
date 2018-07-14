package main

import (
	"compress/gzip"
	"encoding/xml"
	"log"
	"os"

	"./conjugate"
)

var entityNames = []string{
	"Buddh", "MA", "Shinto", "X", "abbr", "adj-f", "adj-i", "adj-ix", "adj-kari",
	"adj-ku", "adj-na", "adj-nari", "adj-no", "adj-pn", "adj-shiku", "adj-t",
	"adv", "adv-to", "anat", "arch", "archit", "astron", "ateji", "aux", "aux-adj",
	"aux-v", "baseb", "biol", "bot", "bus", "chem", "chn", "col", "comp", "conj",
	"cop-da", "ctr", "derog", "eK", "econ", "ek", "engr", "exp", "fam", "fem",
	"finc", "food", "geol", "geom", "gikun", "hob", "hon", "hum", "iK", "id", "ik",
	"int", "io", "iv", "joc", "ksb", "ktb", "kyb", "kyu", "law", "ling", "m-sl",
	"mahj", "male", "male-sl", "math", "med", "mil", "music", "n", "n-adv", "n-pr",
	"n-pref", "n-suf", "n-t", "nab", "num", "oK", "obs", "obsc", "oik", "ok",
	"on-mim", "osb", "physics", "pn", "poet", "pol", "pref", "proverb", "prt",
	"quote", "rare", "rkb", "sens", "shogi", "sl", "sports", "suf", "sumo", "thb",
	"tsb", "tsug", "uK", "uk", "unc", "v-unspec", "v1", "v1-s", "v2a-s", "v2b-k",
	"v2b-s", "v2d-k", "v2d-s", "v2g-k", "v2g-s", "v2h-k", "v2h-s", "v2k-k",
	"v2k-s", "v2m-k", "v2m-s", "v2n-s", "v2r-k", "v2r-s", "v2s-s", "v2t-k",
	"v2t-s", "v2w-s", "v2y-k", "v2y-s", "v2z-s", "v4b", "v4g", "v4h", "v4k", "v4m",
	"v4n", "v4r", "v4s", "v4t", "v5aru", "v5b", "v5g", "v5k", "v5k-s", "v5m",
	"v5n", "v5r", "v5r-i", "v5s", "v5t", "v5u", "v5u-s", "v5uru", "vi", "vk", "vn",
	"vr", "vs", "vs-c", "vs-i", "vs-s", "vt", "vulg", "vz", "yoji", "zool",
}

var xmlEntities map[string]string

func init() {
	xmlEntities = make(map[string]string)
	for _, v := range entityNames {
		xmlEntities[v] = v
	}
}

type reading struct {
	XMLName  xml.Name `xml:"r_ele"`
	Reading  string   `xml:"reb"`
	Info     []string `xml:"re_inf"`
	Priority []string `xml:"re_pri"`
	NoKanji  bool     `xml:"re_nokanji"`
}

type kanji struct {
	XMLName  xml.Name `xml:"k_ele"`
	Kanji    string   `xml:"keb"`
	Info     []string `xml:"ke_inf"`
	Priority []string `xml:"ke_pri"`
}

type sense struct {
	XMLName      xml.Name `xml:"sense"`
	Reference    []string `xml:"xref"`
	Antonym      []string `xml:"ant"`
	PartOfSpeech []string `xml:"pos"`
	Misc         []string `xml:"misc"`
	Field        []string `xml:"field"`
	Gloss        []string `xml:"gloss"`
}

type entry struct {
	XMLName xml.Name  `xml:"entry"`
	Reading []reading `xml:"r_ele"`
	Kanji   []kanji   `xml:"k_ele"`
	Sense   []sense   `xml:"sense"`
}

func (e entry) Conjugate() map[string]*conjugate.Conjugations {
	c := make(map[string]*conjugate.Conjugations)
	for _, pos := range e.Sense[0].PartOfSpeech {
		for _, k := range e.Kanji {
			if c[k.Kanji] == nil {
				c[k.Kanji] = conjugate.Conjugate(k.Kanji, pos)
			}
		}
		for _, r := range e.Reading {
			if c[r.Reading] == nil {
				c[r.Reading] = conjugate.Conjugate(r.Reading, pos)
			}
		}
	}
	return c
}

type jmDict struct {
	XMLName xml.Name `xml:"JMdict"`
	Entry   []entry  `xml:"entry"`
}

type byReading []entry

func (a byReading) Len() int           { return len(a) }
func (a byReading) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byReading) Less(i, j int) bool { return a[i].Reading[0].Reading < a[j].Reading[0].Reading }

func parseDict() *jmDict {
	f, err := os.Open("JMdict_e.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()
	decoder := xml.NewDecoder(gr)
	decoder.Entity = xmlEntities
	dict := jmDict{}
	if err := decoder.Decode(&dict); err != nil {
		log.Fatal(err)
	}
	return &dict
}
