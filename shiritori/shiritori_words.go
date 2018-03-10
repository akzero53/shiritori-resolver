package shiritori

type ShiritoriWords struct {
	Words []ShiritoriWord
}

func InitShiritoriWords(words []string) ShiritoriWords {
	sws := ShiritoriWords{Words: make([]ShiritoriWord, len(words))}
	for i, word := range words {
		sws.Words[i] = ShiritoriWord{Word: word}
	}
	return sws
}

func Sort(words *[]ShiritoriWord) {

}
