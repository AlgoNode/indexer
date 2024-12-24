// Package common provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package common

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns 200 if healthy.
	// (GET /health)
	MakeHealthCheck(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// MakeHealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) MakeHealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MakeHealthCheck(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/health", wrapper.MakeHealthCheck, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
