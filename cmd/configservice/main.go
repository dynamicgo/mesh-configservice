package main

import (
	app "github.com/dynamicgo/mesh-app"
	_ "github.com/dynamicgo/mesh-configservice/load"
	_ "github.com/dynamicgo/mesh-libp2p-network"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app.Run()
}
