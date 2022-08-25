package server

import (
	"Crunch-Garage/go-video-call/config"
	controller "Crunch-Garage/go-video-call/controllers"
	"flag"
	"time"

	w "Crunch-Garage/go-video-call/packages/webrtc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)

var (
	addr = flag.String("addr", ":"+config.EnvPort(), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())
	app.Use(cors.New())

	/*routes*/
	app.Get("/", controller.Welcome)
	app.Get("/room/create", controller.RoomCreate)
	app.Get("/room/:uuid", controller.Room)
	app.Get("/room/:uuid/websocket", websocket.New(controller.RoomWebSocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/room/:uuid/chat", controller.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(controller.RoomChatWebsocket))
	app.Get("/room/uuid/viewer/websocket", websocket.New(controller.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", controller.Stream)
	app.Get("/stream/:ssuid/websocket", websocket.New(controller.StreamWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/stream/ssuid/chat/websocket", websocket.New(controller.StreamChatWebsocket))
	app.Get("/stream/:ssuid/viewer/websocket", websocket.New(controller.StreamViewerWebsocket))
	app.Static("/", "./assets")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)
	go dispatchKeyFrames()
	/*check for certificates*/
	if *cert != "" {
		return app.ListenTLS(*addr, *cert, *key)
	}
	return app.Listen(*addr)

}

/*go routine*/
func dispatchKeyFrames() {
	for range time.NewTicker(time.Second * 3).C {
		for _, room := range w.Rooms {
			room.Peers.DispatchKeyFrames()
		}
	}

}

func HandleRoutes()
