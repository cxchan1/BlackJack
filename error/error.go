package errs

type (

	// Error Interface
	Error interface {
		Error() string
		MarshalText() ([]byte, error)
	}

	// Type represents the Error types
	Type int

	// System is an server-side error
	System struct {
		ErrorString string
	}

	// Client is the client input error
	Client struct {
		ErrorString string
	}
)

// Error implements the error interface.
func (t System) Error() string {
	return t.ErrorString
}

// MarshalText implements the TextMarshaller
func (t System) MarshalText() ([]byte, error) {
	return []byte(t.ErrorString), nil
}

// Error implements the error interface.
func (t Client) Error() string {
	return t.ErrorString
}

// MarshalText implements the TextMarshaller
func (t Client) MarshalText() ([]byte, error) {
	return []byte(t.ErrorString), nil
}
