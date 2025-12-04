## TP_GO — Mini CRM en ligne de commande

### Prérequis

- **Go ≥ 1.20** installé sur la machine

Vérification :

```bash
go version
```

### Description

Application CLI écrite en Go permettant de gérer un carnet de contacts.  
Fonctionnalités :

- **Ajouter** un contact
- **Lister** les contacts
- **Supprimer** un contact
- **Mettre à jour** un contact

Le stockage des contacts est **interchangeable** : fichier JSON ou base SQLite via GORM.

### Structure du projet

```text
TP1_GO/
├── cmd/
│   ├── root.go        # Commande racine, sélection du type de stockage
│   ├── add.go         # Sous-commande: ajout de contact
│   ├── list.go        # Sous-commande: liste des contacts
│   ├── update.go      # Sous-commande: mise à jour
│   └── delete.go      # Sous-commande: suppression
├── internal/
│   ├── app/
│   │   └── app.go     # Ancienne logique de menu interactif (non utilisée par Cobra)
│   ├── config/
│   │   └── config.go  # Configuration de l’application (BDD, etc.)
│   └── storage/
│       ├── storage.go # Interface Storer + modèle Contact
│       ├── json.go    # Implémentation JSON (fichier users.json)
│       └── gorm.go    # Implémentation GORM (SQLite)
├── users.json         # Fichier de stockage JSON (si utilisé)
└── main.go            # Point d’entrée qui appelle cmd.Execute()
```

### Installation

1. Cloner le dépôt :

```bash
git clone https://github.com/ZainaDali/TP1_GO.git
cd TP1_GO
```

2. (Optionnel) Récupérer les dépendances :

```bash
go mod tidy
```

3. Construire le binaire :

```bash
go build -o crm .
```

### Utilisation de base (stockage par défaut : SQLite / GORM)

Par défaut, le stockage est fait en **SQLite** (fichier `crm_users.db`) via GORM.

- **Lister les contacts** :

```bash
./crm list
```

- **Ajouter un contact (mode interactif)** :

```bash
./crm add
```

- **Ajouter un contact avec flags** :

```bash
./crm add -n "Toto" -e "toto@mail.com"
```

- **Mettre à jour un contact** :

```bash
./crm update -i 1 -n "Nouveau Nom" -e "nouvel@mail.com"
```

- **Supprimer un contact** :

```bash
./crm delete -i 1
```

### Choisir le type de stockage avec le flag `--storage` / `-s`

Le type de stockage se choisit **à l’exécution**, grâce au flag global :

- **SQLite / GORM (par défaut)** :

```bash
./crm list                    # équivalent à: ./crm list -s gorm
./crm add -n "Toto" -e "toto@mail.com"
./crm add -s gorm -n "Toto" -e "toto@mail.com"
```

- **JSON (fichier `users.json`)** :

```bash
./crm list -s json
./crm add -s json -n "Toto" -e "toto@mail.com"
./crm update -s json -i 1 -n "Tata"
./crm delete -s json -i 1
```

Le flag `--storage` est **persistant** (global) : tu peux aussi écrire :

```bash
./crm -s json list
./crm -s json add -n "Alice" -e "alice@mail.com"
```

### Aide et documentation des commandes

Afficher l’aide générale :

```bash
./crm --help
```

Afficher l’aide d’une sous-commande :

```bash
./crm add --help
./crm list --help
./crm update --help
./crm delete --help
```
