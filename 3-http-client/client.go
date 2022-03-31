package cloudpayments

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	defaultBaseURL = "https://api.cloudpayments.ru/"
)

type client struct {
	PublicId   string
	APISecret  string
	HTTPClient *http.Client
}

func NewClient(publicId, apiSecret string) *client {
	return &client{
		PublicId:   publicId,
		APISecret:  apiSecret,
		HTTPClient: &http.Client{},
	}
}

func (c *client) sendRequest(endpoint, requestId string, v, body interface{}) (*http.Response, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPost, defaultBaseURL+endpoint, buf)
	if err != nil {
		return nil, err
	}
	if requestId != "" {
		req.Header.Set("X-Request-ID", requestId)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.SetBasicAuth(c.PublicId, c.APISecret)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, errors.New("invalid credentials")
	case http.StatusTooManyRequests:
		return nil, errors.New("too many requests")
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, res.Body)
	default:
		decErr := json.NewDecoder(res.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return res, err
}

func (c *client) Test(requestId string) error {
	res := new(CPResponse)
	if _, err := c.sendRequest("test", requestId, res, nil); err != nil {
		return err
	}
	if !res.Success {
		return errors.New(res.Message)
	}
	return nil
}

func (c *client) ChargeCard(input ChargeCardInput, requestId string) (*Transaction, *Secure3D, error) {
	endpoint := "payments/cards/charge"
	if input.RequireConfirmation {
		endpoint = "payments/cards/auth"
	}

	res := new(CPResponse)
	if _, err := c.sendRequest(endpoint, requestId, res, input); err != nil {
		return nil, nil, err
	}

	m, _ := json.Marshal(res.Model)

	if res.Message != "" {
		return nil, nil, errors.New(res.Message)
	}

	if !res.Success && res.Model["PaReq"] != "" {
		var s Secure3D

		if err := json.Unmarshal(m, &s); err != nil {
			return nil, nil, err
		}
		return nil, &s, nil
	}

	var t Transaction

	if err := json.Unmarshal(m, &t); err != nil {
		return nil, nil, err
	}
	if t.ReasonCode > 0 {
		return nil, nil, NewError(t.ReasonCode)
	}
	return &t, nil, nil
}
