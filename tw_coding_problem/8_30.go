package tw_coding_problem

import "errors"

func Solution(A []int, X int, Y int, Z int) int {
	TimePassed := -1
	station := InitGasStation(X, Y, Z)

	for i := 0; i < len(A); i++ {
		WasPumpAvailable, err := station.TryToPump(A[i])
		if err != nil {
			// all pumps are clear and cant fuel
			// queue is blocked
			return TimePassed
		} else if !WasPumpAvailable {
			// retry car when pump is clear
			station.PassOneSecondInAllStation()
			TimePassed++
			i = i - 1
		}
	}

	return TimePassed
}

type GasStation struct {
	X FuelStation
	Y FuelStation
	Z FuelStation
}

type FuelStation struct {
	GasReserve  int
	FillRequest int
	TimePassed  int
	InUse       bool
}

func InitGasStation(X int, Y int, Z int) GasStation {
	return GasStation{
		X: FuelStation{GasReserve: X},
		Y: FuelStation{GasReserve: Y},
		Z: FuelStation{GasReserve: Z},
	}
}

func (g *GasStation) TryToPump(pumpAmount int) (bool, error) {
	if g.PumpAtX(pumpAmount) {
		return true, nil
	} else if g.PumpAtY(pumpAmount) {
		return true, nil
	} else if g.PumpAtZ(pumpAmount) {
		return true, nil
	} else if g.ArePumpsClear() {
		return false, errors.New("gas station cant handle gas request")
	}
	return false, nil

}

func (g *GasStation) PumpAtX(pumpAmount int) bool {
	return attemptPumpAtStation(pumpAmount, &g.X)
}

func (g *GasStation) PumpAtY(pumpAmount int) bool {
	return attemptPumpAtStation(pumpAmount, &g.Y)

}

func (g *GasStation) PumpAtZ(pumpAmount int) bool {
	return attemptPumpAtStation(pumpAmount, &g.Z)

}

func attemptPumpAtStation(gasAmount int, X *FuelStation) bool {
	// handles if a station can be used for Gas Request
	if X.GasReserve-gasAmount < 0 || X.InUse == true {
		return false
	}
	X.InUse = true
	X.FillRequest = gasAmount
	X.GasReserve = X.GasReserve - gasAmount
	return true
}

func (g *GasStation) PassOneSecondInAllStation() {
	passOneSecondInStation(&g.X)
	passOneSecondInStation(&g.Y)
	passOneSecondInStation(&g.Z)
}

func passOneSecondInStation(Z *FuelStation) {
	// Handle logic for passing 1 second in station and if station is now available
	if Z.InUse {
		Z.TimePassed++
		if Z.TimePassed > Z.FillRequest {
			Z.InUse = false
			Z.FillRequest = 0
			Z.TimePassed = 0
		}
	}
}

func (g *GasStation) ArePumpsClear() bool {
	if g.X.InUse {
		return false
	}
	if g.Y.InUse {
		return false
	}
	if g.Z.InUse {
		return false
	}
	return true
}
