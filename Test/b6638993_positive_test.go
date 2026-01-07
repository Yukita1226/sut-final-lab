package test

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"

	entity "lab/Entity"
)

func TestXxx(t *testing.T) {

	t.Run(`test correct data`, func(t *testing.T) {
		g := NewGomegaWithT(t)

		e := &entity.Book{
			Title: "hello",
			Price: 50.1,
			Code:  "BK123456",
		}

		ok, err := govalidator.ValidateStruct(e)
		if err != nil {
			fmt.Printf("t: %v\n", err.Error())
		}

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}
