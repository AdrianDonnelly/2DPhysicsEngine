package renderer

func gravity(g *Game, p *placedRect) {
	var YForce float32 = 0.0

	YForce += p.Mass * 9.82

	if p.Yaccel > 100 {
		p.Yaccel = 100
	}

	YForce += -1 * 0.5 * g.AirDens * p.CDrag * p.Area * p.Yvel * p.Yvel

	dy := p.Yvel*g.timeStep + (0.5 * p.Yaccel * g.timeStep * g.timeStep)
	p.Y += dy * 100

	newYaccel := YForce / p.Mass
	avgYaccel := 0.5 * (newYaccel + p.Yaccel)
	p.Yvel += avgYaccel * g.timeStep
	p.Yaccel += newYaccel

	//collisions
	if p.Y+p.Height > g.Height && p.Yvel > 0 {
		p.Yvel *= p.bounce
		p.Y = g.Height - p.Height
	}
}
