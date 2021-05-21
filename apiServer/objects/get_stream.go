package objects

import (
	"XCloud/apiServer/locate"
	"XCloud/lib/objectstream"
	"fmt"
	"io"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fall", object)
	}
	return objectstream.NewGetStream(server, object)
}
