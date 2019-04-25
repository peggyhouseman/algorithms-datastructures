package percolation

type RunConfig struct {
	NumberGrids int
	GridType    UnionFindDsType
	Height      int
	Width       int
}

func NewRunConfigs(numGrids int, height int, width int) []RunConfig {
	if numGrids <= 0 {
		return []RunConfig{}
	}

	return []RunConfig{
		RunConfig{
			NumberGrids: numGrids,
			GridType:    ArrayQuickFind,
			Height:      height,
			Width:       width,
		},
		RunConfig{
			NumberGrids: numGrids,
			GridType:    ArrayQuickUnion,
			Height:      height,
			Width:       width,
		},
	}
}
