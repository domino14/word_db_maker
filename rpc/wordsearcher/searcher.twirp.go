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
	// Search takes in a search request and returns a search response.
	// This response can be expanded or not, depending on the `expand` field
	// in SearchRequest.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)

	// Expand takes in an unexpanded search response and returns a
	// search response (fully expanded). See expandedRepr above in
	// the Alphagram field.
	Expand(context.Context, *SearchResponse) (*SearchResponse, error)
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

func (c *questionSearcherProtobufClient) Search(ctx context.Context, in *SearchRequest) (*SearchResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "wordsearcher")
	ctx = ctxsetters.WithServiceName(ctx, "QuestionSearcher")
	ctx = ctxsetters.WithMethodName(ctx, "Search")
	out := new(SearchResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionSearcherProtobufClient) Expand(ctx context.Context, in *SearchResponse) (*SearchResponse, error) {
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

func (c *questionSearcherJSONClient) Search(ctx context.Context, in *SearchRequest) (*SearchResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "wordsearcher")
	ctx = ctxsetters.WithServiceName(ctx, "QuestionSearcher")
	ctx = ctxsetters.WithMethodName(ctx, "Search")
	out := new(SearchResponse)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionSearcherJSONClient) Expand(ctx context.Context, in *SearchResponse) (*SearchResponse, error) {
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
	var respContent *SearchResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.QuestionSearcher.Search(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SearchResponse and nil error while calling Search. nil responses are not supported"))
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
	var respContent *SearchResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.QuestionSearcher.Search(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SearchResponse and nil error while calling Search. nil responses are not supported"))
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

	reqContent := new(SearchResponse)
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
	reqContent := new(SearchResponse)
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
	// 866 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x35, 0x2d, 0x93, 0x16, 0x87, 0x96, 0xbc, 0x1e, 0xa4, 0x29, 0xe1, 0xa6, 0x8d, 0xa0, 0x20,
	0x8d, 0x0e, 0x85, 0x03, 0xb8, 0x87, 0x5e, 0x72, 0xa1, 0x14, 0x46, 0x22, 0x42, 0x91, 0xee, 0x52,
	0x96, 0xd3, 0x13, 0x41, 0x49, 0x8c, 0x4d, 0x44, 0x5a, 0xaa, 0x94, 0xdc, 0xca, 0xbf, 0xa3, 0xe7,
	0xfe, 0xaa, 0xa2, 0xfd, 0x3b, 0x2d, 0x76, 0x97, 0x14, 0xa9, 0xc0, 0x50, 0x73, 0xdb, 0x79, 0xfb,
	0xe6, 0xcd, 0xc7, 0x23, 0x96, 0xf0, 0x3c, 0x5b, 0x4e, 0x5f, 0xff, 0x9e, 0x66, 0xb3, 0x55, 0x1c,
	0x65, 0xd3, 0xbb, 0x38, 0x7b, 0x5d, 0x1c, 0x2e, 0x96, 0x59, 0xba, 0x4e, 0xf1, 0xa4, 0x7a, 0xd9,
	0xfe, 0x5b, 0x01, 0xdd, 0x9a, 0x2f, 0xef, 0xa2, 0xdb, 0x2c, 0x5a, 0xe0, 0x33, 0xd0, 0xa3, 0x22,
	0x30, 0x95, 0x96, 0xd2, 0xd1, 0x69, 0x09, 0x60, 0x07, 0x54, 0x91, 0x6b, 0x1e, 0xb6, 0x6a, 0x1d,
	0xe3, 0x12, 0x2f, 0xaa, 0x4a, 0x17, 0x37, 0x69, 0x36, 0xa3, 0x92, 0x80, 0x6d, 0x38, 0x89, 0x37,
	0xcb, 0x88, 0xcd, 0xe2, 0x19, 0x8d, 0x97, 0x99, 0x59, 0x6b, 0x29, 0x9d, 0x3a, 0xdd, 0xc1, 0xf0,
	0x29, 0x68, 0xf3, 0x98, 0xdd, 0xae, 0xef, 0xcc, 0xa3, 0x96, 0xd2, 0x51, 0x69, 0x1e, 0x61, 0x0b,
	0x8c, 0x65, 0x96, 0x4e, 0xa2, 0x49, 0x32, 0x4f, 0xd6, 0x0f, 0xa6, 0x2a, 0x2e, 0xab, 0x10, 0x57,
	0x9f, 0xa6, 0x8b, 0x49, 0xc2, 0xa2, 0x75, 0x92, 0xb2, 0x95, 0xa9, 0xb5, 0x94, 0x4e, 0x8d, 0xee,
	0x60, 0xed, 0x3f, 0x0e, 0xe1, 0x88, 0x77, 0x84, 0x08, 0x47, 0xbc, 0xa7, 0x7c, 0x1a, 0x71, 0xde,
	0x1d, 0xf3, 0xf0, 0xf3, 0x31, 0xbf, 0x03, 0x98, 0xc5, 0x1f, 0x13, 0x96, 0x70, 0x25, 0xd1, 0xba,
	0x4e, 0x2b, 0x08, 0x3e, 0x07, 0xe3, 0x63, 0x96, 0xb2, 0x75, 0x78, 0x97, 0xa6, 0x9f, 0x56, 0xa2,
	0x7b, 0x9d, 0x82, 0x80, 0x06, 0x1c, 0xc1, 0x6f, 0x01, 0x26, 0xd1, 0xf4, 0x53, 0x7e, 0xaf, 0x4a,
	0x7d, 0x8e, 0xc8, 0xeb, 0x57, 0x70, 0x3a, 0x8f, 0x37, 0xc9, 0x34, 0x65, 0xe1, 0xea, 0x61, 0x31,
	0x49, 0xe7, 0x72, 0x02, 0x9d, 0x36, 0x73, 0x38, 0x90, 0x28, 0x76, 0x80, 0x24, 0x8c, 0xc5, 0x59,
	0x58, 0x96, 0x33, 0x8f, 0xc5, 0x26, 0x9b, 0x02, 0x7f, 0x57, 0x94, 0xc4, 0xef, 0xe1, 0x54, 0x32,
	0xb7, 0x75, 0xcd, 0xba, 0x20, 0x36, 0x04, 0xdc, 0xcd, 0x6b, 0xb7, 0xff, 0xaa, 0x43, 0x23, 0x10,
	0x86, 0xd1, 0xf8, 0xd7, 0xfb, 0x78, 0xb5, 0xc6, 0xf7, 0x70, 0x22, 0x1d, 0x5c, 0x46, 0x59, 0xb4,
	0x58, 0x99, 0x8a, 0xb0, 0xf6, 0xd5, 0xae, 0xb5, 0x3b, 0x29, 0x79, 0x74, 0xc5, 0xf9, 0x74, 0x27,
	0x99, 0x5b, 0x2a, 0x2d, 0x16, 0x4b, 0xad, 0xd3, 0x3c, 0x3a, 0xff, 0x01, 0xb4, 0x61, 0xc2, 0x86,
	0xd1, 0x06, 0x09, 0xd4, 0x16, 0x09, 0x13, 0x66, 0xa8, 0x94, 0x1f, 0x05, 0x12, 0x6d, 0x44, 0x02,
	0x47, 0xa2, 0xcd, 0xf9, 0x0b, 0x30, 0x82, 0x75, 0x96, 0xb0, 0xdb, 0x71, 0x34, 0xbf, 0x8f, 0xf1,
	0x09, 0xa8, 0xbf, 0xf1, 0x43, 0xee, 0xa0, 0x0c, 0xce, 0x5f, 0x16, 0x24, 0x2b, 0xcb, 0xa2, 0x07,
	0x5e, 0x59, 0xe0, 0x72, 0x00, 0x9d, 0xe6, 0x11, 0xa7, 0x79, 0xf7, 0x8b, 0x49, 0x9c, 0x3d, 0x46,
	0x53, 0xb7, 0xb4, 0x17, 0x05, 0xed, 0x91, 0x92, 0x6a, 0x51, 0xf2, 0x9f, 0x1a, 0x18, 0x95, 0xd9,
	0xb1, 0x07, 0xfa, 0x34, 0x65, 0x33, 0xf9, 0x99, 0x70, 0x66, 0xf3, 0xf2, 0xe5, 0xbe, 0xbd, 0xf5,
	0x0a, 0x32, 0x2d, 0xf3, 0xf0, 0x0d, 0x68, 0x8b, 0x84, 0x15, 0x1b, 0x30, 0x2e, 0xdb, 0xfb, 0x14,
	0xe4, 0x12, 0x07, 0x07, 0x34, 0xcf, 0xc1, 0xf7, 0x60, 0xac, 0xc4, 0x16, 0x64, 0xbb, 0x35, 0x21,
	0xb1, 0xdf, 0xbc, 0x72, 0xb3, 0x83, 0x03, 0x5a, 0xcd, 0x2e, 0xc5, 0x22, 0xbe, 0x2b, 0xf1, 0x5d,
	0x7f, 0x91, 0x98, 0x58, 0x6d, 0x29, 0x26, 0xb2, 0xb9, 0x18, 0x13, 0x1b, 0x95, 0x62, 0xea, 0xff,
	0x8b, 0x55, 0x7c, 0xe2, 0x62, 0x95, 0xec, 0x52, 0x4c, 0x8e, 0xa9, 0x7d, 0xa9, 0xd8, 0x76, 0xcc,
	0x4a, 0x76, 0x97, 0x40, 0x73, 0xbb, 0x7e, 0xf1, 0xdd, 0xb6, 0xff, 0x55, 0x40, 0xdf, 0x9a, 0x83,
	0x06, 0x1c, 0xbb, 0xf6, 0x07, 0xa7, 0xe7, 0x7b, 0xe4, 0x00, 0x01, 0x34, 0xd7, 0xf6, 0xfa, 0xa3,
	0x01, 0x51, 0xf0, 0x2b, 0x38, 0xbb, 0xa2, 0x7e, 0xd7, 0xea, 0x3a, 0xae, 0x33, 0xfa, 0x25, 0xa4,
	0x96, 0xd7, 0xb7, 0xc9, 0x21, 0x3e, 0x01, 0x52, 0x85, 0x5d, 0x27, 0x18, 0x91, 0xda, 0xe7, 0x64,
	0xd7, 0x19, 0x3a, 0x23, 0x72, 0x84, 0x4f, 0x01, 0xbd, 0xeb, 0x61, 0xd7, 0xa6, 0xa1, 0xff, 0x2e,
	0xb4, 0x3c, 0xab, 0x4f, 0xad, 0x61, 0x40, 0x54, 0x2e, 0x52, 0xe2, 0x63, 0xff, 0xc6, 0x76, 0x03,
	0xa2, 0xe1, 0x29, 0x18, 0x57, 0xbe, 0xe3, 0x8d, 0xc2, 0xb1, 0xe5, 0x5e, 0xdb, 0xa4, 0xce, 0x69,
	0x43, 0x6b, 0xd4, 0x1b, 0x38, 0x5e, 0xbf, 0xc8, 0x26, 0x3a, 0x22, 0x34, 0x2d, 0xf7, 0x6a, 0x20,
	0x42, 0x59, 0x1f, 0x38, 0xe6, 0xf9, 0xa3, 0xd0, 0xf1, 0xc2, 0x62, 0x18, 0x03, 0x75, 0x50, 0xa9,
	0xdd, 0xb7, 0x3f, 0x90, 0x13, 0x6c, 0x80, 0x7e, 0xe3, 0xd3, 0xb7, 0x92, 0xdd, 0x68, 0xbf, 0x81,
	0x33, 0x2f, 0x5d, 0x3b, 0xcc, 0x8d, 0x37, 0xe5, 0x22, 0xce, 0xa0, 0xe1, 0x8f, 0x06, 0x36, 0x0d,
	0x6d, 0xaf, 0xef, 0x3a, 0xc1, 0x80, 0x1c, 0xc8, 0x59, 0xed, 0xb1, 0xe3, 0x5f, 0x07, 0xe1, 0xd8,
	0xa6, 0x81, 0xe3, 0x7b, 0x44, 0x69, 0x4f, 0xa1, 0x59, 0x6c, 0x7f, 0xb5, 0x4c, 0xd9, 0x2a, 0xc6,
	0x9f, 0x00, 0xb6, 0xef, 0x69, 0xf1, 0xa6, 0x7c, 0xbd, 0xeb, 0xd7, 0xf6, 0xa7, 0x43, 0x2b, 0x54,
	0x34, 0xe1, 0x38, 0x7f, 0x04, 0xf3, 0x77, 0xb9, 0x08, 0x2f, 0xff, 0x54, 0x80, 0xfc, 0xcc, 0xbd,
	0x4d, 0x52, 0x16, 0xe4, 0x22, 0xd8, 0x03, 0x4d, 0x9e, 0xf1, 0x9b, 0x3d, 0x5f, 0xc3, 0xf9, 0xb3,
	0xc7, 0x2f, 0xf3, 0x66, 0xdf, 0x82, 0x66, 0x8b, 0x77, 0x0a, 0xf7, 0xf2, 0xf6, 0xab, 0x4c, 0x34,
	0xf1, 0x77, 0xfd, 0xf1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xc4, 0x54, 0xd1, 0x80, 0x07,
	0x00, 0x00,
}
