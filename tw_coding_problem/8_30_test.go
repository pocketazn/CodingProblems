package tw_coding_problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	type TestCase struct {
		Name             string
		Sample           []int
		TankX            int
		TankY            int
		TankZ            int
		ExpectedSolution int
	}

	Cases := []TestCase{
		{
			Name:             "Sample A",
			Sample:           []int{2, 8, 4, 3, 2},
			TankX:            7,
			TankY:            11,
			TankZ:            3,
			ExpectedSolution: 8,
		},
		{
			Name:             "Queue is Blocked by [2]",
			Sample:           []int{2, 13, 4, 3, 2},
			TankX:            7,
			TankY:            11,
			TankZ:            3,
			ExpectedSolution: 2,
		},
		{
			Name:             "Cant Pump",
			Sample:           []int{5},
			TankX:            4,
			TankY:            0,
			TankZ:            3,
			ExpectedSolution: -1,
		},
		{
			Name:             "Cant Pump all empty",
			Sample:           []int{5},
			TankX:            0,
			TankY:            0,
			TankZ:            0,
			ExpectedSolution: -1,
		},
		{
			Name:             "Cant Pump all empty",
			Sample:           []int{0},
			TankX:            0,
			TankY:            0,
			TankZ:            0,
			ExpectedSolution: -1,
		},
		{
			Name:             "No Cars",
			Sample:           []int{},
			TankX:            5,
			TankY:            5,
			TankZ:            5,
			ExpectedSolution: -1,
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			sol := Solution(tt.Sample, tt.TankX, tt.TankY, tt.TankZ)
			assert.Equal(t, tt.ExpectedSolution, sol)
		})
	}
}
