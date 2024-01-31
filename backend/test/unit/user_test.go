package unit

import (
	"fmt"
	"testing"

	"github.com/B6406325/labtest/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUserOK(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`OK`, func(t *testing.T){
		user := entity.User{
			ID: "B6406325",
			Name: "Pariwat",
			Phone: "0123456789",
			Age: 20,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestUserID(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`ID is required`, func(t *testing.T){
		user := entity.User{
			ID: "",
			Name: "Pariwat",
			Phone: "0123456789",
			Age: 20,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("ID is required"))
	})

	t.Run(`ID invalid`, func(t *testing.T){
		user := entity.User{
			ID: "N1234567",
			Name: "Pariwat",
			Phone: "0123456789",
			Age: 20,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("ID: %s does not validate as matches(^[B]\\d{7}$)", user.ID)))
	})
}

func TestUserPhone(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Phone is required`, func(t *testing.T){
		user := entity.User{
			ID: "B6400000",
			Name: "Pariwat",
			Phone: "",
			Age: 20,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run(`Phone invalid`, func(t *testing.T){
		user := entity.User{
			ID: "B6400000",
			Name: "Pariwat",
			Phone: "012345",
			Age: 20,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", user.Phone)))
	})
}