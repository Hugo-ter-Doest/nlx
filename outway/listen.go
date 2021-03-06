// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"encoding/binary"
	"hash/crc64"
	"net/http"
	"strconv"
	"strings"

	"go.uber.org/zap/zapcore"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) ListenAndServe(address string) error {
	err := http.ListenAndServe(address, o)
	if err != nil {
		return errors.Wrap(err, "failed to run http server")
	}
	return nil
}

// ServeHTTP handles requests from the organization to the outway, it selects the correct service backend and lets it handle the request further.
func (o *Outway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := o.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)
	logger.Debug("received request")
	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 3)
	if len(urlparts) != 3 {
		http.Error(w, "nlx outway error: invalid path in url", http.StatusBadRequest)
		logger.Warn("received request with invalid path")
		return
	}
	organizationName := urlparts[0]
	serviceName := urlparts[1]
	r.URL.Path = urlparts[2] // retain original path

	o.servicesLock.RLock()
	service := o.services[organizationName+"."+serviceName]
	o.servicesLock.RUnlock()
	if service == nil {
		http.Error(w, "nlx outway error: unknown service", http.StatusBadRequest)
		logger.Warn("received request for unknown service")
		return
	}

	var logFields = []zapcore.Field{zap.String("doelbinding-log", "yes")}
	if userID := r.Header.Get("X-NLX-Request-User-Id"); userID != "" {
		logFields = append(logFields, zap.String("doelbinding-user-id", userID))
		r.Header.Del("X-NLX-Request-User-Id")
	}
	if applicationID := r.Header.Get("X-NLX-Request-Application-Id"); applicationID != "" {
		logFields = append(logFields, zap.String("doelbinding-application-id", applicationID))
		r.Header.Del("X-NLX-Request-Application-Id")
	}
	if subjectIdentifier := r.Header.Get("X-NLX-Request-Subject-Identifier"); subjectIdentifier != "" {
		logFields = append(logFields, zap.String("doelbinding-subject-identifier", subjectIdentifier))
		r.Header.Del("X-NLX-Request-Subject-Identifier")
	}
	if processID := r.Header.Get("X-NLX-Request-Process-Id"); processID != "" {
		logFields = append(logFields, zap.String("doelbinding-process-id", processID))
	}
	if dataElements := r.Header.Get("X-NLX-Request-Data-Elements"); dataElements != "" {
		logFields = append(logFields, zap.String("doelbinding-data-elements", dataElements))
	}

	requestIDFlake, err := o.requestFlake.NextID()
	if err != nil {
		logger.Error("could not get new request ID", zap.Error(err))
		http.Error(w, "outway: internal server error", http.StatusInternalServerError)
		return
	}
	requestIDFlakeBytes := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(requestIDFlakeBytes, requestIDFlake)
	requestIDNum := crc64.Checksum(requestIDFlakeBytes, o.ecmaTable)
	requestID := strconv.FormatUint(requestIDNum, 32)
	logFields = append(logFields, zap.String("request-id", requestID))
	r.Header.Set("X-NLX-Request-Id", requestID)

	logger.Info("sending request", logFields...)

	service.proxyRequest(w, r)

	logger.Info("sending request finished", logFields...)
}
