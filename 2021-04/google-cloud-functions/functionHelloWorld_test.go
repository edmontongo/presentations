package google_cloud_functions

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/hello", nil)

	HelloWorld(recorder, request)

	require.Contains(t, recorder.Body.String(), "Hello, world!")
}
