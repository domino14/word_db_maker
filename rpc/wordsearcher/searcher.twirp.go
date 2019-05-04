// Code generated by protoc-gen-twirp v5.7.0, DO NOT EDIT.
// source: rpc/wordsearcher/searcher.proto

/*
Package wordsearcher is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v5.7.0.

It is generated from these files:
	rpc/wordsearcher/searcher.proto
*/
package wordsearcher

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"
import strconv "strconv"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Imports only used by utility functions:
import io "io"
import json "encoding/json"
import url "net/url"

// ==========================
// QuestionSearcher Interface
// ==========================

// QuestionSearcher service searches for questions (duh!)
type QuestionSearcher interface {
	// Search takes in a search request and returns an unexpanded
	// search response.
	Search(context.Context, *SearchRequest) (*UnexpandedSearchResponse, error)

	// Expand takes in an unexpanded search response and returns a
	// search response (fully expanded).
	Expand(context.Context, *UnexpandedSearchResponse) (*SearchResponse, error)
}

// ================================
// QuestionSearcher Protobuf Client
// ================================

type questionSearcherProtobufClient struct {
	client HTTPClient
	urls   [2]string
}

// NewQuestionSearcherProtobufClient creates a Protobuf client that implements the QuestionSearcher interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewQuestionSearcherProtobufClient(addr string, client HTTPClient) QuestionSearcher {
	prefix := urlBase(addr) + QuestionSearcherPathPrefix
	urls := [2]string{
		prefix + "Search",
		prefix + "Expand",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &questionSearcherProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &questionSearcherProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *questionSearcherProtobufClient) Search(ctx context.Context, in *SearchRequest) (*UnexpandedSearchResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "wordsearcher")
	ctx = ctxsetters.WithServiceName(ctx, "QuestionSearcher")
	ctx = ctxsetters.WithMethodName(ctx, "Search")
	out := new(UnexpandedSearchResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionSearcherProtobufClient) Expand(ctx context.Context, in *UnexpandedSearchResponse) (*SearchResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "wordsearcher")
	ctx = ctxsetters.WithServiceName(ctx, "QuestionSearcher")
	ctx = ctxsetters.WithMethodName(ctx, "Expand")
	out := new(SearchResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ============================
// QuestionSearcher JSON Client
// ============================

type questionSearcherJSONClient struct {
	client HTTPClient
	urls   [2]string
}

// NewQuestionSearcherJSONClient creates a JSON client that implements the QuestionSearcher interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewQuestionSearcherJSONClient(addr string, client HTTPClient) QuestionSearcher {
	prefix := urlBase(addr) + QuestionSearcherPathPrefix
	urls := [2]string{
		prefix + "Search",
		prefix + "Expand",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &questionSearcherJSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &questionSearcherJSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *questionSearcherJSONClient) Search(ctx context.Context, in *SearchRequest) (*UnexpandedSearchResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "wordsearcher")
	ctx = ctxsetters.WithServiceName(ctx, "QuestionSearcher")
	ctx = ctxsetters.WithMethodName(ctx, "Search")
	out := new(UnexpandedSearchResponse)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionSearcherJSONClient) Expand(ctx context.Context, in *UnexpandedSearchResponse) (*SearchResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "wordsearcher")
	ctx = ctxsetters.WithServiceName(ctx, "QuestionSearcher")
	ctx = ctxsetters.WithMethodName(ctx, "Expand")
	out := new(SearchResponse)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ===============================
// QuestionSearcher Server Handler
// ===============================

type questionSearcherServer struct {
	QuestionSearcher
	hooks *twirp.ServerHooks
}

func NewQuestionSearcherServer(svc QuestionSearcher, hooks *twirp.ServerHooks) TwirpServer {
	return &questionSearcherServer{
		QuestionSearcher: svc,
		hooks:            hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *questionSearcherServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// QuestionSearcherPathPrefix is used for all URL paths on a twirp QuestionSearcher server.
// Requests are always: POST QuestionSearcherPathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const QuestionSearcherPathPrefix = "/twirp/wordsearcher.QuestionSearcher/"

func (s *questionSearcherServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "wordsearcher")
	ctx = ctxsetters.WithServiceName(ctx, "QuestionSearcher")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/wordsearcher.QuestionSearcher/Search":
		s.serveSearch(ctx, resp, req)
		return
	case "/twirp/wordsearcher.QuestionSearcher/Expand":
		s.serveExpand(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *questionSearcherServer) serveSearch(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSearchJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSearchProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *questionSearcherServer) serveSearchJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Search")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(SearchRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to parse request json"))
		return
	}

	// Call service method
	var respContent *UnexpandedSearchResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.QuestionSearcher.Search(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UnexpandedSearchResponse and nil error while calling Search. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *questionSearcherServer) serveSearchProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Search")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(SearchRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to parse request proto"))
		return
	}

	// Call service method
	var respContent *UnexpandedSearchResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.QuestionSearcher.Search(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UnexpandedSearchResponse and nil error while calling Search. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *questionSearcherServer) serveExpand(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveExpandJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveExpandProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *questionSearcherServer) serveExpandJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Expand")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(UnexpandedSearchResponse)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to parse request json"))
		return
	}

	// Call service method
	var respContent *SearchResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.QuestionSearcher.Expand(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SearchResponse and nil error while calling Expand. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *questionSearcherServer) serveExpandProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Expand")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(UnexpandedSearchResponse)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to parse request proto"))
		return
	}

	// Call service method
	var respContent *SearchResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.QuestionSearcher.Expand(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SearchResponse and nil error while calling Expand. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *questionSearcherServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *questionSearcherServer) ProtocGenTwirpVersion() string {
	return "v5.7.0"
}

func (s *questionSearcherServer) PathPrefix() string {
	return QuestionSearcherPathPrefix
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// github.com/golang/protobuf/protoc-gen-go/descriptor.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)
	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string
	// PathPrefix returns the HTTP URL path prefix for all methods handled by this
	// service. This can be used with an HTTP mux to route twirp requests
	// alongside non-twirp requests on one HTTP listener.
	PathPrefix() string
}

// WriteError writes an HTTP response with a valid Twirp error format (code, msg, meta).
// Useful outside of the Twirp server (e.g. http middleware), but does not trigger hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	respBody := marshalErrorToJSON(twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBody)))
	resp.WriteHeader(statusCode) // set HTTP status code and send response

	_, writeErr := resp.Write(respBody)
	if writeErr != nil {
		// We have three options here. We could log the error, call the Error
		// hook, or just silently ignore the error.
		//
		// Logging is unacceptable because we don't have a user-controlled
		// logger; writing out to stderr without permission is too rude.
		//
		// Calling the Error hook would confuse users: it would mean the Error
		// hook got called twice for one request, which is likely to lead to
		// duplicated log messages and metrics, no matter how well we document
		// the behavior.
		//
		// Silently ignoring the error is our least-bad option. It's highly
		// likely that the connection is broken and the original 'err' says
		// so anyway.
		_ = writeErr
	}

	callResponseSent(ctx, hooks)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v5.7.0")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrapInternal(err, "failed to read server error response body")
	}
	var tj twerrJSON
	if err := json.Unmarshal(respBodyBytes, &tj); err != nil {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg)
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429, 502, 503, 504: // Too Many Requests, Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}

func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrapInternal wraps an error with a prefix as an Internal error.
// The original error cause is accessible by github.com/pkg/errors.Cause.
func wrapInternal(err error, prefix string) twirp.Error {
	return twirp.InternalErrorWith(&wrappedError{prefix: prefix, cause: err})
}

type wrappedError struct {
	prefix string
	cause  error
}

func (e *wrappedError) Cause() error  { return e.cause }
func (e *wrappedError) Error() string { return e.prefix + ": " + e.cause.Error() }

// ensurePanicResponses makes sure that rpc methods causing a panic still result in a Twirp Internal
// error response (status 500), and error hooks are properly called with the panic wrapped as an error.
// The panic is re-raised so it can be handled normally with middleware.
func ensurePanicResponses(ctx context.Context, resp http.ResponseWriter, hooks *twirp.ServerHooks) {
	if r := recover(); r != nil {
		// Wrap the panic as an error so it can be passed to error hooks.
		// The original error is accessible from error hooks, but not visible in the response.
		err := errFromPanic(r)
		twerr := &internalWithCause{msg: "Internal service panic", cause: err}
		// Actually write the error
		writeError(ctx, resp, twerr, hooks)
		// If possible, flush the error to the wire.
		f, ok := resp.(http.Flusher)
		if ok {
			f.Flush()
		}

		panic(r)
	}
}

// errFromPanic returns the typed error if the recovered panic is an error, otherwise formats as error.
func errFromPanic(p interface{}) error {
	if err, ok := p.(error); ok {
		return err
	}
	return fmt.Errorf("panic: %v", p)
}

// internalWithCause is a Twirp Internal error wrapping an original error cause, accessible
// by github.com/pkg/errors.Cause, but the original error message is not exposed on Msg().
type internalWithCause struct {
	msg   string
	cause error
}

func (e *internalWithCause) Cause() error                                { return e.cause }
func (e *internalWithCause) Error() string                               { return e.msg + ": " + e.cause.Error() }
func (e *internalWithCause) Code() twirp.ErrorCode                       { return twirp.Internal }
func (e *internalWithCause) Msg() string                                 { return e.msg }
func (e *internalWithCause) Meta(key string) string                      { return "" }
func (e *internalWithCause) MetaMap() map[string]string                  { return nil }
func (e *internalWithCause) WithMeta(key string, val string) twirp.Error { return e }

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// withoutRedirects makes sure that the POST request can not be redirected.
// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest makes a Protobuf request to the remote Twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return wrapInternal(err, "failed to marshal proto request")
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return wrapInternal(err, "could not build request")
	}
	resp, err := client.Do(req)
	if err != nil {
		return wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrapInternal(err, "failed to read response body")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return wrapInternal(err, "failed to unmarshal proto response")
	}
	return nil
}

// doJSONRequest makes a JSON request to the remote Twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return wrapInternal(err, "failed to marshal json request")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return wrapInternal(err, "could not build request")
	}
	resp, err := client.Do(req)
	if err != nil {
		return wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return wrapInternal(err, "failed to unmarshal json response")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}
	return nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

var twirpFileDescriptor0 = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x25, 0x8b, 0x36, 0x87, 0x8a, 0xbc, 0x1d, 0xa4, 0x29, 0xab, 0xa6, 0x8d, 0xca, 0x22,
	0x89, 0x0e, 0x85, 0x0d, 0xa8, 0x87, 0x5e, 0x8a, 0x02, 0x94, 0xc1, 0x48, 0x44, 0x28, 0xca, 0x5d,
	0xca, 0x4e, 0xda, 0x0b, 0xb1, 0x92, 0x68, 0x9b, 0x88, 0xb4, 0x54, 0x48, 0xa5, 0x95, 0x9f, 0xa3,
	0x40, 0xdf, 0x23, 0xf7, 0x3e, 0x5b, 0x51, 0xec, 0x92, 0x14, 0x29, 0xb7, 0x11, 0x7c, 0xdb, 0xfd,
	0xe6, 0x9b, 0x6f, 0xfe, 0x96, 0x43, 0x78, 0x96, 0xac, 0x66, 0x67, 0x7f, 0xc4, 0xc9, 0x3c, 0x0d,
	0x59, 0x32, 0xbb, 0x0d, 0x93, 0xb3, 0xe2, 0x70, 0xba, 0x4a, 0xe2, 0x75, 0x8c, 0xcd, 0xaa, 0xd1,
	0x9c, 0x01, 0x5e, 0xf2, 0x70, 0xb3, 0x62, 0x7c, 0x1e, 0xce, 0x7f, 0xf9, 0x10, 0xa6, 0xeb, 0x28,
	0xe6, 0xf8, 0x14, 0x34, 0xb6, 0x58, 0xdd, 0xb2, 0x9b, 0x84, 0x2d, 0x0d, 0xa5, 0xa3, 0x74, 0x35,
	0x5a, 0x02, 0x78, 0x06, 0x0d, 0xa9, 0x61, 0xd4, 0x3a, 0x4a, 0x57, 0xef, 0x7d, 0x79, 0x5a, 0x55,
	0x3c, 0x75, 0xa3, 0x74, 0x3d, 0xbe, 0x7e, 0x23, 0x20, 0x9a, 0xf1, 0xcc, 0x8f, 0x0a, 0x68, 0xd6,
	0xd6, 0x7d, 0xbf, 0x78, 0xb7, 0x14, 0xaf, 0x77, 0xf5, 0x1e, 0xee, 0x8a, 0x0b, 0xd9, 0x5c, 0x15,
	0x9f, 0x80, 0xba, 0x08, 0xf9, 0xcd, 0xfa, 0xd6, 0xa8, 0x77, 0x94, 0x6e, 0x83, 0xe6, 0x37, 0xec,
	0x80, 0xbe, 0x4a, 0xe2, 0x29, 0x9b, 0x46, 0x8b, 0x68, 0x7d, 0x67, 0x1c, 0x4a, 0x63, 0x15, 0x42,
	0x13, 0x9a, 0xb3, 0x78, 0x39, 0x8d, 0x38, 0x13, 0xd5, 0xa6, 0x46, 0xa3, 0xa3, 0x74, 0xeb, 0x74,
	0x07, 0x33, 0xff, 0xac, 0xc1, 0xa1, 0x88, 0x86, 0x08, 0x87, 0x22, 0x5e, 0x9e, 0xa9, 0x3c, 0xef,
	0x96, 0x50, 0xbb, 0x5f, 0xc2, 0x37, 0x00, 0xf3, 0xf0, 0x3a, 0xe2, 0x91, 0x50, 0x92, 0xc9, 0x69,
	0xb4, 0x82, 0xe0, 0x33, 0xd0, 0xaf, 0x93, 0x98, 0xaf, 0x83, 0xdb, 0x38, 0x7e, 0x97, 0xca, 0x04,
	0x35, 0x0a, 0x12, 0x1a, 0x0a, 0x04, 0xbf, 0x06, 0x98, 0xb2, 0xd9, 0xbb, 0xdc, 0xde, 0xc8, 0xf4,
	0x05, 0x92, 0x99, 0x5f, 0xc2, 0xc9, 0x22, 0xdc, 0x44, 0xb3, 0x98, 0x07, 0xe9, 0xdd, 0x72, 0x1a,
	0x2f, 0x52, 0x43, 0x95, 0x9c, 0x56, 0x0e, 0xfb, 0x19, 0x8a, 0x5d, 0x20, 0x11, 0xe7, 0x61, 0x12,
	0x94, 0xe1, 0x8c, 0xa3, 0x8e, 0xd2, 0x3d, 0xa6, 0x2d, 0x89, 0xbf, 0x2a, 0x42, 0xe2, 0x0b, 0x38,
	0xc9, 0x98, 0xdb, 0xb8, 0xc6, 0xb1, 0x24, 0x3e, 0x92, 0x70, 0x3f, 0x8f, 0x6d, 0x7e, 0x0b, 0x7a,
	0x65, 0xbe, 0x95, 0xde, 0xd4, 0x8b, 0xde, 0x98, 0x7f, 0xa9, 0xf0, 0xc8, 0x97, 0xf3, 0xa2, 0xe1,
	0x7b, 0xf1, 0xa0, 0xf0, 0x35, 0x34, 0xb3, 0x01, 0xae, 0x58, 0xc2, 0x96, 0xa9, 0x64, 0xeb, 0xbd,
	0x97, 0xbb, 0x93, 0xdd, 0x71, 0xc9, 0x6f, 0x17, 0x82, 0x4f, 0x77, 0x9c, 0xdb, 0xdf, 0x83, 0x3a,
	0x8a, 0xf8, 0x88, 0x6d, 0x90, 0x40, 0x7d, 0x19, 0x71, 0x39, 0x97, 0x06, 0x15, 0x47, 0x89, 0xb0,
	0x8d, 0x1c, 0x88, 0x40, 0xd8, 0xa6, 0xfd, 0x1d, 0xe8, 0xfe, 0x3a, 0x89, 0xf8, 0xcd, 0x15, 0x5b,
	0x7c, 0x08, 0xf1, 0x31, 0x34, 0x7e, 0x17, 0x87, 0x7c, 0x98, 0xd9, 0xa5, 0xfd, 0xbc, 0x20, 0x59,
	0x49, 0xc2, 0xee, 0xc4, 0xbb, 0x92, 0x78, 0x9a, 0x97, 0x95, 0xdf, 0xda, 0x7f, 0xd7, 0x40, 0xaf,
	0xe4, 0x85, 0xe7, 0xa0, 0xcd, 0x62, 0x3e, 0xcf, 0xa6, 0x2c, 0x04, 0x5b, 0xbd, 0xe7, 0xfb, 0x6a,
	0x3a, 0x2f, 0xc8, 0xb4, 0xf4, 0xc3, 0x9f, 0x40, 0x5d, 0x46, 0xbc, 0xc8, 0x5a, 0xef, 0x99, 0xfb,
	0x14, 0xb2, 0xc2, 0x87, 0x07, 0x34, 0xf7, 0xc1, 0xd7, 0xa0, 0xa7, 0x32, 0xf3, 0xac, 0xaa, 0xba,
	0x94, 0xd8, 0xdf, 0xd8, 0xb2, 0x1b, 0xc3, 0x03, 0x5a, 0xf5, 0x2e, 0xc5, 0x98, 0x68, 0x83, 0x7c,
	0x96, 0x0f, 0x12, 0x93, 0x5d, 0x2b, 0xc5, 0xa4, 0x77, 0x9f, 0x40, 0x6b, 0x5b, 0xa4, 0x9c, 0x9c,
	0xf9, 0x8f, 0x02, 0xda, 0xb6, 0x05, 0xa8, 0xc3, 0x91, 0x6b, 0xbf, 0x75, 0xce, 0xc7, 0x1e, 0x39,
	0x40, 0x00, 0xd5, 0xb5, 0xbd, 0xc1, 0x64, 0x48, 0x14, 0xfc, 0x1c, 0x3e, 0xbb, 0xa0, 0xe3, 0xbe,
	0xd5, 0x77, 0x5c, 0x67, 0xf2, 0x6b, 0x40, 0x2d, 0x6f, 0x60, 0x93, 0x1a, 0x3e, 0x06, 0x52, 0x85,
	0x5d, 0xc7, 0x9f, 0x90, 0xfa, 0x7d, 0xb2, 0xeb, 0x8c, 0x9c, 0x09, 0x39, 0xc4, 0x27, 0x80, 0xde,
	0xe5, 0xa8, 0x6f, 0xd3, 0x60, 0xfc, 0x2a, 0xb0, 0x3c, 0x6b, 0x40, 0xad, 0x91, 0x4f, 0x1a, 0x42,
	0xa4, 0xc4, 0xaf, 0xc6, 0x6f, 0x6c, 0xd7, 0x27, 0x2a, 0x36, 0xe1, 0x78, 0x68, 0xf9, 0xc1, 0xc4,
	0x1a, 0xf8, 0xe4, 0x08, 0x4f, 0x40, 0xbf, 0x18, 0x3b, 0xde, 0x24, 0xb8, 0xb2, 0xdc, 0x4b, 0x9b,
	0x1c, 0x0b, 0xa7, 0x91, 0x35, 0x39, 0x1f, 0x3a, 0xde, 0xa0, 0xd0, 0x22, 0x1a, 0x22, 0xb4, 0x2c,
	0xf7, 0x62, 0x28, 0xaf, 0x59, 0x36, 0x20, 0x30, 0x6f, 0x3c, 0x09, 0x1c, 0x2f, 0x28, 0x4a, 0xd3,
	0x51, 0x83, 0x06, 0xb5, 0x07, 0xf6, 0x5b, 0xd2, 0x34, 0x7f, 0x03, 0xa3, 0x5c, 0xb5, 0x45, 0x23,
	0xd3, 0x55, 0xcc, 0xd3, 0x10, 0x7f, 0x06, 0xed, 0x7d, 0xbe, 0x7c, 0x8b, 0xef, 0xa3, 0xb3, 0xdb,
	0xf9, 0xff, 0x6e, 0x69, 0x5a, 0xba, 0x98, 0x0e, 0xb4, 0xee, 0x29, 0xfe, 0x08, 0xb0, 0xdd, 0x48,
	0x85, 0xe4, 0x17, 0xbb, 0x92, 0xdb, 0x95, 0x4c, 0x2b, 0xd4, 0xde, 0x47, 0x05, 0x48, 0x11, 0xc2,
	0xcf, 0xa9, 0x38, 0x02, 0x35, 0x3b, 0xe3, 0x57, 0x7b, 0x1e, 0x44, 0xfb, 0xc5, 0xa7, 0x72, 0xbe,
	0x97, 0x9c, 0x07, 0xaa, 0x2d, 0x2d, 0xf8, 0x40, 0x8f, 0xf6, 0xd3, 0xff, 0x0f, 0x9b, 0x59, 0xa7,
	0xaa, 0xfc, 0xb5, 0xfd, 0xf0, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0xd0, 0xd2, 0x12, 0xfd,
	0x06, 0x00, 0x00,
}
