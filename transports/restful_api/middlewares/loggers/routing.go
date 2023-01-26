package loggers

func (m *LoggerMiddleware) Routing() {
	m.httpServer.Use(m.Log)
}
