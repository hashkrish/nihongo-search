package dictionary

import (
	"encoding/json"
	"os"
	"strings"

	"nihongo-search/internal/models"
)

type Dictionary struct {
	KanjiDetails []models.KanjiData
	JMDictWords  []models.JMDictWord
}

func New() *Dictionary {
	return &Dictionary{
		KanjiDetails: []models.KanjiData{},
		JMDictWords:  []models.JMDictWord{},
	}
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

func (d *Dictionary) LoadKanji(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var kanji [][]interface{}
	err = json.Unmarshal(data, &kanji)
	if err != nil {
		return err
	}

	for _, k := range kanji {
		onyomi := k[1].(string)
		kunyomi := k[2].(string)
		kanjiData := models.KanjiData{
			Kanji:          k[0].(string),
			Onyomi:         strings.Fields(onyomi),
			Kunyomi:        strings.Fields(kunyomi),
			Type:           k[3].(string),
			Meanings:       toStringSlice(k[4]),
			AdditionalInfo: toStringMap(k[5]),
		}
		d.KanjiDetails = append(d.KanjiDetails, kanjiData)
	}
	return nil
}

func (d *Dictionary) LoadJMDict(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var iwords [][]interface{}
	err = json.Unmarshal(data, &iwords)
	if err != nil {
		return err
	}

	for _, iw := range iwords {
		d.JMDictWords = append(d.JMDictWords, models.JMDictWord{
			Word:       iw[0].(string),
			Reading:    iw[1].(string),
			Category:   iw[2].(string),
			Meanings:   toStringSlice(iw[5]),
			Identifier: int(iw[6].(float64)),
		})
	}
	return nil
}

func (d *Dictionary) GetKanji(kanji string) []models.KanjiData {
	var results []models.KanjiData
	for _, kanjiData := range d.KanjiDetails {
		if kanjiData.Kanji == kanji {
			results = append(results, kanjiData)
		}
	}
	return results
}

func (d *Dictionary) SearchKanjiByMeaning(meaning string) []models.KanjiData {
	var results []models.KanjiData
	for _, kanjiData := range d.KanjiDetails {
		for _, m := range kanjiData.Meanings {
			if m == meaning {
				results = append(results, kanjiData)
			}
		}
	}
	return results
}

func (d *Dictionary) SearchKanjiByReading(reading string, type_ string) []models.KanjiData {
	var results []models.KanjiData
	if type_ == "onyomi" {
		for _, kanjiData := range d.KanjiDetails {
			for _, onyomi := range kanjiData.Onyomi {
				if strings.Replace(onyomi, ".", "", -1) == reading {
					results = append(results, kanjiData)
				}
			}
		}
	} else if type_ == "kunyomi" {
		for _, kanjiData := range d.KanjiDetails {
			for _, kunyomi := range kanjiData.Kunyomi {
				if strings.Replace(kunyomi, ".", "", -1) == reading {
					results = append(results, kanjiData)
				}
			}
		}
	}
	return results
}

func (d *Dictionary) SearchJMDictByMeaning(meaning string) []models.JMDictWord {
	var results []models.JMDictWord
	for _, word := range d.JMDictWords {
		for _, m := range word.Meanings {
			if strings.ToLower(m) == meaning {
				results = append(results, word)
			}
		}
	}
	return results
}

func (d *Dictionary) SearchJMDictByReading(reading string) []models.JMDictWord {
	var results []models.JMDictWord
	// Check if reading has no dots. Wait, why check? The original code did:
	// if strings.Replace(reading, ".", "", -1) == reading { ... }
	// This implies that if `reading` has a dot, we don't search.
	// But `reading` passed here is usually RomajiToKana result, which shouldn't have dots.
	// The original code was:
	/*
		if strings.Replace(reading, ".", "", -1) == reading {
			for _, word := range words {
				if word.Reading == reading {
					results = append(results, word)
				}
			}
		}
	*/
	// I'll keep the logic same.
	if strings.Replace(reading, ".", "", -1) == reading {
		for _, word := range d.JMDictWords {
			if word.Reading == reading {
				results = append(results, word)
			}
		}
	}
	return results
}

func (d *Dictionary) GetJMDictWord(word string) []models.JMDictWord {
	var results []models.JMDictWord
	for _, w := range d.JMDictWords {
		if w.Word == word {
			results = append(results, w)
		}
	}
	return results
}
