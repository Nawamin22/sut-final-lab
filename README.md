# sut-final-lab
func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Check name not blank", func(t *testing.T) {
		u := User{
			Name: "",
			Url:  "https://elearning2.sut.ac.th/",
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
		}
		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("name not blank"))
	})
	
git checkout main
git checkout -b issue-3
git add .
git commit -m “ui module equipment - close #3”
git remote update
git rebase origin/main
git checkout main
git merge issue-3 –no-ff
git push origin main
