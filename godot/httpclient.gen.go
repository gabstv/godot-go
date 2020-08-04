package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

// HTTPClientMethod is an enum for Method values.
type HTTPClientMethod int

const (
	HTTPClientMethodConnect HTTPClientMethod = 7
	HTTPClientMethodDelete  HTTPClientMethod = 4
	HTTPClientMethodGet     HTTPClientMethod = 0
	HTTPClientMethodHead    HTTPClientMethod = 1
	HTTPClientMethodMax     HTTPClientMethod = 9
	HTTPClientMethodOptions HTTPClientMethod = 5
	HTTPClientMethodPatch   HTTPClientMethod = 8
	HTTPClientMethodPost    HTTPClientMethod = 2
	HTTPClientMethodPut     HTTPClientMethod = 3
	HTTPClientMethodTrace   HTTPClientMethod = 6
)

// HTTPClientResponseCode is an enum for ResponseCode values.
type HTTPClientResponseCode int

const (
	HTTPClientResponseAccepted                     HTTPClientResponseCode = 202
	HTTPClientResponseAlreadyReported              HTTPClientResponseCode = 208
	HTTPClientResponseBadGateway                   HTTPClientResponseCode = 502
	HTTPClientResponseBadRequest                   HTTPClientResponseCode = 400
	HTTPClientResponseConflict                     HTTPClientResponseCode = 409
	HTTPClientResponseContinue                     HTTPClientResponseCode = 100
	HTTPClientResponseCreated                      HTTPClientResponseCode = 201
	HTTPClientResponseExpectationFailed            HTTPClientResponseCode = 417
	HTTPClientResponseFailedDependency             HTTPClientResponseCode = 424
	HTTPClientResponseForbidden                    HTTPClientResponseCode = 403
	HTTPClientResponseFound                        HTTPClientResponseCode = 302
	HTTPClientResponseGatewayTimeout               HTTPClientResponseCode = 504
	HTTPClientResponseGone                         HTTPClientResponseCode = 410
	HTTPClientResponseHttpVersionNotSupported      HTTPClientResponseCode = 505
	HTTPClientResponseImATeapot                    HTTPClientResponseCode = 418
	HTTPClientResponseImUsed                       HTTPClientResponseCode = 226
	HTTPClientResponseInsufficientStorage          HTTPClientResponseCode = 507
	HTTPClientResponseInternalServerError          HTTPClientResponseCode = 500
	HTTPClientResponseLengthRequired               HTTPClientResponseCode = 411
	HTTPClientResponseLocked                       HTTPClientResponseCode = 423
	HTTPClientResponseLoopDetected                 HTTPClientResponseCode = 508
	HTTPClientResponseMethodNotAllowed             HTTPClientResponseCode = 405
	HTTPClientResponseMisdirectedRequest           HTTPClientResponseCode = 421
	HTTPClientResponseMovedPermanently             HTTPClientResponseCode = 301
	HTTPClientResponseMultipleChoices              HTTPClientResponseCode = 300
	HTTPClientResponseMultiStatus                  HTTPClientResponseCode = 207
	HTTPClientResponseNetworkAuthRequired          HTTPClientResponseCode = 511
	HTTPClientResponseNonAuthoritativeInformation  HTTPClientResponseCode = 203
	HTTPClientResponseNotAcceptable                HTTPClientResponseCode = 406
	HTTPClientResponseNotExtended                  HTTPClientResponseCode = 510
	HTTPClientResponseNotFound                     HTTPClientResponseCode = 404
	HTTPClientResponseNotImplemented               HTTPClientResponseCode = 501
	HTTPClientResponseNotModified                  HTTPClientResponseCode = 304
	HTTPClientResponseNoContent                    HTTPClientResponseCode = 204
	HTTPClientResponseOk                           HTTPClientResponseCode = 200
	HTTPClientResponsePartialContent               HTTPClientResponseCode = 206
	HTTPClientResponsePaymentRequired              HTTPClientResponseCode = 402
	HTTPClientResponsePermanentRedirect            HTTPClientResponseCode = 308
	HTTPClientResponsePreconditionFailed           HTTPClientResponseCode = 412
	HTTPClientResponsePreconditionRequired         HTTPClientResponseCode = 428
	HTTPClientResponseProcessing                   HTTPClientResponseCode = 102
	HTTPClientResponseProxyAuthenticationRequired  HTTPClientResponseCode = 407
	HTTPClientResponseRequestedRangeNotSatisfiable HTTPClientResponseCode = 416
	HTTPClientResponseRequestEntityTooLarge        HTTPClientResponseCode = 413
	HTTPClientResponseRequestHeaderFieldsTooLarge  HTTPClientResponseCode = 431
	HTTPClientResponseRequestTimeout               HTTPClientResponseCode = 408
	HTTPClientResponseRequestUriTooLong            HTTPClientResponseCode = 414
	HTTPClientResponseResetContent                 HTTPClientResponseCode = 205
	HTTPClientResponseSeeOther                     HTTPClientResponseCode = 303
	HTTPClientResponseServiceUnavailable           HTTPClientResponseCode = 503
	HTTPClientResponseSwitchingProtocols           HTTPClientResponseCode = 101
	HTTPClientResponseSwitchProxy                  HTTPClientResponseCode = 306
	HTTPClientResponseTemporaryRedirect            HTTPClientResponseCode = 307
	HTTPClientResponseTooManyRequests              HTTPClientResponseCode = 429
	HTTPClientResponseUnauthorized                 HTTPClientResponseCode = 401
	HTTPClientResponseUnavailableForLegalReasons   HTTPClientResponseCode = 451
	HTTPClientResponseUnprocessableEntity          HTTPClientResponseCode = 422
	HTTPClientResponseUnsupportedMediaType         HTTPClientResponseCode = 415
	HTTPClientResponseUpgradeRequired              HTTPClientResponseCode = 426
	HTTPClientResponseUseProxy                     HTTPClientResponseCode = 305
	HTTPClientResponseVariantAlsoNegotiates        HTTPClientResponseCode = 506
)

// HTTPClientStatus is an enum for Status values.
type HTTPClientStatus int

const (
	HTTPClientStatusBody              HTTPClientStatus = 7
	HTTPClientStatusCantConnect       HTTPClientStatus = 4
	HTTPClientStatusCantResolve       HTTPClientStatus = 2
	HTTPClientStatusConnected         HTTPClientStatus = 5
	HTTPClientStatusConnecting        HTTPClientStatus = 3
	HTTPClientStatusConnectionError   HTTPClientStatus = 8
	HTTPClientStatusDisconnected      HTTPClientStatus = 0
	HTTPClientStatusRequesting        HTTPClientStatus = 6
	HTTPClientStatusResolving         HTTPClientStatus = 1
	HTTPClientStatusSslHandshakeError HTTPClientStatus = 9
)

//func NewHTTPClientFromPointer(ptr gdnative.Pointer) HTTPClient {
func newHTTPClientFromPointer(ptr gdnative.Pointer) HTTPClient {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HTTPClient{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Hyper-text transfer protocol client (sometimes called "User Agent"). Used to make HTTP requests to download web content, upload files and other data or to communicate with various services, among other use cases. See [HTTPRequest] for an higher-level alternative. [b]Note:[/b] This client only needs to connect to a host once (see [method connect_to_host]) to send multiple requests. Because of this, methods that take URLs usually take just the part after the host instead of the full URL, as the client is already connected to a host. See [method request] for a full example and to get started. A [HTTPClient] should be reused between multiple requests or to connect to different hosts instead of creating one client per request. Supports SSL and SSL server certificate verification. HTTP status codes in the 2xx range indicate success, 3xx redirection (i.e. "try again, but over here"), 4xx something was wrong with the request, and 5xx something went wrong on the server's side. For more information on HTTP, see https://developer.mozilla.org/en-US/docs/Web/HTTP (or read RFC 2616 to get it straight from the source: https://tools.ietf.org/html/rfc2616). [b]Note:[/b] When performing HTTP requests from a project exported to HTML5, keep in mind the remote server may not allow requests from foreign origins due to [url=https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS]CORS[/url]. If you host the server in question, you should modify its backend to allow requests from foreign origins by adding the [code]Access-Control-Allow-Origin: *[/code] HTTP header.
*/
type HTTPClient struct {
	Reference
	owner gdnative.Object
}

func (o *HTTPClient) BaseClass() string {
	return "HTTPClient"
}

/*
        Closes the current connection, allowing reuse of this [HTTPClient].
	Args: [], Returns: void
*/
func (o *HTTPClient) Close() {
	//log.Println("Calling HTTPClient.Close()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "close")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Connects to a host. This needs to be done before any requests are sent. The host should not have http:// prepended but will strip the protocol identifier if provided. If no [code]port[/code] is specified (or [code]-1[/code] is used), it is automatically set to 80 for HTTP and 443 for HTTPS (if [code]use_ssl[/code] is enabled). [code]verify_host[/code] will check the SSL identity of the host if set to [code]true[/code].
	Args: [{ false host String} {-1 true port int} {False true use_ssl bool} {True true verify_host bool}], Returns: enum.Error
*/
func (o *HTTPClient) ConnectToHost(host gdnative.String, port gdnative.Int, useSsl gdnative.Bool, verifyHost gdnative.Bool) gdnative.Error {
	//log.Println("Calling HTTPClient.ConnectToHost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(host)
	ptrArguments[1] = gdnative.NewPointerFromInt(port)
	ptrArguments[2] = gdnative.NewPointerFromBool(useSsl)
	ptrArguments[3] = gdnative.NewPointerFromBool(verifyHost)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "connect_to_host")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [], Returns: StreamPeer
*/
func (o *HTTPClient) GetConnection() StreamPeerImplementer {
	//log.Println("Calling HTTPClient.GetConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_connection")

	// Call the parent method.
	// StreamPeer
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newStreamPeerFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(StreamPeerImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "StreamPeer" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(StreamPeerImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *HTTPClient) GetReadChunkSize() gdnative.Int {
	//log.Println("Calling HTTPClient.GetReadChunkSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_read_chunk_size")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the response's body length. [b]Note:[/b] Some Web servers may not send a body length. In this case, the value returned will be [code]-1[/code]. If using chunked transfer encoding, the body length will also be [code]-1[/code].
	Args: [], Returns: int
*/
func (o *HTTPClient) GetResponseBodyLength() gdnative.Int {
	//log.Println("Calling HTTPClient.GetResponseBodyLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_body_length")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the response's HTTP status code.
	Args: [], Returns: int
*/
func (o *HTTPClient) GetResponseCode() gdnative.Int {
	//log.Println("Calling HTTPClient.GetResponseCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_code")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the response headers.
	Args: [], Returns: PoolStringArray
*/
func (o *HTTPClient) GetResponseHeaders() gdnative.PoolStringArray {
	//log.Println("Calling HTTPClient.GetResponseHeaders()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_headers")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Returns all response headers as a Dictionary of structure [code]{ "key": "value1; value2" }[/code] where the case-sensitivity of the keys and values is kept like the server delivers it. A value is a simple String, this string can have more than one value where "; " is used as separator. [b]Example:[/b] [codeblock] { "content-length": 12, "Content-Type": "application/json; charset=UTF-8", } [/codeblock]
	Args: [], Returns: Dictionary
*/
func (o *HTTPClient) GetResponseHeadersAsDictionary() gdnative.Dictionary {
	//log.Println("Calling HTTPClient.GetResponseHeadersAsDictionary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_response_headers_as_dictionary")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Returns a [enum Status] constant. Need to call [method poll] in order to get status updates.
	Args: [], Returns: enum.HTTPClient::Status
*/
func (o *HTTPClient) GetStatus() HTTPClientStatus {
	//log.Println("Calling HTTPClient.GetStatus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "get_status")

	// Call the parent method.
	// enum.HTTPClient::Status
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return HTTPClientStatus(ret)
}

/*
        If [code]true[/code], this [HTTPClient] has a response available.
	Args: [], Returns: bool
*/
func (o *HTTPClient) HasResponse() gdnative.Bool {
	//log.Println("Calling HTTPClient.HasResponse()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "has_response")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *HTTPClient) IsBlockingModeEnabled() gdnative.Bool {
	//log.Println("Calling HTTPClient.IsBlockingModeEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "is_blocking_mode_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code], this [HTTPClient] has a response that is chunked.
	Args: [], Returns: bool
*/
func (o *HTTPClient) IsResponseChunked() gdnative.Bool {
	//log.Println("Calling HTTPClient.IsResponseChunked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "is_response_chunked")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        This needs to be called in order to have any request processed. Check results with [method get_status].
	Args: [], Returns: enum.Error
*/
func (o *HTTPClient) Poll() gdnative.Error {
	//log.Println("Calling HTTPClient.Poll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "poll")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Generates a GET/POST application/x-www-form-urlencoded style query string from a provided dictionary, e.g.: [codeblock] var fields = {"username": "user", "password": "pass"} var query_string = http_client.query_string_from_dict(fields) # Returns "username=user&password=pass" [/codeblock] Furthermore, if a key has a [code]null[/code] value, only the key itself is added, without equal sign and value. If the value is an array, for each value in it a pair with the same key is added. [codeblock] var fields = {"single": 123, "not_valued": null, "multiple": [22, 33, 44]} var query_string = http_client.query_string_from_dict(fields) # Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44" [/codeblock]
	Args: [{ false fields Dictionary}], Returns: String
*/
func (o *HTTPClient) QueryStringFromDict(fields gdnative.Dictionary) gdnative.String {
	//log.Println("Calling HTTPClient.QueryStringFromDict()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(fields)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "query_string_from_dict")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Reads one chunk from the response.
	Args: [], Returns: PoolByteArray
*/
func (o *HTTPClient) ReadResponseBodyChunk() gdnative.PoolByteArray {
	//log.Println("Calling HTTPClient.ReadResponseBodyChunk()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "read_response_body_chunk")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Sends a request to the connected host. The URL parameter is just the part after the host, so for [code]http://somehost.com/index.php[/code], it is [code]index.php[/code]. Headers are HTTP request headers. For available HTTP methods, see [enum Method]. To create a POST request with query strings to push to the server, do: [codeblock] var fields = {"username" : "user", "password" : "pass"} var query_string = http_client.query_string_from_dict(fields) var headers = ["Content-Type: application/x-www-form-urlencoded", "Content-Length: " + str(query_string.length())] var result = http_client.request(http_client.METHOD_POST, "index.php", headers, query_string) [/codeblock]
	Args: [{ false method int} { false url String} { false headers PoolStringArray} { true body String}], Returns: enum.Error
*/
func (o *HTTPClient) Request(method gdnative.Int, url gdnative.String, headers gdnative.PoolStringArray, body gdnative.String) gdnative.Error {
	//log.Println("Calling HTTPClient.Request()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(method)
	ptrArguments[1] = gdnative.NewPointerFromString(url)
	ptrArguments[2] = gdnative.NewPointerFromPoolStringArray(headers)
	ptrArguments[3] = gdnative.NewPointerFromString(body)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "request")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Sends a raw request to the connected host. The URL parameter is just the part after the host, so for [code]http://somehost.com/index.php[/code], it is [code]index.php[/code]. Headers are HTTP request headers. For available HTTP methods, see [enum Method]. Sends the body data raw, as a byte array and does not encode it in any way.
	Args: [{ false method int} { false url String} { false headers PoolStringArray} { false body PoolByteArray}], Returns: enum.Error
*/
func (o *HTTPClient) RequestRaw(method gdnative.Int, url gdnative.String, headers gdnative.PoolStringArray, body gdnative.PoolByteArray) gdnative.Error {
	//log.Println("Calling HTTPClient.RequestRaw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(method)
	ptrArguments[1] = gdnative.NewPointerFromString(url)
	ptrArguments[2] = gdnative.NewPointerFromPoolStringArray(headers)
	ptrArguments[3] = gdnative.NewPointerFromPoolByteArray(body)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "request_raw")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *HTTPClient) SetBlockingMode(enabled gdnative.Bool) {
	//log.Println("Calling HTTPClient.SetBlockingMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "set_blocking_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false connection StreamPeer}], Returns: void
*/
func (o *HTTPClient) SetConnection(connection StreamPeerImplementer) {
	//log.Println("Calling HTTPClient.SetConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(connection.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "set_connection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bytes int}], Returns: void
*/
func (o *HTTPClient) SetReadChunkSize(bytes gdnative.Int) {
	//log.Println("Calling HTTPClient.SetReadChunkSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bytes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPClient", "set_read_chunk_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// HTTPClientImplementer is an interface that implements the methods
// of the HTTPClient class.
type HTTPClientImplementer interface {
	ReferenceImplementer
	Close()
	GetConnection() StreamPeerImplementer
	GetReadChunkSize() gdnative.Int
	GetResponseBodyLength() gdnative.Int
	GetResponseCode() gdnative.Int
	GetResponseHeaders() gdnative.PoolStringArray
	GetResponseHeadersAsDictionary() gdnative.Dictionary
	HasResponse() gdnative.Bool
	IsBlockingModeEnabled() gdnative.Bool
	IsResponseChunked() gdnative.Bool
	QueryStringFromDict(fields gdnative.Dictionary) gdnative.String
	ReadResponseBodyChunk() gdnative.PoolByteArray
	SetBlockingMode(enabled gdnative.Bool)
	SetConnection(connection StreamPeerImplementer)
	SetReadChunkSize(bytes gdnative.Int)
}
