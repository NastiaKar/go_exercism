package nucleotidecount

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
    var histogram Histogram = Histogram{'A' : 0, 'C' : 0, 'G' : 0, 'T' : 0}
	for _, i := range d {
        _, ok := histogram[i]
        if !ok {
            return nil, errors.New("Incorrect nucleotide")
        } else {
            histogram[i] += 1
        }
    }

    return histogram, nil
}
