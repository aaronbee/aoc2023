package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/aaronbee/aoc2023"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	var part1 int
	hm := hashmap{}
	s := bufio.NewScanner(f)
	s.Split(scanSeq)
	for s.Scan() {
		part1 += hash(s.Bytes())
		doStep(&hm, s.Text())
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", hm.power())
}

func scanSeq(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		// We have a full instruction
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), bytes.TrimSpace(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func doStep(hm *hashmap, step string) {
	if label, ok := strings.CutSuffix(step, "-"); ok {
		hm.delete(label)
		return
	}
	label, val, ok := strings.Cut(step, "=")
	if !ok {
		panic(fmt.Errorf("unexpected step: %q", step))
	}
	hm.set(label, aoc2023.Atoi(val))
}

func hash(seq []byte) int {
	var h int
	for _, b := range seq {
		h += int(b)
		h *= 17
		h %= 256
	}
	return h
}

type hashmap struct {
	boxes [256]*slot
}

type slot struct {
	k    string
	v    int
	next *slot
}

func (hm *hashmap) set(label string, v int) {
	h := hash([]byte(label))
	sPtr := &hm.boxes[h]
	for *sPtr != nil {
		if (*sPtr).k == label {
			(*sPtr).v = v
			return
		}
		sPtr = &(*sPtr).next
	}
	*sPtr = &slot{k: label, v: v}
}

func (hm *hashmap) delete(label string) {
	h := hash([]byte(label))
	sPtr := &hm.boxes[h]
	for *sPtr != nil {
		if (*sPtr).k == label {
			*sPtr = (*sPtr).next
			return
		}
		sPtr = &(*sPtr).next
	}
}

func (hm *hashmap) power() int {
	var power int
	for i, s := range hm.boxes {
		j := 1
		for s != nil {
			power += (i + 1) * j * s.v
			s = s.next
			j++
		}
	}
	return power
}
