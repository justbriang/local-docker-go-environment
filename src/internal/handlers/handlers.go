package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

//payloads for websocket
var wsChan = make(chan WsPayload)

//list of clients online
var clients = make(map[WsConnection]string)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

//upgrade regular http connection to websocket
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, r, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

type WsConnection struct {
	*websocket.Conn
}

//defines response sent back from a websocket connection
type WsResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

//defines typical payload sent from/to a websocket connection
type WsPayload struct {
	Action      string       `json:"action"`
	Username    string       `json:"username"`
	Message     string       `json:"message"`
	currentConn WsConnection `json:"_"`
}

//upgrades connection to websocket
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client connected at endpoint")
	var response WsResponse
	response.Message = `<em><small>Welcome to the websocket endpoint</small></em>`
	conn := WsConnection{Conn: ws}
	clients[conn] = ""
	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
	go ListenForWs(&conn)
}
func ListenForWs(conn *WsConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error in ListenForWs: ", fmt.Sprintf("%v", r))
		}

	}()
	var payload WsPayload
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {

		} else {
			payload.currentConn = *conn
			wsChan <- payload

		}
	}

}
func ListenToWsChannel() {
	var response WsResponse
	for {
		payload := <-wsChan
		response.Action = "Got Here"
		response.Message = payload.Message
		log.Println("Received payload: ", payload)

		switch payload.Action {
		case "username":
			clients[payload.currentConn] = payload.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			BroadCastToAll(response)
		case "left":
			response.Action = "list_users"
			delete(clients, payload.currentConn)
			users := getUserList()
			response.ConnectedUsers = users
			BroadCastToAll(response)
		case "message":
			response.Action = "message"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", payload.Username, payload.Message)
			BroadCastToAll(response)
		}

	}
}
func getUserList() []string {
	var userList []string
	for _, client := range clients {
		if client != "" {
			userList = append(userList, client)
		}
	}
	sort.Strings(userList)
	return userList
}

func BroadCastToAll(response WsResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println(err)
		}
	}
}
func renderPage(w http.ResponseWriter, r *http.Request, page string, data jet.VarMap) error {
	view, err := views.GetTemplate(page)
	if err != nil {
		log.Println(err)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		return err
	}
	return nil
}