package main

import (
	"gopkg.in/macaron.v1"

	"log"
	"os"
	"strconv"

	config "github.com/helber/myMercurius/conf"
	conf "github.com/helber/myMercurius/conf/app"
)

// application entrypoint
func main() {
	app := macaron.New()
	conf.SetupMiddlewares(app)
	conf.SetupRoutes(app)
	/*
		Generated using http://www.kammerl.de/ascii/AsciiSignature.php - (Font: 'starwars')
		All signatures are made with FIGlet (c) 1991, 1993, 1994 Glenn Chappell and Ian Chai
		All fonts are taken from figlet.org and jave.de.
		Please check for Font Credits the figlet font database!
		Figlet Frontend - Written by Julius Kammerl - 2005
	*/
	log.Println(" __         _                                                         _                 __                 ")
	log.Println("[  |       / |_                                                      / |_              |  ]                ")
	log.Println("| | .---.`| |-'.--.  .---.  _ .--.   .---.  _ .--.  _   __  _ .--. `| |-'______   .--.| |  _ .--.   .--.  ")
	log.Println("| |/ /__\\| | ( (`\\]/ /__\\[ `.-. | / /'`\\][ `/'`\\][ \\ [  ][ '/'`\\ \\| | |______|/ /'`\\' | [ `.-. | ( (`\\] ")
	log.Println("| || \\__.,| |, `'.'.| \\__., | | | | | \\__.  | |     \\ '/ /  | \\__/ || |,        | \\__/  |  | | | |  `'.'. ")
	log.Println("[___]'.__.'\\__/[\\__) )'.__.'[___||__]'.___.'[___]  [\\_:  /   | ;.__/ \\__/         '.__.;__][___||__][\\__) )")
	log.Println("\\__.'   [__|                                           ")
	app.Run(port())
}

// configure http port
func port() int {
	port, err := config.Cfg.Section("").Key("http_port").Int()
	if err != nil {
		log.Fatal(err)
	}

	if forceLocal, _ := config.Cfg.Section("").Key("force_local_http_port").Bool(); forceLocal == false {
		if i, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
			port = i
		}
	}

	return port
}
