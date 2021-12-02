package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"rarebnb/internal/models"
	"strings"
	"testing"
)

type postData struct {
	key string
	value string 
}

var theTests = []struct {
	name string
	url string
	method string
	expectedStatusCode int 
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"hurtyurt", "/hurtyurt", "GET", http.StatusOK},
	{"timothys", "/timothys", "GET", http.StatusOK},
	{"searchavailability", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	// {"post-search-avail", "/search-availability", "POST", []postData{
	// 	{key: "start", value: "01-01-2020"},
	// 	{key: "end", value: "01-03-2020"},
	// }, http.StatusOK},
	// {"post-search-avail-json", "/search-availability-json", "POST", []postData{
	// 	{key: "start", value: "01-01-2020"},
	// 	{key: "end", value: "01-03-2020"},
	// }, http.StatusOK},
	// {"make-reservation-post", "/make-reservation", "POST", []postData{
	// 	{key: "first_name", value: "Sir"},
	// 	{key: "last_name", value: "SAL"},
	// 	{key: "email", value: "ssmal@niceguy.com"},
	// 	{key: "phone", value: "1-800-NICE-GUY"},
	// }, http.StatusOK},
}

func TestHandlers(t *testing.T){
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID: 1,
			RoomName: "Hurt Yurt",
		},
	}
	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test case where reservation is not in session
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test with nonexistent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

}

func TestRepository_PostReservation(t *testing.T){
	reqBody := "start_date=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Sir")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Sal")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=sir@sal.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=111222333")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test for missing post body
	req, _ = http.NewRequest("POST", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for missing post body. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test for invalid start date
	reqBody = "start_date=invalid"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Sir")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Sal")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=sir@sal.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=111222333")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid start date. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test for invalid end date
	reqBody = "start_date=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=invalid")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Sir")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Sal")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=sir@sal.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=111222333")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid end date. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test for invalid room id
	reqBody = "start_date=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Sir")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Sal")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=sir@sal.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=111222333")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=invalid")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid room id. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test to see if post fails forms validation
	reqBody = "start_date=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=S")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Sal")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=sir@sal.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=111222333")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code for invalid data. Got %d, wanted %d", rr.Code, http.StatusOK)
	}

	//Test to see if insert into database failed
	reqBody = "start_date=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Sir")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Sal")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=sir@sal.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=111222333")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=2")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned failed to insert. Got %d status code, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	//Test to see if creating room restriction failed
	reqBody = "start_date=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-03")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=Sir")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Sal")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=sir@sal.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=111222333")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1000")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded ")

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned failed to create room restriction. Got %d status code, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

}

func getCtx(req *http.Request) context.Context{
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}