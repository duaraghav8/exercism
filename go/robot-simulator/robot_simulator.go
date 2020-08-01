package robot

const (
	N Dir = 0
	S     = 1
	E     = 2
	W     = 3
)

func (d Dir) String() string {
	switch d {
	case N:
		return "North"
	case S:
		return "South"
	case E:
		return "East"
	default:
		return "West"
	}
}

func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case S:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = S
	default:
		Step1Robot.Dir = N
	}
}

func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case S:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = N
	default:
		Step1Robot.Dir = S
	}
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
	case S:
		Step1Robot.Y -= 1
	case E:
		Step1Robot.X += 1
	default:
		Step1Robot.X -= 1
	}
}
