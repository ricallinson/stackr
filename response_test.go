package stackr

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestResponse(t *testing.T) {

	Describe("Response.Emit()", func() {

		It("should return [0]", func() {
			res := createResponse(NewMockResponseWriter(false))
			test := 0
			res.On("foo", func() {
				test++
			})
			AssertEqual(test, 0)
		})

		It("should return [1]", func() {
			res := createResponse(NewMockResponseWriter(false))
			test := 0
			res.On("foo", func() {
				test++
			})
			res.Emit("foo")
			AssertEqual(test, 1)
		})

		It("should return [3]", func() {
			res := createResponse(NewMockResponseWriter(false))
			test := 0
			res.On("foo", func() {
				test++
			})
			res.On("foo", func() {
				test++
			})
			res.On("foo", func() {
				test++
			})
			res.Emit("foo")
			AssertEqual(test, 3)
		})
	})

	Describe("Response.Write()", func() {

		It("should return [true] after writing foo", func() {
			res := createResponse(NewMockResponseWriter(false))
			test := res.Write("foo")
			AssertEqual(test, true)
		})

		It("should return [false] after writing foo", func() {
			res := createResponse(NewMockResponseWriter(true))
			test := res.Write("foo")
			AssertEqual(test, false)
		})
	})

	Describe("Response.WriteBytes()", func() {

		It("should return [true] after writing 1, 2, 3", func() {
			res := createResponse(NewMockResponseWriter(false))
			test := res.WriteBytes([]byte{1, 2, 3})
			AssertEqual(test, true)
		})

		It("should return [false] after writing 1, 2, 3", func() {
			res := createResponse(NewMockResponseWriter(true))
			test := res.WriteBytes([]byte{1, 2, 3})
			AssertEqual(test, false)
		})
	})

	Describe("Response.SetHeaders()", func() {

		res := createResponse(NewMockResponseWriter(false))

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

	Describe("Response.RemoveHeader()", func() {

		res := createResponse(NewMockResponseWriter(false))

		It("should return [value1] from setting the headers", func() {
			res.SetHeader("foo", "bar")
			AssertEqual(res.Writer.Header().Get("foo"), "bar")
			res.RemoveHeader("foo")
			AssertEqual(res.Writer.Header().Get("foo"), "")
		})
	})

	Report(t)
}
