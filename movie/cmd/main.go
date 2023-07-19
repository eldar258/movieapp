package main

import (
	"log"
	"movieapp/movie/internal/controller/movie"
	metadatagateway "movieapp/movie/internal/gateway/metadata/http"
	retinggateway "movieapp/movie/internal/gateway/rating/http"
	httphandler "movieapp/movie/internal/handelr/http"
	"net/http"
)

func main() {
	log.Println("Start the movie service")

	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := retinggateway.New("localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
