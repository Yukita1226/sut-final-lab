package test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"

	entity "lab/Entity"
)

func TestXxx2(t *testing.T) {

	t.Run(`test wrong price(toohigh)`, func(t *testing.T) {
		g := NewGomegaWithT(t)

		e := &entity.Book{
			Title: "hello",
			Price: 5001, // overflow
			Code:  "BK123456",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))

	})

	t.Run(`test wrong price(toolow)`, func(t *testing.T) {
		g := NewGomegaWithT(t)

		e := &entity.Book{
			Title: "hello",
			Price: 49.9, // low
			Code:  "BK123456",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))

	})
}
