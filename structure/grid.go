package structure

import "time"

type ProgrammeContainer struct {
	ID        string    `json:"id"`
	SecretID  string    `json:"secret_id"`
	Programme Programme `json:"programme"`
	ValidKey  string    `json:"valid_key"`
}
type GridInfos struct {
	Id            string           `json:"id"`
	Taille        int              `json:"taille"`
	ZoneTransfert Position         `json:"zone_transfert"`
	Iteration     int              `json:"iteration"`
	Cycle         int              `json:"cycle"`
	NbrProgrammes int              `json:"nbr_programmes"`
	Programmes    []ProgrammeInfos `json:"programmes"`
	Status        bool             `json:"status"`
	Version       string           `json:"version"`
	FlagCapture   bool             `json:"flag_capture"`
	IndiceFlag    IndiceFlag       `json:"indice_flag"`
}
type IndiceFlag struct {
	IndiceValue string `json:"indice_value"`
	Instruction string `json:"instruction"`
}
type ProgrammeStatusInfos struct {
	Programme     Programme            `json:"programme"`
	LockProgramme map[string]Programme `json:"lock_programme"`
}
type ZoneInfos struct {
	ID         int              `json:"id"`
	Actif      bool             `json:"actif"`
	Cellules   []CelluleInfos   `json:"cellule"`
	Programmes []ProgrammeInfos `json:"programmes"`
}
type ProgrammeInfos struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	NbrCellules int    `json:"nbr_cellules"`
	ValeurTotal int    `json:"valeur_total"`
	EnergyTotal int    `json:"energy_total"`
	Status      bool   `json:"status"`
}
type CelluleInfos struct {
	ID        int         `json:"id"`
	Valeur    int         `json:"valeur"`
	Energy    int         `json:"energy"`
	DataType  CelluleData `json:"data_type"`
	DataCount int         `json:"data_count"`
	Status    bool        `json:"status"`
	Destroy   bool        `json:"destroy"`
	Rebuild   bool        `json:"rebuild"`
	Capture   bool        `json:"capture"`
}
type Programme struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Position  Position         `json:"position"`
	Cellules  map[int]*Cellule `json:"cellules"`
	Level     int              `json:"level"`
	GridFlags []string         `json:"grid_flags"`
	Status    bool             `json:"status"`
}
type Position struct {
	SecteurID int `json:"secteur_id"`
	ZoneID    int `json:"zone_id"`
}
type Cellule struct {
	ID              int                 `json:"id"`
	Valeur          int                 `json:"valeur"`
	Energy          int                 `json:"energy"`
	Datas           map[int]CelluleData `json:"datas"`
	Status          bool                `json:"status"`
	Destroy         bool                `json:"destroy"`
	Rebuild         bool                `json:"rebuild"`
	CurrentAccesLog CelluleLog          `json:"current_acces_log"`
	AccesLog        []CelluleLog        `json:"acces_log"`
	Capture         bool                `json:"capture"`
}
type CelluleLog struct {
	PID            string    `json:"pid"`
	TargetID       string    `json:"target_id"`
	ReceiveDestroy bool      `json:"receive_destroy"`
	ActiveDestroy  bool      `json:"active_destroy"`
	ActiveRebuild  bool      `json:"active_rebuild"`
	ActiveCapture  bool      `json:"active_capture"`
	CTime          time.Time `json:"c_time"`
}
type CelluleData struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Energy  int    `json:"energy"`
	IsFlag  bool   `json:"is_flag"`
}
