package filenaming

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

func Dump(filename string) {
	s := filepath.ToSlash(filename)
	dir := filepath.Dir(filename)
	fr := filepath.FromSlash(filename)
	base := filepath.Base(filename)

	data := struct {
		ToSlash   string
		Dir       string
		FromSlash string
		Base      string
	}{
		ToSlash:   s,
		Dir:       dir,
		FromSlash: fr,
		Base:      base,
	}

	bb, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bb))
}
