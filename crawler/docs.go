package main

import (
	"github.com/daviddengcn/gcse"
	"log"
)

var docDB *gcse.MemDB

const (
	fieldImports = "i"
)

func processDocument(d *gcse.DocInfo) error {
	pkg := d.Package

	// fetch saved DocInfo
	var savedD gcse.DocInfo
	exists := docDB.Get(pkg, &savedD)
	if exists && d.StarCount < 0 {
		d.StarCount = savedD.StarCount
	}
	if d.StarCount < 0 {
		d.StarCount = 0
	}

	log.Printf("Package %s processed!", pkg)

	docDB.Put(pkg, *d)

	return nil
}