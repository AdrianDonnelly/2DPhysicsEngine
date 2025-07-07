package renderer

import "fmt"

func gravity(g *Game, p *PlacedCircle) {
	var fy float32 = 0.0
	fy += p.M * 9.82
	fmt.Println(p.M)
	fmt.Println("Force ", fy)
	fy += -1 * 0.5 * g.rho * p.C_d * p.A * p.vy * p.vy
	fmt.Println("Force after ", fy)

	dy := p.vy*g.timeStep + (0.5 * p.ay * g.timeStep * g.timeStep)
	p.Y = dy * 1000

	new_ay := fy / p.M
	avg_ay := 0.5 * (new_ay + p.ay)
	p.vy = avg_ay * g.timeStep

	if p.Y+p.Radius > g.Height && p.vy > 0 {
		p.vy *= p.e
		p.Y = g.Height - p.Radius
	}

}
