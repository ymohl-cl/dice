package loto

import (
	"errors"
	"fmt"

	"github.com/ymohl-cl/dice/pkg/dice"
)

// Loto data interface
type Loto interface {
	Prediction(opt Option) (rest int64, err error)
	resetPrediction() error
	History() History
	String() string
}

type defaultLoto struct {
	history History
	balls   dice.Dice
	joker   dice.Dice
}

type variant1Loto struct {
	history History
	ball1   dice.Dice
	ball2   dice.Dice
	ball3   dice.Dice
	ball4   dice.Dice
	ball5   dice.Dice
	joker   dice.Dice
}

// New loto instance with the loading of loto hitory and return the default instance
func New() (Loto, error) {
	var l defaultLoto
	var err error

	if l.history, err = NewHistory(); err != nil {
		return nil, err
	}

	return &l, nil
}

func (l *defaultLoto) resetPrediction() error {
	var err error

	if l.balls, err = dice.New(49); err != nil {
		return err
	}
	if l.joker, err = dice.New(10); err != nil {
		return err
	}
	return nil
}

// Prediction analyse the history describe by opt to get the balls (include joker) score.
// On the defaultLoto all balls are groups on the same dice.
func (l *defaultLoto) Prediction(opt Option) (int64, error) {
	var indexOut int
	var err error
	var lenSection int64

	if err = l.resetPrediction(); err != nil {
		return 0, err
	}

	if opt.Index > len(l.history.RecentDraws) {
		return 0, errors.New("invalid index to start exploration draw")
	}
	lenSection = int64(len(l.history.RecentDraws[opt.Index:]))
	if opt.NDraws > lenSection {
		opt.NDraws = lenSection
	}
	indexOut = opt.Index + int(opt.NDraws)
	if opt.NDraws == 0 {
		indexOut = opt.Index + int(lenSection)
		opt.NDraws = lenSection
	}

	for _, d := range l.history.RecentDraws[opt.Index:indexOut] {
		l.balls.SetThrow(d.B1)
		l.balls.SetThrow(d.B2)
		l.balls.SetThrow(d.B3)
		l.balls.SetThrow(d.B4)
		l.balls.SetThrow(d.B5)
		l.joker.SetThrow(d.Joker)
	}
	rest := lenSection - opt.NDraws
	if !opt.Old {
		// dont use the optionnal old throws
		return rest, nil
	}

	for _, d := range l.history.OldDraws {
		l.balls.SetThrow(d.B1)
		l.balls.SetThrow(d.B2)
		l.balls.SetThrow(d.B3)
		l.balls.SetThrow(d.B4)
		l.balls.SetThrow(d.B5)
		l.joker.SetThrow(d.Joker)
		if opt.Old6thBall {
			// optionnal 6th ball (value 1 to 49)
			l.balls.SetThrow(d.B6)
		}
		if opt.OldLuckyBall {
			// optionnal lucky ball (value 1 to 49)
			l.balls.SetThrow(d.Joker)
		}
	}
	return rest, nil
}

// History getter
func (l defaultLoto) History() History {
	return l.history
}

// String printable
func (l defaultLoto) String() string {
	var str string

	str = fmt.Sprintf("the default order prediction in follow about the picks balls:\n")
	faces, picks := l.balls.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}
	str += fmt.Sprintf("the default order prediction in follow about the pick joker:\n")
	faces, picks = l.joker.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}

	return str
}

// NewVariant1 analyse ball by ball separetly with a strict pick order
func NewVariant1() (Loto, error) {
	var l variant1Loto
	var err error

	if l.history, err = NewHistory(); err != nil {
		return nil, err
	}

	return &l, nil
}

func (l *variant1Loto) resetPrediction() error {
	var err error

	if l.ball1, err = dice.New(49); err != nil {
		return err
	}
	if l.ball2, err = dice.New(49); err != nil {
		return err
	}
	if l.ball3, err = dice.New(49); err != nil {
		return err
	}
	if l.ball4, err = dice.New(49); err != nil {
		return err
	}
	if l.ball5, err = dice.New(49); err != nil {
		return err
	}
	if l.joker, err = dice.New(10); err != nil {
		return err
	}
	return nil
}

// Prediction analyse the history describe by opt to get the balls (include joker) score.
// On the variant1Loto all balls are separetly on the differs dices.
func (l *variant1Loto) Prediction(opt Option) (int64, error) {
	var indexOut int
	var err error
	var lenSection int64

	if err = l.resetPrediction(); err != nil {
		return 0, err
	}

	if opt.Index > len(l.history.RecentDraws) {
		return 0, errors.New("invalid index to start exploration draw")
	}
	lenSection = int64(len(l.history.RecentDraws[opt.Index:]))
	if opt.NDraws > lenSection {
		opt.NDraws = lenSection
	}
	indexOut = opt.Index + int(opt.NDraws)
	if opt.NDraws == 0 {
		indexOut = opt.Index + int(lenSection)
		opt.NDraws = lenSection
	}

	for _, d := range l.history.RecentDraws[opt.Index:indexOut] {
		l.ball1.SetThrow(d.B1)
		l.ball2.SetThrow(d.B2)
		l.ball3.SetThrow(d.B3)
		l.ball4.SetThrow(d.B4)
		l.ball5.SetThrow(d.B5)
		l.joker.SetThrow(d.Joker)
	}
	rest := lenSection - opt.NDraws
	if !opt.Old {
		// dont use the optionnal old throws
		return rest, nil
	}

	for _, d := range l.history.OldDraws {
		l.ball1.SetThrow(d.B1)
		l.ball2.SetThrow(d.B2)
		l.ball3.SetThrow(d.B3)
		l.ball4.SetThrow(d.B4)
		l.ball5.SetThrow(d.B5)
		l.joker.SetThrow(d.Joker)
		if opt.Old6thBall {
			// optionnal 6th ball (value 1 to 49)
			l.ball1.SetThrow(d.B6)
			l.ball2.SetThrow(d.B6)
			l.ball3.SetThrow(d.B6)
			l.ball4.SetThrow(d.B6)
			l.ball5.SetThrow(d.B6)
		}
		if opt.OldLuckyBall {
			// optionnal lucky ball (value 1 to 49)
			l.ball1.SetThrow(d.Joker)
			l.ball2.SetThrow(d.Joker)
			l.ball3.SetThrow(d.Joker)
			l.ball4.SetThrow(d.Joker)
			l.ball5.SetThrow(d.Joker)
		}
	}
	return rest, nil
}

// History getter
func (l *variant1Loto) History() History {
	return l.history
}

// String printable
func (l variant1Loto) String() string {
	var str string

	str = fmt.Sprintf("the variant 1 order prediction in follow about the picks ball_1:\n")
	faces, picks := l.ball1.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}
	str += fmt.Sprintf("the variant 1 order prediction in follow about the picks ball_2:\n")
	faces, picks = l.ball2.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}
	str += fmt.Sprintf("the variant 1 order prediction in follow about the picks ball_3:\n")
	faces, picks = l.ball3.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}
	str += fmt.Sprintf("the variant 1 order prediction in follow about the picks ball_4:\n")
	faces, picks = l.ball4.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}
	str += fmt.Sprintf("the variant 1 order prediction in follow about the picks ball_5:\n")
	faces, picks = l.ball5.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}

	str += fmt.Sprintf("the variant 1 order prediction in follow about the pick joker:\n")
	faces, picks = l.joker.Order()
	for i, pick := range picks {
		str += fmt.Sprintf("\tface number %d pick %d time(s)\n", faces[i], pick)
	}

	return str
}
