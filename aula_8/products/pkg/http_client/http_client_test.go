package httpclient_test

import (
	httpclient "github.com/jhonyzam/curso_golang/aula_8/products/pkg/http_client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewHTTPClient(t *testing.T) {
	client := httpclient.NewHTTPClient(time.Minute)
	assert.NotNil(t, client)
}
