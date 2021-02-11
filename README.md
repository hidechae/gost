# gost

Generate golang struct definitions from mysql table schema.

# Install

Run

```sh
go install github.com/hidechae/gost
```

# Usage

```sh
$ gost -h
Generate golang struct definitions from MySQL table schema.

Usage:
  gost -u root --host 127.0.0.1 -P 3306 -d test_db -t suffix_% [flags]

Flags:
  -d, --database string   Database
      --encoding string   Encoding (default "utf8mb4")
  -h, --help              help for gost
      --host string       Host address (default "127.0.0.1")
  -p, --password string   Password
  -P, --port string       Port (default "3306")
  -t, --table string      table name
  -u, --user string       User name (default "root")
```

For example, following table exists.
```sql
> desc user;
+------------+------------------+------+-----+-------------------+-----------------------------+
| Field      | Type             | Null | Key | Default           | Extra                       |
+------------+------------------+------+-----+-------------------+-----------------------------+
| id         | int(10) unsigned | NO   | PRI | NULL              | auto_increment              |
| email      | varchar(255)     | NO   |     | NULL              |                             |
| name       | varchar(255)     | YES  |     | NULL              |                             |
| created_at | timestamp        | NO   |     | CURRENT_TIMESTAMP |                             |
| updated_at | timestamp        | NO   |     | CURRENT_TIMESTAMP | on update CURRENT_TIMESTAMP |
+------------+------------------+------+-----+-------------------+-----------------------------+
```

gost generate struct definition from table schema.
```go
$ gost -uroot -h127.0.0.1 -P3306 -d test -t user
type User struct {
        Id uint
        Email string
        Name *string
        Gender int8
        CreatedAt time.Time
        UpdatedAt time.Time
}
```

# Feature

- Table and column name converted to camel case.
- Unsigned integer types mapped uint.
- Nullable column mapped pointer.
