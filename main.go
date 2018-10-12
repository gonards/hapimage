package main

import (
	"flag"
	"hapimage/api"
	"log"
)

var validType = map[string]bool{
	"sqlite": true,
}

func main() {
	// Get command line flags
	modeFlag := flag.String("m", "server", "Launch server or create DB. Allowed values are : server|skeleton")
	dbTypeFlag := flag.String("t", "sqlite", "Type of DB used. Allowed values are : sqlite")
	flag.Parse()
	switch *modeFlag {
	case "server":
		api.NewServer()
	case "skeleton":
		validateDBType(*dbTypeFlag)
		api.InitSkeleton(*dbTypeFlag)
	}
}

func validateDBType(dbType string) {
	if !validType[dbType] {
		log.Fatal("Invalid DB type. Allowd modes are : sqlite")
	}
}
