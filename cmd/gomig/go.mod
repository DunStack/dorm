module github.com/dunstack/dorm/cmd/gomig

go 1.22.4

replace github.com/dunstack/dorm/mig => ../../mig

require (
	github.com/dunstack/dorm/mig v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.27.5
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
)
