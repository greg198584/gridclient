# Apprendre à coder en jouant

## Grille Client CLI 

### Description 

https://codeurlibre.systeme.io/commencez--programmer-en-jouant--ce-nouveau-jeu

L'API fournit une grille de jeu avec des secteurs et des zones, et vous devrez créer des algorithmes pour faire avancer vos PAG ( Programme Api Grid ) dans la bonne direction.

Explorez les différentes régions, attaquez les autres programmes et capturez leurs ressources afin de progresser dans le jeu.

Trouvez le drapeau et transférez-le au programme principal pour gagner.

Le jeu est gagné par le premier joueur qui atteint l'objectif final avec son programme.

En cours de route, vous découvrirez des concepts de programmation tels que les variables, les boucles, les conditionnels et les fonctions.

### Règles

Le but du jeu est de trouver le drapeau caché et de détruire la zone de transfert, qui déconnectera tous les programmes du réseau et les empêchera d'économiser de l'énergie.

Les joueurs peuvent s'attaquer mutuellement.


- Force défense = energy quantity
- Force attaque = energy quantity
- Zone transfert = Zone de dépose de drapeau
- Trouver le drapeau cacher
- Le drapeau augmente le level + permet destruction zone de transfert et du cycle en cours
- Les cellules peuvent être piéger et une requête capture cause des dégâts dans ce cas la

### Concept

#### Le projet principal est l'API qui est en libre accès et sans authentification. 

- Vous pouvez utiliser le langage de votre choix et composer votre client comme vous le souhaitez. 
- Le client qui est inclus dans ce projet est un tutoriel pour utiliser l'API. 
- Vous avez tous les modes de jeu disponibles et un exemple de retour JSON par les routes. 
- Vous choisissez comment et quand vous voulez utiliser une route. 
- L'API vous fournit toutes les informations dont vous avez besoin pour jouer. 
- Et vous avez également dans les liens ci-dessus un tutoriel vidéo complet.

#### Mis a jour API

`v2.0.0` pour usage avec jeu en dev sur unity3D

taille zone = 150 000 000 km
deplacement vitesse par default = 300 000 km/s * 4

1. Simplification (mode verouillage de zone supprimer)
2. Sauvegarde supprimer (le flag permet de sauvegarder est passer au level suivant)
3. Attaque possible sur cellule target en mode exploration (surtout pour mode avec unity)
4. Mode deplacement et cout. ( distance vitesse de deplacement )
5. ajout infos distance de zone dans grid status.
6. suppression infos programme ( seulement sur scan zone )
7. Ajout estimation distance temp trajet et cout energy et iteration.

`v1.11.2`   

1. Destruction des cellules piégées pour capture sans risque.
2. Mode verrouillage de zone (emprisonnements du programme).           
3. La lock zone peut etre egalement déverrouiller par force brute de mot de passe (parametre founi mode random, mais simple)

`v1.9.10`

Cellules piège ( capture cause des degat ) | zone de transfert = zone de sauvegarde | zone de transfert pour celui qui capture le drapeau.

`v1.0.0`    

Déplacement | capture de drapeau | combat | scan | explore | capture | infos programmes | infos grille

### Client GO Installation

installer go https://go.dev/doc/install

```bash
go install github.com/greg198584/gridclient@latest
```
### Tuto complet en vidéo

- https://youtu.be/DlN74mHg0bw

### API

- API PROJET https://gitlab.com/greg198584/ApiGrid

- Aucune base de donnée est utiliser par l'API. 
- 
- Il suffit de garder le JSON représenter par la structure ProgrammeContainer signification PAG ( Programme Api Grid )

[v1_programme_new](_exemple_retour_json/v1_programme_new.json).

```go
type ProgrammeContainer struct {
	ID        string    `json:"id"`
	SecretID  string    `json:"secret_id"`
	Programme Programme `json:"programme"`
	ValidKey  string    `json:"valid_key"`
}
```

**JSON obtenu par deux routes**

```bash
GET /v1/programme/new/:name
GET /v1/programme/push/flag/:id/:secretid
```

#### Routes API

```bash
GET /v1/grid
GET /v1/programme/infos/:id/:secretid
GET /v1/programme/save/:id/:secretid
GET /v1/programme/new/:name
POST /v1/programme/load
POST /v1/programme/upgrade
GET /v1/programme/unset/:id/:secretid
GET /v1/programme/move/:id/:secretid/:zone_id
GET /v1/programme/scan/:id/:secretid
GET /v1/programme/cellule/log/:id/:secretid/:celluleid
GET /v1/programme/explore/:id/:secretid/:celluleid
GET /v1/programme/destroy/:id/:secretid/:celluleid/:targetid
GET /v1/programme/rebuild/:id/:secretid/:celluleid/:targetid
GET /v1/programme/capture/cellule/data/:id/:secretid/:celluleid/:index
GET /v1/programme/capture/cellule/energy/:id/:secretid/:celluleid/:index
GET /v1/programme/capture/target/data/:id/:secretid/:celluleid/:targetid
GET /v1/programme/capture/target/energy/:id/:secretid/:celluleid/:targetid
GET /v1/programme/equilibrium/:id/:secretid
GET /v1/programme/push/flag/:id/:secretid
GET v1/programme/destroy/zone/:id/:secretid/:celluleid/:energy
GET v1/programme/navigation/stop/:id/:secretid
GET v1/programme/exploration/stop/:id/:secretid
GET v1/programme/stop/move/:id/:secretid/:secteur_id/:zone_id
GET v1/programme/estimate/move/:id/:secretid/:secteur_id/:zone_id
```

#### Route Destruction cellules de zone (zone de transfert destruction avec flag seulement)

```bash
GET /v1/programme/destroy/zone/:id/:secretid/:celluleid
```

-La route `/v1/programme/destroy/zone/` est un moyen sûr de déclencher une réinitialisation du réseau et d'empêcher tout programme d'économiser de l'énergie à la fin d'un cycle (si destruction zone de transfert).

### Exemple retour API par route

[v1_grid](_exemple_retour_json/v1_grid.json).

[v1_programme_capture_cellule_data](_exemple_retour_json/v1_programme_capture_cellule_data.json).

[v1_programme_capture_cellule_energy](_exemple_retour_json/v1_programme_capture_cellule_energy.json).

[v1_programme_capture_target_data](_exemple_retour_json/v1_programme_capture_target_data.json).

[v1_programme_capture_target_energy](_exemple_retour_json/v1_programme_capture_target_energy.json).

[v1_programme_cellule_log](_exemple_retour_json/v1_programme_cellule_log.json).

[v1_programme_destroy](_exemple_retour_json/v1_programme_destroy.json).

[v1_programme_equilibrium](_exemple_retour_json/v1_programme_equilibrium.json).

[v1_programme_explore](_exemple_retour_json/v1_programme_explore.json).

[v1_programme_infos](_exemple_retour_json/v1_programme_infos.json).

[v1_programme_jump_down](_exemple_retour_json/v1_programme_jump_down.json).

[v1_programme_jump_up](_exemple_retour_json/v1_programme_jump_up.json).

[v1_programme_load](_exemple_retour_json/v1_programme_load.json).

[v1_programme_move](_exemple_retour_json/v1_programme_move.json).

[v1_programme_new](_exemple_retour_json/v1_programme_new.json).

[v1_programme_push_flag](_exemple_retour_json/v1_programme_push_flag.json).

[v1_programme_rebuild](_exemple_retour_json/v1_programme_rebuild.json).

[v1_programme_save](_exemple_retour_json/v1_programme_save.json).

[v1_programme_scan](_exemple_retour_json/v1_programme_scan.json).

[v1_programme_unset](_exemple_retour_json/v1_programme_unset.json).

[v1_programme_upgrade](_exemple_retour_json/v1_programme_upgrade.json).

[v1_programme_zone_actif](_exemple_retour_json/v1_programme_zone_actif.json).

[v1_programme_destroy_zone](_exemple_retour_json/v1_programme_destroy_zone.json).

### Usage 

```bash
> $ go run main.go                                                                                                                                         
Usage: main [OPTIONS] COMMAND [arg...]

Concepteur Console

Options:
  -v, --version      Show the version and exit

Commands:
  create             creation programme et chargement sur la grille
  save               sauvegarde programme
  load               charger programme existant sur la grille
  upgrade            mis a jour programme
  delete             deconnecter un programme de la grille
  move               deplacer un programme sur la grille
  scan               scan infos de la zone pour
  explore            exploration de cellule de zone
  destroy            destroy cellule programme
  rebuild            reconstruire cellule programme
  capture            capture data-energy cellule programme et zone
  equilibrium        repartir energie du programme uniformement
  pushflag           push drapeau dans zone de transfert
  status             status grille
  infos              infos programme
  attack             mode attaque - tous programme dans la zone
  defense            mode defense + attaque simultanement programme hostile
  quick_move         deplacement secteur + zone voulu
  search_flag        current + scan + explore (all zone secteur current) + capture >>> FLAG
  search_energy      scan + explore + capture >>> ENERGY
  search_programme   recherche programme
  search_trap        recherche cellule trap (cellule dangereuse)
  monitoring         position + status programme monitoring
  log                info log cellule
  destroy_zone       destroy cellule zone current (avec flag)
  zone_actif         affiche zone actif + programme sur zone (avec flag)

Run 'main COMMAND --help' for more information on a command.
```

### Exemple usage

#### Ajouter un programme

- create

```bash
curl --request GET \
  --url http://195.154.84.18:20080/v1/programme/new/BBB \
  --header 'content-type: application/json'
```

```json
{
	"id": "7378faf7afa4ccbc9c8ec25fdbb9ad62de4d21bc",
	"secret_id": "7d4b71be1a053f7b32fed407885ecff2830b4b44",
	"programme": {
		"id": "7378faf7afa4ccbc9c8ec25fdbb9ad62de4d21bc",
		"name": "BBB",
		"position": {
			"secteur_id": 0,
			"zone_id": 0
		},
		"cellules": {
			"0": {
				"id": 0,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"1": {
				"id": 1,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"2": {
				"id": 2,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"3": {
				"id": 3,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"4": {
				"id": 4,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"5": {
				"id": 5,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"6": {
				"id": 6,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"7": {
				"id": 7,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"8": {
				"id": 8,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			}
		},
		"status": true
	},
	"valid_key": "$2a$14$DB6BXPvYvLPZVvgr0UCeE.DmehETOOtayCh.dU1qaCYfVPqXWPbHa"
}
```

- load ( utiliser le json recuperer lors du create )

```bash
curl --request POST \
  --url http://195.154.84.18:20080/v1/programme/load \
  --header 'content-type: application/json' \
  --data '{
	"id": "7378faf7afa4ccbc9c8ec25fdbb9ad62de4d21bc",
	"secret_id": "7d4b71be1a053f7b32fed407885ecff2830b4b44",
	"programme": {
		"id": "7378faf7afa4ccbc9c8ec25fdbb9ad62de4d21bc",
		"name": "BBB",
		"position": {
			"secteur_id": 0,
			"zone_id": 0
		},
		"cellules": {
			"0": {
				"id": 0,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"1": {
				"id": 1,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"2": {
				"id": 2,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"3": {
				"id": 3,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"4": {
				"id": 4,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"5": {
				"id": 5,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"6": {
				"id": 6,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"7": {
				"id": 7,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			},
			"8": {
				"id": 8,
				"valeur": 10,
				"energy": 10,
				"datas": null,
				"status": true,
				"destroy": true,
				"rebuild": true,
				"current_acces_log": {
					"pid": "",
					"target_id": "",
					"receive_destroy": false,
					"active_destroy": false,
					"active_rebuild": false,
					"active_capture": false,
					"c_time": "0001-01-01T00:00:00Z"
				},
				"acces_log": null,
				"capture": false
			}
		},
		"status": true
	},
	"valid_key": "$2a$14$DB6BXPvYvLPZVvgr0UCeE.DmehETOOtayCh.dU1qaCYfVPqXWPbHa"
}'
```
### AG-1 monitoring

#### Get Started

```bash
> $ cd ag1-monitoring
> $ python3 -m venv .env 
> $ source .env/bin/activate
> $ pip install -r requirements.txt
```

```bash
> $ python3 main.py --id your_id --secret-id your_secret_id --host 195.154.84.18:20080
L'ID du programme est 69a2e0368897ff
Le Secret ID est b73c621bf0b699384d9
host: 195.154.84.18:20080
```

## Screen exemple

![alt text](img/moni_demo.png)
