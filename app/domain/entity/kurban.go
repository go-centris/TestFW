package entity

import (
	"fmt"
	"html"
	"stncCms/app/domain/dto"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

/*
yeni turler
iki grup arasi degisim yapilmis mi turu ---.> yok bu noticaftion mi olmali ??
grup durumlari olmali ilk 3 madde ona eklenebilir --> yada gerek yok mu durum degisir surekli cunku


	//Kurbanlar BoRC DURUM BİLGİLERİ
	//TODO: sirac e gore eklenecekler
	//parcalama ve TAKSİMATTA
	//SEVK EDİLİYOR
	//TESLİMAT NOKTASINDA
	//KESİLDİ
	//teslim edildi ise teslim edilen kisi bilgileri
*/

const (
	/*
		Dikkat buradaki verileri degistireceksiniz veritabanindan da degisiklik yapiniz
	*/

	/************************************/
	/**********Kurbanlar  DURUM BİLGİLERİ ********/
	/************************************/
	//KurbanDurumKurbanEklendiKurbanBayraminaAitDegil = 10 // kurban eklenmiş ama hisseli yada kurban bayramina ait olmayan kucukbas
	KurbanDurumKurbanGirisKaydiYapildi       = 1 // giris kaydi
	KurbanDurumGrupOlusmusKurbanYok          = 2 // grup oluşmuş ama kimse atanmamış , yani kesimlik kurban verilmemiştir
	KurbanDurumGrupOlusmusKesimlikHayvaniVar = 3 // grup atanmış yani bir kesimlik inek verilmiş
	KurbanDurumKurbanKesimiTamamlanmis       = 4 // kurban kesimi tamamlanmış
	KurbanDurumKurbanParcalamaTaksimatta     = 5 // kurban kesimi tamamlanmış Kurban Parcalama Taksimatta
	KurbanDurumSevkEdiliyor                  = 6 // teslim edilmek icin hazirlaniacak
	KurbanDurumTeslimatNoktasinda            = 7 // teslimat yeri gelmis
	KurbanDurumKurbanKesildiTeslimEdildi     = 7 //  kurban kesildi teslim edildi
	KurbanDurumKurbanKesildiTeslimEdilemedi  = 8 //  kurban kesildi teslim edilemedi
	KurbanKaydiSilindi                       = 9 //  kurban bilgisi silinmis TODO: soft delete kontrol

	//KurbanDurumIkiGrupYerDegistirdi                 = 5 //  iki grup arasi degisim yapildi i  //TODO: ayri bir alan olabilir mi aslinda olmali

	/************************************/
	/********** BORCLAR **BORC DURUM ********/
	/************************************/
	KurbanBorcDurumIlkEklenenFiyat        = 1 // ilk eklenen fiyat değeri
	KurbanBorcDurumKasaBorcluDurumda      = 2 //  kasa borçlu kalmışsa
	KurbanBorcDurumBorcuDevamEdiyor       = 3 //  borcu devam ediyor , taksit yada kapora gibi bir ekleme olmussa bu durum isaretlenir
	KurbanBorcDurumHesapKapandi           = 4 //  tüm borcunu ödemiş
	KurbanBorcDurumKaporaOdemesiHayvanBos = 5 //  kapora odendi ama hayvan atanmamışdır yani bos grup olusmus daha hayvan verilmemesi fakat hayirsever on odeme yapmis

	//KurbanBorcDurumIkiGrupYerDegistirdi   = 5 //  iki grup arasi degisim yapildi  //TODO: ayri bir alan olabilir mi aslinda olmali

	//KurbanBorcDurumTaksitSilindi = 11 //  taksit silndi //TODO:  taksit durum mu bu sadece odemelerde mi tutulsa odeme event surekli degisir burada cok degisim olmaz

	/************************************/
	/********** Vekalet  ********/
	/************************************/
	VekaletDurumuAlinmadi = 1
	VekaletDurumuAlindi   = 2
	//****kisi degisiklik kilidi -- kisi odeme yapmissa artik o kurbandaki kisi bilgileri degistirilmez onun icin onlem
	KisiDegisiklikKilidiKisiDegistirilemez   = 0
	KisiDegisiklikKilidiKisiDegistirilebilir = 1

	/************************************/
	/********** KuRBAN TURLERI  ********/
	/************************************/
	KurbanTuruAdak                       = 1
	KurbanTuruAkika                      = 2
	KurbanTuruSukur                      = 3
	KurbanTuruNiyet                      = 4
	KurbanTuruBagis                      = 5
	KurbanTuruNafile                     = 6
	KurbanTuruSifa                       = 7
	KurbanTuruKurbanBayramiKucukbas      = 8
	KurbanTuruKurbanBayramiHisseliKurban = 9
)

/*
kurban türleri
1- Adak Olarak
2- Akika Olarak Kesilecek
3- Şükür Olarak
4-SAHİBİNİN NİYETİNE Olarak
5-BAĞIŞ Olarak Kesilecek
6-NAFİLE Olarak Kesilecek
7-Şifa Olarak Kesilecek
8- Kurban Bayramı Kesilecek küçük baş
9- Hisseli Büyükbaş
*/

// KurbanTableName table name
var KurbanTableName string = "sacrifice_kurbanlar"

// TODO: bunun dto kopyası olsun orada validate olsun
// Kurban strcut
type SacrificeKurbanlar struct {
	ID            uint64  `gorm:"primary_key;auto_increment" json:"id"`
	UserID        uint64  `gorm:"not null;" json:"user_id"`
	GrupID        uint64  `gorm:"not null;DEFAULT:'0'" json:"grup_id"`
	KisiID        uint64  `gorm:"not null;DEFAULT:'0'" validate:"numeric,omitempty"  json:"kisi_id"`
	Aciklama      string  `gorm:"type:text;" validate:"omitempty,required"  json:"aciklama" `
	KimAdina      string  `gorm:"type:text;" validate:"omitempty,required"  json:"KimAdina" ` //---
	Sertifika     string  `gorm:"type:varchar(255); ;null;" json:"sertifika"`                 //---yenı
	VekaletDurumu int     `gorm:"type:smallint ;NOT NULL;DEFAULT:'0'" validate:"required,omitempty"  json:"vekalet"`
	Agirlik       int     `gorm:"type:smallint ; NULL;"  json:"agirlik"`
	Slug          string  `gorm:"type:varchar(255); ;null;" json:"slug"`
	KurbanTuru    int     `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required,omitempty"  json:"kurbanTuru"`
	Durum         int     `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"omitempty,required"  json:"durum"`
	BorcDurum     int     `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required,omitempty"  json:"BorcDurum"`
	GrupLideri    int     `gorm:"type:smallint ;NOT NULL;DEFAULT:'0'" validate:"omitempty,required"  json:"grupLideri"`
	KurbanFiyati  float64 `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"required,omitempty"  json:"kurbanFiyati"`
	KasaBorcu     float64 `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"numeric"  json:"kasaBorcu"` //TODO: bunun kasa borcu gibi bir isim olmasi lazim
	Alacak        float64 `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"numeric,omitempty"  json:"alacak"`
	Bakiye        float64 `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"numeric,omitempty"  json:"bakiye"`
	HayvanCinsi   int     `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required" json:"hayvanCinsi"`
	//eğer kisi Degisiklik Kilidi 0 ise kiside değişiklik yapılamaz eğer 1 ise değiştireleibr ,
	//-- yenı kurbanda odeme falan yapmıssa kurbana kayıtlı kısının degıstırılmesı sıkıntı olabılır, nosql gibi bısey de log tutmadan bunu yapmak dogru degıldır
	KisiDegisiklikKilidi int                     `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'"  json:"kisiDegisiklikKilidi"` //---burası default degerı daha sonra 1 olacak
	KurbanBayramiYili    int                     `gorm:"type:smallint ;NOT NULL;" validate:"omitempty,required,numeric"`   //otomaitk atılacak isitem tarafından ayarlradan ceksin
	Tarih                time.Time               `gorm:"type:date ; NULL;"  json:"tarih"`
	CreatedAt            time.Time               `json:"created_at"`
	UpdatedAt            time.Time               `json:"updated_at"`
	DeletedAt            *time.Time              `json:"deleted_at"`
	Odemeler             []dto.SacrificeOdemeler `gorm:"foreignKey:KurbanID;references:ID;AssociationForeignKey:ID"`
}

// BeforeSave init
func (gk *SacrificeKurbanlar) BeforeSave() {
	gk.Aciklama = html.EscapeString(strings.TrimSpace(gk.Aciklama))
}

// Prepare init
func (gk *SacrificeKurbanlar) Prepare() {
	gk.CreatedAt = time.Now()
	gk.UpdatedAt = time.Now()
}

// TableName override
func (gk *SacrificeKurbanlar) TableName() string {
	return KurbanTableName
}

// Validate fluent validation
func (gk *SacrificeKurbanlar) Validate() map[string]string {
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
