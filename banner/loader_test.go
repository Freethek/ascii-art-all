package banner

import (
	"testing"
)

func TestLoad(t *testing.T) {
	// loading banner map
	bannerMap := Load("../banners/standard.txt")

	// checking that all 95 ascii are present
	if len(bannerMap) != 95 {
		t.Errorf("incomplete graphic respresntation for all ascii")
	}

	// checking if the 8line are complete and accurate
	// by cunting the slice of string index
	if len(bannerMap['A']) != 8 {
		t.Errorf("invalid graphic representation")
	}

}
