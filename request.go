package stackr

import(
    "net/http"
    "github.com/ricallinson/httphelp"
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

    // Holds custom values set by functions in the request flow.
    Map map[string]interface{}

    //
    accepted []string

    //
    acceptedLanguages []string

    //
    acceptedCharsets []string
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

    // A map for storing general key/values over the lifetime of the request.
    if this.Map == nil {
        this.Map = map[string]interface{}{}
    }

    return this
}

/*
    Return an slice of Accepted media types ordered from highest quality to lowest.
*/
func (this *Request) Accepted() ([]string) {
    if this.accepted == nil {
        a := this.Header.Get("Accept")
        this.accepted = this.processAccepted(a)
    }
    return this.accepted
}

/*
    Return an slice of Accepted languages ordered from highest quality to lowest.
*/
func (this *Request) AcceptedLanguages() ([]string) {
    if this.acceptedLanguages == nil {
        a := this.Header.Get("Accept-Language")
        this.acceptedLanguages = this.processAccepted(a)
    }
    return this.acceptedLanguages
}

/*
    Return an slice of Accepted charsets ordered from highest quality to lowest.
*/
func (this *Request) AcceptedCharsets() ([]string) {
    if this.acceptedCharsets == nil {
        a := this.Header.Get("Accept-Charset")
        this.acceptedCharsets = this.processAccepted(a)
    }
    return this.acceptedCharsets
}

/*
    Return an slice of "accepted" ordered from highest quality to lowest.
*/
func (this *Request) processAccepted(a string) (list []string) {
    for _, accept := range httphelp.ParseAccept(a) {
        if len(accept.SubType) > 0 {
            list = append(list, accept.Type + "/" + accept.SubType)
        } else {
            list = append(list, accept.Type)
        }
    }
    return
}