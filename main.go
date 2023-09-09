package main

import "fmt"

type SafetyPlacer interface {
	placeSafeties()
}
type RockClimber struct {
	rocksClimbed int
	sp           SafetyPlacer
}

func (rc *RockClimber) newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed += 1
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

type IceSafetyPlacer struct {
	//we can add db api call etc. here
}

func (sp *IceSafetyPlacer) placeSafeties() {
	fmt.Println("placing my ice safety")
}

type NoSafetyPlacer struct {
}

func (sp *NoSafetyPlacer) placeSafeties() {
	fmt.Println("cant place safety here")
}

func main() {

	rc := RockClimber{sp: &NoSafetyPlacer{}}
	for i := 0; i < 11; i++ {
		rc.climbRock()
		fmt.Println(rc.rocksClimbed)
	}
}
