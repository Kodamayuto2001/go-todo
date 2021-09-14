package main

import (
	"fmt"
	"net/http"
	
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func makeUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	uu := u.String()
	return uu 
}


//	ToDoアプリを作る
//	DBテーブルを作成する
//	Userテーブルを作成する

//	/api/v1/users/:id/boardにGETしてきた場合、
//	DBからユーザーIDに合致する全ボードを取得し、Jsonで返す

func main() {
	makeUUID()

	engine := gin.Default()

	engine.GET("/api/v1/users/:id/board",func(c *gin.Context){

	})
}
