package component

import (
    "log"
    "github.com/mypianoplayer/ragtime/game"
    )

type View struct {
    game.ComponentBase
    text string
}

func NewView(t string) *View {
    v := &View {
        ComponentBase:*game.NewComponent(),
        text:t,
    }
    return v
}

func (v *View) Start() {
    
}

func (v *View) Update() {
    log.Println("view update")
}

func (v *View) UpdateOrder() int {
    return 1;
}
