package main
//使用并发测试速度
import (
	"database/sql"
	"fmt"
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

func get_data(id int,db *sql.DB){

	rows,err := db.Query("select * from try where ida%4 = $1",id)
	checkErr(err)
	for rows.Next(){
		var IDa int
		var IDb int
		err = rows.Scan(&IDa,&IDb)//这行读取了每一个row里面的数据ida,idb
		checkErr(err)

	}
	quit <- 0        //用一个quitchannel来保证主线程在所有子线程结束后再结束.
}

func closed(db *sql.DB){
	db.Close()
	fmt.Println(time.Now())
}
func main(){

	start:=time.Now().UnixNano()//打印时间.
	println("开始时间是",start)


	//runtime.GOMAXPROCS(2)
	db,err := sql.Open("postgres","host=localhost port=5432 database= test user=postgres sslmode=disable password=root")
	checkErr(err)
	fmt.Println(time.Now())
 get_data(0,db)
 get_data(1,db)
 get_data(2,db)
 get_data(3,db)
	defer closed(db)
	for i := 0; i < 4; i++ {
		<- quit   //这样代码quit会保证了.最后运行完.
	}



	end:=time.Now().UnixNano()
	println("结束了")
	println(   (float64(end-start))/math.Pow10(9)   )
}
