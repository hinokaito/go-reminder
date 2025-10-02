package main

import "fmt"

/*
	戦士を選定するプログラムにミスがあり、エラーが出てしまいました。修正してください。
	なお、if文等、未学習のものがありますが、そこにミスは無いので触れないでください。
	checkWarrior関数は、戦士に適していればTrue、適していなければFalseを返す関数です。

	期待する結果:
	True
	False
	False

*/

const (
	heightLowerLimit = 170 // 戦士の身長下限 変えないでください
	heightUpperLimit = 200 // 戦士の身長上限 変えないでください
	humanHeight      = 0
)

func main() {
	humanHeight := 175
	fmt.Print("審査の結果、%vでした。", checkWarrior(humanHeight))

	humanHeight := 160
	fmt.Print("審査の結果、%vでした。", checkWarrior(humanHeight))

	humanHeight := 201
	fmt.Print("審査の結果、%vでした。", checkWarrior(humanHeight))
}

// 変えないでください
func checkWarrior(height int) bool {
	if height >= heightLowerLimit && height <= heightUpperLimit {
		return true
	}
	return false
}
