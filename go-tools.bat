mkdir  %GOPATH%\src\golang.org\x
git clone https://github.com/golang/tools.git %GOPATH%\src\golang.org\x\tools

go get -v github.com/mdempsky/gocode
go get -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/acroca/go-symbols
go get -v golang.org/x/tools/cmd/guru
go get -v golang.org/x/tools/cmd/gorename
go get -v github.com/derekparker/delve/cmd/dlv
go get -v github.com/stamblerre/gocode
go get -v github.com/rogpeppe/godef
go get -v github.com/ianthehat/godef
go get -v github.com/sqs/goreturns
go get -v github.com/golang/lint
git clone https://github.com/golang/lint.git %GOPATH%\src\golang.org\x\lint

go build -o %GOPATH%\bin\gocode.exe github.com/mdempsky/gocode
go build -o %GOPATH%\bin\gopkgs.exe github.com/uudashr/gopkgs/cmd/gopkgs
go build -o %GOPATH%\bin\go-outline.exe github.com/ramya-rao-a/go-outline
go build -o %GOPATH%\bin\go-symbols.exe github.com/acroca/go-symbols
go build -o %GOPATH%\bin\guru.exe golang.org/x/tools/cmd/guru
go build -o %GOPATH%\bin\gorename.exe golang.org/x/tools/cmd/gorename
go build -o %GOPATH%\bin\dlv.exe github.com/derekparker/delve/cmd/dlv
go build -o %GOPATH%\bin\gocode-gomod.exe github.com/stamblerre/gocode
go build -o %GOPATH%\bin\godef.exe github.com/rogpeppe/godef
go build -o %GOPATH%\bin\godef-gomod.exe github.com/ianthehat/godef
go build -o %GOPATH%\bin\goreturns.exe github.com/sqs/goreturns
go build -o %GOPATH%\bin\golint.exe golang.org/x/lint/golint

pause