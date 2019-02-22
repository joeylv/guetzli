package dbcon

import (
	"../module"
	"fmt"
	"time"

	//"github.com/astaxie/beego/orm"
	"github.com/go-xorm/xorm"
	"os"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

//var db, _ = sql.Open("sqlite3", pwd+"/foo.db")
var pwd, _ = os.Getwd()
var engine *xorm.Engine

func init() {
	var err error
	//与数据库建立链接
	engine, err = xorm.NewEngine("sqlite3", pwd+"/foo.db")
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	checkErr(err)
	//在控制台打印sql语句，默认为false
	engine.ShowSQL(false)
	err = engine.Sync(new(module.Path), new(module.Top))
	checkErr(err)
	// set default database
	//orm.RegisterDataBase("default", "sqlite3", pwd+"/foo.db", 30)
	//orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)

	// register model
	//orm.RegisterModel(new(module.DirInfo))

	// create table
	//orm.RunSyncdb("default", false, true)
	//pwd, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

}
func Insert(arg ...string) int64 {

	if arg[0] == "Path" {
		p := module.Path{
			Name: arg[1], Path: arg[2], Thumb: arg[3],
		}
		count, _ := engine.Insert(&p)
		return count

	} else if arg[0] == "Top" {
		//fmt.Println(arg)
		//var ret module.Path
		//engine.Where("thumb = ?", arg[2]).Get(&ret)
		//fmt.Println(ret)
		t := module.Top{
			File:  arg[1],
			Thumb: arg[2],
		}
		count, _ := engine.Insert(&t)
		return count
	}

	//o := orm.NewOrm()
	//dir := module.DirInfo{Id: id, Name: strconv.Itoa(int(id)), Dir: dirName, Thumb: thumb}
	//id, err := o.Insert(&dir)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	return 0
}

func Update(arg ...string) int64 {
	var count int64 = 0
	if arg[0] == "Path" {
		p := module.Path{
			Name: arg[1], Path: arg[2], Thumb: arg[3],
		}
		count, _ = engine.Update(&p)
		//checkErr(err)
		fmt.Println(count)
	} else if arg[0] == "Top" {
		//fmt.Println(arg)
		//var ret module.Path
		//engine.Where("thumb = ?", arg[2]).Get(&ret)
		//fmt.Println(ret)

		top := module.Top{File: arg[1], Thumb: arg[2]}
		has, err := engine.Get(&top)
		checkErr(err)
		fmt.Println(has)
		fmt.Println(top)
		fmt.Println(arg[3] == "h")
		if !has {
			if arg[3] == "h" {
				top.H = 1
			} else {
				top.L = 1
			}
			top.Time = time.Now()
			engine.Insert(&top)
		} else {
			if arg[3] == "h" {
				top.H = top.H + 1
			} else {
				top.L = top.L + 1
			}
			//count, err := strconv.Atoi(arg[1])
			//checkErr(err)

			//t := module.Top{
			//	Count: count,
			//	Thumb: arg[2],
			//}
			count, _ = engine.Id(top.Id).Update(&top)
			fmt.Println(count)
		}
		//return count
	}
	return count
}

func Get(dir *module.Path) {
	//o := orm.NewOrm()
	//u := module.Path{Id: dir.Id}
	//err := o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)

}
func SearchPath(arg ...string) *module.Path {
	path := &module.Path{Path: arg[0]}
	has, _ := engine.Get(path)
	//fmt.Println(has)

	//checkErr(err)
	//fmt.Println(path)
	if has {
		return path
	}
	return nil
}
func SearchTop(arg ...string) *module.Top {
	top := &module.Top{File: arg[1], Thumb: arg[2]}
	has, err := engine.Get(top)
	fmt.Println(has)

	checkErr(err)
	fmt.Println(top)
	return top
}

func Remove(dir *module.Path) {
	//o := orm.NewOrm()
	//
	//num, err := o.Delete(dir)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//删除数据
	//stmt, err := db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res, err := stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
	//
	//db.Close()
}

//CreateTable
func Create() {
	//sql_table := `
	//CREATE TABLE IF NOT EXISTS dir_info(
	//    uid INTEGER PRIMARY KEY AUTOINCREMENT,
	//	name VARCHAR(64) NULL,
	//    dirName VARCHAR(256) NULL,
	//    created DATE NULL
	//);
	//`

	//db.Exec(sql_table)
	//db.Close()
}

func Count(m string) int64 {
	if m == "Path" {
		path := new(module.Path)
		total, err := engine.Count(path)
		checkErr(err)
		return total
	}

	return 0
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
