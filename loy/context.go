package loy

import (
	"bytes"
	"mime/multipart"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type Context struct {
	Ctx *fiber.Ctx
}

func (c *Context) SendString(s string) error {
	return c.Ctx.SendString(s)
}

func (c *Context) BodyParser(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *Context) QueryParser(v interface{}) error {
	return c.Ctx.QueryParser(v)
}

func (c *Context) Params(key string) string {
	return c.Ctx.Params(key)
}

func (c *Context) Status(status int) *Context {
	c.Ctx.Status(status)
	return c
}

func (c *Context) SendStatus(status int) error {
	return c.Ctx.SendStatus(status)
}

func (c *Context) JSON(v interface{}) error {
	return c.Ctx.JSON(v)
}

func (c *Context) SetHeader(key, value string) {
	c.Ctx.Set(key, value)
}

func (c *Context) GetHeader(key string) string {
	return c.Ctx.Get(key)
}

func (c *Context) Next() error {
	return c.Ctx.Next()
}

func (c *Context) GetCookie(key string) string {
	return c.Ctx.Cookies(key)
}

func (c *Context) SetCookie(cookie *fiber.Cookie) {
	c.Ctx.Cookie(cookie)
}

func (c *Context) ClearCookie(key string) {
	c.Ctx.ClearCookie(key)
}

func (c *Context) Render(component templ.Component) error {
	var buff bytes.Buffer
	w := &BytesWriter{&buff}
	err := component.Render(c.Ctx.Context(), w)
	if err != nil {
		return err
	}
	c.SetHeader("Content-Type", "text/html")
	return c.Ctx.Send(buff.Bytes())
}

func (c *Context) FormFile(key string) (*multipart.FileHeader, error) {
	return c.Ctx.FormFile(key)
}
