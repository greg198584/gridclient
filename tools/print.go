package tools

import (
	"fmt"
	"github.com/greg198584/gridclient/structure"
	"github.com/logrusorgru/aurora"
	"strconv"
	"time"
)

func PrintProgramme(psi structure.ProgrammeStatusInfos) {
	var header []string
	var dataList [][]string
	Info(fmt.Sprintf("SECTEUR-ZONE [%d][%d]", psi.Programme.Position.SecteurID, psi.Programme.Position.ZoneID))
	if len(psi.LockProgramme) != 0 {
		header = []string{"name", "cell", "energy", "valeur", "indicator", "status", "capture", "exploration", "---", "exploration", "capture", "status", "indicator", "valeur", "energy", "cell", "name"}
		for _, lockProgramme := range psi.LockProgramme {
			for i := 0; i < len(psi.Programme.Cellules); i++ {
				valeur := psi.Programme.Cellules[i].Valeur
				prefixValeur := psi.Programme.Level
				valeurString := ""
				count := valeur / (prefixValeur * 10)
				for j := 0; j < count; j++ {
					valeurString += "-"
					if valeur > 5 {
						valeurString += aurora.Green(fmt.Sprintf("%d", prefixValeur)).String()
					} else if valeur > 3 && valeur < 7 {
						valeurString += aurora.Yellow(fmt.Sprintf("%d", prefixValeur)).String()
					} else {
						valeurString += aurora.Red(fmt.Sprintf("%d", prefixValeur)).String()
					}
				}
				valeurString += "-"
				status := ""
				if psi.Programme.Cellules[i].Status {
					status = aurora.Green("OK").String()
				} else {
					status = aurora.Red("NOK").String()
				}
				exploration := ""
				if psi.Programme.Cellules[i].Exploration {
					exploration = aurora.Green("OK").String()
				} else {
					exploration = aurora.Red("NOK").String()
				}
				capture := ""
				if psi.Programme.Cellules[i].Capture {
					capture = aurora.Green("OK").String()
				} else {
					capture = aurora.Red("NOK").String()
				}

				lpValeur := lockProgramme.Cellules[i].Valeur
				lpPrefixValeur := lockProgramme.Level
				lpValeurString := ""
				lpCount := lpValeur / (lpPrefixValeur * 10)
				for j := 0; j < lpCount; j++ {
					lpValeurString += "-"
					if lpValeur > 5 {
						lpValeurString += aurora.Green(fmt.Sprintf("%d", lpPrefixValeur)).String()
					} else if lpValeur > 3 && lpValeur < 7 {
						lpValeurString += aurora.Yellow(fmt.Sprintf("%d", lpPrefixValeur)).String()
					} else {
						lpValeurString += aurora.Red(fmt.Sprintf("%d", lpPrefixValeur)).String()
					}
				}
				lpValeurString += "-"
				lpStatus := ""
				if lockProgramme.Cellules[i].Status {
					lpStatus = aurora.Green("OK").String()
				} else {
					lpStatus = aurora.Red("NOK").String()
				}
				lpExploration := ""
				if lockProgramme.Cellules[i].Exploration {
					lpExploration = aurora.Green("OK").String()
				} else {
					lpExploration = aurora.Red("NOK").String()
				}
				lpCapture := ""
				if lockProgramme.Cellules[i].Capture {
					lpCapture = aurora.Green("OK").String()
				} else {
					lpCapture = aurora.Red("NOK").String()
				}
				dataList = append(dataList, []string{
					aurora.Cyan(psi.Programme.Name).String(),
					strconv.FormatInt(int64(psi.Programme.Cellules[i].ID), 10),
					strconv.FormatInt(int64(psi.Programme.Cellules[i].Energy), 10),
					strconv.FormatInt(int64(psi.Programme.Cellules[i].Valeur), 10),
					valeurString,
					status,
					capture,
					exploration,
					"***",
					lpExploration,
					lpCapture,
					lpStatus,
					lpValeurString,
					strconv.FormatInt(int64(lockProgramme.Cellules[i].Valeur), 10),
					strconv.FormatInt(int64(lockProgramme.Cellules[i].Energy), 10),
					strconv.FormatInt(int64(lockProgramme.Cellules[i].ID), 10),
					aurora.Red(lockProgramme.Name).String(),
				})
			}
			//fmt.Printf(
			//	"<--- Programme Info status [%s][%s] %s [%s][%s] --->\n",
			//	aurora.Cyan(psi.Programme.Name),
			//	aurora.Green(psi.Programme.ID),
			//	aurora.White("VS"),
			//	aurora.Red(lockProgramme.Name),
			//	aurora.Green(lockProgramme.ID),
			//)
			PiStatus := ""
			if psi.Programme.Status {
				PiStatus = aurora.Green("OK").String()
			} else {
				PiStatus = aurora.Red("NOK").String()
			}
			LPStatus := ""
			if lockProgramme.Status {
				LPStatus = aurora.Green("OK").String()
			} else {
				LPStatus = aurora.Red("NOK").String()
			}
			// Infos programme
			var PiHeader = []string{
				"Name",
				"ID",
				"Status",
				"VS",
				"Status",
				"ID",
				"Name",
			}
			var PiData [][]string

			PiData = append(PiData, []string{
				aurora.Cyan(psi.Programme.Name).String(),
				aurora.Green(psi.Programme.ID).String(),
				PiStatus,
				"X",
				LPStatus,
				aurora.Red(lockProgramme.ID).String(),
				aurora.Red(lockProgramme.Name).String(),
			})
			PrintColorTable(PiHeader, PiData, "<---[ Programme infos ]--->")
			PrintColorTable(header, dataList)
			dataList = nil
		}
	} else {
		header = []string{"name", "cell", "energy", "valeur", "indicator", "status", "capture", "exploration"}
		for i := 0; i < len(psi.Programme.Cellules); i++ {
			valeur := psi.Programme.Cellules[i].Valeur
			prefixValeur := psi.Programme.Level
			valeurString := ""
			count := valeur / (prefixValeur * 10)
			for j := 0; j < count; j++ {
				valeurString += "-"
				if valeur > 5 {
					valeurString += aurora.Green(fmt.Sprintf("%d", prefixValeur)).String()
				} else if valeur > 3 && valeur < 7 {
					valeurString += aurora.Yellow(fmt.Sprintf("%d", prefixValeur)).String()
				} else {
					valeurString += aurora.Red(fmt.Sprintf("%d", prefixValeur)).String()
				}
			}
			valeurString += "-"
			status := ""
			if psi.Programme.Cellules[i].Status {
				status = aurora.Green("OK").String()
			} else {
				status = aurora.Red("NOK").String()
			}
			exploration := ""
			if psi.Programme.Cellules[i].Exploration {
				exploration = aurora.Green("OK").String()
			} else {
				exploration = aurora.Red("NOK").String()
			}
			capture := ""
			if psi.Programme.Cellules[i].Capture {
				capture = aurora.Green("OK").String()
			} else {
				capture = aurora.Red("NOK").String()
			}
			dataList = append(dataList, []string{
				aurora.Cyan(psi.Programme.Name).String(),
				strconv.FormatInt(int64(psi.Programme.Cellules[i].ID), 10),
				strconv.FormatInt(int64(psi.Programme.Cellules[i].Energy), 10),
				strconv.FormatInt(int64(psi.Programme.Cellules[i].Valeur), 10),
				valeurString,
				status,
				capture,
				exploration,
			})
		}
		fmt.Printf("<---[ Programme info ID [%s] ]--->\n", aurora.Green(psi.Programme.ID))
		PiStatus := ""
		if psi.Programme.Status {
			PiStatus = aurora.Green("OK").String()
		} else {
			PiStatus = aurora.Red("NOK").String()
		}
		PiNagivation := ""
		if psi.Navigation {
			PiNagivation = aurora.Green("OK").String()
		} else {
			PiNagivation = aurora.Red("NOK").String()
		}
		PiExploration := ""
		if psi.Programme.Exploration {
			PiExploration = aurora.Green("OK").String()
		} else {
			PiExploration = aurora.Red("NOK").String()
		}
		// Infos programme
		var PiHeader = []string{"Name", "Status", "Exploration", "Navigation", "Destination", "Temp arriver"}
		var PiData [][]string

		var timeDiff time.Duration
		if psi.Navigation {
			timeNow := time.Now()
			timeDiff = timeNow.Sub(psi.NavigationTimeArrived)
		}
		PiData = append(PiData, []string{
			aurora.Cyan(psi.Programme.Name).String(),
			PiStatus,
			PiExploration,
			PiNagivation,
			fmt.Sprintf("[ S %d- Z %d ]", psi.Programme.NextPosition.SecteurID, psi.Programme.NextPosition.ZoneID),
			aurora.Yellow(timeDiff.String()).String(),
		})
		PrintColorTable(PiHeader, PiData)
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

/*func PrintGridPosition(programme structure.Programme, size int) {
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
}*/
func PrintGridPosition(programme structure.Programme, size int) {
	var header []string
	var data [][]string
	for i := 0; i < size; i++ {
		header = append(header, fmt.Sprintf("- ZONE %d -", i))
	}
	var tmpData []string
	for j := 0; j < size; j++ {
		value := aurora.Red("0").String()
		if programme.Position.ZoneID == j {
			value = aurora.Green("1").String()
		}
		tmpData = append(tmpData, value)
	}
	data = append(data, tmpData)
	PrintColorTable(header, data, fmt.Sprintf(
		"SECTEUR [%d] ZONE [%d]",
		programme.Position.SecteurID,
		programme.Position.ZoneID,
	))
}
func PrintInfosGrille(infos structure.GridInfos) {
	var header = []string{"ID", "Taille", "ZoneTransfert", "Iteration", "Cycle", "NbrProgramme", "status", "Version", "FlagCapture"}
	var InfosTab [][]string

	flagCapture := aurora.Red("FALSE")
	if infos.FlagCapture {
		flagCapture = aurora.Green("TRUE")
	}
	statusGrid := aurora.Red("FALSE")
	if infos.Status {
		statusGrid = aurora.Green("TRUE")
	}
	InfosTab = append(InfosTab, []string{
		infos.Id,
		strconv.FormatInt(int64(infos.Taille), 10),
		fmt.Sprintf("S%d-Z%d", infos.ZoneTransfert.SecteurID, infos.ZoneTransfert.ZoneID),
		strconv.FormatInt(int64(infos.Iteration), 10),
		strconv.FormatInt(int64(infos.Cycle), 10),
		strconv.FormatInt(int64(infos.NbrProgrammes), 10),
		statusGrid.String(),
		infos.Version,
		flagCapture.String(),
	})
	PrintColorTable(header, InfosTab, "<---[ Infos grille ]--->")
	return
}
func PrintZoneInfos(infos structure.ZoneInfos) {
	var header = []string{"CELL ID", "VALEUR", "STATUS", "DATA", "DETECT TRAP"}
	var cellData [][]string
	for i := 0; i < len(infos.Cellules); i++ {
		statusCell := aurora.Red("FALSE")
		if infos.Cellules[i].Status {
			statusCell = aurora.Green("TRUE")
		}
		valeurCell := aurora.Red("0")
		if infos.Cellules[i].Valeur > 0 {
			valeurCell = aurora.Green(strconv.Itoa(infos.Cellules[i].Valeur))
		}
		safeCell := aurora.Green("SAFE")
		if infos.Cellules[i].Trapped {
			safeCell = aurora.Red("DANGER")
		}
		cellData = append(cellData, []string{
			fmt.Sprintf("%d", infos.Cellules[i].ID),
			valeurCell.String(),
			statusCell.String(),
			fmt.Sprintf("%d", infos.Cellules[i].DataCount),
			safeCell.String(),
		})
	}
	statusZone := aurora.Red("FALSE")
	if infos.Status {
		statusZone = aurora.Green("TRUE")
	}
	PrintColorTable(header, cellData, fmt.Sprintf("<---[ Infos programme sur Zone [%d] - status [%s]]--->", infos.ID, statusZone))
	header = []string{"Name", "Valeur total", "Energy total", "Status", "Exploration"}
	var progrData [][]string
	for i := 0; i < len(infos.Programmes); i++ {
		status := ""
		if infos.Programmes[i].Status {
			status = aurora.Green("OK").String()
		} else {
			status = aurora.Red("NOK").String()
		}
		exploration := ""
		if infos.Programmes[i].Exploration {
			exploration = aurora.Green("OK").String()
		} else {
			exploration = aurora.Red("NOK").String()
		}
		progrData = append(progrData, []string{
			infos.Programmes[i].Name,
			fmt.Sprintf("%d", infos.Programmes[i].ValeurTotal),
			fmt.Sprintf("%d", infos.Programmes[i].EnergyTotal),
			status,
			exploration,
		})
	}
	PrintColorTable(header, progrData, fmt.Sprintf("<---[ Infos Zone [%d] ]--->", infos.ID))
	return
}
func PrintExplore(celluleID string, data map[int]structure.CelluleData) {
	var header = []string{"ID", "ENERGY", "COMPETENCE", "IS_FLAG"}
	var cellData [][]string

	for i := 0; i < len(data); i++ {
		isFlag := aurora.Red("false")
		if data[i].IsFlag {
			isFlag = aurora.Green("true")
		}
		competence := aurora.Red("false")
		if data[i].Competence {
			competence = aurora.Green("true")
		}
		cellData = append(cellData, []string{
			fmt.Sprintf("%d", data[i].ID),
			fmt.Sprintf("%d", data[i].Energy),
			fmt.Sprintf("%s", competence),
			fmt.Sprintf("%s", isFlag),
		})
	}
	PrintColorTable(header, cellData, fmt.Sprintf("<---[ data cellule [%s] ]--->", celluleID))
	return
}
func PrintCelluleLogs(celluleLogs map[int]structure.CelluleLog) {
	var header = []string{"PID", "ACTIVE_CAPTURE", "C_TIME"}
	var cellData [][]string
	for _, log := range celluleLogs {
		activeCapture := aurora.Red("false")
		if log.ActiveCapture {
			activeCapture = aurora.Green("true")
		}
		cellData = append(cellData, []string{
			log.PID,
			activeCapture.String(),
			fmt.Sprintf("%s", log.CTime),
		})
	}
	PrintColorTable(header, cellData, fmt.Sprintf("<---[ Log cellule ]--->"))
	return
}
func PrintZoneActif(Zones []structure.ZonesGrid) {
	var header = []string{"Secteur", "Zone", "Status", "Actif", "Distance"}
	var dataList [][]string
	for _, zone := range Zones {
		status := ""
		if zone.Status {
			status = aurora.Green("OK").String()
		} else {
			status = aurora.Red("NOK").String()
		}
		actif := ""
		if zone.Actif {
			actif = aurora.Green("OK").String()
		} else {
			actif = aurora.Red("NOK").String()
		}
		dataList = append(dataList, []string{
			fmt.Sprintf("%d", zone.SecteurID),
			fmt.Sprintf("%d", zone.ZoneID),
			status,
			actif,
			fmt.Sprintf("%d", zone.Distance),
		})
	}
	if len(dataList) > 0 {
		PrintColorTable(header, dataList, "<---[ zone actif et status ( actif = programme sur zone | status true = objets sur zone ]--->")
	}
	return
}
