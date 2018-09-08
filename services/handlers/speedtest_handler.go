package handlers

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) (int, error) {
	return http.StatusOK, nil
}

func Upload(w http.ResponseWriter, r *http.Request) (int, error) {
	return http.StatusOK, nil
}

func Download(w http.ResponseWriter, r *http.Request) (int, error) {
	return http.StatusOK, nil
}
