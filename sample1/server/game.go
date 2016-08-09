package main

import (
    "log"
    "strconv"
    "fmt"
    "github.com/mypianoplayer/ragtime/game"
    "github.com/mypianoplayer/ragtime/server"
    _ "github.com/mypianoplayer/ragtime_sample/sample1/server/component"
    "net/http"
    )

type Game struct {
    scene *game.Scene
    server *server.Server
    
    players map[int]*Player
}

func NewGame() *Game {
    sv := server.New("/game")
    sc := game.NewScene(sv.RecvCh(), nil)
    g := &Game{
        scene:sc,
        server:sv,
        players:make(map[int]*Player),
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

    // test := component.NewPosition()
    // for t := range game.EachComponent(test) {
    //     log.Println(t)
    // }


	http.Handle("/", http.FileServer(http.Dir("../client/")))

	g.scene.Start()
	g.server.Start()

}

func (g *Game) OnMessage(msg *server.Message) {
    log.Println("onmsg", msg)
    if msg.Params[0] == "start" {
        name := fmt.Sprintf("player%d", msg.ID)
        player := NewPlayer(name, g.server.SendAllCh())
        g.scene.AddObject(player)
        g.started = true
        
        g.players[msg.ID] = player
    }
    
    
    if g.started && msg.Params[0] == "click" {
        x, _ := strconv.Atoi(msg.Params[1])
        y, _ := strconv.Atoi(msg.Params[2])
        
        p, ok := g.players[msg.ID]
        if ok {
            p.InputComponent().SetPos([2]float32{ float32(x), float32(y) })
        }
    }
}
