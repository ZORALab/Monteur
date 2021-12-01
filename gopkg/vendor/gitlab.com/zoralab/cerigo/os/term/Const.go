package term

// statusID for all PrintStatus status value.
const (
	// NoTagStatus is the statusID for PrintStatus not to append any tag
	// This is the Default.
	NoTagStatus = uint(0)

	// InfoStatus is the statusID for PrintStatus to append tag
	// [ INFO ]
	InfoStatus = uint(1)

	// ErrorStatus is the statusID for PrintStatus to append tag
	// [ ERROR ]
	ErrorStatus = uint(2)

	// WarningStatus is the statusID for PrintStatus to append tag
	// [ WARNING ]
	WarningStatus = uint(3)

	// DebugStatus is the statusID for PrintStatus to append tag
	// [ DEBUG ]
	DebugStatus = uint(4)
)
