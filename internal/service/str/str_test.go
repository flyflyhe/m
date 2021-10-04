package str

import (
	"fmt"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	fmt.Println(LicenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(LicenseKeyFormatting("2-5g-3-J", 2))
}