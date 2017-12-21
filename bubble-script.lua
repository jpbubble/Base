--[[
        bubble-script.lua
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.21
]]

-----

--[[

    This script is ignored by the Go compiler, but it was put in a 
    separate .lua file in order to have full advantage of syntax
    highlight and external parse error checkers.
    
    It will have to be converted into a .go file in order 
    to function properly.
    
    Well, in a bit of self reliance this file can "build itself".
    All you need is a cli Lua tool.
    
    If the two versions already match, hey no bother. (although
    the script has been set up to make the stuff between the start
    build and end build lines into a big comment in the .go
    file. In the respository this will mostly be the case, but
    this file has just been sent with it, as it will be handy
    if this file alone is used to modify the script, also in any
    possible pull requests.
    
    When using this lua script to build the .go file please note
    it will define all functions and variables, even though all 
    they can do is take up memory space. This is to make sure no
    parse errors or definitions errors pop up. If they are there, 
    building won't happen.
    
    Please note, use a regular lua cli tool, written in C or C++
    Do not use the Go variant. The Go script only contains stuff
    go-lua supports, but the builder does call to functions the
    Go version of Lua does not yet support!
    
    ]]
    
    
-- START BUILD

mkl = { version=function(a,b) lversion=b lname=a end, lic=function(a,b) llicense=b end }
mkl.version("Bubble Base - bubble-script.lua","17.12.21")
mkl.lic    ("Bubble Base - bubble-script.lua","Mozilla Public License 2.0")

function go_buildme()
	-- read and block out the building script.
	file = "bubble-script.lua"
	print("Converting: "..file)		
	lines = {}
	for line in io.lines(file) do 
		if line=="-- START BUILD" then 
			lines[#lines + 1] = "--[[ START BUILD"
		else
			lines[#lines + 1] = line
		end
	end
	-- Convert into .Go and write it
	goversion = os.date('%y.%m.%d')
	bt = io.open("bubble-script.go","w")
	assert(bt,"Somehow it was not possible to create the file bubbles-script.go")
	bt:write("// For the license of this .go file -- It's the same as the Lua script it echoes\n\n")
	bt:write("package bubble\n\n")
	bt:write("import(\n\t\"trickyunits/mkl\"\n)\n\n\n")
	bt:write("const sysscript = "..string.char(96).."\n")
	for i,l in ipairs(lines) do
		bt:write("\t"..l.."\t\t-- line #"..i.."\n")
	end
	bt:write(string.char(96).."\n\n")
	bt:write("func init(){\n")
	bt:write('\tmkl.Version("Bubble Base - bubble-script.go"," '..goversion..'")\n')
	bt:write('\tmkl.Version("Bubble Base - bubble-script.lua","'.. lversion..'")\n')
	bt:write('\tmkl.Lic    ("Bubble Base - bubble-script.go"," '.. llicense..'")\n')
	bt:write('\tmkl.Lic    ("Bubble Base - bubble-script.lua","'.. llicense..'")\n')
	bt:write('}\n\n');
	bt:close()
	print("Done")
end


-- END BUILD ]]



-- Desplays Script version. Meant for debugging only
function ScriptVersion()
	return lversion
end






-- START BUILD
go_buildme()
-- END BUILD ]]
