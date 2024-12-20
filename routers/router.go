package routers

import (
	"cat-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/custom", &controllers.CustomController{})
	beego.Router("/custom/breeds", &controllers.CustomController{}, "get:GetBreeds")
	beego.Router("/custom/breed_images", &controllers.CustomController{}, "get:GetBreedImages")
	beego.Router("/custom/vote", &controllers.CustomController{}, "post:CreateVote")
	beego.Router("/custom/votes", &controllers.CustomController{}, "get:GetVotes")

	// New routes for favourites
	beego.Router("/custom/favourite", &controllers.CustomController{}, "post:CreateFavourite") // For creating a favourite
	beego.Router("/custom/favourites", &controllers.CustomController{}, "get:GetFavourites")
}
