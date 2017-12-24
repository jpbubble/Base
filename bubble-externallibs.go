/*
        bubble-externallibs.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.24
*/
package bubble

import (
	"github.com/Shopify/go-lua"
)


type bublib struct {
	f func(l *lua.State)
}


var bublibs = []bublib{}

// Registers a function that needs to be executed when a new Lua VM is 
// created. Basically this routine has to make sure all libraries added
// to Lua are automatically added.
func BREG(f func(l *lua.State)){
	bublibs = append(bublibs,bublib{f})
}

func bINI(l *lua.State){
	for _,bl:=range(bublibs){ bl.f(l) }
}

