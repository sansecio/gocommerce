package main

import (
	"fmt"
	"os"

	"github.com/sansecio/gocommerce"
)

func main() {
	var stores []*gocommerce.Store
	if len(os.Args) <= 1 {
		stores = gocommerce.DiscoverStores()
	} else {
		for _, arg := range os.Args[1:] {
			if store := gocommerce.FindStoreAtRoot(arg); store != nil {
				stores = append(stores, store)
			}
		}
	}

	fmt.Println("Found", len(stores), "stores")
	for _, store := range stores {
		ver, err := store.Platform.Version(store.DocRoot)
		if err != nil {
			ver = "unknown"
		}
		fmt.Printf("- %s (ver: %s) at %s\n", store.Platform.Name(), ver, store.DocRoot)
		fmt.Printf("DBC: %+v\n", store.Config.DB)
	}
}
