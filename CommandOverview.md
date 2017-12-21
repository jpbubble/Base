# Command overview

This is just a list of the commands that bubble adds to Lua.
This is by no means a guide to Lua itself, as you can find that on
http://www.lua.org



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
