package main

import (
	"github.com/greg198584/gridclient/programme"
	mowcli "github.com/jawher/mow.cli"
	"os"
	"strconv"
)

func main() {
	app := mowcli.App("main", "Concepteur Console")
	app.Version("v version", "0.1")

	app.Command("create", "creation programme et chargement sur la grille", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.New(*pname)
		}
	})
	//app.Command("save", "sauvegarde programme", func(cmd *mowcli.Cmd) {
	//	var (
	//		pname = cmd.StringOpt("n name", "", "nom du programme")
	//	)
	//	cmd.Action = func() {
	//		programme.Save(*pname)
	//	}
	//})
	app.Command("load", "charger programme existant sur la grille", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Load(*pname)
		}
	})
	app.Command("upgrade", "mis a jour programme", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Upgrade(*pname)
		}
	})
	app.Command("disconnect", "deconnecter un programme de la grille", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Delete(*pname)
		}
	})
	app.Command("move", "deplacer un programme sur la grille", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			secteurID = cmd.StringOpt("s secteur", "", "ID du secteur")
			zoneID    = cmd.StringOpt("z zone", "", "ID zone")
		)
		cmd.Action = func() {
			programme.MovePosition(*pname, *secteurID, *zoneID)
		}
	})
	app.Command("estimate_move", "estimation temp deplacement sur zone", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			secteurID = cmd.StringOpt("s secteur", "", "ID du secteur")
			zoneID    = cmd.StringOpt("z zone", "", "ID zone")
		)
		cmd.Action = func() {
			programme.EstimateMove(*pname, *secteurID, *zoneID)
		}
	})
	app.Command("stop_move", "stopper navigation en cours ( retour zone de depart )", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.StopMove(*pname)
		}
	})
	app.Command("scan", "scan infos de la zone pour", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Scan(*pname)
		}
	})
	app.Command("explore", "exploration de cellule de zone", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			celluleID = cmd.StringOpt("c cellule", "", "ID cellule")
		)
		cmd.Action = func() {
			programme.Explore(*pname, *celluleID)
		}
	})
	app.Command("destroy", "destroy cellule programme", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			celluleID = cmd.StringOpt("c cellule", "", "ID cellule")
			targetID  = cmd.StringOpt("t target", "", "ID programme cible")
			energy    = cmd.StringOpt("e energy", "", "quantiter energy a utiliser")
		)
		cmd.Action = func() {
			CelluleID, _ := strconv.Atoi(*celluleID)
			Energy, _ := strconv.Atoi(*energy)
			programme.Destroy(*pname, CelluleID, *targetID, Energy)
		}
	})
	app.Command("rebuild", "reconstruire cellule programme", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			celluleID = cmd.StringOpt("c cellule", "", "ID cellule")
			targetID  = cmd.StringOpt("t target", "", "ID programme cible")
			energy    = cmd.StringOpt("e energy", "", "quantiter energy a utiliser")
		)
		cmd.Action = func() {
			CelluleID, _ := strconv.Atoi(*celluleID)
			Energy, _ := strconv.Atoi(*energy)
			programme.Rebuild(*pname, CelluleID, *targetID, Energy)
		}
	})
	app.Command("capture", "capture data-energy cellule programme et zone", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			celluleID = cmd.StringOpt("c cellule", "", "ID cellule")
			target    = cmd.StringOpt("t target", "", "cible [cell-pid]")
			mode      = cmd.StringOpt("m mode", "", "mode [data-energy]")
			id        = cmd.StringOpt("i id", "", "index cellule ou pid - ou index multiple [id_debut-id_fin] ex (12 10-20)")
		)
		cmd.Action = func() {
			CelluleID, _ := strconv.Atoi(*celluleID)
			switch *mode {
			case "data":
				if *target == "pid" {
					programme.CaptureTargetData(*pname, CelluleID, *id)
				} else {
					//index, _ := strconv.Atoi(*id)
					programme.CaptureCellData(*pname, CelluleID, *id)
				}
				break
			case "energy":
				if *target == "pid" {
					programme.CaptureTargetEnergy(*pname, CelluleID, *id)
				} else {
					//index, _ := strconv.Atoi(*id)
					programme.CaptureCellEnergy(*pname, CelluleID, *id)
				}
				break
			default:
			}
		}
	})
	app.Command("equilibrium", "repartir energie du programme uniformement", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Equilibrium(*pname)
		}
	})
	app.Command("pushflag", "push drapeau dans zone de transfert", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.PushFlag(*pname)
		}
	})
	app.Command("status", "status grille", func(cmd *mowcli.Cmd) {
		var (
			zoneActif = cmd.BoolOpt("a actif", false, "afficher que les zone status true")
		)
		cmd.Action = func() {
			programme.GetStatusGrid(*zoneActif)
		}
	})
	app.Command("infos", "infos programme", func(cmd *mowcli.Cmd) {
		var (
			pname         = cmd.StringOpt("n name", "", "nom du programme")
			printPosition = cmd.BoolOpt("p position", false, "afficher position")
		)
		cmd.Action = func() {
			programme.GetInfoProgramme(*pname, *printPosition)
		}
	})
	app.Command("navigation_stop", "stop mode navigation", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Navigation(*pname)
		}
	})
	app.Command("exploration_stop", "stop exploration", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.ExplorationStop(*pname)
		}
	})
	app.Command("monitoring", "position + status programme monitoring", func(cmd *mowcli.Cmd) {
		var (
			pname         = cmd.StringOpt("n name", "", "nom du programme")
			printPosition = cmd.BoolOpt("p position", false, "afficher position")
		)
		cmd.Action = func() {
			programme.Monitoring(*pname, *printPosition)
		}
	})
	app.Command("log", "info log cellule", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			celluleID = cmd.StringOpt("c cellule", "", "ID cellule")
		)
		cmd.Action = func() {
			programme.GetCelluleLog(*pname, *celluleID)
		}
	})
	app.Command("clean_log", "clean log cellule", func(cmd *mowcli.Cmd) {
		var (
			pname     = cmd.StringOpt("n name", "", "nom du programme")
			celluleID = cmd.StringOpt("c cellule", "", "ID cellule")
		)
		cmd.Action = func() {
			programme.CleanCelluleLog(*pname, *celluleID)
		}
	})
	app.Command("destroy_zone", "destroy cellule zone current", func(cmd *mowcli.Cmd) {
		var (
			pname      = cmd.StringOpt("n name", "", "nom du programme")
			celluleID  = cmd.StringOpt("c cellule", "", "ID cellule")
			energy     = cmd.StringOpt("e energy", "", "quantite energy utiliser par cellule")
			allCellule = cmd.BoolOpt("a all", false, "toutes les cellules")
		)
		cmd.Action = func() {
			celluleIDint, _ := strconv.Atoi(*celluleID)
			energyint, _ := strconv.Atoi(*energy)
			programme.DestroyZone(*pname, celluleIDint, energyint, *allCellule)
		}
	})
	app.Action = func() {
		app.PrintHelp()
	}
	err := app.Run(os.Args)
	if err != nil {
		app.PrintHelp()
	}
	os.Exit(0)
}
