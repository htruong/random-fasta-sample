package main

import (
	"code.google.com/p/biogo/alphabet"
	"code.google.com/p/biogo/io/seqio/fasta"
	"code.google.com/p/biogo/seq/linear"
	"flag"
	"fmt"
	"math/rand"
	"os"
	//"strconv"
	//"strings"
	"time"
	//"unsafe"
)

var numkeep *int = flag.Int("k", 100, "Number of sequences in the sample")
var total *int = flag.Int("t", 100, "Number of total sequences in the fasta file")
var in *string = flag.String("in", "in.fas", "Input file")
var out *string = flag.String("out", "out.fas", "Output file")

func main() {

	flag.Parse()

	// Now really print out the connected components
	seqtype := alphabet.DNA

	f1, err := os.Open(*in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.\n", err)
		os.Exit(1)
	}
	defer f1.Close()

	f2, err := os.Create(*out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.\n", err)
		os.Exit(1)
	}
	defer f2.Close()

	fastaSeqs := fasta.NewReader(f1, linear.NewSeq("", nil, seqtype))
	fastaOut := fasta.NewWriter(f2, 80)

	rand.Seed(time.Now().UnixNano())
	list := rand.Perm(*total)
	keep := make(map[int]bool)

	for i := 0; i < *numkeep; i++ {
		keep[list[i]] = true
	}

	counter := 0

	for true {
		seq, e := fastaSeqs.Read()
		if e != nil {
			break
		}

		if keep[counter] {
			fastaOut.Write(seq)
		}

		counter++
	}

	fmt.Fprintf(os.Stderr, "Done.\n")

}
