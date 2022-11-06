package algo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/greg198584/gridclient/api"
	"github.com/greg198584/gridclient/structure"
	"github.com/greg198584/gridclient/tools"
	"net/http"
	"os"
	"time"
)

const (
	TIME_MILLISECONDE = 500
)

type Algo struct {
	Name       string
	ID         string
	Pc         structure.ProgrammeContainer
	InfosGrid  structure.GridInfos
	Psi        structure.ProgrammeStatusInfos
	StatusCode int
}

func _LoadProgramme(name string) (psi structure.ProgrammeStatusInfos, pc structure.ProgrammeContainer, err error) {
	pc, err = _GetProgrammeFile(name)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(pc)
	res, statusCode, err := api.RequestApi(
		"POST",
		fmt.Sprintf("%s/%s", api.API_URL, api.ROUTE_LOAD_PROGRAMME),
		reqBodyBytes.Bytes(),
	)
	if statusCode == http.StatusCreated || statusCode == http.StatusOK {
		err = json.Unmarshal(res, &psi)
	} else {
		pc, err = _CreateProgramme(name)
	}
	return
}
func _CreateProgramme(name string) (programme structure.ProgrammeContainer, err error) {
	tools.Title(fmt.Sprintf("création programme [%s]", name))
	if _IsExistFile(name) == false {
		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s", api.API_URL, api.ROUTE_NEW_PROGRAMME, name),
			nil,
		)
		if err != nil {
			return programme, err
		}
		if statusCode == http.StatusCreated {
			err = json.Unmarshal(res, &programme)
			tools.CreateJsonFile(fmt.Sprintf("%s.json", name), programme)
		} else {
			err = errors.New("erreur creation programme")
			return programme, err
		}
	} else {
		tools.Warning(fmt.Sprintf("programme file exist"))
	}
	return programme, err
}
func _IsExistFile(name string) bool {
	filename := fmt.Sprintf("%s.json", name)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func _GetProgrammeFile(name string) (pc structure.ProgrammeContainer, err error) {
	file, err := tools.GetJsonFile(fmt.Sprintf("%s.json", name))
	if err != nil {
		return pc, err
	}
	err = json.Unmarshal(file, &pc)
	if err != nil {
		return pc, err
	}
	return
}
func NewAlgo(name string) (algo *Algo, err error) {
	tools.Title(fmt.Sprintf("chargement programme [%s]", name))
	psi, pc, err := _LoadProgramme(name)
	algo = &Algo{
		Name: name,
		Psi:  psi,
		Pc:   pc,
	}
	if algo.Psi.Programme.ID == "" {
		algo.GetInfosProgramme()
	}
	algo.ID = algo.Psi.Programme.ID
	return algo, err
}
func (a *Algo) GetInfosProgramme() (ok bool, err error) {
	tools.Title(fmt.Sprintf("infos programme [%s]", a.Name))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_STATUS_PROGRAMME, a.Pc.ID, a.Pc.SecretID),
		nil,
	)
	a.StatusCode = statusCode
	if err != nil || statusCode != http.StatusOK {
		return false, err
	}
	a.Psi = structure.ProgrammeStatusInfos{}
	err = json.Unmarshal(res, &a.Psi)
	return true, err
}
func (a *Algo) Delete() (ok bool, err error) {
	tools.Title(fmt.Sprintf("suppression programme [%s]", a.Name))
	_, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_UNSET_PROGRAMME, a.Pc.ID, a.Pc.SecretID),
		nil,
	)
	a.StatusCode = statusCode
	if err != nil || statusCode != http.StatusOK {
		return false, err
	}
	return true, err
}
func (a *Algo) JumpUp(valeur int) (ok bool, err error) {
	tools.Title(fmt.Sprintf("Programme [%s] JumpUP [%d]", a.Name, valeur))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s/%d", api.API_URL, api.ROUTE_JUMPUP_PROGRAMME, a.Pc.ID, a.Pc.SecretID, valeur),
		nil,
	)
	a.StatusCode = statusCode
	if err != nil || statusCode != http.StatusOK {
		return false, err
	}
	a.Psi = structure.ProgrammeStatusInfos{}
	err = json.Unmarshal(res, &a.Psi)
	return true, err
}
func (a *Algo) JumpDown(valeur int) (ok bool, err error) {
	tools.Title(fmt.Sprintf("Programme [%s] JumpDown [%d]", a.Name, valeur))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s/%d", api.API_URL, api.ROUTE_JUMPDOWN_PROGRAMME, a.Pc.ID, a.Pc.SecretID, valeur),
		nil,
	)
	a.StatusCode = statusCode
	if err != nil || statusCode != http.StatusOK {
		return false, err
	}
	a.Psi = structure.ProgrammeStatusInfos{}
	err = json.Unmarshal(res, &a.Psi)
	return true, err
}
func (a *Algo) Move(valeur int) (ok bool, err error) {
	tools.Title(fmt.Sprintf("Programme [%s] Move to Zone [%d]", a.Name, valeur))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s/%d", api.API_URL, api.ROUTE_MOVE_PROGRAMME, a.Pc.ID, a.Pc.SecretID, valeur),
		nil,
	)
	if err != nil || statusCode != http.StatusOK {
		return false, err
	}
	a.Psi = structure.ProgrammeStatusInfos{}
	err = json.Unmarshal(res, &a.Psi)
	return true, err
}
func (a *Algo) Scan() (ok bool, res []byte, err error) {
	tools.Title(fmt.Sprintf("Programme [%s] scan", a.Name))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_SCAN_PROGRAMME, a.Pc.ID, a.Pc.SecretID),
		nil,
	)
	if err != nil || statusCode != http.StatusOK {
		return false, res, err
	}
	return true, res, err
}
func (a *Algo) Explore(celluleID string, startID string, endID string) (ok bool, res []byte, err error) {
	tools.Title(fmt.Sprintf("Programme [%s] explore cellule [%s] [%s-%s]", a.Name, celluleID, startID, endID))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", api.API_URL, api.ROUTE_EXPLORE_PROGRAMME, a.Pc.ID, a.Pc.SecretID, celluleID, startID, endID),
		nil,
	)
	if err != nil || statusCode != http.StatusOK {
		return false, res, err
	}
	return true, res, err
}
func (a *Algo) Destroy(celluleID int, targetID string) (ok bool, err error) {
	tools.Title(fmt.Sprintf("Programme [%s] destroy -> [%s] cellule [%d]", a.Name, targetID, celluleID))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s/%d/%s", api.API_URL, api.ROUTE_DESTROY_PROGRAMME, a.Pc.ID, a.Pc.SecretID, celluleID, targetID),
		nil,
	)
	if err != nil || statusCode != http.StatusOK {
		return false, err
	}
	a.Psi = structure.ProgrammeStatusInfos{}
	err = json.Unmarshal(res, &a.Psi)
	return true, err
}
func (a *Algo) Rebuild(celluleID int, targetID string) (ok bool, err error) {
	tools.Title(fmt.Sprintf("Programme [%s] rebuild -> [%s] cellule [%s]", a.Name, celluleID, targetID))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s/%d/%s", api.API_URL, api.ROUTE_REBUILD_PROGRAMME, a.Pc.ID, a.Pc.SecretID, celluleID, targetID),
		nil,
	)
	if err != nil || statusCode != http.StatusOK {
		return false, err
	}
	a.Psi = structure.ProgrammeStatusInfos{}
	err = json.Unmarshal(res, &a.Psi)
	return true, err
}
func (a *Algo) GetStatusGrid() (err error) {
	tools.Title(fmt.Sprintf("Status grid"))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s", api.API_URL, api.ROUTE_STATUS_GRID),
		nil,
	)
	if err != nil {
		err = errors.New(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
	} else {
		a.InfosGrid = structure.GridInfos{}
		err = json.Unmarshal(res, &a.InfosGrid)
	}
	return
}
func (a *Algo) Unset() {
	_, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_UNSET_PROGRAMME, a.Pc.ID, a.Pc.SecretID),
		nil,
	)
	if err != nil {
		tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
	} else {

	}
}
func (a *Algo) PrintInfo(printGrid bool) {
	if printGrid {
		tools.PrintGridPosition(a.Psi.Programme, a.InfosGrid.Taille)
	}
	tools.PrintProgramme(a.Psi)
	tools.PrintInfosGrille(a.InfosGrid)
}
func (a *Algo) GetProgramme() (ok bool, programmes []string) {
	if scanOK, res, _ := a.Scan(); !scanOK {
		return scanOK, programmes
	} else {
		var zoneInfos structure.ZoneInfos
		err := json.Unmarshal(res, &zoneInfos)
		if err != nil {
			return false, programmes
		}
		for _, programme := range zoneInfos.Programmes {
			if programme.Status {
				programmes = append(programmes, programme.ID)
			}
		}
	}
	return true, programmes
}
func (a *Algo) Attack(targetID string) {
	for i := 0; i < 10; i++ {
		time.Sleep(TIME_MILLISECONDE * time.Millisecond)
		if ok, _ := a.Destroy(i, targetID); !ok {
			return
		}
		tools.PrintProgramme(a.Psi)
		tools.PrintInfosGrille(a.InfosGrid)
	}
}
func (a *Algo) CheckAttack() {
	maxValeur := a.Psi.Programme.Level * 10
	for len(a.Psi.LockProgramme) > 0 {
		receive_destroy := false
		_ = a.GetStatusGrid()
		for _, cellule := range a.Psi.Programme.Cellules {
			if cellule.Valeur < maxValeur {
				nbrRebuild := maxValeur - cellule.Valeur
				for i := 0; i < nbrRebuild; i++ {
					time.Sleep(TIME_MILLISECONDE * time.Millisecond)
					if ok, _ := a.Rebuild(cellule.ID, a.ID); !ok {
						return
					}
					tools.PrintProgramme(a.Psi)
					tools.PrintInfosGrille(a.InfosGrid)
					if cellule.CurrentAccesLog.ReceiveDestroy {
						receive_destroy = true
						time.Sleep(TIME_MILLISECONDE * time.Millisecond)
						if ok, _ := a.Destroy(cellule.ID, cellule.CurrentAccesLog.PID); !ok {
							return
						}
						tools.PrintProgramme(a.Psi)
						tools.PrintInfosGrille(a.InfosGrid)
					} else {
						receive_destroy = false
					}
				}
			}
		}
		if receive_destroy == false {
			break
		}
		if len(a.Psi.LockProgramme) == 1 {
			if _, ok := a.Psi.LockProgramme[a.ID]; ok {
				break
			}
		}
	}
	return
}
