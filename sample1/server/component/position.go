package component

import (
    "log"
    "github.com/mypianoplayer/ragtime/game"
    )


type Position struct {
    game.ComponentBase
    pos [2]float32
    input *Input
}

func NewPosition() *Position {
    return &Position {
        ComponentBase:*game.NewComponentBase(OrderPosition),
    }
}

func (p *Position) IsDeleted() bool {
    return false
}

func (p *Position) Pos() [2]float32 {
    return p.pos
}

func (p *Position) Start() {
    
    for c := range game.EachComponent(p.Object()) {
        
        log.Println("component : ", c)
        
        in,ok := c.(*Input)
        if ok {
            p.input = in
            log.Println(" set ref input ok" )
        }
    }
}

func (p *Position) Update() {
    p.pos[0] = p.pos[0] + (p.input.Pos()[0] - p.pos[0]) / 20.0
    p.pos[1] = p.pos[1] + (p.input.Pos()[1] - p.pos[1]) / 20.0
}

