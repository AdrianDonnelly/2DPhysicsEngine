package renderer

import "fmt"

func gravity(g *Game, p *PlacedCircle) {
	var YForce float32 = 0.0
	YForce += p.Mass * 9.82

	YForce += -1 * 0.5 * g.AirDens * p.CDrag * p.Area * p.Yvel * p.Yvel
	fmt.Println(p.Y)
	dy := p.Yvel*g.timeStep + (0.5 * p.Yaccel * g.timeStep * g.timeStep)
	p.Y += dy * 100

	new_Yaccel := YForce / p.Mass
	avg_Yaccel := 0.5 * (new_Yaccel + p.Yaccel)
	p.Yvel += avg_Yaccel * g.timeStep
	p.Yaccel += new_Yaccel
	fmt.Printf("dy=%.4f, Yvel=%.4f, Yaccel=%.4f, Y=%.2f\n", dy, p.Yvel, p.Yaccel, p.Y)

	if p.Y+p.Radius > g.Height && p.Yvel > 0 {
		p.Yvel *= p.bounce
		p.Y = g.Height - p.Radius
	}

}
