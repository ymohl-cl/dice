package loto

// Loto data interface
type Loto interface {
	GeneralPrediction(opt Option) (GeneralPredict, error)
	BallsPrediction(opt Option) (BallsPredict, error)
}

type loto struct {
	history History
}

// New loto instance with the loading of loto hitory
func New() (Loto, error) {
	var l loto
	var err error

	if l.history, err = NewHistory(); err != nil {
		return nil, err
	}
	return &l, nil
}

func (l loto) GeneralPrediction(opt Option) (GeneralPredict, error) {
	var predict GeneralPredict
	var nb int64
	var err error

	if predict, err = NewGeneralPredict(); err != nil {
		return GeneralPredict{}, err
	}

	nb = opt.NumberDraws
	if opt.NumberDraws == 0 {
		nb = int64(len(l.history.RecentDraws))
	}
	for _, d := range l.history.RecentDraws[:nb] {
		predict.Balls.SetThrow(d.B1)
		predict.Balls.SetThrow(d.B2)
		predict.Balls.SetThrow(d.B3)
		predict.Balls.SetThrow(d.B4)
		predict.Balls.SetThrow(d.B5)
		predict.Joker.SetThrow(d.Joker)
	}
	if !opt.Old {
		// dont use the optionnal old throws
		return predict, nil
	}

	for _, d := range l.history.OldDraws {
		predict.Balls.SetThrow(d.B1)
		predict.Balls.SetThrow(d.B2)
		predict.Balls.SetThrow(d.B3)
		predict.Balls.SetThrow(d.B4)
		predict.Balls.SetThrow(d.B5)
		if opt.Old6thBall {
			// optionnal 6th ball (value 1 to 49)
			predict.Balls.SetThrow(d.B6)
		}
		if opt.OldLuckyBall {
			// optionnal lucky ball (value 1 to 49)
			predict.Balls.SetThrow(d.Joker)
		}
	}
	return predict, nil
}

func (l loto) BallsPrediction(opt Option) (BallsPredict, error) {
	var err error
	var predict BallsPredict

	if predict, err = NewBallsPredict(); err != nil {
		return BallsPredict{}, err
	}
	for _, d := range l.history.RecentDraws {
		predict.Ball1.SetThrow(d.B1)
		predict.Ball2.SetThrow(d.B2)
		predict.Ball3.SetThrow(d.B3)
		predict.Ball4.SetThrow(d.B4)
		predict.Ball5.SetThrow(d.B5)
		predict.Joker.SetThrow(d.Joker)
	}
	if !opt.Old {
		// dont use the optionnal old throws
		return predict, nil
	}

	for _, d := range l.history.OldDraws {
		predict.Ball1.SetThrow(d.B1)
		predict.Ball2.SetThrow(d.B2)
		predict.Ball3.SetThrow(d.B3)
		predict.Ball4.SetThrow(d.B4)
		predict.Ball5.SetThrow(d.B5)
		if opt.Old6thBall {
			// optionnal 6th ball (value 1 to 49)
			predict.Ball1.SetThrow(d.B6)
			predict.Ball2.SetThrow(d.B6)
			predict.Ball3.SetThrow(d.B6)
			predict.Ball4.SetThrow(d.B6)
			predict.Ball5.SetThrow(d.B6)
		}
		if opt.OldLuckyBall {
			// optionnal lucky ball (value 1 to 49)
			predict.Ball1.SetThrow(d.Joker)
			predict.Ball2.SetThrow(d.Joker)
			predict.Ball3.SetThrow(d.Joker)
			predict.Ball4.SetThrow(d.Joker)
			predict.Ball5.SetThrow(d.Joker)
		}
	}
	return predict, nil
}
