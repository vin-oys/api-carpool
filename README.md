# api-carpool

### **Repo setup is referencing from [Techschool](https://github.com/techschool/simplebank)**

---

Install docker

[Docker download](https://docs.docker.com/get-docker/)

---

Install sqlc (for mac only)

[sqlc documentation](https://docs.sqlc.dev/en/stable/)
>`brew install sqlc`

---

Install golang-migrate.

[golang-migrate documentation](https://github.com/golang-migrate/migrate) 
>`brew install golang-migrate`
---
Please refer to **Makerfile** for scripts

Build network for DB and App connection
>`make network`

Run docker container
>`make postgres`

Create DB
>`make createdb`

Run DB migration script
>`make migrateup`

Build golang app
>`make server-build`

Run golang app
>`make server-run`
---
