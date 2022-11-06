package main

import (
	"gitlab.com/greg198584/gridclient/algo"
	"gitlab.com/greg198584/gridclient/tools"
	"net/http"
	"time"
)

func main() {
	current, err := algo.NewAlgo("demo_b")
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	if err != nil {
		panic(err)
	}
	taille := current.InfosGrid.Taille - 1
	nbrJump := current.Psi.Programme.Position.SecteurID
	if nbrJump > 0 {
		nbrJump = taille - current.Psi.Programme.Position.SecteurID
	} else {
		nbrJump = taille
	}
	current.Move(taille - 1)
	current.JumpDown(nbrJump)
	status := true
	for status {
		current.GetStatusGrid()
		tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
		tools.PrintInfosGrille(current.InfosGrid)
		for j := taille; j > 0; j-- {
			time.Sleep(500 * time.Millisecond)
			if ok, _ := current.Move(j); !ok {
				status = false
				break
			}
			current.CheckAttack()
			tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
			tools.PrintInfosGrille(current.InfosGrid)
			ok, programmes := current.GetProgramme()
			if !ok {
				status = false
				break
			}
			for _, pid := range programmes {
				//tools.Success(fmt.Sprintf("programme trouver [%s]", pid))
				//SendToDiscord(pid)
				for i := 0; i < 11; i++ {
					current.Attack(pid)
					current.CheckAttack()
					if current.Psi.LockProgramme[pid].Status == false {
						break
					}
				}
			}
		}
		if ok, _ := current.Move(taille); !ok {
			break
		}
		current.CheckAttack()
		tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
		tools.PrintInfosGrille(current.InfosGrid)
		time.Sleep(500 * time.Millisecond)
		if ok, _ := current.JumpUp(1); !ok {
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
