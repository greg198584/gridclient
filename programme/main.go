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
	psi, err := _LoadProgramme(name)
	if err != nil {
		tools.Fail(err.Error())
	} else {
		tools.Success(fmt.Sprintf("programme ajouter [%s]", name))
		tools.PrintProgramme(psi)
	}
}
func Delete(name string) {
	tools.Title(fmt.Sprintf("suppression programme [%s]", name))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_UNSET_PROGRAMME, pc.ID, pc.SecretID),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			jsonPretty, _ := tools.PrettyString(res)
			tools.Info(fmt.Sprintf("status = [%d]", statusCode))
			fmt.Println(jsonPretty)
		}
	}
}
func JumpUp(name string, valeur string) {
	tools.Title(fmt.Sprintf("Programme [%s] JumpUP [%s]", name, valeur))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%s", api.API_URL, api.ROUTE_JUMPUP_PROGRAMME, pc.ID, pc.SecretID, valeur),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				tools.PrintProgramme(psi)
			}
		}
	}
}
func JumpDown(name string, valeur string) {
	tools.Title(fmt.Sprintf("Programme [%s] JumpDown [%s]", name, valeur))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%s", api.API_URL, api.ROUTE_JUMPDOWN_PROGRAMME, pc.ID, pc.SecretID, valeur),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				tools.PrintProgramme(psi)
			}
		}
	}
}
func Move(name string, valeur string) {
	tools.Title(fmt.Sprintf("Programme [%s] Move to Zone [%s]", name, valeur))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%s", api.API_URL, api.ROUTE_MOVE_PROGRAMME, pc.ID, pc.SecretID, valeur),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				GetStatusGrid()
			}
		}
	}
}
func Scan(name string) {
	tools.Title(fmt.Sprintf("Programme [%s] scan", name))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_SCAN_PROGRAMME, pc.ID, pc.SecretID),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
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
}
func Explore(name string, celluleID string) {
	tools.Title(fmt.Sprintf("Programme [%s] explore cellule [%s]", name, celluleID))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%s", api.API_URL, api.ROUTE_EXPLORE_PROGRAMME, pc.ID, pc.SecretID, celluleID),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
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
}
func Destroy(name string, celluleID int, targetID string, energy int) {
	tools.Title(fmt.Sprintf("Programme [%s] destroy -> [%s] cellule [%s] energy [%s]", name, celluleID, targetID, energy))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		for i := 0; i < energy; i++ {
			res, statusCode, err := api.RequestApi(
				"GET",
				fmt.Sprintf("%s/%s/%s/%s/%d/%s", api.API_URL, api.ROUTE_DESTROY_PROGRAMME, pc.ID, pc.SecretID, celluleID, targetID),
				nil,
			)
			if err != nil {
				tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
			} else {
				//jsonPretty, _ := tools.PrettyString(res)
				//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
				//fmt.Println(jsonPretty)
				var psi structure.ProgrammeStatusInfos
				err = json.Unmarshal(res, &psi)
				if err != nil {
					tools.Fail(err.Error())
				} else {
					tools.PrintGridPosition(psi.Programme, 10)
					tools.PrintProgramme(psi)
				}
			}
		}
	}
	return
}
func Rebuild(name string, celluleID int, targetID string, energy int) {
	tools.Title(fmt.Sprintf("Programme [%s] rebuild -> [%s] cellule [%s] energy [%s]", name, celluleID, targetID, energy))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)
		for i := 0; i < energy; i++ {
			res, statusCode, err := api.RequestApi(
				"GET",
				fmt.Sprintf("%s/%s/%s/%s/%d/%s", api.API_URL, api.ROUTE_REBUILD_PROGRAMME, pc.ID, pc.SecretID, celluleID, targetID),
				nil,
			)
			if err != nil {
				tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
			} else {
				//jsonPretty, _ := tools.PrettyString(res)
				//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
				//fmt.Println(jsonPretty)
				var psi structure.ProgrammeStatusInfos
				err = json.Unmarshal(res, &psi)
				if err != nil {
					tools.Fail(err.Error())
				} else {
					tools.PrintGridPosition(psi.Programme, 10)
					tools.PrintProgramme(psi)
				}
			}
		}
	}
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
		//jsonPretty, _ := tools.PrettyString(res)
		//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
		//fmt.Println(jsonPretty)
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

	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
	} else {
		res, _, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_STATUS_PROGRAMME, pc.ID, pc.SecretID),
			nil,
		)
		if err != nil {
			tools.Fail(err.Error())
		}
		var psi structure.ProgrammeStatusInfos
		err = json.Unmarshal(res, &psi)
		if err != nil {
			tools.Fail(err.Error())
		}

		tools.PrintGridPosition(psi.Programme, 10)
		tools.PrintProgramme(psi)
	}
}

func CaptureTargetData(name string, celluleID int, targetID string) {
	tools.Title(fmt.Sprintf("[%s] Capture data target [%s] - cellule [%s]", name, targetID, celluleID))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)

		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%d/%s", api.API_URL, api.ROUTE_CAPTURE_TARGET_DATA, pc.ID, pc.SecretID, celluleID, targetID),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			//jsonPretty, _ := tools.PrettyString(res)
			//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
			//fmt.Println(jsonPretty)
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				tools.PrintProgramme(psi)
			}
		}
	}
	return
}
func CaptureCellData(name string, celluleID int, index int) {
	tools.Title(fmt.Sprintf("[%s] Capture data cellule [%d] - index [%d]", name, celluleID, index))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)

		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%d/%d", api.API_URL, api.ROUTE_CAPTURE_CELL_DATA, pc.ID, pc.SecretID, celluleID, index),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			//jsonPretty, _ := tools.PrettyString(res)
			//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
			//fmt.Println(jsonPretty)
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				tools.PrintProgramme(psi)
			}
		}
	}
	return
}
func CaptureTargetEnergy(name string, celluleID int, targetID string) {
	tools.Title(fmt.Sprintf("[%s] Capture energy target [%s] - cellule [%s]", name, targetID, celluleID))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)

		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%d/%s", api.API_URL, api.ROUTE_CAPTURE_TARGET_ENERGY, pc.ID, pc.SecretID, celluleID, targetID),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			//jsonPretty, _ := tools.PrettyString(res)
			//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
			//fmt.Println(jsonPretty)
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				tools.PrintProgramme(psi)
			}
		}
	}
	return
}
func CaptureCellEnergy(name string, celluleID int, index int) {
	tools.Title(fmt.Sprintf("[%s] Capture energy cellule [%s] - index [%d]", name, celluleID, index))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)

		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s/%d/%d", api.API_URL, api.ROUTE_CAPTURE_CELL_ENERGY, pc.ID, pc.SecretID, celluleID, index),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			//jsonPretty, _ := tools.PrettyString(res)
			//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
			//fmt.Println(jsonPretty)
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				tools.PrintProgramme(psi)
			}
		}
	}
	return
}
func Equilibrium(name string) {
	tools.Title(fmt.Sprintf("Equilibrium energy programme [%s]", name))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)

		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_EQUILIBRiUM, pc.ID, pc.SecretID),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			//jsonPretty, _ := tools.PrettyString(res)
			//tools.Info(fmt.Sprintf("status = [%d]", statusCode))
			//fmt.Println(jsonPretty)
			var psi structure.ProgrammeStatusInfos
			err = json.Unmarshal(res, &psi)
			if err != nil {
				tools.Fail(err.Error())
			} else {
				tools.PrintGridPosition(psi.Programme, 10)
				tools.PrintProgramme(psi)
			}
		}
	}
	return
}
func PushFlag(name string) {
	tools.Title(fmt.Sprintf("Push flag - programme [%s]", name))
	pc, err := _GetProgrammeFile(name)
	if err != nil {
		tools.Fail(err.Error())
		return
	} else {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(pc)

		res, statusCode, err := api.RequestApi(
			"GET",
			fmt.Sprintf("%s/%s/%s/%s", api.API_URL, api.ROUTE_PUSH_FLAG, pc.ID, pc.SecretID),
			nil,
		)
		if err != nil {
			tools.Fail(fmt.Sprintf("status code [%d] - [%s]", statusCode, err.Error()))
		} else {
			if statusCode != http.StatusOK {
				tools.Fail(fmt.Sprintf("erreur status code [%d]"))
			} else {
				var newPC structure.ProgrammeContainer
				err = json.Unmarshal(res, &newPC)
				if err != nil {
					tools.Fail(err.Error())
				} else {
					err = tools.CreateJsonFile(fmt.Sprintf("%s.json", name), newPC)
					if err != nil {
						tools.Fail("erreur sauvegarde programme")
					} else {
						tools.Success("sauvegarde effectuer")
					}
				}
			}
		}
	}
	return
}
func Attack(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	if err != nil {
		panic(err)
	}
	status := true
	for status {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		status, _ = current.GetInfosProgramme()
		current.CheckAttack()
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
func CheckAttack(name string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	if err != nil {
		panic(err)
	}
	status := true
	for status {
		time.Sleep(algo.TIME_MILLISECONDE * time.Millisecond)
		status, _ = current.GetInfosProgramme()
		current.CheckAttack()
		status = current.Psi.Programme.Status
		tools.PrintInfosGrille(current.InfosGrid)
	}
}
func MovePosition(name string, position string) {
	current, err := algo.NewAlgo(name)
	if err != nil {
		//panic(err)
	}
	err = current.GetStatusGrid()
	if err != nil {
		panic(err)
	}
	splitPosition := strings.Split(position, "-")
	secteurID, _ := strconv.Atoi(splitPosition[0])
	zoneID, _ := strconv.Atoi(splitPosition[1])
	currentSecteurID := current.Psi.Programme.Position.SecteurID
	nbrJump := 0
	fmt.Printf("S-Z = [%d-%d] - current secteur ID [%d]\n", secteurID, zoneID, currentSecteurID)
	if secteurID > currentSecteurID {
		nbrJump = secteurID - currentSecteurID
		current.JumpDown(nbrJump)
	} else {
		nbrJump = currentSecteurID - secteurID
		current.JumpUp(nbrJump)
	}
	current.Move(zoneID)
	tools.PrintGridPosition(current.Psi.Programme, current.InfosGrid.Taille)
	tools.PrintInfosGrille(current.InfosGrid)
}
