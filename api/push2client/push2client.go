package push2client

import (
	"encoding/json"
	"go-websocket/api"
	"go-websocket/define/retcode"
	"go-websocket/servers/server"
	"net/http"
)

type Controller struct {
}

type pushToClientInputData struct {
	ClientId string `json:"clientId"`
	Message  string `json:"message"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//解析参数
	_ = r.ParseForm()
	var inputData pushToClientInputData
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//发送信息
	server.SendMessage2Client(&inputData.ClientId, &inputData.Message)

	api.Render(w, retcode.SUCCESS, "success", []string{})
	return
}
