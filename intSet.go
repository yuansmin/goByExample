// bit 数组
// Add(n)：将数组第 n 位置为 1
// Has(n): 位数组第 n 位是否为 1
// IntSet 使用一个 []uint64 ，每个元素有 64 位，按照其二进制的表现形式，如果第 n 位为 1，
// 则表示该数组包含 n。第一个元素表示 0～63，第二个元素 64～127，以此类推

package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := IntSet{}
	s.Add(1)
	s.Add(20)
	s.Add(300)
	fmt.Println(s.String())
	fmt.Println(s.Has(20))
	fmt.Printf("length s: %d\n", s.Len())
	s.Remove(20)
	fmt.Printf("remove 20: %s\n", &s)
	s.Add(20)
	fmt.Printf("copy set: %s\n", s.String())
	s.Clear()
	s.AddAll(2, 3, 4)
	fmt.Printf("AddAll %s\n", s.String())
}

const wordLenght = 64

type wordType uint64

type IntSet struct {
	word []wordType
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/wordLenght, uint(x%wordLenght)
	return len(s.word) > word && s.word[word]&1<<bit != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/wordLenght, uint(x%wordLenght)
	for word >= len(s.word) {
		s.word = append(s.word, 0)
	}
	s.word[word] |= 1 << bit
}

func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.word {
		if i >= len(s.word) {
			s.word = append(s.word, word)
		} else {
			s.word[i] |= word
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.word {
		if word == 0 {
			continue
		}

		for j := 0; j < wordLenght; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				num := i*wordLenght + j
				fmt.Fprintf(&buf, "%d", num)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	length := 0
	for w := 0; w < len(s.word); w++ {
		for bit := 0; bit < wordLenght; bit++ {
			if s.word[w]&(1<<uint(bit)) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Remove(x int) {
	word, bit := x/wordLenght, uint(x%wordLenght)
	if word >= len(s.word) {
		return
	}
	s.word[word] = s.word[word] &^ (1 << bit)
}

func (s *IntSet) Clear() {
	for i := 0; i < len(s.word); i++ {
		s.word[i] = wordType(0)
	}
}

func (s *IntSet) Copy() *IntSet {
	newIntSet := IntSet{}
	copy(newIntSet.word, s.word)
	return &newIntSet
}
