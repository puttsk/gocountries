package gocountries_test

import (
	"testing"

	gocountries "github.com/puttsk/go-countries"
)

func TestGetCountryFromAlpha3(t *testing.T) {
    var c gocountries.Country

	c = gocountries.GetCountryFromAlpha3("afg")
    if c.ID != 4 {
		t.Errorf("invalid value: expect: %d actual: %d", 4, c.ID)
	}
    if c.EN != "Afghanistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Afghanistan", c.EN)
	}
	if c.TH != "อัฟกานิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อัฟกานิสถาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("alb")
    if c.ID != 8 {
		t.Errorf("invalid value: expect: %d actual: %d", 8, c.ID)
	}
    if c.EN != "Albania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Albania", c.EN)
	}
	if c.TH != "แอลเบเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลเบเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("dza")
    if c.ID != 12 {
		t.Errorf("invalid value: expect: %d actual: %d", 12, c.ID)
	}
    if c.EN != "Algeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Algeria", c.EN)
	}
	if c.TH != "แอลจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลจีเรีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("and")
    if c.ID != 20 {
		t.Errorf("invalid value: expect: %d actual: %d", 20, c.ID)
	}
    if c.EN != "Andorra" {
		t.Errorf("invalid value: expect: %s actual: %s", "Andorra", c.EN)
	}
	if c.TH != "อันดอร์รา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อันดอร์รา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ago")
    if c.ID != 24 {
		t.Errorf("invalid value: expect: %d actual: %d", 24, c.ID)
	}
    if c.EN != "Angola" {
		t.Errorf("invalid value: expect: %s actual: %s", "Angola", c.EN)
	}
	if c.TH != "แองโกลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แองโกลา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("atg")
    if c.ID != 28 {
		t.Errorf("invalid value: expect: %d actual: %d", 28, c.ID)
	}
    if c.EN != "Antigua and Barbuda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Antigua and Barbuda", c.EN)
	}
	if c.TH != "แอนติกาและบาร์บูดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอนติกาและบาร์บูดา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("arg")
    if c.ID != 32 {
		t.Errorf("invalid value: expect: %d actual: %d", 32, c.ID)
	}
    if c.EN != "Argentina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Argentina", c.EN)
	}
	if c.TH != "อาร์เจนตินา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์เจนตินา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("arm")
    if c.ID != 51 {
		t.Errorf("invalid value: expect: %d actual: %d", 51, c.ID)
	}
    if c.EN != "Armenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Armenia", c.EN)
	}
	if c.TH != "อาร์มีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์มีเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("aus")
    if c.ID != 36 {
		t.Errorf("invalid value: expect: %d actual: %d", 36, c.ID)
	}
    if c.EN != "Australia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Australia", c.EN)
	}
	if c.TH != "ออสเตรเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรเลีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("aut")
    if c.ID != 40 {
		t.Errorf("invalid value: expect: %d actual: %d", 40, c.ID)
	}
    if c.EN != "Austria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Austria", c.EN)
	}
	if c.TH != "ออสเตรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("aze")
    if c.ID != 31 {
		t.Errorf("invalid value: expect: %d actual: %d", 31, c.ID)
	}
    if c.EN != "Azerbaijan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Azerbaijan", c.EN)
	}
	if c.TH != "อาเซอร์ไบจาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาเซอร์ไบจาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bhs")
    if c.ID != 44 {
		t.Errorf("invalid value: expect: %d actual: %d", 44, c.ID)
	}
    if c.EN != "Bahamas" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahamas", c.EN)
	}
	if c.TH != "บาฮามาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาฮามาส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bhr")
    if c.ID != 48 {
		t.Errorf("invalid value: expect: %d actual: %d", 48, c.ID)
	}
    if c.EN != "Bahrain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahrain", c.EN)
	}
	if c.TH != "บาห์เรน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาห์เรน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bgd")
    if c.ID != 50 {
		t.Errorf("invalid value: expect: %d actual: %d", 50, c.ID)
	}
    if c.EN != "Bangladesh" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bangladesh", c.EN)
	}
	if c.TH != "บังกลาเทศ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บังกลาเทศ", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("brb")
    if c.ID != 52 {
		t.Errorf("invalid value: expect: %d actual: %d", 52, c.ID)
	}
    if c.EN != "Barbados" {
		t.Errorf("invalid value: expect: %s actual: %s", "Barbados", c.EN)
	}
	if c.TH != "บาร์เบโดส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาร์เบโดส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("blr")
    if c.ID != 112 {
		t.Errorf("invalid value: expect: %d actual: %d", 112, c.ID)
	}
    if c.EN != "Belarus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belarus", c.EN)
	}
	if c.TH != "เบลารุส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลารุส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bel")
    if c.ID != 56 {
		t.Errorf("invalid value: expect: %d actual: %d", 56, c.ID)
	}
    if c.EN != "Belgium" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belgium", c.EN)
	}
	if c.TH != "เบลเยียม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลเยียม", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("blz")
    if c.ID != 84 {
		t.Errorf("invalid value: expect: %d actual: %d", 84, c.ID)
	}
    if c.EN != "Belize" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belize", c.EN)
	}
	if c.TH != "เบลีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลีซ", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ben")
    if c.ID != 204 {
		t.Errorf("invalid value: expect: %d actual: %d", 204, c.ID)
	}
    if c.EN != "Benin" {
		t.Errorf("invalid value: expect: %s actual: %s", "Benin", c.EN)
	}
	if c.TH != "เบนิน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบนิน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("btn")
    if c.ID != 64 {
		t.Errorf("invalid value: expect: %d actual: %d", 64, c.ID)
	}
    if c.EN != "Bhutan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bhutan", c.EN)
	}
	if c.TH != "ภูฏาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ภูฏาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bol")
    if c.ID != 68 {
		t.Errorf("invalid value: expect: %d actual: %d", 68, c.ID)
	}
    if c.EN != "Bolivia (Plurinational State of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bolivia (Plurinational State of)", c.EN)
	}
	if c.TH != "โบลิเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โบลิเวีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bih")
    if c.ID != 70 {
		t.Errorf("invalid value: expect: %d actual: %d", 70, c.ID)
	}
    if c.EN != "Bosnia and Herzegovina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bosnia and Herzegovina", c.EN)
	}
	if c.TH != "บอสเนียและเฮอร์เซโกวีนา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอสเนียและเฮอร์เซโกวีนา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bwa")
    if c.ID != 72 {
		t.Errorf("invalid value: expect: %d actual: %d", 72, c.ID)
	}
    if c.EN != "Botswana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Botswana", c.EN)
	}
	if c.TH != "บอตสวานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอตสวานา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bra")
    if c.ID != 76 {
		t.Errorf("invalid value: expect: %d actual: %d", 76, c.ID)
	}
    if c.EN != "Brazil" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brazil", c.EN)
	}
	if c.TH != "บราซิล" {
		t.Errorf("invalid value: expect: %s actual: %s", "บราซิล", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("brn")
    if c.ID != 96 {
		t.Errorf("invalid value: expect: %d actual: %d", 96, c.ID)
	}
    if c.EN != "Brunei Darussalam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brunei Darussalam", c.EN)
	}
	if c.TH != "บรูไน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บรูไน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bgr")
    if c.ID != 100 {
		t.Errorf("invalid value: expect: %d actual: %d", 100, c.ID)
	}
    if c.EN != "Bulgaria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bulgaria", c.EN)
	}
	if c.TH != "บัลแกเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "บัลแกเรีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bfa")
    if c.ID != 854 {
		t.Errorf("invalid value: expect: %d actual: %d", 854, c.ID)
	}
    if c.EN != "Burkina Faso" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burkina Faso", c.EN)
	}
	if c.TH != "บูร์กินาฟาโซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บูร์กินาฟาโซ", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("bdi")
    if c.ID != 108 {
		t.Errorf("invalid value: expect: %d actual: %d", 108, c.ID)
	}
    if c.EN != "Burundi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burundi", c.EN)
	}
	if c.TH != "บุรุนดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "บุรุนดี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("khm")
    if c.ID != 116 {
		t.Errorf("invalid value: expect: %d actual: %d", 116, c.ID)
	}
    if c.EN != "Cambodia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cambodia", c.EN)
	}
	if c.TH != "กัมพูชา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัมพูชา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cmr")
    if c.ID != 120 {
		t.Errorf("invalid value: expect: %d actual: %d", 120, c.ID)
	}
    if c.EN != "Cameroon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cameroon", c.EN)
	}
	if c.TH != "แคเมอรูน" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคเมอรูน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("can")
    if c.ID != 124 {
		t.Errorf("invalid value: expect: %d actual: %d", 124, c.ID)
	}
    if c.EN != "Canada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Canada", c.EN)
	}
	if c.TH != "แคนาดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคนาดา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cpv")
    if c.ID != 132 {
		t.Errorf("invalid value: expect: %d actual: %d", 132, c.ID)
	}
    if c.EN != "Cabo Verde" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cabo Verde", c.EN)
	}
	if c.TH != "กาบูเวร์ดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบูเวร์ดี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("caf")
    if c.ID != 140 {
		t.Errorf("invalid value: expect: %d actual: %d", 140, c.ID)
	}
    if c.EN != "Central African Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Central African Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐแอฟริกากลาง" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐแอฟริกากลาง", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tcd")
    if c.ID != 148 {
		t.Errorf("invalid value: expect: %d actual: %d", 148, c.ID)
	}
    if c.EN != "Chad" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chad", c.EN)
	}
	if c.TH != "ชาด" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชาด", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("chl")
    if c.ID != 152 {
		t.Errorf("invalid value: expect: %d actual: %d", 152, c.ID)
	}
    if c.EN != "Chile" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chile", c.EN)
	}
	if c.TH != "ชิลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชิลี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("chn")
    if c.ID != 156 {
		t.Errorf("invalid value: expect: %d actual: %d", 156, c.ID)
	}
    if c.EN != "China" {
		t.Errorf("invalid value: expect: %s actual: %s", "China", c.EN)
	}
	if c.TH != "จีน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จีน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("col")
    if c.ID != 170 {
		t.Errorf("invalid value: expect: %d actual: %d", 170, c.ID)
	}
    if c.EN != "Colombia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Colombia", c.EN)
	}
	if c.TH != "โคลอมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โคลอมเบีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("com")
    if c.ID != 174 {
		t.Errorf("invalid value: expect: %d actual: %d", 174, c.ID)
	}
    if c.EN != "Comoros" {
		t.Errorf("invalid value: expect: %s actual: %s", "Comoros", c.EN)
	}
	if c.TH != "คอโมโรส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอโมโรส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cog")
    if c.ID != 178 {
		t.Errorf("invalid value: expect: %d actual: %d", 178, c.ID)
	}
    if c.EN != "Congo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo", c.EN)
	}
	if c.TH != "สาธารณรัฐคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐคองโก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cod")
    if c.ID != 180 {
		t.Errorf("invalid value: expect: %d actual: %d", 180, c.ID)
	}
    if c.EN != "Congo, Democratic Republic of the" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo, Democratic Republic of the", c.EN)
	}
	if c.TH != "สาธารณรัฐประชาธิปไตยคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐประชาธิปไตยคองโก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cri")
    if c.ID != 188 {
		t.Errorf("invalid value: expect: %d actual: %d", 188, c.ID)
	}
    if c.EN != "Costa Rica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Costa Rica", c.EN)
	}
	if c.TH != "คอสตาริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอสตาริกา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("civ")
    if c.ID != 384 {
		t.Errorf("invalid value: expect: %d actual: %d", 384, c.ID)
	}
    if c.EN != "Côte d'Ivoire" {
		t.Errorf("invalid value: expect: %s actual: %s", "Côte d'Ivoire", c.EN)
	}
	if c.TH != "โกตดิวัวร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โกตดิวัวร์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("hrv")
    if c.ID != 191 {
		t.Errorf("invalid value: expect: %d actual: %d", 191, c.ID)
	}
    if c.EN != "Croatia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Croatia", c.EN)
	}
	if c.TH != "โครเอเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โครเอเชีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cub")
    if c.ID != 192 {
		t.Errorf("invalid value: expect: %d actual: %d", 192, c.ID)
	}
    if c.EN != "Cuba" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cuba", c.EN)
	}
	if c.TH != "คิวบา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิวบา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cyp")
    if c.ID != 196 {
		t.Errorf("invalid value: expect: %d actual: %d", 196, c.ID)
	}
    if c.EN != "Cyprus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cyprus", c.EN)
	}
	if c.TH != "ไซปรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไซปรัส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("cze")
    if c.ID != 203 {
		t.Errorf("invalid value: expect: %d actual: %d", 203, c.ID)
	}
    if c.EN != "Czechia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Czechia", c.EN)
	}
	if c.TH != "เช็กเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เช็กเกีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("dnk")
    if c.ID != 208 {
		t.Errorf("invalid value: expect: %d actual: %d", 208, c.ID)
	}
    if c.EN != "Denmark" {
		t.Errorf("invalid value: expect: %s actual: %s", "Denmark", c.EN)
	}
	if c.TH != "เดนมาร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เดนมาร์ก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("dji")
    if c.ID != 262 {
		t.Errorf("invalid value: expect: %d actual: %d", 262, c.ID)
	}
    if c.EN != "Djibouti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Djibouti", c.EN)
	}
	if c.TH != "จิบูตี" {
		t.Errorf("invalid value: expect: %s actual: %s", "จิบูตี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("dma")
    if c.ID != 212 {
		t.Errorf("invalid value: expect: %d actual: %d", 212, c.ID)
	}
    if c.EN != "Dominica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominica", c.EN)
	}
	if c.TH != "ดอมินีกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ดอมินีกา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("dom")
    if c.ID != 214 {
		t.Errorf("invalid value: expect: %d actual: %d", 214, c.ID)
	}
    if c.EN != "Dominican Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominican Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐโดมินิกัน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐโดมินิกัน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ecu")
    if c.ID != 218 {
		t.Errorf("invalid value: expect: %d actual: %d", 218, c.ID)
	}
    if c.EN != "Ecuador" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ecuador", c.EN)
	}
	if c.TH != "เอกวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอกวาดอร์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("egy")
    if c.ID != 818 {
		t.Errorf("invalid value: expect: %d actual: %d", 818, c.ID)
	}
    if c.EN != "Egypt" {
		t.Errorf("invalid value: expect: %s actual: %s", "Egypt", c.EN)
	}
	if c.TH != "อียิปต์" {
		t.Errorf("invalid value: expect: %s actual: %s", "อียิปต์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("slv")
    if c.ID != 222 {
		t.Errorf("invalid value: expect: %d actual: %d", 222, c.ID)
	}
    if c.EN != "El Salvador" {
		t.Errorf("invalid value: expect: %s actual: %s", "El Salvador", c.EN)
	}
	if c.TH != "เอลซัลวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอลซัลวาดอร์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gnq")
    if c.ID != 226 {
		t.Errorf("invalid value: expect: %d actual: %d", 226, c.ID)
	}
    if c.EN != "Equatorial Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Equatorial Guinea", c.EN)
	}
	if c.TH != "อิเควทอเรียลกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิเควทอเรียลกินี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("eri")
    if c.ID != 232 {
		t.Errorf("invalid value: expect: %d actual: %d", 232, c.ID)
	}
    if c.EN != "Eritrea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eritrea", c.EN)
	}
	if c.TH != "เอริเทรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอริเทรีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("est")
    if c.ID != 233 {
		t.Errorf("invalid value: expect: %d actual: %d", 233, c.ID)
	}
    if c.EN != "Estonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Estonia", c.EN)
	}
	if c.TH != "เอสโตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสโตเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("eth")
    if c.ID != 231 {
		t.Errorf("invalid value: expect: %d actual: %d", 231, c.ID)
	}
    if c.EN != "Ethiopia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ethiopia", c.EN)
	}
	if c.TH != "เอธิโอเปีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอธิโอเปีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("fji")
    if c.ID != 242 {
		t.Errorf("invalid value: expect: %d actual: %d", 242, c.ID)
	}
    if c.EN != "Fiji" {
		t.Errorf("invalid value: expect: %s actual: %s", "Fiji", c.EN)
	}
	if c.TH != "ฟีจี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟีจี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("fin")
    if c.ID != 246 {
		t.Errorf("invalid value: expect: %d actual: %d", 246, c.ID)
	}
    if c.EN != "Finland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Finland", c.EN)
	}
	if c.TH != "ฟินแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟินแลนด์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("fra")
    if c.ID != 250 {
		t.Errorf("invalid value: expect: %d actual: %d", 250, c.ID)
	}
    if c.EN != "France" {
		t.Errorf("invalid value: expect: %s actual: %s", "France", c.EN)
	}
	if c.TH != "ฝรั่งเศส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฝรั่งเศส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gab")
    if c.ID != 266 {
		t.Errorf("invalid value: expect: %d actual: %d", 266, c.ID)
	}
    if c.EN != "Gabon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gabon", c.EN)
	}
	if c.TH != "กาบอง" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบอง", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gmb")
    if c.ID != 270 {
		t.Errorf("invalid value: expect: %d actual: %d", 270, c.ID)
	}
    if c.EN != "Gambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gambia", c.EN)
	}
	if c.TH != "แกมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แกมเบีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("geo")
    if c.ID != 268 {
		t.Errorf("invalid value: expect: %d actual: %d", 268, c.ID)
	}
    if c.EN != "Georgia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Georgia", c.EN)
	}
	if c.TH != "จอร์เจีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์เจีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("deu")
    if c.ID != 276 {
		t.Errorf("invalid value: expect: %d actual: %d", 276, c.ID)
	}
    if c.EN != "Germany" {
		t.Errorf("invalid value: expect: %s actual: %s", "Germany", c.EN)
	}
	if c.TH != "เยอรมนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยอรมนี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gha")
    if c.ID != 288 {
		t.Errorf("invalid value: expect: %d actual: %d", 288, c.ID)
	}
    if c.EN != "Ghana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ghana", c.EN)
	}
	if c.TH != "กานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กานา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("grc")
    if c.ID != 300 {
		t.Errorf("invalid value: expect: %d actual: %d", 300, c.ID)
	}
    if c.EN != "Greece" {
		t.Errorf("invalid value: expect: %s actual: %s", "Greece", c.EN)
	}
	if c.TH != "กรีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "กรีซ", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("grd")
    if c.ID != 308 {
		t.Errorf("invalid value: expect: %d actual: %d", 308, c.ID)
	}
    if c.EN != "Grenada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Grenada", c.EN)
	}
	if c.TH != "เกรเนดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกรเนดา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gtm")
    if c.ID != 320 {
		t.Errorf("invalid value: expect: %d actual: %d", 320, c.ID)
	}
    if c.EN != "Guatemala" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guatemala", c.EN)
	}
	if c.TH != "กัวเตมาลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัวเตมาลา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gin")
    if c.ID != 324 {
		t.Errorf("invalid value: expect: %d actual: %d", 324, c.ID)
	}
    if c.EN != "Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea", c.EN)
	}
	if c.TH != "กินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gnb")
    if c.ID != 624 {
		t.Errorf("invalid value: expect: %d actual: %d", 624, c.ID)
	}
    if c.EN != "Guinea-Bissau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea-Bissau", c.EN)
	}
	if c.TH != "กินี-บิสเซา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี-บิสเซา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("guy")
    if c.ID != 328 {
		t.Errorf("invalid value: expect: %d actual: %d", 328, c.ID)
	}
    if c.EN != "Guyana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guyana", c.EN)
	}
	if c.TH != "กายอานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กายอานา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("hti")
    if c.ID != 332 {
		t.Errorf("invalid value: expect: %d actual: %d", 332, c.ID)
	}
    if c.EN != "Haiti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Haiti", c.EN)
	}
	if c.TH != "เฮติ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เฮติ", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("hnd")
    if c.ID != 340 {
		t.Errorf("invalid value: expect: %d actual: %d", 340, c.ID)
	}
    if c.EN != "Honduras" {
		t.Errorf("invalid value: expect: %s actual: %s", "Honduras", c.EN)
	}
	if c.TH != "ฮอนดูรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮอนดูรัส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("hun")
    if c.ID != 348 {
		t.Errorf("invalid value: expect: %d actual: %d", 348, c.ID)
	}
    if c.EN != "Hungary" {
		t.Errorf("invalid value: expect: %s actual: %s", "Hungary", c.EN)
	}
	if c.TH != "ฮังการี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮังการี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("isl")
    if c.ID != 352 {
		t.Errorf("invalid value: expect: %d actual: %d", 352, c.ID)
	}
    if c.EN != "Iceland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iceland", c.EN)
	}
	if c.TH != "ไอซ์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอซ์แลนด์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ind")
    if c.ID != 356 {
		t.Errorf("invalid value: expect: %d actual: %d", 356, c.ID)
	}
    if c.EN != "India" {
		t.Errorf("invalid value: expect: %s actual: %s", "India", c.EN)
	}
	if c.TH != "อินเดีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินเดีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("idn")
    if c.ID != 360 {
		t.Errorf("invalid value: expect: %d actual: %d", 360, c.ID)
	}
    if c.EN != "Indonesia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Indonesia", c.EN)
	}
	if c.TH != "อินโดนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินโดนีเซีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("irn")
    if c.ID != 364 {
		t.Errorf("invalid value: expect: %d actual: %d", 364, c.ID)
	}
    if c.EN != "Iran (Islamic Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iran (Islamic Republic of)", c.EN)
	}
	if c.TH != "อิหร่าน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิหร่าน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("irq")
    if c.ID != 368 {
		t.Errorf("invalid value: expect: %d actual: %d", 368, c.ID)
	}
    if c.EN != "Iraq" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iraq", c.EN)
	}
	if c.TH != "อิรัก" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิรัก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("irl")
    if c.ID != 372 {
		t.Errorf("invalid value: expect: %d actual: %d", 372, c.ID)
	}
    if c.EN != "Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ireland", c.EN)
	}
	if c.TH != "ไอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอร์แลนด์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("isr")
    if c.ID != 376 {
		t.Errorf("invalid value: expect: %d actual: %d", 376, c.ID)
	}
    if c.EN != "Israel" {
		t.Errorf("invalid value: expect: %s actual: %s", "Israel", c.EN)
	}
	if c.TH != "อิสราเอล" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิสราเอล", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ita")
    if c.ID != 380 {
		t.Errorf("invalid value: expect: %d actual: %d", 380, c.ID)
	}
    if c.EN != "Italy" {
		t.Errorf("invalid value: expect: %s actual: %s", "Italy", c.EN)
	}
	if c.TH != "อิตาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิตาลี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("jam")
    if c.ID != 388 {
		t.Errorf("invalid value: expect: %d actual: %d", 388, c.ID)
	}
    if c.EN != "Jamaica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jamaica", c.EN)
	}
	if c.TH != "จาเมกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "จาเมกา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("jpn")
    if c.ID != 392 {
		t.Errorf("invalid value: expect: %d actual: %d", 392, c.ID)
	}
    if c.EN != "Japan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Japan", c.EN)
	}
	if c.TH != "ญี่ปุ่น" {
		t.Errorf("invalid value: expect: %s actual: %s", "ญี่ปุ่น", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("jor")
    if c.ID != 400 {
		t.Errorf("invalid value: expect: %d actual: %d", 400, c.ID)
	}
    if c.EN != "Jordan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jordan", c.EN)
	}
	if c.TH != "จอร์แดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์แดน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("kaz")
    if c.ID != 398 {
		t.Errorf("invalid value: expect: %d actual: %d", 398, c.ID)
	}
    if c.EN != "Kazakhstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kazakhstan", c.EN)
	}
	if c.TH != "คาซัคสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คาซัคสถาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ken")
    if c.ID != 404 {
		t.Errorf("invalid value: expect: %d actual: %d", 404, c.ID)
	}
    if c.EN != "Kenya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kenya", c.EN)
	}
	if c.TH != "เคนยา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เคนยา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("kir")
    if c.ID != 296 {
		t.Errorf("invalid value: expect: %d actual: %d", 296, c.ID)
	}
    if c.EN != "Kiribati" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kiribati", c.EN)
	}
	if c.TH != "คิริบาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิริบาส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("prk")
    if c.ID != 408 {
		t.Errorf("invalid value: expect: %d actual: %d", 408, c.ID)
	}
    if c.EN != "Korea (Democratic People's Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea (Democratic People's Republic of)", c.EN)
	}
	if c.TH != "เกาหลีเหนือ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีเหนือ", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("kor")
    if c.ID != 410 {
		t.Errorf("invalid value: expect: %d actual: %d", 410, c.ID)
	}
    if c.EN != "Korea, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea, Republic of", c.EN)
	}
	if c.TH != "เกาหลีใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีใต้", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("kwt")
    if c.ID != 414 {
		t.Errorf("invalid value: expect: %d actual: %d", 414, c.ID)
	}
    if c.EN != "Kuwait" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kuwait", c.EN)
	}
	if c.TH != "คูเวต" {
		t.Errorf("invalid value: expect: %s actual: %s", "คูเวต", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("kgz")
    if c.ID != 417 {
		t.Errorf("invalid value: expect: %d actual: %d", 417, c.ID)
	}
    if c.EN != "Kyrgyzstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kyrgyzstan", c.EN)
	}
	if c.TH != "คีร์กีซสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คีร์กีซสถาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lao")
    if c.ID != 418 {
		t.Errorf("invalid value: expect: %d actual: %d", 418, c.ID)
	}
    if c.EN != "Lao People's Democratic Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lao People's Democratic Republic", c.EN)
	}
	if c.TH != "ลาว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลาว", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lva")
    if c.ID != 428 {
		t.Errorf("invalid value: expect: %d actual: %d", 428, c.ID)
	}
    if c.EN != "Latvia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Latvia", c.EN)
	}
	if c.TH != "ลัตเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลัตเวีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lbn")
    if c.ID != 422 {
		t.Errorf("invalid value: expect: %d actual: %d", 422, c.ID)
	}
    if c.EN != "Lebanon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lebanon", c.EN)
	}
	if c.TH != "เลบานอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลบานอน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lso")
    if c.ID != 426 {
		t.Errorf("invalid value: expect: %d actual: %d", 426, c.ID)
	}
    if c.EN != "Lesotho" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lesotho", c.EN)
	}
	if c.TH != "เลโซโท" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลโซโท", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lbr")
    if c.ID != 430 {
		t.Errorf("invalid value: expect: %d actual: %d", 430, c.ID)
	}
    if c.EN != "Liberia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liberia", c.EN)
	}
	if c.TH != "ไลบีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไลบีเรีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lby")
    if c.ID != 434 {
		t.Errorf("invalid value: expect: %d actual: %d", 434, c.ID)
	}
    if c.EN != "Libya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Libya", c.EN)
	}
	if c.TH != "ลิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิเบีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lie")
    if c.ID != 438 {
		t.Errorf("invalid value: expect: %d actual: %d", 438, c.ID)
	}
    if c.EN != "Liechtenstein" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liechtenstein", c.EN)
	}
	if c.TH != "ลิกเตนสไตน์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิกเตนสไตน์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ltu")
    if c.ID != 440 {
		t.Errorf("invalid value: expect: %d actual: %d", 440, c.ID)
	}
    if c.EN != "Lithuania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lithuania", c.EN)
	}
	if c.TH != "ลิทัวเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิทัวเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lux")
    if c.ID != 442 {
		t.Errorf("invalid value: expect: %d actual: %d", 442, c.ID)
	}
    if c.EN != "Luxembourg" {
		t.Errorf("invalid value: expect: %s actual: %s", "Luxembourg", c.EN)
	}
	if c.TH != "ลักเซมเบิร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลักเซมเบิร์ก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mkd")
    if c.ID != 807 {
		t.Errorf("invalid value: expect: %d actual: %d", 807, c.ID)
	}
    if c.EN != "North Macedonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "North Macedonia", c.EN)
	}
	if c.TH != "นอร์ทมาซิโดเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์ทมาซิโดเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mdg")
    if c.ID != 450 {
		t.Errorf("invalid value: expect: %d actual: %d", 450, c.ID)
	}
    if c.EN != "Madagascar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Madagascar", c.EN)
	}
	if c.TH != "มาดากัสการ์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาดากัสการ์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mwi")
    if c.ID != 454 {
		t.Errorf("invalid value: expect: %d actual: %d", 454, c.ID)
	}
    if c.EN != "Malawi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malawi", c.EN)
	}
	if c.TH != "มาลาวี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลาวี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mys")
    if c.ID != 458 {
		t.Errorf("invalid value: expect: %d actual: %d", 458, c.ID)
	}
    if c.EN != "Malaysia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malaysia", c.EN)
	}
	if c.TH != "มาเลเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาเลเซีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mdv")
    if c.ID != 462 {
		t.Errorf("invalid value: expect: %d actual: %d", 462, c.ID)
	}
    if c.EN != "Maldives" {
		t.Errorf("invalid value: expect: %s actual: %s", "Maldives", c.EN)
	}
	if c.TH != "มัลดีฟส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มัลดีฟส์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mli")
    if c.ID != 466 {
		t.Errorf("invalid value: expect: %d actual: %d", 466, c.ID)
	}
    if c.EN != "Mali" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mali", c.EN)
	}
	if c.TH != "มาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mlt")
    if c.ID != 470 {
		t.Errorf("invalid value: expect: %d actual: %d", 470, c.ID)
	}
    if c.EN != "Malta" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malta", c.EN)
	}
	if c.TH != "มอลตา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลตา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mhl")
    if c.ID != 584 {
		t.Errorf("invalid value: expect: %d actual: %d", 584, c.ID)
	}
    if c.EN != "Marshall Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Marshall Islands", c.EN)
	}
	if c.TH != "หมู่เกาะมาร์แชลล์" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะมาร์แชลล์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mrt")
    if c.ID != 478 {
		t.Errorf("invalid value: expect: %d actual: %d", 478, c.ID)
	}
    if c.EN != "Mauritania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritania", c.EN)
	}
	if c.TH != "มอริเตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเตเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mus")
    if c.ID != 480 {
		t.Errorf("invalid value: expect: %d actual: %d", 480, c.ID)
	}
    if c.EN != "Mauritius" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritius", c.EN)
	}
	if c.TH != "มอริเชียส" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเชียส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mex")
    if c.ID != 484 {
		t.Errorf("invalid value: expect: %d actual: %d", 484, c.ID)
	}
    if c.EN != "Mexico" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mexico", c.EN)
	}
	if c.TH != "เม็กซิโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เม็กซิโก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("fsm")
    if c.ID != 583 {
		t.Errorf("invalid value: expect: %d actual: %d", 583, c.ID)
	}
    if c.EN != "Micronesia (Federated States of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Micronesia (Federated States of)", c.EN)
	}
	if c.TH != "ไมโครนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไมโครนีเซีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mar")
    if c.ID != 504 {
		t.Errorf("invalid value: expect: %d actual: %d", 504, c.ID)
	}
    if c.EN != "Morocco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Morocco", c.EN)
	}
	if c.TH != "โมร็อกโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมร็อกโก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mda")
    if c.ID != 498 {
		t.Errorf("invalid value: expect: %d actual: %d", 498, c.ID)
	}
    if c.EN != "Moldova, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Moldova, Republic of", c.EN)
	}
	if c.TH != "มอลโดวา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลโดวา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mco")
    if c.ID != 492 {
		t.Errorf("invalid value: expect: %d actual: %d", 492, c.ID)
	}
    if c.EN != "Monaco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Monaco", c.EN)
	}
	if c.TH != "โมนาโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมนาโก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mng")
    if c.ID != 496 {
		t.Errorf("invalid value: expect: %d actual: %d", 496, c.ID)
	}
    if c.EN != "Mongolia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mongolia", c.EN)
	}
	if c.TH != "มองโกเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มองโกเลีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mne")
    if c.ID != 499 {
		t.Errorf("invalid value: expect: %d actual: %d", 499, c.ID)
	}
    if c.EN != "Montenegro" {
		t.Errorf("invalid value: expect: %s actual: %s", "Montenegro", c.EN)
	}
	if c.TH != "มอนเตเนโกร" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอนเตเนโกร", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("moz")
    if c.ID != 508 {
		t.Errorf("invalid value: expect: %d actual: %d", 508, c.ID)
	}
    if c.EN != "Mozambique" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mozambique", c.EN)
	}
	if c.TH != "โมซัมบิก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมซัมบิก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("mmr")
    if c.ID != 104 {
		t.Errorf("invalid value: expect: %d actual: %d", 104, c.ID)
	}
    if c.EN != "Myanmar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Myanmar", c.EN)
	}
	if c.TH != "พม่า" {
		t.Errorf("invalid value: expect: %s actual: %s", "พม่า", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("nam")
    if c.ID != 516 {
		t.Errorf("invalid value: expect: %d actual: %d", 516, c.ID)
	}
    if c.EN != "Namibia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Namibia", c.EN)
	}
	if c.TH != "นามิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นามิเบีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("nru")
    if c.ID != 520 {
		t.Errorf("invalid value: expect: %d actual: %d", 520, c.ID)
	}
    if c.EN != "Nauru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nauru", c.EN)
	}
	if c.TH != "นาอูรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "นาอูรู", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("npl")
    if c.ID != 524 {
		t.Errorf("invalid value: expect: %d actual: %d", 524, c.ID)
	}
    if c.EN != "Nepal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nepal", c.EN)
	}
	if c.TH != "เนปาล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนปาล", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("nld")
    if c.ID != 528 {
		t.Errorf("invalid value: expect: %d actual: %d", 528, c.ID)
	}
    if c.EN != "Netherlands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Netherlands", c.EN)
	}
	if c.TH != "เนเธอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนเธอร์แลนด์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("nzl")
    if c.ID != 554 {
		t.Errorf("invalid value: expect: %d actual: %d", 554, c.ID)
	}
    if c.EN != "New Zealand" {
		t.Errorf("invalid value: expect: %s actual: %s", "New Zealand", c.EN)
	}
	if c.TH != "นิวซีแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิวซีแลนด์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("nic")
    if c.ID != 558 {
		t.Errorf("invalid value: expect: %d actual: %d", 558, c.ID)
	}
    if c.EN != "Nicaragua" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nicaragua", c.EN)
	}
	if c.TH != "นิการากัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิการากัว", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ner")
    if c.ID != 562 {
		t.Errorf("invalid value: expect: %d actual: %d", 562, c.ID)
	}
    if c.EN != "Niger" {
		t.Errorf("invalid value: expect: %s actual: %s", "Niger", c.EN)
	}
	if c.TH != "ไนเจอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนเจอร์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("nga")
    if c.ID != 566 {
		t.Errorf("invalid value: expect: %d actual: %d", 566, c.ID)
	}
    if c.EN != "Nigeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nigeria", c.EN)
	}
	if c.TH != "ไนจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนจีเรีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("nor")
    if c.ID != 578 {
		t.Errorf("invalid value: expect: %d actual: %d", 578, c.ID)
	}
    if c.EN != "Norway" {
		t.Errorf("invalid value: expect: %s actual: %s", "Norway", c.EN)
	}
	if c.TH != "นอร์เวย์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์เวย์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("omn")
    if c.ID != 512 {
		t.Errorf("invalid value: expect: %d actual: %d", 512, c.ID)
	}
    if c.EN != "Oman" {
		t.Errorf("invalid value: expect: %s actual: %s", "Oman", c.EN)
	}
	if c.TH != "โอมาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "โอมาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("pak")
    if c.ID != 586 {
		t.Errorf("invalid value: expect: %d actual: %d", 586, c.ID)
	}
    if c.EN != "Pakistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Pakistan", c.EN)
	}
	if c.TH != "ปากีสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปากีสถาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("plw")
    if c.ID != 585 {
		t.Errorf("invalid value: expect: %d actual: %d", 585, c.ID)
	}
    if c.EN != "Palau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Palau", c.EN)
	}
	if c.TH != "ปาเลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาเลา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("pan")
    if c.ID != 591 {
		t.Errorf("invalid value: expect: %d actual: %d", 591, c.ID)
	}
    if c.EN != "Panama" {
		t.Errorf("invalid value: expect: %s actual: %s", "Panama", c.EN)
	}
	if c.TH != "ปานามา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปานามา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("png")
    if c.ID != 598 {
		t.Errorf("invalid value: expect: %d actual: %d", 598, c.ID)
	}
    if c.EN != "Papua New Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Papua New Guinea", c.EN)
	}
	if c.TH != "ปาปัวนิวกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาปัวนิวกินี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("pry")
    if c.ID != 600 {
		t.Errorf("invalid value: expect: %d actual: %d", 600, c.ID)
	}
    if c.EN != "Paraguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Paraguay", c.EN)
	}
	if c.TH != "ปารากวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปารากวัย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("per")
    if c.ID != 604 {
		t.Errorf("invalid value: expect: %d actual: %d", 604, c.ID)
	}
    if c.EN != "Peru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Peru", c.EN)
	}
	if c.TH != "เปรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "เปรู", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("phl")
    if c.ID != 608 {
		t.Errorf("invalid value: expect: %d actual: %d", 608, c.ID)
	}
    if c.EN != "Philippines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Philippines", c.EN)
	}
	if c.TH != "ฟิลิปปินส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟิลิปปินส์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("pol")
    if c.ID != 616 {
		t.Errorf("invalid value: expect: %d actual: %d", 616, c.ID)
	}
    if c.EN != "Poland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Poland", c.EN)
	}
	if c.TH != "โปแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปแลนด์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("prt")
    if c.ID != 620 {
		t.Errorf("invalid value: expect: %d actual: %d", 620, c.ID)
	}
    if c.EN != "Portugal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Portugal", c.EN)
	}
	if c.TH != "โปรตุเกส" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปรตุเกส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("qat")
    if c.ID != 634 {
		t.Errorf("invalid value: expect: %d actual: %d", 634, c.ID)
	}
    if c.EN != "Qatar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Qatar", c.EN)
	}
	if c.TH != "กาตาร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาตาร์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("rou")
    if c.ID != 642 {
		t.Errorf("invalid value: expect: %d actual: %d", 642, c.ID)
	}
    if c.EN != "Romania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Romania", c.EN)
	}
	if c.TH != "โรมาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โรมาเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("rus")
    if c.ID != 643 {
		t.Errorf("invalid value: expect: %d actual: %d", 643, c.ID)
	}
    if c.EN != "Russian Federation" {
		t.Errorf("invalid value: expect: %s actual: %s", "Russian Federation", c.EN)
	}
	if c.TH != "รัสเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "รัสเซีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("rwa")
    if c.ID != 646 {
		t.Errorf("invalid value: expect: %d actual: %d", 646, c.ID)
	}
    if c.EN != "Rwanda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Rwanda", c.EN)
	}
	if c.TH != "รวันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "รวันดา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("kna")
    if c.ID != 659 {
		t.Errorf("invalid value: expect: %d actual: %d", 659, c.ID)
	}
    if c.EN != "Saint Kitts and Nevis" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Kitts and Nevis", c.EN)
	}
	if c.TH != "เซนต์คิตส์และเนวิส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์คิตส์และเนวิส", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lca")
    if c.ID != 662 {
		t.Errorf("invalid value: expect: %d actual: %d", 662, c.ID)
	}
    if c.EN != "Saint Lucia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Lucia", c.EN)
	}
	if c.TH != "เซนต์ลูเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์ลูเชีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("vct")
    if c.ID != 670 {
		t.Errorf("invalid value: expect: %d actual: %d", 670, c.ID)
	}
    if c.EN != "Saint Vincent and the Grenadines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Vincent and the Grenadines", c.EN)
	}
	if c.TH != "เซนต์วินเซนต์และเกรนาดีนส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์วินเซนต์และเกรนาดีนส์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("wsm")
    if c.ID != 882 {
		t.Errorf("invalid value: expect: %d actual: %d", 882, c.ID)
	}
    if c.EN != "Samoa" {
		t.Errorf("invalid value: expect: %s actual: %s", "Samoa", c.EN)
	}
	if c.TH != "ซามัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซามัว", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("smr")
    if c.ID != 674 {
		t.Errorf("invalid value: expect: %d actual: %d", 674, c.ID)
	}
    if c.EN != "San Marino" {
		t.Errorf("invalid value: expect: %s actual: %s", "San Marino", c.EN)
	}
	if c.TH != "ซานมารีโน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซานมารีโน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("stp")
    if c.ID != 678 {
		t.Errorf("invalid value: expect: %d actual: %d", 678, c.ID)
	}
    if c.EN != "Sao Tome and Principe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sao Tome and Principe", c.EN)
	}
	if c.TH != "เซาตูเมและปรินซีปี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาตูเมและปรินซีปี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("sau")
    if c.ID != 682 {
		t.Errorf("invalid value: expect: %d actual: %d", 682, c.ID)
	}
    if c.EN != "Saudi Arabia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saudi Arabia", c.EN)
	}
	if c.TH != "ซาอุดีอาระเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซาอุดีอาระเบีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("sen")
    if c.ID != 686 {
		t.Errorf("invalid value: expect: %d actual: %d", 686, c.ID)
	}
    if c.EN != "Senegal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Senegal", c.EN)
	}
	if c.TH != "เซเนกัล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเนกัล", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("srb")
    if c.ID != 688 {
		t.Errorf("invalid value: expect: %d actual: %d", 688, c.ID)
	}
    if c.EN != "Serbia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Serbia", c.EN)
	}
	if c.TH != "เซอร์เบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซอร์เบีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("syc")
    if c.ID != 690 {
		t.Errorf("invalid value: expect: %d actual: %d", 690, c.ID)
	}
    if c.EN != "Seychelles" {
		t.Errorf("invalid value: expect: %s actual: %s", "Seychelles", c.EN)
	}
	if c.TH != "เซเชลส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเชลส์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("sle")
    if c.ID != 694 {
		t.Errorf("invalid value: expect: %d actual: %d", 694, c.ID)
	}
    if c.EN != "Sierra Leone" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sierra Leone", c.EN)
	}
	if c.TH != "เซียร์ราลีโอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซียร์ราลีโอน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("sgp")
    if c.ID != 702 {
		t.Errorf("invalid value: expect: %d actual: %d", 702, c.ID)
	}
    if c.EN != "Singapore" {
		t.Errorf("invalid value: expect: %s actual: %s", "Singapore", c.EN)
	}
	if c.TH != "สิงคโปร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สิงคโปร์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("svk")
    if c.ID != 703 {
		t.Errorf("invalid value: expect: %d actual: %d", 703, c.ID)
	}
    if c.EN != "Slovakia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovakia", c.EN)
	}
	if c.TH != "สโลวาเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวาเกีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("svn")
    if c.ID != 705 {
		t.Errorf("invalid value: expect: %d actual: %d", 705, c.ID)
	}
    if c.EN != "Slovenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovenia", c.EN)
	}
	if c.TH != "สโลวีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวีเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("slb")
    if c.ID != 90 {
		t.Errorf("invalid value: expect: %d actual: %d", 90, c.ID)
	}
    if c.EN != "Solomon Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Solomon Islands", c.EN)
	}
	if c.TH != "หมู่เกาะโซโลมอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะโซโลมอน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("som")
    if c.ID != 706 {
		t.Errorf("invalid value: expect: %d actual: %d", 706, c.ID)
	}
    if c.EN != "Somalia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Somalia", c.EN)
	}
	if c.TH != "โซมาเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โซมาเลีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("zaf")
    if c.ID != 710 {
		t.Errorf("invalid value: expect: %d actual: %d", 710, c.ID)
	}
    if c.EN != "South Africa" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Africa", c.EN)
	}
	if c.TH != "แอฟริกาใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอฟริกาใต้", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ssd")
    if c.ID != 728 {
		t.Errorf("invalid value: expect: %d actual: %d", 728, c.ID)
	}
    if c.EN != "South Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Sudan", c.EN)
	}
	if c.TH != "เซาท์ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาท์ซูดาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("esp")
    if c.ID != 724 {
		t.Errorf("invalid value: expect: %d actual: %d", 724, c.ID)
	}
    if c.EN != "Spain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Spain", c.EN)
	}
	if c.TH != "สเปน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สเปน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("lka")
    if c.ID != 144 {
		t.Errorf("invalid value: expect: %d actual: %d", 144, c.ID)
	}
    if c.EN != "Sri Lanka" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sri Lanka", c.EN)
	}
	if c.TH != "ศรีลังกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ศรีลังกา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("sdn")
    if c.ID != 729 {
		t.Errorf("invalid value: expect: %d actual: %d", 729, c.ID)
	}
    if c.EN != "Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sudan", c.EN)
	}
	if c.TH != "ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูดาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("sur")
    if c.ID != 740 {
		t.Errorf("invalid value: expect: %d actual: %d", 740, c.ID)
	}
    if c.EN != "Suriname" {
		t.Errorf("invalid value: expect: %s actual: %s", "Suriname", c.EN)
	}
	if c.TH != "ซูรินาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูรินาม", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("swz")
    if c.ID != 748 {
		t.Errorf("invalid value: expect: %d actual: %d", 748, c.ID)
	}
    if c.EN != "Eswatini" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eswatini", c.EN)
	}
	if c.TH != "เอสวาตีนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสวาตีนี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("swe")
    if c.ID != 752 {
		t.Errorf("invalid value: expect: %d actual: %d", 752, c.ID)
	}
    if c.EN != "Sweden" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sweden", c.EN)
	}
	if c.TH != "สวีเดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวีเดน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("che")
    if c.ID != 756 {
		t.Errorf("invalid value: expect: %d actual: %d", 756, c.ID)
	}
    if c.EN != "Switzerland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Switzerland", c.EN)
	}
	if c.TH != "สวิตเซอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวิตเซอร์แลนด์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("syr")
    if c.ID != 760 {
		t.Errorf("invalid value: expect: %d actual: %d", 760, c.ID)
	}
    if c.EN != "Syrian Arab Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Syrian Arab Republic", c.EN)
	}
	if c.TH != "ซีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซีเรีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tjk")
    if c.ID != 762 {
		t.Errorf("invalid value: expect: %d actual: %d", 762, c.ID)
	}
    if c.EN != "Tajikistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tajikistan", c.EN)
	}
	if c.TH != "ทาจิกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ทาจิกิสถาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tza")
    if c.ID != 834 {
		t.Errorf("invalid value: expect: %d actual: %d", 834, c.ID)
	}
    if c.EN != "Tanzania, United Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tanzania, United Republic of", c.EN)
	}
	if c.TH != "แทนซาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แทนซาเนีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tha")
    if c.ID != 764 {
		t.Errorf("invalid value: expect: %d actual: %d", 764, c.ID)
	}
    if c.EN != "Thailand" {
		t.Errorf("invalid value: expect: %s actual: %s", "Thailand", c.EN)
	}
	if c.TH != "ไทย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไทย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tls")
    if c.ID != 626 {
		t.Errorf("invalid value: expect: %d actual: %d", 626, c.ID)
	}
    if c.EN != "Timor-Leste" {
		t.Errorf("invalid value: expect: %s actual: %s", "Timor-Leste", c.EN)
	}
	if c.TH != "ติมอร์-เลสเต" {
		t.Errorf("invalid value: expect: %s actual: %s", "ติมอร์-เลสเต", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tgo")
    if c.ID != 768 {
		t.Errorf("invalid value: expect: %d actual: %d", 768, c.ID)
	}
    if c.EN != "Togo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Togo", c.EN)
	}
	if c.TH != "โตโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โตโก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ton")
    if c.ID != 776 {
		t.Errorf("invalid value: expect: %d actual: %d", 776, c.ID)
	}
    if c.EN != "Tonga" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tonga", c.EN)
	}
	if c.TH != "ตองงา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตองงา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tto")
    if c.ID != 780 {
		t.Errorf("invalid value: expect: %d actual: %d", 780, c.ID)
	}
    if c.EN != "Trinidad and Tobago" {
		t.Errorf("invalid value: expect: %s actual: %s", "Trinidad and Tobago", c.EN)
	}
	if c.TH != "ตรินิแดดและโตเบโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตรินิแดดและโตเบโก", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tun")
    if c.ID != 788 {
		t.Errorf("invalid value: expect: %d actual: %d", 788, c.ID)
	}
    if c.EN != "Tunisia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tunisia", c.EN)
	}
	if c.TH != "ตูนิเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูนิเซีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tur")
    if c.ID != 792 {
		t.Errorf("invalid value: expect: %d actual: %d", 792, c.ID)
	}
    if c.EN != "Turkey" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkey", c.EN)
	}
	if c.TH != "ตุรกี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตุรกี", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tkm")
    if c.ID != 795 {
		t.Errorf("invalid value: expect: %d actual: %d", 795, c.ID)
	}
    if c.EN != "Turkmenistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkmenistan", c.EN)
	}
	if c.TH != "เติร์กเมนิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เติร์กเมนิสถาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("tuv")
    if c.ID != 798 {
		t.Errorf("invalid value: expect: %d actual: %d", 798, c.ID)
	}
    if c.EN != "Tuvalu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tuvalu", c.EN)
	}
	if c.TH != "ตูวาลู" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูวาลู", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("uga")
    if c.ID != 800 {
		t.Errorf("invalid value: expect: %d actual: %d", 800, c.ID)
	}
    if c.EN != "Uganda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uganda", c.EN)
	}
	if c.TH != "ยูกันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูกันดา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ukr")
    if c.ID != 804 {
		t.Errorf("invalid value: expect: %d actual: %d", 804, c.ID)
	}
    if c.EN != "Ukraine" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ukraine", c.EN)
	}
	if c.TH != "ยูเครน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูเครน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("are")
    if c.ID != 784 {
		t.Errorf("invalid value: expect: %d actual: %d", 784, c.ID)
	}
    if c.EN != "United Arab Emirates" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Arab Emirates", c.EN)
	}
	if c.TH != "สหรัฐอาหรับเอมิเรตส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอาหรับเอมิเรตส์", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("gbr")
    if c.ID != 826 {
		t.Errorf("invalid value: expect: %d actual: %d", 826, c.ID)
	}
    if c.EN != "United Kingdom of Great Britain and Northern Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Kingdom of Great Britain and Northern Ireland", c.EN)
	}
	if c.TH != "สหราชอาณาจักร" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหราชอาณาจักร", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("usa")
    if c.ID != 840 {
		t.Errorf("invalid value: expect: %d actual: %d", 840, c.ID)
	}
    if c.EN != "United States of America" {
		t.Errorf("invalid value: expect: %s actual: %s", "United States of America", c.EN)
	}
	if c.TH != "สหรัฐอเมริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอเมริกา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ury")
    if c.ID != 858 {
		t.Errorf("invalid value: expect: %d actual: %d", 858, c.ID)
	}
    if c.EN != "Uruguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uruguay", c.EN)
	}
	if c.TH != "อุรุกวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุรุกวัย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("uzb")
    if c.ID != 860 {
		t.Errorf("invalid value: expect: %d actual: %d", 860, c.ID)
	}
    if c.EN != "Uzbekistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uzbekistan", c.EN)
	}
	if c.TH != "อุซเบกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุซเบกิสถาน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("vut")
    if c.ID != 548 {
		t.Errorf("invalid value: expect: %d actual: %d", 548, c.ID)
	}
    if c.EN != "Vanuatu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Vanuatu", c.EN)
	}
	if c.TH != "วานูอาตู" {
		t.Errorf("invalid value: expect: %s actual: %s", "วานูอาตู", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("ven")
    if c.ID != 862 {
		t.Errorf("invalid value: expect: %d actual: %d", 862, c.ID)
	}
    if c.EN != "Venezuela (Bolivarian Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Venezuela (Bolivarian Republic of)", c.EN)
	}
	if c.TH != "เวเนซุเอลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวเนซุเอลา", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("vnm")
    if c.ID != 704 {
		t.Errorf("invalid value: expect: %d actual: %d", 704, c.ID)
	}
    if c.EN != "Viet Nam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Viet Nam", c.EN)
	}
	if c.TH != "เวียดนาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวียดนาม", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("yem")
    if c.ID != 887 {
		t.Errorf("invalid value: expect: %d actual: %d", 887, c.ID)
	}
    if c.EN != "Yemen" {
		t.Errorf("invalid value: expect: %s actual: %s", "Yemen", c.EN)
	}
	if c.TH != "เยเมน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยเมน", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("zmb")
    if c.ID != 894 {
		t.Errorf("invalid value: expect: %d actual: %d", 894, c.ID)
	}
    if c.EN != "Zambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zambia", c.EN)
	}
	if c.TH != "แซมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แซมเบีย", c.TH)
	}
	c = gocountries.GetCountryFromAlpha3("zwe")
    if c.ID != 716 {
		t.Errorf("invalid value: expect: %d actual: %d", 716, c.ID)
	}
    if c.EN != "Zimbabwe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zimbabwe", c.EN)
	}
	if c.TH != "ซิมบับเว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซิมบับเว", c.TH)
	}
}

func BenchmarkGetCountryFromAlpha3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gocountries.GetCountryFromAlpha3("tha")
	}
}

func BenchmarkGetCountryFromAlpha3Map(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gocountries.GetCountryFromAlpha3Map("tha")
	}
}
