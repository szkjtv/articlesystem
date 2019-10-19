package databases

import (
	"articlesystem/models"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// db, err := sql.Open("mysql", "root:aa14514118@tcp(127.0.0.1:3306)/articlesystem")
	// checkError(err)
	// defer db.Close()
	Dbcoom()
}

func Dbcoom() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:.aA1451418@tcp(123.207.88.76:3306)/articlesystem")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:.aA1451418@tcp(123.207.88.76:3306)/articlesystem")
	checkError(err)
	return
}

// FindUserByUsernameAndPassword 通过 username 和 password 查找 User
func FindUserByUsernameAndPassword(account string, password string) (user *models.User) {
	sql := "select id, account from user where account=? and password=?"

	db := getDB()
	defer db.Close()

	rows, err := db.Query(sql, account, password)
	checkError(err)
	defer rows.Close()

	if rows.Next() {
		var id int
		// var email string
		rows.Scan(&id)

		user = &models.User{
			Id:       id,
			Account:  account,
			Password: password,
		}
	}
	return
}

// AddUser 添加新 User
func AddUser(user *models.User) {
	sql := "insert into user(account, password) values(?,?)"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, user.Account, user.Password)
	checkError(err)
}

func AddContent(article *models.Article) {
	sql := "insert into article(title,content) values(?,?)"
	db := getDB()
	defer db.Close()
	_, err := db.Exec(sql, article.Title, article.Content)
	checkError(err)
}

// UpdateUser 更新 User
func UpdateUser(user *models.User) {
	sql := "update user set account=?, password=? where id=?"

	db := getDB()
	defer db.Close()

	_, err := db.Exec(sql, user.Account, user.Password, user.Id)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
