package alphabet

import (
	"log"
)

// LetterDistribution encodes the tile distribution for the relevant game.
type LetterDistribution struct {
	Distribution map[rune]uint8
	PointValues  map[rune]uint8
	SortOrder    map[rune]int
	numLetters   int
}

func EnglishLetterDistribution() LetterDistribution {
	dist := map[rune]uint8{
		'A': 9, 'B': 2, 'C': 2, 'D': 4, 'E': 12, 'F': 2, 'G': 3, 'H': 2,
		'I': 9, 'J': 1, 'K': 1, 'L': 4, 'M': 2, 'N': 6, 'O': 8, 'P': 2,
		'Q': 1, 'R': 6, 'S': 4, 'T': 6, 'U': 4, 'V': 2, 'W': 2, 'X': 1,
		'Y': 2, 'Z': 1, '?': 2,
	}
	ptValues := map[rune]uint8{
		'A': 1, 'B': 3, 'C': 3, 'D': 2, 'E': 1, 'F': 4, 'G': 2, 'H': 4,
		'I': 1, 'J': 8, 'K': 5, 'L': 1, 'M': 3, 'N': 1, 'O': 1, 'P': 3,
		'Q': 10, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 4, 'W': 4, 'X': 8,
		'Y': 4, 'Z': 10, '?': 0,
	}
	return LetterDistribution{dist, ptValues,
		makeSortMap("ABCDEFGHIJKLMNOPQRSTUVWXYZ?"), 100}
}

func SpanishLetterDistribution() LetterDistribution {
	dist := map[rune]uint8{
		'1': 1, '2': 1, '3': 1, // 1: CH, 2: LL, 3: RR
		'A': 12, 'B': 2, 'C': 4, 'D': 5, 'E': 12, 'F': 1, 'G': 2, 'H': 2,
		'I': 6, 'J': 1, 'L': 4, 'M': 2, 'N': 5, 'Ñ': 1, 'O': 9, 'P': 2,
		'Q': 1, 'R': 5, 'S': 6, 'T': 4, 'U': 5, 'V': 1, 'X': 1, 'Y': 1,
		'Z': 1, '?': 2,
	}
	ptValues := map[rune]uint8{
		'1': 5, '2': 8, '3': 8, // 1: CH, 2: LL, 3: RR
		'A': 1, 'B': 3, 'C': 3, 'D': 2, 'E': 1, 'F': 4, 'G': 2, 'H': 4,
		'I': 1, 'J': 8, 'L': 1, 'M': 3, 'N': 1, 'Ñ': 8, 'O': 1, 'P': 3,
		'Q': 5, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 4, 'X': 8, 'Y': 4,
		'Z': 10, '?': 0,
	}
	return LetterDistribution{dist, ptValues,
		makeSortMap("ABC1DEFGHIJL2MNÑOPQR3STUVXYZ?"), 100}
}

// MakeBag returns a bag of tiles.
func (ld LetterDistribution) MakeBag(alph *Alphabet) *Bag {
	bag := make([]MachineLetter, ld.numLetters)
	idx := 0
	for rn, val := range ld.Distribution {
		for j := uint8(0); j < val; j++ {
			val, err := alph.Val(rn)
			if err != nil {
				log.Fatalf("Attempt to initialize bag failed: %v", err)
			}
			bag[idx] = val
			idx++
		}
	}
	// Make an array of scores in alphabet order
	scores := make([]int, len(ld.Distribution))
	for rn, ptVal := range ld.PointValues {
		ml, err := alph.Val(rn)
		if err != nil {
			panic("Wrongly initialized")
		}
		if ml == BlankMachineLetter {
			ml = MachineLetter(len(ld.Distribution) - 1)
		}
		scores[ml] = int(ptVal)
	}
	b := NewBag(bag, len(ld.Distribution), alph, scores)
	b.Shuffle()

	return b
}

func makeSortMap(order string) map[rune]int {
	sortMap := make(map[rune]int)
	for idx, letter := range order {
		sortMap[letter] = idx
	}
	return sortMap
}
