# hack_assembler
HACK Assembler by Go

This repository is to the assembler that convert HACK Assemblies into binary code.

Since only considering the completion of project 6 of nand2tetris, so this don't consider any invalid assembly cases.

# execute
```
go run .\main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\Add.asm" --out="{{REPLACE_YOUR_DIRECTORY}}\Add.hack"
go run .\main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\Max.asm" --out="{{REPLACE_YOUR_DIRECTORY}}\Max.hack"
go run .\main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\MaxL.asm" --out="{{REPLACE_YOUR_DIRECTORY}}\MaxL.hack"
go run .\main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\Pong.asm" --out="{{REPLACE_YOUR_DIRECTORY}}\Pong.hack"
go run .\main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\PongL.asm" --out="{{REPLACE_YOUR_DIRECTORY}}\PongL.hack"
go run .\main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\Rect.asm" --out="{{REPLACE_YOUR_DIRECTORY}}\Rect.hack"
go run .\main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\RectL.asm" --out="{{REPLACE_YOUR_DIRECTORY}}\RectL.hack"
```