MKDIR  src/golang.org/x

git clone https://github.com/golang/tools.git src/golang.org/x/tools

go get -v github.com/mdempsky/gocode
go get -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/acroca/go-symbols
go get -v golang.org/x/tools/cmd/guru
go get -v golang.org/x/tools/cmd/gorename
go get -v github.com/go-delve/delve/cmd/dlv
go get -v github.com/stamblerre/gocode
go get -v github.com/rogpeppe/godef
go get -v github.com/ianthehat/godef
go get -v github.com/sqs/goreturns
go get -v github.com/golang/lint

git clone https://github.com/golang/lint.git src/golang.org/x/lint

go build -o bin/gocode.exe github.com/mdempsky/gocode
go build -o bin/gopkgs.exe github.com/uudashr/gopkgs/cmd/gopkgs
go build -o bin/go-outline.exe github.com/ramya-rao-a/go-outline
go build -o bin/go-symbols.exe github.com/acroca/go-symbols
go build -o bin/guru.exe golang.org/x/tools/cmd/guru
go build -o bin/gorename.exe golang.org/x/tools/cmd/gorename
go build -o bin/dlv.exe github.com/go-delve/delve/cmd/dlv
go build -o bin/gocode-gomod.exe github.com/stamblerre/gocode
go build -o bin/godef.exe github.com/rogpeppe/godef
go build -o bin/godef-gomod.exe github.com/ianthehat/godef
go build -o bin/goreturns.exe github.com/sqs/goreturns
go build -o bin/golint.exe golang.org/x/lint/golint

PAUSE