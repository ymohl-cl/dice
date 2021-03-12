package loto

// score point to match ball or joker
var (
	JokerValue int64 = 100
	BallValue  int64 = 10
)

// Score is a status resolve to the prediction
// Draw is the future draw which try to find with the current balls and joker
type Score struct {
	Value int64
	Draw  Draw
	Balls []int32
	Joker int32
}

// NewScore instanciate and evaluate score
func NewScore(d Draw, balls []int32, joker int32) Score {
	var sc int64

	for _, ball := range balls {
		switch ball {
		case d.B1, d.B2, d.B3, d.B4, d.B5:
			sc += BallValue
		}
	}
	if joker == d.Joker {
		sc += JokerValue
	}
	return Score{
		Value: sc,
		Draw:  d,
		Balls: balls,
		Joker: joker,
	}
}
