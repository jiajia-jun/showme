package main

import (
	"webproject/dao"
	"webproject/router"
	"log"
	"os"
)

func main() {
	dao.Pickdata()
	dao.LoadProfile()
	dao.LoadMessages()

	r := router.InitRouter()

	// 证书路径
	//crtPath := "ssl/server.crt"
	//keyPath := "ssl/server.key"
	//_, errC := os.Stat(crtPath)
	//_, errK := os.Stat(keyPath)
	//
	//if errC != nil || errK != nil {
	//	log.Println("启用HTTP")
	//	err := r.Run(":8080") // http启动
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//} else { // https启动
	//	log.Println("启用HTTPS")
	//	err := r.RunTLS("10.17.250.224:8443", crtPath, keyPath)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	lcrtPath := "ssl/Radmin_LAN/server_LAN.crt"
	lkeyPath := "ssl/Radmin_LAN/server_LAN.key"
	_, lerrC := os.Stat(lcrtPath)
	_, lerrK := os.Stat(lkeyPath)

	if lerrC != nil || lerrK != nil {
		log.Println("启用HTTP")
		err := r.Run(":8080") // http启动
		if err != nil {
			log.Fatal(err)
		}

	} else { // https启动
		log.Println("启用HTTPS")
		err := r.RunTLS("26.126.204.192:8443", lcrtPath, lkeyPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
