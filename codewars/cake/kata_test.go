// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package cake_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/panpito/youtube/codewars/cake"
)

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200})).To(Equal(2))
		Expect(Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}, map[string]int{"sugar": 500, "flour": 2000, "milk": 2000})).To(Equal(0))
	})
})
