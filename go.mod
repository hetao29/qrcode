module main

go 1.12

replace (
	lib/daemon v1.0.0 => ./src/lib/daemon
	rsc.io/qr => github.com/hetao29/qr v0.0.0-20151201054752-25b6bdf44e67
)

require (
	github.com/SKatiyar/qr v0.0.0-20151201054752-25b6bdf44e67 // indirect
	lib/daemon v1.0.0
	rsc.io/qr v0.2.0
)
