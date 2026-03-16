package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("Sequences have different lengths")
    }

    distance := 0
    for x, _ := range a {
        if a[x] != b[x] {
            distance += 1
        }
    }

    return distance, nil
}
