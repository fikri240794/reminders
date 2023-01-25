package logger

func (m *LoggerMiddleware) Routing() {
	m.httpServer.Use(m.Log)
}
