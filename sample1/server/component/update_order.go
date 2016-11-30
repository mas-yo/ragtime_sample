package component

import (
    _ "log"
    "github.com/mypianoplayer/ragtime/game"
    )

const (
    OrderInput int = iota
    OrderPosition
    OrderView
)

const (
    ComponentType_Input game.ComponentType = iota
    ComponentType_Position
    ComponentType_View
)