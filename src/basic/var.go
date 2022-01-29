package basic

func studyVar() {
	var a int
	var f float32 = 11.

	a = 10
	f = 12.0

	println(a, f)

	var i, j, k int = 1, 2, 3
	println(i, j, k)

	var (
		x = 1
		y = "string"
		z = true
	)
	println(x, y, z)

	sas := "short assignment statement"
	println(sas)

	const (
		c1 = 123
		c2 = "const string"
	)
	println(c1, c2)

	const (
		N0 = iota
		N1
		N2
		N3
	)
	println(N0, N1, N2, N3)
}
