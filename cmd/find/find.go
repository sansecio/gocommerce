package main

import (
	"context"
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
		fmt.Printf("%s (ver: %s) at %s\n", store.Platform.Name(), ver, store.DocRoot)
		if store.Config != nil && store.Config.DB != nil {
			fmt.Printf("DBC: %+v\n", store.Config.DB)
		}

		if urls, err := store.Platform.BaseURLs(context.Background(), store.DocRoot); err == nil && len(urls) > 0 {
			fmt.Println("Base URLs:")
			for _, url := range urls {
				fmt.Printf("- %s\n", url)
			}
		}
		fmt.Println()
	}
}
