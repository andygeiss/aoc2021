package assert

import (
	"fmt"
	"testing"
)

// That ...
func That(t *testing.T, value, expected interface{}) {
	if fmt.Sprintf("%v", value) != fmt.Sprintf("%v", expected) {
		t.Fatalf("got value [%v], but it's not like expected [%v]", value, expected)
	}
}
