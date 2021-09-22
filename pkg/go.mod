module gitlab.com/zoralab/monteur

go 1.17

replace (
	github.com/pelletier/go-toml/v2 => ./monteur/internal/endec/go-toml
	gitlab.com/zoralab/monteur/pkg => ./
)

require (
	github.com/pelletier/go-toml/v2 v2.0.0-00010101000000-000000000000
	gitlab.com/zoralab/cerigo v0.0.2
	gitlab.com/zoralab/monteur/pkg v0.0.0-00010101000000-000000000000
)
