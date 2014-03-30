package fasta

import (
	"bufio"
	"os"
	"strings"
)

type Dna struct {
	Sequence []byte
}

type FastaFile struct {
	DnaSeqs map[string]Dna
}

func NewFastaFile(filename string) *FastaFile {
	file, err := os.Open(filename)

	if err != nil {
		panic("Not a valid file.")
	}

	var (
		text       string
		dna        []byte
		identifier string
		dnaMap     = make(map[string]Dna)
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text = strings.TrimSpace(scanner.Text())

		// Hit an identifier line
		if text[0] == '>' {
			// If we stored a previous identifier, get the DNA string and map to the
			// identifier and clear the string
			if identifier != "" {
				dnaMap[identifier] = Dna{dna}
				dna = make([]byte, 0)
				identifier = ""
			}

			// Standard FASTA identifiers look like: ">foo_<id>"
			identifier = strings.Split(text, ">")[1]
		} else {
			// Append here since multi-line DNA strings are possible
			dna = append(dna, []byte(text)...)
		}
	}

	// EOF, there's one last identifier to store
	dnaMap[identifier] = Dna{dna}

	f := FastaFile{DnaSeqs: dnaMap}
	return &f
}
