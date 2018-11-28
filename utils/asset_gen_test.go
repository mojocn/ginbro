package utils

import "testing"

func TestEmbed(t *testing.T) {
	err := Embed("_test_out.go", "utils", ".")
	if err != nil {
		t.Fatal(err)
	}
}
