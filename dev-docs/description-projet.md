# Descriptions du projet (fr)

Site internet de gestion de concours de programmation, dans le style d'[adventofcode](https://adventofcode.com/). 
# Pour les utilisateurs

Pour chaque problème, chaque participant ou chaque équipe reçoit une entrée (*input*) unique associée à son compte (ou à son équipe) et doit générer une sortie (*output*). Une entrée possède une seule sortie possible.

> [!info] Exemple
>  **Problème :** donnez la somme des nombres reçus. Les nombres sont séparés par des virgules.
> **Input :** `2,5,3,4,5,7,12`
> 
> L'utilisateur développe un programme dans le langage de programmation de son choix lui permettant de résoudre le problème. Le programme pourrait par exemple être :
> 
```python
input = "2,5,3,4,5,7,12"
sum = 0
for s in input.split(","):
  sum += int(s)
print(sum)
```
> 
> **Sortie attendue :** `38`
## Création de compte utilisateur

L'utilisateur reçoit un lien lui permettant d'activer son compte et se créer un mot de passe.

## Notifications

Une section "notification" permet de lister des informations envoyées par les admins lors du déroulement du concours.

# Pour les administrateurs 

Un panel d'administration contient différents onglets permettant d'administrer un concours avant et pendant son déroulement.

Il est constitué de plusieurs sections, une pour les différents concours, avec un bouton pour créer un nouveau concours, et une pour administrer la plateforme en général (création des utilisateurs, etc...)

## Gestion des concours

Pour un concours, il est possible d'ajouter des problèmes, et d'éditer leur contenu en markdown. 
Pour importer les données par utilisateur, dans la première version, il faudra importer un ficher zip contenant une structure de ce type :

```
dossier_principal/
├─ identifiant_utilisateur1/
│  ├─ input.txt # exemple : 2,5,3,4,5,7,12
│  ├─ output.txt # exemple : 38
├─ identifiant_utilisateur2/
│  ├─ input.txt
│  ├─ output.txt
├─ [...]
```

Ce type de donnée peut ainsi être généré par les créateurs de problème dans le langage de leur choix, ce qui permet à plus de monde de créer des problèmes.

## Gestion des utilisateurs

Un formulaire d'inscription permet de se créer un compte, et les comptes doivent être validés par les administrateurs. Ces mêmes administrateurs peuvent également créer les comptes manuellement, et importer un fichier CSV pour créer plusieurs comptes d'un coup. Dans ce cas, les utilisateurs sont prévenus par mail avec un lien leur permettant de se créer un mot de passe. Cette méthode permet de faire un formulaire d'inscription (google form par exemple) qui permet d'exporter les données en CSV.
