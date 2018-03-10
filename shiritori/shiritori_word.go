package shiritori

import "github.com/akzero/shiritori-resolver/utils"

type ShiritoriWord struct {
	Word string
}

func (sw *ShiritoriWord) GetWord() string {
	return sw.Word
}

func (sw *ShiritoriWord) GetHeadCharacter() string {
	return utils.GetStringAt(sw.Word, 0)
}

func (sw *ShiritoriWord) GetLastCharacter() string {
	return utils.GetStringAt(sw.Word, len([]rune(sw.Word))-1)
}

func (sw *ShiritoriWord) IsChainable(next ShiritoriWord) bool {
	return sw.GetLastCharacter() == next.GetHeadCharacter()
}
