# sut-final-lab
func TestVideo(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Check name not blank", func(t *testing.T) {
		u := Video{
			Name: "",
			Url:  "https://elearning2.sut.ac.th/",
		}
		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))

	})
}
	
git checkout main
git checkout -b issue-3
git add .
git commit -m “ui module equipment - close #3”
git remote update
git rebase origin/main
git checkout main
git merge issue-3 –no-ff
git push origin main
