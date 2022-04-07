package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText string
	capitalize []bool
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText, make([]bool, len(plainText))}
}

func (f *FormattedText) Capitalise(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func main() {
	fmt.Printf("app is starting ..")

	text := "This is a brave new world"

	ft := NewFormattedText(text)
	ft.Capitalise(10, 15)
	fmt.Println(ft.String())
}
