package loto

// Option to adjust the data prediction
// Index to start in the history
// NDraw to the number draw usage in the prediction
// Old to get the old loto history
type Option struct {
	Index        int
	NDraws       int64
	Old          bool
	Old6thBall   bool
	OldLuckyBall bool
}
