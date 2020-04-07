package util

type bind interface {
	Check()(error)
}
