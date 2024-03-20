package ctx

import (
	"context"

	"github.com/sirupsen/logrus"
)

type logKey string

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// CTX extends the standard context from the standard library,
// integrating it with the logrus package for structured logging.
// It allows passing around a single object in function calls
// that provides both context management (like cancellation and timeouts)
// and the ability to log structured data.
type CTX struct {
	context.Context
	logrus.FieldLogger
}

// Background returns a non-nil, empty Context.
// It is typically used for initialization and in situations where there is
// no existing context (like in the main function or tests).
func Background() CTX {
	return CTX{
		Context:     context.Background(),
		FieldLogger: logrus.StandardLogger(),
	}
}

// WithValue returns a copy of the parent for adding metadata to the context and logs,
// which can then be passed down through function calls.
func WithValue(parent CTX, key string, val interface{}) CTX {
	return CTX{
		Context:     context.WithValue(parent, logKey(key), val),
		FieldLogger: parent.FieldLogger.WithField(key, val),
	}
}

// WithValues is similar to WithValue, but allows adding multiple key-value pairs at once.
func WithValues(parent CTX, kvs map[string]interface{}) CTX {
	c := parent
	for k, v := range kvs {
		c = WithValue(c, k, v)
	}
	return c
}

// WithCancel creates a new CTX and a cancel function from the provided parent CTX.
// The cancel function, when called, cancels the new CTX, releasing any resources associated with it,
// such as child goroutines waiting on the CTX's Done channel.
func WithCancel(parent CTX) (CTX, context.CancelFunc) {
	context, cancel := context.WithCancel(parent)
	return CTX{
		Context:     context,
		FieldLogger: parent.FieldLogger,
	}, cancel
}
