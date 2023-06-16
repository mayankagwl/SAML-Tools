package main

import (
	"fmt"
	_config "github.com/Compile7/SAML-Tools/config"
	_handlers "github.com/Compile7/SAML-Tools/handlers"
	_middleware "github.com/Compile7/SAML-Tools/middleware"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config := _config.GetInstance()
	router := mux.NewRouter()
	middleware := _middleware.NewMiddleware()
	handler := _handlers.NewHandler()

	router.Use(middleware.PostBodyMiddleware)
	handler.SetRoutes(router)

	if len(config.LambdaTaskRoot) > 0 {
		adapter := gorillamux.NewV2(router)
		lambda.Start(adapter.ProxyWithContext)
	} else {
		fmt.Printf(`Server listing on htp://%s`, config.ServiceAddress)
		log.Fatal(http.ListenAndServe(config.ServiceAddress, router))
	}
}
