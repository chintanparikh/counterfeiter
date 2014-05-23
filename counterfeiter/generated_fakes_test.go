package counterfeiter_test

import (
	"errors"

	"github.com/maxbrunsfeld/counterfeiter/fixtures/fakes"
	"github.com/maxbrunsfeld/counterfeiter/fixtures/interfaces"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generated fakes", func() {
	var fake *fakes.FakeSomeInterface

	BeforeEach(func() {
		fake = fakes.NewFakeSomeInterface()
	})

	It("implements the interface", func() {
		var fake interfaces.SomeInterface
		fake = fakes.NewFakeSomeInterface()
		Expect(fake).NotTo(BeNil())
	})

	It("can have its behavior configured", func() {
		fake.Method1Stub = func(arg1 string, arg2 uint64) error {
			Expect(arg1).To(Equal("stuff"))
			Expect(arg2).To(Equal(uint64(5)))
			return errors.New("hi")
		}

		ret := fake.Method1("stuff", 5)

		Expect(ret).To(Equal(errors.New("hi")))
	})

	It("records the arguments it was called with", func() {
		Expect(fake.Method1Calls()).To(HaveLen(0))

		fake.Method1Stub = func(arg1 string, arg2 uint64) error {
			return nil
		}
		fake.Method1("stuff", 5)

		Expect(fake.Method1Calls()).To(HaveLen(1))
		Expect(fake.Method1Calls()[0].Arg1).To(Equal("stuff"))
		Expect(fake.Method1Calls()[0].Arg2).To(Equal(uint64(5)))
	})

	It("records its calls without race conditions", func() {
		fake.Method2Stub = func() {}

		go fake.Method2()

		Eventually(fake.Method2Calls, 1.0).Should(HaveLen(1))
	})
})
