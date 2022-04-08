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

type TextRange struct {
	Start, End int
	Capitalise, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText string
	formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)
	return r
}

func main() {
	fmt.Printf("app is starting ..")

	text := "This is a brave new world"

	ft := NewFormattedText(text)
	ft.Capitalise(10, 15)
	fmt.Println(ft.String())
}
