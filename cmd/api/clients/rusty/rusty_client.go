package rusty

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"net/http"
	"net/url"
	"taska-core-me-go/cmd/api/models"
	"time"

	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/rusty"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/transport/http_client"
)

//go:generate mockgen -destination=../../testutils/mocks/rusty_client_mock.go -package=mocks -source=./rusty_client.go

type IRustyClient interface {
	Get(ctx context.Context, url string, headers map[string]string, queryParams map[string]string) (*models.RustyResponse, error)
	Post(ctx context.Context, url string, headers map[string]string, body interface{}) (*models.RustyResponse, error)
	Patch(ctx context.Context, url string, headers map[string]string, body interface{}) (*models.RustyResponse, error)
	Put(ctx context.Context, url string, headers map[string]string, body interface{}) (*models.RustyResponse, error)
	Delete(ctx context.Context, url string, headers map[string]string, params map[string]interface{}) (*models.RustyResponse, error)
}

type RustyClientConfig struct {
	DefaultTimeOut time.Duration
	RetryCount     int
}

type RustyClient struct {
	cfg RustyClientConfig
}

func NewRustyClient(cfg RustyClientConfig) *RustyClient {
	return &RustyClient{cfg: cfg}
}

func (c *RustyClient) newRequester() *http_client.RetryableClient {
	return http_client.NewRetryable(
		c.cfg.RetryCount,
		http_client.WithTimeout(c.cfg.DefaultTimeOut),
	)
}

func (c *RustyClient) getEndpointOptions(headers map[string]string) []rusty.EndpointOption {
	opts := make([]rusty.EndpointOption, 0, len(headers))
	for k, v := range headers {
		opts = append(opts, rusty.WithHeader(k, v))
	}
	return opts
}

func (c *RustyClient) getQueryParamsOptions(params map[string]string) url.Values {
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}
	return q
}

func (c *RustyClient) getRequestOptions(params map[string]interface{}) []rusty.RequestOption {
	opts := make([]rusty.RequestOption, 0, len(params))
	for k, v := range params {
		opts = append(opts, rusty.WithParam(k, v))
	}
	return opts
}

func (c *RustyClient) generateResponse(ctx context.Context, response *rusty.Response, err error) (*models.RustyResponse, error) {
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[RUSTY] error: %v", err), err)
		return nil, err
	}

	if response == nil {
		return nil, fmt.Errorf("empty response from server")
	}

	logger.Info(ctx, fmt.Sprintf("[RUSTY] response status: %d", response.StatusCode))
	return &models.RustyResponse{
		Body:       response.Body,
		StatusCode: response.StatusCode,
	}, nil
}

func (c *RustyClient) doRequestWithBody(ctx context.Context, method, urlStr string, headers map[string]string, body interface{}) (*models.RustyResponse, error) {
	h := make(map[string]string, len(headers)+1)
	for k, v := range headers {
		h[k] = v
	}
	h["Content-Type"] = "application/json"

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error serializing body: %w", err)
	}

	logger.Info(ctx, fmt.Sprintf("[RUSTY] %s %s, headers=%v, body=%s", method, urlStr, h, string(bodyJSON)))
	requester := c.newRequester()
	endpoint, err := rusty.NewEndpoint(requester, urlStr, c.getEndpointOptions(h)...)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	var response *rusty.Response
	switch method {
	case http.MethodPost:
		response, err = endpoint.Post(ctx, rusty.WithBody(body))
	case http.MethodPatch:
		response, err = endpoint.Patch(ctx, rusty.WithBody(body))
	case http.MethodPut:
		response, err = endpoint.Put(ctx, rusty.WithBody(body))
	default:
		return nil, fmt.Errorf("unsupported method: %s", method)
	}

	return c.generateResponse(ctx, response, err)
}

func (c *RustyClient) Get(ctx context.Context, url string, headers map[string]string, queryParams map[string]string) (*models.RustyResponse, error) {
	logger.Info(ctx, fmt.Sprintf("[RUSTY] GET %s", url))
	requester := c.newRequester()

	endpoint, err := rusty.NewEndpoint(requester, url, c.getEndpointOptions(headers)...)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	response, err := endpoint.Get(ctx, rusty.WithQuery(c.getQueryParamsOptions(queryParams)))
	return c.generateResponse(ctx, response, err)
}

func (c *RustyClient) Post(ctx context.Context, url string, headers map[string]string, body interface{}) (*models.RustyResponse, error) {
	return c.doRequestWithBody(ctx, http.MethodPost, url, headers, body)
}

func (c *RustyClient) Patch(ctx context.Context, url string, headers map[string]string, body interface{}) (*models.RustyResponse, error) {
	return c.doRequestWithBody(ctx, http.MethodPatch, url, headers, body)
}

func (c *RustyClient) Put(ctx context.Context, url string, headers map[string]string, body interface{}) (*models.RustyResponse, error) {
	return c.doRequestWithBody(ctx, http.MethodPut, url, headers, body)
}

func (c *RustyClient) Delete(ctx context.Context, url string, headers map[string]string, params map[string]interface{}) (*models.RustyResponse, error) {
	logger.Info(ctx, fmt.Sprintf("[RUSTY] DELETE %s", url))
	requester := c.newRequester()

	endpoint, err := rusty.NewEndpoint(requester, url, c.getEndpointOptions(headers)...)
	if err != nil {
		return nil, fmt.Errorf("failed to create endpoint: %w", err)
	}

	response, err := endpoint.Delete(ctx, c.getRequestOptions(params)...)
	return c.generateResponse(ctx, response, err)
}
