package compresses

func (m *CompressMiddleware) Routing() {
	m.httpServer.Use(m.Compress)
}
