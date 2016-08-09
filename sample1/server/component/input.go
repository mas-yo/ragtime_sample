package component

import (
    _ "log"
    "github.com/mypianoplayer/ragtime/game"
    )


type Input struct {
    game.ComponentBase
    pos [2]float32
}

func NewInput() *Input {
    return &Input {
        ComponentBase:*game.NewComponentBase(OrderInput),
    }
}

func (i *Input) SetPos(pos [2]float32) {
    i.pos = pos
}

func (i *Input) Pos() [2]float32 {
    return i.pos
}

func (p *Input) Start() {
    
}

func (p *Input) Update() {
    // log.Println("pos update")

    // server.SendAll pos
}
