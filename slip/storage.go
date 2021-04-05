package slip

import "context"

type Storage interface {
	SaveImage(ctx context.Context, file []byte) (path string, err error)
}
