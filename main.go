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
var kanjiDetails []ja.KanjiData
var JMDictWords []ja.JMDictWord

type PageData struct {
	Title string
}

type SearchData struct {
	KanjiDataList []ja.KanjiData
	WordDataList  []ja.JMDictWord
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{
		Title: "Nihongo search!",
	}
	tmpl.Execute(w, data)

}

// TODO: Refactor this function to use a single search function
func handlePartialSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	query = strings.TrimSpace(query)
	fmt.Println("Query:", query)

	tmpl := template.Must(template.ParseFiles("partials/search.html"))
	data := SearchData{
		KanjiDataList: []ja.KanjiData{},
		WordDataList:  []ja.JMDictWord{},
	}
	counter++
	if query == "" {
		tmpl.Execute(w, data)
		return
	}

	data.KanjiDataList = append(data.KanjiDataList, ja.GetKanji(kanjiDetails, query)...)
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByMeaning(kanjiDetails, query)...)

	kunyomi := ja.RomajiToKana(query, "hiragana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDetails, kunyomi, "kunyomi")...)

	onyomi := ja.RomajiToKana(query, "katakana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDetails, onyomi, "onyomi")...)

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

	data.WordDataList = append(data.WordDataList, ja.GetJMDictyWord(JMDictWords, query)...)
	data.WordDataList = append(data.WordDataList, ja.SearchJMDictByMeaning(JMDictWords, query)...)
	data.WordDataList = append(data.WordDataList, ja.SearchJMDictByReading(JMDictWords, ja.RomajiToKana(query, "hiragana"))...)
	fmt.Println("Word data list:", data.WordDataList)

	tmpl.Execute(w, data)
}

func handlePangoPartialSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	query = strings.TrimSpace(query)
	fmt.Println("Query:", query)

	tmpl := template.Must(template.ParseFiles("partials/pango/search.html"))
	data := SearchData{
		KanjiDataList: []ja.KanjiData{},
		WordDataList:  []ja.JMDictWord{},
	}
	counter++
	if query == "" {
		tmpl.Execute(w, data)
		return
	}

	data.KanjiDataList = append(data.KanjiDataList, ja.GetKanji(kanjiDetails, query)...)
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByMeaning(kanjiDetails, query)...)

	kunyomi := ja.RomajiToKana(query, "hiragana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDetails, kunyomi, "kunyomi")...)

	onyomi := ja.RomajiToKana(query, "katakana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDetails, onyomi, "onyomi")...)

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

	if len(data.KanjiDataList) > 3 {
		data.KanjiDataList = data.KanjiDataList[:3]
	}

	data.WordDataList = append(data.WordDataList, ja.GetJMDictyWord(JMDictWords, query)...)
	data.WordDataList = append(data.WordDataList, ja.SearchJMDictByMeaning(JMDictWords, query)...)
	data.WordDataList = append(data.WordDataList, ja.SearchJMDictByReading(JMDictWords, ja.RomajiToKana(query, "hiragana"))...)
	fmt.Println("Word data list:", data.WordDataList)

	if len(data.WordDataList) > 3 {
		data.WordDataList = data.WordDataList[:3]
	}

	tmpl.Execute(w, data)
}

func handleTextPartialSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	query = strings.TrimSpace(query)
	fmt.Println("Query:", query)

	tmpl := template.Must(template.ParseFiles("partials/text/search.html"))
	data := SearchData{
		KanjiDataList: []ja.KanjiData{},
		WordDataList:  []ja.JMDictWord{},
	}
	counter++
	if query == "" {
		tmpl.Execute(w, data)
		return
	}

	data.KanjiDataList = append(data.KanjiDataList, ja.GetKanji(kanjiDetails, query)...)
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByMeaning(kanjiDetails, query)...)

	kunyomi := ja.RomajiToKana(query, "hiragana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDetails, kunyomi, "kunyomi")...)

	onyomi := ja.RomajiToKana(query, "katakana")
	data.KanjiDataList = append(data.KanjiDataList, ja.SearchKanjiByReading(kanjiDetails, onyomi, "onyomi")...)

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

	if len(data.KanjiDataList) > 3 {
		data.KanjiDataList = data.KanjiDataList[:3]
	}

	data.WordDataList = append(data.WordDataList, ja.GetJMDictyWord(JMDictWords, query)...)
	data.WordDataList = append(data.WordDataList, ja.SearchJMDictByMeaning(JMDictWords, query)...)
	data.WordDataList = append(data.WordDataList, ja.SearchJMDictByReading(JMDictWords, ja.RomajiToKana(query, "hiragana"))...)
	fmt.Println("Word data list:", data.WordDataList)

	if len(data.WordDataList) > 3 {
		data.WordDataList = data.WordDataList[:3]
	}

	tmpl.Execute(w, data)
}

func handleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	var err error
	var kanjiDetailsList []ja.KanjiData
	for i := 1; i <= 2; i++ {
		filename := "./data/ja/kanji_bank_" + strconv.Itoa(i) + ".json"
		kanjiDetailsList, err = ja.LoadKanjiFromJsonFile(filename)
		if err != nil {
			fmt.Println("Error loading kanji data:", err)
			return
		}
		kanjiDetails = append(kanjiDetails, kanjiDetailsList...)
	}

	var JMDictWordsData []ja.JMDictWord
	for i := 1; i <= 75; i++ {
		filename := "./data/ja/term_bank_" + strconv.Itoa(i) + ".json"
		JMDictWordsData, err = ja.LoadJMDictFromJsonFile(filename)
		if err != nil {
			fmt.Println("Error loading JMDict data:", err)
			return
		}
		JMDictWords = append(JMDictWords, JMDictWordsData...)

	}
	fmt.Println("JMDict words loaded")

	// Route handlers
	http.HandleFunc("GET /", handleHome)
	http.HandleFunc("GET /healthcheck", handleHealthCheck)
	http.HandleFunc("GET /partial/search", handlePartialSearch)
	http.HandleFunc("GET /partial/pango/search", handlePangoPartialSearch)
	http.HandleFunc("GET /partial/text/search", handleTextPartialSearch)

	// Start server
	fmt.Println("Server running on port 8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
