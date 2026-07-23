package blog

type Status int

const (
	STATUS_DRAFT = iota
	STATUS_PUBLISH
)

type UPDATE_MODE int

const (
	UPDATE_MODE_PUT = iota
	UPDATE_MODE_PATCH
)
