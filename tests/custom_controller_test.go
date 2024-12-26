package tests

import (
	// "encoding/json"
	"testing"
	"net/http"
	"net/http/httptest"
	"strings" // Required for checking the URL prefix
	"github.com/stretchr/testify/assert"
	beego "github.com/beego/beego/v2/server/web"
	"cat-api/controllers" // Adjust import path as necessary
)

// Mock response structure for breeds
type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Mock response structure for breed images
type BreedImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Breeds []Breed `json:"breeds"`
}

type response struct {
	data []byte
	err  error
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


