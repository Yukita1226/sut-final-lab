package test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"

	entity "lab/Entity"
)

func TestXxx2(t *testing.T) {

	t.Run(`test wrong code 1`, func(t *testing.T) {
		g := NewGomegaWithT(t)

		e := &entity.Book{
			Title: "hello",
			Price: 50,
			Code:  "BK-123456",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))

	})

	t.Run(`test wrong code 2`, func(t *testing.T) {
		g := NewGomegaWithT(t)

		e := &entity.Book{
			Title: "hello",
			Price: 50,
			Code:  "BK1234566456",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))

	})

	t.Run(`test wrong code 3`, func(t *testing.T) {
		g := NewGomegaWithT(t)

		e := &entity.Book{
			Title: "hello",
			Price: 50,
			Code:  "BK1fw456",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))

	})

	t.Run(`test wrong code 4`, func(t *testing.T) {
		g := NewGomegaWithT(t)

		e := &entity.Book{
			Title: "hello",
			Price: 50,
			Code:  "Bf1234566456",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))

	})
}
