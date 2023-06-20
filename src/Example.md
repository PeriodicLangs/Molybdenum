# Lexer Example

```
@import hello
@import goodbye as g

edef main() {
    string name = "Joe"
    name = "John"
    hello.Hello(name)
    g.Goodbye("Joe Bloggs")
}

```

lexed as 

```
IMPORT ["filename":"hello", "importedname": "hello"]
IMPORT ["filename":"goodbye", "importedname": "g"]
ENTRYPOINT "main"
BLOCKSTART
VARDEFINE ["ident":"name", "type": "string"]
ASSIGN ["ident":"name", "val":"Joe"]
ASSIGN ["ident":"name", "val":"John"] // NOTE! Check that types are the same
CALL ["importedname": "hello", "functionname": "Hello", "args": ["ident": "name"]]
CALL ["importedname": "g", "functionname": "Goodbye", "args": ["val": "Joe Bloggs"]]
BLOCKEND
```