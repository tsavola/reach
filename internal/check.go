// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path"
	"sort"
	"strconv"
)

const CoverImportPath = "github.com/tsavola/reach/cover"

func Check(verbose bool) (ok bool, err error) {
	srcdir, err := os.Getwd()
	if err != nil {
		return
	}

	pkg, err := build.Default.Import(".", srcdir, 0)
	if err != nil {
		return
	}

	pkgdir := pkg.Dir
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, pkgdir, nil, 0)
	if err != nil {
		return
	}

	var sites []site

	for _, pkg := range pkgs {
		for name, file := range pkg.Files {
			var ident string

			for _, imp := range file.Imports {
				if imp.Path.Kind == token.STRING {
					if s, err := strconv.Unquote(imp.Path.Value); err == nil && s == CoverImportPath {
						if imp.Name != nil {
							ident = imp.Name.Name
						} else {
							ident = path.Base(CoverImportPath)
						}
						break
					}
				}
			}

			if ident != "" {
				v := &visitor{fset, name, ident, sites} // TODO: look into imports
				ast.Walk(v, file)
				sites = v.sites
			}
		}
	}

	sort.Slice(sites, func(i, j int) bool {
		return sites[i].file < sites[j].file || (sites[i].file == sites[j].file && sites[i].line < sites[j].line)
	})

	coverLock.Lock()
	defer coverLock.Unlock()

	ok = true
	width := 0

	for _, site := range sites {
		cs, reached := coverage[site]
		if reached && all(cs) {
			if !verbose {
				continue
			}
		} else {
			ok = false
		}

		if n := len(site.String()); n > width {
			width = n
		}
	}

	if ok && !verbose {
		return
	}

	for _, site := range sites {
		cs, reached := coverage[site]
		if reached && all(cs) && !verbose {
			continue
		}

		line := fmt.Sprintf(fmt.Sprintf("%%-%ds", width+1), site.String()+":")

		switch {
		case !reached:
			line += " unreached"

		case len(cs) == 0:
			line += " covered"

		default:
			delim := " coverage:"
			for _, c := range cs {
				var n int
				if c {
					n = 1
				}
				line += fmt.Sprintf("%s %d", delim, n)
				delim = ""
			}
		}

		fmt.Fprintln(os.Stderr, line)
	}

	return
}

type visitor struct {
	fset  *token.FileSet
	file  string
	ident string
	sites []site
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if id, ok := node.(*ast.Ident); ok && id.Name == v.ident && id.Obj == nil {
		v.sites = append(v.sites, site{
			v.file,
			v.fset.Position(id.NamePos).Line,
		})
	}

	return v
}

func all(cs []bool) bool {
	for _, c := range cs {
		if !c {
			return false
		}
	}

	return true
}
