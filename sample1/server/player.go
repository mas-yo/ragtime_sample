package main

import (
    "log"
    // "reflect"
    "github.com/mypianoplayer/ragtime/game"
    "github.com/mypianoplayer/ragtime/server"
    "github.com/mypianoplayer/ragtime_sample/sample1/server/component"
    )

type Player struct {
    game.ObjectBase
    component.Input
    component.Position
    component.View
}

func NewPlayer(sendMsgCh chan *server.Message) *Player {
    p := Player {
        ObjectBase:*game.NewObject("player"),
        Input:*component.NewInput(),
        Position:*component.NewPosition(),
        View:*component.NewView("[-]", sendMsgCh),
    }
    game.SetupComponent(&p)
    
    return &p
}

func (p *Player) InputComponent() *component.Input {
    return &p.Input
}

func (p *Player) Test() {
    log.Println("TEST")
}

