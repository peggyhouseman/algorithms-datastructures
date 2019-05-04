package simulation

type PercolationRunConfig struct {
	NumberGrids int
	GridType    UnionFindDsType
	Height      int
	Width       int
}

func NewPercolationRunConfigs(numGrids int, height int, width int) []PercolationRunConfig {
	if numGrids <= 0 {
		return []PercolationRunConfig{}
	}

	return []PercolationRunConfig{
		PercolationRunConfig{
			NumberGrids: numGrids,
			GridType:    ArrayQuickFind,
			Height:      height,
			Width:       width,
		},
		PercolationRunConfig{
			NumberGrids: numGrids,
			GridType:    ArrayQuickUnion,
			Height:      height,
			Width:       width,
		},
	}
}
