package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	isGameStarted         = false
	titleFont             font.Face
	titleOptionsFont      font.Face
	game                  *Game
	maxCameraZoom         float32 = 50
	minCameraZoom         float32 = 10
	wallWidth             float32 = 1
	currentlySelectedUnit *Unit
)

func init() {
	snftFont, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	titleFont, err = opentype.NewFace(snftFont, &opentype.FaceOptions{
		Size:    250,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	titleOptionsFont, err = opentype.NewFace(snftFont, &opentype.FaceOptions{
		Size:    100,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 2000, 1350
}

func main() {
	units := []*Unit{}
	units = append(units, &Unit{-5.0, 5.0, 10.0, 100, 1.0, 1, false, []*Command{}})
	units = append(units, &Unit{-5.0, -5.0, 10.0, 100, 1.0, 1, false, []*Command{}})
	units = append(units, &Unit{5.0, 5.0, 10.0, 100, 1.0, 2, false, []*Command{}})
	units = append(units, &Unit{5.0, -5.0, 10.0, 100, 1.0, 2, false, []*Command{}})
	walls := []*Wall{}
	walls = append(walls, &Wall{-10, -10, 10, 0})
	game = &Game{units, walls, -10, -10, maxCameraZoom}

	ebiten.SetWindowResizable(true)
	ebiten.RunGame(game)
}
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return &EndError{time.Now(), "Program Ended"}
	}

	if !isGameStarted {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			isGameStarted = true
		}
	} else {
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			game.cameraScale = ceilingClamp(game.cameraScale*(1/float32(ebiten.ActualTPS())+1), maxCameraZoom)
		} else if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyArrowDown)) {
			game.cameraScale = floorClamp(game.cameraScale*(1-0.9/float32(ebiten.ActualTPS())), minCameraZoom)
		}

		if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyW)) {
			game.cameraY -= 1000 / game.cameraScale / float32(ebiten.ActualTPS())
		}
		if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyS)) {
			game.cameraY += 1000 / game.cameraScale / float32(ebiten.ActualTPS())
		}

		if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyA)) {
			game.cameraX -= 1000 / game.cameraScale / float32(ebiten.ActualTPS())
		}
		if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyD)) {
			game.cameraX += 1000 / game.cameraScale / float32(ebiten.ActualTPS())
		}
	}

	return g.Simulate()
}

func (g *Game) Simulate() error {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		if currentlySelectedUnit == nil { //no unit selected ==> find unit and select it
			var clickX, clickY float32 = 0, 0
			clickX, clickY = g.getGamePositionOfMouse()
			unitclick := false
			for i := 0; i < len(g.units) && !unitclick; i++ {

				if float64(g.units[i].clickRadius) >= math.Sqrt(float64((g.units[i].positionX-clickX)*(g.units[i].positionX-clickX))+float64((g.units[i].positionY-clickY)*(g.units[i].positionY-clickY))) {
					currentlySelectedUnit = g.units[i]
					currentlySelectedUnit.isSelected = true
					unitclick = true

				}
			}
			if !unitclick && currentlySelectedUnit != nil {
				currentlySelectedUnit.isSelected = false
				currentlySelectedUnit = nil
			}
		} else if currentlySelectedUnit != nil { //unit selected ==>  select new unit,if another unit not clicked, move current unit to that spot that was clicked
			var clickX, clickY float32 = 0, 0
			clickX, clickY = g.getGamePositionOfMouse()
			unitclick := false
			for i := 0; i < len(g.units) && !unitclick; i++ {

				if float64(g.units[i].clickRadius) >= math.Sqrt(float64((g.units[i].positionX-clickX)*(g.units[i].positionX-clickX))+float64((g.units[i].positionY-clickY)*(g.units[i].positionY-clickY))) {

					if currentlySelectedUnit != g.units[i] {
						currentlySelectedUnit.isSelected = false
						currentlySelectedUnit = g.units[i]
						currentlySelectedUnit.isSelected = true
						unitclick = true
					} else {
						currentlySelectedUnit.isSelected = false
						currentlySelectedUnit = nil
						unitclick = true
					}

				}
			}
			if !unitclick { //no new unit clicked

				currentlySelectedUnit.commands = []*Command{&Command{clickX, clickY}}
			}
		}
	}

	for i := 0; i < len(game.units); i++ {
		if len(game.units[i].commands) != 0 {

			if true { //insert collision check here.
				if math.Abs(float64(game.units[i].commands[0].targetX-game.units[i].positionX)) < 0.1 && math.Abs(float64(game.units[i].commands[0].targetY-game.units[i].positionY)) < 0.1 {
					game.units[i].commands = []*Command{}
				} else {

					game.units[i].positionX += (game.units[i].commands[0].targetX - game.units[i].positionX) / float32(math.Sqrt(math.Pow(float64(game.units[i].commands[0].targetX-game.units[i].positionX), 2)+math.Pow(float64(game.units[i].commands[0].targetY-game.units[i].positionY), 2))) / float32(ebiten.TPS())
					game.units[i].positionY += (game.units[i].commands[0].targetY - game.units[i].positionY) / float32(math.Sqrt(math.Pow(float64(game.units[i].commands[0].targetX-game.units[i].positionX), 2)+math.Pow(float64(game.units[i].commands[0].targetY-game.units[i].positionY), 2))) / float32(ebiten.TPS())

				}
			}
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if !isGameStarted { // game not started
		screen.Fill(color.RGBA{89, 164, 230, 255})
		text.Draw(screen, "Simple RTS", titleFont, 330, 300, color.Black)

		text.Draw(screen, "(Space) Start", titleOptionsFont, 650, 700, color.Black)
		text.Draw(screen, "(ESC) Quit", titleOptionsFont, 725, 850, color.Black)
	} else { //game Started
		screen.Fill(color.RGBA{0, 170, 18, 255})

		for i := 0; i < len(game.walls); i++ {
			vector.StrokeLine(
				screen,
				(game.walls[i].startX-game.cameraX)*game.cameraScale+1000,
				(game.walls[i].startY-game.cameraY)*game.cameraScale+675,
				(game.walls[i].endX-game.cameraX)*game.cameraScale+1000,
				(game.walls[i].endY-game.cameraY)*game.cameraScale+675,
				game.cameraScale*wallWidth,
				color.RGBA{151, 151, 151, 255},
				false)

		}

		for i := 0; i < len(game.units); i++ {
			if game.units[i].team == 1 {

				if game.units[i].isSelected {
					vector.DrawFilledCircle(
						screen,
						(game.units[i].positionX-game.cameraX)*game.cameraScale+1000,
						(game.units[i].positionY-game.cameraY)*game.cameraScale+675,
						game.cameraScale*1.3,
						color.RGBA{255, 252, 38, 255},
						false)
				}
				vector.DrawFilledCircle(
					screen,
					(game.units[i].positionX-game.cameraX)*game.cameraScale+1000,
					(game.units[i].positionY-game.cameraY)*game.cameraScale+675,
					game.cameraScale,
					color.RGBA{202, 0, 0, 255},
					false)

			} else if game.units[i].team == 2 {
				if game.units[i].isSelected {
					vector.DrawFilledCircle(
						screen,
						(game.units[i].positionX-game.cameraX)*game.cameraScale+1000,
						(game.units[i].positionY-game.cameraY)*game.cameraScale+675,
						game.cameraScale*1.3,
						color.RGBA{255, 252, 38, 255},
						false)
				}
				vector.DrawFilledCircle(
					screen,
					(game.units[i].positionX-game.cameraX)*game.cameraScale+1000,
					(game.units[i].positionY-game.cameraY)*game.cameraScale+675,
					game.cameraScale,
					color.RGBA{0, 8, 255, 255},
					false)

			}
		}
	}
}

func floorClamp(v, lim float32) float32 {
	if v <= lim {
		return lim
	}
	return v
}

func (e *EndError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (g *Game) getGamePositionOfMouse() (x, y float32) {
	cursorX, cursorY := ebiten.CursorPosition()
	return (float32(cursorX)-1000.0)/g.cameraScale + g.cameraX, (float32(cursorY)-675)/g.cameraScale + g.cameraY
}

func doubleClamp(min, max, v float32) float32 {
	if v >= max {
		return max
	}
	if v <= min {
		return min
	}
	return v
}

func ceilingClamp(v, lim float32) float32 {
	if v >= lim {
		return lim
	}
	return v
}

func isCircleCircleIntersecting(c1 CircleCollider, c2 CircleCollider) *Vector2 {

	if math.Sqrt(math.Pow(float64(c1.c.x-c2.c.x), 2)+math.Pow(float64(c1.c.y-c2.c.y), 2)) > float64(c1.r+c2.r) {
		return nil
	}

	return &Vector2{(c1.r*c2.c.x + c2.r*c1.c.x) / (c1.r + c2.r), (c1.r*c2.c.y + c2.r*c1.c.y) / (c1.r + c2.r)}
}

func isCircleLineIntersecting(cc CircleCollider, lc LineCollider) *Vector2 {
	s := (lc.y1 - lc.y2) / (lc.x1 - lc.x2)
	h := (lc.y1 - cc.c.y - (lc.x1-cc.c.x)*s)

	positivePossibleX := (-2*s*h + float32(math.Sqrt(float64((2*s*h)*(2*s*h)-4*(s*s+1)*(h*h-cc.r*cc.r))))) / (2*s*s + 2)
	negativePossibleX := (-2*s*h - float32(math.Sqrt(float64((2*s*h)*(2*s*h)-4*(s*s+1)*(h*h-cc.r*cc.r))))) / (2*s*s + 2)

	fmt.Println(lc.isWithinDomain(positivePossibleX+cc.c.x), " - ", lc.isWithinDomain(negativePossibleX+cc.c.x))

	if lc.isWithinDomain(positivePossibleX+cc.c.x) && lc.isWithinDomain(negativePossibleX+cc.c.x) {
		return &Vector2{(negativePossibleX+positivePossibleX)/2 + cc.c.x, ((negativePossibleX+positivePossibleX)/2)*s + h + cc.c.y}
	}

	if !lc.isWithinDomain(positivePossibleX+cc.c.x) && lc.isWithinDomain(negativePossibleX+cc.c.x) {
		return &Vector2{negativePossibleX + cc.c.x, negativePossibleX*s + h + cc.c.y}
	}

	if lc.isWithinDomain(positivePossibleX+cc.c.x) && !lc.isWithinDomain(negativePossibleX+cc.c.x) {
		return &Vector2{positivePossibleX + cc.c.x, positivePossibleX*s + h + cc.c.y}
	}

	fmt.Println("no collision")
	return nil
}

func (line *LineCollider) getSlope() float32 {
	return (line.y2 - line.y1) / (line.x2 - line.x1)
}

func (line *LineCollider) getIntercept() float32 {
	return line.y1 - line.x1*line.getSlope()
}

func (line *LineCollider) getRange() (float32, float32) {
	if line.y1 > line.y2 {
		return line.y2, line.y1
	}
	return line.y1, line.y2
}

// returns lower ==> upper
func (line *LineCollider) getDomain() (float32, float32) {

	if line.x1 > line.x2 {
		return line.x2, line.x1
	}
	return line.x1, line.x2

}

func (line *LineCollider) isWithinDomain(possibleX float32) bool {
	lower, upper := line.getDomain()

	if possibleX < lower {
		return false
	}

	if possibleX > upper {
		return false
	}

	return true
}

func newLineCollider(X1 float32, Y1 float32, X2 float32, Y2 float32) *LineCollider {
	returnthing := LineCollider{X1, Y1, X2, Y2}
	return &returnthing
}

type CircleCollider struct {
	c Vector2
	r float32
}

type Vector2 struct {
	x float32
	y float32
}

type Game struct {
	units       []*Unit
	walls       []*Wall
	cameraX     float32
	cameraY     float32
	cameraScale float32
}

type Unit struct {
	positionX   float32
	positionY   float32
	movement    float32
	health      int
	clickRadius float32
	team        int
	isSelected  bool
	commands    []*Command
}

type LineCollider struct {
	x1 float32
	y1 float32

	x2 float32
	y2 float32
}

type Wall struct {
	startX float32
	startY float32
	endX   float32
	endY   float32
}

type Command struct {
	targetX float32
	targetY float32
}

type EndError struct {
	When time.Time
	What string
}
