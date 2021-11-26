package attributes

import (
	"github.com/pkg/errors"
	"reflect"
)

type Attributes map[string]interface{}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func New() Attributes {
	return Attributes{}
}

func (attr Attributes) WithError(err error) Attributes {
	if err == nil {
		return attr
	}

	attr["exception.type"] = reflect.TypeOf(err).Elem().String()
	attr["exception.message"] = err.Error()

	if cause := errors.Cause(err); cause != nil {
		attr["exception.cause"] = cause.Error()
	}

	if st, ok := err.(stackTracer); ok {
		attr["exception.stacktrace"] = st.StackTrace()
	}

	return attr
}
