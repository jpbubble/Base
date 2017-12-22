# Command overview

This is just a list of the commands that bubble adds to Lua.
This is by no means a guide to Lua itself, as you can find that on
http://www.lua.org

## Lua functions commands


### CRASH(errormessage)

Will crash the application with an error message.
Contrary to Lua own features the Bubble Features will be called to make it happen and they can be different per application using Bubble.

### Use(Script)

Use will take another script from the used JCR resource and add it to the existing VM.
- Use does not return anything, it just loads and executes and if new functions and variables are defined they are added to the existing VM
- If the script has already been loaded before within that VM it will not be loaded again. This makes Use() safe to use in basically all called subscripts and it also prevent "cyclic reading" which could lead to a stack overflow (which finally breaks an otherwise infinit loop).
- Use() cannot look outside of its JCR resource.
- Since JCR is case INSENSITIVE so in the Use() command.
- If the loaded script contains parse errors or if the script is not even present within the JCR resource a fatal error will be thrown.
- File names can be given a few directives by using "__" as separators. A file having "linux" will only be seen when you are on linus and when it has "not_linux" it will be ignored by Linux but not by other systems. You get the idea?


### UseDir(Dir)

Use all scripts in a certain directory
- All notes set in Use() basically apply here as well




## Preprocessor commands

Bubble does feature a preprocessor.
In order not to make them spook up any syntax highlighters or parse checkers, the pre-processor commands are all brought in comment form.

~~~Lua
-- #if $darwin
	print("I am on mac")
-- #else
	print("I am not on mac")
-- #fi
~~~
This is the basic idea. All preprocessor commands are prefixed with # in order to make Bubble able to tell them apart from regular comments.

Below is an overview of them

### -- #define opt

Define a custom value for preprocessing.
options should only contain numbers, underscores and roman letters.
prefix "$" is reserved for system definitions so using those in a define statement will make Bubble throw an error
prefix "@" means the definition will affect all scripts loaded regardless of the vm loading them, all other defs are tied to the vm loading them
prefix "!" should never be used as #if and #demand use it as a "not" statement.

### -- #undef opt

Undoes #define
This can also be used on options not defined before.

### -- #if opt [opt2] [opt3]

If all options (in AND form) are set the code coming next will be read. If not it will be ignored. You can prefix an option with "!" to make it act as a "not".
You can only have one level of #if. Trying to set more levels will lead to trouble!

### -- #else

if the last #if was false then the next code will be read, if it was true it will be ignored. Ah you know how if..else... statements work.
Please note there is no elseif for this (nor is it planned).


### -- #fi

Ends the last if.

### -- #demand opt [opt2] [opt3]

Almost the same as "if" but where "if" only affects the code until "fi", demand will ignore the complete code file if the definitions turn out a "false".
It does not matter where in the script this tag is located, it will always produce the same results, with the exception of #define and #undef commands affecting global defs, the the ones prior to the demand command will always be executed.
Of course if speed is important in the pre-processing, it can be recommended to place it as much on top of your code as possible, as it anything pre-processed before will get lost and anything after that will just not be processed.

