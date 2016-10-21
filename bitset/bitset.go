package bitset

const (
	WS = 64 // word size
)

type BitSet struct {
	set []uint64
}

func (b *BitSet) Add(i int) *BitSet {
	r, c := i/WS, uint(i%WS)
	if r >= len(b.set) {
		for j := 0; j <= r-len(b.set); j++ {
			b.set = append(b.set, 0)
		}
	}
	b.set[r] = b.set[r] | 1<<c
	return b
}

func (b *BitSet) Has(i int) bool {
	r, c := i/WS, uint(i%WS)
	if r < len(b.set) {
		return b.set[r]&1<<c == 1<<c
	} else {
		return false
	}

}

/*
Add(int) Self
Remove(int) Self
Add(BitSet) Self
Remove(BitSet) Self
Size() int
IsSubsetOf(BitSet)
*/
