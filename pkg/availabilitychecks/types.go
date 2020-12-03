package availabilitychecks

type CheckType string

const (
	TypeHTTP     CheckType = "http"
	TypeRegistry CheckType = "registry"
)
