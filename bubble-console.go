/*
        bubble-console.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.21
*/
package bubble

import(
	"trickyunits/mkl"
	"trickyunits/ansistring"
	"log"
)



type tBubbleConsole struct {
	Write func(color, txt string)
	WriteLn func(color,txt string)
	Error func(txt string, fatal ...bool) // Only meant to make the "fatal" parameter optional.
	Warn func(txt string)
}


// Contains the functions you can use to output log data.
var Console tBubbleConsole


func default_bc_write(color,txt string) {
	c:=ansistring.A_White
	f:=0
	switch color {
		case "Yellow":
			c=ansistring.A_Yellow
		case "Amber":
			c=ansistring.A_Yellow
			f=ansistring.A_Dark
		case "Blue":
			c=ansistring.A_Blue
		case "LightBlue","SkyBlue":
			c=ansistring.A_Blue
			f=ansistring.A_Bright
		case "Red":
			c=ansistring.A_Red
		case "Pink":
			c=ansistring.A_Red
			f=ansistring.A_Bright
		case "Cyan":
			c=ansistring.A_Cyan
		case "LightCyan":
			c=ansistring.A_Cyan
			f=ansistring.A_Bright
		case "Magenta":
			c=ansistring.A_Magenta
		case "Purple":
			c=ansistring.A_Magenta
			f=ansistring.A_Dark
		case "Green":
			c=ansistring.A_Green
		case "LightGreen":
			c=ansistring.A_Green
			f=ansistring.A_Bright
	}
	log.Print(ansistring.SCol(txt,c,f))
}


func default_bc_writeln(color,txt string) {
	default_bc_write(color,txt+"\n")
}

func default_bc_error(txt string,fatal ...bool) {
	default_bc_writeln("Red","ERROR!")
	default_bc_writeln("Yellow",txt)
	if len(fatal)>0 {
		if fatal[0] { log.Fatal(ansistring.SCol("This is a fatal error! Terminating!",Magenta,0)) }
	}
}

func default_bc_warn(txt string) {
	default_bc_writeln("Amber","WARNING!")
	default_bc_writeln("Yellow",txt)
}


// You can set your own functions with this to make Bubble write its logs and throw its errors.
func SetConsole(wr,wrln func(col,txt string),er func(txt,fatal ...bool),wr func(txt)) {
	Console = tBubbleConsole { Write:wr, WriteLn:wrln, Error:er, Warn:wr }
}

// Resets the console to the default setting, which is through the standard "log" out put of Go.
func ConsoleToDefault(){
	SetConsole(default_bc_write,default_bc_writeln,default_bc_error,default_bc_warn)
}

// Writes to the console
func Write(col,txt string) { Console.Write(col,txt) }

// Writes to the console and goes to a new line
func WriteLn(col,txt string) { Console.Write(col,txt) }

// Throws an error
func Error(txt string) { Console.Error(txt) }

// Throws a warning
func Warn(txt string) { Console.Warn(txt) }



func init(){
mkl.Version("Bubble Base - bubble-console.go","17.12.21")
mkl.Lic    ("Bubble Base - bubble-console.go","Mozilla Public License 2.0")
ConsoleToDefault()
}
