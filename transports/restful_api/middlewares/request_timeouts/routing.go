package request_timeouts

func (m *RequestTimeoutMiddleware) Routing() {
	m.httpServer.Use(m.RequestTimeout)
}
