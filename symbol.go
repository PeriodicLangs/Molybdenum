package main

type Symbol interface {
	Name() string
	Metadata() map[string] string
	SetMetadata(map[string] string)
}

type IDENTIFIER struct {
	NAME     string
	DATATYPE BasicType
}

func (i IDENTIFIER) Name() string {
	return i.NAME
}

func (i IDENTIFIER) Metadata() map[string]string {
  m := make(map[string]string)
  n, err := i.DATATYPE.Name()
  if err != nil {
    panic(err)
  }
	m["DATATYPE"] = n
  return m
}

func (i IDENTIFIER) SetMetadata(map[string]string) {}

var KEYWORDS = []string{"edef", "def", "@import", "mdef"}
type KEYWORD struct {
	NAME   string
	META map[string]string // function name, file import data etc
}

func (k KEYWORD) Name() string {
	return k.NAME
}

func (k KEYWORD) Metadata() map[string]string {
	return k.META
}

func (k KEYWORD) SetMetadata(meta map[string]string) {
  m := deepCopy(k.META)
  for i, v := range meta {
    if m[i] != v {
      m[i] = v
    }
  }
}

func deepCopy(m map[string]string) map[string]string{
  rm := make(map[string]string)
  for i, v := range m {
    rm[i] = v
  }
  return rm
}
