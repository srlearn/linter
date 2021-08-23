# Inductive Logic Programming (ILP) Grammar and Linter

This defines a simple grammar (`cmd/ILPLang.g4`) and a command-line
tool which can be used to lint for problems in dataset formatting.

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/srlearn/linter)](https://github.com/srlearn/linter/releases)
[![GitHub](https://img.shields.io/github/license/srlearn/linter)](https://github.com/srlearn/linter/blob/main/LICENSE)
[![Test Parsing](https://github.com/srlearn/linter/actions/workflows/package-test.yml/badge.svg)](https://github.com/srlearn/linter/actions/workflows/package-test.yml)

## Overview

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
./linter -tokens -file=examples/pos/pos1.txt
./linter -file=examples/pos/pos1.txt
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
./linter -tokens -file=examples/neg/neg1.txt
line 2:0 token recognition error at: 'B'
line 2:3 token recognition error at: ' '
line 2:4 token recognition error at: 'D'
./linter -file=examples/neg/neg1.txt
line 2:0 token recognition error at: 'B'
line 2:3 token recognition error at: ' '
line 2:4 token recognition error at: 'D'
line 2:5 missing '(' at 'ata'
line 2:8 mismatched input '.' expecting {')', ','}
```

**Example 3: Regression Examples**

The parser can also look for `regressionExample` values, used in regression
data sets.

The parser **will not** check whether an *entire* dataset is correct
(`regressionExample` in labeled as positive, empty negative examples, and
facts). But this could be accomplished fairly easily elsewhere.

```prolog
regressionExample(medv(id100),33.2).
regressionExample(medv(id101),27.5).
regressionExample(medv(id10),18.9).
regressionExample(medv(id102),26.5).
```

## Usage

### Download a Binary

Precompiled binaries are listed on the
[GitHub Releases page](https://github.com/srlearn/linter/releases),
and the latest version can be downloaded with these links:

| Platform | Link |
| :--- | :--- |
| Linux/amd64 | [Download](https://github.com/srlearn/linter/releases/latest/download/linter-linux-amd64) |
| macOS/amd64 | [Download](https://github.com/srlearn/linter/releases/latest/download/linter-darwin-amd64) |
| Windows/amd64 | [Download](https://github.com/srlearn/linter/releases/latest/download/linter-windows-amd64.exe) |

### Build from Source

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
allowed are lowercase characters, integers, and underscores.

```prolog
a(x_1,y_1).
b(x_1).
```

## Contributions

- [Alexander L. Hayes](https://hayesall.com) - *Indiana University, Bloomington*

Some ideas were taken from the `FOPC_MLN_ILP_Parser` developed by
Jude Shavlik and Trevor Walker (and possibly contributed to by many others
who went unnamed in the source code). There are a few versions of their
Tokenizers
([StreamTokenizerJWS](https://github.com/hayesall/SRLBoost/blob/master/src/main/java/edu/wisc/cs/will/FOPC_MLN_ILP_Parser/StreamTokenizerJWS.java)
and
[StreamTokenizerTAW](https://github.com/hayesall/SRLBoost/blob/master/src/main/java/edu/wisc/cs/will/FOPC_MLN_ILP_Parser/StreamTokenizerTAW.java))
and [Parser](https://github.com/hayesall/SRLBoost/blob/master/src/main/java/edu/wisc/cs/will/FOPC_MLN_ILP_Parser/FileParser.java)
currently used in other projects.
