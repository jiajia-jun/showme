package main

import (
	"log"
	"os"
	"webproject/dao"
	"webproject/router"
)

func main() {
	dao.Pickdata()
	dao.LoadProfile()
	dao.LoadMessages()
	dao.InitImageCache()

	r := router.InitRouter()

	// RadminLan启动
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
		err := r.RunTLS("localhost:8443", lcrtPath, lkeyPath)
		

		if err != nil {
			log.Fatal(err)
		}
	}

}

// 26.126.204.192:8443
