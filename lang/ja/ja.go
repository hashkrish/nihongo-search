package ja

// Romaji to Hiragana and Katakana mappings
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

	// Yōon sounds (combination sounds)
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
		}

		// Handle vowel lengthening (ー)
		if to == "katakana" && i < len(romaji) && romaji[i] == romaji[i-1] {
			kana += choonpu
		}
	}

	return kana
}
