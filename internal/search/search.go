package search

import (
	"sort"
	"strconv"
	"strings"

	"nihongo-search/internal/dictionary"
	"nihongo-search/internal/models"
	"nihongo-search/lang/ja"
)

type SearchData struct {
	KanjiDataList []models.KanjiData
	WordDataList  []models.JMDictWord
}

type Service struct {
	dict *dictionary.Dictionary
}

func NewService(dict *dictionary.Dictionary) *Service {
	return &Service{
		dict: dict,
	}
}

func (s *Service) Search(query string) *SearchData {
	query = strings.TrimSpace(query)
	data := &SearchData{
		KanjiDataList: []models.KanjiData{},
		WordDataList:  []models.JMDictWord{},
	}
	if query == "" {
		return data
	}

	// Search Kanji
	data.KanjiDataList = append(data.KanjiDataList, s.dict.GetKanji(query)...)
	data.KanjiDataList = append(data.KanjiDataList, s.dict.SearchKanjiByMeaning(query)...)

	kunyomi := ja.RomajiToKana(query, "hiragana")
	data.KanjiDataList = append(data.KanjiDataList, s.dict.SearchKanjiByReading(kunyomi, "kunyomi")...)

	onyomi := ja.RomajiToKana(query, "katakana")
	data.KanjiDataList = append(data.KanjiDataList, s.dict.SearchKanjiByReading(onyomi, "onyomi")...)

	// Sort Kanji
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

	// Search Words
	data.WordDataList = append(data.WordDataList, s.dict.GetJMDictWord(query)...)
	data.WordDataList = append(data.WordDataList, s.dict.SearchJMDictByMeaning(query)...)
	data.WordDataList = append(data.WordDataList, s.dict.SearchJMDictByReading(ja.RomajiToKana(query, "hiragana"))...)

	return data
}
