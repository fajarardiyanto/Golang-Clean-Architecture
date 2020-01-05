package controllers

import (
	"net/http"

	"github.com/fajarardiyanto/portal/api/response"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome to this Awsome API")
}
