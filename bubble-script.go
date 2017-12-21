// For the license of this .go file -- It's the same as the Lua script it echoes

package bubble

import(
	"trickyunits/mkl"
)


const sysscript = `
	--[[		-- line #1
	        bubble-script.lua		-- line #2
		(c) 2017 Jeroen Petrus Broks.		-- line #3
				-- line #4
		This Source Code Form is subject to the terms of the 		-- line #5
		Mozilla Public License, v. 2.0. If a copy of the MPL was not 		-- line #6
		distributed with this file, You can obtain one at 		-- line #7
		http://mozilla.org/MPL/2.0/.		-- line #8
	        Version: 17.12.21		-- line #9
	]]		-- line #10
			-- line #11
	-----		-- line #12
			-- line #13
	--[[		-- line #14
			-- line #15
	    This script is ignored by the Go compiler, but it was put in a 		-- line #16
	    separate .lua file in order to have full advantage of syntax		-- line #17
	    highlight and external parse error checkers.		-- line #18
	    		-- line #19
	    It will have to be converted into a .go file in order 		-- line #20
	    to function properly.		-- line #21
	    		-- line #22
	    Well, in a bit of self reliance this file can "build itself".		-- line #23
	    All you need is a cli Lua tool.		-- line #24
	    		-- line #25
	    If the two versions already match, hey no bother. (although		-- line #26
	    the script has been set up to make the stuff between the start		-- line #27
	    build and end build lines into a big comment in the .go		-- line #28
	    file. In the respository this will mostly be the case, but		-- line #29
	    this file has just been sent with it, as it will be handy		-- line #30
	    if this file alone is used to modify the script, also in any		-- line #31
	    possible pull requests.		-- line #32
	    		-- line #33
	    When using this lua script to build the .go file please note		-- line #34
	    it will define all functions and variables, even though all 		-- line #35
	    they can do is take up memory space. This is to make sure no		-- line #36
	    parse errors or definitions errors pop up. If they are there, 		-- line #37
	    building won't happen.		-- line #38
	    		-- line #39
	    Please note, use a regular lua cli tool, written in C or C++		-- line #40
	    Do not use the Go variant. The Go script only contains stuff		-- line #41
	    go-lua supports, but the builder does call to functions the		-- line #42
	    Go version of Lua does not yet support!		-- line #43
	    		-- line #44
	    ]]		-- line #45
	    		-- line #46
	    		-- line #47
	--[[ START BUILD		-- line #48
			-- line #49
	mkl = { version=function(a,b) lversion=b lname=a end, lic=function(a,b) llicense=b end }		-- line #50
	mkl.version("Bubble Base - bubble-script.lua","17.12.21")		-- line #51
	mkl.lic    ("Bubble Base - bubble-script.lua","Mozilla Public License 2.0")		-- line #52
			-- line #53
	function go_buildme()		-- line #54
		-- read and block out the building script.		-- line #55
		file = "bubble-script.lua"		-- line #56
		print("Converting: "..file)				-- line #57
		lines = {}		-- line #58
		for line in io.lines(file) do 		-- line #59
			if line=="-- START BUILD" then 		-- line #60
				lines[#lines + 1] = "--[[ START BUILD"		-- line #61
			else		-- line #62
				lines[#lines + 1] = line		-- line #63
			end		-- line #64
		end		-- line #65
		-- Convert into .Go and write it		-- line #66
		goversion = os.date('%y.%m.%d')		-- line #67
		bt = io.open("bubble-script.go","w")		-- line #68
		assert(bt,"Somehow it was not possible to create the file bubbles-script.go")		-- line #69
		bt:write("// For the license of this .go file -- It's the same as the Lua script it echoes\n\n")		-- line #70
		bt:write("package bubble\n\n")		-- line #71
		bt:write("import(\n\t\"trickyunits/mkl\"\n)\n\n\n")		-- line #72
		bt:write("const sysscript = "..string.char(96).."\n")		-- line #73
		for i,l in ipairs(lines) do		-- line #74
			bt:write("\t"..l.."\t\t-- line #"..i.."\n")		-- line #75
		end		-- line #76
		bt:write(string.char(96).."\n\n")		-- line #77
		bt:write("func init(){\n")		-- line #78
		bt:write('\tmkl.Version("Bubble Base - bubble-script.go"," '..goversion..'")\n')		-- line #79
		bt:write('\tmkl.Version("Bubble Base - bubble-script.lua","'.. lversion..'")\n')		-- line #80
		bt:write('\tmkl.Lic    ("Bubble Base - bubble-script.go"," '.. llicense..'")\n')		-- line #81
		bt:write('\tmkl.Lic    ("Bubble Base - bubble-script.lua","'.. llicense..'")\n')		-- line #82
		bt:write('}\n\n');		-- line #83
		bt:close()		-- line #84
		print("Done")		-- line #85
	end		-- line #86
			-- line #87
			-- line #88
	-- END BUILD ]]		-- line #89
			-- line #90
			-- line #91
			-- line #92
	-- Desplays Script version. Meant for debugging only		-- line #93
	function ScriptVersion()		-- line #94
		return lversion		-- line #95
	end		-- line #96
			-- line #97
			-- line #98
			-- line #99
			-- line #100
			-- line #101
			-- line #102
	--[[ START BUILD		-- line #103
	go_buildme()		-- line #104
	-- END BUILD ]]		-- line #105
`

func init(){
	mkl.Version("Bubble Base - bubble-script.go"," 17.12.21")
	mkl.Version("Bubble Base - bubble-script.lua","17.12.21")
	mkl.Lic    ("Bubble Base - bubble-script.go"," Mozilla Public License 2.0")
	mkl.Lic    ("Bubble Base - bubble-script.lua","Mozilla Public License 2.0")
}

