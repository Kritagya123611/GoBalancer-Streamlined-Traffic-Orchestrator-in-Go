package main

import "testing"

func TestForceFailure(t *testing.T) {
    t.Fatal("Forced failure to test CI logs categorization")
}
