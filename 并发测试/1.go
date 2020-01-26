//这个sql是创建数据库里面的数据,,    首先自己手动简历一个表交try里面的column是ida, idb 即可.
package main

import(
"fmt"
"database/sql"
_ "github.com/lib/pq"
	"math"
	"time"
)
var quit chan int = make(chan int)

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func main6767(){
	start:=time.Now().UnixNano()//打印时间.
	println("开始时间是",start)
	db,err := sql.Open("postgres","host=localhost port=5432 database= test user=postgres sslmode=disable password=root")
	checkErr(err)
	for i:=0; i<10000;i++{
		stmt,err := db.Prepare("insert into try values($1,1)")
		checkErr(err)
		stmt.Exec(i+1)
	}
	db.Close()
	end:=time.Now().UnixNano()
	println(   (float64(end-start))/math.Pow10(9)   )

}
