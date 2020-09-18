package main

// Go连接MySQL示例
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DSN: Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		fmt.Printf("dsn %s invalid, err: %v", dsn, err) // 格式错误
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err: %v", dsn, err)
		return
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
	fmt.Println("linked success!")
}
