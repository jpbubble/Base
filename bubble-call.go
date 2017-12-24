/*
        bubble-call.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.24
*/
package bubble

import (
	"trickyunits/mkl"
	"github.com/Shopify/go-lua"
)

// Works like a regular Lua call, but in stead of regular Lua crashouts
// through the panic system and such, this routine will still crash, but
// through all regularly set up routines, meaning you will see error
// messages the way the creator of the engine intended and that the
// any undo routines can run like planned.
// If you need not to crashout, you can still use Lua's own protected 
// call. For most bubble engines it's however best that you never
// use the regular Call as "panic" can cause memory leaks.
func Call(l *lua.State,ac,rc int){
	err:=l.ProtectedCall(ac,rc,0)
	if err!=nil {
		Fatal("LUA> "+err.Error())
	}
}

func tqCall(l *lua.State, f string, rc,eh int, funcargs []string) error{
	l.Global(f)
	for _,a :=range funcargs{
		l.PushString(a)
	}
	return l.ProtectedCall(len(funcargs),rc,0)
}

// Protected veresion of QCall
func QPCall(l *lua.State,f string, rc, eh int, funcargs ...string) error{
	return tqCall(l,f,rc,eh,funcargs)
}

// This routine accepts a function and a line of strings as parameters.
// Please note strings only for parameters here.
// Values returned by the lua call are ignored, BUT they can still read 
// out with regular lua routines.
func QCall(l *lua.State,f string,rc int,funcargs ...string){
	err:=tqCall(l,f,rc,0,funcargs)
	if err!=nil {
		Fatal("LUA> "+err.Error())
	}
}



func init(){
mkl.Version("Bubble Base - bubble-call.go","17.12.24")
mkl.Lic    ("Bubble Base - bubble-call.go","Mozilla Public License 2.0")
}
