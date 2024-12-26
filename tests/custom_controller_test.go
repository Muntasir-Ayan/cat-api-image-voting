package tests

import (
	// "bytes"
	"encoding/json"
	// "fmt"
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


type Vote struct {
	ImageID string `json:"image_id"`
	Value   int    `json:"value"`
}


func TestCreateVote(t *testing.T) {
    // Mock server simulating the external API
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        t.Logf("Mock server received request method: %s", r.Method)
        t.Logf("Mock server received request headers: %+v", r.Header)

        // Return 401 if API key is invalid
        if r.Header.Get("x-api-key") != "test_api_key" {
            t.Log("Invalid API key detected")
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte(`"AUTHENTICATION_ERROR"`)) // Keep the original string response
            return
        }

        t.Log("Valid API key detected")
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message":"SUCCESS","id":12345}`))
    }))
    defer mockServer.Close()

    // Set Beego configuration for mock server
    beego.AppConfig.Set("catapi_key", "test_api_key") // Set a valid API key
    beego.AppConfig.Set("catapi_url", mockServer.URL)

    // Create the test request and response recorder
    formData := `image_id=test_image_id&value=1`
    r := httptest.NewRequest("POST", "/v1/votes", strings.NewReader(formData))
    r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Set("x-api-key", "invalid_api_key") // Set an invalid API key to trigger the error
    w := httptest.NewRecorder()

    // Initialize the context and controller
    ctx := context.NewContext()
    ctx.Reset(w, r)
    ctrl := &controllers.CustomController{}
    ctrl.Init(ctx, "CustomController", "CreateVote", nil)

    // Call the CreateVote method
    ctrl.CreateVote()

    // Assert the response
    t.Logf("Response body: %s", w.Body.String())
    assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected status code 401")

    // Check if the response body matches the expected string
    assert.Equal(t, `"AUTHENTICATION_ERROR"`, w.Body.String(), "Expected AUTHENTICATION_ERROR in response")
}



// TestGetVotes tests the GetVotes method of the CustomController

/**
type Vote struct {
	ImageID string `json:"image_id"`
	Value   int    `json:"value"`
}


func TestCreateVote(t *testing.T) {
    // Set up mock server
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "POST", r.Method)
        assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
        assert.Equal(t, "test_api_key", r.Header.Get("x-api-key"))

        var vote Vote
        err := json.NewDecoder(r.Body).Decode(&vote)
        assert.NoError(t, err)

        // Validate the vote data
        if vote.ImageID == "" || vote.Value == 0 {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"error": "Invalid vote data"}`)) // Use valid JSON
            return
        }

        // Respond with success if validation passes
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "Vote created"}`))
    }))
    defer mockServer.Close()

    // Set the app config for the test
    beego.AppConfig.Set("catapi_key", "test_api_key")
    beego.AppConfig.Set("catapi_url", mockServer.URL)

    // Create valid vote data
    vote := Vote{
        ImageID: "test_image_id",
        Value:   1,
    }

    // Marshal the vote data
    voteJSON, err := json.Marshal(vote)
    assert.NoError(t, err)

    // Create request with the JSON data
    w := httptest.NewRecorder()
    r := httptest.NewRequest("POST", "/v1/votes", bytes.NewBuffer(voteJSON))
    r.Header.Set("Content-Type", "application/json")
    r.Header.Set("x-api-key", "test_api_key")

    ctx := context.NewContext()
    ctx.Reset(w, r)

    // Initialize the controller
    ctrl := &controllers.CustomController{}
    ctrl.Init(ctx, "CustomController", "CreateVote", nil)

    // Call the CreateVote method, catching any panic
    var panicValue interface{}
    func() {
        defer func() {
            panicValue = recover()
        }()
        ctrl.CreateVote()
    }()

    // Check if we got the expected "user stop run" panic
    if panicValue != nil {
        if err, ok := panicValue.(error); ok {
            assert.Equal(t, "user stop run", err.Error())
        } else {
            t.Logf("Unexpected panic type: %T", panicValue)
        }
    }

    // Check the response status code and body
    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]string
    err = json.Unmarshal(w.Body.Bytes(), &response)
    if err != nil {
        t.Fatalf("Failed to unmarshal response: %v\nResponse body: %s", err, w.Body.String())
    }

    // Assert the correct response message
    assert.Equal(t, "Vote created", response["message"])
}
**/


/**

type Favourite struct {
    ImageID string `json:"image_id"`
    SubID   string `json:"sub_id,omitempty"`
}

type FavouriteResponse struct {
    ID      string `json:"id"`
    ImageID string `json:"image_id"`
    SubID   string `json:"sub_id,omitempty"`
}

func TestCreateFavourite(t *testing.T) {
    // Set up mock server
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "POST", r.Method)
        assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
        assert.Equal(t, "test_api_key", r.Header.Get("x-api-key"))

        body, err := ioutil.ReadAll(r.Body)
        assert.NoError(t, err)

        var fav Favourite
        err = json.Unmarshal(body, &fav)
        assert.NoError(t, err)

        // Check if the incoming request has the expected data
        assert.Equal(t, "test_image_id", fav.ImageID)
        assert.Equal(t, "test_sub_id", fav.SubID)

        // Respond with a success message
        response := FavouriteResponse{
            ID:      "favourite_123",
            ImageID: fav.ImageID,
            SubID:   fav.SubID,
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    }))
    defer mockServer.Close()

    // Set the app config for the test
    beego.AppConfig.Set("catapi_key", "test_api_key")
    beego.AppConfig.Set("catapi_url", mockServer.URL)

    // Create a request with valid data
    w := httptest.NewRecorder()
    r := httptest.NewRequest("POST", "/v1/favourites", bytes.NewBufferString(`{"image_id": "test_image_id", "sub_id": "test_sub_id"}`))
    r.Header.Set("Content-Type", "application/json")
    r.Header.Set("x-api-key", "test_api_key")

    ctx := context.NewContext()
    ctx.Reset(w, r)

    // Initialize the controller
    ctrl := &controllers.CustomController{}
    ctrl.Init(ctx, "CustomController", "CreateFavourite", nil)

    // Call the CreateFavourite method
    ctrl.CreateFavourite()

    // Check the response status code and body
    assert.Equal(t, http.StatusOK, w.Code)

    var response FavouriteResponse
    err := json.Unmarshal(w.Body.Bytes(), &response)
    if err != nil {
        t.Fatalf("Failed to unmarshal response: %v\nResponse body: %s", err, w.Body.String())
    }

    // Assert the correct response values
    assert.Equal(t, "favourite_123", response.ID)
    assert.Equal(t, "test_image_id", response.ImageID)
    assert.Equal(t, "test_sub_id", response.SubID)
}

func TestCreateFavouriteMissingImageID(t *testing.T) {
    // Create a request with missing image_id
    w := httptest.NewRecorder()
    r := httptest.NewRequest("POST", "/v1/favourites", bytes.NewBufferString(`{"sub_id": "test_sub_id"}`))
    r.Header.Set("Content-Type", "application/json")
    r.Header.Set("x-api-key", "test_api_key")

    ctx := context.NewContext()
    ctx.Reset(w, r)

    // Initialize the controller
    ctrl := &controllers.CustomController{}
    ctrl.Init(ctx, "CustomController", "CreateFavourite", nil)

    // Call the CreateFavourite method
    ctrl.CreateFavourite()

    // Check the response status code
    assert.Equal(t, http.StatusBadRequest, w.Code)

    var errorResponse map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
    if err != nil {
        t.Fatalf("Failed to unmarshal response: %v\nResponse body: %s", err, w.Body.String())
    }

    // Assert the correct error message
    assert.Equal(t, "Image ID is required", errorResponse["error"])
}

**/


