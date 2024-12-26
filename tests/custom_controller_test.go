package tests

import (
	"encoding/json"
	"testing"
	"net/http"
	"net/http/httptest"
	// "io/ioutil"
	"strings" // Required for checking the URL prefix
	"github.com/stretchr/testify/assert"
	beego "github.com/beego/beego/v2/server/web"
	 "github.com/beego/beego/v2/server/web/context"
	"cat-api/controllers" // Adjust import path as necessary
)


type Breed struct {
    ID               string `json:"id"`
    Name             string `json:"name"`
    Description      string `json:"description"`
    Origin           string `json:"origin"`
    ReferenceImageID string `json:"reference_image_id"`
}

type BreedImage struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Breeds []interface{} `json:"breeds"`
}
// TestGet tests the Get method of the CustomController
func TestGet(t *testing.T) {
	// Step 1: Create a mock server that simulates the external API (TheCatAPI)
	mockResponse := `[{"id":"dqg", "url":"https://cdn2.thecatapi.com/images/dqg.jpg"}]`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method) // Ensure it's a GET request
		assert.Equal(t, "x-api-key", r.Header.Get("x-api-key")) // Check for the correct API key header
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse)) // Send the mock response
	}))
	defer mockServer.Close()

	// Step 2: Set the Beego configuration directly (bypassing app.conf file)
	beego.AppConfig.Set("catapi_key", "test_api_key")
	beego.AppConfig.Set("catapi_url", mockServer.URL)

	// Step 3: Create an instance of the CustomController and perform the GET action
	ctrl := &controllers.CustomController{}
	ctrl.Data = make(map[interface{}]interface{}) // Initialize Data map
	ctrl.TplName = ""

	// Perform the GET action (simulate the controller action)
	ctrl.Get()

	// Step 4: Validate the results
	// Check that the URL starts with the expected domain
	assert.True(t, strings.HasPrefix(ctrl.Data["CatImageURL"].(string), "https://cdn2.thecatapi.com/images/"))
	assert.Equal(t, "custom_page.tpl", ctrl.TplName)
}


func TestGetBreeds(t *testing.T) {
    mockResponse := `[
        {
            "id": "abys",
            "name": "Abyssinian",
            "description": "The Abyssinian is easy to care for, and a joy to have in your home. They’re affectionate cats and love both people and other animals.",
            "origin": "Egypt",
            "reference_image_id": "0XYvRd7oD"
        },
        {
            "id": "aege",
            "name": "Aegean",
            "description": "Native to the Greek islands known as the Cyclades in the Aegean Sea, these are natural cats, meaning they developed without humans getting involved in their breeding. As a breed, Aegean Cats are rare, although they are numerous on their home islands. They are generally friendly toward people and can be excellent cats for families with children.",
            "origin": "Greece",
            "reference_image_id": "ozEvzdVM-"
        }
    ]`

    var expectedBreeds []Breed
    err := json.Unmarshal([]byte(mockResponse), &expectedBreeds)
    if err != nil {
        t.Fatalf("Error unmarshalling mock response: %v", err)
    }

    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "GET", r.Method)
        assert.Equal(t, "test_api_key", r.Header.Get("x-api-key"))
        assert.Equal(t, "/v1/breeds", r.URL.Path)

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(mockResponse))
    }))
    defer mockServer.Close()

    beego.AppConfig.Set("catapi_key", "test_api_key")
    beego.AppConfig.Set("catapi_url", mockServer.URL)

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/v1/breeds", nil)
    ctx := context.NewContext()
    ctx.Reset(w, r)

    ctrl := &controllers.CustomController{}
    ctrl.Init(ctx, "CustomController", "GetBreeds", nil)

    ctrl.GetBreeds()

    assert.Equal(t, http.StatusOK, w.Code)

    var actualBreeds []Breed
    err = json.Unmarshal(w.Body.Bytes(), &actualBreeds)
    if err != nil {
        t.Fatalf("Error unmarshalling actual response: %v", err)
    }

    // Limit actual breeds to expected values for comparison
    actualBreeds = actualBreeds[:len(expectedBreeds)]

    assert.Equal(t, expectedBreeds, actualBreeds)
}




func TestGetBreedImages(t *testing.T) {
	// Correct Mock Response with only 2 images and the correct structure
	mockResponse := `[
		{
			"id": "image1",
			"url": "https://example.com/image1.jpg",
			"breeds": null
		},
		{
			"id": "image2",
			"url": "https://example.com/image2.jpg",
			"breeds": null
		}
	]`

	// Set up mock server -  Crucially, use mockResponse in the handler
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Received request: %s %s", r.Method, r.URL.Path)
		t.Logf("Request headers: %+v", r.Header)

		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "test_api_key", r.Header.Get("x-api-key"))
		breedID := r.URL.Query().Get("breed_id")
		assert.Equal(t, "abys", breedID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse)) // Use the mockResponse here
	}))
	defer mockServer.Close()

	beego.AppConfig.Set("catapi_key", "test_api_key")
	beego.AppConfig.Set("catapi_url", mockServer.URL)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/images/search?breed_id=abys", nil)
	r.Header.Set("x-api-key", "test_api_key")
	ctx := context.NewContext()
	ctx.Reset(w, r)

	ctrl := &controllers.CustomController{}
	ctrl.Init(ctx, "CustomController", "GetBreedImages", nil)

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok && err.Error() != "user stop run" {
				t.Fatalf("Unexpected panic: %v", r)
			}
		}
	}()
	ctrl.GetBreedImages()

	assert.Equal(t, http.StatusOK, w.Code)
	t.Logf("Raw response body: %s", w.Body.String())

	var images []BreedImage
	err := json.Unmarshal(w.Body.Bytes(), &images)
	if err != nil {
		t.Fatalf("Error unmarshalling response: %v", err)
	}


}




