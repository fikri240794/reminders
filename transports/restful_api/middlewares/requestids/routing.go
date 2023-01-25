package requestids

func (m *RequestIDMiddleware) Routing() {
	m.httpServer.Use(m.RequestID)
}
