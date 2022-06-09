package types

func Ref[T any](t T) *T {
	return &t
}

func Ok(err error) {
	if err != nil {
		panic(err)
	}
}

func Unwrap[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

func UnwrapError[T any](t T, err error) error {
	return err
}
