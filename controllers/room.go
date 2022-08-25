package controller

import (
	"fmt"

	w "Crunch-Garage/go-video-call/packages/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	uuid, suuid, _ := createOrGetRoom(uuid)
}

func RoomWebSocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	/*create room if it does not exist else get room if ut exists*/
	_, _, room := createOrGetRoom(uuid)
	w.RoomConn(c, room.Peer)
}

/*create or get room*/
func createOrGetRoom(uuid string) (string, string, *w.Room) {

}

func RoomViewerWebsocket(c *websocket.Conn){

}

func RoomViewerConn(c *websocket.Conn, p *w.Peers){

}

type WebsocketMessage struct {
	Event string `json:"event"`
	Data string `json:"data"`
}