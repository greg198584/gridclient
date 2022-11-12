package programme

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/greg198584/gridclient/algo"
	"github.com/greg198584/gridclient/api"
	"github.com/greg198584/gridclient/structure"
	"github.com/greg198584/gridclient/tools"
	"github.com/logrusorgru/aurora"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func _IsExistFile(name string) bool {
	filename := fmt.Sprintf("%s.json", name)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func _CreateProgramme(name string) (programme structure.ProgrammeContainer, err error) {
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s", api.API_URL, api.ROUTE_NEW_PROGRAMME, name),
		nil,
	)
	if err != nil {
		return
	}
	if statusCode == http.StatusCreated {
		err = json.Unmarshal(res, &programme)
		tools.CreateJsonFile(fmt.Sprintf("%s.json", name), programme)
	} else {
		err = errors.New("erreur creation programme")
	}
	return
}
func _LoadProgramme(name string) (psi structure.ProgrammeStatusInfos, err error) {
	pc, err := _GetProgrammeFile(name)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(pc)
	res, statusCode, err := api.RequestApi(
		"POST",
		fmt.Sprintf("%s/%s", api.API_URL, api.ROUTE_LOAD_PROGRAMME),
		reqBodyBytes.Bytes(),
	)
	if statusCode == http.StatusCreated {
		err = json.Unmarshal(res, &psi)
	} else {
		err = errors.New("erreur chargement programme")
		jsonPretty, _ := tools.PrettyString(res)
		tools.Info(fmt.Sprintf("status = [%d]", statusCode))
		fmt.Println(jsonPretty)
	}
	return
}
func _UpgradeProgramme(name string) (pc structure.ProgrammeContainer, err error) {
	currentPC, err := _GetProgrammeFile(name)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(currentPC)
	res, statusCode, err := api.RequestApi(
		"POST",
		fmt.Sprintf("%s/%s", api.API_URL, api.ROUTE_UPGRADE_PROGRAMME),
		reqBodyBytes.Bytes(),
	)
	if statusCode == http.StatusCreated {
		err = json.Unmarshal(res, &pc)
		tools.CreateJsonFile(fmt.Sprintf("%s.json", name), pc)
	} else {
		err = errors.New("erreur chargement programme")
		jsonPretty, _ := tools.PrettyString(res)
		tools.Info(fmt.Sprintf("status = [%d]", statusCode))
		fmt.Println(jsonPretty)
	}
	return
}
func _GetProgrammeFile(name string) (pc *structure.ProgrammeContainer, err error) {
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
func New(name string) {
	tools.Title(fmt.Sprintf("création programme [%s]", name))
	if _IsExistFile(name) == false {

		programmeContainer, err := _CreateProgramme(name)
		if err != nil {
			tools.Fail(err.Error())
		} else {
			tools.Success(fmt.Sprintf("programme ajouter ID = [%s]", programmeContainer.ID))
			tools.Info(fmt.Sprintf("programme info"))
			Info(&programmeContainer)
		}
	} else {
		tools.Warning(fmt.Sprintf("programme file exist"))
	}
}
func Info(pc *structure.ProgrammeContainer) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(pc.Programme)
	jsonPretty, _ := tools.PrettyString(reqBodyBytes.Bytes())
	fmt.Println(jsonPretty)
}
func Load(name string) {
	tools.Title(fmt.Sprintf("chargement programme [%s]", name))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.PrintInfo(true)
}
func Upgrade(name string) {
	tools.Title(fmt.Sprintf("chargement programme [%s]", name))
	_, err := _UpgradeProgramme(name)
	if err != nil {
		tools.Fail(err.Error())
	} else {
		tools.Success(fmt.Sprintf("programme ajouter [%s]", name))
		GetInfoProgramme(name)
	}
}
func Delete(name string) {
	tools.Title(fmt.Sprintf("suppression programme [%s]", name))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.Unset()
}
func JumpUp(name string, valeur string) {
	tools.Title(fmt.Sprintf("Programme [%s] JumpUP [%s]", name, valeur))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	valeurInt, err := strconv.Atoi(valeur)
	if err != nil {
		return
	}
	current.JumpUp(valeurInt)
	current.PrintInfo(true)
}
func JumpDown(name string, valeur string) {
	tools.Title(fmt.Sprintf("Programme [%s] JumpDown [%s]", name, valeur))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	valeurInt, err := strconv.Atoi(valeur)
	if err != nil {
		return
	}
	current.JumpDown(valeurInt)
	current.PrintInfo(true)
}
func Move(name string, valeur string) {
	tools.Title(fmt.Sprintf("Programme [%s] Move to Zone [%s]", name, valeur))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	valeurInt, err := strconv.Atoi(valeur)
	if err != nil {
		return
	}
	current.Move(valeurInt)
	current.PrintInfo(true)
}
func Scan(name string) {
	tools.Title(fmt.Sprintf("Programme [%s] scan", name))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	_, res, err := current.Scan()
	if err != nil {
		tools.Fail(err.Error())
	} else {
		var zoneInfos structure.ZoneInfos
		err = json.Unmarshal(res, &zoneInfos)
		if err != nil {
			tools.Fail(err.Error())
		} else {
			tools.PrintZoneInfos(zoneInfos)
		}
	}
}
func Explore(name string, celluleID string) {
	tools.Title(fmt.Sprintf("Programme [%s] explore cellule [%s]", name, celluleID))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	celluleIdInt, err := strconv.Atoi(celluleID)
	_, res, err := current.Explore(celluleIdInt)
	if err != nil {
		tools.Fail(err.Error())
	} else {
		var celluleData map[int]structure.CelluleData
		err = json.Unmarshal(res, &celluleData)
		if err != nil {
			tools.Fail(err.Error())
		} else {
			tools.PrintExplore(celluleID, celluleData)
		}
	}
}
func Destroy(name string, celluleID int, targetID string, energy int) {
	tools.Title(fmt.Sprintf("Programme [%s] destroy -> [%s] cellule [%s] energy [%s]", name, celluleID, targetID, energy))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.Attack(celluleID, targetID, energy)
	current.PrintInfo(false)
	return
}
func Rebuild(name string, celluleID int, targetID string, energy int) {
	tools.Title(fmt.Sprintf("Programme [%s] rebuild -> [%s] cellule [%s] energy [%s]", name, celluleID, targetID, energy))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.Rebuild(celluleID, targetID)
	current.PrintInfo(false)
	return
}
func GetStatusGrid() {
	tools.Title(fmt.Sprintf("Status grid"))
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s", api.API_URL, api.ROUTE_STATUS_GRID),
		nil,
	)
	if err != nil {
		tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
	} else {
		var infos structure.GridInfos
		err = json.Unmarshal(res, &infos)
		if err != nil {
			tools.Fail(err.Error())
		} else {
			tools.PrintInfosGrille(infos)
		}
	}
	return
}
func GetInfoProgramme(name string) {
	tools.Title(fmt.Sprintf("infos programme"))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.GetInfosProgramme()
	current.PrintInfo(true)
}

func CaptureTargetData(name string, celluleID int, targetID string) {
	tools.Title(fmt.Sprintf("[%s] Capture data target [%s] - cellule [%s]", name, targetID, celluleID))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.CaptureTargetData(celluleID, targetID)
	current.PrintInfo(false)
	return
}
func CaptureCellData(name string, celluleID int, index int) {
	tools.Title(fmt.Sprintf("[%s] Capture data cellule [%d] - index [%d]", name, celluleID, index))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.CaptureCellData(celluleID, index)
	current.PrintInfo(false)
	return
}
func CaptureTargetEnergy(name string, celluleID int, targetID string) {
	tools.Title(fmt.Sprintf("[%s] Capture energy target [%s] - cellule [%s]", name, targetID, celluleID))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.CaptureTargetEnergy(celluleID, targetID)
	current.PrintInfo(false)
	return
}
func CaptureCellEnergy(name string, celluleID int, index int) {
	tools.Title(fmt.Sprintf("[%s] Capture energy cellule [%s] - index [%d]", name, celluleID, index))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.CaptureCellEnergy(celluleID, index)
	current.PrintInfo(false)
	return
}
func Equilibrium(name string) {
	tools.Title(fmt.Sprintf("Equilibrium energy programme [%s]", name))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.Equilibrium()
	current.PrintInfo(false)
}
func PushFlag(name string) {
	tools.Title(fmt.Sprintf("Push flag - programme [%s]", name))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.PushFlag()
	current.PrintInfo(false)
}
func Attack(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	status := true
	for status {
		status = current.Psi.Programme.Status
		ok, programmes := current.GetProgramme()
		if !ok {
			status = false
			break
		}
		current.CheckAttack()
		for _, pid := range programmes {
			for i := 0; i < algo.MAX_CELLULES; i++ {
				if current.Psi.Programme.Cellules[i].Status {
					statusTarget := true
					if _, okLP := current.Psi.LockProgramme[pid]; okLP {
						statusTarget = current.Psi.LockProgramme[pid].Cellules[i].Status
					}
					if statusTarget {
						current.Attack(i, pid, algo.ENERGY_MAX_ATTACK/5)
					}
				}
				current.PrintInfo(false)
			}
			if current.Psi.LockProgramme[pid].Status == false {
				break
			}
		}
	}
}
func CheckAttack(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	status := true
	for status {
		status = current.Psi.Programme.Status
		current.CheckAttack()
	}
}
func MovePosition(name string, position string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	splitPosition := strings.Split(position, "-")
	secteurID, _ := strconv.Atoi(splitPosition[0])
	zoneID, _ := strconv.Atoi(splitPosition[1])
	current.QuickMove(secteurID, zoneID)
	current.PrintInfo(true)
}

func SearchFlag(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	current.Move(0)
	for i := 0; i <= current.InfosGrid.Taille; i++ {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		if ok, _ := current.Move(i); !ok {
			if current.StatusCode == http.StatusUnauthorized {
				return
			}
		}
		if scanOK, scanRes, _ := current.Scan(); !scanOK {
			tools.Fail("erreur scan")
			return
		} else {
			var zoneInfos structure.ZoneInfos
			err := json.Unmarshal(scanRes, &zoneInfos)
			if err != nil {
				tools.Fail(err.Error())
				return
			} else {
				if ok := current.SearchFlag(zoneInfos.Cellules); ok {
					return
				}
			}
		}
		//current.PrintInfo(false)
	}
}
func SearchEnergy(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	current.Move(0)
	for i := 0; i <= current.InfosGrid.Taille; i++ {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		if ok, _ := current.Move(i); !ok {
			if current.StatusCode == http.StatusUnauthorized {
				break
			}
		}
		if scanOK, scanRes, _ := current.Scan(); !scanOK {
			tools.Fail("erreur scan")
			return
		} else {
			var zoneInfos structure.ZoneInfos
			err := json.Unmarshal(scanRes, &zoneInfos)
			if err != nil {
				tools.Fail(err.Error())
				return
			} else {
				current.SearchEnergy(zoneInfos.Cellules)
			}
			//current.PrintInfo(false)
		}
	}
}
func SearchProgramme(name string, all bool) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	if all {
		current.QuickMove(0, 0)
	}
	status := true
	for status {
		//current.PrintInfo(true)
		for i := 0; i <= current.InfosGrid.Taille; i++ {
			time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
			if ok, _ := current.Move(i); !ok {
				if current.StatusCode == http.StatusBadRequest {
					break
				}
			}
			if scanOK, scanRes, _ := current.Scan(); !scanOK {
				tools.Fail("erreur scan")
			} else {
				var zoneInfos structure.ZoneInfos
				json.Unmarshal(scanRes, &zoneInfos)
				programmeFound := false
				for _, programme := range zoneInfos.Programmes {
					programmeFound = true
					tools.Success("PROGRAMME FOUND")
					fmt.Printf("\n\t>>> pprogramme trouver [%s] [%d] [%t]\n", aurora.Green(programme.Name), aurora.Cyan(programme.ID), programme.Status)
					break
				}
				if programmeFound {
					return
				}
				//current.PrintInfo(false)
			}
		}
		if ok, _ := current.Move(0); !ok {
			continue
		}
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		if ok, _ := current.JumpDown(1); !ok {
			if current.StatusCode == http.StatusBadRequest {
				status = false
				break
			}
		}
	}
}
func Monitoring(name string, printGrid bool) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	for {
		time.Sleep(500 * time.Millisecond)
		current.GetInfosProgramme()
		current.PrintInfo(printGrid)
	}
}
func GetCelluleLog(name string, celluleID string) {
	tools.Title(fmt.Sprintf("GET LOG cellule [%s] - programme [%s]", celluleID, name))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s/%s", api.API_URL, api.ROUTE_GET_CELLULE_LOG, current.Pc.ID, current.Pc.SecretID, celluleID),
		nil,
	)
	if err != nil {
		tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
	} else {
		var celluleLogs []structure.CelluleLog
		err = json.Unmarshal(res, &celluleLogs)
		if err != nil {
			tools.Fail(err.Error())
		} else {
			tools.PrintCelluleLogs(celluleLogs)
		}
	}
	return
}
