# TP_GO — Mini CRM en ligne de commande

## Prérequis

- **Go ≥ 1.20** installé sur la machine

Vérification :

```bash
go version
```

## Description

Application CLI écrite en Go permettant de gérer un carnet de contacts.  
Fonctionnalités :

- Ajouter un contact
- Lister les contacts
- Supprimer un contact
- Mettre à jour un contact

## Structure du projet

```
TP1_GO/
├── cmd/
│   └── main.go               # Point d’entrée du programme
├── internal/
│   └── app/
│       └── app.go            # Logique principale de l’application
│   └── storage/
│       └── memory.go         # Implémentation en mémoire du stockage des contacts
│       └── storage.go        # Interface définissant les opérations de stockage
│       └── json.go           # Stockage dans un fichier json
```

## Installation

1. Cloner le dépôt :

```bash
git clone https://github.com/ZainaDali/TP1_GO.git
cd TP1_GO
```

## Utilisation

1. Lancer le programme

```bash
go build -o crm .
```

2. Voir les utilisateurs

```bash
./crm list
```

3. Ajouter un utilisateur

```bash
./crm add
```

(avec flag)

```bash
./crm add -n=toto -e=toto@mail.com
```

4. Modifier un utilisateur

```bash
./crm update -i 1
```

4. Ajouter un utilisateur (avec CLI)

```bash
go run . add -n "Alice" -e "alice@mail.com"

go run . update -i 1 -n "Alice Dubois"

go run . delete -i 2

go run . list
```

5. Voir les flags disponibles

```bash
go run . --help
```
