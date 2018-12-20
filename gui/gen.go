
//+build generate

package main

import "github.com/zserge/lorca"
//run vuejs's npm run build first to create static files
//go:generate npm run build --prefix ./ginbro-spa
//convert
func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	lorca.Embed("assets.go", "ginbro-spa/dist")

}
