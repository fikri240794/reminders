package corses

func (m *CORSMiddleware) Routing() {
	m.httpServer.Use(m.CORS)
}
