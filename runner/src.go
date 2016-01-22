package runner
import (
	"go/token"
	"io/ioutil"
	"log"
	"go/ast"
)

type Src struct {
	code []byte
}
func (s *Src) Code(a, b token.Pos) string {
	return string(s.code[a-1:b-1])
}
// todo: change method to return error
func NewSrc(filename string) *Src {
	bb, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	return &Src{
		code:bb,
	}
}
func (src *Src) Find(fields []*ast.Field, val string) string {
	for _,field := range fields {
		s := src.Code(field.Type.Pos(), field.Type.End())
		if s == val {
			return s
		}
	}
	return ""
}
func (src *Src) FindRcv(list *ast.FieldList) string {
	if list == nil || list.List == nil {
		return ""
	}
	for _,field := range list.List {
		return src.Code(field.Type.Pos(), field.Type.End())
	}
	return ""
}