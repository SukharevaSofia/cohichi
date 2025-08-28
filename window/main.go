package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - keyboard input")
	ballPosition := rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)

	cohi_texture := rl.LoadTexture("../assets/baby.png")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyRight) {
			ballPosition.X += 0.8
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			ballPosition.X -= 0.8
		}
		if rl.IsKeyDown(rl.KeyUp) {
			ballPosition.Y -= 0.8
		}
		if rl.IsKeyDown(rl.KeyDown) {
			ballPosition.Y += 0.8
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Magenta)

		rl.DrawText("move the ball with arrow keys", 10, 10, 20, rl.DarkGray)
		rl.DrawTextureV(cohi_texture, ballPosition, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
