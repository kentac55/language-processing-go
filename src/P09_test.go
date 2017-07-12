package exercise

import (
	"testing"
)

func TestP09(t *testing.T) {
	const i = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	const o = "I c'nlduot beielve that I cuold aualctly unaserntdd what I was ridneag : the phmnenaoel peowr of the huamn mind ."
	const seed = 1
	if x := P09(i, seed); x != o {
		t.Errorf("P09() may be broken... \ni: %v\ne: %v\no: %v", i, o, x)
	}
}
