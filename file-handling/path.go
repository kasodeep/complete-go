package handling

import (
	"fmt"
	"path/filepath"
	"strings"
)

func PathsDemo() {
	// joining paths.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// it will also take into account the ..
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// finding the directories and file.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// absolute vs relative paths.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// getting the extension of the file.
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	// returns a relative path between the base and the target.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
