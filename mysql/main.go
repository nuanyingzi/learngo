package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局对象db
var db *sql.DB

type user struct {
	id, age int64
	name    string
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 查询单条数据
func queryRowDemo() {
	sqlStr := "select id,name,age from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 查询多条数据
func queryMultiRowDemo() {
	sqlStr := "select id, name,age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?,?)"
	ret, err := db.Exec(sqlStr, "李小萌", 18)
	if err != nil {
		fmt.Printf("insert to failed,err:%v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d\n", theId)
}

// 修改数据
func updateRowDemo() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 100, 2)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get rowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%v\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get rowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%v\n", n)
}

// 预处理查询示例
func prepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
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

// sql注入
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}

// 事务
func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "update user set age = 30 where id = ?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected failed, err: %v\n", err)
		return
	}

	sqlStr2 := "update user set age = 40 where id = ?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret2.RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交了")
		tx.Commit()
	} else {
		tx.Rollback()
		fmt.Println("事务回滚了")
	}
	fmt.Println("exec trans success!")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	// insertRowDemo() // 插入数据
	// queryRowDemo()      // 查询单条数据
	// queryMultiRowDemo() // 查询多条数据
	// updateRowDemo() // 修改数据
	//deleteRowDemo() // 删除数据
	//prepareQueryDemo() // 预处理查询
	//sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
	transactionDemo() // 事务
}
