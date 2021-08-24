package chain

import "testing"

func TestLinkChain(t *testing.T) {
	words := "test"
	f := &AdSensitiveWordFilter{}
	f.SetNext(&PoliticalWordFilter{})
	f.Filter(words)
}

func TestArrayChain(t *testing.T) {
	fChain := &SensitiveWordFilterChain{}
	fChain.AddFilter(&AryAdSensitiveWordFilter{})
	fChain.AddFilter(&AryPoliticalWordFilter{})

	fChain.Filter("test")
}
