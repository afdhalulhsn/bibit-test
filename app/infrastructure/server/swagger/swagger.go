package swagger

import (
	"net/http"

	_ "bibit/cmd/server/main/docs"
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
)

func RunSwagger() {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition
	))
	http.ListenAndServe(":1323", r)
}
