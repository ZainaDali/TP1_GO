# Mini-CRM : Étapes finales 

Vous arrivez sur les dernières étapes de ce projet fil rouge : toutes mes félicitations ! 
À l'heure actuelle, vous avez déjà réalisée ces étapes : 
* **Étape 1** : Programme simple sans persistance de données (via des map) 
* **Étape 2** : 
  * 2.1 : Ajout des struct pour modéliser les données de manière plus sûre 
  * 2.2 : Ajout des Interfaces pour avoir un programme modulaire, plus facile à faire évoluer et à maintenir
  
Malheureusement, vous n'êtes pas au bout de vos peines puisque 3 dernières étapes vous attendent pour finaliser cette transformation et en faire un véritable outil en ligne de commande, robuste et configurable.

## Étapes 3

### Étape 3.1 : CLI & Cobra

* Le but de cette étape et de transformer votre programme en vraie CLI grâce à Cobra. Voici quelques indications : 

1. Ajouter les dépendances nécessaires :
```bash
    go get -u github.com/spf13/cobra@latest
```
2. Réorganiser les projets : Adoptez une structure de projet orientée Cobra
3. Implémenter la commande Racine (`cmd/root.go`)
4. Implémenter les sous-commandes : 
    * Pour chaque fonctionnalité (ajouter, lister, mettre à jour, supprimer), créez un fichier .go dédié dans le package cmd.


**Notes pour l'ajout de contact** : Je vous propose ici de garder l'ajout de contact interactif (donc quand on veut créer un contact) le programme nous demande le nom puis le mail du contact (et autres champs si besoin), et de créer des flags (non obligatoires du coup) pour pouvoir créer des utilisateurs à la volée en ligne de commande.


### Étape 3.2 : Persistance de données et JSON

Vous avez (j'espère) une belle interface `storer` qui vous permet de pouvoir intégrer facilement des nouveaux storer au fur et à mesure. Il est temps d'en rajouter une couche et d'assurer la persistance de données grâce à un fichier JSON.

Pas besoin de se prendre la tête : vous pouvez tout faire avec les packages de la bibliothèque standard de Go, notamment `encoding/json`. 

Normalement, vous savez ce qu'il vous reste à faire... (par exemple créer un nouveau fichier JSONstorer avec une struct et toutes les méthodes qui lui permettront d'être considérées comme une interface Storer)


## Étape 4 : Nouveau storer : GORM 

L'objectif ici est d'ajouter un nouveau système de stockage et de passer cette fois-ci par une vraie base de données SQLite via l'ORM GORM. Grâce à l'architecture basée sur les interfaces, cette modification devrait se faire sans impacter la logique métier.

### Steps

1. **Ajouter les dépendances nécessaires**, à la racine du projet ajoutez GORM et son driver SQLite :
```bash
    go get gorm.io/gorm
    go get gorm.io/driver/sqlite 
```

Pour les utilisateurs de windows, vous pouvez importer dans votre db.go le package https://github.com/glebarez/sqlite

2. Mettre à jour la struct `Contact`
3. Créer le `GORMStore`
4. Implémenter l'interface `Storer`
5. Intégrer dans `cmd/root.go`



## Étape 5 : Gestion de la config avec Viper


Pas grand chose à dire ici : Juste une config minimale avec par exemple un yaml qui dit quel type de storer utiliser au lancement du programme et on gère ça avec Viper !


## Fonctionnalités finales attendues

À la fin, vous obtiendrez un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go, incluant :

* **Une architecture en packages découplés.**
* **L'injection de dépendances via les interfaces.**
* **La création d'une CLI professionnelle avec Cobra.**
* **La gestion de configuration externe avec Viper.**
* **Gestion complète des contacts (CRUD)** : Ajouter, Lister, Mettre à jour et Supprimer des contacts.
* **Configuration externe** : Le comportement de l'application (notamment le type de stockage) peut être modifié sans recompiler.
* **Persistance des données** : Support de multiples backends de stockage :
  * GORM/SQLite : Une base de données SQL robuste contenue dans un simple fichier.
  * Fichier JSON : Une sauvegarde simple et lisible.
  * En mémoire : Un stockage éphémère pour les tests.



