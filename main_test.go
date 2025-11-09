package main

import "testing"
//testing high severity
func TestForceFailure(t *testing.T) {
    t.Fatal("panic: Forced CI failure to test high severity")
}
