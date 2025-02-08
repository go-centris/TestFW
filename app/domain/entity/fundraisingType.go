package entity

import (
	// "stncCms/app/domain/dto"
	"html"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

var FundraisingTypeTableName string = "fundraising_type"

// FundraisingType struct
type FundraisingType struct {
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint64     `gorm:"not null;" json:"userId"`
	BranchID     uint64     `gorm:"NOT NUL;DEFAULT:'0'" json:"branchID"  validate:"numeric,omitempty"`
	Title        string     `gorm:"type:varchar(255); null;" json:"title" validate:"required"`
	Picture      string     `gorm:"type:varchar(255); null;" json:"picture"`
	TimeLimit    time.Time  `gorm:"type:date ; NULL;"  json:"tarih"`
	TargetMoney  float64    `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"numeric"  json:"targetMoney"`
	Explanation  string     `gorm:"type:text ;" json:"aciklama" validate:"omitempty"`
	Totalearning float64    `gorm:"-" json:"totalEarning"` // ignore this field when write and read
	CreatedAt    time.Time  ` json:"created_at"`
	UpdatedAt    time.Time  ` json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

// BeforeSave init
func (gk *FundraisingType) BeforeSave() {
	gk.Explanation = html.EscapeString(strings.TrimSpace(gk.Explanation))
}

// Prepare init
func (gk *FundraisingType) Prepare() {
	gk.CreatedAt = time.Now()
	gk.UpdatedAt = time.Now()
}

// TableName override
func (gk *FundraisingType) TableName() string {
	return FundraisingTypeTableName
}

// Validate fluent validation
func (gk *FundraisingType) Validate() map[string]string {
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
	//fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		//fmt.Println(errs)
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
