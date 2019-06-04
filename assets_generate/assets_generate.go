/*
   Copyright 2019 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	flag.Parse()
	pkg := flag.Arg(0)
	path := flag.Arg(1)
	// bundle files
	err := vfsgen.Generate(http.Dir(path), vfsgen.Options{
		PackageName: pkg,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
