package useMysql

import (
	"database/sql"
	_ "databases/mysql"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Insert() {
	db, err := sql.Open("mysql", "root:sa666@tcp(59.110.159.6:3306)/test?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare("insert user_infos set name=?,age=?,sex=?")
	checkErr(err)
	res, err := stmt.Exec("iliya", "21", "0")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func Update() {
	db, err := sql.Open("mysql", "root:sa666@tcp(59.110.159.6:3306)/test?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare("update user_infos set name=? where id=?")
	checkErr(err)
	res, err := stmt.Exec("i丽呀", 14)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func Query() {
	db, err := sql.Open("mysql", "root:sa666@tcp(59.110.159.6:3306)/test?charset=utf8")
	checkErr(err)
	//查询数据
	rows, err := db.Query("select * from user_infos")
	checkErr(err)
	for rows.Next() {
		var uid int
		var name string
		var age int
		var sex int
		err = rows.Scan(&uid, &name, &age, &sex)
		checkErr(err)
		fmt.Print("uid：", uid)
		fmt.Print("  name: ", name)
		fmt.Print("  age: ", age)
		fmt.Println("  sex: ", sex)
	}
}

func Delete() {
	db, err := sql.Open("mysql", "root:sa666@tcp(59.110.159.6:3306)/test?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare("delete from user_infos where id=?")
	checkErr(err)
	res, err := stmt.Exec(12)
	checkErr(err)
	affect, err := res.RowsAffected()
	fmt.Println(affect)

}
