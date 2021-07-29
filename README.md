# Inductive Logic Programming (ILP) Grammar and Linter

This defines a simple grammar (`cmd/ILPLang.g4`) and a command-line
tool which can be used to lint for problems in dataset formatting.

## Usage

The target is a `linter` binary to help point out issues when tokenizing
or parsing a dataset.

**Example 1: No Errors**

When the dataset is well-formatted, nothing is returned:

```prolog
smokes(person1).
friends(person1,person2).
friends(person2,person1).
```

```console
./linter -tokens -file=examples/input.txt
./linter -tokens -file=examples/input.txt
# (No output for either case)
```

**Example 2: Bad Data**

When there is something in the data that cannot be recognized, problems
are directed to stderr:

```
friends(person1,person2).
Bad Data.
```

```console
$ ./linter -tokens -file=examples/bad_data.txt
line 2:0 token recognition error at: 'B'
line 2:3 token recognition error at: ' '
line 2:4 token recognition error at: 'D'
$ ./linter -file=examples/bad_datainput.txt
line 2:0 token recognition error at: 'B'
```

## Building

Building requires a [Go compiler](https://golang.org/).

```
cd cmd
go build
```

A copy of the generated ANTLR parser files are committed to the repository,
and rebuilding them requires an [ANTLR Parser Generator](https://www.antlr.org/).

```
make clean
make linter
```

## Limitations

This grammar is extremely conservative currently: the only tokens
allowed are lowercase characters and underscores, and everything
must start with a lowercase character.

```prolog
a(x_1,y_1).
b(x_1).
```
