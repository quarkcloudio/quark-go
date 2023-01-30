package zeroadapter

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const JSON_RESPONSE = "json"   // json类型响应
const IMAGE_RESPONSE = "image" // 图片类型响应
const EXCEL_RESPONSE = "excel" // Excel文件类型响应

// 将gozero框架的Ctx转换为builder框架Request
func RequestAdapter(r *http.Request, routePath string) (*builder.Request, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	headerString := ""
	for hk, hvs := range r.Header {
		tmp := ""
		for _, v := range hvs {
			tmp = tmp + "," + v
		}
		tmp = strings.Trim(tmp, ",")
		headerString = headerString + hk + ": " + tmp + "\r\n"
	}

	// 将框架请求转换为builder框架请求
	return &builder.Request{
		IPString:       r.RemoteAddr,
		HeaderString:   headerString,
		MethodString:   r.Method,
		FullPathString: routePath,
		HostString:     r.Host,
		PathString:     r.URL.Path,
		QueryString:    r.Response.Request.URL.RawQuery,
		BodyBuffer:     body,
	}, nil
}

// 适配gozero框架响应
func ResponseAdapter(r *builder.Resource, responseType string, w http.ResponseWriter) {
	result, err := r.Run()
	if err != nil {
		httpx.OkJson(w, msg.Error(err.Error(), ""))
		return
	}

	// 响应结果
	switch responseType {
	case JSON_RESPONSE:
		httpx.OkJson(w, result)
		return
	case IMAGE_RESPONSE:
		w.Write(result.([]byte))
		return
	case EXCEL_RESPONSE:
		w.Header().Add("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Write(result.([]byte))
		return
	}
}

// 适配gozero框架路由
func RouteAdapter(b *builder.Resource, routePath string, responseType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			httpx.OkJson(w, msg.Error(err.Error(), ""))
			return
		}

		headerString := ""
		for hk, hvs := range r.Header {
			tmp := ""
			for _, v := range hvs {
				tmp = tmp + "," + v
			}
			tmp = strings.Trim(tmp, ",")
			headerString = headerString + hk + ": " + tmp + "\r\n"
		}

		// 将框架请求转换为builder框架请求
		request := &builder.Request{
			IPString:       r.RemoteAddr,
			HeaderString:   headerString,
			MethodString:   r.Method,
			FullPathString: routePath,
			HostString:     r.Host,
			PathString:     r.URL.Path,
			QueryString:    r.URL.RawQuery,
			BodyBuffer:     body,
		}

		// 转换Request对象
		result, err := b.TransformRequest(request).Run()
		if err != nil {
			httpx.OkJson(w, msg.Error(err.Error(), ""))
			return
		}

		// 响应结果
		switch responseType {
		case JSON_RESPONSE:
			httpx.OkJson(w, result)
			return
		case IMAGE_RESPONSE:

			w.Write(result.([]byte))
			return
		case EXCEL_RESPONSE:
			w.Header().Add("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
			w.Header().Add("Content-Type", "application/octet-stream")
			w.Write(result.([]byte))
			return
		}
	}
}

// 适配gozero框架
func Adapter(b *builder.Resource, server *rest.Server) {

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/login/:resource/index",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/index", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/login/:resource/handle",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/handle", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/login/:resource/captchaId",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/captchaId", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/login/:resource/captcha/:id",
				Handler: RouteAdapter(b, "/api/admin/login/:resource/captcha/:id", IMAGE_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/logout/:resource/handle",
				Handler: RouteAdapter(b, "/api/admin/logout/:resource/handle", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/dashboard/:resource/index",
				Handler: RouteAdapter(b, "/api/admin/dashboard/:resource/index", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/index",
				Handler: RouteAdapter(b, "/api/admin/:resource/index", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/editable",
				Handler: RouteAdapter(b, "/api/admin/:resource/editable", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/action/:uriKey",
				Handler: RouteAdapter(b, "/api/admin/:resource/action/:uriKey", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/action/:uriKey",
				Handler: RouteAdapter(b, "/api/admin/:resource/action/:uriKey", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/create",
				Handler: RouteAdapter(b, "/api/admin/:resource/create", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/store",
				Handler: RouteAdapter(b, "/api/admin/:resource/store", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/edit",
				Handler: RouteAdapter(b, "/api/admin/:resource/edit", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/edit/values",
				Handler: RouteAdapter(b, "/api/admin/:resource/edit/values", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/save",
				Handler: RouteAdapter(b, "/api/admin/:resource/save", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/detail",
				Handler: RouteAdapter(b, "/api/admin/:resource/detail", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/export",
				Handler: RouteAdapter(b, "/api/admin/:resource/export", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/import",
				Handler: RouteAdapter(b, "/api/admin/:resource/import", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/:resource/import",
				Handler: RouteAdapter(b, "/api/admin/:resource/import", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/import/template",
				Handler: RouteAdapter(b, "/api/admin/:resource/import/template", EXCEL_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/:resource/:uriKey/form",
				Handler: RouteAdapter(b, "/api/admin/:resource/:uriKey/form", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/upload/:resource/handle",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/handle", JSON_RESPONSE),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/admin/upload/:resource/getList",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/getList", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/upload/:resource/delete",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/delete", JSON_RESPONSE),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/admin/upload/:resource/crop",
				Handler: RouteAdapter(b, "/api/admin/upload/:resource/crop", JSON_RESPONSE),
			},
		},
	)
}
