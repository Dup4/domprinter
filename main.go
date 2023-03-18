// Code generated by hertz generator.

package main

import (
	"bufio"
	"bytes"
	"context"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/requestid"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"

	"github.com/Dup4/domprinter/biz/constants"
	"github.com/Dup4/domprinter/biz/dal"
	_ "github.com/Dup4/domprinter/swagger"
)

// @BasePath /
// @schemes http
func main() {
	dal.Init()

	h := server.Default()

	h.
		Use(requestid.New()).
		Use(
			accesslog.New(
				accesslog.WithAccessLogFunc(hlog.CtxInfof),
				accesslog.WithTimeFormat(constants.ISO8601TimeFormat),
				accesslog.WithFormat(getAccessLogFormat()),
			),
		)

	initBasicAuth(h)
	register(h)

	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	h.Spin()
}

func initBasicAuth(h *server.Hertz) {
	authUsername := os.Getenv("AUTH_USERNAME")
	authPassword := os.Getenv("AUTH_PASSWORD")

	accesslog.Tags["requestID"] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(requestid.Get(c))
	}

	accesslog.Tags["clientIP"] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(c.ClientIP())
	}

	if len(authUsername) > 0 && len(authPassword) > 0 {
		h.Use(basic_auth.BasicAuth(map[string]string{
			authUsername: authPassword,
		}))
	}
}

func getAccessLogFormat() string {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	writer.WriteString("[${time}] [${requestID}] [${status}] [${latency}]")
	writer.WriteString(" -")
	writer.WriteString(" ${method} ${protocol}://${host}${path}")
	writer.WriteString(" [clientIP=${clientIP}] [path=${path}] [route=${route}] [url=${url}]")
	writer.WriteString(" [ip=${ip}] [ips=${ips}] [ua=${ua}]")
	writer.WriteString(" [bytesSent=${bytesSent}] [bytesReceived=${bytesReceived}]")
	writer.WriteString(" [queryParams=${queryParams}] [reqHeaders=${reqHeaders}] [resHeaders=${resHeaders}]")
	writer.WriteString(" [reqBody=${body}] [resBody=${resBody}]")

	writer.Flush()

	return buf.String()
}
