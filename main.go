package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/akzero/shiritori-resolver/shiritori"
)

func main() {
	bytes, err := ioutil.ReadFile("words.json")
	if err != nil {
		log.Fatal(err)
	}

	var words []string
	if err := json.Unmarshal(bytes, &words); err != nil {
		log.Fatal(err)
	}

	shiritoriWords := shiritori.InitShiritoriWords(words)

	for _, sw := range shiritoriWords.Words {
		fmt.Println("単語: " + sw.GetWord())
		fmt.Println("先頭: " + sw.GetHeadCharacter())
		fmt.Println("末尾: " + sw.GetLastCharacter())
		nextSW := shiritori.ShiritoriWord{Word: "りんご"}
		fmt.Println("「" + nextSW.GetWord() + "」が次に来れるか: " + strconv.FormatBool(sw.IsChainable(nextSW)))
		fmt.Println("-------")
	}
}
