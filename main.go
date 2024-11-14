package main

import( 
	"github.com/hajimehoshi/ebiten/v2"
    "shoot/assets"
)


func (g *Game) Update() error {
	return nil
}

func(g *Game) Draw(screen *ebiten.Image){
	screen.DrawImage(assets.PlayerSprite,nil)
}

func(g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int){
	return outsideWidth, outsideHeight
}

func main(){
	g := &Game{}

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}
 