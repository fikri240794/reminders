package swagger

func (c *SwaggerController) Routing() {
	c.httpServer.Get("/swagger/*", c.Index)
}
