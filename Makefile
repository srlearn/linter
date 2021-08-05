# Copyright 2021 Alexander L. Hayes

.PHONY : clean distribution

linter: cmd/parser
	cd cmd; go build .
	mv cmd/linter .

cmd/parser: cmd/ILPLang.g4
	cd cmd; antlr4 -Dlanguage=Go -o parser ILPLang.g4

clean:
	rm -f linter
	rm -rf cmd/parser
	rm -rf dist/

distribution: cmd/parser
	mkdir dist
	cd cmd; env GOOS=windows GOARCH=amd64 go build .; mv linter.exe ../dist/linter-windows-amd64.exe
	cd cmd; env GOOS=darwin GOARCH=amd64 go build .; mv linter ../dist/linter-darwin-amd64
	cd cmd; env GOOS=linux GOARCH=amd64 go build .; mv linter ../dist/linter-linux-amd64
