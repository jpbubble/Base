/*
        bubble_main.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.21
*/
package bubble

import(
	"log"
	"strings"
	"trickyunits/mkl"
	"trickyunits/ansistring"
	"trickyunits/jcr6/jcr6main"
	"github.com/Shopify/go-lua"
)



var debugchat = false

var bubjcr jcr6main.TJCR6Dir


func dbg(msg string) { if debugchat {log.Print(ansistring.SCol(msg,4,0)) }}

func init(){
mkl.Version("Bubble Base - bubble_main.go","17.12.21")
mkl.Lic    ("Bubble Base - bubble_main.go","Mozilla Public License 2.0")
}

func SetJCR(jcrfile string) jcr6main.TJCR6Dir{
	bubjcr = jcr6main.Dir(jcrfile)
	return bubjcr
}

type TBubble struct {
	l *lua.State
	used map[string] bool
}

var vms = map[string] TBubble {}

func CreateBubble(id string) *lua.State {
	s:=strings.ToUpper(id)
	dbg("Creating bubble: "+s)
	l := lua.NewState()
	lua.BaseOpen(l)
	lua.OpenLibraries(l)
	initbubbleapi(s,l)
	lua.LoadString(l,"BUBBLE_VM_ID = '"+s+"'") // Please note this variable may NEVER be changed or bad stuff may happen!
	l.Call(0,0)
	lua.LoadString(l,sysscript)
	l.Call(0,0)
	vms[s] = TBubble { l,map[string] bool{} }
	return l
}

func GetBubble(id string) *lua.State {
	s:=strings.ToUpper(id)
	return vms[s].l
}

// Loads a script to a VM with id <id>
// If the script does not exist it will be created, if it does exist the current stuff will be added.
func LoadScript(id,script string){
	s:=strings.ToUpper(id)
	var l *lua.State
	if _,ok:=vms[s];!ok{
		l=CreateBubble(id)
	 } else {
		l = GetBubble(id)
	}
	dbg("Loading")
	uscript:="Use('"+script+"')\n"
	dbg(uscript)
	lua.LoadString(l,uscript)
	l.Call(0,0)
}


