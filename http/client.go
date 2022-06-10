package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"form3/model"
	log "github.com/sirupsen/logrus"

	"form3"
)

const baseURL = "https://api.form3.tech/v1/organisation/accounts"

type client struct {
	id             string
	accountType    string
	organizationID string

	options    httpOptions
	httpClient *http.Client
}

type httpOptions struct {
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

	var defaultOptions = []HttpOption{
		Logger(log.StandardLogger()),
		Timeout(2 * time.Second),
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

	req, err := http.NewRequest("POST", getURL(""), encodedBody)
	if err != nil {
		return model.AccountData{}, err
	}

	res, err := c.httpClient.Do(req.WithContext(ctx))
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
	url := getURL("/" + id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.AccountData{}, err
	}

	res, err := c.httpClient.Do(req.WithContext(ctx))
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
	url := getURL("/" + accountID + "?version=" + strconv.Itoa(version))
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status == strconv.Itoa(http.StatusNotFound) {
		return errors.New("resource does not exist")
	}
	if res.Status == strconv.Itoa(http.StatusConflict) {
		return errors.New("incorrect version")
	}

	return nil
}

func encode(value interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(value)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func getURL(url string) string {
	return baseURL + url
}
