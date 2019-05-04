package simulation

type ThreeSumConfig struct {
	Length   int
	SumValue int
}

func NewThreeSumConfig(length int, sumValue int) ThreeSumConfig {
	return ThreeSumConfig{
		Length:   length,
		SumValue: sumValue,
	}
}
