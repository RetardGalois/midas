package parser

import (
	"regexp"
	"bufio"
	"strings"
	"io"
)

var (
	url_regex = regexp.MustCompile("(?i)(?:URL|HOST|DOMAIN)\\s*(?:=>|=|:|\\s)")
	user_regex = regexp.MustCompile("(?i)(?:USER|USERNAME|LOGIN)\\s*(?:=>|=|:|\\s)")
	pass_regex = regexp.MustCompile("(?i)(?:PASS|PASSWORD)\\s*(?:=>|=|:|\\s)")
)

func appendToken(line string, i int, Type TokenType, vault *[]Token) {
	var regex *regexp.Regexp
	switch Type {
	case Host:
		regex = url_regex
	case Username:
		regex = user_regex
	case Password:
		regex = pass_regex
	}
	first_match := regex.FindStringSubmatch(line)[0]
	token := strings.Replace(line, first_match, "", -1)
	*vault = append(*vault, Token{
		Type : Type,
		col   : i,
		Token : token, 
	})
}

func Lexer(file io.Reader) Tokens {
	ret := []Token{}
	lines := bufio.NewScanner(file)
	i := 0
	for lines.Scan() {
		line := lines.Text()
		if url_regex.MatchString(line) {
			appendToken(line, i, Host, &ret)
		} else if user_regex.MatchString(line) {
			appendToken(line, i, Username, &ret)
		} else if pass_regex.MatchString(line) {
			appendToken(line, i, Password, &ret)
		}
		i++
	}
	return Tokens {	
		Tokens : ret,
	}
}
