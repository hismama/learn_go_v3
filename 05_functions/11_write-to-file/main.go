package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	p := person{
		first: "Jenny",
	}

	f, err := os.Create("./05_functions/11_write-to-file/output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	// Writes 'first' to the file f
	p.writeOut(f)
	// Writes 'first' to the buffer b
	p.writeOut(&b)
	fmt.Println(b.String())

	//b := bytes.NewBufferString("Hello ")
	//fmt.Println(b.String())
	//b.WriteString("Gophers!")
	//fmt.Println(b.String())
	//b.Reset()
	//fmt.Println(b.String())
	//b.WriteString("It's Friday")
	//fmt.Println(b.String())
}
