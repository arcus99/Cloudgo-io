package service

import (
	"net/http"
	"github.com/unrolled/render"
)

func RegistHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "regist", nil)
	}

}
