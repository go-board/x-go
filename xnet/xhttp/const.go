package xhttp

import (
	"strings"
)

type HeaderKey string

const (
	HeaderAccept                          HeaderKey = "Accept"
	HeaderAcceptCh                                  = "Accept-CH"
	HeaderAcceptChLifetime                          = "Accept-CH-Lifetime"
	HeaderAcceptCharset                             = "Accept-Charset"
	HeaderAcceptEncoding                            = "Accept-Encoding"
	HeaderAcceptLanguage                            = "Accept-Language"
	HeaderAcceptPatch                               = "Accept-Patch"
	HeaderAcceptRange                               = "Accept-Range"
	HeaderAccessControlAllowCredentials             = "Access-Control-Allow-Credentials"
	HeaderAccessControlAllowHeaders                 = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowMethods                 = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowOrigin                  = "Access-Control-Allow-Origin"
	HeaderAccessControlExposeHeaders                = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge                       = "Access-Control-Max-Age"
	HeaderAccessControlRequestHeaders               = "Access-Control-Request-Headers"
	HeaderAccessControlRequestMethod                = "Access-Control-Request-Method"
	HeaderAge                                       = "Age"
	HeaderAllow                                     = "Allow"
	HeaderAltSvc                                    = "Alt-Svc"
	HeaderAuthorization                             = "Authorization"
	HeaderCacheControl                              = "Cache-Control"
	HeaderClearSiteData                             = "Clear-Site-Data"
	HeaderConnection                                = "Connection"
	HeaderContentDisposition                        = "Content-Disposition"
	HeaderContentEncoding                           = "Content-Encoding"
	HeaderContentLanguage                           = "Content-Language"
	HeaderContentLength                             = "Content-Length"
	HeaderContentRange                              = "Content-Range"
	HeaderContentSecurityPolicy                     = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly           = "Content-Security-Policy-Report-Only"
	HeaderContentType                               = "Content-Type"
	HeaderCookie                                    = "Cookie"
	HeaderCrossOriginResourcePolicy                 = "Cross-Origin-Resource-Policy"
	HeaderDNT                                       = "DNT"
	HeaderDPR                                       = "DPR"
	HeaderDate                                      = "Date"
	HeaderDeviceMemory                              = "Device-Memory"
	HeaderDigest                                    = "Digest"
	HeaderETag                                      = "ETag"
	HeaderEarlyData                                 = "Early-Data"
	HeaderExpect                                    = "Expect"
	HeaderExpectCT                                  = "Expect-CT"
	HeaderExpires                                   = "Expires"
	HeaderFeaturePolicy                             = "Feature-Policy" // exp
	HeaderForwarded                                 = "Forwarded"
	HeaderFrom                                      = "From"
	HeaderHost                                      = "Host"
	HeaderIfMatch                                   = "If-Match"
	HeaderIfModifiedSince                           = "If-Modified-Since"
	HeaderIfNoneMatch                               = "If-None-Match"
	HeaderIfRange                                   = "If-Range"
	HeaderIfUnmodifiedSince                         = "If-Unmodified-Since"
	HeaderIndex                                     = "Index"
	HeaderKeepAlive                                 = "Keep-Alive"
	HeaderLargeAllocation                           = "Large-Allocation"
	HeaderLastModified                              = "Last-Modified"
	HeaderLink                                      = "Link"
	HeaderLocation                                  = "Location"
	HeaderOrigin                                    = "Origin"
	HeaderPragma                                    = "Pragma"
	HeaderProxyAuthenticate                         = "Proxy-Authenticate"
	HeaderProxyAuthorization                        = "Proxy-Authorization"
	HeaderPublicKeyPins                             = "Public-Key-Pins"
	HeaderPublicKeyPinsReportOnly                   = "Public-Key-Pins-Report-Only"
	HeaderRange                                     = "Range"
	HeaderReferer                                   = "Referer"
	HeaderRefererPolicy                             = "Referer-Policy"
	HeaderRetryAfter                                = "Retry-After"
	HeaderSaveData                                  = "Save-Data"
	HeaderSecFetchDest                              = "Sec-Fetch-Dest"
	HeaderSecFetchMode                              = "Sec-Fetch-Mode"
	HeaderSecFetchSite                              = "Sec-Fetch-Site"
	HeaderSecFetchUser                              = "Sec-Fetch-User"
	HeaderSecWebSocketAccept                        = "Sec-WebSocket-Accept"
	HeaderServer                                    = "Server"
	HeaderServerTiming                              = "Server-Timing"
	HeaderSetCookie                                 = "Set-Cookie"
	HeaderSourceMap                                 = "SourceMap"
	HeaderHTTPStrictTransportSecurity               = "Strict-Transport-Security:"
	HeaderTE                                        = "TE"
	HeaderTimingAllowOrigin                         = "Timing-Allow-Origin"
	HeaderTk                                        = "Tk"
	HeaderTrailer                                   = "Trailer"
	HeaderTransferEncoding                          = "Transfer-Encoding"
	HeaderUpgradeInsecureRequests                   = "Upgrade-Insecure-Requests"
	HeaderUserAgent                                 = "User-Agent"
	HeaderVary                                      = "Vary"
	HeaderVia                                       = "Via"
	HeaderWWWAuthenticate                           = "WWW-Authenticate"
	HeaderWantDigest                                = "Want-Digest"
	HeaderWarning                                   = "Warning"
	HeaderXContentTypeOptions                       = "X-Content-Type-Options"
	HeaderXDNSPrefetchControl                       = "X-DNS-Prefetch-Control"
	HeaderXForwardedFor                             = "X-Forwarded-For"
	HeaderXForwardedHost                            = "X-Forwarded-Host"
	HeaderXForwardedProto                           = "X-Forwarded-Proto"
	HeaderXFrameOptions                             = "X-Frame-Options"
	HeaderXXSSProtection                            = "X-XSS-Protection"

	HeaderXRequestID     = "X-Request-ID"
	HeaderXRequestDevice = "X-Request-Device"
	HeaderXRequestUser   = "X-Request-User"
)

func LowerHeader(h string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r - 26
		}
		return r
	}, h)
}

type Mime string

const (
	charsetUTF8 = "charset=UTF-8"

	MIMEApplicationJSON                  = "application/json"
	MIMEApplicationJSONCharsetUTF8       = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationJavaScript            = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 = MIMEApplicationJavaScript + "; " + charsetUTF8
	MIMEApplicationXML                   = "application/xml"
	MIMEApplicationXMLCharsetUTF8        = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                          = "text/xml"
	MIMETextXMLCharsetUTF8               = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm                  = "application/x-www-form-urlencoded"
	MIMEApplicationProtobuf              = "application/protobuf"
	MIMEApplicationMsgpack               = "application/msgpack"
	MIMETextHTML                         = "text/html"
	MIMETextHTMLCharsetUTF8              = MIMETextHTML + "; " + charsetUTF8
	MIMETextPlain                        = "text/plain"
	MIMETextPlainCharsetUTF8             = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm                    = "multipart/form-data"
	MIMEOctetStream                      = "application/octet-stream"
)
