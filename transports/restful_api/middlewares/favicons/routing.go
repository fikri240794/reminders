package favicons

func (m *FaviconMiddleware) Routing() {
	m.httpServer.Use(m.Favicon)
}
