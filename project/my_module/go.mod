module my_module

go 1.20

replace myAPI => ../myAPI

require myAPI v0.0.0-00010101000000-000000000000 // indirect
