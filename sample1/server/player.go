package main

import (
    "log"
    // "reflect"
    "github.com/mypianoplayer/ragtime/game"
    "github.com/mypianoplayer/ragtime_sample/sample1/server/component"
    )

type Player struct {
    game.ObjectBase
    component.Position
    component.View
}

func NewPlayer() *Player {
    p := Player {
        ObjectBase:*game.NewObject("player"),
        Position:component.Position{},
        View:*component.NewView("{-]"),
    }
    game.SetupComponent(&p)
    
    return &p
}


func (p *Player) Test() {
    log.Println("TEST")
}

