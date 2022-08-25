package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mcos/schemabuf/schemabuf"
	"io/ioutil"
	"log"
)

var serverPath = "system"

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/grpc_admin?charset=utf8mb4&parseTime=true"
	pkg := "pb"
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	s, err := schemabuf.GenerateSchema(db, pkg, []string{
		"casbin_rule",
		"ga_roles",
		"ga_user_roles",
		"ga_user_signin_logs",
		"ga_users",
	})
	if err != nil {
		log.Fatal(err)
	}
	
	if s != nil {
		fmt.Println(s)
		protoFile := fmt.Sprintf("app/%s/%s/%s.proto", serverPath, serverPath, serverPath)
		err := ioutil.WriteFile(protoFile, []byte(s.String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
