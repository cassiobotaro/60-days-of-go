package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// DatFile is a file in dat format
type DatFile struct{}

// Write persists the content
func (d *DatFile) Write(p []byte) (n int, err error) {
	// open a dat file and write p
	return
}

// PreffixOutput preffix content writed with some preffix
func PreffixOutput(w io.Writer, preffix []byte, content []byte) {
	// receive a writer, then you can pass whatever that implements writer interface
	// it's called polymorphism
	content = append(preffix, content...)
	w.Write(content)
}

// Barber cuts hair
type Barber interface {
	CutHair() string
}

// Singer sings
type Singer interface {
	Sing() string
}

// SingerBarber sings and cut hair
type SingerBarber interface {
	Barber
	Singer
}

//Person can sing and cut hair
type Person struct {
	Name string
}

// Sing ...
func (p Person) Sing() string {
	return "La la la!"
}

// CutHair ...
func (p Person) CutHair() string {
	return "Almost cutted"
}

// CutHairAndSing cut hair while sing
func CutHairAndSing(sb SingerBarber) {
	for i := 0; i < 3; i++ {
		fmt.Println(sb.Sing())
		fmt.Println(sb.CutHair())
	}
}

func main() {
	// standard output implements the method write, so can be used here
	PreffixOutput(os.Stdout, []byte(">>> "), []byte("ls -a\n"))
	// the same occurs with stdeer
	PreffixOutput(os.Stderr, []byte(">>> "), []byte("ls -a\n"))
	// you can use a file too
	f, err := os.Create("/tmp/dat2")
	if err != nil {
		log.Fatal(err)
	}
	PreffixOutput(f, []byte(">>> "), []byte("ls -a\n"))
	// or customized writers, it only have to implements writer interface
	PreffixOutput(&DatFile{}, []byte(">>> "), []byte("ls -a\n"))
	// person implements barber and singer interface then can be used as SingerBarber
	p := Person{}
	CutHairAndSing(p)
}
