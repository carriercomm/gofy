package main

import "runtime"

func hex(n uint64, b bool) string {
	if n == 0 {
		if b {
			return "0"
		} else {
			return ""
		}
	}
	return hex(n / 16, false) + string("0123456789ABCDEF"[n & 15])
}

func fuck(s string) {
	println("SHIT IS BROKEN")
	println(s)
	runtime.Halt()
}


func main() {
	var initp Process
//	runtime.EndCritical()
	initrd := make(Initrd)
	initrd["hello"] = testbinary[:]
	rootns := Namespace{NamespaceEntry{string: "/", Filesystem: initrd}}
	f, err := rootns.Open("/hello", ORD, 0)
	if err != nil {
		println(err.String())
		for {}
	}
	err = initp.Exec(f)
	if err != nil {
		println(err.String())
		for {}
	}
	initp.Run()
}
