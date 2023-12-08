package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/aaronbee/aoc2023"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var plays []play
	for s.Scan() {
		l := s.Text()
		hand, bid, ok := strings.Cut(l, " ")
		if !ok {
			panic(fmt.Errorf("unexpected line: %q", l))
		}
		plays = append(plays, play{
			h:   parseHand(hand),
			bid: aoc2023.Atoi(bid),
		})
	}

	slices.SortFunc(plays, func(a, b play) int {
		if a.h.t != b.h.t {
			return cmp.Compare(a.h.t, b.h.t)
		}
		return slices.Compare(a.h.cards[:], b.h.cards[:])
	})
	var rank int
	var part1 int
	for _, p := range plays {
		rank++
		part1 += p.bid * rank
	}
	fmt.Println("Part 1:", part1)
}

type play struct {
	h   hand
	bid int
}

func (p play) String() string {
	return fmt.Sprintf("%s bid=%d", p.h, p.bid)
}

type handType byte

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func (t handType) String() string {
	switch t {
	case highCard:
		return "high card"
	case onePair:
		return "one pair"
	case twoPair:
		return "two pair"
	case threeOfAKind:
		return "three of a kind"
	case fullHouse:
		return "full house"
	case fourOfAKind:
		return "four of a kind"
	case fiveOfAKind:
		return "five of a kind"
	}
	return fmt.Sprintf("<unexpected hand type: %d>", t)
}

type card byte

func (c card) String() string {
	switch c {
	case 2, 3, 4, 5, 6, 7, 8, 9:
		return string(rune(c + '0'))
	case 10:
		return "T"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	}
	return fmt.Sprintf("<unexpected card: %d>", c)
}

type hand struct {
	cards [5]card
	t     handType
}

func (h hand) String() string {
	return fmt.Sprintf("%v %v", h.cards, h.t)
}

func parseHand(s string) hand {
	var cards [5]card
	var counts [15]card
	for i := 0; i < len(s); i++ {
		var c card
		switch b := s[i]; b {
		case '2', '3', '4', '5', '6', '7', '8', '9':
			c = card(b - '0')
		case 'T':
			c = 10
		case 'J':
			c = 11
		case 'Q':
			c = 12
		case 'K':
			c = 13
		case 'A':
			c = 14
		default:
			panic(fmt.Errorf("unexpected card %s from hand[%d]: %s",
				string(rune(c)), i, s))
		}
		cards[i] = c
		counts[c]++
	}

	var t handType
	slices.Sort(counts[:])
	last := len(counts) - 1
	switch counts[last] {
	case 5:
		t = fiveOfAKind
	case 4:
		t = fourOfAKind
	case 3:
		if counts[last-1] == 2 {
			t = fullHouse
		} else {
			t = threeOfAKind
		}
	case 2:
		if counts[last-1] == 2 {
			t = twoPair
		} else {
			t = onePair
		}
	}

	return hand{cards: cards, t: t}
}
