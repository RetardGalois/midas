package parser

import (
	"fmt"
	"errors"
)

type TokenType int

const (
	Host TokenType = iota
	Username
	Password
)

type Token struct {
	Type TokenType 
	col  int
	Token string
}

type Tokens struct {
	Tokens []Token
}

func absInt(value int) int {
	if value < 0 {
	   return -value
	}
	return value
}

func min(x int, y int) int {
	if x < y {
	   return x
	}
	return y
}

func (t Tokens) IsRightOriented() bool {
	if len(t.Tokens) > 0 {
		return t.Tokens[0].Type == Host
	}
	return false
}

func (t Tokens) IsLeftOriented() bool {
	if len(t.Tokens) > 1 {
		return t.Tokens[len(t.Tokens)-1].Type == Host
	}
	return false
}

func (t Tokens) ParseRight(offset int, _type TokenType, max_steps int) (string, error) {
	if offset + max_steps > len(t.Tokens) {
		max_steps = len(t.Tokens) - offset - 1
	}
	if t.Tokens[offset].Type == _type {
		return "", errors.New(fmt.Sprintf("the element in offset %d is the same type as the choosen element type.", offset))
	}
	i := 1
	for i <= max_steps  {
		// FIXME: offset + i out of range. Maybe is fixed; don't know.
		if t.Tokens[offset + i].Type == _type {
			return t.Tokens[offset + i].Token, nil
		}
		i++
	}
	return "", errors.New(fmt.Sprintf("starting from the offset %d, unable to find a element in %d steps.", offset, max_steps))
}

func (t Tokens) ParseLeft(offset int, _type TokenType, max_steps int) (string, error) {
	if offset - max_steps < 0 {
		max_steps = offset
	}
	if t.Tokens[offset].Type == _type {
		return "", errors.New(fmt.Sprintf("the element in offset %d is the same type as the choosen element type.", offset))
	}
	i := 1
	for i <= max_steps {
		// FIXME: offset - i out of range. Maybe is fixed; don't know.
		if t.Tokens[offset - i].Type == _type {
			return t.Tokens[offset - i].Token, nil
		}
		i++
	}
	return "", errors.New(fmt.Sprintf("starting from the offset %d, unable to find a element in %d steps.", offset, max_steps))
}

func (t Tokens) ParseClosest(offset int, _type TokenType, max_steps int) (string, error) {
	max_steps = min(offset + max_steps, len(t.Tokens) - offset - 1)
	if t.Tokens[offset].Type == _type {
		return "", errors.New(fmt.Sprintf("the element in offset %d is the same type as the choosen element type.", offset))
	}
	i := 1
	abs_right := 0
	abs_left := 0
	offset_right := 0
	offset_left := 0
	for i <= max_steps {
		if t.Tokens[offset + i].Type == _type && abs_right == 0 {
			abs_right = absInt(t.Tokens[offset].col - t.Tokens[offset + i].col)
			offset_right = i
		}
		if t.Tokens[offset - i].Type == _type && abs_left == 0 {
			abs_left = absInt(t.Tokens[offset].col - t.Tokens[offset - i].col)
			offset_left = i
		}
		i++
	}
	if abs_right == 0 || abs_left == 0 {
		return "", errors.New(fmt.Sprintf("can't find a element in %d steps.", max_steps))
	}
	if abs_right == abs_left {
		return "", errors.New(fmt.Sprintf("distance in right and left direction is the same.", max_steps))
	}
	if abs_right < abs_left {
		return t.Tokens[offset + offset_right].Token, nil
	} else {
		return t.Tokens[offset - offset_left].Token, nil
	}
}

func (t Tokens) Display() {
	for i, Token := range t.Tokens {
		fmt.Println(fmt.Sprintf("%d - %+v", i, Token))
	}
}
