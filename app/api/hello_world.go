package api

import (
	"github.com/jinzhu/gorm"
	"github.com/k8s-go-demo/utils"
	"net/http"
)

func GetHelloWorld(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, "Hello World")
}