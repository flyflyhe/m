package str

import (
	"fmt"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	fmt.Println(LicenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(LicenseKeyFormatting("2-5g-3-J", 2))
}

func TestCountSegments(t *testing.T) {
	fmt.Println(CountSegments(""))
	fmt.Println(CountSegments("    "))
	fmt.Println(CountSegments(" a   "))
	fmt.Println(CountSegments("a   b"))
}

func TestFindRepeatedDnaSequences(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(FindRepeatedDnaSequences(s))
}