package bitset

import (
	"bytes"
	"fmt"
)

// blockSize defines the max blocksize size
const blockSize = 64

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set
type IntSet struct {
	words []uint
}

// Creates new bitset using the size argument
func New(l int) *IntSet {
	// So that it'll be possible to union logic of the map implemetation New uses the variadic parameter size and takes the 0th element
	set := &IntSet{words: make([]uint, l)}
	return set
}

// String returns a string representation of bitset
func (s IntSet) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')

	for i := 0; i < len(s.words); i++ {
		if s.words[i] == 0 {
			continue
		}
		for j := 0; j < blockSize; j++ {
			if (s.words[i]>>j)&1 != 0 {
				buf.WriteByte(' ')
				fmt.Fprintf(&buf, "%d", blockSize*i+j)
			}
		}
	}

	buf.WriteByte(' ')
	buf.WriteByte('}')

	return buf.String()
}

// Elements return all the elements represented as slice of all the blocks
func (s *IntSet) Elements() []int {
	var elements []int
	var addedElementsCount int

	for i := 0; i < len(s.words); i++ {
		if s.words[i] == 0 {
			continue
		}
		for j := 0; j < blockSize; j++ {
			if (s.words[i]>>j)&1 != 0 {
				elements = append(elements, blockSize*i+j)
				addedElementsCount++
			}
		}
	}

	return elements[:addedElementsCount]
}

// Add adds all the passed elements to the set
func (s *IntSet) Add(values ...int) {
	for _, value := range values {
		s.subAdd(value)
	}
}

// Remove removes all the passed elements from the set
func (s *IntSet) Remove(values ...int) {
	for _, value := range values {
		s.subRemove(value)
	}
}

// Has reports whether the set contains non-negative value x
func (s *IntSet) Has(x int) bool {
	// word - the bit block index
	// bit - the index of element in the its block
	word, bit := x/blockSize, uint(x%blockSize)
	return word < len(s.words) && // Whether the block for this number exist
		s.words[word]&(1<<bit) != 0 // Has the corresponding bit is 1
}

// Len returns a count of non-zero bits
func (s *IntSet) Len() int {
	var length int
	for _, block := range s.words {
		for block > 0 {
			length++
			block &= (block - 1)
		}
	}
	return length
}

// Clear does a bitset empty
func (s *IntSet) Clear() error {
	if s != nil {
		s.words = nil
		return nil
	}
	return fmt.Errorf("the bitset ptr is nil")
}

// Copy makes a copy of the initial bitset ane returns it
func (s *IntSet) Copy() *IntSet {
	newWords := make([]uint, len(s.words))
	copy(newWords, s.words)

	newBitSet := &IntSet{words: newWords}

	return newBitSet
}

// UnionWith(t *IntSet) sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		// Union blocks if the block at index i exists in the s.words
		if i < len(s.words) {
			s.words[i] |= tword
			// Otherwise just append the block at index i from t.words
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	var intersectedBlocksCount int
	maxBlockLevel := minLen(s, t)

	for i := 0; i < maxBlockLevel; i++ {
		s.words[i] &= t.words[i]
		intersectedBlocksCount++
	}

	s.words = s.words[:intersectedBlocksCount]
}

// DifferenceWith sets s to the difference of s and t (excludes all the elements of t from s)
func (s *IntSet) DifferenceWith(t *IntSet) {
	maxBlockLevel := minLen(s, t)

	for i := 0; i < maxBlockLevel; i++ {
		s.words[i] &^= t.words[i]
	}
}

// SymmetricDifference returns the symmetric difference of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	sWithoutT := (s.Copy())
	sWithoutT.DifferenceWith(t)

	tWithoutS := (t.Copy())
	tWithoutS.DifferenceWith(s)

	sWithoutT.UnionWith(tWithoutS)

	return &IntSet{words: sWithoutT.words}
}

// subAdd adds the non-negative value x to the set
func (s *IntSet) subAdd(x int) {
	word, bit := x/blockSize, uint(x%blockSize)
	// While the count of blocks is less than word, add new blocks.
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// Set the bit at the number's index in the block
	s.words[word] |= 1 << bit
}

// subRemove turns an element at words[blockIndex][indexInBlock] into 0
func (s *IntSet) subRemove(x int) bool {
	blockIndex, indexInBlock := s.getBlockAndIndex(x)
	if blockIndex >= 0 {
		s.words[blockIndex] &^= (1 << indexInBlock)
		return true
	}

	return false
}

// getBlockAndIndex returns the block index and the index in this block if the element is consisted by set
// otherwise it returns -1,-1
func (s *IntSet) getBlockAndIndex(x int) (int, int) {
	if s.Has(x) {
		blockIndex, indexInBlock := x/blockSize, x%blockSize
		return blockIndex, indexInBlock
	}
	return -1, -1
}

// minLen int returns the minimal length of underlying slices of s and t
func minLen(s, t *IntSet) int {
	if len(s.words) > len(t.words) {
		return len(t.words)
	}
	return len(s.words)
}
