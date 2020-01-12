package bindgroup

import (
	"encoding/json"
	"fmt"
	"go-websocket/api"
	"go-websocket/define/retcode"
	"go-websocket/servers/server"
	"net/http"
)

type Controller struct {
}

type bindToGroupInputData struct {
	ClientId  string `json:"clientId"`
	GroupName string `json:"groupName"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//解析参数
	_ = r.ParseForm()
	var inputData bindToGroupInputData
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(inputData.ClientId) > 0 && len(inputData.GroupName) > 0 {
		server.AddClient2Group(&inputData.GroupName, &inputData.ClientId)
	} else {
		fmt.Println("参数错误")
	}

	api.Render(w, retcode.SUCCESS, "success", []string{})
}
