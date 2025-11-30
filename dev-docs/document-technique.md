# Document technique (fr)

## Tech stack

- [go](https://go.dev) : langage de programmation utilisé sur le projet
- [htmx](https://htmx.org) : librairie utilisée pour permettre au frontend de communiquer avec le backend
- [gorm](https://gorm.io) : librarire d'ORM en go permettant la définition du schéma de base de donnée, et la connexion aux principaux SGBD 

## Base de données

| settings |
| -------- |
|          |

| user      |
| --------- |
| id        |
| email     |
| validated |
|           |

| activation_codes     |
| -------------------- |
| id                   |
| code                 |
| expiration_timestamp |

| contest         |
| --------------- |
| id              |
| name            |
| start_timestamp |
| end_timestamp   |

| problem       |
| ------------- |
| id            |
| contest_id    |
| name          |
| points        |
| difficulty_id |
# Backend

Le backend est séparé en plusieurs modules go :

- **iosu** : le serveur 
- 


