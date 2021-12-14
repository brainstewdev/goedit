# goedit
a line based text editor with syntax highlighting and themes
# usage
the editor is *not* ready for every day use but it can still perform basic editing
on small text files.
languages support is limited to cpp and golang for now but it can be
expanded easly.
# build and run
the editor has to be built by the user. no precompiled binaries are provided.
you should have go installed on your system and make.
if you don't have make you can still build the editor, you just need to copy configuration file by hand.
if you have make and golang and you would like to build the program you need to clone this repository,
go into the folder in which you have cloned the repo and type "make"
you will find the binary file and the configuration files in the build directory
# customise
you can add language support by adding a file called {extension of the source files}.json
inside of that you then need to insert all the keywords into a field called keywords and all of the types into a array called types
to customise existing themes or to add others you need to modify the colors.json file inside the config folder
all of the colors are in RGB format and are applied to text using the SetColors function.
