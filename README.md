# Coordinator Backend

Project dependencies:
- Install [Go](https://go.dev/)
- Install [go-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
- Install [taskfile](https://taskfile.dev/installation/)
- Be sure to have installed docker and docker-compose

### First run
```bash
task setup
```

If you already setup the project, you can choose to run only the database and run the project locally by your IDE with debug or run everything together.

To run only the database:
```bash
task db
```

To run database and code:
```bash
task run
```

To clear your docker containers:
```bash
task clear
```

#### TODO:

- [ ] Include hot reload on development
- [ ] Finish CRUDs (student, professor, school ...)
