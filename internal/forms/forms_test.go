package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("form is invalid")
	}
}

func TestForm_Required(t * testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid(){
		t.Error("Form found valid when required values are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid(){
		t.Error("form found invalid with all required fields")
	}
}

func TestForm_Has(t *testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("Form validates incomplete field")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form fails to validate complete field")
	}
}

func TestForm_MinLength(t *testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("validates nonexistent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should throw an error, but didn't")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid(){
		t.Error("validates field length of 100 when field is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)

	form.MinLength("another_field", 1)
	if !form.Valid(){
		t.Error("fails to validate field length of 1 when length is satisfied")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("threw an error, but shouldn't")
	}
}

func TestForm_IsEmail(t *testing.T){
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid(){
		t.Error("form validates email for nonexistent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "ssmal@niceguy.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid(){
		t.Error("form fails to validate properly formatted email")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "ssmal")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid(){
		t.Error("form validates improperly formatted email")
	}
}