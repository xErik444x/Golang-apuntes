The code above can be built using the command go build . and then executed either by running the hello command or by double clicking the icon. You could also bypass the compiling step and just run the code directly using go run ..

compilar:
go build -ldflags "-H windowsgui" -o myapp.exe