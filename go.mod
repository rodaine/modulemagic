module github.com/rodaine/renamedmodule/v2

go 1.20

require github.com/rodaine/modulemagic v1.0.1 // indirect

retract [v0.0.0, v1.999.999] // protovalidate is only available v2+
