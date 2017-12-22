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
			io.write(i.."/"..#l.."\r")		-- line #75
			bt:write("\t"..l.."\t\t-- line #"..i.."\n")		-- line #76
		end		-- line #77
		bt:write(string.char(96).."\n\n")		-- line #78
		bt:write("func init(){\n")		-- line #79
		bt:write('\tmkl.Version("Bubble Base - bubble-script.go"," '..goversion..'")\n')		-- line #80
		bt:write('\tmkl.Version("Bubble Base - bubble-script.lua","'.. lversion..'")\n')		-- line #81
		bt:write('\tmkl.Lic    ("Bubble Base - bubble-script.go"," '.. llicense..'")\n')		-- line #82
		bt:write('\tmkl.Lic    ("Bubble Base - bubble-script.lua","'.. llicense..'")\n')		-- line #83
		bt:write('}\n\n');		-- line #84
		bt:close()		-- line #85
		print("All Done!")		-- line #86
	end		-- line #87
			-- line #88
			-- line #89
	-- END BUILD ]]		-- line #90
			-- line #91
			-- line #92
	-- Boolean features		-- line #93
	boolyn = { [true] = "yes", [false] =  "no" }		-- line #94
	boolbt = { [true] =     1, [false] =     0 }		-- line #95
	boolon = { [true] =  "on", [false] = "off" }		-- line #96
			-- line #97
			-- line #98
	-- system value		-- line #99
	svaltypes = {		-- line #100
			-- line #101
	   ['string']   = function(a) return a end,		-- line #102
	   ['number']   = function(a) return a end,		-- line #103
	   ['table']    = function(a) return serialize('<table>',a) end,		-- line #104
	   ['function'] = function(a) return "<function>" end,		-- line #105
	   ['nil']      = function(a) return "<nil>" end,		-- line #106
	   ['unknown']  = function(a) return "<unknown type: "..type(a)..">" end,		-- line #107
	   ['boolean']  = function(a) if a then return "true" else return "false" end end		-- line #108
	   }		-- line #109
			-- line #110
	function sval(a)		-- line #111
	local f = svaltypes[type(a)] or svaltypes.unknown		-- line #112
	return f(a)		-- line #113
	end		-- line #114
			-- line #115
	function each(a) -- BLD: Can be used if you only need the values in a nummeric indexed tabled. (as ipairs will always return the indexes as well, regardeless if you need them or not)		-- line #116
		local i=0		-- line #117
		if type(a)~="table" then		-- line #118
			--Console.Write("Each received a "..type(a).."!",255,0,0)		-- line #119
			return nil		-- line #120
		end		-- line #121
		return function()		-- line #122
			i=i+1		-- line #123
			if a[i] then return a[i] end		-- line #124
	    end		-- line #125
	end		-- line #126
			-- line #127
	function ieach(a) -- BLD: Same as each, but not in reversed order		-- line #128
		local i=#a+1		-- line #129
		if type(a)~="table" then		-- line #130
		--Console.Write("IEach received a "..type(a).."!",255,0,0)		-- line #131
			return nil		-- line #132
		end		-- line #133
		return function()		-- line #134
			i=i-1		-- line #135
			if a[i] then return a[i] end		-- line #136
		end		-- line #137
	end		-- line #138
			-- line #139
			-- line #140
	--[[		-- line #141
			-- line #142
	    This function is written by Michal Kottman.		-- line #143
	    http://stackoverflow.com/questions/15706270/sort-a-table-in-lua		-- line #144
			-- line #145
	]]		-- line #146
			-- line #147
	function spairs(t, order)		-- line #148
	    -- collect the keys		-- line #149
	    local keys = {}		-- line #150
	    for k in pairs(t) do keys[#keys+1] = k end		-- line #151
			-- line #152
	    -- if order function given, sort by it by passing the table and keys a, b,		-- line #153
	    -- otherwise just sort the keys 		-- line #154
	    if order then		-- line #155
	        table.sort(keys, function(a,b) return order(t, a, b) end)		-- line #156
	    else		-- line #157
	        table.sort(keys)		-- line #158
	    end		-- line #159
			-- line #160
	    -- return the iterator function		-- line #161
	    local i = 0		-- line #162
	    return function()		-- line #163
	        i = i + 1		-- line #164
	        if keys[i] then		-- line #165
	            return keys[i], t[keys[i]]		-- line #166
	        end		-- line #167
	    end		-- line #168
	end		-- line #169
			-- line #170
			-- line #171
	-- String features --		-- line #172
	upper = string.upper		-- line #173
	lower = string.lower		-- line #174
	chr = string.char		-- line #175
	printf = string.format		-- line #176
	replace = string.gsub		-- line #177
	rep = string.rep		-- line #178
	substr = string.sub		-- line #179
			-- line #180
			-- line #181
	function cprintf(a,b)		-- line #182
		print(printf(a,b))		-- line #183
	end		-- line #184
			-- line #185
	function left(s,l)		-- line #186
	return substr(s,1,l)		-- line #187
	end		-- line #188
			-- line #189
	function right(s,l)		-- line #190
	local ln = l or 1		-- line #191
	local st = s or "nostring"		-- line #192
	-- return substr(st,string.len(st)-ln,string.len(st))		-- line #193
	return substr(st,-ln,-1)		-- line #194
	end 		-- line #195
			-- line #196
	function mid(s,o,l)		-- line #197
	local ln=l or 1		-- line #198
	local of=o or 1		-- line #199
	local st=s or ""		-- line #200
	return substr(st,of,(of+ln)-1)		-- line #201
	end 		-- line #202
			-- line #203
			-- line #204
	function trim(s)		-- line #205
	  return (s:gsub("^%s*(.-)%s*$", "%1"))		-- line #206
	end		-- line #207
	-- from PiL2 20.4		-- line #208
			-- line #209
	function findstuff(haystack,needle) -- BLD: Returns the position on which a substring (needle) is found inside a string or (array)table (haystrack). If nothing if found it will return nil.<p>Needle must be a string if haystack is a string, if haystack is a table, needle can be any type.		-- line #210
	local ret = nil		-- line #211
	local i		-- line #212
	for i=1,len(haystack) do		-- line #213
	    if type(haystack)=='table'  and needle==haystack[i] then ret = ret or i end		-- line #214
	    if type(haystack)=='string' and needle==mid(haystack,i,len(needle)) then ret = ret or i end		-- line #215
	    -- rint("finding needle: "..needle) if ret then print("found at: "..ret) end print("= Checking: "..i.. " >> "..mid(haystack,i,len(needle)))		-- line #216
	    end		-- line #217
	return ret    		-- line #218
	end		-- line #219
			-- line #220
	function safestring(s)		-- line #221
	local allowed = "qwertyuiopasdfghjklzxcvbnmmQWERTYUIOPASDFGHJKLZXCVBNM 12345678890-_=+!@#$%^&*()[]{};:|,.<>/?"		-- line #222
	local i		-- line #223
	local safe = true		-- line #224
	local alt = ""		-- line #225
	assert ( type(s)=='string' , "safestring expects a string not a "..type(s) )		-- line #226
	for i=1,len(s) do		-- line #227
	    safe = safe and (findstuff(allowed,mid(s,i,1))~=nil)		-- line #228
	    alt = alt .."\\"..string.byte(mid(s,i,1),1)		-- line #229
	    end		-- line #230
	-- print("DEBUG: Testing string"); if safe then print("The string "..s.." was safe") else print("The string "..s.." was not safe and was reformed to: "..alt) end    		-- line #231
	local ret = { [true] = s, [false]=alt }		-- line #232
	-- print("returning "..ret[safe])		-- line #233
	return ret[safe]     		-- line #234
	end 		-- line #235
			-- line #236
			-- line #237
			-- line #238
			-- line #239
			-- line #240
	-- Desplays Script version. Meant for debugging only		-- line #241
	function ScriptVersion()		-- line #242
		return lversion		-- line #243
	end		-- line #244
			-- line #245
	function b_assert(condition,errmsg)		-- line #246
		if not condition then CRASH(errormsg) end		-- line #247
	end		-- line #248
			-- line #249
	function JCR_Dir()		-- line #250
		local t = BUBBLE_JCR_DIR()		-- line #251
		local ggetdir,e = load(JCR_GetDir(path))		-- line #252
		if not ggetdir then Error(e) return nil end		-- line #253
		--print(type(ggetdir))		-- line #254
		return ggetdir()		-- line #255
	end getdir = GetDir		-- line #256
			-- line #257
	function Use(scriptfile)		-- line #258
		b_assert(type(scriptfile)=="string","I expected a file name as a string as parameter of the use file, but I received a '"..type(scriptfile).."'")		-- line #259
		BUBBLE_TRUE_USE(scriptfile,BUBBLE_VM_ID)		-- line #260
	end		-- line #261
			-- line #262
	function UseDir(directory)		-- line #263
		local d = JCR_Dir()		-- line #264
		for f in each(d) do		-- line #265
			if left(f,#directory)==directory then Use(f) end		-- line #266
		end		-- line #267
	end		-- line #268
			-- line #269
			-- line #270
			-- line #271
			-- line #272
			-- line #273
	--[[ START BUILD		-- line #274
	go_buildme()		-- line #275
	-- END BUILD ]]		-- line #276
`

func init(){
	mkl.Version("Bubble Base - bubble-script.go"," 17.12.22")
	mkl.Version("Bubble Base - bubble-script.lua","17.12.21")
	mkl.Lic    ("Bubble Base - bubble-script.go"," Mozilla Public License 2.0")
	mkl.Lic    ("Bubble Base - bubble-script.lua","Mozilla Public License 2.0")
}

