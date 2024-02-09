package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sansecio/gocommerce"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("No document root specified.")
	}

	store := gocommerce.FindStoreAtRoot(os.Args[1])
	if store == nil {
		log.Fatalf("Unable to find store at %s\n", os.Args[1])
	}

	ver, err := store.Platform.Version(os.Args[1])
	if err != nil {
		ver = "unknown"
	}
	fmt.Printf("Found %s (ver: %s) at %s\n", store.Platform.Name(), ver, os.Args[1])
	fmt.Printf("DBC: %+v\n", store.Config.DB)
}
