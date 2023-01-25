package etags

func (m *ETagMiddleware) Routing() {
	m.httpServer.Use(m.ETag)
}
