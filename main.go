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
	app.Command("delete", "deconnecter un programme de la grille", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Delete(*pname)
		}
	})
	app.Command("move", "deplacer un programme sur la grille", func(cmd *mowcli.Cmd) {
		var (
			pname    = cmd.StringOpt("n name", "", "nom du programme")
			typemove = cmd.StringOpt("t type", "", "type move (jumpup, jumpdown, move)")
			valeur   = cmd.StringOpt("v valeur", "", "valeur entre 0 et taille max (9)")
		)
		cmd.Action = func() {
			switch *typemove {
			case "jumpup":
				programme.JumpUp(*pname, *valeur)
				break
			case "jumpdown":
				programme.JumpDown(*pname, *valeur)
				break
			case "move":
				programme.Move(*pname, *valeur)
				break
			default:
				break
			}
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
			energy    = cmd.StringOpt("e energy", "", "nombre energy a utiliser")
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
			energy    = cmd.StringOpt("e energy", "", "nombre energy a utiliser")
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
			id        = cmd.StringOpt("i id", "", "index cellule ou pid")
		)
		cmd.Action = func() {
			CelluleID, _ := strconv.Atoi(*celluleID)
			switch *mode {
			case "data":
				if *target == "pid" {
					programme.CaptureTargetData(*pname, CelluleID, *id)
				} else {
					index, _ := strconv.Atoi(*id)
					programme.CaptureCellData(*pname, CelluleID, index)
				}
				break
			case "energy":
				if *target == "pid" {
					programme.CaptureTargetEnergy(*pname, CelluleID, *id)
				} else {
					index, _ := strconv.Atoi(*id)
					programme.CaptureCellEnergy(*pname, CelluleID, index)
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
		cmd.Action = func() {
			programme.GetStatusGrid()
		}
	})
	app.Command("infos", "infos programme", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.GetInfoProgramme(*pname)
		}
	})
	app.Command("attack", "mode attaque - tous programme dans la zone", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.Attack(*pname)
		}
	})
	app.Command("defense", "mode defense + attaque simultanement programme hostile", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.CheckAttack(*pname)
		}
	})
	app.Command("quick_move", "deplacement secteur + zone voulu", func(cmd *mowcli.Cmd) {
		var (
			pname    = cmd.StringOpt("n name", "", "nom du programme")
			position = cmd.StringOpt("p position", "", "secteur-zone exemple: [7-2]")
		)
		cmd.Action = func() {
			programme.MovePosition(*pname, *position)
		}
	})
	app.Command("search_flag", "current + scan + explore (all zone secteur current) + capture >>> FLAG", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.SearchFlag(*pname)
		}
	})
	app.Command("search_energy", "scan + explore + capture >>> ENERGY", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
		)
		cmd.Action = func() {
			programme.SearchEnergy(*pname)
		}
	})
	app.Command("search_programme", "recherche programme", func(cmd *mowcli.Cmd) {
		var (
			pname = cmd.StringOpt("n name", "", "nom du programme")
			all   = cmd.BoolOpt("a all", false, "toutes la grille")
		)
		cmd.Action = func() {
			programme.SearchProgramme(*pname, *all)
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
	app.Action = func() {
		app.PrintHelp()
	}
	err := app.Run(os.Args)
	if err != nil {
		app.PrintHelp()
	}
	os.Exit(0)
}
