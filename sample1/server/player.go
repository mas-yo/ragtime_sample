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
}

func NewPlayer(name string, sendMsgCh chan *server.Message) *Player {
    p := Player {
        ObjectBase:*game.NewObjectBase(name),
    }
    p.ObjectBase.AddComponent( component.NewInput() )
    p.ObjectBase.AddComponent( component.NewPosition() )
    p.ObjectBase.AddComponent( component.NewView(name, "[-]", sendMsgCh) )
    p.SetupComponent()
    
    return &p
}

func (p *Player) InputComponent() *component.Input {
    ret,_ := p.Components()[component.ComponentType_Input]
    return ret.(*component.Input)
}

func (p *Player) Test() {
    log.Println("TEST")
}
