package component

import (
    _ "log"
    "strconv"
    "github.com/mypianoplayer/ragtime/game"
    "github.com/mypianoplayer/ragtime/server"
    )

type View struct {
    game.ComponentBase
    text string
    pos *Position
    sendMsgCh chan *server.Message
}

func NewView(t string, sendMsgCh chan *server.Message) *View {
    
    msg := server.Message{
        Params:[]string{"newObject", "player", t},
    }
    sendMsgCh <- &msg
    
    v := View {
        ComponentBase:*game.NewComponent(),
        text:t,
        sendMsgCh:sendMsgCh,
    }
    return &v
}

func (v *View) Start() {
    for c := range game.EachComponent(v.Object()) {
        pos,ok := c.(*Position)
        if ok {
            v.pos = pos
        }
    }
}

func (v *View) Update() {
    // log.Println("view update")
    
    msg := server.Message {
        Params:[]string{"setpos", "player", strconv.FormatFloat(float64(v.pos.Pos()[0]),'f',-1,32), strconv.FormatFloat(float64(v.pos.Pos()[1]),'f',-1,32)},
    }
    v.sendMsgCh <- &msg
}

func (v *View) UpdateOrder() int {
    return 1;
}
