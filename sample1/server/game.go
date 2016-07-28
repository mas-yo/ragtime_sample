package main

import (
    "log"
    "github.com/mypianoplayer/ragtime/game"
    "github.com/mypianoplayer/ragtime/server"
    "net/http"
    )

type Game struct {
    scene *game.Scene
    server *server.Server
}

func NewGame() *Game {
    sv := server.New("/game")
    sc := game.NewScene(sv.RecvCh(), nil)
    g := &Game{
        scene:sc,
        server:sv,
    }
    g.scene.SetReceiver(g)
    return g
}

func (g *Game) Scene() *game.Scene {
    return g.scene
}

func (g *Game) Server() *server.Server {
    return g.server
}

func (g* Game) Start() {


// 	player := NewPlayer()

// 	g.scene.AddObject(player)

	http.Handle("/", http.FileServer(http.Dir("../client/")))

	g.scene.Start()
	g.server.Start()

}

func (g *Game) OnMessage(msg *server.Message) {
    log.Println("onmsg", msg)
    if msg.Params[0] == "start" {
        player := NewPlayer()
        g.scene.AddObject(player)
    }
}
