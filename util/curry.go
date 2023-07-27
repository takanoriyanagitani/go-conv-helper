package util

func CurryIII[T, U, V, W any](f func(T, U, V) W) func(T) func(U) func(V) W {
	return func(t T) func(U) func(V) W {
		return func(u U) func(V) W {
			return func(v V) W {
				return f(t, u, v)
			}
		}
	}
}
