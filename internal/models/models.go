package models

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
