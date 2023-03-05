package src

import (
	"reflect"

	"github.com/gofiber/websocket/v2"
)

func removeClient(ws *websocket.Conn) {
	// Remove the client from the clients slice
	for i, client := range clients {
		if client == ws {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	//close client if there is empty players
	if len(clients) == 0 {
		ws.Close()
	}
}

func indexOf(data []*websocket.Conn, element *websocket.Conn) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func removeElement(slice interface{}, elem interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("removeElement: first argument must be a slice")
	}

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(s.Index(i).Interface(), elem) {
			reflect.Copy(s.Slice(i, s.Len()), s.Slice(i+1, s.Len()))
			s.SetLen(s.Len() - 1)
			return
		}
	}
}
