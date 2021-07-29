# Copyright 2021 Alexander L. Hayes

.PHONY : clean

linter: cmd/parser
	cd cmd; go build .
	mv cmd/linter .

cmd/parser: cmd/ILPLang.g4
	cd cmd; antlr4 -Dlanguage=Go -o parser ILPLang.g4

clean:
	rm -f linter
	rm -rf cmd/parser
