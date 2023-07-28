package main

import (
	"fmt"
	"os"
	"strings"

	"midas/fs"
	"midas/parser"
)

// TODO: change error handling pattern.
func parse(file string) {
	zipped_content, err := fs.File(file)
	if err != nil {
		panic(fmt.Sprintf("ERROR ON FILE %s: fs.File: %s", file, err))
	}
	// FIXME: don't panic. just log to stderr.
	for _, content := range zipped_content {
		reader := strings.NewReader(content)
		tokens := parser.Lexer(reader)
		if len(tokens.Tokens) == 0 {
			continue
		}
		if tokens.IsRightOriented() {
			for i, t := range tokens.Tokens {
				if t.Type == parser.Host {

					host := t.Token
					username, err := tokens.ParseRight(i, parser.Username, 10)
					if err != nil {
						fmt.Fprintf(os.Stderr, "WARNING ON FILE %s: ParseRight: %s\n", file, err)
						continue
					}
					password, err := tokens.ParseRight(i, parser.Password, 10)
					if err != nil {
						fmt.Fprintf(os.Stderr, "WARNING ON FILE %s: ParseRight: %s\n", file, err)
						continue
					}
					fmt.Println(fmt.Sprintf("%s:%s:%s", strings.Trim(host, " "), strings.Trim(username, " "), strings.Trim(password, " ")))
				}
			}
		} else if tokens.IsLeftOriented() {
			for i, t := range tokens.Tokens {
				if t.Type == parser.Host {
					host := t.Token
					username, err := tokens.ParseLeft(i, parser.Username, 10)
					if err != nil {
						panic(fmt.Sprintf("WARNING ON FILE %s: ParseLeft: %s\n", file, err))
					}
					password, err := tokens.ParseLeft(i, parser.Password, 10)
					if err != nil {
						panic(fmt.Sprintf("WARNING ON FILE %s: ParseLeft: %s\n", file, err))
					}
					fmt.Println(fmt.Sprintf("%s:%s:%s", strings.Trim(host, " "), strings.Trim(username, " "), strings.Trim(password, " ")))
				}
			}
		} else {
			// FIXME: index out of range [-1]
			continue
			for i, t := range tokens.Tokens {
				if t.Type == parser.Host {
					host := t.Token
					username, err := tokens.ParseClosest(i, parser.Username, 10)
					if err != nil {
						panic(fmt.Sprintf("Can't parse username. ERROR: %s", err))
					}
					password, err := tokens.ParseClosest(i, parser.Password, 10)
					if err != nil {
						panic(fmt.Sprintf("Can't parse password. ERROR: %s", err))
					}
					fmt.Println(fmt.Sprintf("%s:%s:%s", strings.Trim(host, " "), strings.Trim(username, " "), strings.Trim(password, " ")))
				}
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("not enough arguments.")
	}
	parse(os.Args[1])
}
