// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example - Demonstrates texture coordinates.
package main

import (
	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/window"
)

var game *Game

// gfxLoop is responsible for drawing things to the window.
func gfxLoop(w window.Window, d gfx.Device) {

	game.Init(w, d)

	for {
		game.Update(w, d)
	}
}

func main() {
	game = NewGame()
	window.Run(gfxLoop, nil)
}
