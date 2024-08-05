package ja

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Romaji to Hiragana  mappings
var romajiToHiragana = map[string]string{
	// Basic syllables
	"a": "あ", "i": "い", "u": "う", "e": "え", "o": "お",
	"ka": "か", "ki": "き", "ku": "く", "ke": "け", "ko": "こ",
	"sa": "さ", "shi": "し", "su": "す", "se": "せ", "so": "そ",
	"ta": "た", "chi": "ち", "tsu": "つ", "te": "て", "to": "と",
	"na": "な", "ni": "に", "nu": "ぬ", "ne": "ね", "no": "の",
	"ha": "は", "hi": "ひ", "fu": "ふ", "he": "へ", "ho": "ほ",
	"ma": "ま", "mi": "み", "mu": "む", "me": "め", "mo": "も",
	"ya": "や", "yu": "ゆ", "yo": "よ",
	"ra": "ら", "ri": "り", "ru": "る", "re": "れ", "ro": "ろ",
	"wa": "わ", "wo": "を", "n": "ん",

	// Voiced sounds
	"ga": "が", "gi": "ぎ", "gu": "ぐ", "ge": "げ", "go": "ご",
	"za": "ざ", "ji": "じ", "zu": "ず", "ze": "ぜ", "zo": "ぞ",
	"da": "だ", "dji": "ぢ", "dzu": "づ", "de": "で", "do": "ど",
	"ba": "ば", "bi": "び", "bu": "ぶ", "be": "べ", "bo": "ぼ",
	"pa": "ぱ", "pi": "ぴ", "pu": "ぷ", "pe": "ぺ", "po": "ぽ",

	// Yoon sounds (combination sounds)
	"kya": "きゃ", "kyu": "きゅ", "kyo": "きょ",
	"sha": "しゃ", "shu": "しゅ", "sho": "しょ",
	"cha": "ちゃ", "chu": "ちゅ", "cho": "ちょ",
	"nya": "にゃ", "nyu": "にゅ", "nyo": "にょ",
	"hya": "ひゃ", "hyu": "ひゅ", "hyo": "ひょ",
	"mya": "みゃ", "myu": "みゅ", "myo": "みょ",
	"rya": "りゃ", "ryu": "りゅ", "ryo": "りょ",
	"gya": "ぎゃ", "gyu": "ぎゅ", "gyo": "ぎょ",
	"ja": "じゃ", "ju": "じゅ", "jo": "じょ",
	"bya": "びゃ", "byu": "びゅ", "byo": "びょ",
	"pya": "ぴゃ", "pyu": "ぴゅ", "pyo": "ぴょ",
}

// Romaji to Katakana  mappings
var romajiToKatakana = map[string]string{
	// Basic syllables
	"a": "ア", "i": "イ", "u": "ウ", "e": "エ", "o": "オ",
	"ka": "カ", "ki": "キ", "ku": "ク", "ke": "ケ", "ko": "コ",
	"sa": "サ", "shi": "シ", "su": "ス", "se": "セ", "so": "ソ",
	"ta": "タ", "chi": "チ", "tsu": "ツ", "te": "テ", "to": "ト",
	"na": "ナ", "ni": "ニ", "nu": "ヌ", "ne": "ネ", "no": "ノ",
	"ha": "ハ", "hi": "ヒ", "fu": "フ", "he": "ヘ", "ho": "ホ",
	"ma": "マ", "mi": "ミ", "mu": "ム", "me": "メ", "mo": "モ",
	"ya": "ヤ", "yu": "ユ", "yo": "ヨ",
	"ra": "ラ", "ri": "リ", "ru": "ル", "re": "レ", "ro": "ロ",
	"wa": "ワ", "wo": "ヲ", "n": "ン",

	// Voiced sounds
	"ga": "ガ", "gi": "ギ", "gu": "グ", "ge": "ゲ", "go": "ゴ",
	"za": "ザ", "ji": "ジ", "zu": "ズ", "ze": "ゼ", "zo": "ゾ",
	"da": "ダ", "dji": "ヂ", "dzu": "ヅ", "de": "デ", "do": "ド",
	"ba": "バ", "bi": "ビ", "bu": "ブ", "be": "ベ", "bo": "ボ",
	"pa": "パ", "pi": "ピ", "pu": "プ", "pe": "ペ", "po": "ポ",

	// Yōon sounds (combination sounds)
	"kya": "キャ", "kyu": "キュ", "kyo": "キョ",
	"sha": "シャ", "shu": "シュ", "sho": "ショ",
	"cha": "チャ", "chu": "チュ", "cho": "チョ",
	"nya": "ニャ", "nyu": "ニュ", "nyo": "ニョ",
	"hya": "ヒャ", "hyu": "ヒュ", "hyo": "ヒョ",
	"mya": "ミャ", "myu": "ミュ", "myo": "ミョ",
	"rya": "リャ", "ryu": "リュ", "ryo": "リョ",
	"gya": "ギャ", "gyu": "ギュ", "gyo": "ギョ",
	"ja": "ジャ", "ju": "ジュ", "jo": "ジョ",
	"bya": "ビャ", "byu": "ビュ", "byo": "ビョ",
	"pya": "ピャ", "pyu": "ピュ", "pyo": "ピョ",
}

func RomajiToKana(romaji string, to string) string {
	var kanaMap map[string]string
	var smallTsu string
	var choonpu string

	if to == "hiragana" {
		kanaMap = romajiToHiragana
		smallTsu = "っ"
		choonpu = "" // No lengthening character in Hiragana
	} else if to == "katakana" {
		kanaMap = romajiToKatakana
		smallTsu = "ッ"
		choonpu = "ー"
	} else {
		return "Invalid conversion type"
	}

	kana := ""
	i := 0

	for i < len(romaji) {
		// Handle double consonants (っ/ッ)
		if i+1 < len(romaji) && romaji[i] == romaji[i+1] {
			kana += smallTsu
			i++
			continue
		}

		// Check for matching substrings from longest to shortest
		found := false
		for length := 3; length > 0; length-- {
			if i+length > len(romaji) {
				continue
			}
			substr := romaji[i : i+length]
			if char, exists := kanaMap[substr]; exists {
				kana += char
				i += length
				found = true
				break
			}
		}

		// If no match, just add the character itself (for unsupported cases)
		if !found {
			kana += string(romaji[i])
			i++
			continue
		}

		// Handle vowel lengthening (ー)
		if to == "katakana" && i < len(romaji) && romaji[i] == romaji[i-1] {
			kana += choonpu
			i++
			continue
		}
	}

	return kana
}

func PrintPWD() {
	fmt.Println("PWD:", os.Getenv("PWD"))
}

/*
 * Sample code for loading kanji data from a JSON file
 [
    [
        "亜",
        "ア",
        "つ.ぐ",
        "jouyou",
        [
            "Asia",
            "rank next",
            "come after",
            "-ous"
        ],
        {
            "deroo": "3273",
            "four_corner": "1010.6",
            "freq": "1509",
            "gakken": "1331",
            "grade": "8",
            "halpern_kkd": "4354",
            "halpern_kkld": "2204",
            "halpern_kkld_2ed": "2966",
            "halpern_njecd": "3540",
            "heisig": "1809",
            "heisig6": "1950",
            "henshall": "997",
            "jf_cards": "1032",
            "jis208": "1-16-01",
            "jlpt": "1",
            "kanji_in_context": "1818",
            "kodansha_compact": "35",
            "maniette": "1827",
            "moro": "272",
            "nelson_c": "43",
            "nelson_n": "81",
            "oneill_kk": "1788",
            "oneill_names": "525",
            "sh_desc": "0a7.14",
            "sh_kk": "1616",
            "sh_kk2": "1724",
            "skip": "4-7-1",
            "strokes": "7",
            "tutt_cards": "1092",
            "ucs": "4e9c"
        }
    ]
 ]

*/

type KanjiData struct {
	Kanji          string
	Onyomi         []string
	Kunyomi        []string
	Type           string
	Meanings       []string
	AdditionalInfo map[string]string
}

type JMDictWord struct {
	Word       string
	Reading    string
	Category   string
	Meanings   []string
	Identifier int
}

func toStringSlice(data interface{}) []string {
	var result []string
	for _, v := range data.([]interface{}) {
		result = append(result, v.(string))
	}
	return result
}

func toStringMap(data interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range data.(map[string]interface{}) {
		result[k] = v.(string)
	}
	return result
}

func LoadKanjiFromJsonFile(filename string) ([]KanjiData, error) {
	// Load kanji from JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var kanji [][]interface{}
	var kanjiDataList []KanjiData

	err = json.Unmarshal(data, &kanji)
	if err != nil {
		log.Fatal(err)
	}

	for _, k := range kanji {
		onyomi := k[1].(string)
		kunyomi := k[2].(string)
		kanjiData := KanjiData{
			Kanji:          k[0].(string),
			Onyomi:         strings.Fields(onyomi),
			Kunyomi:        strings.Fields(kunyomi),
			Type:           k[3].(string),
			Meanings:       toStringSlice(k[4]),
			AdditionalInfo: toStringMap(k[5]),
		}
		kanjiDataList = append(kanjiDataList, kanjiData)
	}
	return kanjiDataList, nil
}

func LoadJMDictFromJsonFile(filename string) ([]JMDictWord, error) {
	// Load kanji from JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var iwords [][]interface{}
	var words []JMDictWord

	err = json.Unmarshal(data, &iwords)
	if err != nil {
		log.Fatal(err)
	}

	for _, iw := range iwords {
		words = append(words, JMDictWord{
			Word:       iw[0].(string),
			Reading:    iw[1].(string),
			Category:   iw[2].(string),
			Meanings:   toStringSlice(iw[5]),
			Identifier: int(iw[6].(float64)),
		})
	}

	return words, nil
}

func GetKanji(kanjiDataList []KanjiData, kanji string) []KanjiData {
	var results []KanjiData
	for _, kanjiData := range kanjiDataList {
		if kanjiData.Kanji == kanji {
			results = append(results, kanjiData)
		}
	}
	return results
}

func SearchKanjiByMeaning(kanjiDataList []KanjiData, meaning string) []KanjiData {
	var results []KanjiData
	for _, kanjiData := range kanjiDataList {
		for _, m := range kanjiData.Meanings {
			if m == meaning {
				results = append(results, kanjiData)
			}
		}
	}
	return results
}

func SearchKanjiByReading(kanjiDataList []KanjiData, reading string, type_ string) []KanjiData {
	var results []KanjiData
	if type_ == "onyomi" {
		for _, kanjiData := range kanjiDataList {
			for _, onyomi := range kanjiData.Onyomi {
				if strings.Replace(onyomi, ".", "", -1) == reading {
					results = append(results, kanjiData)
				}
			}
		}
	} else if type_ == "kunyomi" {
		for _, kanjiData := range kanjiDataList {
			for _, kunyomi := range kanjiData.Kunyomi {
				if strings.Replace(kunyomi, ".", "", -1) == reading {
					results = append(results, kanjiData)
				}
			}
		}
	}
	return results
}

func SearchJMDictByMeaning(words []JMDictWord, meaning string) []JMDictWord {
	var results []JMDictWord
	for _, word := range words {
		for _, m := range word.Meanings {
			if strings.ToLower(m) == meaning {
				results = append(results, word)
			}
		}
	}
	return results
}

func SearchJMDictByReading(words []JMDictWord, reading string) []JMDictWord {
	var results []JMDictWord
	if strings.Replace(reading, ".", "", -1) == reading {
		for _, word := range words {
			if word.Reading == reading {
				results = append(results, word)
			}
		}
	}
	return results
}

func GetJMDictyWord(words []JMDictWord, word string) []JMDictWord {
	var results []JMDictWord
	for _, w := range words {
		if w.Word == word {
			results = append(results, w)
		}
	}
	return results
}
