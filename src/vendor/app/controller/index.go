package controller

import (
	"net/http"
	
	"app/shared/view"
)

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "index/index"
	v.Render(w)
}
