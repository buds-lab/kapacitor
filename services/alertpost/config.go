package alertpost

import (
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Config struct {
	Endpoint string            `toml:"endpoint" override:"endpoint"`
	URL      string            `toml:"url" override:"url"`
	Headers  map[string]string `toml:"headers" override:"headers,redact"`
}

// TODO: fix
func NewConfig() Config {
	return Config{
		Endpoint: "test",
		URL:      "http://localhost:3000",
	}
}

func (c Config) NewRequest(body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", c.URL, body)
	if err != nil {
		return nil, err
	}

	for k, v := range c.Headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

func (c Config) Validate() error {
	if c.Endpoint == "" {
		return errors.New("must specify endpoint name")
	}

	if c.URL == "" {
		return errors.New("must specify url")
	}

	if _, err := url.Parse(c.URL); err != nil {
		return errors.Wrapf(err, "invalid URL %q", c.URL)
	}

	return nil
}

type Configs []Config

func (cs Configs) Validate() error {
	for _, c := range cs {
		err := c.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// Will create some garbage, but That should be fine
func (cs Configs) index() map[string]Config {
	m := map[string]Config{}

	for _, c := range cs {
		m[c.Endpoint] = c
	}

	return m
}
