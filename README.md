## Grille Client CLI

### Description

https://codeurlibre.systeme.io/commencez--programmer-en-jouant--ce-nouveau-jeu

L'API fournit une grille de jeu avec des secteurs et des zones, et vous devrez créer des algorithmes pour faire avancer vos programmes dans la bonne direction.

Explorez les différentes régions, attaquez les autres programmes et capturez leurs ressources afin de progresser dans le jeu.

Trouvez le drapeau et transférez-le au programme principal pour gagner.

Le jeu est gagné par le premier joueur qui atteint l'objectif final avec son programme.

En cours de route, vous découvrirez des concepts de programmation tels que les variables, les boucles, les conditionnels et les fonctions.

### Installation

installer go https://go.dev/doc/install

```bash
go get gitlab.com/greg198584/gridclient
```

### route API

- API BETA: http://195.154.84.18:20080/v1/grid

```bash
GET /v1/grid
GET /v1/programme/infos/:id/:secretid
GET /v1/programme/new/:name
POST /v1/programme/load
GET /v1/programme/unset/:id/:secretid
GET /v1/programme/save/:id/:secretid
GET /v1/programme/jump/up/:id/:secretid/:jumpnbr
GET /v1/programme/jump/down/:id/:secretid/:jumpnbr
GET /v1/programme/move/:id/:secretid/:zone_id
GET /v1/programme/scan/:id/:secretid
GET /v1/programme/explore/:id/:secretid/:celluleid/:startid/:endid
GET /v1/programme/destroy/:id/:secretid/:celluleid/:targetid
GET /v1/programme/rebuild/:id/:secretid/:celluleid/:targetid
```

### Usage 

```bash
> $ go run main.go                                                                                                                                               [±main ●●]

Usage: main [OPTIONS] COMMAND [arg...]

Concepteur Console
                  
Options:          
  -v, --version   Show the version and exit
                  
Commands:         
  create          creation programme et chargement sur la grille
  load            charger programme existant sur la grille
  save            sauvegarder un programme
  delete          deconnecter un programme de la grille
  move            deplacer un programme sur la grille
  scan            scan la zone pour get informations (cellules infos / programmes present et infos)
  explore         exploration de cellule de zone
  destroy         destroy cellule programme
  rebuild         reconstruire cellule programme
  status          status grille
  infos           infos programme
  algo            lancer un algo
                  
Run 'main COMMAND --help' for more information on a command.
```

### Exemple usage

#### Ajouter un programme

- create

```bash
curl --request GET \
  --url http://localhost/v1/programme/new/BBB \
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
  --url http://localhost/v1/programme/load \
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

#### Deplacer un programme

- Move (changer de zone)

```bash
curl --request GET \
  --url http://localhost/v1/programme/move/7378faf7afa4ccbc9c8ec25fdbb9ad62de4d21bc/7d4b71be1a053f7b32fed407885ecff2830b4b44/1 \
  --header 'content-type: application/json'
```

- JumpDown (descendre de secteur de 0 vers taille max )

```bash
curl --request GET \
  --url http://localhost/v1/programme/jump/down/7378faf7afa4ccbc9c8ec25fdbb9ad62de4d21bc/7d4b71be1a053f7b32fed407885ecff2830b4b44/1 \
  --header 'content-type: application/json'
```

- JumpUp (monter de secteur de taille max vers 0)
```bash
curl --request GET \
  --url http://localhost/v1/programme/jump/up/7378faf7afa4ccbc9c8ec25fdbb9ad62de4d21bc/7d4b71be1a053f7b32fed407885ecff2830b4b44/1 \
  --header 'content-type: application/json'
```
