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
		if c.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest" {
			c.CustomAbort(http.StatusInternalServerError, "Failed to create request")
		} else {
			c.Data["CatImageURL"] = ""
			c.TplName = "custom_page.tpl"
		}
		return
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		if c.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest" {
			c.CustomAbort(http.StatusInternalServerError, "Failed to fetch data from API")
		} else {
			c.Data["CatImageURL"] = ""
			c.TplName = "custom_page.tpl"
		}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest" {
			c.CustomAbort(http.StatusInternalServerError, "Failed to read response body")
		} else {
			c.Data["CatImageURL"] = ""
			c.TplName = "custom_page.tpl"
		}
		return
	}

	var catImages []CatImage
	if err := json.Unmarshal(body, &catImages); err != nil || len(catImages) == 0 {
		if c.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest" {
			c.CustomAbort(http.StatusInternalServerError, "Failed to parse response")
		} else {
			c.Data["CatImageURL"] = ""
			c.TplName = "custom_page.tpl"
		}
		return
	}

	if c.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest" {
		c.Data["json"] = map[string]string{"url": catImages[0].URL}
		c.ServeJSON()
	} else {
		c.Data["CatImageURL"] = catImages[0].URL
		c.TplName = "custom_page.tpl"
	}
}
