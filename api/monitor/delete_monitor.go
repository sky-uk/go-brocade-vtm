package monitor

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// DeleteMonitorAPI : object used to call delete on monitor
type DeleteMonitorAPI struct {
	*rest.BaseAPI
}

// NewDelete : returns a new DeleteMonitorAPI object
func NewDelete(name string) *DeleteMonitorAPI {
	this := new(DeleteMonitorAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/monitors/"+name, nil, nil, new(api.VTMError))
	return this
}
