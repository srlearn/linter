#!/usr/bin/env bash

# Copyright 2021 Alexander L. Hayes
# Apache 2.0 License

# SYNOPSIS:
#   An example for how you might test an input file.
# USAGE:
#   ./lint.sh input.txt

TOKENS=$(./linter -tokens -file=$1 2>&1)
PARSE=$(./linter -file=$1 2>&1)
EXIT=0

if [[ ! -z $TOKENS ]]; then
    echo "$TOKENS"
    EXIT=2
fi

if [[ ! -z $PARSE ]]; then
    echo "$PARSE"
    EXIT=2
fi

exit $EXIT
