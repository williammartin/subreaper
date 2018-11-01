package subreaper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/williammartin/subreaper"
)

var _ = Describe("This test process", func() {

	It("reports as a non-subreaper by default", func() {
		isSubreaper, err := subreaper.IsSubreaper()
		Expect(err).NotTo(HaveOccurred())
		Expect(isSubreaper).To(BeFalse())
	})

	When("the subreaper attribute is set", func() {
		var setErr error

		BeforeEach(func() {
			setErr = subreaper.Set()
		})

		It("does not error", func() {
			Expect(setErr).NotTo(HaveOccurred())
		})

		It("reports as a subreaper", func() {
			isSubreaper, err := subreaper.IsSubreaper()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSubreaper).To(BeTrue())
		})

		When("the subreaper attribute is later unset", func() {
			var unsetErr error

			BeforeEach(func() {
				unsetErr = subreaper.Unset()
			})

			It("does not error", func() {
				Expect(unsetErr).NotTo(HaveOccurred())
			})

			It("no longer reports as a subreaper", func() {
				isSubreaper, err := subreaper.IsSubreaper()
				Expect(err).NotTo(HaveOccurred())
				Expect(isSubreaper).To(BeFalse())
			})
		})
	})

})
