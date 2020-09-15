package nexusiq

import (
	"crypto/tls"
	"encoding/xml"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq/client"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Host     string
	BasePath *string
	Schemes  []string
	Insecure bool

	Username string
	Password string
	Timeout  time.Duration

	Debug bool

	Root *client.NexusIq

	transport *httptransport.Runtime
}

func NewClient(c *Client) *Client {
	if c == nil {
		c = &Client{}
	}

	// Default host
	if c.Host == "" {
		c.Host = "127.0.0.1:8070"
	}

	// Default base path
	if c.BasePath == nil {
		c.BasePath = swag.String(client.DefaultBasePath)
	}

	// Default schemes
	if len(c.Schemes) == 0 {
		c.Schemes = client.DefaultSchemes
	}

	// Default timeout
	if c.Timeout == 0 {
		c.Timeout = 30 * time.Second
	}

	// Client transport
	c.transport = httptransport.New(c.Host, swag.StringValue(c.BasePath), c.Schemes)

	// Register custom producers
	c.transport.Producers[runtime.XMLMime] = XMLProducer()

	// Debug
	c.transport.Debug = c.Debug

	// Ignore TLS certificate
	if c.Insecure {
		c.transport.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// Create root of generated client
	c.Root = client.New(c.transport, nil)

	return c
}

func (c *Client) BasicAuth() runtime.ClientAuthInfoWriter {
	return httptransport.BasicAuth(c.Username, c.Password)
}

// XMLProducer creates a new XML producer
// The created Producer writes the raw XML if the provided data implements XmlContent
func XMLProducer() runtime.Producer {
	return runtime.ProducerFunc(func(writer io.Writer, data interface{}) error {
		if _, is := data.(XmlContent); is {
			_, err := io.WriteString(writer, data.(XmlContent).ToXml())
			return err
		}

		enc := xml.NewEncoder(writer)
		return enc.Encode(data)
	})
}
