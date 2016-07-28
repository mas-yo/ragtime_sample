package component

import (
    "log"
    "github.com/mypianoplayer/ragtime/game"
    )


type Position struct {
    game.ComponentBase
    pos [2]float32
}

func NewPosition() *Position {
    return &Position {
        ComponentBase:*game.NewComponent(),
    }
}

func (p *Position) IsDeleted() bool {
    return false
}

func (p *Position) Start() {
    
}

func (p *Position) Update() {
    log.Println("pos update")

    // server.SendAll pos
}

func (p *Position) UpdateOrder() int {
    return 0
}
