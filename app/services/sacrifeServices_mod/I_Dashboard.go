package sacrifeServices_mod

import "stncCms/app/domain/dto"

// DashboardAppInterface service
type DashboardAppInterface interface {
	TotalPrice(*float64)                                                                // genel toplam kesilen kurban parasi
	SacrifeSharedPriceTotal(*float64)                                                   //sadece hisseli kurban parasi toplami
	TotalSacrife(*int64)                                                                // genel toplam kesilen kurban parasi
	RemainingDebt(*float64)                                                             // Toplam Kalan Borc
	SharedSacrifeTotal(*int64)                                                          // Hisseli Kesilen Toplam Kurban
	SharedSacrifeRemainingDebt(*float64)                                                // sadece hisseli  Kalan toplam Borc
	ShareSacrifeCount2021(*int64)                                                       //sadece hisseli kurban miktari yil 2021
	ShareSacrifeCount2022(*int64)                                                       //sadece hisseli kurban miktari yil 2022
	ShareSacrifeCount2023(*int64)                                                       //sadece hisseli kurban miktari yil 2023
	SharedSacrifeRemainingDebt2021(*float64)                                            //sadece hisseli  Kalan toplam Borc -- 2021
	SharedSacrifeRemainingDebt2022(*float64)                                            //sadece hisseli  Kalan toplam Borc -- 2022
	SharedSacrifeRemainingDebt2023(*float64)                                            //sadece hisseli  Kalan toplam Borc -- 2023
	SacrifeSharedPriceTotal2021(*float64)                                               //sadece hisseli kurban paRASI 2021
	SacrifeSharedPriceTotal2022(*float64)                                               //sadece hisseli kurban paRASI 2022
	SacrifeSharedPriceTotal2023(*float64)                                               //sadece hisseli kurban paRASI 2023
	CharitableWhoAddedMostSacrife() (*dto.CharitableWhoAddedMostSacrife, error)         // En cok kurban kestiren hayirsever
	UsersWhoAddedMostSacrifeAndBranch() (*dto.UsersWhoAddedMostSacrifeAndBranch, error) //En cok kurban ekleyen  subemiz
	UsersWhoAddedMostSacrifeAndUser() (*dto.UsersWhoAddedMostSacrifeAndBranch, error)   //En cok kurban ekleyen  hocamiz
}

// DashboardApp struct  init
type DashboardApp struct {
	request DashboardAppInterface
}

var _ DashboardAppInterface = &DashboardApp{}

// TotalPrice genel toplam kesilen kurban parasi
// TotalPrice genel toplam kesilen kurban parasi
func (f *DashboardApp) TotalPrice(returnValue *float64) {
	f.request.TotalPrice(returnValue)
}

// SacrifeSharedPriceTotal sadece hisseli kurban miktari
func (f *DashboardApp) SacrifeSharedPriceTotal(returnValue *float64) {
	f.request.SacrifeSharedPriceTotal(returnValue)
}

// TotalSacrife genel toplam kesilen kurban adet
func (f *DashboardApp) TotalSacrife(returnValue *int64) {
	f.request.TotalSacrife(returnValue)
}

// RemainingDebt Toplam Kalan Borc
func (f *DashboardApp) RemainingDebt(returnValue *float64) {
	f.request.RemainingDebt(returnValue)
}

// SharedSacrifeTotal Hisseli Kesilen Toplam Kurban
func (f *DashboardApp) SharedSacrifeTotal(returnValue *int64) {
	f.request.SharedSacrifeTotal(returnValue)
}

// SharedSacrifeRemainingDebt  sadece hisseli  Kalan toplam Borc
func (f *DashboardApp) SharedSacrifeRemainingDebt(returnValue *float64) {
	f.request.SharedSacrifeRemainingDebt(returnValue)
}

// ShareSacrifeCount2021  sadece hisseli kurban miktari yil 2022
func (f *DashboardApp) ShareSacrifeCount2021(returnValue *int64) {
	f.request.ShareSacrifeCount2021(returnValue)
}

// ShareSacrifeCount2022  sadece hisseli kurban miktari yil 2023
func (f *DashboardApp) ShareSacrifeCount2022(returnValue *int64) {
	f.request.ShareSacrifeCount2022(returnValue)
}

// ShareSacrifeCount2023  sadece hisseli kurban miktari yil 2023
func (f *DashboardApp) ShareSacrifeCount2023(returnValue *int64) {
	f.request.ShareSacrifeCount2023(returnValue)
}

// SharedSacrifeRemainingDebt2021  sadece hisseli  Kalan toplam Borc -- 2021
func (f *DashboardApp) SharedSacrifeRemainingDebt2021(returnValue *float64) {
	f.request.SharedSacrifeRemainingDebt2021(returnValue)
}

// SharedSacrifeRemainingDebt2022  sadece hisseli  Kalan toplam Borc -- 2022
func (f *DashboardApp) SharedSacrifeRemainingDebt2022(returnValue *float64) {
	f.request.SharedSacrifeRemainingDebt2022(returnValue)
}

// SharedSacrifeRemainingDebt2023  sadece hisseli  Kalan toplam Borc -- 2023
func (f *DashboardApp) SharedSacrifeRemainingDebt2023(returnValue *float64) {
	f.request.SharedSacrifeRemainingDebt2023(returnValue)
}

// SacrifeSharedPriceTotal2021  sadece hisseli kurban paRASI 2021
func (f *DashboardApp) SacrifeSharedPriceTotal2021(returnValue *float64) {
	f.request.SacrifeSharedPriceTotal2021(returnValue)
}

// SacrifeSharedPriceTotal2022  sadece hisseli kurban paRASI 2021
func (f *DashboardApp) SacrifeSharedPriceTotal2022(returnValue *float64) {
	f.request.SacrifeSharedPriceTotal2022(returnValue)
}

// SacrifeSharedPriceTotal2023  sadece hisseli kurban paRASI 2023
func (f *DashboardApp) SacrifeSharedPriceTotal2023(returnValue *float64) {
	f.request.SacrifeSharedPriceTotal2023(returnValue)
}

// CharitableWhoAddedMostSacrife  En cok kurban kestiren hayirsever
func (f *DashboardApp) CharitableWhoAddedMostSacrife() (*dto.CharitableWhoAddedMostSacrife, error) {
	return f.request.CharitableWhoAddedMostSacrife()
}

// UsersWhoAddedMostSacrifeAndBranch  En cok kurban ekleyen  subemiz
func (f *DashboardApp) UsersWhoAddedMostSacrifeAndBranch() (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	return f.request.UsersWhoAddedMostSacrifeAndBranch()
}

// UsersWhoAddedMostSacrifeAndUser  En cok kurban ekleyen hocamiz
func (f *DashboardApp) UsersWhoAddedMostSacrifeAndUser() (*dto.UsersWhoAddedMostSacrifeAndBranch, error) {
	return f.request.UsersWhoAddedMostSacrifeAndUser()
}
