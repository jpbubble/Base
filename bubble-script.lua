--[[
        bubble-script.lua
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.22
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
mkl.version("Bubble Base - bubble-script.lua","17.12.22")
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
		io.write(i.."/"..#l.."\r")
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
	print("All Done!")
end


-- END BUILD ]]


-- Boolean features
boolyn = { [true] = "yes", [false] =  "no" }
boolbt = { [true] =     1, [false] =     0 }
boolon = { [true] =  "on", [false] = "off" }


-- system value
svaltypes = {

   ['string']   = function(a) return a end,
   ['number']   = function(a) return a end,
   ['table']    = function(a) return serialize('<table>',a) end,
   ['function'] = function(a) return "<function>" end,
   ['nil']      = function(a) return "<nil>" end,
   ['unknown']  = function(a) return "<unknown type: "..type(a)..">" end,
   ['boolean']  = function(a) if a then return "true" else return "false" end end
   }

function sval(a)
local f = svaltypes[type(a)] or svaltypes.unknown
return f(a)
end

function each(a) -- BLD: Can be used if you only need the values in a nummeric indexed tabled. (as ipairs will always return the indexes as well, regardeless if you need them or not)
	local i=0
	if type(a)~="table" then
		--Console.Write("Each received a "..type(a).."!",255,0,0)
		return nil
	end
	return function()
		i=i+1
		if a[i] then return a[i] end
    end
end

function ieach(a) -- BLD: Same as each, but not in reversed order
	local i=#a+1
	if type(a)~="table" then
	--Console.Write("IEach received a "..type(a).."!",255,0,0)
		return nil
	end
	return function()
		i=i-1
		if a[i] then return a[i] end
	end
end


--[[

    This function is written by Michal Kottman.
    http://stackoverflow.com/questions/15706270/sort-a-table-in-lua

]]

function spairs(t, order)
    -- collect the keys
    local keys = {}
    for k in pairs(t) do keys[#keys+1] = k end

    -- if order function given, sort by it by passing the table and keys a, b,
    -- otherwise just sort the keys 
    if order then
        table.sort(keys, function(a,b) return order(t, a, b) end)
    else
        table.sort(keys)
    end

    -- return the iterator function
    local i = 0
    return function()
        i = i + 1
        if keys[i] then
            return keys[i], t[keys[i]]
        end
    end
end


-- String features --
upper = string.upper
lower = string.lower
chr = string.char
printf = string.format
replace = string.gsub
rep = string.rep
substr = string.sub


function cprintf(a,b)
	print(printf(a,b))
end

function left(s,l)
return substr(s,1,l)
end

function right(s,l)
local ln = l or 1
local st = s or "nostring"
-- return substr(st,string.len(st)-ln,string.len(st))
return substr(st,-ln,-1)
end 

function mid(s,o,l)
local ln=l or 1
local of=o or 1
local st=s or ""
return substr(st,of,(of+ln)-1)
end 


function trim(s)
  return (s:gsub("^%s*(.-)%s*$", "%1"))
end
-- from PiL2 20.4

function findstuff(haystack,needle) -- BLD: Returns the position on which a substring (needle) is found inside a string or (array)table (haystrack). If nothing if found it will return nil.<p>Needle must be a string if haystack is a string, if haystack is a table, needle can be any type.
local ret = nil
local i
for i=1,len(haystack) do
    if type(haystack)=='table'  and needle==haystack[i] then ret = ret or i end
    if type(haystack)=='string' and needle==mid(haystack,i,len(needle)) then ret = ret or i end
    -- rint("finding needle: "..needle) if ret then print("found at: "..ret) end print("= Checking: "..i.. " >> "..mid(haystack,i,len(needle)))
    end
return ret    
end

function safestring(s)
local allowed = "qwertyuiopasdfghjklzxcvbnmmQWERTYUIOPASDFGHJKLZXCVBNM 12345678890-_=+!@#$%^&*()[]{};:|,.<>/?"
local i
local safe = true
local alt = ""
assert ( type(s)=='string' , "safestring expects a string not a "..type(s) )
for i=1,len(s) do
    safe = safe and (findstuff(allowed,mid(s,i,1))~=nil)
    alt = alt .."\\"..string.byte(mid(s,i,1),1)
    end
-- print("DEBUG: Testing string"); if safe then print("The string "..s.." was safe") else print("The string "..s.." was not safe and was reformed to: "..alt) end    
local ret = { [true] = s, [false]=alt }
-- print("returning "..ret[safe])
return ret[safe]     
end 





-- Desplays Script version. Meant for debugging only
function ScriptVersion()
	return lversion
end

function b_assert(condition,errmsg)
	if not condition then CRASH(errormsg) end
end

function JCR_Dir()
	local t = BUBBLE_JCR_DIR()
	local ggetdir,e = load(JCR_GetDir(path))
	if not ggetdir then Error(e) return nil end
	--print(type(ggetdir))
	return ggetdir()
end getdir = GetDir

function Use(scriptfile)
	b_assert(type(scriptfile)=="string","I expected a file name as a string as parameter of the use file, but I received a '"..type(scriptfile).."'")
	BUBBLE_TRUE_USE(scriptfile,BUBBLE_VM_ID)
end

function UseDir(directory)
	local d = JCR_Dir()
	for f in each(d) do
		if left(f,#directory)==directory then Use(f) end
	end
end





-- START BUILD
go_buildme()
-- END BUILD ]]
