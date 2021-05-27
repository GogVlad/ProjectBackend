package backend

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func encryptTest(symbols string, inputFile string, outputFile string) {
	originalText, _ := ioutil.ReadFile(inputFile)
	words, _ := os.Create(outputFile)
	w := bufio.NewWriter(words)
	for i := 0; i < len(originalText)-1; i = i + 2 {
		p1 := strings.Index(symbols, string(originalText[i]))
		p2 := strings.Index(symbols, string(originalText[i+1]))
		if p1 == -1 || p2 == -1 {
			_, _ = fmt.Fprintf(w, "%c%c", originalText[i], originalText[i+1])
		} else {
			l1 := p1 / 9
			c1 := p1 % 9
			l2 := p2 / 9
			c2 := p2 % 9
			p1 = l2*9 + c1
			p2 = l1*9 + c2
			_, _ = fmt.Fprintf(w, "%c%c", symbols[p1], symbols[p2])
		}
	}
	_ = w.Flush()
}
