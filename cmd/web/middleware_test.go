package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
		//Do nothing since this is what we expect.
		break
	default:
		t.Errorf("type is not http.Handler it is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
	case http.Handler:
		//Do nothing since this is what we expect.
		break
	default:
		t.Errorf("type is not http.Handler it is %T", v)
	}
}
