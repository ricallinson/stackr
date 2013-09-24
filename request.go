package stackr

import(
    "net/http"
)

/*
    A Request represents an HTTP request received by the server.
*/
type Request struct {

    // The standard http.Request type
    *http.Request

    // The value of .URL.RequestURI() for easy access.
    // Note: this value may be changed by the Stackr.Handle() function.
    Url string

    // Set to the vlue of the matched portion of the .URL.RequestURI()
    MatchedUrl string

    // The value of .URL.RequestURI() for easy access.
    // Note: this value should NEVER be changed.
    OriginalUrl string

    // This property is a map containing the parsed request body. 
    // This feature is provided by the bodyParser() middleware, though other body 
    // parsing middleware may follow this convention as well.
    // This property defaults to {} when bodyParser() is used.
    Body map[string]string

    // This property is a map containing the parsed query-string, defaulting to {}.
    Query map[string]string

    // This property is a map of the files uploaded. This feature is provided 
    // by the bodyParser() middleware, though other body parsing middleware may 
    // follow this convention as well. This property defaults to {} when bodyParser() is used.
    Files map[string]interface{}

    // When the cookieParser() middleware is not used this object defaults to {}, 
    // otherwise contains the cookies sent by the user-agent.
    Cookies map[string]string

    // When the cookieParser(secret) middleware is not used this object defaults to {}, 
    // otherwise contains the signed cookies sent by the user-agent, unsigned and ready for use. 
    // Signed cookies reside in a different object to show developer intent, otherwise a 
    // malicious attack could be placed on `req.cookie` values which are easy to spoof. 
    // Note that signing a cookie does not mean it is "hidden" nor encrypted, this simply 
    // prevents tampering as the secret used to sign is private.
    SignedCookies map[string]string

    // Return an slice of Accepted media types ordered from highest quality to lowest.
    Accepted []string

    // Return the remote address, or when "trust proxy" is enabled - the upstream address.
    Ip string

    // When "trust proxy" is `true`, parse the "X-Forwarded-For" ip address list and return a slice, 
    // otherwise an empty slice is returned. For example if the value were "client, proxy1, proxy2" 
    // you would receive the slice {"client", "proxy1", "proxy2"} where "proxy2" is the furthest down-stream.
    Ips []string

    // Returns the request URL pathname.
    Path string

    // Check if the request is fresh - aka Last-Modified and/or the ETag still match, 
    // indicating that the resource is "fresh".
    Fresh bool

    // Check if the request is stale - aka Last-Modified and/or the ETag do not match, 
    // indicating that the resource is "stale".
    Stale bool

    // Check if the request was issued with the "X-Requested-With" header field set to "XMLHttpRequest" (jQuery etc).
    Xhr bool

    // Return the protocol string "http" or "https" when requested with TLS. 
    // When the "trust proxy" setting is enabled the "X-Forwarded-Proto" header field will be trusted. 
    // If you're running behind a reverse proxy that supplies https for you this may be enabled.
    Protocol string

    // Check if a TLS connection is established. This is a short-hand for: "https" == req.Protocol
    Secure bool

    // Return an slice of Accepted languages ordered from highest quality to lowest.
    AcceptedLanguages []string

    // Return an slice of Accepted charsets ordered from highest quality to lowest.
    AcceptedCharsets []string

    // Holds custom values set by functions in the request flow.
    Map map[string]interface{}
}

/*
    Returns a new Request.
*/

func createRequest(raw *http.Request) (*Request) {

    // Create the Request type.
    this := &Request{
        Request: raw,
        Url: raw.URL.RequestURI(),
        OriginalUrl: raw.URL.RequestURI(),
    }

    // Could be set by middleware.
    if this.Body == nil {
        this.Body = map[string]string{}
    }

    // Could be set by middleware.
    if this.Query == nil {
        this.Query = map[string]string{}
    }

    // Could be set by middleware.
    if this.Files == nil {
        this.Files = map[string]interface{}{}
    }

    // Could be set by middleware.
    if this.Cookies == nil {
        this.Cookies = map[string]string{}
    }

    // Could be set by middleware.
    if this.SignedCookies == nil {
        this.SignedCookies = map[string]string{}
    }

    // Helpers for standard headers.
    this.Accepted = this.getAccepted()
    this.Ip = this.RemoteAddr
    this.Ips = this.getTrustProxy()
    this.Path = this.URL.Path
    this.Fresh = this.getFresh()
    this.Stale = this.Fresh == false
    this.Xhr = this.Header.Get("X-Requested-With") == "XMLHttpRequest"
    this.Protocol = this.URL.Scheme
    this.Secure = this.Protocol == "https"
    this.AcceptedLanguages = this.getAcceptedLanguages()
    this.AcceptedCharsets = this.getAcceptedCharsets()

    // A map for storing general key/values over the lifetime of the request.
    if this.Map == nil {
        this.Map = map[string]interface{}{}
    }

    return this
}

func (this *Request) getTrustProxy() ([]string) {
    return []string{}
}

func (this *Request) getFresh() (bool) {
    return false
}

func (this *Request) getAccepted() ([]string) {
    return []string{}
}

func (this *Request) getAcceptedLanguages() ([]string) {
    return []string{}
}

func (this *Request) getAcceptedCharsets() ([]string) {
    return []string{}
}