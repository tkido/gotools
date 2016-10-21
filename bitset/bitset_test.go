package bitset

import "testing"

func TestHasBitSet(t *testing.T) {
	var b BitSet
	if b.Has(1) {
		t.Errorf("Empty set should not have 1")
	}
}
