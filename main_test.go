package main

import "testing"

func test(t *testing.T) {
	result, err := selecDataReturnJsonFormat("SELECT * FROM t016ffukzsi0y5ie.master_technical_order")
	if err != nil {
		println(result)
	}
}
