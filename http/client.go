package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"form3/model"
	log "github.com/sirupsen/logrus"

	"form3"
)

const (
	defaultBaseURL = "https://api.staging-form3.tech"
	path           = "/v1/organisation/accounts"
)

// client implement a http around package http.Client.
//
// To use client, create an instance with NewClient.
// Use HttpOptions (Functional Options) to set every feature need it.
//
// The list of HttpOptions can be found un the correspondent go file
type client struct {
	id             string
	accountType    string
	organizationID string

	options    httpOptions
	httpClient *http.Client
}

type httpOptions struct {
	headers map[string]string
	baseURL string
	logger  *log.Logger
	timeout time.Duration
}

func NewClient(id string, accountType string, organizationId string, opts ...HttpOption) http3_go_lib.Client {
	c := &client{
		id:             id,
		accountType:    accountType,
		organizationID: organizationId,
		httpClient:     &http.Client{},
	}

	defHeaders := map[string]string{
		"Content-Type": "application/vnd.api+json",
	}
	// default options need it to initialize and can be overriden
	var defaultOptions = []HttpOption{
		BaseURL(defaultBaseURL),
		Logger(log.StandardLogger()),
		Timeout(2 * time.Second),
		Headers(defHeaders),
	}

	options := append(defaultOptions, opts...)

	for _, op := range options {
		op(c)
	}

	return c
}

func (c *client) Create(ctx context.Context, name []string, country string, options ...http3_go_lib.AccountOption) (model.AccountData, error) {
	var account model.AccountData
	account.Attributes.Name = name
	account.Attributes.Country = country
	for _, op := range options {
		op(&account)
	}

	encodedBody, err := encode(account)
	if err != nil {
		return model.AccountData{}, err
	}

	req, err := http.NewRequest("POST", c.getURL(""), encodedBody)
	if err != nil {
		return model.AccountData{}, err
	}

	res, err := c.do(req.WithContext(ctx))
	if err != nil {
		return model.AccountData{}, err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&account); err != nil {
		return model.AccountData{}, err
	}

	return account, nil
}

func (c *client) Fetch(ctx context.Context, id string) (model.AccountData, error) {
	url := c.getURL("/" + id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.AccountData{}, err
	}

	res, err := c.do(req.WithContext(ctx))
	if err != nil {
		return model.AccountData{}, err
	}
	defer res.Body.Close()

	var account model.AccountData
	if err = json.NewDecoder(res.Body).Decode(&account); err != nil {
		return model.AccountData{}, err
	}

	return account, nil
}

func (c *client) Delete(ctx context.Context, accountID string, version int) error {
	url := c.getURL("/" + accountID + "?version=" + strconv.Itoa(version))
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	res, err := c.do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (c *client) do(req *http.Request) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("account api error: %s", res.Status)
	}

	return res, nil
}

func encode(value interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(value)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func (c *client) getURL(url string) string {
	return c.options.baseURL + path + url
}
