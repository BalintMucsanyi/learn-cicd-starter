package auth

import (
	"testing"
	"reflect"
)

func TestGetAPIKey(t *testing.T) {
	got := GetAPIKey ("Authorization XXX")
	want := []string{"Authorization", "XXX"}

	if !reflect.DeepEqual(want,got){
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}