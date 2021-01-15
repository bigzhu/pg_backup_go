used to backup remote postgresql DB to local

## install

```
go get github.com/bigzhu/pg_backup_go
```

## Create db config file

create `db.toml` file with content:

```toml
host = "127.0.0.1"
user = "bigzhu"
db_name = "bigzhu"
password = "bigzhu"
```

## usage

run

```bash
backup_pg_go
```

in the db.toml file path
