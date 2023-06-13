package morse_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/panpito/youtube/codewars/morse"
)

var _ = Describe("Test Example", func() {
	It("Example from description", func() {
		Expect(DecodeMorse(".... . -.--   .--- ..- -.. .")).To(Equal("HEY JUDE"))
	})
})
