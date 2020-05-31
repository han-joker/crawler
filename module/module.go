package module

type Module interface {
	ID() MID
	Addr() string
	Score() uint64
	SetScore(score uint64)
	ScoreCalculator() CalculatorScore
	CalledCount() uint64
	AcceptedCount() uint64
	CompletedCount() uint64
	HandingNumber() uint64
	Counts() Counts
	Summary() SummaryStruct
}
