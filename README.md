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
├── main.go # Point d’entrée du programme
├── contact/
│ └── contact.go # Gestion des contacts (ajout, suppression, affichage)
└── menu/
└── menu.go # Affichage du menu principal
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
go run .
```
