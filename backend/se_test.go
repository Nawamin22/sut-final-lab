package goapitest

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

type User struct {
	Name string `valid:"required~name not blank"`
	Url  string `valid:"url" gorm:"uniqueIndex"`
	// Email    string
	// Phone    string
	// Password string
}

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Check name not blank", func(t *testing.T) {
		u := User{
			Name: "",
			Url:  "https://elearning2.sut.ac.th/",
			// Email:    "nawa@gmail.com",
			// Phone:    "0973377714",
			// Password: "123456",
		}
		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("name not blank"))
	})
	t.Run("Check Url invalid", func(t *testing.T) {
		u := User{
			Name: "Mana",
			Url:  "://elearning2.sut.ac.th/",
			// Email:    "nawa@gmail.com",
			// Phone:    "0973377714",
			// Password: "123456",
		}
		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url: ://elearning2.sut.ac.th/ does not validate as url"))
	})
	t.Run("Check Url invalid and name not blank", func(t *testing.T) {
		u := User{
			Name: "",
			Url:  "://elearning2.sut.ac.th/",
			// Email:    "nawa@gmail.com",
			// Phone:    "0973377714",
			// Password: "123456",
		}
		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url: ://elearning2.sut.ac.th/ does not validate as url;name not blank"))
	})

}
