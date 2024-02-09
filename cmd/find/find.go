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

	fmt.Printf("Found %s at %s\n", store.Platform.Name(), os.Args[1])
	fmt.Printf("DBC: %+v\n", store.Config.DB)
}
