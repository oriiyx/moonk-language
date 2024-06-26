package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/oriiyx/moonk-language/lexer"
	"github.com/oriiyx/moonk-language/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
			fmt.Printf("%+v\n", tok)
		}
		// same as this:
		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
	}
}
