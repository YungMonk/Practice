git submodule add https://github.com/mdempsky/gocode.git src/github.com/mdempsky/gocode
git submodule add https://github.com/uudashr/gopkgs.git src/github.com/uudashr/gopkgs
git submodule add https://github.com/karrick/godirwalk.git src/github.com/karrick/godirwalk
git submodule add https://github.com/pkg/errors.git src/github.com/pkg/errors
git submodule add https://github.com/ramya-rao-a/go-outline.git src/github.com/ramya-rao-a/go-outline
git submodule add https://github.com/acroca/go-symbols.git src/github.com/acroca/go-symbols
git submodule add https://github.com/golang/tools.git src/golang.org/x/tools
git submodule add https://github.com/go-delve/delve.git src/github.com/go-delve/delve
git submodule add https://github.com/stamblerre/gocode.git src/github.com/stamblerre/gocode
git submodule add https://github.com/rogpeppe/godef.git src/github.com/rogpeppe/godef
git submodule add https://github.com/ianthehat/godef.git src/github.com/ianthehat/godef
git submodule add https://github.com/sqs/goreturns.git src/github.com/sqs/goreturns
git submodule add https://github.com/golang/lint.git src/golang.org/x/lint

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
go build -o bin/golint.exe golang.org/x/lint

PAUSE