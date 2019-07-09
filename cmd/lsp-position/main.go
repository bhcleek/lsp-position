package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf16"
	"unicode/utf8"
)

const replacement = '\ufffd'

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	offset := flag.Uint64("offset", 0, "the byte offset")
	flag.Parse()

	out, err := codeUnitOf(os.Stdin, *offset)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}

func codeUnitOf(r io.Reader, offset uint64) (uint64, error) {
	bufr := bufio.NewReader(r)
	col := uint64(0)

	current := uint64(0)

	for current < offset {
		c, _, err := bufr.ReadRune()
		if err != nil {
			return 0, err
		}

		//log.Printf("%c", c)

		i := utf8.RuneLen(c)
		if i == -1 {
			return 0, errors.New("invalid utf-8 code point")
		}
		current += uint64(i)

		col++
		p1, p2 := utf16.EncodeRune(c)
		if !(p1 == replacement && p2 == replacement) {
			col++
		}
	}

	return col, nil
}
