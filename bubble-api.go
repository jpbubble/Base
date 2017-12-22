package bubble


import(
	//"fmt"
	"path"
	"strings"
	"trickyunits/jcr6/jcr6main"
	"github.com/Shopify/go-lua"
)


// When set true the styem will note what it compiles. When set to false it will keep this silent.
var CompileChat = true


func bubble_CRASH(l *lua.State) int {
	e:=lua.OptString(l,1,"Unknown error")
	Fatal(e)
	return 0
}

func bubble_Use(l *lua.State) int {
	scriptfile:=lua.CheckString(l,1)
	vm:=lua.CheckString(l,2)
	s:=strings.ToUpper(vm)
	dbg("Use >> vm="+s+"; script="+scriptfile)
	//fmt.Print("Use("+scriptfile+","+vm+")")
	if _,ok:=vms[s];!ok { Fatal("Unknown VM: "+s) }
	ll:=vms[s].l
	us:=strings.ToUpper(scriptfile)
	if _,ok:=vms[s].used[us];!ok{
		if CompileChat {
			Console.Write("Yellow","Compiling: ")
			Console.WriteLn("Cyan",scriptfile)
		}
		vms[s].used[us]=true
		b:=jcr6main.JCR_String(bubjcr,scriptfile)
		if jcr6main.JCR6Error!="" { Fatal(jcr6main.JCR6Error) }
		err:=lua.LoadBuffer(ll , b, path.Base(scriptfile), "")
		if err!=nil { Fatal(err.Error()) }
		l.Call(0,0)
	} else {
		//fmt.Print("Loaded") // debug line
	}
	return 0
}

func bubble_jcrdir(l *lua.State) int {
	ret:="local ret = {}\n"
	for _,e:=range bubjcr.Entries {
		ret+="ret[#ret+1]=\""..e.Entry.."\"\n"
	}
	ret+="\n\n\nreturn ret"
	lua.PushString(l,ret)
	return 1
}











func initbubbleapi(bub string,l *lua.State){
		lua.SetFunctions(l, []lua.RegistryFunction{ 
			{ "BUBBLE_TRUE_USE",bubble_Use },
			{ "BUBBLE_JCR_DIR",bubble_jcrdir },
			{ "CRASH",bubble_CRASH },
			},0)

}
