module services/auth

go 1.17

replace sharedInterfaces => ../../../sharedInterfaces

require (
	github.com/stretchr/testify v1.7.2
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	sharedInterfaces v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
