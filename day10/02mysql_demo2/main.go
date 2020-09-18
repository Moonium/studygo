package main

// Go连接MySQL示例
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // 连接池

func initDB() (err error) {
	// DSN: Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	name string
	age  int
}

func queryOne(id int) {
	var u1 user

	sqlStr := `select id, name, age from user where id=?;`
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)

	fmt.Printf("u1: %#v\n", u1)
}

// 查询多条数据示例
func queryMultiRow(n int) {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func insert() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "周林", 29)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo(newAge, id int) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo(id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	var m = map[string]int{
		"小王子":  18,
		"沙河娜扎": 18,
	}
	for k, v := range m {
		_, err = stmt.Exec(k, v)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err: %v", err) // 格式错误
		return
	}
	fmt.Println("连接成功！")
	// queryOne(2)
	// queryMultiRow(2)
	// insert()
	// updateRowDemo(9000, 5)
	// deleteRowDemo(5)
	prepareInsertDemo()
}
