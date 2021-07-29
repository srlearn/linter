// Copyright 2021 Alexander L. Hayes
// Apache 2.0 License

package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/srlearn/linter/parser"
)

type ilplangListener struct {
	*parser.BaseILPLangListener
}

func main() {

	tokenize := flag.Bool("tokens", false, "Tokenize instead of parsing")
	filename := flag.String("file", "input.txt", "Input filename")
	flag.Parse()

	content, err := ioutil.ReadFile((*filename))
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	is := antlr.NewInputStream(text)
	lexer := parser.NewILPLangLexer(is)

	if *tokenize {
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
		}
	} else {
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewILPLangParser(stream)
		antlr.ParseTreeWalkerDefault.Walk(&ilplangListener{}, p.Start())
	}
}
