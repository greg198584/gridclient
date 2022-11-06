package main

import (
	"github.com/greg198584/gridclient/algo"
	"github.com/greg198584/gridclient/tools"
	"time"
)

func main() {
	current, err := algo.NewAlgo("demo_a")
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	if err != nil {
		panic(err)
	}
	status := true
	for status {
		time.Sleep(1000 * time.Millisecond)
		current.CheckAttack()
		tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
		tools.PrintInfosGrille(current.InfosGrid)
		ok, programmes := current.GetProgramme()
		if !ok {
			status = false
			break
		}
		for _, pid := range programmes {
			for i := 0; i < 11; i++ {
				current.Attack(pid)
				current.CheckAttack()
				if current.Psi.LockProgramme[pid].Status == false {
					break
				}
			}
		}
	}
}
