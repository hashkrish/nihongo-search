package search

import (
    "testing"
    "nihongo-search/internal/dictionary"
    "nihongo-search/internal/models"
)

func TestSearch(t *testing.T) {
    dict := dictionary.New()
    // Add some dummy data
    dict.KanjiDetails = append(dict.KanjiDetails, models.KanjiData{
        Kanji: "日",
        Meanings: []string{"day", "sun"},
        AdditionalInfo: map[string]string{"freq": "1"},
    })

    srv := NewService(dict)
    results := srv.Search("day")

    if len(results.KanjiDataList) == 0 {
        t.Errorf("Expected results, got 0")
    }
}
