package renderer

func gravity(g *Game, p *PlacedCircle) {
	var fy float32 = 0.0

	fy += p.M * g.Gravity

	dy := p.vy*g.timeStep + (0.5 * p.ay * g.timeStep * g.timeStep)

	p.Y = dy * 100
	new_ay := fy / p.M
	avg_ay := 0.5 * (new_ay + p.ay)
	p.vy = avg_ay * g.timeStep

	if p.Y+p.Radius > g.Height && p.vy > 0 {
		p.Y = g.Height - p.Radius
	}

}
