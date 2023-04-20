package router

import (
	"fmt"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

func APIController() {
	var (
		handler    *mux.Router
		prefixPath string
	)

	handler = mux.NewRouter()
	prefixPath = config.ApplicationConfiguration.GetServer().PrefixPath
	if prefixPath != "" {
		prefixPath = "/" + prefixPath
	}

	auth := handler.PathPrefix("/v1" + prefixPath + "/auth").Subrouter()
	auth.HandleFunc("/register", endpoint.AuthEndpoint.RegisterUser).Methods("POST")
	auth.HandleFunc("/login", endpoint.AuthEndpoint.Login).Methods("POST")
	auth.HandleFunc("/logout", endpoint.AuthEndpoint.Logout).Methods("GET")
	auth.Use(Middleware)

	api := handler.PathPrefix("/v1" + prefixPath + "/api").Subrouter()
	api.HandleFunc("/recruitment/position", endpoint.RecruitmentEndpoint.GetListPosition).Methods("GET")
	api.HandleFunc("/recruitment/position/{ID}", endpoint.RecruitmentEndpoint.GetDetailPosition).Methods("GET")
	api.Use(JWTMiddleware)

	fmt.Print(http.ListenAndServe(fmt.Sprintf(`%s:%s`, config.ApplicationConfiguration.GetServer().Host, config.ApplicationConfiguration.GetServer().Port), handler))
}
