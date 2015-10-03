package writefile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"types"
)

var dbPath = "./biblioteca"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteBookToFile(b *types.Book) {

	f, err := os.OpenFile(dbPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(b.Titolo + "#" + b.Autore + "\n")

}

func ElencoLibri() {
	// read whole the file
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func ElencoLibri2() map[int]*types.Book {

	books := make(map[int]*types.Book)

	f, _ := os.Open(dbPath)
	defer f.Close()

	bf := bufio.NewReader(f)

	i := 0
	for {
		// reader.ReadLine does a buffered read up to a line terminator,
		// handles either /n or /r/n, and returns just the line without
		// the /r or /r/n.
		line, isPrefix, err := bf.ReadLine()

		// loop termination condition 1:  EOF.
		// this is the normal loop termination condition.
		if err == io.EOF {
			break
		}

		// loop termination condition 2: some other error.
		// Errors happen, so check for them and do something with them.
		if err != nil {
			log.Fatal(err)
		}
		// loop termination condition 3: line too long to fit in buffer
		// without multiple reads.  Bufio's default buffer size is 4K.
		// Chances are if you haven't seen a line terminator after 4k
		// you're either reading the wrong file or the file is corrupt.
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", f.Name())
		}
		// success.  The variable line is now a byte slice based on on
		// bufio's underlying buffer.  This is the minimal churn necessary
		// to let you look at it, but note! the data may be overwritten or
		// otherwise invalidated on the next read.  Look at it and decide
		// if you want to keep it.  If so, copy it or copy the portions
		// you want before iterating in this loop.  Also note, it is a byte
		// slice.  Often you will want to work on the data as a string,
		// and the string type conversion (shown here) allocates a copy of
		// the data.  It would be safe to send, store, reference, or otherwise
		// hold on to this string, then continue iterating in this loop.
		b := &types.Book{}
		l := string(line)
		s := strings.Split(l, "#")
		b.Titolo, b.Autore = s[0], s[1]
		books[i] = b
		i = i + 1
	}

	return books
}
