// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package logger

const (

	// HTTP
	EcsHttpVersion           = "http.version"
	EcsHttpRequestId         = "http.request.id"
	EcsHttpRequestMethod     = "http.request.method"
	EcsHttpRequestBodyBytes  = "http.request.body.bytes"
	EcsHttpResponseCode      = "http.response.status_code"
	EcsHttpResponseBodyBytes = "http.response.body.bytes"

	// URL
	EcsUrlFull   = "url.full"
	EcsUrlDomain = "url.domain"
	EcsUrlPort   = "url.port"

	// Client
	EcsClientAddress = "client.address"
	EcsClientIp      = "client.ip"
	EcsClientPort    = "client.port"

	// TLS
	EcsTlsEstablished = "tls.established"

	// Event
	EcsEventDuration = "event.duration"
)
