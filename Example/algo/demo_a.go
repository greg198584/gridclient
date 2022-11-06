package main

import (
	"gitlab.com/greg198584/gridclient/algo"
	"gitlab.com/greg198584/gridclient/tools"
	"net/http"
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
	count := current.InfosGrid.Taille
	nbrJump := current.Psi.Programme.Position.SecteurID
	current.Move(0)
	current.JumpUp(nbrJump)
	status := true
	for status {
		current.GetStatusGrid()
		tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
		tools.PrintInfosGrille(current.InfosGrid)
		for j := 0; j <= count; j++ {
			time.Sleep(250 * time.Millisecond)
			if ok, _ := current.Move(j); !ok {
				if current.StatusCode == http.StatusUnauthorized {
					status = false
					break
				}
			}
			current.CheckAttack()
			current.GetStatusGrid()
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
		if ok, _ := current.Move(0); !ok {
			continue
		}
		current.CheckAttack()
		tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
		tools.PrintInfosGrille(current.InfosGrid)
		time.Sleep(250 * time.Millisecond)
		if ok, _ := current.JumpDown(1); !ok {
			if current.StatusCode == http.StatusUnauthorized {
				status = false
				break
			} else {
				current.Unset()
				break
			}
		}
		current.CheckAttack()
	}
}
