package main

import (
    "github.com/mypianoplayer/ragtime/game"
    )

type Input struct {
    game.Object
    clicked bool
    pos [2]float32
}


