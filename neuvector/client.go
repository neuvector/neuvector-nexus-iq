package neuvector

import (
	"crypto/tls"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	"net/http"
	"time"
)

type ClientError string

func (e ClientError) Error() string {
	return string(e)
}

// Additional client implementations beyond the generated client

type Client struct {
	Host     string
	BasePath string
	Schemes  []string
	Insecure bool

	Username string
	Password string
	Timeout  time.Duration

	Debug bool

	Root *client.NeuVector

	transport   *httptransport.Runtime
	token       *models.RESTToken
	tokenIssued time.Time
	tokenExpiry time.Time
}

func NewClient(c *Client) *Client {
	if c == nil {
		c = &Client{}
	}

	// Default host
	if c.Host == "" {
		c.Host = "127.0.0.1:10443"
	}

	// Default base path
	if c.BasePath == "" {
		c.BasePath = client.DefaultBasePath
	}

	// Default schemas
	if len(c.Schemes) == 0 {
		c.Schemes = client.DefaultSchemes
	}

	// Default timeout
	if c.Timeout == 0 {
		c.Timeout = 30 * time.Second
	}

	// Client transport
	c.transport = httptransport.New(c.Host, c.BasePath, c.Schemes)

	// Ignore TLS certificate
	if c.Insecure {
		c.transport.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// Debug
	c.transport.Debug = c.Debug

	// Create root of generated client
	c.Root = client.New(c.transport, nil)

	return c
}
