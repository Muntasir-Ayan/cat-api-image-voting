// custom_controller.go
package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type CatImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type CustomController struct {
	beego.Controller
}

func (c *CustomController) Get() {
	apiKey, _ := beego.AppConfig.String("catapi_key")
	url := "https://api.thecatapi.com/v1/images/search"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Data["CatImageURL"] = ""
		c.TplName = "custom_page.tpl"
		return
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		c.Data["CatImageURL"] = ""
		c.TplName = "custom_page.tpl"
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["CatImageURL"] = ""
		c.TplName = "custom_page.tpl"
		return
	}

	var catImages []CatImage
	if err := json.Unmarshal(body, &catImages); err != nil || len(catImages) == 0 {
		c.Data["CatImageURL"] = ""
	} else {
		c.Data["CatImageURL"] = catImages[0].URL
	}

	c.TplName = "custom_page.tpl"
}
