module gitlab.com/zoralab/monteur

go 1.17

replace (
	github.com/pelletier/go-toml/v2 => ./monteur/internal/endec/toml/internal/go-toml
	gitlab.com/zoralab/monteur/gopkg => ./
)

require (
	github.com/pelletier/go-toml/v2 v2.0.0-00010101000000-000000000000
	gitlab.com/zoralab/cerigo v0.0.2
)

require (
	gitlab.com/zoralab/monteur/gopkg v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
)
