package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/greg198584/gridclient/structure"
	"github.com/logrusorgru/aurora"
	"strconv"
)

func PrintProgramme(psi structure.ProgrammeStatusInfos) {
	var header []string
	var dataList [][]string
	Info(fmt.Sprintf("SECTEUR-ZONE [%d][%d]", psi.Programme.Position.SecteurID, psi.Programme.Position.ZoneID))
	if len(psi.LockProgramme) != 0 {
		header = []string{"name", "cell", "energy", "valeur", "indicator", "status", "---", "status", "indicator", "valeur", "energy", "cell", "name"}
		for _, lockProgramme := range psi.LockProgramme {
			for i := 0; i < len(psi.Programme.Cellules); i++ {
				valeur := psi.Programme.Cellules[i].Valeur
				prefixValeur := psi.Programme.Level
				valeurString := ""
				for j := 0; j < (valeur / prefixValeur); j++ {
					if valeur > 5 {
						valeurString += aurora.Green(fmt.Sprintf("-%d-", prefixValeur)).String()
					} else if valeur > 3 && valeur < 7 {
						valeurString += aurora.Yellow(fmt.Sprintf("-%d-", prefixValeur)).String()
					} else {
						valeurString += aurora.Red(fmt.Sprintf("-%d-", prefixValeur)).String()
					}
				}
				status := ""
				if psi.Programme.Cellules[i].Status {
					status = aurora.Green("OK").String()
				} else {
					status = aurora.Red("NOK").String()
				}

				lpValeur := lockProgramme.Cellules[i].Valeur
				lpPrefixValeur := lockProgramme.Level
				lpValeurString := ""
				for j := 0; j < (lpValeur / lpPrefixValeur); j++ {
					if lpValeur > 5 {
						lpValeurString += aurora.Green(fmt.Sprintf("-%d-", lpPrefixValeur)).String()
					} else if lpValeur > 3 && lpValeur < 7 {
						lpValeurString += aurora.Yellow(fmt.Sprintf("-%d-", lpPrefixValeur)).String()
					} else {
						lpValeurString += aurora.Red(fmt.Sprintf("-%d-", lpPrefixValeur)).String()
					}
				}
				lpStatus := ""
				if lockProgramme.Cellules[i].Status {
					lpStatus = aurora.Green("OK").String()
				} else {
					lpStatus = aurora.Red("NOK").String()
				}
				dataList = append(dataList, []string{
					aurora.Cyan(psi.Programme.Name).String(),
					strconv.FormatInt(int64(psi.Programme.Cellules[i].ID), 10),
					strconv.FormatInt(int64(psi.Programme.Cellules[i].Energy), 10),
					strconv.FormatInt(int64(psi.Programme.Cellules[i].Valeur), 10),
					valeurString,
					status,
					"***",
					lpStatus,
					lpValeurString,
					strconv.FormatInt(int64(lockProgramme.Cellules[i].Valeur), 10),
					strconv.FormatInt(int64(lockProgramme.Cellules[i].Energy), 10),
					strconv.FormatInt(int64(lockProgramme.Cellules[i].ID), 10),
					aurora.Red(lockProgramme.Name).String(),
				})
			}
			fmt.Printf(
				"<--- Programme Info status [%s][%s] %s [%s][%s] --->\n",
				aurora.Cyan(psi.Programme.Name),
				aurora.Green(psi.Programme.ID),
				aurora.White("VS"),
				aurora.Red(lockProgramme.Name),
				aurora.Green(lockProgramme.ID),
			)
			PrintColorTable(header, dataList)
			dataList = nil
		}
	} else {
		header = []string{"name", "cell", "energy", "valeur", "indicator", "status"}
		for i := 0; i < len(psi.Programme.Cellules); i++ {
			valeur := psi.Programme.Cellules[i].Valeur
			prefixValeur := psi.Programme.Level
			valeurString := ""
			for j := 0; j < (valeur / prefixValeur); j++ {
				if valeur > 5 {
					valeurString += aurora.Green(fmt.Sprintf("-%d-", prefixValeur)).String()
				} else if valeur > 3 && valeur < 7 {
					valeurString += aurora.Yellow(fmt.Sprintf("-%d-", prefixValeur)).String()
				} else {
					valeurString += aurora.Red(fmt.Sprintf("-%d-", prefixValeur)).String()
				}
			}
			status := ""
			if psi.Programme.Cellules[i].Status {
				status = aurora.Green("OK").String()
			} else {
				status = aurora.Red("NOK").String()
			}
			dataList = append(dataList, []string{
				aurora.Cyan(psi.Programme.Name).String(),
				strconv.FormatInt(int64(psi.Programme.Cellules[i].ID), 10),
				strconv.FormatInt(int64(psi.Programme.Cellules[i].Energy), 10),
				strconv.FormatInt(int64(psi.Programme.Cellules[i].Valeur), 10),
				valeurString,
				status,
			})
		}
		fmt.Printf("<---[ Programme Info status [%s] [%s] ]--->\n", aurora.Cyan(psi.Programme.Name), aurora.Green(psi.Programme.ID))
		PrintColorTable(header, dataList)
		dataList = nil
	}
	flag := ""
	for _, cellule := range psi.Programme.Cellules {
		for _, data := range cellule.Datas {
			if data.IsFlag {
				flag = data.Content
			}
		}

	}
	if flag != "" {
		fmt.Printf("<---[ %s - [%s] ]--->\n", aurora.Green("FLAG CAPTURER"), aurora.Cyan(flag))
	}
	return
}
func PrintGridPosition(programme structure.Programme, size int) {
	var header []string
	var data [][]string
	header = append(header, "secteur")
	for i := 0; i < size; i++ {
		header = append(header, fmt.Sprintf("zone %d", i))
	}
	for i := 0; i < size; i++ {
		var tmpData []string
		tmpData = append(tmpData, aurora.Cyan(fmt.Sprintf("%d", i)).String())
		for j := 0; j < size; j++ {
			value := aurora.Red("0").String()
			if programme.Position.SecteurID == i && programme.Position.ZoneID == j {
				value = aurora.Green("1").String()
			}
			tmpData = append(tmpData, value)
		}
		data = append(data, tmpData)
	}
	PrintColorTable(header, data)
}
func PrintInfosGrille(infos structure.GridInfos) {
	var header = []string{"ID", "Taille", "ZoneTransfert", "Iteration", "Cycle", "NbrProgramme", "status", "Version"}
	var InfosTab [][]string

	InfosTab = append(InfosTab, []string{
		infos.Id,
		strconv.FormatInt(int64(infos.Taille), 10),
		fmt.Sprintf("S%d-Z%d", infos.ZoneTransfert.SecteurID, infos.ZoneTransfert.ZoneID),
		strconv.FormatInt(int64(infos.Iteration), 10),
		strconv.FormatInt(int64(infos.Cycle), 10),
		strconv.FormatInt(int64(infos.NbrProgrammes), 10),
		fmt.Sprintf("%t", infos.Status),
		infos.Version,
	})
	PrintColorTable(header, InfosTab, "<---[ Infos grille ]--->")
	return
}
func PrintZoneInfos(infos structure.ZoneInfos) {
	var header = []string{"CELL ID", "STATUS", "DATA", "DATA_TYPE"}
	var cellData [][]string
	for i := 0; i < len(infos.Cellules); i++ {
		dataTypeBytes := new(bytes.Buffer)
		json.NewEncoder(dataTypeBytes).Encode(infos.Cellules[i].DataType)
		//dataType, _ := PrettyString(dataTypeBytes.Bytes())
		cellData = append(cellData, []string{
			fmt.Sprintf("%d", infos.Cellules[i].ID),
			fmt.Sprintf("%t", infos.Cellules[i].Status),
			fmt.Sprintf("%d", infos.Cellules[i].DataCount),
			fmt.Sprintf("%s", dataTypeBytes.String()),
		})
	}
	PrintColorTable(header, cellData, fmt.Sprintf("<---[ Infos programme sur Zone [%d] ]--->", infos.ID))
	header = []string{"PID", "NAME", "VALEUR TOTAL", "ENERGY TOTAL"}
	var progrData [][]string
	for i := 0; i < len(infos.Programmes); i++ {
		progrData = append(progrData, []string{
			infos.Programmes[i].ID,
			infos.Programmes[i].Name,
			fmt.Sprintf("%d", infos.Programmes[i].ValeurTotal),
			fmt.Sprintf("%d", infos.Programmes[i].EnergyTotal),
		})
	}
	PrintColorTable(header, progrData, fmt.Sprintf("<---[ Infos Zone [%d] ]--->", infos.ID))
	return
}
func PrintExplore(celluleID string, data map[int]structure.CelluleData) {
	var header = []string{"ID", "ENERGY", "IS_FLAG"}
	var cellData [][]string

	for i := 0; i < len(data); i++ {
		isFlag := aurora.Red("false")
		if data[i].IsFlag {
			isFlag = aurora.Green("true")
		}
		cellData = append(cellData, []string{
			fmt.Sprintf("%d", data[i].ID),
			fmt.Sprintf("%d", data[i].Energy),
			fmt.Sprintf("%s", isFlag),
		})
	}
	PrintColorTable(header, cellData, fmt.Sprintf("<---[ data cellule [%s] ]--->", celluleID))
	return
}
