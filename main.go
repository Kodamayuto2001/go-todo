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

type Boards struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
}

func AddBoard(board *Boards) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql",os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
}

func main() {
	makeUUID()

	engine := gin.Default()

	engine.POST("/api/v1/users/:id/boards",func(c *gin.Context){
		var boards Boards
		err := c.ShouldBindJSON(&(boards.Name))
		if err != nil {
			c.JSON(http.Status.BadRequest,gin.H{
				"error": err.Error(),
			})
			return
		}


		c.JSON(
			http.StatusOK,
			gin.H{
				"message":"Board is created"
			}
		)
	})

	engine.GET("/api/v1/users/:id/boards",func(c *gin.Context){

	})
}
