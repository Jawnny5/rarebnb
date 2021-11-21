package main

import (
	"fmt"
	"rarebnb/internal/config"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type){
	case *chi.Mux:
		//Do nothing, test passed
	default: 
		t.Error(fmt.Sprintf("Type is not chi.mux, rather %T", v))
	}
}