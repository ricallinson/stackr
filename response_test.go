package stackr

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestResponse(t *testing.T) {

    Describe("Response.Write()", func() {

        It("should return [true] after writing foo", func() {
            res := CreateResponse(NewMockResponseWriter(false))
            test := res.Write("foo")
            AssertEqual(test, true)
        })

        It("should return [false] after writing foo", func() {
            res := CreateResponse(NewMockResponseWriter(true))
            test := res.Write("foo")
            AssertEqual(test, false)
        })
    })

    Describe("Response.WriteBytes()", func() {

        It("should return [true] after writing 1, 2, 3", func() {
            res := CreateResponse(NewMockResponseWriter(false))
            test := res.WriteBytes([]byte{1, 2, 3})
            AssertEqual(test, true)
        })

        It("should return [false] after writing 1, 2, 3", func() {
            res := CreateResponse(NewMockResponseWriter(true))
            test := res.WriteBytes([]byte{1, 2, 3})
            AssertEqual(test, false)
        })
    })

    Describe("Response.SetHeaders()", func() {

        res := CreateResponse(NewMockResponseWriter(false))

        It("should return [value1] from setting the headers", func() {
            headers := map[string]string{"key0": "value0", "key1": "value1"}
            test := res.SetHeaders(headers)
            AssertEqual(test, true)
            AssertEqual(res.Writer.Header().Get("key1"), "value1")
        })

        It("should return [value] from setting KEY", func() {
            headers := map[string]string{"KEY": "value"}
            res.SetHeaders(headers)
            AssertEqual(res.Writer.Header().Get("key"), "value")
        })

        It("should return [false] from setting key", func() {
            res.HeaderSent = true
            headers := map[string]string{"key": "value"}
            test := res.SetHeaders(headers)
            AssertEqual(test, false)
        })
    })

    Report(t)
}