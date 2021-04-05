package slip

import "context"

type Storage interface {
	StoreImage(ctx context.Context, file []byte) (path string, err error)
}
