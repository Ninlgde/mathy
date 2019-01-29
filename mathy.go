package mathy

func Fib(n int) int64 {
    if n <= 1 {
        return 1
	}
	var a int64 = 1
	var b int64 = 1
    for i:=2; i<=n; i++ {
		a, b = a+b, a
    }
    return a
}

func Fact(n int) int64 {
    if n <= 1 {
        return 1
    }
    var s int64 = 1
    for i := 2; i <= n; i++ {
        s *= int64(i)
    }
    return s
}
