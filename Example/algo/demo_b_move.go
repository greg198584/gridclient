package main

import (
	"encoding/json"
	"fmt"
	"github.com/greg198584/gridclient/algo"
	"github.com/greg198584/gridclient/structure"
	"github.com/greg198584/gridclient/tools"
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
			time.Sleep(500 * time.Millisecond)
			if ok, _ := current.Move(j); !ok {
				if current.StatusCode == http.StatusUnauthorized {
					status = false
					break
				}
			}
			if scanOK, res, _ := current.Scan(); !scanOK {
				tools.Fail("faille scan")
			} else {
				var zoneInfos structure.ZoneInfos
				_ := json.Unmarshal(res, &zoneInfos)
				for _, programme := range zoneInfos.Programmes {
					if programme.Status {
						tools.Info(fmt.Sprintf("programme trouver [%d]"))
					}
				}
			}
			current.GetStatusGrid()
			tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
			tools.PrintInfosGrille(current.InfosGrid)
		}
		if ok, _ := current.Move(0); !ok {
			continue
		}
		tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
		tools.PrintInfosGrille(current.InfosGrid)
		time.Sleep(500 * time.Millisecond)
		if ok, _ := current.JumpDown(1); !ok {
			if current.StatusCode == http.StatusUnauthorized {
				status = false
				break
			} else {
				current.Unset()
				break
			}
		}
	}
}
