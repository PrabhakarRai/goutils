# Color
color package for go
====================

This is a package to change the color of the text and background in the console, working both under Windows and other systems.

Under Windows, the console APIs are used. Otherwise, ANSI texts are output.


Usage:
```go
color.Foreground(Green, false)
fmt.Println("Green text starts here...")
color.ChangeColor(Red, true, White, false)
fmt.Println(...)
color.ResetColor()
```

This package is light version of - https://github.com/daviddengcn/go-colortext