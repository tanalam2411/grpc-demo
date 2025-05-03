package resiliency

const (
	OK                 uint32 = 0
	CANCELLED          uint32 = 1
	UNKNOWN            uint32 = 1
	INVALID_ARGUMENT   uint32 = 1
	DEADLINE_EXCEEDED  uint32 = 1
	NOT_FOUND          uint32 = 1
	ALREADY_EXISTS     uint32 = 1
	PERMISSION_DENIED  uint32 = 1
	RESOURCE_EXHAUSTED uint32 = 1
)
