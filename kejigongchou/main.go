package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"kejigongchou/lzu_kjfz"
	"log"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(lzu_kjfz.Cors())
	// 服务静态文件目录
	//router.Static("/static", "./static")
	// 服务虚拟静态文件系统
	router.StaticFS("/lzu_kjfz", http.Dir("./lzu_kjfz"))
	// 服务单个静态文件
	//router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.POST("/getall", postData)
	router.GET("/getdata", getData)
	err := router.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func getData(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./receiveddata.db")
	checkErr(err)
	rows, err := db.Query("select * from data")
	for rows.Next() {
		var id int
		var name string
		var phone string
		var dwmc string
		var zwzc string
		var kyfzsl string
		var kyfzfx string
		var kygljzyhgg string
		var kytdzj string
		var zdkjxmgg string
		var qykjhz string
		var qt string
		err = rows.Scan(&id, &name, &phone, &dwmc, &zwzc, &kyfzsl, &kyfzfx, &kygljzyhgg, &kytdzj, &zdkjxmgg, &qykjhz, &qt)
		checkErr(err)
		fmt.Println(id, name, phone, dwmc, zwzc, kyfzsl, kyfzfx, kygljzyhgg, kytdzj, zdkjxmgg, qykjhz, qt)
	}
}

func postData(c *gin.Context) {
	//fmt.Println(c)
	json := make(map[string]interface{}) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("%v", &json)
	c.JSON(http.StatusOK, gin.H{
		"name":     json["name"],
		"password": json["password"],
	})
	name := json["name"]
	phone := json["phone"]
	dwmc := json["dwmc"]             // 单位名称
	zwzc := json["zwzc"]             // 职务/职称
	kyfzsl := json["kyfzsl"]         // 科研发展思路
	kyfzfx := json["kyfzfx"]         // 科研发展方向
	kygljzyhgg := json["kygljzyhgg"] // 科研管理机制优化改革
	kytdzj := json["kytdzj"]         // 科研团队组建
	zdkjxmgg := json["zdkjxmgg"]     // 重大科技项目攻关
	qykjhz := json["qykjhz"]         // 区域科技合作
	qt := json["qt"]                 // 其它
	fmt.Println(name, phone, dwmc, zwzc, kyfzsl, kyfzfx, kygljzyhgg, kytdzj, zdkjxmgg, qykjhz, qt)
	insertData(json)
}

func insertData(json map[string]interface{}) {
	db, err := sql.Open("sqlite3", "./receiveddata.db")
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO data(name, phone, dwmc,zwzc,kyfzsl,kyfzfx,kygljzyhgg,kytdzj,zdkjxmgg,qykjhz,qt) values(?,?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)
	res, err := stmt.Exec(json["name"], json["phone"], json["dwmc"], json["zwzc"], json["kyfzsl"], json["kyfzfx"], json["kygljzyhgg"], json["kytdzj"], json["zdkjxmgg"], json["qykjhz"], json["qt"])
	checkErr(err)
	fmt.Println(res.LastInsertId())
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
