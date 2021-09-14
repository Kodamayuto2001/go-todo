package main

import (
	"fmt"
	"net/http"
	"log"
	"database/sql"
	"os"
	
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

func AddBoard(board *Boards) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return "godotenv err!"
	}
	db, err := sql.Open("mysql",os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
		return "db open err!"
	}
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO boards(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
		return "db.Prepare err!"
	}

	ins.Exec(board.Name)

	id, err := db.Query("SELECT id FROM boards")
	if err != nil {
		log.Fatal(err)
		return "db select err!"
	}
	defer id.Close()

	err = id.Scan(board.ID)
	if err != nil {
		log.Fatal(err)
		return "id.Scan err!"
	}

	return "nil"
}

func main() {
	makeUUID()

	engine := gin.Default()

	engine.POST("/api/v1/users/:id/boards",func(c *gin.Context){
		var boards Boards
		err := c.ShouldBindJSON(&(boards.Name))
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
			return
		}

		message := AddBoard(&boards)
		if message != "nil" {
			fmt.Println(message)
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"message":"Board is created",
			},
		)
	})

	// engine.GET("/api/v1/users/:id/boards",func(c *gin.Context){

	// })

	engine.Run(":3000")
}
