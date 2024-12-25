package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	// Dynamically resolve the absolute path to the application
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	web.TestBeegoInit(apppath)
}

// TestGetCatImage tests the / endpoint
func TestGetCatImage(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestGetCatImage", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Get Cat Image\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestGetBreeds tests the /custom/breeds endpoint
func TestGetBreeds(t *testing.T) {
	r, _ := http.NewRequest("GET", "/custom/breeds", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestGetBreeds", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Get Breeds\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Breeds Should Not Be Empty", func() {
			var breeds []map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &breeds)
			So(err, ShouldBeNil)
			So(breeds, ShouldNotBeEmpty)
		})
	})
}

// TestGetBreedImages tests the /custom/breed_images endpoint
func TestGetBreedImages(t *testing.T) {
	r, _ := http.NewRequest("GET", "/custom/breed_images?breed_id=beng", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestGetBreedImages", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Get Breed Images\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Breed Images Should Not Be Empty", func() {
			var images []map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &images)
			So(err, ShouldBeNil)
			So(images, ShouldNotBeEmpty)
		})
	})
}

// TestCreateVote tests the /custom/vote endpoint
func TestCreateVote(t *testing.T) {
	vote := map[string]interface{}{
		"image_id": "abc123",
		"value":    1,
	}
	voteData, _ := json.Marshal(vote)

	r, _ := http.NewRequest("POST", "/custom/vote", bytes.NewBuffer(voteData))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestCreateVote", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Create Vote\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Response Should Not Be Empty", func() {
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)
			So(response, ShouldNotBeEmpty)
		})
	})
}

// TestGetVotes tests the /custom/votes endpoint
func TestGetVotes(t *testing.T) {
	r, _ := http.NewRequest("GET", "/custom/votes?limit=5&order=ASC", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestGetVotes", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Get Votes\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Votes Should Not Be Empty", func() {
			var votes []map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &votes)
			So(err, ShouldBeNil)
			So(votes, ShouldNotBeEmpty)
		})
	})
}

// TestCreateFavourite tests the /custom/favourite endpoint
func TestCreateFavourite(t *testing.T) {
	fav := map[string]interface{}{
		"image_id": "abc123",
		"sub_id":   "user123",
	}
	favData, _ := json.Marshal(fav)

	r, _ := http.NewRequest("POST", "/custom/favourite", bytes.NewBuffer(favData))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestCreateFavourite", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Create Favourite\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Response Should Not Be Empty", func() {
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)
			So(response, ShouldNotBeEmpty)
		})
	})
}

// TestGetFavourites tests the /custom/favourites endpoint
func TestGetFavourites(t *testing.T) {
	r, _ := http.NewRequest("GET", "/custom/favourites", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestGetFavourites", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Get Favourites\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Favourites Should Not Be Empty", func() {
			var favourites []map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &favourites)
			So(err, ShouldBeNil)
			So(favourites, ShouldNotBeEmpty)
		})
	})
}

// TestDeleteFavourite tests the /custom/favourites/1 endpoint
func TestDeleteFavourite(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/custom/favourites/1", nil)
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestDeleteFavourite", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Delete Favourite\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Response Should Confirm Deletion", func() {
			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			So(err, ShouldBeNil)
			So(response["message"], ShouldEqual, "Favourite deleted successfully")
		})
	})
}
