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
    return &Player {
        ObjectBase:*game.NewObject("player"),
        Position:component.Position{},
        View:*component.NewView("{-]"),
    }
}


func (p *Player) Test() {
    log.Println("TEST")
}

