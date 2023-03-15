package number

import (
	"testing"
)

func TestBasicUsage(t *testing.T) {
	want := Int64Ptr(77)

	if want == nil {
		t.Fatalf(`"pointer" Should not be nil`)
	}
}

func TestReturnValue(t *testing.T) {
	want := Int64Ptr(89)

	if *want != 89 {
		t.Fatalf(`"pointer" has the wrong value %d`, *want)
	}
}
