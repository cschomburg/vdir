// Copyright (C) 2015 Constantin Schomburg <me@cschomburg.com>

package vdir

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	r := NewDecoder(strings.NewReader(TestVCard))
	_, err := r.ReadObject()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnmarshal(t *testing.T) {
	var c Card
	if err := Unmarshal([]byte(TestVCard), &c); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", c)
}
