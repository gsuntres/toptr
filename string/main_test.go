package string

import (
	"log"
	"testing"
)

func TestBasicUsage(t *testing.T) {
	want := StringPtr("hi")

	if want == nil {
		t.Fatalf(`"pointer" Should not be nil`)
	}
}

func TestReturnValue(t *testing.T) {
	want := StringPtr("Hello!")
	log.Println(want)
	if *want != "Hello!" {
		t.Fatalf(`"pointer" has the wrong value %s`, *want)
	}
}
