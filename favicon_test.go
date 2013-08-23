package stack

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestFavicon(t *testing.T) {
    Describe("AssertEqual()", func() {
        It("should return true", func() {
            AssertEqual(true, true)
        })
    })
    Report(t)
}