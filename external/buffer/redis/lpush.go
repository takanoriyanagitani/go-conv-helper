package ir

import (
	"context"

	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

type LPush[K Key, V Val] func(context.Context, K, V) (length int64, e error)

func (p LPush[K, V]) CreateLimited(getter LengthGetter[K]) func(LengthChecker) LPush[K, V] {
	return func(checker LengthChecker) LPush[K, V] {
		var builder func(K) LengthLimiter = getter.NewLimiter(checker)
		return func(ctx context.Context, key K, val V) (length int64, e error) {
			var limiter LengthLimiter = builder(key)
			var tooMany error = limiter(ctx)
			return util.Select(
				func() (int64, error) { return 0, tooMany },
				func() (int64, error) { return p(ctx, key, val) },
				nil == tooMany,
			)()
		}
	}
}
