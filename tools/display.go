package tools

import (
	"github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func PrintColorTable(header []string, dataList [][]string, title_opt ...string) {
	if len(title_opt) == 1 {
		log.Printf("[%s] %s", aurora.Magenta("---"), aurora.Cyan("[ "+title_opt[0]+" ]"))
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, data := range dataList {
		var dataColor []tablewriter.Colors
		for _, dataValue := range data {
			if dataValue == "OK" {
				dataColor = append(dataColor, tablewriter.Colors{
					tablewriter.Bold,
					tablewriter.FgGreenColor,
				})
			} else if dataValue == "NOK" {
				dataColor = append(dataColor, tablewriter.Colors{
					tablewriter.Bold,
					tablewriter.FgRedColor,
				})
			}
		}
		table.Rich(data, dataColor)
	}
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}

func PrintColorTableNoBorder(header []string, dataList [][]string, title_opt ...string) {
	if len(title_opt) == 1 {
		log.Printf("%s %s", aurora.Blue(">>>"), aurora.Cyan("[ "+title_opt[0]+" ]"))
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, data := range dataList {
		var dataColor []tablewriter.Colors
		for _, dataValue := range data {
			if dataValue == "OK" || dataValue == "YES" || dataValue == "+" {
				dataColor = append(dataColor, tablewriter.Colors{
					tablewriter.Bold,
					tablewriter.FgGreenColor,
				})
			} else if dataValue == "FAIL" || dataValue == "NO" || dataValue == "-" {
				dataColor = append(dataColor, tablewriter.Colors{
					tablewriter.Bold,
					tablewriter.FgRedColor,
				})
			} else {
				dataColor = append(dataColor, tablewriter.Colors{})
			}
		}
		table.Rich(data, dataColor)
	}
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.Render()
}

func Title(title string) {
	log.Printf("[%s]", aurora.Magenta(title))
}

func Info(message string, tab ...bool) {
	isTab := ""
	if len(tab) > 0 {
		if tab[0] {
			isTab = "\t"
		}
	}
	log.Printf("%s[%s] [%s]", isTab, aurora.Blue(">>>"), aurora.Cyan(message))
}

func Success(message string, tab ...bool) {
	isTab := ""
	if len(tab) > 0 {
		if tab[0] {
			isTab = "\t"
		}
	}
	log.Printf("%s[%s] [%s] [%s]", isTab, aurora.Green("+"), aurora.Yellow(message), aurora.Green("OK"))
}

func Warning(message string, tab ...bool) {
	isTab := ""
	if len(tab) > 0 {
		if tab[0] {
			isTab = "\t"
		}
	}
	log.Printf("%s[%s] [%s]", isTab, aurora.Yellow("***"), aurora.White(message).Bold())
}

func Fail(message string, tab ...bool) {
	isTab := ""
	if len(tab) > 0 {
		if tab[0] {
			isTab = "\t"
		}
	}
	log.Printf("%s[%s] [%s] [%s]", isTab, aurora.Red("-"), aurora.Yellow(message), aurora.Red("FAIL"))
}

func Error(message string) {
	log.Printf("\t[%s] [%s]", aurora.Red("X"), aurora.Red(message))
}

func Log(message string, logData string, tab ...bool) {
	isTab := ""
	if len(tab) > 0 {
		if tab[0] {
			isTab = "\t"
		}
	}
	log.Printf("%s[%s] [%s] (%s)", isTab, aurora.Yellow("LOG"), aurora.Yellow(message), aurora.Yellow(logData))
}
