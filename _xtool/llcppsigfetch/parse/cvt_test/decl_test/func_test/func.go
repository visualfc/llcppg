package main

import test "github.com/goplus/llcppg/_xtool/llcppsigfetch/parse/cvt_test"

func main() {
	TestFuncDecl()
}

func TestFuncDecl() {
	testCases := []string{
		`void foo();`,
		`void foo(int a);`,
		`void foo(int a,...);`,
		`float* foo(int a,double b);`,
		`static inline int add(int a, int b);`,
	}
	test.RunTest("TestFuncDecl", testCases)
}
