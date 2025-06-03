package creator

import "os"

type FileCreator interface {
	Create() (*os.File, error)
}
