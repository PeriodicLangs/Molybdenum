# Lexer Example

```
@import hello
@import goodbye as g

edef main() {
    string name = "Joe"
    name = "John"
    int number = 10 * 0.5 - 1
    hello.Hello(name)
    g.Goodbye("Joe Bloggs")
}

```

intermedially lexed as 

```
IMPORT
IDENT "hello"
NEWLINE
IMPORT
IDENT "goodbye"
KEYWORD "as"
IDENT "g"
NEWLINE
KEYWORD "edef"
IDENT "main"
LBRAC
RBRAC
BLOCKSTART
NEWLINE
TYPEANNOT "string"
IDENT "name"
ASSIGN
LITERAL "Joe"
NEWLINE
IDENT "name"
ASSIGN
LITERAL "John"
NEWLINE
TYPEANNOT "int"
IDENT "number"
ASSIGN
LITERAL 10
MUL
LITERAL 0.5
SUB
1
NEWLINE
IDENT "hello"
DOT
IDENT "Hello"
LBRAC
IDENT "name"
RBRAC
NEWLINE
IDENT "g"
DOT
IDENT "Goodbye"
LBRAC
LITERAL "Joe Bloggs"
RBRAC
NEWLINE
BLOCKEND
EOF
```

finally lexed as 

```
IMPORT ["filename":"hello", "importedname": "hello"]
IMPORT ["filename":"goodbye", "importedname": "g"]
ENTRYPOINT "main"
BLOCKSTART
VARDEFINE ["ident":"name", "type": "string"]
ASSIGN ["ident":"name", "val":"Joe"]
ASSIGN ["ident":"name", "val":"John"] // NOTE! Check that types are the same
VARDEFINE ["ident":"number", "type": "int"]
ASSIGN ["ident":"number", "val": 4] // implicit int cast
CALL ["importedname": "hello", "functionname": "Hello", "args": ["ident": "name"]]
CALL ["importedname": "g", "functionname": "Goodbye", "args": ["val": "Joe Bloggs"]]
BLOCKEND
EOF
```