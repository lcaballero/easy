package runner
import (
	"encoding/json"
)



type Imports struct {
	Filename, File, Package string
	UsageRcv, ExecRcv map[string]bool
}

func NewImports(pkg, filename string) *Imports {
	return &Imports{
		Filename: filename,
		Package: pkg,
		UsageRcv: make(map[string]bool),
		ExecRcv: make(map[string]bool),
	}
}

func (m *Imports) AddUsage(name string) {
	m.UsageRcv[name] = true
}

func (m *Imports) AddExec(name string) {
	m.ExecRcv[name] = true
}

func (m *Imports) String() string {
	bb, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(bb)
}