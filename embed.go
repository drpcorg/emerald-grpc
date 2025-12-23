package emerald_grpc

import (
	"embed"
	"io/fs"
)

//go:embed proto
var protoFS embed.FS

func GetProtoEmbedded() fs.FS {
	return protoFS
}
