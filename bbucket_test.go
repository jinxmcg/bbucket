// This file tests the buckets

package bbucket

import (
//	"math"
//	"math/rand"
	"testing"
//	"time"
)

func TestEmptyBitSet(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("A zero-length bitset should be fine")
		}
	}()
	b := New()
	if b.Len() != 0 {
		t.Errorf("Empty set should have length 0, not %d", b.Len())
	}
}

