package service
import (
	"net/http"
	"github.com/unrolled/render"
)
func TableHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.HTML(w, http.StatusOK, "table", struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} {
			Username: req.Form["username"][0],
			Password: req.Form["password"][0],
		})
	}
}