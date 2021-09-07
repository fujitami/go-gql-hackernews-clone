package links

import (
	database "go-gql-hackernews-clone/internal/pkg/db/mysql"
	"go-gql-hackernews-clone/internal/users"
	"log"
)

// 構造体Linkの定義
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// DBに保存したLinkオブジェクトのidを返す関数
func (link Link) Save() int64 {
	// linksテーブルにオブジェクトを挿入するSQLクエリ
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	// SQLステートメントの実行
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	// linkオブジェクトのidを取得
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select id, title, address from Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil{
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
