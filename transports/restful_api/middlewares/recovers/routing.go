package recovers

func (m *RecoverMiddleware) Routing() {
	m.httpServer.Use(m.Recover)
}
