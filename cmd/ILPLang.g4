// Copyright 2021 Alexander L. Hayes
// Apache 2.0 License

/*
Some portions of this are based on the ANTLR `prolog.g4` example,
Copyright (c) 2013, Tom Everett
Used under the terms of a 3-Clause BSD License, copied below:

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. Neither the name of Tom Everett nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

grammar ILPLang;

LPAREN: '(';
RPAREN: ')';
COMMA: ',';
PERIOD: '.';
NEWLINE: [\n];

fragment DIGIT: [0-9];
fragment SMALL_LETTER: [a-z];
fragment ALPHANUMERIC
    : SMALL_LETTER
    | DIGIT
    | '_'
    ;

start
    : term
    | lastterm
    | EOF
    ;

// This recognizes both relations and entities.
// The grammar would be more generally useful if this recognized entities and
// relations on those entities.
OBJECT: SMALL_LETTER (ALPHANUMERIC)*;

term
    : OBJECT LPAREN OBJECT (COMMA OBJECT)* RPAREN PERIOD NEWLINE
    ;

// The 'lastterm' is not required to have a newline after it, but can have an end-of-file.
lastterm
    : OBJECT LPAREN OBJECT (COMMA OBJECT)* RPAREN PERIOD EOF
    ;
