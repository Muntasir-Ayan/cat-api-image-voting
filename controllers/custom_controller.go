package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"

	beego "github.com/beego/beego/v2/server/web"
)

type Vote struct {
	ImageID string `json:"image_id"`
	Value   int    `json:"value"`
}

type VoteResponse struct {
	ID      int    `json:"id"`
	ImageID string `json:"image_id"`
	Value   int    `json:"value"`
}

type CatImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
	ImageID     string `json:"reference_image_id"`
}

type BreedImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Breeds []struct {
		Name        string `json:"name"`
		Origin      string `json:"origin"`
		Description string `json:"description"`
		Wikipedia   string `json:"wikipedia_url"`
	} `json:"breeds"`
}

type CustomController struct {
	beego.Controller
}

// Fetches a random cat image
func (c *CustomController) Get() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	url := "https://api.thecatapi.com/v1/images/search"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.handleError(err)
		return
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		c.handleError(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.handleError(err)
		return
	}

	var catImages []CatImage
	if err := json.Unmarshal(body, &catImages); err != nil || len(catImages) == 0 {
		c.handleError(err)
		return
	}

	c.Data["CatImageURL"] = catImages[0].URL
	c.TplName = "custom_page.tpl"
}

// Fetches list of breeds
func (c *CustomController) GetBreeds() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	url := "https://api.thecatapi.com/v1/breeds"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Failed to fetch breeds")
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var breeds []Breed
	if err := json.Unmarshal(body, &breeds); err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Failed to parse breeds")
		return
	}

	c.Data["json"] = breeds
	c.ServeJSON()
}

// Fetches images and details for a specific breed
// Fetches images and details for a specific breed
func (c *CustomController) GetBreedImages() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	breedID := c.GetString("breed_id")
	if breedID == "" {
		c.CustomAbort(http.StatusBadRequest, "Missing breed ID")
		return
	}

	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=8&breed_ids=%s", breedID)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Failed to fetch breed images")
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var images []BreedImage
	if err := json.Unmarshal(body, &images); err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Failed to parse breed images")
		return
	}

	// Log breed information for debugging
	for _, image := range images {
		for _, breed := range image.Breeds {
			fmt.Printf("Breed Name: %s, Wikipedia: %s\n", breed.Name, breed.Wikipedia)
		}
	}

	c.Data["json"] = images
	c.ServeJSON()
}

func (c *CustomController) CreateVote() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	imageID := c.GetString("image_id")
	voteValue, err := c.GetInt("value")
	if err != nil || imageID == "" {
		c.CustomAbort(http.StatusBadRequest, "Invalid vote data")
		return
	}

	url := "https://api.thecatapi.com/v1/votes"

	vote := Vote{
		ImageID: imageID,
		Value:   voteValue,
	}

	voteData, _ := json.Marshal(vote)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(voteData))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Failed to create vote")
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	c.Data["json"] = string(body) // Return response from the API
	c.ServeJSON()
}

func (c *CustomController) GetVotes() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	
	// Get query parameters
	limit := c.GetString("limit") // No default value here
	order := c.GetString("order", "DESC") // Default to DESC if not provided

	// Construct base URL
	url := "https://api.thecatapi.com/v1/votes"

	// Add query parameters if provided
	query := url + "?"
	if limit != "" {
		query += "limit=" + limit + "&"
	}
	query += "order=" + order

	req, _ := http.NewRequest("GET", query, nil)
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Failed to fetch votes")
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var votes []VoteResponse
	if err := json.Unmarshal(body, &votes); err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Failed to parse votes")
		return
	}

	c.Data["json"] = votes
	c.ServeJSON()
}




// Error handling for XMLHttpRequest
func (c *CustomController) handleError(err error) {
	if c.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest" {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	} else {
		c.Data["CatImageURL"] = ""
		c.TplName = "custom_page.tpl"
	}
}
