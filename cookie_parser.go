package stackr

import(
    // "github.com/gorilla/securecookie"
)

func CookieParser(s ...string) (func(req *Request, res *Response, next func())) {

    secret := ""

    if len(s) == 1 {
        secret = s[0]
    }

    return func(req *Request, res *Response, next func()) {

        if req.Cookies != nil {
            next()
            return
        }

        req.Map["cookie-secret"] = secret
        req.Cookies = map[string]string{}
        req.SignedCookies = map[string]string{}

        if len(secret) > 0 {
            for _, c := range req.Request.Cookies() {
                req.SignedCookies[c.Name] = unsignCookie(c.Value, secret)
            }
        }
        for _, c := range req.Request.Cookies() {
            req.Cookies[c.Name] = c.Value
        }
    }
}

// Use github.com/gorilla/securecookie
func unsignCookie(v string, s string) (string) {
    panic(HALT)
}

// Use github.com/gorilla/securecookie
func signCookie(v string, s string) (string) {
    panic(HALT)
}