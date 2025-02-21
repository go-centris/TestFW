package entity

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
	"strings"
)

// TODO: bunun dto kopyası olsun orada validate olsun
// Kurban strcut
type Currency struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Country  string `gorm:"type:varchar(100); ;null;" json:"country"`
	Currency string `gorm:"type:varchar(100); ;null;" json:"currency"`
	Code     string `gorm:"type:varchar(100); ;null;" json:"code"`
	Symbol   string `gorm:"type:varchar(100); ;null;" json:"symbol"`

	//CreatedAt time.Time  `json:"created_at"`
	//UpdatedAt time.Time  `json:"updated_at"`
	//DeletedAt *time.Time `json:"deleted_at"`
}

// Prepare init
//func (gk *Currency) Prepare() {
//	gk.CreatedAt = time.Now()
//	gk.UpdatedAt = time.Now()
//}

// Validate fluent validation
func (gk *Currency) Validate() map[string]string {
	var (
		validate *validator.Validate
		uni      *ut.UniversalTranslator
	)
	tr := en.New()
	uni = ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")
	validate = validator.New()
	tr_translations.RegisterDefaultTranslations(validate, trans)
	errorLog := make(map[string]string)
	err := validate.Struct(gk)
	fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			// can translate each error one at a time.
			lng := strings.Replace(e.Translate(trans), e.Field(), "Burası", 1)
			errorLog[e.Field()+"_error"] = e.Translate(trans)
			// errorLog[e.Field()] = e.Translate(trans)
			errorLog[e.Field()] = lng
			errorLog[e.Field()+"_valid"] = "is-invalid"
		}
	}
	return errorLog
}
