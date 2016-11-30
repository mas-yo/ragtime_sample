package component

import (
    _ "log"
    "github.com/mypianoplayer/ragtime/game"
    )


type Position struct {
    game.ComponentBase
    pos [2]float32
    input *Input
}

func NewPosition() *Position {
    return &Position {
        ComponentBase:*game.NewComponentBase(ComponentType_Position,OrderPosition),
    }
}

func (p *Position) IsDeleted() bool {
    return false
}

func (p *Position) Pos() [2]float32 {
    return p.pos
}

func (p *Position) Start() {
    p.input, _ = p.Object().Components()[ComponentType_Input].(*Input)
}

func (p *Position) Update() {
    p.pos[0] = p.pos[0] + (p.input.Pos()[0] - p.pos[0]) / 20.0
    p.pos[1] = p.pos[1] + (p.input.Pos()[1] - p.pos[1]) / 20.0
}

