# go / pkg / database / sql 教程

---
## overview
### sql.DB
sql.DB 并不是实际的数据库连接，也不是任何数据库系统的 database 或 schema 的映射，他只是一个接口和数据库存在的抽象。
sql.DB 会在幕后处理某些重要的任务：
- 通过驱动，打开和关闭与实际底层数据库的连接。
- 根据需要管理连接池。
sql.DB 抽象的设计就是为了让使用者不必担心如何管理对基础数据存储的并发访问。一个连接当在执行任务时会被标记为使用中，使用完成后则返回到连接池。需要注意的是，如果未释放连接到返回给连接池，将导致 sql.DB 打开大量的连接，可能耗尽资源。
创建 sql.DB 后，你可以使用它去查询它所代表的数据，以及创建语句和事务。

## Importing a Database Driver
database/sql 需要与特定数据库的驱动一起使用。
你通常不应该直接使用第三方驱动程序包，有些会鼓励你这样做，但在我们看来，这并不好。相反，如果可能，你的代码应仅引用 database/sql 中定义的类型。这有助于避免使代码依赖于驱动包，因此你可以使用更少的代码更改来更改底层驱动程序。它还会强制你使用 GO 习语，而不是特定驱动包的特殊习语。

mysql driver 示例：
```
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
```
匿名导入 _ ，驱动会将自己注册为可用于 database/sql 的包，通常除了运行 init 函数之外没有其他事情发生。 

---
## Accessing the Database
为了创建 sql.DB , 需要使用 sql.Open() , 其返回 *sql.DB :
```
func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
```
详细说明：
1. sql.Open 的第一个参数是 driver name 驱动名称。这是驱动程序用于向database/sql注册自身的字符串，并且通常与包名称相同以避免混淆。例如，mysql 对应的  github.com/go-sql-driver/mysql 。某些驱动没有遵循预定，但使用了数据库名称，例如 sqlite3 对应 github.com/mattn/go-sqlite3 ， postgres 对应 github.com/lib/pq 。
2. 第二个参数是驱动程序特定的语法，说明如何访问数据库系统。示例中连接的是本地mysql的 hello 数据库。
3. 你应该始终检查并处理所有 database/sql 参数返回的错误。稍后会讨论一些特殊情况，这样做是没有意义的。
4. 如果 sql.DB 的生命周期不应该超出函数的范围，通常会使用 `defer db.Close()` 。

与直觉相反的是，`sql.Open()`不会与数据库建立任何连接，也不会验证驱动连接参数。他只是准备了一个数据库的抽象以供后续使用。当需要时他才会建立第一个连接。如果你需要检查数据库是否可用且可访问（例如，检查是否可以建立网络连接并登录），请使用`db.Ping()`，别忘了检查错误：
```
err = db.Ping()
if err != nil {
	// do something here
}
```
尽管在完成使用时通常会 `Close()`，但是 sql.DB 对象被设计为 long-lived 。不要频繁的 `Open()` 和 `Close()` 数据库。你应该只创建一个 sql.DB 对象为你需要访问的每个不同的数据存储，并保留它直到程序完成访问该数据库。根据需要传递它，或者以某种方式在全局范围内提供它，但保持打开状态。不要在短期函数中 `Open()` 和 `Close()` ，相反，将 `sql.DB` 作为参数传递给该短期函数。

如果不将 sql.DB 作为长期存在的对象，你将会遇到很多连接上的问题。

---
## Retrieving Result Sets
从数据库中取回结果有以下几种常用方式：
1. Execute a query that returns rows.
2. Prepare a statement for repeated use, execute it multiple times, and destroy it.
3. Execute a statement in a once-off fashion, without preparing it for repeated use.
4. Execute a query that returns a single row. There is a shortcut for this special case.
database/sql 的函数名称很重要，如果函数名称包含 `Query` ，则它将返回一组 rows ，即使它是空的。不返回 rows 的 statements 语句，你应该使用 `Exec()` 。

### Fetching Data from the Dababase
```
var (
	id int
	name string
)
rows, err := db.Query("select id, name from users where id = ?", 1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, name)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}
```
说明：
1. 使用 `db.Query()` 返回 query 给数据库，检查错误。
2. `defer rows.Close()` , 这很重要。
3. 通过 `rows.Next()` 迭代 rows。
4. 通过 `rows.Scan()` 将每行中的列读入变量。
5. 完成迭代后检查错误。

这几乎是在 GO 中的唯一方式，你无法将 row 作为一个 map，因为强类型，您需要创建正确类型的变量并将指针传递给它们。

注意可能会出错的地方：
- 你应该在 `for rows.Next()` 循环结束时检查错误。循环期间出现错误，你也需要了解它。
- 使用 `rows.Next()` 迭代所有行，当读取最后一行时将会产生一个 EOF 错误并调用 `rows.Close()` 。如果由于某种原因提前退出循环，那么 rows 不会被关闭，连接仍然是打开的。（如果 `rows.Next()` 由于错误而返回false，则会自动关闭）。这是耗尽资源的简便方法。
- `rows.Close()` 是一个无害的操作，你可以多次调用它，即使已经关闭。但是，我们应该先检查错误，没有错误再调用，以避免运行时 panic 。
- 你应该总是 `defer rows.Close()` ，即使已经在循环结束时显式的调用。
- 不要在循环中 `defer` ，在函数退出之前，延迟的语句都不会执行，因此长时间运行的函数不应该使用它。如果你这样做，内存将会被慢慢占用。如果你在循环中反复查询和消费结果集，则应该在完成每个结果时显式的调用 `rows.Close()` ，而不是使用 `defer` 。

### How Scan() Works
迭代 rows 并 scan 到目标变量时，GO 会在后台执行数据类型转换。它基于目标变量的类型。意识到这一点可以清理代码并帮助避免重复性工作。

例如，假设某个表中的列被定义为 VARCHAR(45) 或类似的，但是实际值都是数字，如果你直接传递指针给字符串，GO 将会复制 bytes 给字符串。当然你也可以使用 `strconv.ParseInt()` 或类似方法进行转换，但是你不得不检查错误，这是很繁琐的。
或者，你可以直接传递 `Scan()` 指针给一个整形。Go 将会自动调用 `strconv.ParseInt()` ，这种方式代码更整洁更少，也是使用 database/sql 推荐的。

### Preparing Queries
多次使用的查询，你应该 prepare 。准备查询的结果是预准备 statement 语句，它为执行语句时提供的参数提供了占位符，这比拼接字符串好多了（例如，避免SQL注入）。

mysql 中的占位符是 `?` 
```
stmt, err := db.Prepare("select id, name from users where id = ?")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close()
rows, err := stmt.Query(1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	// ...
}
if err = rows.Err(); err != nil {
	log.Fatal(err)
}
```
`db.Query()` 实际上做了 prepares, executes, closes a prepared statement 三次与数据库的交互。如果您不小心，您可以将应用程序的数据库交互次数增加三倍！某些驱动程序可以在特定情况下避免这种情况，但并非所有驱动。

### Single-Row Queries
针对最多只返回一行的查询：
```
var name string
err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)
```
查询中的错误将被 defer ，直到调用 `Scan()`，然后从中返回。对于预准备的语句也可以使用 `QueryRow()` :
```
stmt, err := db.Prepare("select name from users where id = ?")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close()
var name string
err = stmt.QueryRow(1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)
```

---
## Modifying Data and Using Transactions
### Statements that Modify Data
使用 `Exec()`，最好是预处理的语句，来完成 `INSERT`, `UPDATE`, `DELETE` 或其他不返回行的语句。
```
stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
if err != nil {
	log.Fatal(err)
}
res, err := stmt.Exec("Dolly")
if err != nil {
	log.Fatal(err)
}
lastId, err := res.LastInsertId()
if err != nil {
	log.Fatal(err)
}
rowCnt, err := res.RowsAffected()
if err != nil {
	log.Fatal(err)
}
log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
```
执行语言将会生成一个 `sql.Result` ，它提供对语言元数据的访问：最后插入的ID和受影响的行数。

如果你不关心结果怎么办？如果您只想执行一个语句并检查是否有错误，但忽略结果怎么办？以下两个语句是否会做同样的事情？
```
_, err := db.Exec("DELETE FROM users")  // OK
_, err := db.Query("DELETE FROM users") // BAD
```
答案是否。两种方式并不相同，你也不应该这样使用 `Query()` 。`Query()` 将会返回一个 `sql.Rows()` , 它保留了数据库的连接直到其关闭。示例中永远不会释放连接。一定要注意。

### Working with Transactions
在Go中，事务本质上是一个保留与数据存储区连接的对象。它允许您执行我们迄今为止看到的所有操作，但保证它们将在同一连接上执行。

调用 `db.Begin()` 开始事务，并在生成的 Tx 变量上调用 `Commit()` 或者 `Rollback()` 关闭。Tx从池中获取连接，并保留它仅用于该事务。 Tx上的方法可以一对一地映射到可以在数据库本身上调用的方法，例如Query（）等。

在事务中创建的预处理语句专门绑定于该事务。

你不应该在代码的 SQL 语句中使用 BEGIN 和 COMMIT 。

注意不要在事务中调用 db 变量，db 是事务外的其他连接，事务必须使用 Tx 。

---
## Using Prepared Statements
预处理语句安全、高效、方便。但是它们的实现方式与您可能习惯的方式略有不同，特别是关于它们如何与 database/sql 的某些内部进行交互。

### Prepared Statements And Connections
在数据库级别，预处理语句绑定到一个数据库连接。典型的流程是客户端将带有占位符的SQL语句发送到服务器进行准备，服务器使用语句ID进行响应，然后客户端通过发送其ID和参数来执行该语句。

在 GO 中，连接并不直接暴露给使用 datebase/sql 包的用户。并不是在连接上预处理，而是在 DB 或者 Tx 上。并且 database/sql 具有一些便利行为，例如自动重试。由于这些原因，在驱动程序级别存在的预准备语句和连接之间的底层关联对代码是隐藏的。

他的工作原理：
1. 预处理语言时，基于连接池中的某个连接。
2. `Stmt` 对于记录所使用的连接。
3. 执行 Stmt 时，它会尝试使用该连接。如果由于它已关闭或忙于执行其他操作而无法使用，它将从池中获取另一个连接，并在另一个连接上使用数据库重新准备语句。
因为在原始连接繁忙时将根据需要重新准备语句，所以数据库的高并发使用可能会使很多连接繁忙，从而创建大量预准备语句。这可能导致语句的明显泄漏，正在准备和重新准备的语句比您想象的更频繁，甚至在语句数量上遇到服务器端限制。

### Avoiding Prepared Statements
GO 隐藏了预处理 statement 的操作。例如一条简单的 `db.Query(sql, param1, param2)` 会先准备sql，然后使用参数执行它，最后关闭语句。


### Prepared Statements in Transactions
在Tx中创建的预处理语句不能与它分开使用。同样，在DB上创建的预准备语句不能在事务中使用，因为它们将绑定到不同的连接。

在事务中想是使用事务外的 statement 可以调用 `Tx.Stmt()` ，但不建议这么做。

```
tx, err := db.Begin()
if err != nil {
	log.Fatal(err)
}
defer tx.Rollback()
stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close() // danger!
for i := 0; i < 10; i++ {
	_, err = stmt.Exec(i)
	if err != nil {
		log.Fatal(err)
	}
}
err = tx.Commit()
if err != nil {
	log.Fatal(err)
}
// stmt.Close() runs here!
```
GO 1.4 版本前，一定要确保在 commit 或者 rollback 之前关闭 statement 。

### Parameter Placeholder Syntax
参数占位符语法


---
## Handing Errors
几乎所有 database/sql 操作都会返回错误，你应该始终检查这些错误，不要忽略他们。
在某些地方，错误行为是特殊情况，或者您可能需要了解其他一些内容。

### Errors From Iterating Resultsets
思考以下代码：
```
for rows.Next() {
	// ...
}
if err = rows.Err(); err != nil {
	// handle the error here
}
```
`rows.Err()` 可能是 `rows.Next()` 循环中的任何错误。除了正常完成循环之外，循环可能由于某种原因而退出，因此您始终需要检查循环是否正常终止。异常终止会自动调用`rows.Close()`，尽管多次调用它是无害的。

### Errors From Closing Resultsets
如前所述，如果过早地退出循环，则应始终显式关闭 `sql.Rows`。如果循环正常退出或通过错误退出，它会自动关闭，但您可能会错误地执行此操作：
```
for rows.Next() {
	// ...
	break; // whoops, rows is not closed! memory leak...
}
// do the usual "if err = rows.Err()" [omitted here]...
// it's always safe to [re?]close here:
if err = rows.Close(); err != nil {
	// but what should we do if there's an error?
	log.Println(err)
}
```
对于检查所有数据库操作的错误，`rows.Close()` 是一个例外情况。如果 `rows.Close()` 返回错误，则不清楚应该怎么做。记录错误消息或 panicing 可能是唯一明智的事情，如果这不合理，那么也许你应该忽略错误。

### Errors From QueryRow()
思考以下代码：
```
var name string
err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)
```
如果查询结果为空会怎样？GO 定义了一个特定的错误常量 `sql.ErrNoRows` 针对 `QueryRow()` 返回结果为空的情况。这个错误会被延迟直到 `Scan()` 调用完成，最好处理如下：
```
var name string
err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
if err != nil {
	if err == sql.ErrNoRows {
		// there were no rows, but otherwise no error occurred
	} else {
		log.Fatal(err)
	}
}
fmt.Println(name)
```
这个错误只会在 `QueryRow()` 返回。

### Identifying Specific Database Errors
思考以下代码：
```
rows, err := db.Query("SELECT someval FROM sometable")
// err contains:
// ERROR 1045 (28000): Access denied for user 'foo'@'::1' (using password: NO)
if strings.Contains(err.Error(), "Access denied") {
	// Handle the permission-denied error
}
```
最好不要这样比较字符串，而应该比较错误号。但是这与驱动相关，mysql 是这样：
```
if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
	if driverErr.Number == 1045 {
		// Handle the permission-denied error
	}
}
```
同样，此处的MySQLError类型由此特定驱动程序提供，并且.Number字段可能因驱动程序而异。但是，数字的值取自MySQL的错误消息，因此是特定于数据库的，而不是特定于驱动程序的。
使用数字仍不太好，最好使用错误列表：
```
if driverErr, ok := err.(*mysql.MySQLError); ok {
	if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
		// Handle the permission-denied error
	}
}
```

### Handing Connection Errors
如果连接异常怎么办？
database/sql 已经处理了这种情况，连接异常会自动重建新的连接并重试最多 10 次。
但是，可能会产生一些意想不到的后果。当发生其他错误情况时，可能会重试某些类型的错误。这可能也是特定于驱动程序的。比如 KILL 了某个 statement 仍会导致重试 10 次。

---
## Working with NULLs
NULL 值能避免就避免，如果避免不了需要使用 `database/sql` 中的特殊类型或者你自己定义的类型去处理它们。

```
for rows.Next() {
	var s sql.NullString
	err := rows.Scan(&s)
	// check err
	if s.Valid {
	   // use s.String
	} else {
	   // NULL value
	}
}
```
如果你需要自定义类型处理空值，可以参考以上 `sql.NullString` 。

另一种方式是使用数据库提供的 `COALESCE()` :
```
rows, err := db.Query(`
	SELECT
		name,
		COALESCE(other_field, '') as otherField
	WHERE id = ?
`, 42)

for rows.Next() {
	err := rows.Scan(&name, &otherField)
	// ..
	// If `other_field` was NULL, `otherField` is now an empty string. This works with other data types as well.
}
```

---
## Working with Unknown Columns
`Scan()` 要求你明确目标变量，但是如果你不知道查询的返回结果呢？

如果你不知道返回结果有多少列，你可以使用 `Columns()` 列举出来。
您可以检查此列表的长度以查看有多少列，并且可以使用正确数量的值将切片传递到 `Scan()`。
例如处理 `SHOW PROCESSLIST` ：
```
cols, err := rows.Columns()
if err != nil {
	// handle the error
} else {
	dest := []interface{}{ // Standard MySQL columns
		new(uint64), // id
		new(string), // host
		new(string), // user
		new(string), // db
		new(string), // command
		new(uint32), // time
		new(string), // state
		new(string), // info
	}
	if len(cols) == 11 {
		// Percona Server
	} else if len(cols) > 8 {
		// Handle this case
	}
	err = rows.Scan(dest...)
	// Work with the values in dest
}
```
如果你既不知道列，又不知道类型，则应该使用 `sql.RawBytes` :
```
cols, err := rows.Columns() // Remember to check err afterwards
vals := make([]interface{}, len(cols))
for i, _ := range cols {
	vals[i] = new(sql.RawBytes)
}
for rows.Next() {
	err = rows.Scan(vals...)
	// Now you can check each element of vals for nil-ness,
	// and you can use type introspection and type assertions
	// to fetch the column into a typed variable.
}
```

---
## The Connection Pool
database/sql 包提供了基本的连接池，但并没有很多控制或者检查的能力。
- 连接池意味着两个语句可能分别执行在两个不同的连接上。
- 在需要时创建连接，并且池中没有空闲连接。
- 默认情况下，连接数没有限制。如果您尝试同时执行大量操作，则可以创建任意数量的连接。这可能导致数据库返回错误，例如“连接太多”。
- 在Go 1.1或更高版本中，您可以使用 `db.SetMaxIdleConns(N)`来限制池中的空闲连接数。但是，这并不限制池大小。
- 在Go 1.2.1或更高版本中，您可以使用 `db.SetMaxOpenConns(N)` 来限制到数据库的总打开连接数。不幸的是，死锁错误（修复）会阻止 `db.SetMaxOpenConns(N)` 安全地在1.2中使用。
- 连接回收速度相当快。使用 `db.SetMaxIdleConns(N)` 设置大量空闲连接可以减少此流失，并有助于保持连接以便重用。
- 长时间保持连接空闲可能会导致问题（例如Microsoft Azure上的MySQL问题）。如果连接超时，请尝试使用 `db.SetMaxIdleConns(0)`，因为连接空闲时间过长。
- 您还可以通过设置 `db.SetConnMaxLifetime（duration）` 来指定可以重用连接的最长时间，因为重用长期连接可能会导致网络问题。这懒惰地关闭未使用的连接，即可以延迟关闭过期的连接。

---
## Surprises, Antipatterns and Limitations

### Resource Exhaustion
注意资源耗尽

### Large uint64 Values
```
_, err := db.Exec("INSERT INTO users(id) VALUES", math.MaxUint64) // Error
```
使用 `uint64` ，较小值可能运行正常，但随着值增大可能会导致意想不到的错误。

### Connection State Mismatch
注意执行 statement 的连接是不同的。

### Database-Specific Syntax
注意某些特定语法，例如占位符是不同的。

### Multiple Result Sets
不支持

### Invoking Stored Procedures
调用存储过程是基于特定驱动的，但 mysql 驱动无法实现。

### Multiple Statement Support
单次执行多个 statement 并不明确支持：
```
_, err := db.Exec("DELETE FROM tbl1; DELETE FROM tbl2") // Error/unpredictable result
```
你应该拆分成多次分别执行。
