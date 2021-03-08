package dice

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// Dice describe a dice and the history
type Dice interface {
	Throw() int32
	SetThrow(int32)
	History() []int32
	Result(face int32) int64
	NbThrow() int64
	WeaklestFaces() []int32
}

type dice struct {
	nbThrow     int64
	facesResult map[int32]int64
	history     []int32
	nbFace      int32
	random      *rand.Rand
}

// New instance dice returnment
func New(nbFace int32) (Dice, error) {
	if nbFace <= 0 {
		return nil, errors.New("The dice can't have a zero face or negative face")
	}

	d := dice{
		facesResult: make(map[int32]int64),
		nbFace:      nbFace,
		random:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for i := 0; i < int(nbFace); i++ {
		d.facesResult[int32(i)+1] = 0
	}
	return &d, nil
}

// Throw dice simulation. Ret get + 1 value because a dice don't have a face with zero value.
func (d *dice) Throw() int32 {
	d.nbThrow++
	ret := d.random.Int31n(d.nbFace)
	ret++

	d.facesResult[ret]++
	d.history = append(d.history, ret)
	return ret
}

// SetThrow dice manual setting
func (d *dice) SetThrow(ret int32) {
	d.nbThrow++

	d.facesResult[ret]++
	d.history = append(d.history, ret)
	return
}

// History (getter)
func (d dice) History() []int32 {
	return d.history
}

// Result return the number time result to face number parameter
func (d dice) Result(face int32) int64 {
	return d.facesResult[face]
}

// NbThrow getter value
func (d dice) NbThrow() int64 {
	return d.nbThrow
}

// WeaklestFaces return the weakles faces scores
func (d dice) WeaklestFaces() []int32 {
	var values []int32
	var min int64

	//	fmt.Printf("RESULT: ---------------------------\n\n")
	min = math.MaxInt64
	for face, v := range d.facesResult {
		//		fmt.Printf("face %d with number result %d\n", face, v)
		if v < min {
			min = v
			values = append([]int32{}, face)
		} else if v == min {
			values = append(values, face)
		}
	}
	return values
}
