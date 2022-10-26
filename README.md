# Go postgres enums

Examples of how-to-use postgres enums in go.

## Run

```
make run
```

## Help

```
make help
```

## Enums conversion

First of all one should know about postgres enums:

- they are stored as integers (internal representation)
- they are returned by postgres protocol as strings (external representation)

I've tested two postgres drivers:

- [github.com/lib/pq](https://github.com/lib/pq)
- [github.com/jackc/pgx](https://github.com/jackc/pgx)

### String enums

When your enum type in go declared as

```go
type MyEnum string
```

then you're good to go. Both drivers would successfuly scan pgenums into
goenums and use goenums as parameters in queries/executions.

If your enum is not string the fun begins:

### Non-string enum as param to query/exec

For jackc/pgx you need to implement one of the following:

- [fmt.Stringer](https://pkg.go.dev/fmt#Stringer) for your type. pgx will cast
  param to string internally if your type satisfies the stringer interface
- [driver.Valuer](https://pkg.go.dev/database/sql/driver#Valuer) that returns
  `string`

For lib/pq you need to implement
[driver.Valuer](https://pkg.go.dev/database/sql/driver#Valuer) that returns
`string`.

### Non-string enums as scannable values

For scanning pgenum into your goenum you need to implement
[database.Scanner](https://pkg.go.dev/database/sql#Scanner) interface for your
type.

```go
func (me *MyEnum) Scan(value interface{}) error
```

BUT! jackc/pgx passes `string` as `value` parameter and lib/pq passes `[]byte`.
So whenever you want your type to be portable you need to support both possible
incoming types.

