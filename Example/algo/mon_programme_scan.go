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
	current, err := algo.NewAlgo("test")
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
			time.Sleep(2 * time.Second)
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
			if scanOK, scanRes, _ := current.Scan(); !scanOK {
				tools.Fail("erreur scan")
			} else {
				var zoneInfos structure.ZoneInfos
				err := json.Unmarshal(scanRes, &zoneInfos)
				if err != nil {
					tools.Fail(err.Error())
				} else {
					for _, cellule := range zoneInfos.Cellules {
						if cellule.Status {
							tools.Success(fmt.Sprintf("zone [%d] - cellule [%d] etat [%t] - data presente ou etat true", zoneInfos.ID, cellule.ID, cellule.Status))
							if exploreOK, exploreRes, _ := current.Explore(cellule.ID); !exploreOK {
								tools.Fail("erreur explore")
							} else {
								var datas map[int]structure.CelluleData
								err := json.Unmarshal(exploreRes, &datas)
								if err != nil {
									tools.Fail(err.Error())
								} else {
									for _, data := range datas {
										if data.IsFlag {
											tools.Success(fmt.Sprintf("FLAG TROUVER - Cellule [%d] - index [%d]", cellule.ID, data.ID))
										}
									}
								}
							}
						} else {
							//tools.Warning(fmt.Sprintf("cellule [%d] etat [%t] - aucune data ou etat false", cellule.ID, cellule.Status))
						}
					}
				}
			}
		}
		if ok, _ := current.Move(0); !ok {
			continue
		}
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
