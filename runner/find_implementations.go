package runner

import (
	"go/token"
	"go/parser"
	"go/ast"
)


type Finder struct {
	Imports []*Imports
}

func NewFinder() *Finder {
	return &Finder{
		Imports: make([]*Imports, 0),
	}
}

func (fp *Finder) NewImport(pkg, filename string) *Imports {
	imp := NewImports(pkg, filename)
	fp.Imports = append(fp.Imports, imp)
	return imp
}

func (fp *Finder) ParseFile(filename string) *ast.File {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	return f
}

func (fp *Finder) Run() {
	filename := "tasks/copy_dx_logs/copy_logs.go"
	src := NewSrc(filename)
	f := fp.ParseFile(filename)
	imp := fp.NewImport(f.Name.Name, filename)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			fp.CheckExec(src, imp, x)
			fp.CheckUsage(src, imp, x)

		case *ast.File:
			imp.File = x.Name.Name
		}
		return true
	})
}

func (fp *Finder) CheckUsage(src *Src, imp *Imports, x *ast.FuncDecl) {
	name := x.Name.Name
	params := ""
	numParams := len(x.Type.Params.List)
	result := src.Find(x.Type.Results.List, "string")
	rcv := src.FindRcv(x.Recv)

	if fp.MeetsUsage(numParams, name, params, result, rcv) {
		imp.AddUsage(rcv)
	}
}

func (fp *Finder) CheckExec(src *Src, imp *Imports, x *ast.FuncDecl) {
	name := x.Name.Name
	params := src.Find(x.Type.Params.List, "*conf.Conf")
	result := src.Find(x.Type.Results.List, "chan tasks.Feedback")
	rcv := src.FindRcv(x.Recv)

	if fp.MeetsExec(name, params, rcv, result) {
		imp.AddExec(rcv)
	}
}

func (fp *Finder) MeetsExec(name, params, rcv, result string) bool {
	hasName := name == "Exec"
	hasParams := params != ""
	hasRcv := rcv != ""
	hasResults := result != ""
	return hasName && hasParams && hasRcv && hasResults
}

func (fp *Finder) MeetsUsage(numParams int, name, params, rcv, result string) bool {
	hasNoParams := numParams == 0
	hasName := name == "Usage"
	hasParams := params == ""
	hasRcv := rcv != ""
	hasResults := result != ""
	return hasNoParams && hasName && hasParams && hasRcv && hasResults
}