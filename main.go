package main

import (
	"fmt"
	"html/template"
	"net/http"
	"nihongo-search/lang/ja"
	"sort"
	"strconv"
	"strings"
)

var counter int
var kanjiDataList []ja.KanjiData

type PageData struct {
	Title string
}

type SearchData struct {
	KanjiDataList []ja.KanjiData
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{
		Title: "Nihongo search!",
	}
	tmpl.Execute(w, data)

}

func handlePartialSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	query = strings.TrimSpace(query)
	fmt.Println("Query:", query)

	tmpl := template.Must(template.ParseFiles("partials/search.html"))
	data := SearchData{
		KanjiDataList: []ja.KanjiData{},
	}
	counter++
	// data.KanjiDataList = ja.SearchKanjiByMeaning(kanjiDataList, query)
	data.KanjiDataList = append(data.KanjiDataList, ja.GetKanji(kanjiDataList, query)...)
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByMeaning(kanjiDataList, query)...)

	kunyomi := ja.RomajiToKana(query, "hiragana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDataList, kunyomi, "kunyomi")...)

	onyomi := ja.RomajiToKana(query, "katakana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDataList, onyomi, "onyomi")...)

	sort.Slice(data.KanjiDataList, func(i, j int) bool {
		iFreq, err := strconv.Atoi(data.KanjiDataList[i].AdditionalInfo["freq"])
		if err != nil {
			iFreq = 99999
		}
		jFreq, err := strconv.Atoi(data.KanjiDataList[j].AdditionalInfo["freq"])
		if err != nil {
			jFreq = 99999
		}
		return iFreq < jFreq
	})
	tmpl.Execute(w, data)
}

func main() {
	// romaji := "konnichiwa"
	// fmt.Println("Hiragana:", ja.RomajiToKana(romaji, "hiragana"))
	// fmt.Println("Katakana:", ja.RomajiToKana(romaji, "katakana"))

	// romaji2 := "atakka"
	// fmt.Println("Katakana with lengthening:", ja.RomajiToKana(romaji2, "katakana"))

	// romaji3 := "nippon"
	// fmt.Println("Hiragana with double consonant:", ja.RomajiToKana(romaji3, "hiragana"))

	// romaji4 := "santana kirishunan"
	// fmt.Println("Katakana with double consonant and lengthening:", ja.RomajiToKana(romaji4, "katakana"))

	// romaji5 := "aa"
	// fmt.Println("Katakana with double consonant and lengthening:", ja.RomajiToKana(romaji5, "katakana"))

	var err error
	kanjiDataList, err = ja.LoadKanjiFromJsonFile("kanji_bank_1.json")
	if err != nil {
		fmt.Println("Error loading kanji data:", err)
		return
	}
	// fmt.Println(ja.SearchKanjiByMeaning(kanjiDataList, "to be"))

	http.HandleFunc("GET /", handleHome)
	http.HandleFunc("GET /partial/search", handlePartialSearch)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
