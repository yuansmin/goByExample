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
}

type IntSet struct {
	word []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return len(s.word) > word && s.word[word]&1<<bit != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.word) {
		s.word = append(s.word, 0)
	}
	s.word[word] |= 1 << bit
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

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				num := i*64 + j
				fmt.Fprintf(&buf, "%d", num)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
