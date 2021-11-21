package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t * testing.T){
	var myH myHandler
	h := NoSrv(&myH)

	switch v := h.(type){
	case http.Handler:
		//Do nothing
	default: 
		t.Error(fmt.Sprintf("Type is not http.handler, rather %T", v))
	}
}

func TestSessionLoad(t * testing.T){
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type){
	case http.Handler:
		//Do nothing
	default: 
		t.Error(fmt.Sprintf("Type is not http.handler, rather %T", v))
	}
}