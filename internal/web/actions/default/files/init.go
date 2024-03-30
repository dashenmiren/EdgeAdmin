package files

import "github.com/iwind/TeaGo"

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Prefix("/files").
			Get("/file", new(FileAction)).
			EndAll()
	})
}
