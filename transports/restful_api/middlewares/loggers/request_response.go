package loggers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

type queryHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func newQueryHeader(k string, v string) *queryHeader {
	return &queryHeader{
		Key:   k,
		Value: v,
	}
}

type requestResponse struct {
	StatusCode      int            `json:"status_code"`
	Method          string         `json:"method"`
	Latency         int64          `json:"latency"`
	Path            string         `json:"path,omitempty"`
	Queries         []*queryHeader `json:"queries,omitempty"`
	RequestHeaders  []*queryHeader `json:"request_headers,omitempty"`
	RequestBody     interface{}    `json:"request_body,omitempty"`
	ResponseHeaders []*queryHeader `json:"response_headers,omitempty"`
	ResponseBody    interface{}    `json:"response_body,omitempty"`
}

func newRequestResponse(ctx *fiber.Ctx, startTime int64, endTime int64) *requestResponse {
	var (
		r           *requestResponse
		parsedQuery url.Values
		body        interface{}
		err         error
	)

	r = &requestResponse{
		Method:     ctx.Method(),
		Path:       ctx.Path(),
		StatusCode: ctx.Response().StatusCode(),
		Latency:    endTime - startTime,
	}

	parsedQuery, err = url.ParseQuery(string(ctx.Request().URI().QueryString()))

	if err == nil {
		for k, v := range parsedQuery {
			if r.Queries == nil {
				r.Queries = []*queryHeader{}
			}

			if len(v) > 0 {
				r.Queries = append(r.Queries, newQueryHeader(k, strings.Join(v, ",")))
			}
		}
	}

	for k, v := range ctx.GetReqHeaders() {
		if r.RequestHeaders == nil {
			r.RequestHeaders = []*queryHeader{}
		}

		r.RequestHeaders = append(r.RequestHeaders, newQueryHeader(k, v))
	}

	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		err = sonic.Unmarshal(ctx.Body(), &body)

		if err == nil {
			r.RequestBody = body
		}
	}

	ctx.Response().Header.VisitAll(func(k, v []byte) {
		if r.ResponseHeaders == nil {
			r.ResponseHeaders = []*queryHeader{}
		}

		r.ResponseHeaders = append(r.ResponseHeaders, newQueryHeader(string(k), string(v)))
	})
	err = sonic.Unmarshal(ctx.Response().Body(), &body)

	if err == nil {
		r.ResponseBody = body
	}

	return r
}
