/*
        bubble-preprocessor.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.22
*/
package bubble

import(
	"fmt"
	"strings"
	"runtime"
	"trickyunits/mkl"
	"trickyunits/qstr"
	)
var globdefs = map[string] bool{}

func init(){
globdefs["$"+runtime.GOOS]=true
mkl.Version("Bubble Base - bubble-preprocessor.go","17.12.22")
mkl.Lic    ("Bubble Base - bubble-preprocessor.go","Mozilla Public License 2.0")
}

func checkdirectives(a[] string,vm string) bool{
	if len(a)<3 { Fatal("Preprocessor expression missing!") }
	ret:=true
	for i:=2;i<len(a);i++{
		switch(qstr.Left(a[i],1)) {
			case "@","$":
				if v,ok:=globdefs[a[i]]; ok { ret = ret && v } else {ret=false}
			case "!":
				switch(qstr.Left(a[i],2)) {
					case "@","$":
						if v,ok:=globdefs[a[i][1:]]; ok { ret = ret && (!v) } 
					default:
						if v,ok:=vms[vm].defs[a[i][1:]]; ok { ret = ret && (!v) } 
					}
			default:
				if v,ok:=vms[vm].defs[a[i][1:]]; ok { ret = ret && (!v) }
		} 
	}
	return ret
}


func preprocess(lines []string,vm string) string {
	ret:=""
	workif:=false
	igif:=false
	for i,lin := range lines {
		linenumber:=fmt.Sprintf("%d",i+1)
		line := qstr.MyTrim(lin)
		if qstr.Prefixed(line,"-- #") {
			ps:=strings.Split(line," ")
			switch strings.ToLower(ps[1]){
			case "#demand":
				if !(checkdirectives(ps,vm) && !igif) { return "" }
			case "#if":
				if workif { Fatal("Double-#if in line: "+linenumber) }
				igif = !checkdirectives(ps,vm)
				workif=true
			case "#else":
				if !workif { Fatal("#else without #if in line: "+linenumber) }
				igif=!igif
			case "#fi":
				if !workif { Fatal("#fi without #if in line: "+linenumber) }
				igif=false
				workif=false
			case "#define":
				if (!igif) {
					if len(ps)<3 { Fatal("Where is the option to define in line: "+linenumber+"?") }
					switch (qstr.Left(ps[2],1)) {
						case "!","$":
							Fatal("Illegal option to define in line: "+linenumber)
						case "@":
							globdefs[ps[1]]=true
						default:
							vms[vm].defs[ps[2]]=true
					}
				}
			case "#undef","#undefine":
				if (!igif) {
					if len(ps)<3 { Fatal("Where is the option to undefine in line: "+linenumber+"?") }
					switch (qstr.Left(ps[2],1)) {
						case "!","$":
							Fatal("Illegal option to define in line: "+linenumber)
						case "@":
							globdefs[ps[1]]=false
						default:
							vms[vm].defs[ps[2]]=false
					}
				}
			default:
				Warn("Unknown preprocessor tag: "+ps[1]+" in line: "+linenumber)
			}
		}
		if igif { 
			ret+="--"  // Make sure Lua will ignore this line if instructed to. I chose this method in stead of just removing to make sure line numbers in run-time errors will remain correct!
		}
		ret +="\t"+line+"\n" 
	}
	return ret
}
