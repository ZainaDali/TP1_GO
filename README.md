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
go run cmd/main.go
```

2. Ajouter un utilisateur (avec flag)

```bash
go run . name=toto email=toto@mail.com
```

3. Voir les flags disponibles

```
go run . --help
```
