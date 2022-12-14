/*
 * Otto notification service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"github.com/danielmunro/otto-notification-service/internal"
	"github.com/danielmunro/otto-notification-service/internal/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	serveHttp()
}

func serveHttp() {
	router := internal.NewRouter()
	handler := cors.AllowAll().Handler(router)
	log.Print("http listening on 8083")
	log.Fatal(http.ListenAndServe(":8083",
		middleware.ContentTypeMiddleware(handler)))
}
