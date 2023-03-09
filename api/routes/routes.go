package routes

import (
	c_auth "MapProject/api/controllers/auth"
	c_inside "MapProject/api/controllers/inside"
	c_pin "MapProject/api/controllers/pin"
	c_polygon "MapProject/api/controllers/polygon"
	jwt2 "MapProject/pkg/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	user := app.Group("/auth")
	user.Post("/login", c_auth.Login) //generates jwt token and send it back //tamamlandi

	create := app.Group("/create", jwt2.JwtMiddlewareWrapper(jwt2.JWTMiddleware))
	create.Post("/polygon", c_polygon.CreatePolygon)              //tammamlandi
	create.Post("/pin", c_pin.CreatePin)                          //tamamlandi
	create.Post("/inside", c_inside.Inside)                       //tamamlandi
	create.Post("/getinfofrompolid", c_inside.GetInfoFromPolygon) //tamamlandi
	create.Post("/getinfofrompinid", c_inside.GetInfoFromPin)     //tamamlandi

	get := app.Group("/get", jwt2.JwtMiddlewareWrapper(jwt2.JWTMiddleware))
	get.Get("/polygons", c_polygon.GetPolygonsByUserID) //tamamlandi
	get.Get("/pins", c_pin.GetPinsByUserID)             //tamamlandi
	//get.Get("getallinfo", c_inside.GetAllInfo) //sonra tamamlanacak (just model)

	delete := app.Group("/delete", jwt2.JwtMiddlewareWrapper(jwt2.JWTMiddleware))
	//kullanici direkt pin silebilir
	delete.Delete("/pin", c_pin.DeletePin)              //tamamlandi
	delete.Delete("/deleteallpins", c_pin.DeleteAllPin) //tamamlandi

	//kullanici direkt poligon silebilir
	delete.Delete("/polygon", c_polygon.DeletePolygon) //tamamlandi

	//kullanici poligon noktasi silebilir
	delete.Delete("/onepoint", c_polygon.DeleteOnePointByUserID) //tamamlandi

	update := app.Group("/update", jwt2.JwtMiddlewareWrapper(jwt2.JWTMiddleware))
	update.Put("/polygon", c_polygon.UpdatePolygon) //tamamlandi
	update.Put("/pin", c_pin.UpdatePin)             //tamamlandi
}
