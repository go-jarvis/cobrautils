package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"
)

//go:embed value_int.go.tmpl
var tmpl_value_int string

//go:embed value_flag.go.tmpl
var tmpl_value_flag string

func toCapital(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}

	s = strings.ToLower(s)
	return strings.ToUpper(s[0:1]) + s[1:]
}

func trimRightNumber(s string) string {
	return strings.TrimRight(s, "0123456789")
}

func trimLeftLetter(s string) string {
	s = strings.TrimLeft(s, "uint")
	if s == "" {
		return "0"
	}
	return s
}

var funcMap = template.FuncMap{
	"ToCapital":       toCapital,
	"TrimRightNumber": trimRightNumber,
	"TrimLeftLetter":  trimLeftLetter,
}

func genFlagValues(types []string) {
	for _, typ := range types {
		genFlagValue(typ)
	}
}

func genFlagValue(typ string) {
	data := mustTmpl("pflag value", tmpl_value_int, typ)
	name := fmt.Sprintf("../pflagvalue/%s.go", typ)
	mustWriteFile(name, data)
}

func genFlag(types []string) {
	data := mustTmpl("flags", tmpl_value_flag, types)

	name := fmt.Sprintf("../pflagvalue/value_flag.go")
	mustWriteFile(name, data)
}

func mustTmpl(name string, tmpl string, data interface{}) []byte {
	tp, err := template.New(name).Funcs(funcMap).Parse(tmpl)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBufferString("")
	err = tp.Execute(buf, data)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func mustWriteFile(name string, data []byte) {
	err := os.WriteFile(name, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	types := []string{
		"int8", "int16", "int32",
		"uint", "uint8", "uint16", "uint32", "uint64",
	}

	genFlag(types)
	genFlagValues(types)
}
