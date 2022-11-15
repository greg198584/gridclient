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
func _Save(name string) (pc structure.ProgrammeContainer, err error) {
	currentPC, err := _GetProgrammeFile(name)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(currentPC)
	res, statusCode, err := api.RequestApi(
		"GET",
		fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_SAVE_PROGRAMME, currentPC.ID, currentPC.SecretID),
		nil,
	)
	if statusCode == http.StatusCreated {
		err = json.Unmarshal(res, &pc)
		tools.CreateJsonFile(fmt.Sprintf("%s.json", name), pc)
		tools.Success("backup OK")
	} else {
		err = errors.New("erreur chargement programme")
		jsonPretty, _ := tools.PrettyString(res)
		tools.Fail(fmt.Sprintf("status = [%d]", statusCode))
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
func Save(name string) {
	tools.Title(fmt.Sprintf("save programme [%s]", name))
	_, err := _Save(name)
	if err != nil {
		tools.Fail(fmt.Sprintf(" backup [%s] FAIL [%s]", name, err.Error()))
	} else {
		tools.Success(fmt.Sprintf(" backup [%s] OK", name))
		GetInfoProgramme(name, false)
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
		tools.Fail(fmt.Sprintf(" upgrade [%s] FAIL [%s]", name, err.Error()))
	} else {
		tools.Success(fmt.Sprintf("upgrade [%s] OK", name))
		GetInfoProgramme(name, false)
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
func Destroy(name string, celluleID int, targetID string) {
	tools.Title(fmt.Sprintf(
		"Programme [%s] destroy -> [%s] cellule [%s] energy [%s]",
		name,
		celluleID,
		targetID,
		algo.ENERGY_MAX_ATTACK,
	))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.Attack(celluleID, targetID)
	current.PrintInfo(false)
	return
}
func Rebuild(name string, celluleID int, targetID string) {
	tools.Title(fmt.Sprintf(
		"Programme [%s] rebuild -> [%s] cellule [%s] energy [%s]",
		name,
		celluleID,
		targetID,
		algo.ENERGY_MAX_ATTACK,
	))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.Defense(celluleID, targetID)
	current.PrintInfo(false)
	return
}
func GetStatusGrid(pid bool) {
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
			if pid {
				tools.PrintInfosProgrammeGrille(infos)
			}
		}
	}
	return
}
func GetZoneActif(name string) {
	tools.Title(fmt.Sprintf("Scan Zone actif"))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	zoneInfos, _ := current.GetZoneActif()
	tools.PrintZoneActif(zoneInfos)
	return
}
func GetInfoProgramme(name string, printPosition bool) {
	tools.Title(fmt.Sprintf("infos programme"))
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	current.GetInfosProgramme()
	current.PrintInfo(printPosition)
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
	current.GetInfosProgramme()
	current.PrintInfo(false)
}
func DestroyZone(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	status := true
	for status {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		current.GetInfosProgramme()
		current.PrintInfo(false)
		if current.Psi.Programme.ID != "" {
			status = current.Psi.Programme.Status
		}
		ok, zoneInfos := current.GetZoneinfos()
		tools.PrintZoneInfos(zoneInfos)
		if ok && zoneInfos.Status {
			for _, cellule := range zoneInfos.Cellules {
				current.AttackZone(cellule.ID)
			}
		}
	}

}
func AttackTarget(current *algo.Algo, pid string) {
	for _, cellule := range current.Psi.Programme.Cellules {
		if cellule.Status && cellule.Energy > 0 {
			statusTarget := true
			if statusTarget {
				current.Attack(cellule.ID, pid)
			}
		}
	}
}
func Attack(name string, PidList []string, printInfo bool) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	status := true
	for status {
		current.GetInfosProgramme()
		if printInfo {
			current.PrintInfo(false)
		}
		if current.Psi.Programme.ID != "" {
			status = current.Psi.Programme.Status
		}
		_, programmes := current.GetProgramme()
		for _, pid := range programmes {
			if len(PidList) > 0 {
				for _, pidTarget := range PidList {
					if pid == pidTarget {
						AttackTarget(current, pid)
					}
				}
			} else {
				AttackTarget(current, pid)
			}
			if current.Psi.LockProgramme[pid].Status == false {
				break
			}
		}
		current.CheckAttack(printInfo)
	}
}
func CheckAttack(name string, printInfo bool) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	status := true
	for status {
		current.GetInfosProgramme()
		current.PrintInfo(false)
		if current.Psi.Programme.ID != "" {
			status = current.Psi.Programme.Status
		}
		current.CheckAttack(printInfo)
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
	currentZoneID := current.Psi.Programme.Position.ZoneID
	for i := currentZoneID; i <= current.InfosGrid.Taille; i++ {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		if ok, _ := current.Move(i); !ok {
			if current.StatusCode == http.StatusUnauthorized {
				return
			}
		}
		if scanOK, scanRes, _ := current.Scan(); !scanOK {
			jsonPretty, _ := tools.PrettyString(scanRes)
			fmt.Println(jsonPretty)
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
		//current.PrintInfo(true)
	}
}
func SearchEnergy(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	currentZoneID := current.Psi.Programme.Position.ZoneID
	for i := currentZoneID; i <= current.InfosGrid.Taille; i++ {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		if ok, _ := current.Move(i); !ok {
			if current.StatusCode == http.StatusUnauthorized {
				break
			}
		}
		if scanOK, scanRes, _ := current.Scan(); !scanOK {
			jsonPretty, _ := tools.PrettyString(scanRes)
			fmt.Println(jsonPretty)
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
func SearchCelluleTrap(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	currentZoneID := current.Psi.Programme.Position.ZoneID
	for i := currentZoneID; i <= current.InfosGrid.Taille; i++ {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		if ok, _ := current.Move(i); !ok {
			if current.StatusCode == http.StatusUnauthorized {
				break
			}
		}
		if scanOK, scanRes, _ := current.Scan(); !scanOK {
			jsonPretty, _ := tools.PrettyString(scanRes)
			fmt.Println(jsonPretty)
			tools.Fail("erreur scan")
			return
		} else {
			var zoneInfos structure.ZoneInfos
			err := json.Unmarshal(scanRes, &zoneInfos)
			if err != nil {
				tools.Fail(err.Error())
				return
			} else {
				for _, cellule := range zoneInfos.Cellules {
					if cellule.Trapped {
						title := aurora.Red("--- CELLULE DANGER")
						tools.Title(fmt.Sprintf(
							"\t%s >>> [%d][%d] cellule [%d]",
							title,
							current.Psi.Programme.Position.SecteurID,
							current.Psi.Programme.Position.ZoneID,
							cellule.ID,
						))
					}
				}
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
	err = current.GetStatusGrid()
	currentZoneID := current.Psi.Programme.Position.ZoneID
	status := true
	for status {
		//current.PrintInfo(true)
		for i := currentZoneID; i <= current.InfosGrid.Taille; i++ {
			time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
			if ok, _ := current.Move(i); !ok {
				if current.StatusCode == http.StatusBadRequest {
					break
				}
			}
			if scanOK, scanRes, _ := current.Scan(); !scanOK {
				jsonPretty, _ := tools.PrettyString(scanRes)
				fmt.Println(jsonPretty)
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
				//current.PrintInfo(true)
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
func Monitoring(name string, printGrid bool, defense bool) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	for {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		current.GetInfosProgramme()
		current.PrintInfo(printGrid)
		if defense {
			current.CheckAttack(true)
		}
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
