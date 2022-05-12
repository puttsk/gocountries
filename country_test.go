package gocountries_test

import (
	"testing"

	gocountries "github.com/puttsk/go-countries"
)

var input = []gocountries.Country{
    {Alpha2: "af", Alpha3: "afg", ID: 4, EN:"Afghanistan", TH:"อัฟกานิสถาน"},
    {Alpha2: "al", Alpha3: "alb", ID: 8, EN:"Albania", TH:"แอลเบเนีย"},
    {Alpha2: "dz", Alpha3: "dza", ID: 12, EN:"Algeria", TH:"แอลจีเรีย"},
    {Alpha2: "ad", Alpha3: "and", ID: 20, EN:"Andorra", TH:"อันดอร์รา"},
    {Alpha2: "ao", Alpha3: "ago", ID: 24, EN:"Angola", TH:"แองโกลา"},
    {Alpha2: "ag", Alpha3: "atg", ID: 28, EN:"Antigua and Barbuda", TH:"แอนติกาและบาร์บูดา"},
    {Alpha2: "ar", Alpha3: "arg", ID: 32, EN:"Argentina", TH:"อาร์เจนตินา"},
    {Alpha2: "am", Alpha3: "arm", ID: 51, EN:"Armenia", TH:"อาร์มีเนีย"},
    {Alpha2: "au", Alpha3: "aus", ID: 36, EN:"Australia", TH:"ออสเตรเลีย"},
    {Alpha2: "at", Alpha3: "aut", ID: 40, EN:"Austria", TH:"ออสเตรีย"},
    {Alpha2: "az", Alpha3: "aze", ID: 31, EN:"Azerbaijan", TH:"อาเซอร์ไบจาน"},
    {Alpha2: "bs", Alpha3: "bhs", ID: 44, EN:"Bahamas", TH:"บาฮามาส"},
    {Alpha2: "bh", Alpha3: "bhr", ID: 48, EN:"Bahrain", TH:"บาห์เรน"},
    {Alpha2: "bd", Alpha3: "bgd", ID: 50, EN:"Bangladesh", TH:"บังกลาเทศ"},
    {Alpha2: "bb", Alpha3: "brb", ID: 52, EN:"Barbados", TH:"บาร์เบโดส"},
    {Alpha2: "by", Alpha3: "blr", ID: 112, EN:"Belarus", TH:"เบลารุส"},
    {Alpha2: "be", Alpha3: "bel", ID: 56, EN:"Belgium", TH:"เบลเยียม"},
    {Alpha2: "bz", Alpha3: "blz", ID: 84, EN:"Belize", TH:"เบลีซ"},
    {Alpha2: "bj", Alpha3: "ben", ID: 204, EN:"Benin", TH:"เบนิน"},
    {Alpha2: "bt", Alpha3: "btn", ID: 64, EN:"Bhutan", TH:"ภูฏาน"},
    {Alpha2: "bo", Alpha3: "bol", ID: 68, EN:"Bolivia (Plurinational State of)", TH:"โบลิเวีย"},
    {Alpha2: "ba", Alpha3: "bih", ID: 70, EN:"Bosnia and Herzegovina", TH:"บอสเนียและเฮอร์เซโกวีนา"},
    {Alpha2: "bw", Alpha3: "bwa", ID: 72, EN:"Botswana", TH:"บอตสวานา"},
    {Alpha2: "br", Alpha3: "bra", ID: 76, EN:"Brazil", TH:"บราซิล"},
    {Alpha2: "bn", Alpha3: "brn", ID: 96, EN:"Brunei Darussalam", TH:"บรูไน"},
    {Alpha2: "bg", Alpha3: "bgr", ID: 100, EN:"Bulgaria", TH:"บัลแกเรีย"},
    {Alpha2: "bf", Alpha3: "bfa", ID: 854, EN:"Burkina Faso", TH:"บูร์กินาฟาโซ"},
    {Alpha2: "bi", Alpha3: "bdi", ID: 108, EN:"Burundi", TH:"บุรุนดี"},
    {Alpha2: "kh", Alpha3: "khm", ID: 116, EN:"Cambodia", TH:"กัมพูชา"},
    {Alpha2: "cm", Alpha3: "cmr", ID: 120, EN:"Cameroon", TH:"แคเมอรูน"},
    {Alpha2: "ca", Alpha3: "can", ID: 124, EN:"Canada", TH:"แคนาดา"},
    {Alpha2: "cv", Alpha3: "cpv", ID: 132, EN:"Cabo Verde", TH:"กาบูเวร์ดี"},
    {Alpha2: "cf", Alpha3: "caf", ID: 140, EN:"Central African Republic", TH:"สาธารณรัฐแอฟริกากลาง"},
    {Alpha2: "td", Alpha3: "tcd", ID: 148, EN:"Chad", TH:"ชาด"},
    {Alpha2: "cl", Alpha3: "chl", ID: 152, EN:"Chile", TH:"ชิลี"},
    {Alpha2: "cn", Alpha3: "chn", ID: 156, EN:"China", TH:"จีน"},
    {Alpha2: "co", Alpha3: "col", ID: 170, EN:"Colombia", TH:"โคลอมเบีย"},
    {Alpha2: "km", Alpha3: "com", ID: 174, EN:"Comoros", TH:"คอโมโรส"},
    {Alpha2: "cg", Alpha3: "cog", ID: 178, EN:"Congo", TH:"สาธารณรัฐคองโก"},
    {Alpha2: "cd", Alpha3: "cod", ID: 180, EN:"Congo, Democratic Republic of the", TH:"สาธารณรัฐประชาธิปไตยคองโก"},
    {Alpha2: "cr", Alpha3: "cri", ID: 188, EN:"Costa Rica", TH:"คอสตาริกา"},
    {Alpha2: "ci", Alpha3: "civ", ID: 384, EN:"Côte d'Ivoire", TH:"โกตดิวัวร์"},
    {Alpha2: "hr", Alpha3: "hrv", ID: 191, EN:"Croatia", TH:"โครเอเชีย"},
    {Alpha2: "cu", Alpha3: "cub", ID: 192, EN:"Cuba", TH:"คิวบา"},
    {Alpha2: "cy", Alpha3: "cyp", ID: 196, EN:"Cyprus", TH:"ไซปรัส"},
    {Alpha2: "cz", Alpha3: "cze", ID: 203, EN:"Czechia", TH:"เช็กเกีย"},
    {Alpha2: "dk", Alpha3: "dnk", ID: 208, EN:"Denmark", TH:"เดนมาร์ก"},
    {Alpha2: "dj", Alpha3: "dji", ID: 262, EN:"Djibouti", TH:"จิบูตี"},
    {Alpha2: "dm", Alpha3: "dma", ID: 212, EN:"Dominica", TH:"ดอมินีกา"},
    {Alpha2: "do", Alpha3: "dom", ID: 214, EN:"Dominican Republic", TH:"สาธารณรัฐโดมินิกัน"},
    {Alpha2: "ec", Alpha3: "ecu", ID: 218, EN:"Ecuador", TH:"เอกวาดอร์"},
    {Alpha2: "eg", Alpha3: "egy", ID: 818, EN:"Egypt", TH:"อียิปต์"},
    {Alpha2: "sv", Alpha3: "slv", ID: 222, EN:"El Salvador", TH:"เอลซัลวาดอร์"},
    {Alpha2: "gq", Alpha3: "gnq", ID: 226, EN:"Equatorial Guinea", TH:"อิเควทอเรียลกินี"},
    {Alpha2: "er", Alpha3: "eri", ID: 232, EN:"Eritrea", TH:"เอริเทรีย"},
    {Alpha2: "ee", Alpha3: "est", ID: 233, EN:"Estonia", TH:"เอสโตเนีย"},
    {Alpha2: "et", Alpha3: "eth", ID: 231, EN:"Ethiopia", TH:"เอธิโอเปีย"},
    {Alpha2: "fj", Alpha3: "fji", ID: 242, EN:"Fiji", TH:"ฟีจี"},
    {Alpha2: "fi", Alpha3: "fin", ID: 246, EN:"Finland", TH:"ฟินแลนด์"},
    {Alpha2: "fr", Alpha3: "fra", ID: 250, EN:"France", TH:"ฝรั่งเศส"},
    {Alpha2: "ga", Alpha3: "gab", ID: 266, EN:"Gabon", TH:"กาบอง"},
    {Alpha2: "gm", Alpha3: "gmb", ID: 270, EN:"Gambia", TH:"แกมเบีย"},
    {Alpha2: "ge", Alpha3: "geo", ID: 268, EN:"Georgia", TH:"จอร์เจีย"},
    {Alpha2: "de", Alpha3: "deu", ID: 276, EN:"Germany", TH:"เยอรมนี"},
    {Alpha2: "gh", Alpha3: "gha", ID: 288, EN:"Ghana", TH:"กานา"},
    {Alpha2: "gr", Alpha3: "grc", ID: 300, EN:"Greece", TH:"กรีซ"},
    {Alpha2: "gd", Alpha3: "grd", ID: 308, EN:"Grenada", TH:"เกรเนดา"},
    {Alpha2: "gt", Alpha3: "gtm", ID: 320, EN:"Guatemala", TH:"กัวเตมาลา"},
    {Alpha2: "gn", Alpha3: "gin", ID: 324, EN:"Guinea", TH:"กินี"},
    {Alpha2: "gw", Alpha3: "gnb", ID: 624, EN:"Guinea-Bissau", TH:"กินี-บิสเซา"},
    {Alpha2: "gy", Alpha3: "guy", ID: 328, EN:"Guyana", TH:"กายอานา"},
    {Alpha2: "ht", Alpha3: "hti", ID: 332, EN:"Haiti", TH:"เฮติ"},
    {Alpha2: "hn", Alpha3: "hnd", ID: 340, EN:"Honduras", TH:"ฮอนดูรัส"},
    {Alpha2: "hu", Alpha3: "hun", ID: 348, EN:"Hungary", TH:"ฮังการี"},
    {Alpha2: "is", Alpha3: "isl", ID: 352, EN:"Iceland", TH:"ไอซ์แลนด์"},
    {Alpha2: "in", Alpha3: "ind", ID: 356, EN:"India", TH:"อินเดีย"},
    {Alpha2: "id", Alpha3: "idn", ID: 360, EN:"Indonesia", TH:"อินโดนีเซีย"},
    {Alpha2: "ir", Alpha3: "irn", ID: 364, EN:"Iran (Islamic Republic of)", TH:"อิหร่าน"},
    {Alpha2: "iq", Alpha3: "irq", ID: 368, EN:"Iraq", TH:"อิรัก"},
    {Alpha2: "ie", Alpha3: "irl", ID: 372, EN:"Ireland", TH:"ไอร์แลนด์"},
    {Alpha2: "il", Alpha3: "isr", ID: 376, EN:"Israel", TH:"อิสราเอล"},
    {Alpha2: "it", Alpha3: "ita", ID: 380, EN:"Italy", TH:"อิตาลี"},
    {Alpha2: "jm", Alpha3: "jam", ID: 388, EN:"Jamaica", TH:"จาเมกา"},
    {Alpha2: "jp", Alpha3: "jpn", ID: 392, EN:"Japan", TH:"ญี่ปุ่น"},
    {Alpha2: "jo", Alpha3: "jor", ID: 400, EN:"Jordan", TH:"จอร์แดน"},
    {Alpha2: "kz", Alpha3: "kaz", ID: 398, EN:"Kazakhstan", TH:"คาซัคสถาน"},
    {Alpha2: "ke", Alpha3: "ken", ID: 404, EN:"Kenya", TH:"เคนยา"},
    {Alpha2: "ki", Alpha3: "kir", ID: 296, EN:"Kiribati", TH:"คิริบาส"},
    {Alpha2: "kp", Alpha3: "prk", ID: 408, EN:"Korea (Democratic People's Republic of)", TH:"เกาหลีเหนือ"},
    {Alpha2: "kr", Alpha3: "kor", ID: 410, EN:"Korea, Republic of", TH:"เกาหลีใต้"},
    {Alpha2: "kw", Alpha3: "kwt", ID: 414, EN:"Kuwait", TH:"คูเวต"},
    {Alpha2: "kg", Alpha3: "kgz", ID: 417, EN:"Kyrgyzstan", TH:"คีร์กีซสถาน"},
    {Alpha2: "la", Alpha3: "lao", ID: 418, EN:"Lao People's Democratic Republic", TH:"ลาว"},
    {Alpha2: "lv", Alpha3: "lva", ID: 428, EN:"Latvia", TH:"ลัตเวีย"},
    {Alpha2: "lb", Alpha3: "lbn", ID: 422, EN:"Lebanon", TH:"เลบานอน"},
    {Alpha2: "ls", Alpha3: "lso", ID: 426, EN:"Lesotho", TH:"เลโซโท"},
    {Alpha2: "lr", Alpha3: "lbr", ID: 430, EN:"Liberia", TH:"ไลบีเรีย"},
    {Alpha2: "ly", Alpha3: "lby", ID: 434, EN:"Libya", TH:"ลิเบีย"},
    {Alpha2: "li", Alpha3: "lie", ID: 438, EN:"Liechtenstein", TH:"ลิกเตนสไตน์"},
    {Alpha2: "lt", Alpha3: "ltu", ID: 440, EN:"Lithuania", TH:"ลิทัวเนีย"},
    {Alpha2: "lu", Alpha3: "lux", ID: 442, EN:"Luxembourg", TH:"ลักเซมเบิร์ก"},
    {Alpha2: "mk", Alpha3: "mkd", ID: 807, EN:"North Macedonia", TH:"นอร์ทมาซิโดเนีย"},
    {Alpha2: "mg", Alpha3: "mdg", ID: 450, EN:"Madagascar", TH:"มาดากัสการ์"},
    {Alpha2: "mw", Alpha3: "mwi", ID: 454, EN:"Malawi", TH:"มาลาวี"},
    {Alpha2: "my", Alpha3: "mys", ID: 458, EN:"Malaysia", TH:"มาเลเซีย"},
    {Alpha2: "mv", Alpha3: "mdv", ID: 462, EN:"Maldives", TH:"มัลดีฟส์"},
    {Alpha2: "ml", Alpha3: "mli", ID: 466, EN:"Mali", TH:"มาลี"},
    {Alpha2: "mt", Alpha3: "mlt", ID: 470, EN:"Malta", TH:"มอลตา"},
    {Alpha2: "mh", Alpha3: "mhl", ID: 584, EN:"Marshall Islands", TH:"หมู่เกาะมาร์แชลล์"},
    {Alpha2: "mr", Alpha3: "mrt", ID: 478, EN:"Mauritania", TH:"มอริเตเนีย"},
    {Alpha2: "mu", Alpha3: "mus", ID: 480, EN:"Mauritius", TH:"มอริเชียส"},
    {Alpha2: "mx", Alpha3: "mex", ID: 484, EN:"Mexico", TH:"เม็กซิโก"},
    {Alpha2: "fm", Alpha3: "fsm", ID: 583, EN:"Micronesia (Federated States of)", TH:"ไมโครนีเซีย"},
    {Alpha2: "ma", Alpha3: "mar", ID: 504, EN:"Morocco", TH:"โมร็อกโก"},
    {Alpha2: "md", Alpha3: "mda", ID: 498, EN:"Moldova, Republic of", TH:"มอลโดวา"},
    {Alpha2: "mc", Alpha3: "mco", ID: 492, EN:"Monaco", TH:"โมนาโก"},
    {Alpha2: "mn", Alpha3: "mng", ID: 496, EN:"Mongolia", TH:"มองโกเลีย"},
    {Alpha2: "me", Alpha3: "mne", ID: 499, EN:"Montenegro", TH:"มอนเตเนโกร"},
    {Alpha2: "mz", Alpha3: "moz", ID: 508, EN:"Mozambique", TH:"โมซัมบิก"},
    {Alpha2: "mm", Alpha3: "mmr", ID: 104, EN:"Myanmar", TH:"พม่า"},
    {Alpha2: "na", Alpha3: "nam", ID: 516, EN:"Namibia", TH:"นามิเบีย"},
    {Alpha2: "nr", Alpha3: "nru", ID: 520, EN:"Nauru", TH:"นาอูรู"},
    {Alpha2: "np", Alpha3: "npl", ID: 524, EN:"Nepal", TH:"เนปาล"},
    {Alpha2: "nl", Alpha3: "nld", ID: 528, EN:"Netherlands", TH:"เนเธอร์แลนด์"},
    {Alpha2: "nz", Alpha3: "nzl", ID: 554, EN:"New Zealand", TH:"นิวซีแลนด์"},
    {Alpha2: "ni", Alpha3: "nic", ID: 558, EN:"Nicaragua", TH:"นิการากัว"},
    {Alpha2: "ne", Alpha3: "ner", ID: 562, EN:"Niger", TH:"ไนเจอร์"},
    {Alpha2: "ng", Alpha3: "nga", ID: 566, EN:"Nigeria", TH:"ไนจีเรีย"},
    {Alpha2: "no", Alpha3: "nor", ID: 578, EN:"Norway", TH:"นอร์เวย์"},
    {Alpha2: "om", Alpha3: "omn", ID: 512, EN:"Oman", TH:"โอมาน"},
    {Alpha2: "pk", Alpha3: "pak", ID: 586, EN:"Pakistan", TH:"ปากีสถาน"},
    {Alpha2: "pw", Alpha3: "plw", ID: 585, EN:"Palau", TH:"ปาเลา"},
    {Alpha2: "pa", Alpha3: "pan", ID: 591, EN:"Panama", TH:"ปานามา"},
    {Alpha2: "pg", Alpha3: "png", ID: 598, EN:"Papua New Guinea", TH:"ปาปัวนิวกินี"},
    {Alpha2: "py", Alpha3: "pry", ID: 600, EN:"Paraguay", TH:"ปารากวัย"},
    {Alpha2: "pe", Alpha3: "per", ID: 604, EN:"Peru", TH:"เปรู"},
    {Alpha2: "ph", Alpha3: "phl", ID: 608, EN:"Philippines", TH:"ฟิลิปปินส์"},
    {Alpha2: "pl", Alpha3: "pol", ID: 616, EN:"Poland", TH:"โปแลนด์"},
    {Alpha2: "pt", Alpha3: "prt", ID: 620, EN:"Portugal", TH:"โปรตุเกส"},
    {Alpha2: "qa", Alpha3: "qat", ID: 634, EN:"Qatar", TH:"กาตาร์"},
    {Alpha2: "ro", Alpha3: "rou", ID: 642, EN:"Romania", TH:"โรมาเนีย"},
    {Alpha2: "ru", Alpha3: "rus", ID: 643, EN:"Russian Federation", TH:"รัสเซีย"},
    {Alpha2: "rw", Alpha3: "rwa", ID: 646, EN:"Rwanda", TH:"รวันดา"},
    {Alpha2: "kn", Alpha3: "kna", ID: 659, EN:"Saint Kitts and Nevis", TH:"เซนต์คิตส์และเนวิส"},
    {Alpha2: "lc", Alpha3: "lca", ID: 662, EN:"Saint Lucia", TH:"เซนต์ลูเชีย"},
    {Alpha2: "vc", Alpha3: "vct", ID: 670, EN:"Saint Vincent and the Grenadines", TH:"เซนต์วินเซนต์และเกรนาดีนส์"},
    {Alpha2: "ws", Alpha3: "wsm", ID: 882, EN:"Samoa", TH:"ซามัว"},
    {Alpha2: "sm", Alpha3: "smr", ID: 674, EN:"San Marino", TH:"ซานมารีโน"},
    {Alpha2: "st", Alpha3: "stp", ID: 678, EN:"Sao Tome and Principe", TH:"เซาตูเมและปรินซีปี"},
    {Alpha2: "sa", Alpha3: "sau", ID: 682, EN:"Saudi Arabia", TH:"ซาอุดีอาระเบีย"},
    {Alpha2: "sn", Alpha3: "sen", ID: 686, EN:"Senegal", TH:"เซเนกัล"},
    {Alpha2: "rs", Alpha3: "srb", ID: 688, EN:"Serbia", TH:"เซอร์เบีย"},
    {Alpha2: "sc", Alpha3: "syc", ID: 690, EN:"Seychelles", TH:"เซเชลส์"},
    {Alpha2: "sl", Alpha3: "sle", ID: 694, EN:"Sierra Leone", TH:"เซียร์ราลีโอน"},
    {Alpha2: "sg", Alpha3: "sgp", ID: 702, EN:"Singapore", TH:"สิงคโปร์"},
    {Alpha2: "sk", Alpha3: "svk", ID: 703, EN:"Slovakia", TH:"สโลวาเกีย"},
    {Alpha2: "si", Alpha3: "svn", ID: 705, EN:"Slovenia", TH:"สโลวีเนีย"},
    {Alpha2: "sb", Alpha3: "slb", ID: 90, EN:"Solomon Islands", TH:"หมู่เกาะโซโลมอน"},
    {Alpha2: "so", Alpha3: "som", ID: 706, EN:"Somalia", TH:"โซมาเลีย"},
    {Alpha2: "za", Alpha3: "zaf", ID: 710, EN:"South Africa", TH:"แอฟริกาใต้"},
    {Alpha2: "ss", Alpha3: "ssd", ID: 728, EN:"South Sudan", TH:"เซาท์ซูดาน"},
    {Alpha2: "es", Alpha3: "esp", ID: 724, EN:"Spain", TH:"สเปน"},
    {Alpha2: "lk", Alpha3: "lka", ID: 144, EN:"Sri Lanka", TH:"ศรีลังกา"},
    {Alpha2: "sd", Alpha3: "sdn", ID: 729, EN:"Sudan", TH:"ซูดาน"},
    {Alpha2: "sr", Alpha3: "sur", ID: 740, EN:"Suriname", TH:"ซูรินาม"},
    {Alpha2: "sz", Alpha3: "swz", ID: 748, EN:"Eswatini", TH:"เอสวาตีนี"},
    {Alpha2: "se", Alpha3: "swe", ID: 752, EN:"Sweden", TH:"สวีเดน"},
    {Alpha2: "ch", Alpha3: "che", ID: 756, EN:"Switzerland", TH:"สวิตเซอร์แลนด์"},
    {Alpha2: "sy", Alpha3: "syr", ID: 760, EN:"Syrian Arab Republic", TH:"ซีเรีย"},
    {Alpha2: "tj", Alpha3: "tjk", ID: 762, EN:"Tajikistan", TH:"ทาจิกิสถาน"},
    {Alpha2: "tz", Alpha3: "tza", ID: 834, EN:"Tanzania, United Republic of", TH:"แทนซาเนีย"},
    {Alpha2: "th", Alpha3: "tha", ID: 764, EN:"Thailand", TH:"ไทย"},
    {Alpha2: "tl", Alpha3: "tls", ID: 626, EN:"Timor-Leste", TH:"ติมอร์-เลสเต"},
    {Alpha2: "tg", Alpha3: "tgo", ID: 768, EN:"Togo", TH:"โตโก"},
    {Alpha2: "to", Alpha3: "ton", ID: 776, EN:"Tonga", TH:"ตองงา"},
    {Alpha2: "tt", Alpha3: "tto", ID: 780, EN:"Trinidad and Tobago", TH:"ตรินิแดดและโตเบโก"},
    {Alpha2: "tn", Alpha3: "tun", ID: 788, EN:"Tunisia", TH:"ตูนิเซีย"},
    {Alpha2: "tr", Alpha3: "tur", ID: 792, EN:"Turkey", TH:"ตุรกี"},
    {Alpha2: "tm", Alpha3: "tkm", ID: 795, EN:"Turkmenistan", TH:"เติร์กเมนิสถาน"},
    {Alpha2: "tv", Alpha3: "tuv", ID: 798, EN:"Tuvalu", TH:"ตูวาลู"},
    {Alpha2: "ug", Alpha3: "uga", ID: 800, EN:"Uganda", TH:"ยูกันดา"},
    {Alpha2: "ua", Alpha3: "ukr", ID: 804, EN:"Ukraine", TH:"ยูเครน"},
    {Alpha2: "ae", Alpha3: "are", ID: 784, EN:"United Arab Emirates", TH:"สหรัฐอาหรับเอมิเรตส์"},
    {Alpha2: "gb", Alpha3: "gbr", ID: 826, EN:"United Kingdom of Great Britain and Northern Ireland", TH:"สหราชอาณาจักร"},
    {Alpha2: "us", Alpha3: "usa", ID: 840, EN:"United States of America", TH:"สหรัฐอเมริกา"},
    {Alpha2: "uy", Alpha3: "ury", ID: 858, EN:"Uruguay", TH:"อุรุกวัย"},
    {Alpha2: "uz", Alpha3: "uzb", ID: 860, EN:"Uzbekistan", TH:"อุซเบกิสถาน"},
    {Alpha2: "vu", Alpha3: "vut", ID: 548, EN:"Vanuatu", TH:"วานูอาตู"},
    {Alpha2: "ve", Alpha3: "ven", ID: 862, EN:"Venezuela (Bolivarian Republic of)", TH:"เวเนซุเอลา"},
    {Alpha2: "vn", Alpha3: "vnm", ID: 704, EN:"Viet Nam", TH:"เวียดนาม"},
    {Alpha2: "ye", Alpha3: "yem", ID: 887, EN:"Yemen", TH:"เยเมน"},
    {Alpha2: "zm", Alpha3: "zmb", ID: 894, EN:"Zambia", TH:"แซมเบีย"},
    {Alpha2: "zw", Alpha3: "zwe", ID: 716, EN:"Zimbabwe", TH:"ซิมบับเว"},
}


func TestGetCountryFromID(t *testing.T) {
    var c,c1 gocountries.Country
	var ok bool

	c, ok = gocountries.GetCountryFromID(4)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 4 {
		t.Errorf("invalid value: expect: %d actual: %d", 4, c.ID)
	}
    if c.EN != "Afghanistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Afghanistan", c.EN)
	}
	if c.TH != "อัฟกานิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อัฟกานิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(4)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(8)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 8 {
		t.Errorf("invalid value: expect: %d actual: %d", 8, c.ID)
	}
    if c.EN != "Albania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Albania", c.EN)
	}
	if c.TH != "แอลเบเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลเบเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(8)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(12)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 12 {
		t.Errorf("invalid value: expect: %d actual: %d", 12, c.ID)
	}
    if c.EN != "Algeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Algeria", c.EN)
	}
	if c.TH != "แอลจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลจีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(12)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(20)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 20 {
		t.Errorf("invalid value: expect: %d actual: %d", 20, c.ID)
	}
    if c.EN != "Andorra" {
		t.Errorf("invalid value: expect: %s actual: %s", "Andorra", c.EN)
	}
	if c.TH != "อันดอร์รา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อันดอร์รา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(20)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(24)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 24 {
		t.Errorf("invalid value: expect: %d actual: %d", 24, c.ID)
	}
    if c.EN != "Angola" {
		t.Errorf("invalid value: expect: %s actual: %s", "Angola", c.EN)
	}
	if c.TH != "แองโกลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แองโกลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(24)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(28)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 28 {
		t.Errorf("invalid value: expect: %d actual: %d", 28, c.ID)
	}
    if c.EN != "Antigua and Barbuda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Antigua and Barbuda", c.EN)
	}
	if c.TH != "แอนติกาและบาร์บูดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอนติกาและบาร์บูดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(28)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(32)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 32 {
		t.Errorf("invalid value: expect: %d actual: %d", 32, c.ID)
	}
    if c.EN != "Argentina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Argentina", c.EN)
	}
	if c.TH != "อาร์เจนตินา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์เจนตินา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(32)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(51)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 51 {
		t.Errorf("invalid value: expect: %d actual: %d", 51, c.ID)
	}
    if c.EN != "Armenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Armenia", c.EN)
	}
	if c.TH != "อาร์มีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์มีเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(51)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(36)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 36 {
		t.Errorf("invalid value: expect: %d actual: %d", 36, c.ID)
	}
    if c.EN != "Australia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Australia", c.EN)
	}
	if c.TH != "ออสเตรเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรเลีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(36)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(40)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 40 {
		t.Errorf("invalid value: expect: %d actual: %d", 40, c.ID)
	}
    if c.EN != "Austria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Austria", c.EN)
	}
	if c.TH != "ออสเตรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(40)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(31)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 31 {
		t.Errorf("invalid value: expect: %d actual: %d", 31, c.ID)
	}
    if c.EN != "Azerbaijan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Azerbaijan", c.EN)
	}
	if c.TH != "อาเซอร์ไบจาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาเซอร์ไบจาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(31)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(44)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 44 {
		t.Errorf("invalid value: expect: %d actual: %d", 44, c.ID)
	}
    if c.EN != "Bahamas" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahamas", c.EN)
	}
	if c.TH != "บาฮามาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาฮามาส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(44)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(48)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 48 {
		t.Errorf("invalid value: expect: %d actual: %d", 48, c.ID)
	}
    if c.EN != "Bahrain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahrain", c.EN)
	}
	if c.TH != "บาห์เรน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาห์เรน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(48)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(50)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 50 {
		t.Errorf("invalid value: expect: %d actual: %d", 50, c.ID)
	}
    if c.EN != "Bangladesh" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bangladesh", c.EN)
	}
	if c.TH != "บังกลาเทศ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บังกลาเทศ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(50)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(52)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 52 {
		t.Errorf("invalid value: expect: %d actual: %d", 52, c.ID)
	}
    if c.EN != "Barbados" {
		t.Errorf("invalid value: expect: %s actual: %s", "Barbados", c.EN)
	}
	if c.TH != "บาร์เบโดส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาร์เบโดส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(52)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(112)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 112 {
		t.Errorf("invalid value: expect: %d actual: %d", 112, c.ID)
	}
    if c.EN != "Belarus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belarus", c.EN)
	}
	if c.TH != "เบลารุส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลารุส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(112)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(56)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 56 {
		t.Errorf("invalid value: expect: %d actual: %d", 56, c.ID)
	}
    if c.EN != "Belgium" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belgium", c.EN)
	}
	if c.TH != "เบลเยียม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลเยียม", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(56)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(84)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 84 {
		t.Errorf("invalid value: expect: %d actual: %d", 84, c.ID)
	}
    if c.EN != "Belize" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belize", c.EN)
	}
	if c.TH != "เบลีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลีซ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(84)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(204)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 204 {
		t.Errorf("invalid value: expect: %d actual: %d", 204, c.ID)
	}
    if c.EN != "Benin" {
		t.Errorf("invalid value: expect: %s actual: %s", "Benin", c.EN)
	}
	if c.TH != "เบนิน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบนิน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(204)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(64)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 64 {
		t.Errorf("invalid value: expect: %d actual: %d", 64, c.ID)
	}
    if c.EN != "Bhutan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bhutan", c.EN)
	}
	if c.TH != "ภูฏาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ภูฏาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(64)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(68)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 68 {
		t.Errorf("invalid value: expect: %d actual: %d", 68, c.ID)
	}
    if c.EN != "Bolivia (Plurinational State of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bolivia (Plurinational State of)", c.EN)
	}
	if c.TH != "โบลิเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โบลิเวีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(68)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(70)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 70 {
		t.Errorf("invalid value: expect: %d actual: %d", 70, c.ID)
	}
    if c.EN != "Bosnia and Herzegovina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bosnia and Herzegovina", c.EN)
	}
	if c.TH != "บอสเนียและเฮอร์เซโกวีนา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอสเนียและเฮอร์เซโกวีนา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(70)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(72)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 72 {
		t.Errorf("invalid value: expect: %d actual: %d", 72, c.ID)
	}
    if c.EN != "Botswana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Botswana", c.EN)
	}
	if c.TH != "บอตสวานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอตสวานา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(72)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(76)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 76 {
		t.Errorf("invalid value: expect: %d actual: %d", 76, c.ID)
	}
    if c.EN != "Brazil" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brazil", c.EN)
	}
	if c.TH != "บราซิล" {
		t.Errorf("invalid value: expect: %s actual: %s", "บราซิล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(76)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(96)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 96 {
		t.Errorf("invalid value: expect: %d actual: %d", 96, c.ID)
	}
    if c.EN != "Brunei Darussalam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brunei Darussalam", c.EN)
	}
	if c.TH != "บรูไน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บรูไน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(96)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(100)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 100 {
		t.Errorf("invalid value: expect: %d actual: %d", 100, c.ID)
	}
    if c.EN != "Bulgaria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bulgaria", c.EN)
	}
	if c.TH != "บัลแกเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "บัลแกเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(100)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(854)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 854 {
		t.Errorf("invalid value: expect: %d actual: %d", 854, c.ID)
	}
    if c.EN != "Burkina Faso" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burkina Faso", c.EN)
	}
	if c.TH != "บูร์กินาฟาโซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บูร์กินาฟาโซ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(854)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(108)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 108 {
		t.Errorf("invalid value: expect: %d actual: %d", 108, c.ID)
	}
    if c.EN != "Burundi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burundi", c.EN)
	}
	if c.TH != "บุรุนดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "บุรุนดี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(108)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(116)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 116 {
		t.Errorf("invalid value: expect: %d actual: %d", 116, c.ID)
	}
    if c.EN != "Cambodia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cambodia", c.EN)
	}
	if c.TH != "กัมพูชา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัมพูชา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(116)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(120)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 120 {
		t.Errorf("invalid value: expect: %d actual: %d", 120, c.ID)
	}
    if c.EN != "Cameroon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cameroon", c.EN)
	}
	if c.TH != "แคเมอรูน" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคเมอรูน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(120)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(124)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 124 {
		t.Errorf("invalid value: expect: %d actual: %d", 124, c.ID)
	}
    if c.EN != "Canada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Canada", c.EN)
	}
	if c.TH != "แคนาดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคนาดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(124)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(132)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 132 {
		t.Errorf("invalid value: expect: %d actual: %d", 132, c.ID)
	}
    if c.EN != "Cabo Verde" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cabo Verde", c.EN)
	}
	if c.TH != "กาบูเวร์ดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบูเวร์ดี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(132)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(140)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 140 {
		t.Errorf("invalid value: expect: %d actual: %d", 140, c.ID)
	}
    if c.EN != "Central African Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Central African Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐแอฟริกากลาง" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐแอฟริกากลาง", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(140)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(148)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 148 {
		t.Errorf("invalid value: expect: %d actual: %d", 148, c.ID)
	}
    if c.EN != "Chad" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chad", c.EN)
	}
	if c.TH != "ชาด" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชาด", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(148)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(152)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 152 {
		t.Errorf("invalid value: expect: %d actual: %d", 152, c.ID)
	}
    if c.EN != "Chile" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chile", c.EN)
	}
	if c.TH != "ชิลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชิลี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(152)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(156)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 156 {
		t.Errorf("invalid value: expect: %d actual: %d", 156, c.ID)
	}
    if c.EN != "China" {
		t.Errorf("invalid value: expect: %s actual: %s", "China", c.EN)
	}
	if c.TH != "จีน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จีน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(156)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(170)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 170 {
		t.Errorf("invalid value: expect: %d actual: %d", 170, c.ID)
	}
    if c.EN != "Colombia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Colombia", c.EN)
	}
	if c.TH != "โคลอมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โคลอมเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(170)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(174)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 174 {
		t.Errorf("invalid value: expect: %d actual: %d", 174, c.ID)
	}
    if c.EN != "Comoros" {
		t.Errorf("invalid value: expect: %s actual: %s", "Comoros", c.EN)
	}
	if c.TH != "คอโมโรส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอโมโรส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(174)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(178)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 178 {
		t.Errorf("invalid value: expect: %d actual: %d", 178, c.ID)
	}
    if c.EN != "Congo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo", c.EN)
	}
	if c.TH != "สาธารณรัฐคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐคองโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(178)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(180)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 180 {
		t.Errorf("invalid value: expect: %d actual: %d", 180, c.ID)
	}
    if c.EN != "Congo, Democratic Republic of the" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo, Democratic Republic of the", c.EN)
	}
	if c.TH != "สาธารณรัฐประชาธิปไตยคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐประชาธิปไตยคองโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(180)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(188)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 188 {
		t.Errorf("invalid value: expect: %d actual: %d", 188, c.ID)
	}
    if c.EN != "Costa Rica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Costa Rica", c.EN)
	}
	if c.TH != "คอสตาริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอสตาริกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(188)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(384)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 384 {
		t.Errorf("invalid value: expect: %d actual: %d", 384, c.ID)
	}
    if c.EN != "Côte d'Ivoire" {
		t.Errorf("invalid value: expect: %s actual: %s", "Côte d'Ivoire", c.EN)
	}
	if c.TH != "โกตดิวัวร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โกตดิวัวร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(384)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(191)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 191 {
		t.Errorf("invalid value: expect: %d actual: %d", 191, c.ID)
	}
    if c.EN != "Croatia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Croatia", c.EN)
	}
	if c.TH != "โครเอเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โครเอเชีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(191)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(192)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 192 {
		t.Errorf("invalid value: expect: %d actual: %d", 192, c.ID)
	}
    if c.EN != "Cuba" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cuba", c.EN)
	}
	if c.TH != "คิวบา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิวบา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(192)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(196)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 196 {
		t.Errorf("invalid value: expect: %d actual: %d", 196, c.ID)
	}
    if c.EN != "Cyprus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cyprus", c.EN)
	}
	if c.TH != "ไซปรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไซปรัส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(196)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(203)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 203 {
		t.Errorf("invalid value: expect: %d actual: %d", 203, c.ID)
	}
    if c.EN != "Czechia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Czechia", c.EN)
	}
	if c.TH != "เช็กเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เช็กเกีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(203)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(208)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 208 {
		t.Errorf("invalid value: expect: %d actual: %d", 208, c.ID)
	}
    if c.EN != "Denmark" {
		t.Errorf("invalid value: expect: %s actual: %s", "Denmark", c.EN)
	}
	if c.TH != "เดนมาร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เดนมาร์ก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(208)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(262)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 262 {
		t.Errorf("invalid value: expect: %d actual: %d", 262, c.ID)
	}
    if c.EN != "Djibouti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Djibouti", c.EN)
	}
	if c.TH != "จิบูตี" {
		t.Errorf("invalid value: expect: %s actual: %s", "จิบูตี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(262)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(212)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 212 {
		t.Errorf("invalid value: expect: %d actual: %d", 212, c.ID)
	}
    if c.EN != "Dominica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominica", c.EN)
	}
	if c.TH != "ดอมินีกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ดอมินีกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(212)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(214)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 214 {
		t.Errorf("invalid value: expect: %d actual: %d", 214, c.ID)
	}
    if c.EN != "Dominican Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominican Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐโดมินิกัน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐโดมินิกัน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(214)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(218)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 218 {
		t.Errorf("invalid value: expect: %d actual: %d", 218, c.ID)
	}
    if c.EN != "Ecuador" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ecuador", c.EN)
	}
	if c.TH != "เอกวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอกวาดอร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(218)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(818)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 818 {
		t.Errorf("invalid value: expect: %d actual: %d", 818, c.ID)
	}
    if c.EN != "Egypt" {
		t.Errorf("invalid value: expect: %s actual: %s", "Egypt", c.EN)
	}
	if c.TH != "อียิปต์" {
		t.Errorf("invalid value: expect: %s actual: %s", "อียิปต์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(818)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(222)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 222 {
		t.Errorf("invalid value: expect: %d actual: %d", 222, c.ID)
	}
    if c.EN != "El Salvador" {
		t.Errorf("invalid value: expect: %s actual: %s", "El Salvador", c.EN)
	}
	if c.TH != "เอลซัลวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอลซัลวาดอร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(222)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(226)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 226 {
		t.Errorf("invalid value: expect: %d actual: %d", 226, c.ID)
	}
    if c.EN != "Equatorial Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Equatorial Guinea", c.EN)
	}
	if c.TH != "อิเควทอเรียลกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิเควทอเรียลกินี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(226)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(232)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 232 {
		t.Errorf("invalid value: expect: %d actual: %d", 232, c.ID)
	}
    if c.EN != "Eritrea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eritrea", c.EN)
	}
	if c.TH != "เอริเทรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอริเทรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(232)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(233)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 233 {
		t.Errorf("invalid value: expect: %d actual: %d", 233, c.ID)
	}
    if c.EN != "Estonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Estonia", c.EN)
	}
	if c.TH != "เอสโตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสโตเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(233)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(231)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 231 {
		t.Errorf("invalid value: expect: %d actual: %d", 231, c.ID)
	}
    if c.EN != "Ethiopia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ethiopia", c.EN)
	}
	if c.TH != "เอธิโอเปีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอธิโอเปีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(231)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(242)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 242 {
		t.Errorf("invalid value: expect: %d actual: %d", 242, c.ID)
	}
    if c.EN != "Fiji" {
		t.Errorf("invalid value: expect: %s actual: %s", "Fiji", c.EN)
	}
	if c.TH != "ฟีจี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟีจี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(242)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(246)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 246 {
		t.Errorf("invalid value: expect: %d actual: %d", 246, c.ID)
	}
    if c.EN != "Finland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Finland", c.EN)
	}
	if c.TH != "ฟินแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟินแลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(246)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(250)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 250 {
		t.Errorf("invalid value: expect: %d actual: %d", 250, c.ID)
	}
    if c.EN != "France" {
		t.Errorf("invalid value: expect: %s actual: %s", "France", c.EN)
	}
	if c.TH != "ฝรั่งเศส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฝรั่งเศส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(250)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(266)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 266 {
		t.Errorf("invalid value: expect: %d actual: %d", 266, c.ID)
	}
    if c.EN != "Gabon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gabon", c.EN)
	}
	if c.TH != "กาบอง" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบอง", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(266)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(270)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 270 {
		t.Errorf("invalid value: expect: %d actual: %d", 270, c.ID)
	}
    if c.EN != "Gambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gambia", c.EN)
	}
	if c.TH != "แกมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แกมเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(270)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(268)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 268 {
		t.Errorf("invalid value: expect: %d actual: %d", 268, c.ID)
	}
    if c.EN != "Georgia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Georgia", c.EN)
	}
	if c.TH != "จอร์เจีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์เจีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(268)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(276)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 276 {
		t.Errorf("invalid value: expect: %d actual: %d", 276, c.ID)
	}
    if c.EN != "Germany" {
		t.Errorf("invalid value: expect: %s actual: %s", "Germany", c.EN)
	}
	if c.TH != "เยอรมนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยอรมนี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(276)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(288)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 288 {
		t.Errorf("invalid value: expect: %d actual: %d", 288, c.ID)
	}
    if c.EN != "Ghana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ghana", c.EN)
	}
	if c.TH != "กานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กานา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(288)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(300)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 300 {
		t.Errorf("invalid value: expect: %d actual: %d", 300, c.ID)
	}
    if c.EN != "Greece" {
		t.Errorf("invalid value: expect: %s actual: %s", "Greece", c.EN)
	}
	if c.TH != "กรีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "กรีซ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(300)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(308)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 308 {
		t.Errorf("invalid value: expect: %d actual: %d", 308, c.ID)
	}
    if c.EN != "Grenada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Grenada", c.EN)
	}
	if c.TH != "เกรเนดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกรเนดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(308)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(320)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 320 {
		t.Errorf("invalid value: expect: %d actual: %d", 320, c.ID)
	}
    if c.EN != "Guatemala" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guatemala", c.EN)
	}
	if c.TH != "กัวเตมาลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัวเตมาลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(320)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(324)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 324 {
		t.Errorf("invalid value: expect: %d actual: %d", 324, c.ID)
	}
    if c.EN != "Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea", c.EN)
	}
	if c.TH != "กินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(324)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(624)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 624 {
		t.Errorf("invalid value: expect: %d actual: %d", 624, c.ID)
	}
    if c.EN != "Guinea-Bissau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea-Bissau", c.EN)
	}
	if c.TH != "กินี-บิสเซา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี-บิสเซา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(624)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(328)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 328 {
		t.Errorf("invalid value: expect: %d actual: %d", 328, c.ID)
	}
    if c.EN != "Guyana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guyana", c.EN)
	}
	if c.TH != "กายอานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กายอานา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(328)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(332)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 332 {
		t.Errorf("invalid value: expect: %d actual: %d", 332, c.ID)
	}
    if c.EN != "Haiti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Haiti", c.EN)
	}
	if c.TH != "เฮติ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เฮติ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(332)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(340)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 340 {
		t.Errorf("invalid value: expect: %d actual: %d", 340, c.ID)
	}
    if c.EN != "Honduras" {
		t.Errorf("invalid value: expect: %s actual: %s", "Honduras", c.EN)
	}
	if c.TH != "ฮอนดูรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮอนดูรัส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(340)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(348)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 348 {
		t.Errorf("invalid value: expect: %d actual: %d", 348, c.ID)
	}
    if c.EN != "Hungary" {
		t.Errorf("invalid value: expect: %s actual: %s", "Hungary", c.EN)
	}
	if c.TH != "ฮังการี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮังการี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(348)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(352)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 352 {
		t.Errorf("invalid value: expect: %d actual: %d", 352, c.ID)
	}
    if c.EN != "Iceland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iceland", c.EN)
	}
	if c.TH != "ไอซ์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอซ์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(352)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(356)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 356 {
		t.Errorf("invalid value: expect: %d actual: %d", 356, c.ID)
	}
    if c.EN != "India" {
		t.Errorf("invalid value: expect: %s actual: %s", "India", c.EN)
	}
	if c.TH != "อินเดีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินเดีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(356)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(360)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 360 {
		t.Errorf("invalid value: expect: %d actual: %d", 360, c.ID)
	}
    if c.EN != "Indonesia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Indonesia", c.EN)
	}
	if c.TH != "อินโดนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินโดนีเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(360)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(364)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 364 {
		t.Errorf("invalid value: expect: %d actual: %d", 364, c.ID)
	}
    if c.EN != "Iran (Islamic Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iran (Islamic Republic of)", c.EN)
	}
	if c.TH != "อิหร่าน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิหร่าน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(364)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(368)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 368 {
		t.Errorf("invalid value: expect: %d actual: %d", 368, c.ID)
	}
    if c.EN != "Iraq" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iraq", c.EN)
	}
	if c.TH != "อิรัก" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิรัก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(368)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(372)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 372 {
		t.Errorf("invalid value: expect: %d actual: %d", 372, c.ID)
	}
    if c.EN != "Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ireland", c.EN)
	}
	if c.TH != "ไอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอร์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(372)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(376)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 376 {
		t.Errorf("invalid value: expect: %d actual: %d", 376, c.ID)
	}
    if c.EN != "Israel" {
		t.Errorf("invalid value: expect: %s actual: %s", "Israel", c.EN)
	}
	if c.TH != "อิสราเอล" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิสราเอล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(376)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(380)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 380 {
		t.Errorf("invalid value: expect: %d actual: %d", 380, c.ID)
	}
    if c.EN != "Italy" {
		t.Errorf("invalid value: expect: %s actual: %s", "Italy", c.EN)
	}
	if c.TH != "อิตาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิตาลี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(380)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(388)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 388 {
		t.Errorf("invalid value: expect: %d actual: %d", 388, c.ID)
	}
    if c.EN != "Jamaica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jamaica", c.EN)
	}
	if c.TH != "จาเมกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "จาเมกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(388)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(392)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 392 {
		t.Errorf("invalid value: expect: %d actual: %d", 392, c.ID)
	}
    if c.EN != "Japan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Japan", c.EN)
	}
	if c.TH != "ญี่ปุ่น" {
		t.Errorf("invalid value: expect: %s actual: %s", "ญี่ปุ่น", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(392)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(400)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 400 {
		t.Errorf("invalid value: expect: %d actual: %d", 400, c.ID)
	}
    if c.EN != "Jordan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jordan", c.EN)
	}
	if c.TH != "จอร์แดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์แดน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(400)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(398)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 398 {
		t.Errorf("invalid value: expect: %d actual: %d", 398, c.ID)
	}
    if c.EN != "Kazakhstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kazakhstan", c.EN)
	}
	if c.TH != "คาซัคสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คาซัคสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(398)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(404)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 404 {
		t.Errorf("invalid value: expect: %d actual: %d", 404, c.ID)
	}
    if c.EN != "Kenya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kenya", c.EN)
	}
	if c.TH != "เคนยา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เคนยา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(404)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(296)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 296 {
		t.Errorf("invalid value: expect: %d actual: %d", 296, c.ID)
	}
    if c.EN != "Kiribati" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kiribati", c.EN)
	}
	if c.TH != "คิริบาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิริบาส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(296)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(408)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 408 {
		t.Errorf("invalid value: expect: %d actual: %d", 408, c.ID)
	}
    if c.EN != "Korea (Democratic People's Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea (Democratic People's Republic of)", c.EN)
	}
	if c.TH != "เกาหลีเหนือ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีเหนือ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(408)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(410)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 410 {
		t.Errorf("invalid value: expect: %d actual: %d", 410, c.ID)
	}
    if c.EN != "Korea, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea, Republic of", c.EN)
	}
	if c.TH != "เกาหลีใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีใต้", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(410)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(414)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 414 {
		t.Errorf("invalid value: expect: %d actual: %d", 414, c.ID)
	}
    if c.EN != "Kuwait" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kuwait", c.EN)
	}
	if c.TH != "คูเวต" {
		t.Errorf("invalid value: expect: %s actual: %s", "คูเวต", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(414)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(417)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 417 {
		t.Errorf("invalid value: expect: %d actual: %d", 417, c.ID)
	}
    if c.EN != "Kyrgyzstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kyrgyzstan", c.EN)
	}
	if c.TH != "คีร์กีซสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คีร์กีซสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(417)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(418)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 418 {
		t.Errorf("invalid value: expect: %d actual: %d", 418, c.ID)
	}
    if c.EN != "Lao People's Democratic Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lao People's Democratic Republic", c.EN)
	}
	if c.TH != "ลาว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลาว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(418)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(428)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 428 {
		t.Errorf("invalid value: expect: %d actual: %d", 428, c.ID)
	}
    if c.EN != "Latvia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Latvia", c.EN)
	}
	if c.TH != "ลัตเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลัตเวีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(428)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(422)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 422 {
		t.Errorf("invalid value: expect: %d actual: %d", 422, c.ID)
	}
    if c.EN != "Lebanon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lebanon", c.EN)
	}
	if c.TH != "เลบานอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลบานอน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(422)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(426)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 426 {
		t.Errorf("invalid value: expect: %d actual: %d", 426, c.ID)
	}
    if c.EN != "Lesotho" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lesotho", c.EN)
	}
	if c.TH != "เลโซโท" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลโซโท", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(426)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(430)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 430 {
		t.Errorf("invalid value: expect: %d actual: %d", 430, c.ID)
	}
    if c.EN != "Liberia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liberia", c.EN)
	}
	if c.TH != "ไลบีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไลบีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(430)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(434)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 434 {
		t.Errorf("invalid value: expect: %d actual: %d", 434, c.ID)
	}
    if c.EN != "Libya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Libya", c.EN)
	}
	if c.TH != "ลิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(434)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(438)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 438 {
		t.Errorf("invalid value: expect: %d actual: %d", 438, c.ID)
	}
    if c.EN != "Liechtenstein" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liechtenstein", c.EN)
	}
	if c.TH != "ลิกเตนสไตน์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิกเตนสไตน์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(438)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(440)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 440 {
		t.Errorf("invalid value: expect: %d actual: %d", 440, c.ID)
	}
    if c.EN != "Lithuania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lithuania", c.EN)
	}
	if c.TH != "ลิทัวเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิทัวเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(440)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(442)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 442 {
		t.Errorf("invalid value: expect: %d actual: %d", 442, c.ID)
	}
    if c.EN != "Luxembourg" {
		t.Errorf("invalid value: expect: %s actual: %s", "Luxembourg", c.EN)
	}
	if c.TH != "ลักเซมเบิร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลักเซมเบิร์ก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(442)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(807)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 807 {
		t.Errorf("invalid value: expect: %d actual: %d", 807, c.ID)
	}
    if c.EN != "North Macedonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "North Macedonia", c.EN)
	}
	if c.TH != "นอร์ทมาซิโดเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์ทมาซิโดเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(807)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(450)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 450 {
		t.Errorf("invalid value: expect: %d actual: %d", 450, c.ID)
	}
    if c.EN != "Madagascar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Madagascar", c.EN)
	}
	if c.TH != "มาดากัสการ์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาดากัสการ์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(450)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(454)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 454 {
		t.Errorf("invalid value: expect: %d actual: %d", 454, c.ID)
	}
    if c.EN != "Malawi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malawi", c.EN)
	}
	if c.TH != "มาลาวี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลาวี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(454)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(458)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 458 {
		t.Errorf("invalid value: expect: %d actual: %d", 458, c.ID)
	}
    if c.EN != "Malaysia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malaysia", c.EN)
	}
	if c.TH != "มาเลเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาเลเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(458)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(462)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 462 {
		t.Errorf("invalid value: expect: %d actual: %d", 462, c.ID)
	}
    if c.EN != "Maldives" {
		t.Errorf("invalid value: expect: %s actual: %s", "Maldives", c.EN)
	}
	if c.TH != "มัลดีฟส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มัลดีฟส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(462)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(466)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 466 {
		t.Errorf("invalid value: expect: %d actual: %d", 466, c.ID)
	}
    if c.EN != "Mali" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mali", c.EN)
	}
	if c.TH != "มาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(466)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(470)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 470 {
		t.Errorf("invalid value: expect: %d actual: %d", 470, c.ID)
	}
    if c.EN != "Malta" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malta", c.EN)
	}
	if c.TH != "มอลตา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลตา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(470)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(584)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 584 {
		t.Errorf("invalid value: expect: %d actual: %d", 584, c.ID)
	}
    if c.EN != "Marshall Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Marshall Islands", c.EN)
	}
	if c.TH != "หมู่เกาะมาร์แชลล์" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะมาร์แชลล์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(584)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(478)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 478 {
		t.Errorf("invalid value: expect: %d actual: %d", 478, c.ID)
	}
    if c.EN != "Mauritania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritania", c.EN)
	}
	if c.TH != "มอริเตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเตเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(478)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(480)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 480 {
		t.Errorf("invalid value: expect: %d actual: %d", 480, c.ID)
	}
    if c.EN != "Mauritius" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritius", c.EN)
	}
	if c.TH != "มอริเชียส" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเชียส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(480)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(484)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 484 {
		t.Errorf("invalid value: expect: %d actual: %d", 484, c.ID)
	}
    if c.EN != "Mexico" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mexico", c.EN)
	}
	if c.TH != "เม็กซิโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เม็กซิโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(484)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(583)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 583 {
		t.Errorf("invalid value: expect: %d actual: %d", 583, c.ID)
	}
    if c.EN != "Micronesia (Federated States of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Micronesia (Federated States of)", c.EN)
	}
	if c.TH != "ไมโครนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไมโครนีเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(583)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(504)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 504 {
		t.Errorf("invalid value: expect: %d actual: %d", 504, c.ID)
	}
    if c.EN != "Morocco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Morocco", c.EN)
	}
	if c.TH != "โมร็อกโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมร็อกโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(504)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(498)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 498 {
		t.Errorf("invalid value: expect: %d actual: %d", 498, c.ID)
	}
    if c.EN != "Moldova, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Moldova, Republic of", c.EN)
	}
	if c.TH != "มอลโดวา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลโดวา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(498)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(492)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 492 {
		t.Errorf("invalid value: expect: %d actual: %d", 492, c.ID)
	}
    if c.EN != "Monaco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Monaco", c.EN)
	}
	if c.TH != "โมนาโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมนาโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(492)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(496)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 496 {
		t.Errorf("invalid value: expect: %d actual: %d", 496, c.ID)
	}
    if c.EN != "Mongolia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mongolia", c.EN)
	}
	if c.TH != "มองโกเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มองโกเลีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(496)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(499)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 499 {
		t.Errorf("invalid value: expect: %d actual: %d", 499, c.ID)
	}
    if c.EN != "Montenegro" {
		t.Errorf("invalid value: expect: %s actual: %s", "Montenegro", c.EN)
	}
	if c.TH != "มอนเตเนโกร" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอนเตเนโกร", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(499)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(508)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 508 {
		t.Errorf("invalid value: expect: %d actual: %d", 508, c.ID)
	}
    if c.EN != "Mozambique" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mozambique", c.EN)
	}
	if c.TH != "โมซัมบิก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมซัมบิก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(508)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(104)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 104 {
		t.Errorf("invalid value: expect: %d actual: %d", 104, c.ID)
	}
    if c.EN != "Myanmar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Myanmar", c.EN)
	}
	if c.TH != "พม่า" {
		t.Errorf("invalid value: expect: %s actual: %s", "พม่า", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(104)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(516)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 516 {
		t.Errorf("invalid value: expect: %d actual: %d", 516, c.ID)
	}
    if c.EN != "Namibia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Namibia", c.EN)
	}
	if c.TH != "นามิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นามิเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(516)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(520)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 520 {
		t.Errorf("invalid value: expect: %d actual: %d", 520, c.ID)
	}
    if c.EN != "Nauru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nauru", c.EN)
	}
	if c.TH != "นาอูรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "นาอูรู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(520)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(524)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 524 {
		t.Errorf("invalid value: expect: %d actual: %d", 524, c.ID)
	}
    if c.EN != "Nepal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nepal", c.EN)
	}
	if c.TH != "เนปาล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนปาล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(524)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(528)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 528 {
		t.Errorf("invalid value: expect: %d actual: %d", 528, c.ID)
	}
    if c.EN != "Netherlands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Netherlands", c.EN)
	}
	if c.TH != "เนเธอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนเธอร์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(528)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(554)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 554 {
		t.Errorf("invalid value: expect: %d actual: %d", 554, c.ID)
	}
    if c.EN != "New Zealand" {
		t.Errorf("invalid value: expect: %s actual: %s", "New Zealand", c.EN)
	}
	if c.TH != "นิวซีแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิวซีแลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(554)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(558)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 558 {
		t.Errorf("invalid value: expect: %d actual: %d", 558, c.ID)
	}
    if c.EN != "Nicaragua" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nicaragua", c.EN)
	}
	if c.TH != "นิการากัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิการากัว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(558)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(562)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 562 {
		t.Errorf("invalid value: expect: %d actual: %d", 562, c.ID)
	}
    if c.EN != "Niger" {
		t.Errorf("invalid value: expect: %s actual: %s", "Niger", c.EN)
	}
	if c.TH != "ไนเจอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนเจอร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(562)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(566)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 566 {
		t.Errorf("invalid value: expect: %d actual: %d", 566, c.ID)
	}
    if c.EN != "Nigeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nigeria", c.EN)
	}
	if c.TH != "ไนจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนจีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(566)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(578)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 578 {
		t.Errorf("invalid value: expect: %d actual: %d", 578, c.ID)
	}
    if c.EN != "Norway" {
		t.Errorf("invalid value: expect: %s actual: %s", "Norway", c.EN)
	}
	if c.TH != "นอร์เวย์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์เวย์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(578)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(512)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 512 {
		t.Errorf("invalid value: expect: %d actual: %d", 512, c.ID)
	}
    if c.EN != "Oman" {
		t.Errorf("invalid value: expect: %s actual: %s", "Oman", c.EN)
	}
	if c.TH != "โอมาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "โอมาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(512)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(586)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 586 {
		t.Errorf("invalid value: expect: %d actual: %d", 586, c.ID)
	}
    if c.EN != "Pakistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Pakistan", c.EN)
	}
	if c.TH != "ปากีสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปากีสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(586)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(585)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 585 {
		t.Errorf("invalid value: expect: %d actual: %d", 585, c.ID)
	}
    if c.EN != "Palau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Palau", c.EN)
	}
	if c.TH != "ปาเลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาเลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(585)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(591)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 591 {
		t.Errorf("invalid value: expect: %d actual: %d", 591, c.ID)
	}
    if c.EN != "Panama" {
		t.Errorf("invalid value: expect: %s actual: %s", "Panama", c.EN)
	}
	if c.TH != "ปานามา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปานามา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(591)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(598)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 598 {
		t.Errorf("invalid value: expect: %d actual: %d", 598, c.ID)
	}
    if c.EN != "Papua New Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Papua New Guinea", c.EN)
	}
	if c.TH != "ปาปัวนิวกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาปัวนิวกินี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(598)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(600)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 600 {
		t.Errorf("invalid value: expect: %d actual: %d", 600, c.ID)
	}
    if c.EN != "Paraguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Paraguay", c.EN)
	}
	if c.TH != "ปารากวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปารากวัย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(600)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(604)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 604 {
		t.Errorf("invalid value: expect: %d actual: %d", 604, c.ID)
	}
    if c.EN != "Peru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Peru", c.EN)
	}
	if c.TH != "เปรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "เปรู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(604)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(608)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 608 {
		t.Errorf("invalid value: expect: %d actual: %d", 608, c.ID)
	}
    if c.EN != "Philippines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Philippines", c.EN)
	}
	if c.TH != "ฟิลิปปินส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟิลิปปินส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(608)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(616)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 616 {
		t.Errorf("invalid value: expect: %d actual: %d", 616, c.ID)
	}
    if c.EN != "Poland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Poland", c.EN)
	}
	if c.TH != "โปแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปแลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(616)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(620)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 620 {
		t.Errorf("invalid value: expect: %d actual: %d", 620, c.ID)
	}
    if c.EN != "Portugal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Portugal", c.EN)
	}
	if c.TH != "โปรตุเกส" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปรตุเกส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(620)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(634)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 634 {
		t.Errorf("invalid value: expect: %d actual: %d", 634, c.ID)
	}
    if c.EN != "Qatar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Qatar", c.EN)
	}
	if c.TH != "กาตาร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาตาร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(634)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(642)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 642 {
		t.Errorf("invalid value: expect: %d actual: %d", 642, c.ID)
	}
    if c.EN != "Romania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Romania", c.EN)
	}
	if c.TH != "โรมาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โรมาเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(642)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(643)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 643 {
		t.Errorf("invalid value: expect: %d actual: %d", 643, c.ID)
	}
    if c.EN != "Russian Federation" {
		t.Errorf("invalid value: expect: %s actual: %s", "Russian Federation", c.EN)
	}
	if c.TH != "รัสเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "รัสเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(643)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(646)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 646 {
		t.Errorf("invalid value: expect: %d actual: %d", 646, c.ID)
	}
    if c.EN != "Rwanda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Rwanda", c.EN)
	}
	if c.TH != "รวันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "รวันดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(646)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(659)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 659 {
		t.Errorf("invalid value: expect: %d actual: %d", 659, c.ID)
	}
    if c.EN != "Saint Kitts and Nevis" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Kitts and Nevis", c.EN)
	}
	if c.TH != "เซนต์คิตส์และเนวิส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์คิตส์และเนวิส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(659)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(662)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 662 {
		t.Errorf("invalid value: expect: %d actual: %d", 662, c.ID)
	}
    if c.EN != "Saint Lucia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Lucia", c.EN)
	}
	if c.TH != "เซนต์ลูเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์ลูเชีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(662)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(670)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 670 {
		t.Errorf("invalid value: expect: %d actual: %d", 670, c.ID)
	}
    if c.EN != "Saint Vincent and the Grenadines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Vincent and the Grenadines", c.EN)
	}
	if c.TH != "เซนต์วินเซนต์และเกรนาดีนส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์วินเซนต์และเกรนาดีนส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(670)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(882)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 882 {
		t.Errorf("invalid value: expect: %d actual: %d", 882, c.ID)
	}
    if c.EN != "Samoa" {
		t.Errorf("invalid value: expect: %s actual: %s", "Samoa", c.EN)
	}
	if c.TH != "ซามัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซามัว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(882)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(674)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 674 {
		t.Errorf("invalid value: expect: %d actual: %d", 674, c.ID)
	}
    if c.EN != "San Marino" {
		t.Errorf("invalid value: expect: %s actual: %s", "San Marino", c.EN)
	}
	if c.TH != "ซานมารีโน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซานมารีโน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(674)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(678)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 678 {
		t.Errorf("invalid value: expect: %d actual: %d", 678, c.ID)
	}
    if c.EN != "Sao Tome and Principe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sao Tome and Principe", c.EN)
	}
	if c.TH != "เซาตูเมและปรินซีปี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาตูเมและปรินซีปี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(678)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(682)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 682 {
		t.Errorf("invalid value: expect: %d actual: %d", 682, c.ID)
	}
    if c.EN != "Saudi Arabia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saudi Arabia", c.EN)
	}
	if c.TH != "ซาอุดีอาระเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซาอุดีอาระเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(682)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(686)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 686 {
		t.Errorf("invalid value: expect: %d actual: %d", 686, c.ID)
	}
    if c.EN != "Senegal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Senegal", c.EN)
	}
	if c.TH != "เซเนกัล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเนกัล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(686)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(688)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 688 {
		t.Errorf("invalid value: expect: %d actual: %d", 688, c.ID)
	}
    if c.EN != "Serbia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Serbia", c.EN)
	}
	if c.TH != "เซอร์เบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซอร์เบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(688)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(690)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 690 {
		t.Errorf("invalid value: expect: %d actual: %d", 690, c.ID)
	}
    if c.EN != "Seychelles" {
		t.Errorf("invalid value: expect: %s actual: %s", "Seychelles", c.EN)
	}
	if c.TH != "เซเชลส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเชลส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(690)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(694)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 694 {
		t.Errorf("invalid value: expect: %d actual: %d", 694, c.ID)
	}
    if c.EN != "Sierra Leone" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sierra Leone", c.EN)
	}
	if c.TH != "เซียร์ราลีโอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซียร์ราลีโอน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(694)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(702)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 702 {
		t.Errorf("invalid value: expect: %d actual: %d", 702, c.ID)
	}
    if c.EN != "Singapore" {
		t.Errorf("invalid value: expect: %s actual: %s", "Singapore", c.EN)
	}
	if c.TH != "สิงคโปร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สิงคโปร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(702)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(703)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 703 {
		t.Errorf("invalid value: expect: %d actual: %d", 703, c.ID)
	}
    if c.EN != "Slovakia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovakia", c.EN)
	}
	if c.TH != "สโลวาเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวาเกีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(703)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(705)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 705 {
		t.Errorf("invalid value: expect: %d actual: %d", 705, c.ID)
	}
    if c.EN != "Slovenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovenia", c.EN)
	}
	if c.TH != "สโลวีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวีเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(705)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(90)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 90 {
		t.Errorf("invalid value: expect: %d actual: %d", 90, c.ID)
	}
    if c.EN != "Solomon Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Solomon Islands", c.EN)
	}
	if c.TH != "หมู่เกาะโซโลมอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะโซโลมอน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(90)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(706)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 706 {
		t.Errorf("invalid value: expect: %d actual: %d", 706, c.ID)
	}
    if c.EN != "Somalia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Somalia", c.EN)
	}
	if c.TH != "โซมาเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โซมาเลีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(706)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(710)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 710 {
		t.Errorf("invalid value: expect: %d actual: %d", 710, c.ID)
	}
    if c.EN != "South Africa" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Africa", c.EN)
	}
	if c.TH != "แอฟริกาใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอฟริกาใต้", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(710)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(728)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 728 {
		t.Errorf("invalid value: expect: %d actual: %d", 728, c.ID)
	}
    if c.EN != "South Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Sudan", c.EN)
	}
	if c.TH != "เซาท์ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาท์ซูดาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(728)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(724)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 724 {
		t.Errorf("invalid value: expect: %d actual: %d", 724, c.ID)
	}
    if c.EN != "Spain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Spain", c.EN)
	}
	if c.TH != "สเปน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สเปน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(724)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(144)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 144 {
		t.Errorf("invalid value: expect: %d actual: %d", 144, c.ID)
	}
    if c.EN != "Sri Lanka" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sri Lanka", c.EN)
	}
	if c.TH != "ศรีลังกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ศรีลังกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(144)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(729)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 729 {
		t.Errorf("invalid value: expect: %d actual: %d", 729, c.ID)
	}
    if c.EN != "Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sudan", c.EN)
	}
	if c.TH != "ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูดาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(729)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(740)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 740 {
		t.Errorf("invalid value: expect: %d actual: %d", 740, c.ID)
	}
    if c.EN != "Suriname" {
		t.Errorf("invalid value: expect: %s actual: %s", "Suriname", c.EN)
	}
	if c.TH != "ซูรินาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูรินาม", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(740)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(748)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 748 {
		t.Errorf("invalid value: expect: %d actual: %d", 748, c.ID)
	}
    if c.EN != "Eswatini" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eswatini", c.EN)
	}
	if c.TH != "เอสวาตีนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสวาตีนี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(748)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(752)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 752 {
		t.Errorf("invalid value: expect: %d actual: %d", 752, c.ID)
	}
    if c.EN != "Sweden" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sweden", c.EN)
	}
	if c.TH != "สวีเดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวีเดน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(752)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(756)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 756 {
		t.Errorf("invalid value: expect: %d actual: %d", 756, c.ID)
	}
    if c.EN != "Switzerland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Switzerland", c.EN)
	}
	if c.TH != "สวิตเซอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวิตเซอร์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(756)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(760)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 760 {
		t.Errorf("invalid value: expect: %d actual: %d", 760, c.ID)
	}
    if c.EN != "Syrian Arab Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Syrian Arab Republic", c.EN)
	}
	if c.TH != "ซีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(760)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(762)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 762 {
		t.Errorf("invalid value: expect: %d actual: %d", 762, c.ID)
	}
    if c.EN != "Tajikistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tajikistan", c.EN)
	}
	if c.TH != "ทาจิกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ทาจิกิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(762)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(834)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 834 {
		t.Errorf("invalid value: expect: %d actual: %d", 834, c.ID)
	}
    if c.EN != "Tanzania, United Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tanzania, United Republic of", c.EN)
	}
	if c.TH != "แทนซาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แทนซาเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(834)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(764)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 764 {
		t.Errorf("invalid value: expect: %d actual: %d", 764, c.ID)
	}
    if c.EN != "Thailand" {
		t.Errorf("invalid value: expect: %s actual: %s", "Thailand", c.EN)
	}
	if c.TH != "ไทย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไทย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(764)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(626)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 626 {
		t.Errorf("invalid value: expect: %d actual: %d", 626, c.ID)
	}
    if c.EN != "Timor-Leste" {
		t.Errorf("invalid value: expect: %s actual: %s", "Timor-Leste", c.EN)
	}
	if c.TH != "ติมอร์-เลสเต" {
		t.Errorf("invalid value: expect: %s actual: %s", "ติมอร์-เลสเต", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(626)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(768)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 768 {
		t.Errorf("invalid value: expect: %d actual: %d", 768, c.ID)
	}
    if c.EN != "Togo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Togo", c.EN)
	}
	if c.TH != "โตโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โตโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(768)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(776)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 776 {
		t.Errorf("invalid value: expect: %d actual: %d", 776, c.ID)
	}
    if c.EN != "Tonga" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tonga", c.EN)
	}
	if c.TH != "ตองงา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตองงา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(776)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(780)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 780 {
		t.Errorf("invalid value: expect: %d actual: %d", 780, c.ID)
	}
    if c.EN != "Trinidad and Tobago" {
		t.Errorf("invalid value: expect: %s actual: %s", "Trinidad and Tobago", c.EN)
	}
	if c.TH != "ตรินิแดดและโตเบโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตรินิแดดและโตเบโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(780)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(788)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 788 {
		t.Errorf("invalid value: expect: %d actual: %d", 788, c.ID)
	}
    if c.EN != "Tunisia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tunisia", c.EN)
	}
	if c.TH != "ตูนิเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูนิเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(788)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(792)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 792 {
		t.Errorf("invalid value: expect: %d actual: %d", 792, c.ID)
	}
    if c.EN != "Turkey" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkey", c.EN)
	}
	if c.TH != "ตุรกี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตุรกี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(792)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(795)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 795 {
		t.Errorf("invalid value: expect: %d actual: %d", 795, c.ID)
	}
    if c.EN != "Turkmenistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkmenistan", c.EN)
	}
	if c.TH != "เติร์กเมนิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เติร์กเมนิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(795)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(798)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 798 {
		t.Errorf("invalid value: expect: %d actual: %d", 798, c.ID)
	}
    if c.EN != "Tuvalu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tuvalu", c.EN)
	}
	if c.TH != "ตูวาลู" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูวาลู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(798)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(800)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 800 {
		t.Errorf("invalid value: expect: %d actual: %d", 800, c.ID)
	}
    if c.EN != "Uganda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uganda", c.EN)
	}
	if c.TH != "ยูกันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูกันดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(800)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(804)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 804 {
		t.Errorf("invalid value: expect: %d actual: %d", 804, c.ID)
	}
    if c.EN != "Ukraine" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ukraine", c.EN)
	}
	if c.TH != "ยูเครน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูเครน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(804)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(784)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 784 {
		t.Errorf("invalid value: expect: %d actual: %d", 784, c.ID)
	}
    if c.EN != "United Arab Emirates" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Arab Emirates", c.EN)
	}
	if c.TH != "สหรัฐอาหรับเอมิเรตส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอาหรับเอมิเรตส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(784)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(826)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 826 {
		t.Errorf("invalid value: expect: %d actual: %d", 826, c.ID)
	}
    if c.EN != "United Kingdom of Great Britain and Northern Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Kingdom of Great Britain and Northern Ireland", c.EN)
	}
	if c.TH != "สหราชอาณาจักร" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหราชอาณาจักร", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(826)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(840)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 840 {
		t.Errorf("invalid value: expect: %d actual: %d", 840, c.ID)
	}
    if c.EN != "United States of America" {
		t.Errorf("invalid value: expect: %s actual: %s", "United States of America", c.EN)
	}
	if c.TH != "สหรัฐอเมริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอเมริกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(840)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(858)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 858 {
		t.Errorf("invalid value: expect: %d actual: %d", 858, c.ID)
	}
    if c.EN != "Uruguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uruguay", c.EN)
	}
	if c.TH != "อุรุกวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุรุกวัย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(858)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(860)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 860 {
		t.Errorf("invalid value: expect: %d actual: %d", 860, c.ID)
	}
    if c.EN != "Uzbekistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uzbekistan", c.EN)
	}
	if c.TH != "อุซเบกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุซเบกิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(860)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(548)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 548 {
		t.Errorf("invalid value: expect: %d actual: %d", 548, c.ID)
	}
    if c.EN != "Vanuatu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Vanuatu", c.EN)
	}
	if c.TH != "วานูอาตู" {
		t.Errorf("invalid value: expect: %s actual: %s", "วานูอาตู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(548)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(862)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 862 {
		t.Errorf("invalid value: expect: %d actual: %d", 862, c.ID)
	}
    if c.EN != "Venezuela (Bolivarian Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Venezuela (Bolivarian Republic of)", c.EN)
	}
	if c.TH != "เวเนซุเอลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวเนซุเอลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(862)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(704)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 704 {
		t.Errorf("invalid value: expect: %d actual: %d", 704, c.ID)
	}
    if c.EN != "Viet Nam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Viet Nam", c.EN)
	}
	if c.TH != "เวียดนาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวียดนาม", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(704)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(887)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 887 {
		t.Errorf("invalid value: expect: %d actual: %d", 887, c.ID)
	}
    if c.EN != "Yemen" {
		t.Errorf("invalid value: expect: %s actual: %s", "Yemen", c.EN)
	}
	if c.TH != "เยเมน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยเมน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(887)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(894)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 894 {
		t.Errorf("invalid value: expect: %d actual: %d", 894, c.ID)
	}
    if c.EN != "Zambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zambia", c.EN)
	}
	if c.TH != "แซมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แซมเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(894)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromID(716)
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 716 {
		t.Errorf("invalid value: expect: %d actual: %d", 716, c.ID)
	}
    if c.EN != "Zimbabwe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zimbabwe", c.EN)
	}
	if c.TH != "ซิมบับเว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซิมบับเว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromID(716)
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}

	c, ok = gocountries.GetCountryFromID(0)
	if ok {
		t.Errorf("invalid value: expect: %t actual: %t", false, ok)
	}
    if c.ID != 0 {
		t.Errorf("invalid value: expect: %d actual: %d", 0, c.ID)
	}
    if c.EN != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.EN)
	}
	if c.TH != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.TH)
	}
}

func TestGetCountryFromAlpha2(t *testing.T) {
    var c,c1 gocountries.Country
	var ok bool

	c, ok = gocountries.GetCountryFromAlpha2("af")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 4 {
		t.Errorf("invalid value: expect: %d actual: %d", 4, c.ID)
	}
    if c.EN != "Afghanistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Afghanistan", c.EN)
	}
	if c.TH != "อัฟกานิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อัฟกานิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("af")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("al")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 8 {
		t.Errorf("invalid value: expect: %d actual: %d", 8, c.ID)
	}
    if c.EN != "Albania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Albania", c.EN)
	}
	if c.TH != "แอลเบเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลเบเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("al")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("dz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 12 {
		t.Errorf("invalid value: expect: %d actual: %d", 12, c.ID)
	}
    if c.EN != "Algeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Algeria", c.EN)
	}
	if c.TH != "แอลจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลจีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("dz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ad")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 20 {
		t.Errorf("invalid value: expect: %d actual: %d", 20, c.ID)
	}
    if c.EN != "Andorra" {
		t.Errorf("invalid value: expect: %s actual: %s", "Andorra", c.EN)
	}
	if c.TH != "อันดอร์รา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อันดอร์รา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ad")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ao")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 24 {
		t.Errorf("invalid value: expect: %d actual: %d", 24, c.ID)
	}
    if c.EN != "Angola" {
		t.Errorf("invalid value: expect: %s actual: %s", "Angola", c.EN)
	}
	if c.TH != "แองโกลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แองโกลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ao")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ag")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 28 {
		t.Errorf("invalid value: expect: %d actual: %d", 28, c.ID)
	}
    if c.EN != "Antigua and Barbuda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Antigua and Barbuda", c.EN)
	}
	if c.TH != "แอนติกาและบาร์บูดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอนติกาและบาร์บูดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ag")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ar")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 32 {
		t.Errorf("invalid value: expect: %d actual: %d", 32, c.ID)
	}
    if c.EN != "Argentina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Argentina", c.EN)
	}
	if c.TH != "อาร์เจนตินา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์เจนตินา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ar")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("am")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 51 {
		t.Errorf("invalid value: expect: %d actual: %d", 51, c.ID)
	}
    if c.EN != "Armenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Armenia", c.EN)
	}
	if c.TH != "อาร์มีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์มีเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("am")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("au")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 36 {
		t.Errorf("invalid value: expect: %d actual: %d", 36, c.ID)
	}
    if c.EN != "Australia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Australia", c.EN)
	}
	if c.TH != "ออสเตรเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรเลีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("au")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("at")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 40 {
		t.Errorf("invalid value: expect: %d actual: %d", 40, c.ID)
	}
    if c.EN != "Austria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Austria", c.EN)
	}
	if c.TH != "ออสเตรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("at")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("az")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 31 {
		t.Errorf("invalid value: expect: %d actual: %d", 31, c.ID)
	}
    if c.EN != "Azerbaijan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Azerbaijan", c.EN)
	}
	if c.TH != "อาเซอร์ไบจาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาเซอร์ไบจาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("az")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bs")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 44 {
		t.Errorf("invalid value: expect: %d actual: %d", 44, c.ID)
	}
    if c.EN != "Bahamas" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahamas", c.EN)
	}
	if c.TH != "บาฮามาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาฮามาส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bs")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bh")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 48 {
		t.Errorf("invalid value: expect: %d actual: %d", 48, c.ID)
	}
    if c.EN != "Bahrain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahrain", c.EN)
	}
	if c.TH != "บาห์เรน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาห์เรน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bh")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 50 {
		t.Errorf("invalid value: expect: %d actual: %d", 50, c.ID)
	}
    if c.EN != "Bangladesh" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bangladesh", c.EN)
	}
	if c.TH != "บังกลาเทศ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บังกลาเทศ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bd")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 52 {
		t.Errorf("invalid value: expect: %d actual: %d", 52, c.ID)
	}
    if c.EN != "Barbados" {
		t.Errorf("invalid value: expect: %s actual: %s", "Barbados", c.EN)
	}
	if c.TH != "บาร์เบโดส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาร์เบโดส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bb")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("by")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 112 {
		t.Errorf("invalid value: expect: %d actual: %d", 112, c.ID)
	}
    if c.EN != "Belarus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belarus", c.EN)
	}
	if c.TH != "เบลารุส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลารุส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("by")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("be")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 56 {
		t.Errorf("invalid value: expect: %d actual: %d", 56, c.ID)
	}
    if c.EN != "Belgium" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belgium", c.EN)
	}
	if c.TH != "เบลเยียม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลเยียม", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("be")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 84 {
		t.Errorf("invalid value: expect: %d actual: %d", 84, c.ID)
	}
    if c.EN != "Belize" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belize", c.EN)
	}
	if c.TH != "เบลีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลีซ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bj")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 204 {
		t.Errorf("invalid value: expect: %d actual: %d", 204, c.ID)
	}
    if c.EN != "Benin" {
		t.Errorf("invalid value: expect: %s actual: %s", "Benin", c.EN)
	}
	if c.TH != "เบนิน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบนิน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bj")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 64 {
		t.Errorf("invalid value: expect: %d actual: %d", 64, c.ID)
	}
    if c.EN != "Bhutan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bhutan", c.EN)
	}
	if c.TH != "ภูฏาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ภูฏาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bt")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bo")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 68 {
		t.Errorf("invalid value: expect: %d actual: %d", 68, c.ID)
	}
    if c.EN != "Bolivia (Plurinational State of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bolivia (Plurinational State of)", c.EN)
	}
	if c.TH != "โบลิเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โบลิเวีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bo")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ba")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 70 {
		t.Errorf("invalid value: expect: %d actual: %d", 70, c.ID)
	}
    if c.EN != "Bosnia and Herzegovina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bosnia and Herzegovina", c.EN)
	}
	if c.TH != "บอสเนียและเฮอร์เซโกวีนา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอสเนียและเฮอร์เซโกวีนา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ba")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 72 {
		t.Errorf("invalid value: expect: %d actual: %d", 72, c.ID)
	}
    if c.EN != "Botswana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Botswana", c.EN)
	}
	if c.TH != "บอตสวานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอตสวานา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bw")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("br")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 76 {
		t.Errorf("invalid value: expect: %d actual: %d", 76, c.ID)
	}
    if c.EN != "Brazil" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brazil", c.EN)
	}
	if c.TH != "บราซิล" {
		t.Errorf("invalid value: expect: %s actual: %s", "บราซิล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("br")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 96 {
		t.Errorf("invalid value: expect: %d actual: %d", 96, c.ID)
	}
    if c.EN != "Brunei Darussalam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brunei Darussalam", c.EN)
	}
	if c.TH != "บรูไน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บรูไน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 100 {
		t.Errorf("invalid value: expect: %d actual: %d", 100, c.ID)
	}
    if c.EN != "Bulgaria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bulgaria", c.EN)
	}
	if c.TH != "บัลแกเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "บัลแกเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bf")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 854 {
		t.Errorf("invalid value: expect: %d actual: %d", 854, c.ID)
	}
    if c.EN != "Burkina Faso" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burkina Faso", c.EN)
	}
	if c.TH != "บูร์กินาฟาโซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บูร์กินาฟาโซ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bf")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("bi")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 108 {
		t.Errorf("invalid value: expect: %d actual: %d", 108, c.ID)
	}
    if c.EN != "Burundi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burundi", c.EN)
	}
	if c.TH != "บุรุนดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "บุรุนดี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("bi")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("kh")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 116 {
		t.Errorf("invalid value: expect: %d actual: %d", 116, c.ID)
	}
    if c.EN != "Cambodia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cambodia", c.EN)
	}
	if c.TH != "กัมพูชา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัมพูชา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("kh")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 120 {
		t.Errorf("invalid value: expect: %d actual: %d", 120, c.ID)
	}
    if c.EN != "Cameroon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cameroon", c.EN)
	}
	if c.TH != "แคเมอรูน" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคเมอรูน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ca")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 124 {
		t.Errorf("invalid value: expect: %d actual: %d", 124, c.ID)
	}
    if c.EN != "Canada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Canada", c.EN)
	}
	if c.TH != "แคนาดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคนาดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ca")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 132 {
		t.Errorf("invalid value: expect: %d actual: %d", 132, c.ID)
	}
    if c.EN != "Cabo Verde" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cabo Verde", c.EN)
	}
	if c.TH != "กาบูเวร์ดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบูเวร์ดี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cv")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cf")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 140 {
		t.Errorf("invalid value: expect: %d actual: %d", 140, c.ID)
	}
    if c.EN != "Central African Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Central African Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐแอฟริกากลาง" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐแอฟริกากลาง", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cf")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("td")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 148 {
		t.Errorf("invalid value: expect: %d actual: %d", 148, c.ID)
	}
    if c.EN != "Chad" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chad", c.EN)
	}
	if c.TH != "ชาด" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชาด", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("td")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 152 {
		t.Errorf("invalid value: expect: %d actual: %d", 152, c.ID)
	}
    if c.EN != "Chile" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chile", c.EN)
	}
	if c.TH != "ชิลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชิลี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cl")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 156 {
		t.Errorf("invalid value: expect: %d actual: %d", 156, c.ID)
	}
    if c.EN != "China" {
		t.Errorf("invalid value: expect: %s actual: %s", "China", c.EN)
	}
	if c.TH != "จีน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จีน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("co")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 170 {
		t.Errorf("invalid value: expect: %d actual: %d", 170, c.ID)
	}
    if c.EN != "Colombia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Colombia", c.EN)
	}
	if c.TH != "โคลอมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โคลอมเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("co")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("km")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 174 {
		t.Errorf("invalid value: expect: %d actual: %d", 174, c.ID)
	}
    if c.EN != "Comoros" {
		t.Errorf("invalid value: expect: %s actual: %s", "Comoros", c.EN)
	}
	if c.TH != "คอโมโรส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอโมโรส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("km")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 178 {
		t.Errorf("invalid value: expect: %d actual: %d", 178, c.ID)
	}
    if c.EN != "Congo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo", c.EN)
	}
	if c.TH != "สาธารณรัฐคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐคองโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 180 {
		t.Errorf("invalid value: expect: %d actual: %d", 180, c.ID)
	}
    if c.EN != "Congo, Democratic Republic of the" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo, Democratic Republic of the", c.EN)
	}
	if c.TH != "สาธารณรัฐประชาธิปไตยคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐประชาธิปไตยคองโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cd")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 188 {
		t.Errorf("invalid value: expect: %d actual: %d", 188, c.ID)
	}
    if c.EN != "Costa Rica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Costa Rica", c.EN)
	}
	if c.TH != "คอสตาริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอสตาริกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ci")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 384 {
		t.Errorf("invalid value: expect: %d actual: %d", 384, c.ID)
	}
    if c.EN != "Côte d'Ivoire" {
		t.Errorf("invalid value: expect: %s actual: %s", "Côte d'Ivoire", c.EN)
	}
	if c.TH != "โกตดิวัวร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โกตดิวัวร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ci")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("hr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 191 {
		t.Errorf("invalid value: expect: %d actual: %d", 191, c.ID)
	}
    if c.EN != "Croatia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Croatia", c.EN)
	}
	if c.TH != "โครเอเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โครเอเชีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("hr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 192 {
		t.Errorf("invalid value: expect: %d actual: %d", 192, c.ID)
	}
    if c.EN != "Cuba" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cuba", c.EN)
	}
	if c.TH != "คิวบา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิวบา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cu")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 196 {
		t.Errorf("invalid value: expect: %d actual: %d", 196, c.ID)
	}
    if c.EN != "Cyprus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cyprus", c.EN)
	}
	if c.TH != "ไซปรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไซปรัส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cy")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("cz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 203 {
		t.Errorf("invalid value: expect: %d actual: %d", 203, c.ID)
	}
    if c.EN != "Czechia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Czechia", c.EN)
	}
	if c.TH != "เช็กเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เช็กเกีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("cz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("dk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 208 {
		t.Errorf("invalid value: expect: %d actual: %d", 208, c.ID)
	}
    if c.EN != "Denmark" {
		t.Errorf("invalid value: expect: %s actual: %s", "Denmark", c.EN)
	}
	if c.TH != "เดนมาร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เดนมาร์ก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("dk")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("dj")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 262 {
		t.Errorf("invalid value: expect: %d actual: %d", 262, c.ID)
	}
    if c.EN != "Djibouti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Djibouti", c.EN)
	}
	if c.TH != "จิบูตี" {
		t.Errorf("invalid value: expect: %s actual: %s", "จิบูตี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("dj")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("dm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 212 {
		t.Errorf("invalid value: expect: %d actual: %d", 212, c.ID)
	}
    if c.EN != "Dominica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominica", c.EN)
	}
	if c.TH != "ดอมินีกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ดอมินีกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("dm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("do")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 214 {
		t.Errorf("invalid value: expect: %d actual: %d", 214, c.ID)
	}
    if c.EN != "Dominican Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominican Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐโดมินิกัน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐโดมินิกัน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("do")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ec")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 218 {
		t.Errorf("invalid value: expect: %d actual: %d", 218, c.ID)
	}
    if c.EN != "Ecuador" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ecuador", c.EN)
	}
	if c.TH != "เอกวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอกวาดอร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ec")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("eg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 818 {
		t.Errorf("invalid value: expect: %d actual: %d", 818, c.ID)
	}
    if c.EN != "Egypt" {
		t.Errorf("invalid value: expect: %s actual: %s", "Egypt", c.EN)
	}
	if c.TH != "อียิปต์" {
		t.Errorf("invalid value: expect: %s actual: %s", "อียิปต์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("eg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 222 {
		t.Errorf("invalid value: expect: %d actual: %d", 222, c.ID)
	}
    if c.EN != "El Salvador" {
		t.Errorf("invalid value: expect: %s actual: %s", "El Salvador", c.EN)
	}
	if c.TH != "เอลซัลวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอลซัลวาดอร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sv")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gq")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 226 {
		t.Errorf("invalid value: expect: %d actual: %d", 226, c.ID)
	}
    if c.EN != "Equatorial Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Equatorial Guinea", c.EN)
	}
	if c.TH != "อิเควทอเรียลกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิเควทอเรียลกินี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gq")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("er")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 232 {
		t.Errorf("invalid value: expect: %d actual: %d", 232, c.ID)
	}
    if c.EN != "Eritrea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eritrea", c.EN)
	}
	if c.TH != "เอริเทรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอริเทรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("er")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ee")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 233 {
		t.Errorf("invalid value: expect: %d actual: %d", 233, c.ID)
	}
    if c.EN != "Estonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Estonia", c.EN)
	}
	if c.TH != "เอสโตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสโตเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ee")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("et")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 231 {
		t.Errorf("invalid value: expect: %d actual: %d", 231, c.ID)
	}
    if c.EN != "Ethiopia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ethiopia", c.EN)
	}
	if c.TH != "เอธิโอเปีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอธิโอเปีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("et")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("fj")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 242 {
		t.Errorf("invalid value: expect: %d actual: %d", 242, c.ID)
	}
    if c.EN != "Fiji" {
		t.Errorf("invalid value: expect: %s actual: %s", "Fiji", c.EN)
	}
	if c.TH != "ฟีจี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟีจี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("fj")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("fi")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 246 {
		t.Errorf("invalid value: expect: %d actual: %d", 246, c.ID)
	}
    if c.EN != "Finland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Finland", c.EN)
	}
	if c.TH != "ฟินแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟินแลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("fi")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("fr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 250 {
		t.Errorf("invalid value: expect: %d actual: %d", 250, c.ID)
	}
    if c.EN != "France" {
		t.Errorf("invalid value: expect: %s actual: %s", "France", c.EN)
	}
	if c.TH != "ฝรั่งเศส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฝรั่งเศส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("fr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ga")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 266 {
		t.Errorf("invalid value: expect: %d actual: %d", 266, c.ID)
	}
    if c.EN != "Gabon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gabon", c.EN)
	}
	if c.TH != "กาบอง" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบอง", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ga")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 270 {
		t.Errorf("invalid value: expect: %d actual: %d", 270, c.ID)
	}
    if c.EN != "Gambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gambia", c.EN)
	}
	if c.TH != "แกมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แกมเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ge")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 268 {
		t.Errorf("invalid value: expect: %d actual: %d", 268, c.ID)
	}
    if c.EN != "Georgia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Georgia", c.EN)
	}
	if c.TH != "จอร์เจีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์เจีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ge")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("de")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 276 {
		t.Errorf("invalid value: expect: %d actual: %d", 276, c.ID)
	}
    if c.EN != "Germany" {
		t.Errorf("invalid value: expect: %s actual: %s", "Germany", c.EN)
	}
	if c.TH != "เยอรมนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยอรมนี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("de")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gh")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 288 {
		t.Errorf("invalid value: expect: %d actual: %d", 288, c.ID)
	}
    if c.EN != "Ghana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ghana", c.EN)
	}
	if c.TH != "กานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กานา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gh")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 300 {
		t.Errorf("invalid value: expect: %d actual: %d", 300, c.ID)
	}
    if c.EN != "Greece" {
		t.Errorf("invalid value: expect: %s actual: %s", "Greece", c.EN)
	}
	if c.TH != "กรีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "กรีซ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 308 {
		t.Errorf("invalid value: expect: %d actual: %d", 308, c.ID)
	}
    if c.EN != "Grenada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Grenada", c.EN)
	}
	if c.TH != "เกรเนดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกรเนดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gd")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 320 {
		t.Errorf("invalid value: expect: %d actual: %d", 320, c.ID)
	}
    if c.EN != "Guatemala" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guatemala", c.EN)
	}
	if c.TH != "กัวเตมาลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัวเตมาลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gt")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 324 {
		t.Errorf("invalid value: expect: %d actual: %d", 324, c.ID)
	}
    if c.EN != "Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea", c.EN)
	}
	if c.TH != "กินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 624 {
		t.Errorf("invalid value: expect: %d actual: %d", 624, c.ID)
	}
    if c.EN != "Guinea-Bissau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea-Bissau", c.EN)
	}
	if c.TH != "กินี-บิสเซา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี-บิสเซา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gw")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 328 {
		t.Errorf("invalid value: expect: %d actual: %d", 328, c.ID)
	}
    if c.EN != "Guyana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guyana", c.EN)
	}
	if c.TH != "กายอานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กายอานา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gy")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ht")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 332 {
		t.Errorf("invalid value: expect: %d actual: %d", 332, c.ID)
	}
    if c.EN != "Haiti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Haiti", c.EN)
	}
	if c.TH != "เฮติ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เฮติ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ht")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("hn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 340 {
		t.Errorf("invalid value: expect: %d actual: %d", 340, c.ID)
	}
    if c.EN != "Honduras" {
		t.Errorf("invalid value: expect: %s actual: %s", "Honduras", c.EN)
	}
	if c.TH != "ฮอนดูรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮอนดูรัส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("hn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("hu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 348 {
		t.Errorf("invalid value: expect: %d actual: %d", 348, c.ID)
	}
    if c.EN != "Hungary" {
		t.Errorf("invalid value: expect: %s actual: %s", "Hungary", c.EN)
	}
	if c.TH != "ฮังการี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮังการี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("hu")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("is")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 352 {
		t.Errorf("invalid value: expect: %d actual: %d", 352, c.ID)
	}
    if c.EN != "Iceland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iceland", c.EN)
	}
	if c.TH != "ไอซ์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอซ์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("is")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("in")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 356 {
		t.Errorf("invalid value: expect: %d actual: %d", 356, c.ID)
	}
    if c.EN != "India" {
		t.Errorf("invalid value: expect: %s actual: %s", "India", c.EN)
	}
	if c.TH != "อินเดีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินเดีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("in")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("id")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 360 {
		t.Errorf("invalid value: expect: %d actual: %d", 360, c.ID)
	}
    if c.EN != "Indonesia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Indonesia", c.EN)
	}
	if c.TH != "อินโดนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินโดนีเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("id")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ir")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 364 {
		t.Errorf("invalid value: expect: %d actual: %d", 364, c.ID)
	}
    if c.EN != "Iran (Islamic Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iran (Islamic Republic of)", c.EN)
	}
	if c.TH != "อิหร่าน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิหร่าน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ir")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("iq")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 368 {
		t.Errorf("invalid value: expect: %d actual: %d", 368, c.ID)
	}
    if c.EN != "Iraq" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iraq", c.EN)
	}
	if c.TH != "อิรัก" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิรัก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("iq")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ie")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 372 {
		t.Errorf("invalid value: expect: %d actual: %d", 372, c.ID)
	}
    if c.EN != "Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ireland", c.EN)
	}
	if c.TH != "ไอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอร์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ie")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("il")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 376 {
		t.Errorf("invalid value: expect: %d actual: %d", 376, c.ID)
	}
    if c.EN != "Israel" {
		t.Errorf("invalid value: expect: %s actual: %s", "Israel", c.EN)
	}
	if c.TH != "อิสราเอล" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิสราเอล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("il")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("it")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 380 {
		t.Errorf("invalid value: expect: %d actual: %d", 380, c.ID)
	}
    if c.EN != "Italy" {
		t.Errorf("invalid value: expect: %s actual: %s", "Italy", c.EN)
	}
	if c.TH != "อิตาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิตาลี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("it")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("jm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 388 {
		t.Errorf("invalid value: expect: %d actual: %d", 388, c.ID)
	}
    if c.EN != "Jamaica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jamaica", c.EN)
	}
	if c.TH != "จาเมกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "จาเมกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("jm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("jp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 392 {
		t.Errorf("invalid value: expect: %d actual: %d", 392, c.ID)
	}
    if c.EN != "Japan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Japan", c.EN)
	}
	if c.TH != "ญี่ปุ่น" {
		t.Errorf("invalid value: expect: %s actual: %s", "ญี่ปุ่น", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("jp")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("jo")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 400 {
		t.Errorf("invalid value: expect: %d actual: %d", 400, c.ID)
	}
    if c.EN != "Jordan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jordan", c.EN)
	}
	if c.TH != "จอร์แดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์แดน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("jo")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("kz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 398 {
		t.Errorf("invalid value: expect: %d actual: %d", 398, c.ID)
	}
    if c.EN != "Kazakhstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kazakhstan", c.EN)
	}
	if c.TH != "คาซัคสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คาซัคสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("kz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ke")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 404 {
		t.Errorf("invalid value: expect: %d actual: %d", 404, c.ID)
	}
    if c.EN != "Kenya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kenya", c.EN)
	}
	if c.TH != "เคนยา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เคนยา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ke")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ki")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 296 {
		t.Errorf("invalid value: expect: %d actual: %d", 296, c.ID)
	}
    if c.EN != "Kiribati" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kiribati", c.EN)
	}
	if c.TH != "คิริบาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิริบาส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ki")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("kp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 408 {
		t.Errorf("invalid value: expect: %d actual: %d", 408, c.ID)
	}
    if c.EN != "Korea (Democratic People's Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea (Democratic People's Republic of)", c.EN)
	}
	if c.TH != "เกาหลีเหนือ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีเหนือ", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("kp")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("kr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 410 {
		t.Errorf("invalid value: expect: %d actual: %d", 410, c.ID)
	}
    if c.EN != "Korea, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea, Republic of", c.EN)
	}
	if c.TH != "เกาหลีใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีใต้", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("kr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("kw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 414 {
		t.Errorf("invalid value: expect: %d actual: %d", 414, c.ID)
	}
    if c.EN != "Kuwait" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kuwait", c.EN)
	}
	if c.TH != "คูเวต" {
		t.Errorf("invalid value: expect: %s actual: %s", "คูเวต", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("kw")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("kg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 417 {
		t.Errorf("invalid value: expect: %d actual: %d", 417, c.ID)
	}
    if c.EN != "Kyrgyzstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kyrgyzstan", c.EN)
	}
	if c.TH != "คีร์กีซสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คีร์กีซสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("kg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("la")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 418 {
		t.Errorf("invalid value: expect: %d actual: %d", 418, c.ID)
	}
    if c.EN != "Lao People's Democratic Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lao People's Democratic Republic", c.EN)
	}
	if c.TH != "ลาว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลาว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("la")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("lv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 428 {
		t.Errorf("invalid value: expect: %d actual: %d", 428, c.ID)
	}
    if c.EN != "Latvia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Latvia", c.EN)
	}
	if c.TH != "ลัตเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลัตเวีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("lv")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("lb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 422 {
		t.Errorf("invalid value: expect: %d actual: %d", 422, c.ID)
	}
    if c.EN != "Lebanon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lebanon", c.EN)
	}
	if c.TH != "เลบานอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลบานอน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("lb")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ls")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 426 {
		t.Errorf("invalid value: expect: %d actual: %d", 426, c.ID)
	}
    if c.EN != "Lesotho" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lesotho", c.EN)
	}
	if c.TH != "เลโซโท" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลโซโท", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ls")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("lr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 430 {
		t.Errorf("invalid value: expect: %d actual: %d", 430, c.ID)
	}
    if c.EN != "Liberia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liberia", c.EN)
	}
	if c.TH != "ไลบีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไลบีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("lr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ly")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 434 {
		t.Errorf("invalid value: expect: %d actual: %d", 434, c.ID)
	}
    if c.EN != "Libya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Libya", c.EN)
	}
	if c.TH != "ลิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ly")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("li")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 438 {
		t.Errorf("invalid value: expect: %d actual: %d", 438, c.ID)
	}
    if c.EN != "Liechtenstein" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liechtenstein", c.EN)
	}
	if c.TH != "ลิกเตนสไตน์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิกเตนสไตน์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("li")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("lt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 440 {
		t.Errorf("invalid value: expect: %d actual: %d", 440, c.ID)
	}
    if c.EN != "Lithuania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lithuania", c.EN)
	}
	if c.TH != "ลิทัวเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิทัวเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("lt")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("lu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 442 {
		t.Errorf("invalid value: expect: %d actual: %d", 442, c.ID)
	}
    if c.EN != "Luxembourg" {
		t.Errorf("invalid value: expect: %s actual: %s", "Luxembourg", c.EN)
	}
	if c.TH != "ลักเซมเบิร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลักเซมเบิร์ก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("lu")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 807 {
		t.Errorf("invalid value: expect: %d actual: %d", 807, c.ID)
	}
    if c.EN != "North Macedonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "North Macedonia", c.EN)
	}
	if c.TH != "นอร์ทมาซิโดเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์ทมาซิโดเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mk")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 450 {
		t.Errorf("invalid value: expect: %d actual: %d", 450, c.ID)
	}
    if c.EN != "Madagascar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Madagascar", c.EN)
	}
	if c.TH != "มาดากัสการ์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาดากัสการ์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 454 {
		t.Errorf("invalid value: expect: %d actual: %d", 454, c.ID)
	}
    if c.EN != "Malawi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malawi", c.EN)
	}
	if c.TH != "มาลาวี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลาวี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mw")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("my")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 458 {
		t.Errorf("invalid value: expect: %d actual: %d", 458, c.ID)
	}
    if c.EN != "Malaysia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malaysia", c.EN)
	}
	if c.TH != "มาเลเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาเลเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("my")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 462 {
		t.Errorf("invalid value: expect: %d actual: %d", 462, c.ID)
	}
    if c.EN != "Maldives" {
		t.Errorf("invalid value: expect: %s actual: %s", "Maldives", c.EN)
	}
	if c.TH != "มัลดีฟส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มัลดีฟส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mv")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ml")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 466 {
		t.Errorf("invalid value: expect: %d actual: %d", 466, c.ID)
	}
    if c.EN != "Mali" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mali", c.EN)
	}
	if c.TH != "มาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ml")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 470 {
		t.Errorf("invalid value: expect: %d actual: %d", 470, c.ID)
	}
    if c.EN != "Malta" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malta", c.EN)
	}
	if c.TH != "มอลตา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลตา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mt")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mh")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 584 {
		t.Errorf("invalid value: expect: %d actual: %d", 584, c.ID)
	}
    if c.EN != "Marshall Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Marshall Islands", c.EN)
	}
	if c.TH != "หมู่เกาะมาร์แชลล์" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะมาร์แชลล์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mh")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 478 {
		t.Errorf("invalid value: expect: %d actual: %d", 478, c.ID)
	}
    if c.EN != "Mauritania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritania", c.EN)
	}
	if c.TH != "มอริเตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเตเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 480 {
		t.Errorf("invalid value: expect: %d actual: %d", 480, c.ID)
	}
    if c.EN != "Mauritius" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritius", c.EN)
	}
	if c.TH != "มอริเชียส" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเชียส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mu")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mx")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 484 {
		t.Errorf("invalid value: expect: %d actual: %d", 484, c.ID)
	}
    if c.EN != "Mexico" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mexico", c.EN)
	}
	if c.TH != "เม็กซิโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เม็กซิโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mx")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("fm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 583 {
		t.Errorf("invalid value: expect: %d actual: %d", 583, c.ID)
	}
    if c.EN != "Micronesia (Federated States of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Micronesia (Federated States of)", c.EN)
	}
	if c.TH != "ไมโครนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไมโครนีเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("fm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ma")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 504 {
		t.Errorf("invalid value: expect: %d actual: %d", 504, c.ID)
	}
    if c.EN != "Morocco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Morocco", c.EN)
	}
	if c.TH != "โมร็อกโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมร็อกโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ma")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("md")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 498 {
		t.Errorf("invalid value: expect: %d actual: %d", 498, c.ID)
	}
    if c.EN != "Moldova, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Moldova, Republic of", c.EN)
	}
	if c.TH != "มอลโดวา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลโดวา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("md")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 492 {
		t.Errorf("invalid value: expect: %d actual: %d", 492, c.ID)
	}
    if c.EN != "Monaco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Monaco", c.EN)
	}
	if c.TH != "โมนาโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมนาโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mc")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 496 {
		t.Errorf("invalid value: expect: %d actual: %d", 496, c.ID)
	}
    if c.EN != "Mongolia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mongolia", c.EN)
	}
	if c.TH != "มองโกเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มองโกเลีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("me")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 499 {
		t.Errorf("invalid value: expect: %d actual: %d", 499, c.ID)
	}
    if c.EN != "Montenegro" {
		t.Errorf("invalid value: expect: %s actual: %s", "Montenegro", c.EN)
	}
	if c.TH != "มอนเตเนโกร" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอนเตเนโกร", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("me")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 508 {
		t.Errorf("invalid value: expect: %d actual: %d", 508, c.ID)
	}
    if c.EN != "Mozambique" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mozambique", c.EN)
	}
	if c.TH != "โมซัมบิก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมซัมบิก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("mm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 104 {
		t.Errorf("invalid value: expect: %d actual: %d", 104, c.ID)
	}
    if c.EN != "Myanmar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Myanmar", c.EN)
	}
	if c.TH != "พม่า" {
		t.Errorf("invalid value: expect: %s actual: %s", "พม่า", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("mm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("na")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 516 {
		t.Errorf("invalid value: expect: %d actual: %d", 516, c.ID)
	}
    if c.EN != "Namibia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Namibia", c.EN)
	}
	if c.TH != "นามิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นามิเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("na")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("nr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 520 {
		t.Errorf("invalid value: expect: %d actual: %d", 520, c.ID)
	}
    if c.EN != "Nauru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nauru", c.EN)
	}
	if c.TH != "นาอูรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "นาอูรู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("nr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("np")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 524 {
		t.Errorf("invalid value: expect: %d actual: %d", 524, c.ID)
	}
    if c.EN != "Nepal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nepal", c.EN)
	}
	if c.TH != "เนปาล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนปาล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("np")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("nl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 528 {
		t.Errorf("invalid value: expect: %d actual: %d", 528, c.ID)
	}
    if c.EN != "Netherlands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Netherlands", c.EN)
	}
	if c.TH != "เนเธอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนเธอร์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("nl")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("nz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 554 {
		t.Errorf("invalid value: expect: %d actual: %d", 554, c.ID)
	}
    if c.EN != "New Zealand" {
		t.Errorf("invalid value: expect: %s actual: %s", "New Zealand", c.EN)
	}
	if c.TH != "นิวซีแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิวซีแลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("nz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ni")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 558 {
		t.Errorf("invalid value: expect: %d actual: %d", 558, c.ID)
	}
    if c.EN != "Nicaragua" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nicaragua", c.EN)
	}
	if c.TH != "นิการากัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิการากัว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ni")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ne")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 562 {
		t.Errorf("invalid value: expect: %d actual: %d", 562, c.ID)
	}
    if c.EN != "Niger" {
		t.Errorf("invalid value: expect: %s actual: %s", "Niger", c.EN)
	}
	if c.TH != "ไนเจอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนเจอร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ne")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ng")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 566 {
		t.Errorf("invalid value: expect: %d actual: %d", 566, c.ID)
	}
    if c.EN != "Nigeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nigeria", c.EN)
	}
	if c.TH != "ไนจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนจีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ng")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("no")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 578 {
		t.Errorf("invalid value: expect: %d actual: %d", 578, c.ID)
	}
    if c.EN != "Norway" {
		t.Errorf("invalid value: expect: %s actual: %s", "Norway", c.EN)
	}
	if c.TH != "นอร์เวย์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์เวย์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("no")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("om")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 512 {
		t.Errorf("invalid value: expect: %d actual: %d", 512, c.ID)
	}
    if c.EN != "Oman" {
		t.Errorf("invalid value: expect: %s actual: %s", "Oman", c.EN)
	}
	if c.TH != "โอมาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "โอมาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("om")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("pk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 586 {
		t.Errorf("invalid value: expect: %d actual: %d", 586, c.ID)
	}
    if c.EN != "Pakistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Pakistan", c.EN)
	}
	if c.TH != "ปากีสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปากีสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("pk")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("pw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 585 {
		t.Errorf("invalid value: expect: %d actual: %d", 585, c.ID)
	}
    if c.EN != "Palau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Palau", c.EN)
	}
	if c.TH != "ปาเลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาเลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("pw")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("pa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 591 {
		t.Errorf("invalid value: expect: %d actual: %d", 591, c.ID)
	}
    if c.EN != "Panama" {
		t.Errorf("invalid value: expect: %s actual: %s", "Panama", c.EN)
	}
	if c.TH != "ปานามา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปานามา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("pa")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("pg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 598 {
		t.Errorf("invalid value: expect: %d actual: %d", 598, c.ID)
	}
    if c.EN != "Papua New Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Papua New Guinea", c.EN)
	}
	if c.TH != "ปาปัวนิวกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาปัวนิวกินี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("pg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("py")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 600 {
		t.Errorf("invalid value: expect: %d actual: %d", 600, c.ID)
	}
    if c.EN != "Paraguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Paraguay", c.EN)
	}
	if c.TH != "ปารากวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปารากวัย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("py")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("pe")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 604 {
		t.Errorf("invalid value: expect: %d actual: %d", 604, c.ID)
	}
    if c.EN != "Peru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Peru", c.EN)
	}
	if c.TH != "เปรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "เปรู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("pe")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ph")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 608 {
		t.Errorf("invalid value: expect: %d actual: %d", 608, c.ID)
	}
    if c.EN != "Philippines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Philippines", c.EN)
	}
	if c.TH != "ฟิลิปปินส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟิลิปปินส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ph")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("pl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 616 {
		t.Errorf("invalid value: expect: %d actual: %d", 616, c.ID)
	}
    if c.EN != "Poland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Poland", c.EN)
	}
	if c.TH != "โปแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปแลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("pl")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("pt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 620 {
		t.Errorf("invalid value: expect: %d actual: %d", 620, c.ID)
	}
    if c.EN != "Portugal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Portugal", c.EN)
	}
	if c.TH != "โปรตุเกส" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปรตุเกส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("pt")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("qa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 634 {
		t.Errorf("invalid value: expect: %d actual: %d", 634, c.ID)
	}
    if c.EN != "Qatar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Qatar", c.EN)
	}
	if c.TH != "กาตาร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาตาร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("qa")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ro")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 642 {
		t.Errorf("invalid value: expect: %d actual: %d", 642, c.ID)
	}
    if c.EN != "Romania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Romania", c.EN)
	}
	if c.TH != "โรมาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โรมาเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ro")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ru")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 643 {
		t.Errorf("invalid value: expect: %d actual: %d", 643, c.ID)
	}
    if c.EN != "Russian Federation" {
		t.Errorf("invalid value: expect: %s actual: %s", "Russian Federation", c.EN)
	}
	if c.TH != "รัสเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "รัสเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ru")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("rw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 646 {
		t.Errorf("invalid value: expect: %d actual: %d", 646, c.ID)
	}
    if c.EN != "Rwanda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Rwanda", c.EN)
	}
	if c.TH != "รวันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "รวันดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("rw")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("kn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 659 {
		t.Errorf("invalid value: expect: %d actual: %d", 659, c.ID)
	}
    if c.EN != "Saint Kitts and Nevis" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Kitts and Nevis", c.EN)
	}
	if c.TH != "เซนต์คิตส์และเนวิส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์คิตส์และเนวิส", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("kn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("lc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 662 {
		t.Errorf("invalid value: expect: %d actual: %d", 662, c.ID)
	}
    if c.EN != "Saint Lucia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Lucia", c.EN)
	}
	if c.TH != "เซนต์ลูเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์ลูเชีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("lc")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("vc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 670 {
		t.Errorf("invalid value: expect: %d actual: %d", 670, c.ID)
	}
    if c.EN != "Saint Vincent and the Grenadines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Vincent and the Grenadines", c.EN)
	}
	if c.TH != "เซนต์วินเซนต์และเกรนาดีนส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์วินเซนต์และเกรนาดีนส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("vc")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ws")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 882 {
		t.Errorf("invalid value: expect: %d actual: %d", 882, c.ID)
	}
    if c.EN != "Samoa" {
		t.Errorf("invalid value: expect: %s actual: %s", "Samoa", c.EN)
	}
	if c.TH != "ซามัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซามัว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ws")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 674 {
		t.Errorf("invalid value: expect: %d actual: %d", 674, c.ID)
	}
    if c.EN != "San Marino" {
		t.Errorf("invalid value: expect: %s actual: %s", "San Marino", c.EN)
	}
	if c.TH != "ซานมารีโน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซานมารีโน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("st")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 678 {
		t.Errorf("invalid value: expect: %d actual: %d", 678, c.ID)
	}
    if c.EN != "Sao Tome and Principe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sao Tome and Principe", c.EN)
	}
	if c.TH != "เซาตูเมและปรินซีปี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาตูเมและปรินซีปี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("st")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 682 {
		t.Errorf("invalid value: expect: %d actual: %d", 682, c.ID)
	}
    if c.EN != "Saudi Arabia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saudi Arabia", c.EN)
	}
	if c.TH != "ซาอุดีอาระเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซาอุดีอาระเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sa")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 686 {
		t.Errorf("invalid value: expect: %d actual: %d", 686, c.ID)
	}
    if c.EN != "Senegal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Senegal", c.EN)
	}
	if c.TH != "เซเนกัล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเนกัล", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("rs")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 688 {
		t.Errorf("invalid value: expect: %d actual: %d", 688, c.ID)
	}
    if c.EN != "Serbia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Serbia", c.EN)
	}
	if c.TH != "เซอร์เบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซอร์เบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("rs")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 690 {
		t.Errorf("invalid value: expect: %d actual: %d", 690, c.ID)
	}
    if c.EN != "Seychelles" {
		t.Errorf("invalid value: expect: %s actual: %s", "Seychelles", c.EN)
	}
	if c.TH != "เซเชลส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเชลส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sc")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 694 {
		t.Errorf("invalid value: expect: %d actual: %d", 694, c.ID)
	}
    if c.EN != "Sierra Leone" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sierra Leone", c.EN)
	}
	if c.TH != "เซียร์ราลีโอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซียร์ราลีโอน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sl")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 702 {
		t.Errorf("invalid value: expect: %d actual: %d", 702, c.ID)
	}
    if c.EN != "Singapore" {
		t.Errorf("invalid value: expect: %s actual: %s", "Singapore", c.EN)
	}
	if c.TH != "สิงคโปร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สิงคโปร์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 703 {
		t.Errorf("invalid value: expect: %d actual: %d", 703, c.ID)
	}
    if c.EN != "Slovakia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovakia", c.EN)
	}
	if c.TH != "สโลวาเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวาเกีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sk")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("si")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 705 {
		t.Errorf("invalid value: expect: %d actual: %d", 705, c.ID)
	}
    if c.EN != "Slovenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovenia", c.EN)
	}
	if c.TH != "สโลวีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวีเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("si")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 90 {
		t.Errorf("invalid value: expect: %d actual: %d", 90, c.ID)
	}
    if c.EN != "Solomon Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Solomon Islands", c.EN)
	}
	if c.TH != "หมู่เกาะโซโลมอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะโซโลมอน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sb")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("so")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 706 {
		t.Errorf("invalid value: expect: %d actual: %d", 706, c.ID)
	}
    if c.EN != "Somalia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Somalia", c.EN)
	}
	if c.TH != "โซมาเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โซมาเลีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("so")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("za")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 710 {
		t.Errorf("invalid value: expect: %d actual: %d", 710, c.ID)
	}
    if c.EN != "South Africa" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Africa", c.EN)
	}
	if c.TH != "แอฟริกาใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอฟริกาใต้", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("za")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ss")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 728 {
		t.Errorf("invalid value: expect: %d actual: %d", 728, c.ID)
	}
    if c.EN != "South Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Sudan", c.EN)
	}
	if c.TH != "เซาท์ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาท์ซูดาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ss")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("es")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 724 {
		t.Errorf("invalid value: expect: %d actual: %d", 724, c.ID)
	}
    if c.EN != "Spain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Spain", c.EN)
	}
	if c.TH != "สเปน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สเปน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("es")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("lk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 144 {
		t.Errorf("invalid value: expect: %d actual: %d", 144, c.ID)
	}
    if c.EN != "Sri Lanka" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sri Lanka", c.EN)
	}
	if c.TH != "ศรีลังกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ศรีลังกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("lk")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 729 {
		t.Errorf("invalid value: expect: %d actual: %d", 729, c.ID)
	}
    if c.EN != "Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sudan", c.EN)
	}
	if c.TH != "ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูดาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sd")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 740 {
		t.Errorf("invalid value: expect: %d actual: %d", 740, c.ID)
	}
    if c.EN != "Suriname" {
		t.Errorf("invalid value: expect: %s actual: %s", "Suriname", c.EN)
	}
	if c.TH != "ซูรินาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูรินาม", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 748 {
		t.Errorf("invalid value: expect: %d actual: %d", 748, c.ID)
	}
    if c.EN != "Eswatini" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eswatini", c.EN)
	}
	if c.TH != "เอสวาตีนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสวาตีนี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("se")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 752 {
		t.Errorf("invalid value: expect: %d actual: %d", 752, c.ID)
	}
    if c.EN != "Sweden" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sweden", c.EN)
	}
	if c.TH != "สวีเดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวีเดน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("se")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ch")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 756 {
		t.Errorf("invalid value: expect: %d actual: %d", 756, c.ID)
	}
    if c.EN != "Switzerland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Switzerland", c.EN)
	}
	if c.TH != "สวิตเซอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวิตเซอร์แลนด์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ch")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("sy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 760 {
		t.Errorf("invalid value: expect: %d actual: %d", 760, c.ID)
	}
    if c.EN != "Syrian Arab Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Syrian Arab Republic", c.EN)
	}
	if c.TH != "ซีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซีเรีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("sy")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tj")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 762 {
		t.Errorf("invalid value: expect: %d actual: %d", 762, c.ID)
	}
    if c.EN != "Tajikistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tajikistan", c.EN)
	}
	if c.TH != "ทาจิกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ทาจิกิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tj")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 834 {
		t.Errorf("invalid value: expect: %d actual: %d", 834, c.ID)
	}
    if c.EN != "Tanzania, United Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tanzania, United Republic of", c.EN)
	}
	if c.TH != "แทนซาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แทนซาเนีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("th")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 764 {
		t.Errorf("invalid value: expect: %d actual: %d", 764, c.ID)
	}
    if c.EN != "Thailand" {
		t.Errorf("invalid value: expect: %s actual: %s", "Thailand", c.EN)
	}
	if c.TH != "ไทย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไทย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("th")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 626 {
		t.Errorf("invalid value: expect: %d actual: %d", 626, c.ID)
	}
    if c.EN != "Timor-Leste" {
		t.Errorf("invalid value: expect: %s actual: %s", "Timor-Leste", c.EN)
	}
	if c.TH != "ติมอร์-เลสเต" {
		t.Errorf("invalid value: expect: %s actual: %s", "ติมอร์-เลสเต", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tl")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 768 {
		t.Errorf("invalid value: expect: %d actual: %d", 768, c.ID)
	}
    if c.EN != "Togo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Togo", c.EN)
	}
	if c.TH != "โตโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โตโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tg")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("to")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 776 {
		t.Errorf("invalid value: expect: %d actual: %d", 776, c.ID)
	}
    if c.EN != "Tonga" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tonga", c.EN)
	}
	if c.TH != "ตองงา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตองงา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("to")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 780 {
		t.Errorf("invalid value: expect: %d actual: %d", 780, c.ID)
	}
    if c.EN != "Trinidad and Tobago" {
		t.Errorf("invalid value: expect: %s actual: %s", "Trinidad and Tobago", c.EN)
	}
	if c.TH != "ตรินิแดดและโตเบโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตรินิแดดและโตเบโก", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tt")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 788 {
		t.Errorf("invalid value: expect: %d actual: %d", 788, c.ID)
	}
    if c.EN != "Tunisia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tunisia", c.EN)
	}
	if c.TH != "ตูนิเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูนิเซีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 792 {
		t.Errorf("invalid value: expect: %d actual: %d", 792, c.ID)
	}
    if c.EN != "Turkey" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkey", c.EN)
	}
	if c.TH != "ตุรกี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตุรกี", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tr")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 795 {
		t.Errorf("invalid value: expect: %d actual: %d", 795, c.ID)
	}
    if c.EN != "Turkmenistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkmenistan", c.EN)
	}
	if c.TH != "เติร์กเมนิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เติร์กเมนิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("tv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 798 {
		t.Errorf("invalid value: expect: %d actual: %d", 798, c.ID)
	}
    if c.EN != "Tuvalu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tuvalu", c.EN)
	}
	if c.TH != "ตูวาลู" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูวาลู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("tv")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ug")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 800 {
		t.Errorf("invalid value: expect: %d actual: %d", 800, c.ID)
	}
    if c.EN != "Uganda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uganda", c.EN)
	}
	if c.TH != "ยูกันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูกันดา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ug")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ua")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 804 {
		t.Errorf("invalid value: expect: %d actual: %d", 804, c.ID)
	}
    if c.EN != "Ukraine" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ukraine", c.EN)
	}
	if c.TH != "ยูเครน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูเครน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ua")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ae")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 784 {
		t.Errorf("invalid value: expect: %d actual: %d", 784, c.ID)
	}
    if c.EN != "United Arab Emirates" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Arab Emirates", c.EN)
	}
	if c.TH != "สหรัฐอาหรับเอมิเรตส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอาหรับเอมิเรตส์", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ae")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("gb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 826 {
		t.Errorf("invalid value: expect: %d actual: %d", 826, c.ID)
	}
    if c.EN != "United Kingdom of Great Britain and Northern Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Kingdom of Great Britain and Northern Ireland", c.EN)
	}
	if c.TH != "สหราชอาณาจักร" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหราชอาณาจักร", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("gb")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("us")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 840 {
		t.Errorf("invalid value: expect: %d actual: %d", 840, c.ID)
	}
    if c.EN != "United States of America" {
		t.Errorf("invalid value: expect: %s actual: %s", "United States of America", c.EN)
	}
	if c.TH != "สหรัฐอเมริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอเมริกา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("us")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("uy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 858 {
		t.Errorf("invalid value: expect: %d actual: %d", 858, c.ID)
	}
    if c.EN != "Uruguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uruguay", c.EN)
	}
	if c.TH != "อุรุกวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุรุกวัย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("uy")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("uz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 860 {
		t.Errorf("invalid value: expect: %d actual: %d", 860, c.ID)
	}
    if c.EN != "Uzbekistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uzbekistan", c.EN)
	}
	if c.TH != "อุซเบกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุซเบกิสถาน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("uz")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("vu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 548 {
		t.Errorf("invalid value: expect: %d actual: %d", 548, c.ID)
	}
    if c.EN != "Vanuatu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Vanuatu", c.EN)
	}
	if c.TH != "วานูอาตู" {
		t.Errorf("invalid value: expect: %s actual: %s", "วานูอาตู", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("vu")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ve")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 862 {
		t.Errorf("invalid value: expect: %d actual: %d", 862, c.ID)
	}
    if c.EN != "Venezuela (Bolivarian Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Venezuela (Bolivarian Republic of)", c.EN)
	}
	if c.TH != "เวเนซุเอลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวเนซุเอลา", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ve")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("vn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 704 {
		t.Errorf("invalid value: expect: %d actual: %d", 704, c.ID)
	}
    if c.EN != "Viet Nam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Viet Nam", c.EN)
	}
	if c.TH != "เวียดนาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวียดนาม", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("vn")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("ye")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 887 {
		t.Errorf("invalid value: expect: %d actual: %d", 887, c.ID)
	}
    if c.EN != "Yemen" {
		t.Errorf("invalid value: expect: %s actual: %s", "Yemen", c.EN)
	}
	if c.TH != "เยเมน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยเมน", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("ye")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("zm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 894 {
		t.Errorf("invalid value: expect: %d actual: %d", 894, c.ID)
	}
    if c.EN != "Zambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zambia", c.EN)
	}
	if c.TH != "แซมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แซมเบีย", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("zm")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}
	c, ok = gocountries.GetCountryFromAlpha2("zw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 716 {
		t.Errorf("invalid value: expect: %d actual: %d", 716, c.ID)
	}
    if c.EN != "Zimbabwe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zimbabwe", c.EN)
	}
	if c.TH != "ซิมบับเว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซิมบับเว", c.TH)
	}
	c1, ok = gocountries.GetCountryFromAlpha2("zw")
	if &c == &c1 {
		t.Errorf("invalid: address: &c: %p &c1: %p", &c, &c1)
	}

	c, ok = gocountries.GetCountryFromAlpha2("aaa")
	if ok {
		t.Errorf("invalid value: expect: %t actual: %t", false, ok)
	}
    if c.ID != 0 {
		t.Errorf("invalid value: expect: %d actual: %d", 0, c.ID)
	}
    if c.EN != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.EN)
	}
	if c.TH != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.TH)
	}
}

func TestGetCountryFromAlpha3(t *testing.T) {
    var c gocountries.Country
	var ok bool

	c, ok = gocountries.GetCountryFromAlpha3("afg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 4 {
		t.Errorf("invalid value: expect: %d actual: %d", 4, c.ID)
	}
    if c.EN != "Afghanistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Afghanistan", c.EN)
	}
	if c.TH != "อัฟกานิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อัฟกานิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("alb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 8 {
		t.Errorf("invalid value: expect: %d actual: %d", 8, c.ID)
	}
    if c.EN != "Albania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Albania", c.EN)
	}
	if c.TH != "แอลเบเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลเบเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("dza")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 12 {
		t.Errorf("invalid value: expect: %d actual: %d", 12, c.ID)
	}
    if c.EN != "Algeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Algeria", c.EN)
	}
	if c.TH != "แอลจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลจีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("and")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 20 {
		t.Errorf("invalid value: expect: %d actual: %d", 20, c.ID)
	}
    if c.EN != "Andorra" {
		t.Errorf("invalid value: expect: %s actual: %s", "Andorra", c.EN)
	}
	if c.TH != "อันดอร์รา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อันดอร์รา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ago")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 24 {
		t.Errorf("invalid value: expect: %d actual: %d", 24, c.ID)
	}
    if c.EN != "Angola" {
		t.Errorf("invalid value: expect: %s actual: %s", "Angola", c.EN)
	}
	if c.TH != "แองโกลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แองโกลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("atg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 28 {
		t.Errorf("invalid value: expect: %d actual: %d", 28, c.ID)
	}
    if c.EN != "Antigua and Barbuda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Antigua and Barbuda", c.EN)
	}
	if c.TH != "แอนติกาและบาร์บูดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอนติกาและบาร์บูดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("arg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 32 {
		t.Errorf("invalid value: expect: %d actual: %d", 32, c.ID)
	}
    if c.EN != "Argentina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Argentina", c.EN)
	}
	if c.TH != "อาร์เจนตินา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์เจนตินา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("arm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 51 {
		t.Errorf("invalid value: expect: %d actual: %d", 51, c.ID)
	}
    if c.EN != "Armenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Armenia", c.EN)
	}
	if c.TH != "อาร์มีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์มีเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("aus")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 36 {
		t.Errorf("invalid value: expect: %d actual: %d", 36, c.ID)
	}
    if c.EN != "Australia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Australia", c.EN)
	}
	if c.TH != "ออสเตรเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรเลีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("aut")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 40 {
		t.Errorf("invalid value: expect: %d actual: %d", 40, c.ID)
	}
    if c.EN != "Austria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Austria", c.EN)
	}
	if c.TH != "ออสเตรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("aze")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 31 {
		t.Errorf("invalid value: expect: %d actual: %d", 31, c.ID)
	}
    if c.EN != "Azerbaijan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Azerbaijan", c.EN)
	}
	if c.TH != "อาเซอร์ไบจาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาเซอร์ไบจาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bhs")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 44 {
		t.Errorf("invalid value: expect: %d actual: %d", 44, c.ID)
	}
    if c.EN != "Bahamas" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahamas", c.EN)
	}
	if c.TH != "บาฮามาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาฮามาส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bhr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 48 {
		t.Errorf("invalid value: expect: %d actual: %d", 48, c.ID)
	}
    if c.EN != "Bahrain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahrain", c.EN)
	}
	if c.TH != "บาห์เรน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาห์เรน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bgd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 50 {
		t.Errorf("invalid value: expect: %d actual: %d", 50, c.ID)
	}
    if c.EN != "Bangladesh" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bangladesh", c.EN)
	}
	if c.TH != "บังกลาเทศ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บังกลาเทศ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("brb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 52 {
		t.Errorf("invalid value: expect: %d actual: %d", 52, c.ID)
	}
    if c.EN != "Barbados" {
		t.Errorf("invalid value: expect: %s actual: %s", "Barbados", c.EN)
	}
	if c.TH != "บาร์เบโดส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาร์เบโดส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("blr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 112 {
		t.Errorf("invalid value: expect: %d actual: %d", 112, c.ID)
	}
    if c.EN != "Belarus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belarus", c.EN)
	}
	if c.TH != "เบลารุส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลารุส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bel")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 56 {
		t.Errorf("invalid value: expect: %d actual: %d", 56, c.ID)
	}
    if c.EN != "Belgium" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belgium", c.EN)
	}
	if c.TH != "เบลเยียม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลเยียม", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("blz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 84 {
		t.Errorf("invalid value: expect: %d actual: %d", 84, c.ID)
	}
    if c.EN != "Belize" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belize", c.EN)
	}
	if c.TH != "เบลีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลีซ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ben")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 204 {
		t.Errorf("invalid value: expect: %d actual: %d", 204, c.ID)
	}
    if c.EN != "Benin" {
		t.Errorf("invalid value: expect: %s actual: %s", "Benin", c.EN)
	}
	if c.TH != "เบนิน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบนิน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("btn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 64 {
		t.Errorf("invalid value: expect: %d actual: %d", 64, c.ID)
	}
    if c.EN != "Bhutan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bhutan", c.EN)
	}
	if c.TH != "ภูฏาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ภูฏาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bol")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 68 {
		t.Errorf("invalid value: expect: %d actual: %d", 68, c.ID)
	}
    if c.EN != "Bolivia (Plurinational State of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bolivia (Plurinational State of)", c.EN)
	}
	if c.TH != "โบลิเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โบลิเวีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bih")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 70 {
		t.Errorf("invalid value: expect: %d actual: %d", 70, c.ID)
	}
    if c.EN != "Bosnia and Herzegovina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bosnia and Herzegovina", c.EN)
	}
	if c.TH != "บอสเนียและเฮอร์เซโกวีนา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอสเนียและเฮอร์เซโกวีนา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bwa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 72 {
		t.Errorf("invalid value: expect: %d actual: %d", 72, c.ID)
	}
    if c.EN != "Botswana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Botswana", c.EN)
	}
	if c.TH != "บอตสวานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอตสวานา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bra")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 76 {
		t.Errorf("invalid value: expect: %d actual: %d", 76, c.ID)
	}
    if c.EN != "Brazil" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brazil", c.EN)
	}
	if c.TH != "บราซิล" {
		t.Errorf("invalid value: expect: %s actual: %s", "บราซิล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("brn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 96 {
		t.Errorf("invalid value: expect: %d actual: %d", 96, c.ID)
	}
    if c.EN != "Brunei Darussalam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brunei Darussalam", c.EN)
	}
	if c.TH != "บรูไน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บรูไน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bgr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 100 {
		t.Errorf("invalid value: expect: %d actual: %d", 100, c.ID)
	}
    if c.EN != "Bulgaria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bulgaria", c.EN)
	}
	if c.TH != "บัลแกเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "บัลแกเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bfa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 854 {
		t.Errorf("invalid value: expect: %d actual: %d", 854, c.ID)
	}
    if c.EN != "Burkina Faso" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burkina Faso", c.EN)
	}
	if c.TH != "บูร์กินาฟาโซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บูร์กินาฟาโซ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("bdi")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 108 {
		t.Errorf("invalid value: expect: %d actual: %d", 108, c.ID)
	}
    if c.EN != "Burundi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burundi", c.EN)
	}
	if c.TH != "บุรุนดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "บุรุนดี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("khm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 116 {
		t.Errorf("invalid value: expect: %d actual: %d", 116, c.ID)
	}
    if c.EN != "Cambodia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cambodia", c.EN)
	}
	if c.TH != "กัมพูชา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัมพูชา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cmr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 120 {
		t.Errorf("invalid value: expect: %d actual: %d", 120, c.ID)
	}
    if c.EN != "Cameroon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cameroon", c.EN)
	}
	if c.TH != "แคเมอรูน" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคเมอรูน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("can")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 124 {
		t.Errorf("invalid value: expect: %d actual: %d", 124, c.ID)
	}
    if c.EN != "Canada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Canada", c.EN)
	}
	if c.TH != "แคนาดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคนาดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cpv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 132 {
		t.Errorf("invalid value: expect: %d actual: %d", 132, c.ID)
	}
    if c.EN != "Cabo Verde" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cabo Verde", c.EN)
	}
	if c.TH != "กาบูเวร์ดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบูเวร์ดี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("caf")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 140 {
		t.Errorf("invalid value: expect: %d actual: %d", 140, c.ID)
	}
    if c.EN != "Central African Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Central African Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐแอฟริกากลาง" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐแอฟริกากลาง", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tcd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 148 {
		t.Errorf("invalid value: expect: %d actual: %d", 148, c.ID)
	}
    if c.EN != "Chad" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chad", c.EN)
	}
	if c.TH != "ชาด" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชาด", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("chl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 152 {
		t.Errorf("invalid value: expect: %d actual: %d", 152, c.ID)
	}
    if c.EN != "Chile" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chile", c.EN)
	}
	if c.TH != "ชิลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชิลี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("chn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 156 {
		t.Errorf("invalid value: expect: %d actual: %d", 156, c.ID)
	}
    if c.EN != "China" {
		t.Errorf("invalid value: expect: %s actual: %s", "China", c.EN)
	}
	if c.TH != "จีน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จีน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("col")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 170 {
		t.Errorf("invalid value: expect: %d actual: %d", 170, c.ID)
	}
    if c.EN != "Colombia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Colombia", c.EN)
	}
	if c.TH != "โคลอมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โคลอมเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("com")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 174 {
		t.Errorf("invalid value: expect: %d actual: %d", 174, c.ID)
	}
    if c.EN != "Comoros" {
		t.Errorf("invalid value: expect: %s actual: %s", "Comoros", c.EN)
	}
	if c.TH != "คอโมโรส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอโมโรส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cog")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 178 {
		t.Errorf("invalid value: expect: %d actual: %d", 178, c.ID)
	}
    if c.EN != "Congo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo", c.EN)
	}
	if c.TH != "สาธารณรัฐคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐคองโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cod")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 180 {
		t.Errorf("invalid value: expect: %d actual: %d", 180, c.ID)
	}
    if c.EN != "Congo, Democratic Republic of the" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo, Democratic Republic of the", c.EN)
	}
	if c.TH != "สาธารณรัฐประชาธิปไตยคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐประชาธิปไตยคองโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cri")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 188 {
		t.Errorf("invalid value: expect: %d actual: %d", 188, c.ID)
	}
    if c.EN != "Costa Rica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Costa Rica", c.EN)
	}
	if c.TH != "คอสตาริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอสตาริกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("civ")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 384 {
		t.Errorf("invalid value: expect: %d actual: %d", 384, c.ID)
	}
    if c.EN != "Côte d'Ivoire" {
		t.Errorf("invalid value: expect: %s actual: %s", "Côte d'Ivoire", c.EN)
	}
	if c.TH != "โกตดิวัวร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โกตดิวัวร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("hrv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 191 {
		t.Errorf("invalid value: expect: %d actual: %d", 191, c.ID)
	}
    if c.EN != "Croatia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Croatia", c.EN)
	}
	if c.TH != "โครเอเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โครเอเชีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cub")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 192 {
		t.Errorf("invalid value: expect: %d actual: %d", 192, c.ID)
	}
    if c.EN != "Cuba" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cuba", c.EN)
	}
	if c.TH != "คิวบา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิวบา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cyp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 196 {
		t.Errorf("invalid value: expect: %d actual: %d", 196, c.ID)
	}
    if c.EN != "Cyprus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cyprus", c.EN)
	}
	if c.TH != "ไซปรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไซปรัส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("cze")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 203 {
		t.Errorf("invalid value: expect: %d actual: %d", 203, c.ID)
	}
    if c.EN != "Czechia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Czechia", c.EN)
	}
	if c.TH != "เช็กเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เช็กเกีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("dnk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 208 {
		t.Errorf("invalid value: expect: %d actual: %d", 208, c.ID)
	}
    if c.EN != "Denmark" {
		t.Errorf("invalid value: expect: %s actual: %s", "Denmark", c.EN)
	}
	if c.TH != "เดนมาร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เดนมาร์ก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("dji")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 262 {
		t.Errorf("invalid value: expect: %d actual: %d", 262, c.ID)
	}
    if c.EN != "Djibouti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Djibouti", c.EN)
	}
	if c.TH != "จิบูตี" {
		t.Errorf("invalid value: expect: %s actual: %s", "จิบูตี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("dma")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 212 {
		t.Errorf("invalid value: expect: %d actual: %d", 212, c.ID)
	}
    if c.EN != "Dominica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominica", c.EN)
	}
	if c.TH != "ดอมินีกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ดอมินีกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("dom")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 214 {
		t.Errorf("invalid value: expect: %d actual: %d", 214, c.ID)
	}
    if c.EN != "Dominican Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominican Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐโดมินิกัน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐโดมินิกัน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ecu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 218 {
		t.Errorf("invalid value: expect: %d actual: %d", 218, c.ID)
	}
    if c.EN != "Ecuador" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ecuador", c.EN)
	}
	if c.TH != "เอกวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอกวาดอร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("egy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 818 {
		t.Errorf("invalid value: expect: %d actual: %d", 818, c.ID)
	}
    if c.EN != "Egypt" {
		t.Errorf("invalid value: expect: %s actual: %s", "Egypt", c.EN)
	}
	if c.TH != "อียิปต์" {
		t.Errorf("invalid value: expect: %s actual: %s", "อียิปต์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("slv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 222 {
		t.Errorf("invalid value: expect: %d actual: %d", 222, c.ID)
	}
    if c.EN != "El Salvador" {
		t.Errorf("invalid value: expect: %s actual: %s", "El Salvador", c.EN)
	}
	if c.TH != "เอลซัลวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอลซัลวาดอร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gnq")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 226 {
		t.Errorf("invalid value: expect: %d actual: %d", 226, c.ID)
	}
    if c.EN != "Equatorial Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Equatorial Guinea", c.EN)
	}
	if c.TH != "อิเควทอเรียลกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิเควทอเรียลกินี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("eri")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 232 {
		t.Errorf("invalid value: expect: %d actual: %d", 232, c.ID)
	}
    if c.EN != "Eritrea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eritrea", c.EN)
	}
	if c.TH != "เอริเทรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอริเทรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("est")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 233 {
		t.Errorf("invalid value: expect: %d actual: %d", 233, c.ID)
	}
    if c.EN != "Estonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Estonia", c.EN)
	}
	if c.TH != "เอสโตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสโตเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("eth")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 231 {
		t.Errorf("invalid value: expect: %d actual: %d", 231, c.ID)
	}
    if c.EN != "Ethiopia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ethiopia", c.EN)
	}
	if c.TH != "เอธิโอเปีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอธิโอเปีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("fji")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 242 {
		t.Errorf("invalid value: expect: %d actual: %d", 242, c.ID)
	}
    if c.EN != "Fiji" {
		t.Errorf("invalid value: expect: %s actual: %s", "Fiji", c.EN)
	}
	if c.TH != "ฟีจี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟีจี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("fin")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 246 {
		t.Errorf("invalid value: expect: %d actual: %d", 246, c.ID)
	}
    if c.EN != "Finland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Finland", c.EN)
	}
	if c.TH != "ฟินแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟินแลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("fra")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 250 {
		t.Errorf("invalid value: expect: %d actual: %d", 250, c.ID)
	}
    if c.EN != "France" {
		t.Errorf("invalid value: expect: %s actual: %s", "France", c.EN)
	}
	if c.TH != "ฝรั่งเศส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฝรั่งเศส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gab")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 266 {
		t.Errorf("invalid value: expect: %d actual: %d", 266, c.ID)
	}
    if c.EN != "Gabon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gabon", c.EN)
	}
	if c.TH != "กาบอง" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบอง", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gmb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 270 {
		t.Errorf("invalid value: expect: %d actual: %d", 270, c.ID)
	}
    if c.EN != "Gambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gambia", c.EN)
	}
	if c.TH != "แกมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แกมเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("geo")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 268 {
		t.Errorf("invalid value: expect: %d actual: %d", 268, c.ID)
	}
    if c.EN != "Georgia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Georgia", c.EN)
	}
	if c.TH != "จอร์เจีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์เจีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("deu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 276 {
		t.Errorf("invalid value: expect: %d actual: %d", 276, c.ID)
	}
    if c.EN != "Germany" {
		t.Errorf("invalid value: expect: %s actual: %s", "Germany", c.EN)
	}
	if c.TH != "เยอรมนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยอรมนี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gha")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 288 {
		t.Errorf("invalid value: expect: %d actual: %d", 288, c.ID)
	}
    if c.EN != "Ghana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ghana", c.EN)
	}
	if c.TH != "กานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กานา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("grc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 300 {
		t.Errorf("invalid value: expect: %d actual: %d", 300, c.ID)
	}
    if c.EN != "Greece" {
		t.Errorf("invalid value: expect: %s actual: %s", "Greece", c.EN)
	}
	if c.TH != "กรีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "กรีซ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("grd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 308 {
		t.Errorf("invalid value: expect: %d actual: %d", 308, c.ID)
	}
    if c.EN != "Grenada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Grenada", c.EN)
	}
	if c.TH != "เกรเนดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกรเนดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gtm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 320 {
		t.Errorf("invalid value: expect: %d actual: %d", 320, c.ID)
	}
    if c.EN != "Guatemala" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guatemala", c.EN)
	}
	if c.TH != "กัวเตมาลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัวเตมาลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gin")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 324 {
		t.Errorf("invalid value: expect: %d actual: %d", 324, c.ID)
	}
    if c.EN != "Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea", c.EN)
	}
	if c.TH != "กินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gnb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 624 {
		t.Errorf("invalid value: expect: %d actual: %d", 624, c.ID)
	}
    if c.EN != "Guinea-Bissau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea-Bissau", c.EN)
	}
	if c.TH != "กินี-บิสเซา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี-บิสเซา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("guy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 328 {
		t.Errorf("invalid value: expect: %d actual: %d", 328, c.ID)
	}
    if c.EN != "Guyana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guyana", c.EN)
	}
	if c.TH != "กายอานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กายอานา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("hti")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 332 {
		t.Errorf("invalid value: expect: %d actual: %d", 332, c.ID)
	}
    if c.EN != "Haiti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Haiti", c.EN)
	}
	if c.TH != "เฮติ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เฮติ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("hnd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 340 {
		t.Errorf("invalid value: expect: %d actual: %d", 340, c.ID)
	}
    if c.EN != "Honduras" {
		t.Errorf("invalid value: expect: %s actual: %s", "Honduras", c.EN)
	}
	if c.TH != "ฮอนดูรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮอนดูรัส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("hun")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 348 {
		t.Errorf("invalid value: expect: %d actual: %d", 348, c.ID)
	}
    if c.EN != "Hungary" {
		t.Errorf("invalid value: expect: %s actual: %s", "Hungary", c.EN)
	}
	if c.TH != "ฮังการี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮังการี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("isl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 352 {
		t.Errorf("invalid value: expect: %d actual: %d", 352, c.ID)
	}
    if c.EN != "Iceland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iceland", c.EN)
	}
	if c.TH != "ไอซ์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอซ์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ind")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 356 {
		t.Errorf("invalid value: expect: %d actual: %d", 356, c.ID)
	}
    if c.EN != "India" {
		t.Errorf("invalid value: expect: %s actual: %s", "India", c.EN)
	}
	if c.TH != "อินเดีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินเดีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("idn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 360 {
		t.Errorf("invalid value: expect: %d actual: %d", 360, c.ID)
	}
    if c.EN != "Indonesia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Indonesia", c.EN)
	}
	if c.TH != "อินโดนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินโดนีเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("irn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 364 {
		t.Errorf("invalid value: expect: %d actual: %d", 364, c.ID)
	}
    if c.EN != "Iran (Islamic Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iran (Islamic Republic of)", c.EN)
	}
	if c.TH != "อิหร่าน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิหร่าน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("irq")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 368 {
		t.Errorf("invalid value: expect: %d actual: %d", 368, c.ID)
	}
    if c.EN != "Iraq" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iraq", c.EN)
	}
	if c.TH != "อิรัก" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิรัก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("irl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 372 {
		t.Errorf("invalid value: expect: %d actual: %d", 372, c.ID)
	}
    if c.EN != "Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ireland", c.EN)
	}
	if c.TH != "ไอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอร์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("isr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 376 {
		t.Errorf("invalid value: expect: %d actual: %d", 376, c.ID)
	}
    if c.EN != "Israel" {
		t.Errorf("invalid value: expect: %s actual: %s", "Israel", c.EN)
	}
	if c.TH != "อิสราเอล" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิสราเอล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ita")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 380 {
		t.Errorf("invalid value: expect: %d actual: %d", 380, c.ID)
	}
    if c.EN != "Italy" {
		t.Errorf("invalid value: expect: %s actual: %s", "Italy", c.EN)
	}
	if c.TH != "อิตาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิตาลี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("jam")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 388 {
		t.Errorf("invalid value: expect: %d actual: %d", 388, c.ID)
	}
    if c.EN != "Jamaica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jamaica", c.EN)
	}
	if c.TH != "จาเมกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "จาเมกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("jpn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 392 {
		t.Errorf("invalid value: expect: %d actual: %d", 392, c.ID)
	}
    if c.EN != "Japan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Japan", c.EN)
	}
	if c.TH != "ญี่ปุ่น" {
		t.Errorf("invalid value: expect: %s actual: %s", "ญี่ปุ่น", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("jor")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 400 {
		t.Errorf("invalid value: expect: %d actual: %d", 400, c.ID)
	}
    if c.EN != "Jordan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jordan", c.EN)
	}
	if c.TH != "จอร์แดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์แดน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("kaz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 398 {
		t.Errorf("invalid value: expect: %d actual: %d", 398, c.ID)
	}
    if c.EN != "Kazakhstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kazakhstan", c.EN)
	}
	if c.TH != "คาซัคสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คาซัคสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ken")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 404 {
		t.Errorf("invalid value: expect: %d actual: %d", 404, c.ID)
	}
    if c.EN != "Kenya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kenya", c.EN)
	}
	if c.TH != "เคนยา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เคนยา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("kir")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 296 {
		t.Errorf("invalid value: expect: %d actual: %d", 296, c.ID)
	}
    if c.EN != "Kiribati" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kiribati", c.EN)
	}
	if c.TH != "คิริบาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิริบาส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("prk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 408 {
		t.Errorf("invalid value: expect: %d actual: %d", 408, c.ID)
	}
    if c.EN != "Korea (Democratic People's Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea (Democratic People's Republic of)", c.EN)
	}
	if c.TH != "เกาหลีเหนือ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีเหนือ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("kor")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 410 {
		t.Errorf("invalid value: expect: %d actual: %d", 410, c.ID)
	}
    if c.EN != "Korea, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea, Republic of", c.EN)
	}
	if c.TH != "เกาหลีใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีใต้", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("kwt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 414 {
		t.Errorf("invalid value: expect: %d actual: %d", 414, c.ID)
	}
    if c.EN != "Kuwait" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kuwait", c.EN)
	}
	if c.TH != "คูเวต" {
		t.Errorf("invalid value: expect: %s actual: %s", "คูเวต", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("kgz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 417 {
		t.Errorf("invalid value: expect: %d actual: %d", 417, c.ID)
	}
    if c.EN != "Kyrgyzstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kyrgyzstan", c.EN)
	}
	if c.TH != "คีร์กีซสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คีร์กีซสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lao")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 418 {
		t.Errorf("invalid value: expect: %d actual: %d", 418, c.ID)
	}
    if c.EN != "Lao People's Democratic Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lao People's Democratic Republic", c.EN)
	}
	if c.TH != "ลาว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลาว", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lva")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 428 {
		t.Errorf("invalid value: expect: %d actual: %d", 428, c.ID)
	}
    if c.EN != "Latvia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Latvia", c.EN)
	}
	if c.TH != "ลัตเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลัตเวีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lbn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 422 {
		t.Errorf("invalid value: expect: %d actual: %d", 422, c.ID)
	}
    if c.EN != "Lebanon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lebanon", c.EN)
	}
	if c.TH != "เลบานอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลบานอน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lso")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 426 {
		t.Errorf("invalid value: expect: %d actual: %d", 426, c.ID)
	}
    if c.EN != "Lesotho" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lesotho", c.EN)
	}
	if c.TH != "เลโซโท" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลโซโท", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lbr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 430 {
		t.Errorf("invalid value: expect: %d actual: %d", 430, c.ID)
	}
    if c.EN != "Liberia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liberia", c.EN)
	}
	if c.TH != "ไลบีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไลบีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lby")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 434 {
		t.Errorf("invalid value: expect: %d actual: %d", 434, c.ID)
	}
    if c.EN != "Libya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Libya", c.EN)
	}
	if c.TH != "ลิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lie")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 438 {
		t.Errorf("invalid value: expect: %d actual: %d", 438, c.ID)
	}
    if c.EN != "Liechtenstein" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liechtenstein", c.EN)
	}
	if c.TH != "ลิกเตนสไตน์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิกเตนสไตน์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ltu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 440 {
		t.Errorf("invalid value: expect: %d actual: %d", 440, c.ID)
	}
    if c.EN != "Lithuania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lithuania", c.EN)
	}
	if c.TH != "ลิทัวเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิทัวเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lux")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 442 {
		t.Errorf("invalid value: expect: %d actual: %d", 442, c.ID)
	}
    if c.EN != "Luxembourg" {
		t.Errorf("invalid value: expect: %s actual: %s", "Luxembourg", c.EN)
	}
	if c.TH != "ลักเซมเบิร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลักเซมเบิร์ก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mkd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 807 {
		t.Errorf("invalid value: expect: %d actual: %d", 807, c.ID)
	}
    if c.EN != "North Macedonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "North Macedonia", c.EN)
	}
	if c.TH != "นอร์ทมาซิโดเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์ทมาซิโดเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mdg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 450 {
		t.Errorf("invalid value: expect: %d actual: %d", 450, c.ID)
	}
    if c.EN != "Madagascar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Madagascar", c.EN)
	}
	if c.TH != "มาดากัสการ์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาดากัสการ์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mwi")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 454 {
		t.Errorf("invalid value: expect: %d actual: %d", 454, c.ID)
	}
    if c.EN != "Malawi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malawi", c.EN)
	}
	if c.TH != "มาลาวี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลาวี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mys")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 458 {
		t.Errorf("invalid value: expect: %d actual: %d", 458, c.ID)
	}
    if c.EN != "Malaysia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malaysia", c.EN)
	}
	if c.TH != "มาเลเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาเลเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mdv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 462 {
		t.Errorf("invalid value: expect: %d actual: %d", 462, c.ID)
	}
    if c.EN != "Maldives" {
		t.Errorf("invalid value: expect: %s actual: %s", "Maldives", c.EN)
	}
	if c.TH != "มัลดีฟส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มัลดีฟส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mli")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 466 {
		t.Errorf("invalid value: expect: %d actual: %d", 466, c.ID)
	}
    if c.EN != "Mali" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mali", c.EN)
	}
	if c.TH != "มาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mlt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 470 {
		t.Errorf("invalid value: expect: %d actual: %d", 470, c.ID)
	}
    if c.EN != "Malta" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malta", c.EN)
	}
	if c.TH != "มอลตา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลตา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mhl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 584 {
		t.Errorf("invalid value: expect: %d actual: %d", 584, c.ID)
	}
    if c.EN != "Marshall Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Marshall Islands", c.EN)
	}
	if c.TH != "หมู่เกาะมาร์แชลล์" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะมาร์แชลล์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mrt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 478 {
		t.Errorf("invalid value: expect: %d actual: %d", 478, c.ID)
	}
    if c.EN != "Mauritania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritania", c.EN)
	}
	if c.TH != "มอริเตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเตเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mus")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 480 {
		t.Errorf("invalid value: expect: %d actual: %d", 480, c.ID)
	}
    if c.EN != "Mauritius" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritius", c.EN)
	}
	if c.TH != "มอริเชียส" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเชียส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mex")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 484 {
		t.Errorf("invalid value: expect: %d actual: %d", 484, c.ID)
	}
    if c.EN != "Mexico" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mexico", c.EN)
	}
	if c.TH != "เม็กซิโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เม็กซิโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("fsm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 583 {
		t.Errorf("invalid value: expect: %d actual: %d", 583, c.ID)
	}
    if c.EN != "Micronesia (Federated States of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Micronesia (Federated States of)", c.EN)
	}
	if c.TH != "ไมโครนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไมโครนีเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mar")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 504 {
		t.Errorf("invalid value: expect: %d actual: %d", 504, c.ID)
	}
    if c.EN != "Morocco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Morocco", c.EN)
	}
	if c.TH != "โมร็อกโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมร็อกโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mda")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 498 {
		t.Errorf("invalid value: expect: %d actual: %d", 498, c.ID)
	}
    if c.EN != "Moldova, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Moldova, Republic of", c.EN)
	}
	if c.TH != "มอลโดวา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลโดวา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mco")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 492 {
		t.Errorf("invalid value: expect: %d actual: %d", 492, c.ID)
	}
    if c.EN != "Monaco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Monaco", c.EN)
	}
	if c.TH != "โมนาโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมนาโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mng")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 496 {
		t.Errorf("invalid value: expect: %d actual: %d", 496, c.ID)
	}
    if c.EN != "Mongolia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mongolia", c.EN)
	}
	if c.TH != "มองโกเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มองโกเลีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mne")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 499 {
		t.Errorf("invalid value: expect: %d actual: %d", 499, c.ID)
	}
    if c.EN != "Montenegro" {
		t.Errorf("invalid value: expect: %s actual: %s", "Montenegro", c.EN)
	}
	if c.TH != "มอนเตเนโกร" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอนเตเนโกร", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("moz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 508 {
		t.Errorf("invalid value: expect: %d actual: %d", 508, c.ID)
	}
    if c.EN != "Mozambique" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mozambique", c.EN)
	}
	if c.TH != "โมซัมบิก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมซัมบิก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("mmr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 104 {
		t.Errorf("invalid value: expect: %d actual: %d", 104, c.ID)
	}
    if c.EN != "Myanmar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Myanmar", c.EN)
	}
	if c.TH != "พม่า" {
		t.Errorf("invalid value: expect: %s actual: %s", "พม่า", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("nam")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 516 {
		t.Errorf("invalid value: expect: %d actual: %d", 516, c.ID)
	}
    if c.EN != "Namibia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Namibia", c.EN)
	}
	if c.TH != "นามิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นามิเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("nru")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 520 {
		t.Errorf("invalid value: expect: %d actual: %d", 520, c.ID)
	}
    if c.EN != "Nauru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nauru", c.EN)
	}
	if c.TH != "นาอูรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "นาอูรู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("npl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 524 {
		t.Errorf("invalid value: expect: %d actual: %d", 524, c.ID)
	}
    if c.EN != "Nepal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nepal", c.EN)
	}
	if c.TH != "เนปาล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนปาล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("nld")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 528 {
		t.Errorf("invalid value: expect: %d actual: %d", 528, c.ID)
	}
    if c.EN != "Netherlands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Netherlands", c.EN)
	}
	if c.TH != "เนเธอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนเธอร์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("nzl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 554 {
		t.Errorf("invalid value: expect: %d actual: %d", 554, c.ID)
	}
    if c.EN != "New Zealand" {
		t.Errorf("invalid value: expect: %s actual: %s", "New Zealand", c.EN)
	}
	if c.TH != "นิวซีแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิวซีแลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("nic")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 558 {
		t.Errorf("invalid value: expect: %d actual: %d", 558, c.ID)
	}
    if c.EN != "Nicaragua" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nicaragua", c.EN)
	}
	if c.TH != "นิการากัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิการากัว", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ner")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 562 {
		t.Errorf("invalid value: expect: %d actual: %d", 562, c.ID)
	}
    if c.EN != "Niger" {
		t.Errorf("invalid value: expect: %s actual: %s", "Niger", c.EN)
	}
	if c.TH != "ไนเจอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนเจอร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("nga")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 566 {
		t.Errorf("invalid value: expect: %d actual: %d", 566, c.ID)
	}
    if c.EN != "Nigeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nigeria", c.EN)
	}
	if c.TH != "ไนจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนจีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("nor")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 578 {
		t.Errorf("invalid value: expect: %d actual: %d", 578, c.ID)
	}
    if c.EN != "Norway" {
		t.Errorf("invalid value: expect: %s actual: %s", "Norway", c.EN)
	}
	if c.TH != "นอร์เวย์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์เวย์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("omn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 512 {
		t.Errorf("invalid value: expect: %d actual: %d", 512, c.ID)
	}
    if c.EN != "Oman" {
		t.Errorf("invalid value: expect: %s actual: %s", "Oman", c.EN)
	}
	if c.TH != "โอมาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "โอมาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("pak")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 586 {
		t.Errorf("invalid value: expect: %d actual: %d", 586, c.ID)
	}
    if c.EN != "Pakistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Pakistan", c.EN)
	}
	if c.TH != "ปากีสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปากีสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("plw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 585 {
		t.Errorf("invalid value: expect: %d actual: %d", 585, c.ID)
	}
    if c.EN != "Palau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Palau", c.EN)
	}
	if c.TH != "ปาเลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาเลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("pan")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 591 {
		t.Errorf("invalid value: expect: %d actual: %d", 591, c.ID)
	}
    if c.EN != "Panama" {
		t.Errorf("invalid value: expect: %s actual: %s", "Panama", c.EN)
	}
	if c.TH != "ปานามา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปานามา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("png")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 598 {
		t.Errorf("invalid value: expect: %d actual: %d", 598, c.ID)
	}
    if c.EN != "Papua New Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Papua New Guinea", c.EN)
	}
	if c.TH != "ปาปัวนิวกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาปัวนิวกินี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("pry")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 600 {
		t.Errorf("invalid value: expect: %d actual: %d", 600, c.ID)
	}
    if c.EN != "Paraguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Paraguay", c.EN)
	}
	if c.TH != "ปารากวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปารากวัย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("per")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 604 {
		t.Errorf("invalid value: expect: %d actual: %d", 604, c.ID)
	}
    if c.EN != "Peru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Peru", c.EN)
	}
	if c.TH != "เปรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "เปรู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("phl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 608 {
		t.Errorf("invalid value: expect: %d actual: %d", 608, c.ID)
	}
    if c.EN != "Philippines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Philippines", c.EN)
	}
	if c.TH != "ฟิลิปปินส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟิลิปปินส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("pol")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 616 {
		t.Errorf("invalid value: expect: %d actual: %d", 616, c.ID)
	}
    if c.EN != "Poland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Poland", c.EN)
	}
	if c.TH != "โปแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปแลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("prt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 620 {
		t.Errorf("invalid value: expect: %d actual: %d", 620, c.ID)
	}
    if c.EN != "Portugal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Portugal", c.EN)
	}
	if c.TH != "โปรตุเกส" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปรตุเกส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("qat")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 634 {
		t.Errorf("invalid value: expect: %d actual: %d", 634, c.ID)
	}
    if c.EN != "Qatar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Qatar", c.EN)
	}
	if c.TH != "กาตาร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาตาร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("rou")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 642 {
		t.Errorf("invalid value: expect: %d actual: %d", 642, c.ID)
	}
    if c.EN != "Romania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Romania", c.EN)
	}
	if c.TH != "โรมาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โรมาเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("rus")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 643 {
		t.Errorf("invalid value: expect: %d actual: %d", 643, c.ID)
	}
    if c.EN != "Russian Federation" {
		t.Errorf("invalid value: expect: %s actual: %s", "Russian Federation", c.EN)
	}
	if c.TH != "รัสเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "รัสเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("rwa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 646 {
		t.Errorf("invalid value: expect: %d actual: %d", 646, c.ID)
	}
    if c.EN != "Rwanda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Rwanda", c.EN)
	}
	if c.TH != "รวันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "รวันดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("kna")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 659 {
		t.Errorf("invalid value: expect: %d actual: %d", 659, c.ID)
	}
    if c.EN != "Saint Kitts and Nevis" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Kitts and Nevis", c.EN)
	}
	if c.TH != "เซนต์คิตส์และเนวิส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์คิตส์และเนวิส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lca")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 662 {
		t.Errorf("invalid value: expect: %d actual: %d", 662, c.ID)
	}
    if c.EN != "Saint Lucia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Lucia", c.EN)
	}
	if c.TH != "เซนต์ลูเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์ลูเชีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("vct")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 670 {
		t.Errorf("invalid value: expect: %d actual: %d", 670, c.ID)
	}
    if c.EN != "Saint Vincent and the Grenadines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Vincent and the Grenadines", c.EN)
	}
	if c.TH != "เซนต์วินเซนต์และเกรนาดีนส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์วินเซนต์และเกรนาดีนส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("wsm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 882 {
		t.Errorf("invalid value: expect: %d actual: %d", 882, c.ID)
	}
    if c.EN != "Samoa" {
		t.Errorf("invalid value: expect: %s actual: %s", "Samoa", c.EN)
	}
	if c.TH != "ซามัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซามัว", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("smr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 674 {
		t.Errorf("invalid value: expect: %d actual: %d", 674, c.ID)
	}
    if c.EN != "San Marino" {
		t.Errorf("invalid value: expect: %s actual: %s", "San Marino", c.EN)
	}
	if c.TH != "ซานมารีโน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซานมารีโน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("stp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 678 {
		t.Errorf("invalid value: expect: %d actual: %d", 678, c.ID)
	}
    if c.EN != "Sao Tome and Principe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sao Tome and Principe", c.EN)
	}
	if c.TH != "เซาตูเมและปรินซีปี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาตูเมและปรินซีปี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("sau")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 682 {
		t.Errorf("invalid value: expect: %d actual: %d", 682, c.ID)
	}
    if c.EN != "Saudi Arabia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saudi Arabia", c.EN)
	}
	if c.TH != "ซาอุดีอาระเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซาอุดีอาระเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("sen")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 686 {
		t.Errorf("invalid value: expect: %d actual: %d", 686, c.ID)
	}
    if c.EN != "Senegal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Senegal", c.EN)
	}
	if c.TH != "เซเนกัล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเนกัล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("srb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 688 {
		t.Errorf("invalid value: expect: %d actual: %d", 688, c.ID)
	}
    if c.EN != "Serbia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Serbia", c.EN)
	}
	if c.TH != "เซอร์เบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซอร์เบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("syc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 690 {
		t.Errorf("invalid value: expect: %d actual: %d", 690, c.ID)
	}
    if c.EN != "Seychelles" {
		t.Errorf("invalid value: expect: %s actual: %s", "Seychelles", c.EN)
	}
	if c.TH != "เซเชลส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเชลส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("sle")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 694 {
		t.Errorf("invalid value: expect: %d actual: %d", 694, c.ID)
	}
    if c.EN != "Sierra Leone" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sierra Leone", c.EN)
	}
	if c.TH != "เซียร์ราลีโอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซียร์ราลีโอน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("sgp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 702 {
		t.Errorf("invalid value: expect: %d actual: %d", 702, c.ID)
	}
    if c.EN != "Singapore" {
		t.Errorf("invalid value: expect: %s actual: %s", "Singapore", c.EN)
	}
	if c.TH != "สิงคโปร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สิงคโปร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("svk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 703 {
		t.Errorf("invalid value: expect: %d actual: %d", 703, c.ID)
	}
    if c.EN != "Slovakia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovakia", c.EN)
	}
	if c.TH != "สโลวาเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวาเกีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("svn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 705 {
		t.Errorf("invalid value: expect: %d actual: %d", 705, c.ID)
	}
    if c.EN != "Slovenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovenia", c.EN)
	}
	if c.TH != "สโลวีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวีเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("slb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 90 {
		t.Errorf("invalid value: expect: %d actual: %d", 90, c.ID)
	}
    if c.EN != "Solomon Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Solomon Islands", c.EN)
	}
	if c.TH != "หมู่เกาะโซโลมอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะโซโลมอน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("som")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 706 {
		t.Errorf("invalid value: expect: %d actual: %d", 706, c.ID)
	}
    if c.EN != "Somalia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Somalia", c.EN)
	}
	if c.TH != "โซมาเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โซมาเลีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("zaf")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 710 {
		t.Errorf("invalid value: expect: %d actual: %d", 710, c.ID)
	}
    if c.EN != "South Africa" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Africa", c.EN)
	}
	if c.TH != "แอฟริกาใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอฟริกาใต้", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ssd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 728 {
		t.Errorf("invalid value: expect: %d actual: %d", 728, c.ID)
	}
    if c.EN != "South Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Sudan", c.EN)
	}
	if c.TH != "เซาท์ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาท์ซูดาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("esp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 724 {
		t.Errorf("invalid value: expect: %d actual: %d", 724, c.ID)
	}
    if c.EN != "Spain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Spain", c.EN)
	}
	if c.TH != "สเปน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สเปน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("lka")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 144 {
		t.Errorf("invalid value: expect: %d actual: %d", 144, c.ID)
	}
    if c.EN != "Sri Lanka" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sri Lanka", c.EN)
	}
	if c.TH != "ศรีลังกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ศรีลังกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("sdn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 729 {
		t.Errorf("invalid value: expect: %d actual: %d", 729, c.ID)
	}
    if c.EN != "Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sudan", c.EN)
	}
	if c.TH != "ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูดาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("sur")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 740 {
		t.Errorf("invalid value: expect: %d actual: %d", 740, c.ID)
	}
    if c.EN != "Suriname" {
		t.Errorf("invalid value: expect: %s actual: %s", "Suriname", c.EN)
	}
	if c.TH != "ซูรินาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูรินาม", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("swz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 748 {
		t.Errorf("invalid value: expect: %d actual: %d", 748, c.ID)
	}
    if c.EN != "Eswatini" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eswatini", c.EN)
	}
	if c.TH != "เอสวาตีนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสวาตีนี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("swe")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 752 {
		t.Errorf("invalid value: expect: %d actual: %d", 752, c.ID)
	}
    if c.EN != "Sweden" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sweden", c.EN)
	}
	if c.TH != "สวีเดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวีเดน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("che")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 756 {
		t.Errorf("invalid value: expect: %d actual: %d", 756, c.ID)
	}
    if c.EN != "Switzerland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Switzerland", c.EN)
	}
	if c.TH != "สวิตเซอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวิตเซอร์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("syr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 760 {
		t.Errorf("invalid value: expect: %d actual: %d", 760, c.ID)
	}
    if c.EN != "Syrian Arab Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Syrian Arab Republic", c.EN)
	}
	if c.TH != "ซีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tjk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 762 {
		t.Errorf("invalid value: expect: %d actual: %d", 762, c.ID)
	}
    if c.EN != "Tajikistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tajikistan", c.EN)
	}
	if c.TH != "ทาจิกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ทาจิกิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tza")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 834 {
		t.Errorf("invalid value: expect: %d actual: %d", 834, c.ID)
	}
    if c.EN != "Tanzania, United Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tanzania, United Republic of", c.EN)
	}
	if c.TH != "แทนซาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แทนซาเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tha")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 764 {
		t.Errorf("invalid value: expect: %d actual: %d", 764, c.ID)
	}
    if c.EN != "Thailand" {
		t.Errorf("invalid value: expect: %s actual: %s", "Thailand", c.EN)
	}
	if c.TH != "ไทย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไทย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tls")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 626 {
		t.Errorf("invalid value: expect: %d actual: %d", 626, c.ID)
	}
    if c.EN != "Timor-Leste" {
		t.Errorf("invalid value: expect: %s actual: %s", "Timor-Leste", c.EN)
	}
	if c.TH != "ติมอร์-เลสเต" {
		t.Errorf("invalid value: expect: %s actual: %s", "ติมอร์-เลสเต", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tgo")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 768 {
		t.Errorf("invalid value: expect: %d actual: %d", 768, c.ID)
	}
    if c.EN != "Togo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Togo", c.EN)
	}
	if c.TH != "โตโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โตโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ton")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 776 {
		t.Errorf("invalid value: expect: %d actual: %d", 776, c.ID)
	}
    if c.EN != "Tonga" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tonga", c.EN)
	}
	if c.TH != "ตองงา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตองงา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tto")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 780 {
		t.Errorf("invalid value: expect: %d actual: %d", 780, c.ID)
	}
    if c.EN != "Trinidad and Tobago" {
		t.Errorf("invalid value: expect: %s actual: %s", "Trinidad and Tobago", c.EN)
	}
	if c.TH != "ตรินิแดดและโตเบโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตรินิแดดและโตเบโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tun")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 788 {
		t.Errorf("invalid value: expect: %d actual: %d", 788, c.ID)
	}
    if c.EN != "Tunisia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tunisia", c.EN)
	}
	if c.TH != "ตูนิเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูนิเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tur")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 792 {
		t.Errorf("invalid value: expect: %d actual: %d", 792, c.ID)
	}
    if c.EN != "Turkey" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkey", c.EN)
	}
	if c.TH != "ตุรกี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตุรกี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tkm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 795 {
		t.Errorf("invalid value: expect: %d actual: %d", 795, c.ID)
	}
    if c.EN != "Turkmenistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkmenistan", c.EN)
	}
	if c.TH != "เติร์กเมนิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เติร์กเมนิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("tuv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 798 {
		t.Errorf("invalid value: expect: %d actual: %d", 798, c.ID)
	}
    if c.EN != "Tuvalu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tuvalu", c.EN)
	}
	if c.TH != "ตูวาลู" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูวาลู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("uga")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 800 {
		t.Errorf("invalid value: expect: %d actual: %d", 800, c.ID)
	}
    if c.EN != "Uganda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uganda", c.EN)
	}
	if c.TH != "ยูกันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูกันดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ukr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 804 {
		t.Errorf("invalid value: expect: %d actual: %d", 804, c.ID)
	}
    if c.EN != "Ukraine" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ukraine", c.EN)
	}
	if c.TH != "ยูเครน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูเครน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("are")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 784 {
		t.Errorf("invalid value: expect: %d actual: %d", 784, c.ID)
	}
    if c.EN != "United Arab Emirates" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Arab Emirates", c.EN)
	}
	if c.TH != "สหรัฐอาหรับเอมิเรตส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอาหรับเอมิเรตส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("gbr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 826 {
		t.Errorf("invalid value: expect: %d actual: %d", 826, c.ID)
	}
    if c.EN != "United Kingdom of Great Britain and Northern Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Kingdom of Great Britain and Northern Ireland", c.EN)
	}
	if c.TH != "สหราชอาณาจักร" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหราชอาณาจักร", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("usa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 840 {
		t.Errorf("invalid value: expect: %d actual: %d", 840, c.ID)
	}
    if c.EN != "United States of America" {
		t.Errorf("invalid value: expect: %s actual: %s", "United States of America", c.EN)
	}
	if c.TH != "สหรัฐอเมริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอเมริกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ury")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 858 {
		t.Errorf("invalid value: expect: %d actual: %d", 858, c.ID)
	}
    if c.EN != "Uruguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uruguay", c.EN)
	}
	if c.TH != "อุรุกวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุรุกวัย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("uzb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 860 {
		t.Errorf("invalid value: expect: %d actual: %d", 860, c.ID)
	}
    if c.EN != "Uzbekistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uzbekistan", c.EN)
	}
	if c.TH != "อุซเบกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุซเบกิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("vut")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 548 {
		t.Errorf("invalid value: expect: %d actual: %d", 548, c.ID)
	}
    if c.EN != "Vanuatu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Vanuatu", c.EN)
	}
	if c.TH != "วานูอาตู" {
		t.Errorf("invalid value: expect: %s actual: %s", "วานูอาตู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("ven")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 862 {
		t.Errorf("invalid value: expect: %d actual: %d", 862, c.ID)
	}
    if c.EN != "Venezuela (Bolivarian Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Venezuela (Bolivarian Republic of)", c.EN)
	}
	if c.TH != "เวเนซุเอลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวเนซุเอลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("vnm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 704 {
		t.Errorf("invalid value: expect: %d actual: %d", 704, c.ID)
	}
    if c.EN != "Viet Nam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Viet Nam", c.EN)
	}
	if c.TH != "เวียดนาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวียดนาม", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("yem")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 887 {
		t.Errorf("invalid value: expect: %d actual: %d", 887, c.ID)
	}
    if c.EN != "Yemen" {
		t.Errorf("invalid value: expect: %s actual: %s", "Yemen", c.EN)
	}
	if c.TH != "เยเมน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยเมน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("zmb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 894 {
		t.Errorf("invalid value: expect: %d actual: %d", 894, c.ID)
	}
    if c.EN != "Zambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zambia", c.EN)
	}
	if c.TH != "แซมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แซมเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3("zwe")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 716 {
		t.Errorf("invalid value: expect: %d actual: %d", 716, c.ID)
	}
    if c.EN != "Zimbabwe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zimbabwe", c.EN)
	}
	if c.TH != "ซิมบับเว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซิมบับเว", c.TH)
	}

	c, ok = gocountries.GetCountryFromAlpha3("aaa")
	if ok {
		t.Errorf("invalid value: expect: %t actual: %t", false, ok)
	}
    if c.ID != 0 {
		t.Errorf("invalid value: expect: %d actual: %d", 0, c.ID)
	}
    if c.EN != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.EN)
	}
	if c.TH != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.TH)
	}
}

func TestGetCountryFromAlpha3Switch(t *testing.T) {
    var c gocountries.Country
	var ok bool

	c, ok = gocountries.GetCountryFromAlpha3Switch("afg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 4 {
		t.Errorf("invalid value: expect: %d actual: %d", 4, c.ID)
	}
    if c.EN != "Afghanistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Afghanistan", c.EN)
	}
	if c.TH != "อัฟกานิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อัฟกานิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("alb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 8 {
		t.Errorf("invalid value: expect: %d actual: %d", 8, c.ID)
	}
    if c.EN != "Albania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Albania", c.EN)
	}
	if c.TH != "แอลเบเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลเบเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("dza")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 12 {
		t.Errorf("invalid value: expect: %d actual: %d", 12, c.ID)
	}
    if c.EN != "Algeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Algeria", c.EN)
	}
	if c.TH != "แอลจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอลจีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("and")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 20 {
		t.Errorf("invalid value: expect: %d actual: %d", 20, c.ID)
	}
    if c.EN != "Andorra" {
		t.Errorf("invalid value: expect: %s actual: %s", "Andorra", c.EN)
	}
	if c.TH != "อันดอร์รา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อันดอร์รา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ago")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 24 {
		t.Errorf("invalid value: expect: %d actual: %d", 24, c.ID)
	}
    if c.EN != "Angola" {
		t.Errorf("invalid value: expect: %s actual: %s", "Angola", c.EN)
	}
	if c.TH != "แองโกลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แองโกลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("atg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 28 {
		t.Errorf("invalid value: expect: %d actual: %d", 28, c.ID)
	}
    if c.EN != "Antigua and Barbuda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Antigua and Barbuda", c.EN)
	}
	if c.TH != "แอนติกาและบาร์บูดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอนติกาและบาร์บูดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("arg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 32 {
		t.Errorf("invalid value: expect: %d actual: %d", 32, c.ID)
	}
    if c.EN != "Argentina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Argentina", c.EN)
	}
	if c.TH != "อาร์เจนตินา" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์เจนตินา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("arm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 51 {
		t.Errorf("invalid value: expect: %d actual: %d", 51, c.ID)
	}
    if c.EN != "Armenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Armenia", c.EN)
	}
	if c.TH != "อาร์มีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาร์มีเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("aus")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 36 {
		t.Errorf("invalid value: expect: %d actual: %d", 36, c.ID)
	}
    if c.EN != "Australia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Australia", c.EN)
	}
	if c.TH != "ออสเตรเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรเลีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("aut")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 40 {
		t.Errorf("invalid value: expect: %d actual: %d", 40, c.ID)
	}
    if c.EN != "Austria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Austria", c.EN)
	}
	if c.TH != "ออสเตรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ออสเตรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("aze")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 31 {
		t.Errorf("invalid value: expect: %d actual: %d", 31, c.ID)
	}
    if c.EN != "Azerbaijan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Azerbaijan", c.EN)
	}
	if c.TH != "อาเซอร์ไบจาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อาเซอร์ไบจาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bhs")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 44 {
		t.Errorf("invalid value: expect: %d actual: %d", 44, c.ID)
	}
    if c.EN != "Bahamas" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahamas", c.EN)
	}
	if c.TH != "บาฮามาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาฮามาส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bhr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 48 {
		t.Errorf("invalid value: expect: %d actual: %d", 48, c.ID)
	}
    if c.EN != "Bahrain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bahrain", c.EN)
	}
	if c.TH != "บาห์เรน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาห์เรน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bgd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 50 {
		t.Errorf("invalid value: expect: %d actual: %d", 50, c.ID)
	}
    if c.EN != "Bangladesh" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bangladesh", c.EN)
	}
	if c.TH != "บังกลาเทศ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บังกลาเทศ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("brb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 52 {
		t.Errorf("invalid value: expect: %d actual: %d", 52, c.ID)
	}
    if c.EN != "Barbados" {
		t.Errorf("invalid value: expect: %s actual: %s", "Barbados", c.EN)
	}
	if c.TH != "บาร์เบโดส" {
		t.Errorf("invalid value: expect: %s actual: %s", "บาร์เบโดส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("blr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 112 {
		t.Errorf("invalid value: expect: %d actual: %d", 112, c.ID)
	}
    if c.EN != "Belarus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belarus", c.EN)
	}
	if c.TH != "เบลารุส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลารุส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bel")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 56 {
		t.Errorf("invalid value: expect: %d actual: %d", 56, c.ID)
	}
    if c.EN != "Belgium" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belgium", c.EN)
	}
	if c.TH != "เบลเยียม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลเยียม", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("blz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 84 {
		t.Errorf("invalid value: expect: %d actual: %d", 84, c.ID)
	}
    if c.EN != "Belize" {
		t.Errorf("invalid value: expect: %s actual: %s", "Belize", c.EN)
	}
	if c.TH != "เบลีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบลีซ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ben")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 204 {
		t.Errorf("invalid value: expect: %d actual: %d", 204, c.ID)
	}
    if c.EN != "Benin" {
		t.Errorf("invalid value: expect: %s actual: %s", "Benin", c.EN)
	}
	if c.TH != "เบนิน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เบนิน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("btn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 64 {
		t.Errorf("invalid value: expect: %d actual: %d", 64, c.ID)
	}
    if c.EN != "Bhutan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bhutan", c.EN)
	}
	if c.TH != "ภูฏาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ภูฏาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bol")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 68 {
		t.Errorf("invalid value: expect: %d actual: %d", 68, c.ID)
	}
    if c.EN != "Bolivia (Plurinational State of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bolivia (Plurinational State of)", c.EN)
	}
	if c.TH != "โบลิเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โบลิเวีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bih")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 70 {
		t.Errorf("invalid value: expect: %d actual: %d", 70, c.ID)
	}
    if c.EN != "Bosnia and Herzegovina" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bosnia and Herzegovina", c.EN)
	}
	if c.TH != "บอสเนียและเฮอร์เซโกวีนา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอสเนียและเฮอร์เซโกวีนา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bwa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 72 {
		t.Errorf("invalid value: expect: %d actual: %d", 72, c.ID)
	}
    if c.EN != "Botswana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Botswana", c.EN)
	}
	if c.TH != "บอตสวานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "บอตสวานา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bra")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 76 {
		t.Errorf("invalid value: expect: %d actual: %d", 76, c.ID)
	}
    if c.EN != "Brazil" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brazil", c.EN)
	}
	if c.TH != "บราซิล" {
		t.Errorf("invalid value: expect: %s actual: %s", "บราซิล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("brn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 96 {
		t.Errorf("invalid value: expect: %d actual: %d", 96, c.ID)
	}
    if c.EN != "Brunei Darussalam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Brunei Darussalam", c.EN)
	}
	if c.TH != "บรูไน" {
		t.Errorf("invalid value: expect: %s actual: %s", "บรูไน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bgr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 100 {
		t.Errorf("invalid value: expect: %d actual: %d", 100, c.ID)
	}
    if c.EN != "Bulgaria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Bulgaria", c.EN)
	}
	if c.TH != "บัลแกเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "บัลแกเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bfa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 854 {
		t.Errorf("invalid value: expect: %d actual: %d", 854, c.ID)
	}
    if c.EN != "Burkina Faso" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burkina Faso", c.EN)
	}
	if c.TH != "บูร์กินาฟาโซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "บูร์กินาฟาโซ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("bdi")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 108 {
		t.Errorf("invalid value: expect: %d actual: %d", 108, c.ID)
	}
    if c.EN != "Burundi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Burundi", c.EN)
	}
	if c.TH != "บุรุนดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "บุรุนดี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("khm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 116 {
		t.Errorf("invalid value: expect: %d actual: %d", 116, c.ID)
	}
    if c.EN != "Cambodia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cambodia", c.EN)
	}
	if c.TH != "กัมพูชา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัมพูชา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cmr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 120 {
		t.Errorf("invalid value: expect: %d actual: %d", 120, c.ID)
	}
    if c.EN != "Cameroon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cameroon", c.EN)
	}
	if c.TH != "แคเมอรูน" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคเมอรูน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("can")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 124 {
		t.Errorf("invalid value: expect: %d actual: %d", 124, c.ID)
	}
    if c.EN != "Canada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Canada", c.EN)
	}
	if c.TH != "แคนาดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "แคนาดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cpv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 132 {
		t.Errorf("invalid value: expect: %d actual: %d", 132, c.ID)
	}
    if c.EN != "Cabo Verde" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cabo Verde", c.EN)
	}
	if c.TH != "กาบูเวร์ดี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบูเวร์ดี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("caf")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 140 {
		t.Errorf("invalid value: expect: %d actual: %d", 140, c.ID)
	}
    if c.EN != "Central African Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Central African Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐแอฟริกากลาง" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐแอฟริกากลาง", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tcd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 148 {
		t.Errorf("invalid value: expect: %d actual: %d", 148, c.ID)
	}
    if c.EN != "Chad" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chad", c.EN)
	}
	if c.TH != "ชาด" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชาด", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("chl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 152 {
		t.Errorf("invalid value: expect: %d actual: %d", 152, c.ID)
	}
    if c.EN != "Chile" {
		t.Errorf("invalid value: expect: %s actual: %s", "Chile", c.EN)
	}
	if c.TH != "ชิลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ชิลี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("chn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 156 {
		t.Errorf("invalid value: expect: %d actual: %d", 156, c.ID)
	}
    if c.EN != "China" {
		t.Errorf("invalid value: expect: %s actual: %s", "China", c.EN)
	}
	if c.TH != "จีน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จีน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("col")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 170 {
		t.Errorf("invalid value: expect: %d actual: %d", 170, c.ID)
	}
    if c.EN != "Colombia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Colombia", c.EN)
	}
	if c.TH != "โคลอมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โคลอมเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("com")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 174 {
		t.Errorf("invalid value: expect: %d actual: %d", 174, c.ID)
	}
    if c.EN != "Comoros" {
		t.Errorf("invalid value: expect: %s actual: %s", "Comoros", c.EN)
	}
	if c.TH != "คอโมโรส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอโมโรส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cog")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 178 {
		t.Errorf("invalid value: expect: %d actual: %d", 178, c.ID)
	}
    if c.EN != "Congo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo", c.EN)
	}
	if c.TH != "สาธารณรัฐคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐคองโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cod")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 180 {
		t.Errorf("invalid value: expect: %d actual: %d", 180, c.ID)
	}
    if c.EN != "Congo, Democratic Republic of the" {
		t.Errorf("invalid value: expect: %s actual: %s", "Congo, Democratic Republic of the", c.EN)
	}
	if c.TH != "สาธารณรัฐประชาธิปไตยคองโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐประชาธิปไตยคองโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cri")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 188 {
		t.Errorf("invalid value: expect: %d actual: %d", 188, c.ID)
	}
    if c.EN != "Costa Rica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Costa Rica", c.EN)
	}
	if c.TH != "คอสตาริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คอสตาริกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("civ")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 384 {
		t.Errorf("invalid value: expect: %d actual: %d", 384, c.ID)
	}
    if c.EN != "Côte d'Ivoire" {
		t.Errorf("invalid value: expect: %s actual: %s", "Côte d'Ivoire", c.EN)
	}
	if c.TH != "โกตดิวัวร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โกตดิวัวร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("hrv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 191 {
		t.Errorf("invalid value: expect: %d actual: %d", 191, c.ID)
	}
    if c.EN != "Croatia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Croatia", c.EN)
	}
	if c.TH != "โครเอเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โครเอเชีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cub")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 192 {
		t.Errorf("invalid value: expect: %d actual: %d", 192, c.ID)
	}
    if c.EN != "Cuba" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cuba", c.EN)
	}
	if c.TH != "คิวบา" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิวบา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cyp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 196 {
		t.Errorf("invalid value: expect: %d actual: %d", 196, c.ID)
	}
    if c.EN != "Cyprus" {
		t.Errorf("invalid value: expect: %s actual: %s", "Cyprus", c.EN)
	}
	if c.TH != "ไซปรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไซปรัส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("cze")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 203 {
		t.Errorf("invalid value: expect: %d actual: %d", 203, c.ID)
	}
    if c.EN != "Czechia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Czechia", c.EN)
	}
	if c.TH != "เช็กเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เช็กเกีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("dnk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 208 {
		t.Errorf("invalid value: expect: %d actual: %d", 208, c.ID)
	}
    if c.EN != "Denmark" {
		t.Errorf("invalid value: expect: %s actual: %s", "Denmark", c.EN)
	}
	if c.TH != "เดนมาร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เดนมาร์ก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("dji")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 262 {
		t.Errorf("invalid value: expect: %d actual: %d", 262, c.ID)
	}
    if c.EN != "Djibouti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Djibouti", c.EN)
	}
	if c.TH != "จิบูตี" {
		t.Errorf("invalid value: expect: %s actual: %s", "จิบูตี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("dma")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 212 {
		t.Errorf("invalid value: expect: %d actual: %d", 212, c.ID)
	}
    if c.EN != "Dominica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominica", c.EN)
	}
	if c.TH != "ดอมินีกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ดอมินีกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("dom")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 214 {
		t.Errorf("invalid value: expect: %d actual: %d", 214, c.ID)
	}
    if c.EN != "Dominican Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Dominican Republic", c.EN)
	}
	if c.TH != "สาธารณรัฐโดมินิกัน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สาธารณรัฐโดมินิกัน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ecu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 218 {
		t.Errorf("invalid value: expect: %d actual: %d", 218, c.ID)
	}
    if c.EN != "Ecuador" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ecuador", c.EN)
	}
	if c.TH != "เอกวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอกวาดอร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("egy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 818 {
		t.Errorf("invalid value: expect: %d actual: %d", 818, c.ID)
	}
    if c.EN != "Egypt" {
		t.Errorf("invalid value: expect: %s actual: %s", "Egypt", c.EN)
	}
	if c.TH != "อียิปต์" {
		t.Errorf("invalid value: expect: %s actual: %s", "อียิปต์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("slv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 222 {
		t.Errorf("invalid value: expect: %d actual: %d", 222, c.ID)
	}
    if c.EN != "El Salvador" {
		t.Errorf("invalid value: expect: %s actual: %s", "El Salvador", c.EN)
	}
	if c.TH != "เอลซัลวาดอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอลซัลวาดอร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gnq")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 226 {
		t.Errorf("invalid value: expect: %d actual: %d", 226, c.ID)
	}
    if c.EN != "Equatorial Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Equatorial Guinea", c.EN)
	}
	if c.TH != "อิเควทอเรียลกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิเควทอเรียลกินี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("eri")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 232 {
		t.Errorf("invalid value: expect: %d actual: %d", 232, c.ID)
	}
    if c.EN != "Eritrea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eritrea", c.EN)
	}
	if c.TH != "เอริเทรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอริเทรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("est")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 233 {
		t.Errorf("invalid value: expect: %d actual: %d", 233, c.ID)
	}
    if c.EN != "Estonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Estonia", c.EN)
	}
	if c.TH != "เอสโตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสโตเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("eth")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 231 {
		t.Errorf("invalid value: expect: %d actual: %d", 231, c.ID)
	}
    if c.EN != "Ethiopia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ethiopia", c.EN)
	}
	if c.TH != "เอธิโอเปีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอธิโอเปีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("fji")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 242 {
		t.Errorf("invalid value: expect: %d actual: %d", 242, c.ID)
	}
    if c.EN != "Fiji" {
		t.Errorf("invalid value: expect: %s actual: %s", "Fiji", c.EN)
	}
	if c.TH != "ฟีจี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟีจี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("fin")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 246 {
		t.Errorf("invalid value: expect: %d actual: %d", 246, c.ID)
	}
    if c.EN != "Finland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Finland", c.EN)
	}
	if c.TH != "ฟินแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟินแลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("fra")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 250 {
		t.Errorf("invalid value: expect: %d actual: %d", 250, c.ID)
	}
    if c.EN != "France" {
		t.Errorf("invalid value: expect: %s actual: %s", "France", c.EN)
	}
	if c.TH != "ฝรั่งเศส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฝรั่งเศส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gab")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 266 {
		t.Errorf("invalid value: expect: %d actual: %d", 266, c.ID)
	}
    if c.EN != "Gabon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gabon", c.EN)
	}
	if c.TH != "กาบอง" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาบอง", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gmb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 270 {
		t.Errorf("invalid value: expect: %d actual: %d", 270, c.ID)
	}
    if c.EN != "Gambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Gambia", c.EN)
	}
	if c.TH != "แกมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แกมเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("geo")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 268 {
		t.Errorf("invalid value: expect: %d actual: %d", 268, c.ID)
	}
    if c.EN != "Georgia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Georgia", c.EN)
	}
	if c.TH != "จอร์เจีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์เจีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("deu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 276 {
		t.Errorf("invalid value: expect: %d actual: %d", 276, c.ID)
	}
    if c.EN != "Germany" {
		t.Errorf("invalid value: expect: %s actual: %s", "Germany", c.EN)
	}
	if c.TH != "เยอรมนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยอรมนี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gha")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 288 {
		t.Errorf("invalid value: expect: %d actual: %d", 288, c.ID)
	}
    if c.EN != "Ghana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ghana", c.EN)
	}
	if c.TH != "กานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กานา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("grc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 300 {
		t.Errorf("invalid value: expect: %d actual: %d", 300, c.ID)
	}
    if c.EN != "Greece" {
		t.Errorf("invalid value: expect: %s actual: %s", "Greece", c.EN)
	}
	if c.TH != "กรีซ" {
		t.Errorf("invalid value: expect: %s actual: %s", "กรีซ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("grd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 308 {
		t.Errorf("invalid value: expect: %d actual: %d", 308, c.ID)
	}
    if c.EN != "Grenada" {
		t.Errorf("invalid value: expect: %s actual: %s", "Grenada", c.EN)
	}
	if c.TH != "เกรเนดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกรเนดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gtm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 320 {
		t.Errorf("invalid value: expect: %d actual: %d", 320, c.ID)
	}
    if c.EN != "Guatemala" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guatemala", c.EN)
	}
	if c.TH != "กัวเตมาลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กัวเตมาลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gin")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 324 {
		t.Errorf("invalid value: expect: %d actual: %d", 324, c.ID)
	}
    if c.EN != "Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea", c.EN)
	}
	if c.TH != "กินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gnb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 624 {
		t.Errorf("invalid value: expect: %d actual: %d", 624, c.ID)
	}
    if c.EN != "Guinea-Bissau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guinea-Bissau", c.EN)
	}
	if c.TH != "กินี-บิสเซา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กินี-บิสเซา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("guy")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 328 {
		t.Errorf("invalid value: expect: %d actual: %d", 328, c.ID)
	}
    if c.EN != "Guyana" {
		t.Errorf("invalid value: expect: %s actual: %s", "Guyana", c.EN)
	}
	if c.TH != "กายอานา" {
		t.Errorf("invalid value: expect: %s actual: %s", "กายอานา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("hti")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 332 {
		t.Errorf("invalid value: expect: %d actual: %d", 332, c.ID)
	}
    if c.EN != "Haiti" {
		t.Errorf("invalid value: expect: %s actual: %s", "Haiti", c.EN)
	}
	if c.TH != "เฮติ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เฮติ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("hnd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 340 {
		t.Errorf("invalid value: expect: %d actual: %d", 340, c.ID)
	}
    if c.EN != "Honduras" {
		t.Errorf("invalid value: expect: %s actual: %s", "Honduras", c.EN)
	}
	if c.TH != "ฮอนดูรัส" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮอนดูรัส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("hun")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 348 {
		t.Errorf("invalid value: expect: %d actual: %d", 348, c.ID)
	}
    if c.EN != "Hungary" {
		t.Errorf("invalid value: expect: %s actual: %s", "Hungary", c.EN)
	}
	if c.TH != "ฮังการี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฮังการี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("isl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 352 {
		t.Errorf("invalid value: expect: %d actual: %d", 352, c.ID)
	}
    if c.EN != "Iceland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iceland", c.EN)
	}
	if c.TH != "ไอซ์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอซ์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ind")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 356 {
		t.Errorf("invalid value: expect: %d actual: %d", 356, c.ID)
	}
    if c.EN != "India" {
		t.Errorf("invalid value: expect: %s actual: %s", "India", c.EN)
	}
	if c.TH != "อินเดีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินเดีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("idn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 360 {
		t.Errorf("invalid value: expect: %d actual: %d", 360, c.ID)
	}
    if c.EN != "Indonesia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Indonesia", c.EN)
	}
	if c.TH != "อินโดนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อินโดนีเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("irn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 364 {
		t.Errorf("invalid value: expect: %d actual: %d", 364, c.ID)
	}
    if c.EN != "Iran (Islamic Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iran (Islamic Republic of)", c.EN)
	}
	if c.TH != "อิหร่าน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิหร่าน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("irq")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 368 {
		t.Errorf("invalid value: expect: %d actual: %d", 368, c.ID)
	}
    if c.EN != "Iraq" {
		t.Errorf("invalid value: expect: %s actual: %s", "Iraq", c.EN)
	}
	if c.TH != "อิรัก" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิรัก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("irl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 372 {
		t.Errorf("invalid value: expect: %d actual: %d", 372, c.ID)
	}
    if c.EN != "Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ireland", c.EN)
	}
	if c.TH != "ไอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไอร์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("isr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 376 {
		t.Errorf("invalid value: expect: %d actual: %d", 376, c.ID)
	}
    if c.EN != "Israel" {
		t.Errorf("invalid value: expect: %s actual: %s", "Israel", c.EN)
	}
	if c.TH != "อิสราเอล" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิสราเอล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ita")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 380 {
		t.Errorf("invalid value: expect: %d actual: %d", 380, c.ID)
	}
    if c.EN != "Italy" {
		t.Errorf("invalid value: expect: %s actual: %s", "Italy", c.EN)
	}
	if c.TH != "อิตาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "อิตาลี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("jam")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 388 {
		t.Errorf("invalid value: expect: %d actual: %d", 388, c.ID)
	}
    if c.EN != "Jamaica" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jamaica", c.EN)
	}
	if c.TH != "จาเมกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "จาเมกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("jpn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 392 {
		t.Errorf("invalid value: expect: %d actual: %d", 392, c.ID)
	}
    if c.EN != "Japan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Japan", c.EN)
	}
	if c.TH != "ญี่ปุ่น" {
		t.Errorf("invalid value: expect: %s actual: %s", "ญี่ปุ่น", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("jor")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 400 {
		t.Errorf("invalid value: expect: %d actual: %d", 400, c.ID)
	}
    if c.EN != "Jordan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Jordan", c.EN)
	}
	if c.TH != "จอร์แดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "จอร์แดน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("kaz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 398 {
		t.Errorf("invalid value: expect: %d actual: %d", 398, c.ID)
	}
    if c.EN != "Kazakhstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kazakhstan", c.EN)
	}
	if c.TH != "คาซัคสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คาซัคสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ken")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 404 {
		t.Errorf("invalid value: expect: %d actual: %d", 404, c.ID)
	}
    if c.EN != "Kenya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kenya", c.EN)
	}
	if c.TH != "เคนยา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เคนยา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("kir")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 296 {
		t.Errorf("invalid value: expect: %d actual: %d", 296, c.ID)
	}
    if c.EN != "Kiribati" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kiribati", c.EN)
	}
	if c.TH != "คิริบาส" {
		t.Errorf("invalid value: expect: %s actual: %s", "คิริบาส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("prk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 408 {
		t.Errorf("invalid value: expect: %d actual: %d", 408, c.ID)
	}
    if c.EN != "Korea (Democratic People's Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea (Democratic People's Republic of)", c.EN)
	}
	if c.TH != "เกาหลีเหนือ" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีเหนือ", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("kor")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 410 {
		t.Errorf("invalid value: expect: %d actual: %d", 410, c.ID)
	}
    if c.EN != "Korea, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Korea, Republic of", c.EN)
	}
	if c.TH != "เกาหลีใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "เกาหลีใต้", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("kwt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 414 {
		t.Errorf("invalid value: expect: %d actual: %d", 414, c.ID)
	}
    if c.EN != "Kuwait" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kuwait", c.EN)
	}
	if c.TH != "คูเวต" {
		t.Errorf("invalid value: expect: %s actual: %s", "คูเวต", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("kgz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 417 {
		t.Errorf("invalid value: expect: %d actual: %d", 417, c.ID)
	}
    if c.EN != "Kyrgyzstan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Kyrgyzstan", c.EN)
	}
	if c.TH != "คีร์กีซสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "คีร์กีซสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lao")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 418 {
		t.Errorf("invalid value: expect: %d actual: %d", 418, c.ID)
	}
    if c.EN != "Lao People's Democratic Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lao People's Democratic Republic", c.EN)
	}
	if c.TH != "ลาว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลาว", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lva")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 428 {
		t.Errorf("invalid value: expect: %d actual: %d", 428, c.ID)
	}
    if c.EN != "Latvia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Latvia", c.EN)
	}
	if c.TH != "ลัตเวีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลัตเวีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lbn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 422 {
		t.Errorf("invalid value: expect: %d actual: %d", 422, c.ID)
	}
    if c.EN != "Lebanon" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lebanon", c.EN)
	}
	if c.TH != "เลบานอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลบานอน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lso")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 426 {
		t.Errorf("invalid value: expect: %d actual: %d", 426, c.ID)
	}
    if c.EN != "Lesotho" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lesotho", c.EN)
	}
	if c.TH != "เลโซโท" {
		t.Errorf("invalid value: expect: %s actual: %s", "เลโซโท", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lbr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 430 {
		t.Errorf("invalid value: expect: %d actual: %d", 430, c.ID)
	}
    if c.EN != "Liberia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liberia", c.EN)
	}
	if c.TH != "ไลบีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไลบีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lby")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 434 {
		t.Errorf("invalid value: expect: %d actual: %d", 434, c.ID)
	}
    if c.EN != "Libya" {
		t.Errorf("invalid value: expect: %s actual: %s", "Libya", c.EN)
	}
	if c.TH != "ลิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lie")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 438 {
		t.Errorf("invalid value: expect: %d actual: %d", 438, c.ID)
	}
    if c.EN != "Liechtenstein" {
		t.Errorf("invalid value: expect: %s actual: %s", "Liechtenstein", c.EN)
	}
	if c.TH != "ลิกเตนสไตน์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิกเตนสไตน์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ltu")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 440 {
		t.Errorf("invalid value: expect: %d actual: %d", 440, c.ID)
	}
    if c.EN != "Lithuania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Lithuania", c.EN)
	}
	if c.TH != "ลิทัวเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลิทัวเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lux")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 442 {
		t.Errorf("invalid value: expect: %d actual: %d", 442, c.ID)
	}
    if c.EN != "Luxembourg" {
		t.Errorf("invalid value: expect: %s actual: %s", "Luxembourg", c.EN)
	}
	if c.TH != "ลักเซมเบิร์ก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ลักเซมเบิร์ก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mkd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 807 {
		t.Errorf("invalid value: expect: %d actual: %d", 807, c.ID)
	}
    if c.EN != "North Macedonia" {
		t.Errorf("invalid value: expect: %s actual: %s", "North Macedonia", c.EN)
	}
	if c.TH != "นอร์ทมาซิโดเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์ทมาซิโดเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mdg")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 450 {
		t.Errorf("invalid value: expect: %d actual: %d", 450, c.ID)
	}
    if c.EN != "Madagascar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Madagascar", c.EN)
	}
	if c.TH != "มาดากัสการ์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาดากัสการ์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mwi")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 454 {
		t.Errorf("invalid value: expect: %d actual: %d", 454, c.ID)
	}
    if c.EN != "Malawi" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malawi", c.EN)
	}
	if c.TH != "มาลาวี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลาวี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mys")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 458 {
		t.Errorf("invalid value: expect: %d actual: %d", 458, c.ID)
	}
    if c.EN != "Malaysia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malaysia", c.EN)
	}
	if c.TH != "มาเลเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาเลเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mdv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 462 {
		t.Errorf("invalid value: expect: %d actual: %d", 462, c.ID)
	}
    if c.EN != "Maldives" {
		t.Errorf("invalid value: expect: %s actual: %s", "Maldives", c.EN)
	}
	if c.TH != "มัลดีฟส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "มัลดีฟส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mli")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 466 {
		t.Errorf("invalid value: expect: %d actual: %d", 466, c.ID)
	}
    if c.EN != "Mali" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mali", c.EN)
	}
	if c.TH != "มาลี" {
		t.Errorf("invalid value: expect: %s actual: %s", "มาลี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mlt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 470 {
		t.Errorf("invalid value: expect: %d actual: %d", 470, c.ID)
	}
    if c.EN != "Malta" {
		t.Errorf("invalid value: expect: %s actual: %s", "Malta", c.EN)
	}
	if c.TH != "มอลตา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลตา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mhl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 584 {
		t.Errorf("invalid value: expect: %d actual: %d", 584, c.ID)
	}
    if c.EN != "Marshall Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Marshall Islands", c.EN)
	}
	if c.TH != "หมู่เกาะมาร์แชลล์" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะมาร์แชลล์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mrt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 478 {
		t.Errorf("invalid value: expect: %d actual: %d", 478, c.ID)
	}
    if c.EN != "Mauritania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritania", c.EN)
	}
	if c.TH != "มอริเตเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเตเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mus")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 480 {
		t.Errorf("invalid value: expect: %d actual: %d", 480, c.ID)
	}
    if c.EN != "Mauritius" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mauritius", c.EN)
	}
	if c.TH != "มอริเชียส" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอริเชียส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mex")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 484 {
		t.Errorf("invalid value: expect: %d actual: %d", 484, c.ID)
	}
    if c.EN != "Mexico" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mexico", c.EN)
	}
	if c.TH != "เม็กซิโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "เม็กซิโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("fsm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 583 {
		t.Errorf("invalid value: expect: %d actual: %d", 583, c.ID)
	}
    if c.EN != "Micronesia (Federated States of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Micronesia (Federated States of)", c.EN)
	}
	if c.TH != "ไมโครนีเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไมโครนีเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mar")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 504 {
		t.Errorf("invalid value: expect: %d actual: %d", 504, c.ID)
	}
    if c.EN != "Morocco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Morocco", c.EN)
	}
	if c.TH != "โมร็อกโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมร็อกโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mda")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 498 {
		t.Errorf("invalid value: expect: %d actual: %d", 498, c.ID)
	}
    if c.EN != "Moldova, Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Moldova, Republic of", c.EN)
	}
	if c.TH != "มอลโดวา" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอลโดวา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mco")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 492 {
		t.Errorf("invalid value: expect: %d actual: %d", 492, c.ID)
	}
    if c.EN != "Monaco" {
		t.Errorf("invalid value: expect: %s actual: %s", "Monaco", c.EN)
	}
	if c.TH != "โมนาโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมนาโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mng")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 496 {
		t.Errorf("invalid value: expect: %d actual: %d", 496, c.ID)
	}
    if c.EN != "Mongolia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mongolia", c.EN)
	}
	if c.TH != "มองโกเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "มองโกเลีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mne")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 499 {
		t.Errorf("invalid value: expect: %d actual: %d", 499, c.ID)
	}
    if c.EN != "Montenegro" {
		t.Errorf("invalid value: expect: %s actual: %s", "Montenegro", c.EN)
	}
	if c.TH != "มอนเตเนโกร" {
		t.Errorf("invalid value: expect: %s actual: %s", "มอนเตเนโกร", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("moz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 508 {
		t.Errorf("invalid value: expect: %d actual: %d", 508, c.ID)
	}
    if c.EN != "Mozambique" {
		t.Errorf("invalid value: expect: %s actual: %s", "Mozambique", c.EN)
	}
	if c.TH != "โมซัมบิก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โมซัมบิก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("mmr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 104 {
		t.Errorf("invalid value: expect: %d actual: %d", 104, c.ID)
	}
    if c.EN != "Myanmar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Myanmar", c.EN)
	}
	if c.TH != "พม่า" {
		t.Errorf("invalid value: expect: %s actual: %s", "พม่า", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("nam")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 516 {
		t.Errorf("invalid value: expect: %d actual: %d", 516, c.ID)
	}
    if c.EN != "Namibia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Namibia", c.EN)
	}
	if c.TH != "นามิเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "นามิเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("nru")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 520 {
		t.Errorf("invalid value: expect: %d actual: %d", 520, c.ID)
	}
    if c.EN != "Nauru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nauru", c.EN)
	}
	if c.TH != "นาอูรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "นาอูรู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("npl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 524 {
		t.Errorf("invalid value: expect: %d actual: %d", 524, c.ID)
	}
    if c.EN != "Nepal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nepal", c.EN)
	}
	if c.TH != "เนปาล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนปาล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("nld")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 528 {
		t.Errorf("invalid value: expect: %d actual: %d", 528, c.ID)
	}
    if c.EN != "Netherlands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Netherlands", c.EN)
	}
	if c.TH != "เนเธอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เนเธอร์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("nzl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 554 {
		t.Errorf("invalid value: expect: %d actual: %d", 554, c.ID)
	}
    if c.EN != "New Zealand" {
		t.Errorf("invalid value: expect: %s actual: %s", "New Zealand", c.EN)
	}
	if c.TH != "นิวซีแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิวซีแลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("nic")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 558 {
		t.Errorf("invalid value: expect: %d actual: %d", 558, c.ID)
	}
    if c.EN != "Nicaragua" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nicaragua", c.EN)
	}
	if c.TH != "นิการากัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "นิการากัว", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ner")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 562 {
		t.Errorf("invalid value: expect: %d actual: %d", 562, c.ID)
	}
    if c.EN != "Niger" {
		t.Errorf("invalid value: expect: %s actual: %s", "Niger", c.EN)
	}
	if c.TH != "ไนเจอร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนเจอร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("nga")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 566 {
		t.Errorf("invalid value: expect: %d actual: %d", 566, c.ID)
	}
    if c.EN != "Nigeria" {
		t.Errorf("invalid value: expect: %s actual: %s", "Nigeria", c.EN)
	}
	if c.TH != "ไนจีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไนจีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("nor")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 578 {
		t.Errorf("invalid value: expect: %d actual: %d", 578, c.ID)
	}
    if c.EN != "Norway" {
		t.Errorf("invalid value: expect: %s actual: %s", "Norway", c.EN)
	}
	if c.TH != "นอร์เวย์" {
		t.Errorf("invalid value: expect: %s actual: %s", "นอร์เวย์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("omn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 512 {
		t.Errorf("invalid value: expect: %d actual: %d", 512, c.ID)
	}
    if c.EN != "Oman" {
		t.Errorf("invalid value: expect: %s actual: %s", "Oman", c.EN)
	}
	if c.TH != "โอมาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "โอมาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("pak")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 586 {
		t.Errorf("invalid value: expect: %d actual: %d", 586, c.ID)
	}
    if c.EN != "Pakistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Pakistan", c.EN)
	}
	if c.TH != "ปากีสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปากีสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("plw")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 585 {
		t.Errorf("invalid value: expect: %d actual: %d", 585, c.ID)
	}
    if c.EN != "Palau" {
		t.Errorf("invalid value: expect: %s actual: %s", "Palau", c.EN)
	}
	if c.TH != "ปาเลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาเลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("pan")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 591 {
		t.Errorf("invalid value: expect: %d actual: %d", 591, c.ID)
	}
    if c.EN != "Panama" {
		t.Errorf("invalid value: expect: %s actual: %s", "Panama", c.EN)
	}
	if c.TH != "ปานามา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปานามา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("png")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 598 {
		t.Errorf("invalid value: expect: %d actual: %d", 598, c.ID)
	}
    if c.EN != "Papua New Guinea" {
		t.Errorf("invalid value: expect: %s actual: %s", "Papua New Guinea", c.EN)
	}
	if c.TH != "ปาปัวนิวกินี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปาปัวนิวกินี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("pry")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 600 {
		t.Errorf("invalid value: expect: %d actual: %d", 600, c.ID)
	}
    if c.EN != "Paraguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Paraguay", c.EN)
	}
	if c.TH != "ปารากวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ปารากวัย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("per")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 604 {
		t.Errorf("invalid value: expect: %d actual: %d", 604, c.ID)
	}
    if c.EN != "Peru" {
		t.Errorf("invalid value: expect: %s actual: %s", "Peru", c.EN)
	}
	if c.TH != "เปรู" {
		t.Errorf("invalid value: expect: %s actual: %s", "เปรู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("phl")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 608 {
		t.Errorf("invalid value: expect: %d actual: %d", 608, c.ID)
	}
    if c.EN != "Philippines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Philippines", c.EN)
	}
	if c.TH != "ฟิลิปปินส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "ฟิลิปปินส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("pol")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 616 {
		t.Errorf("invalid value: expect: %d actual: %d", 616, c.ID)
	}
    if c.EN != "Poland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Poland", c.EN)
	}
	if c.TH != "โปแลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปแลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("prt")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 620 {
		t.Errorf("invalid value: expect: %d actual: %d", 620, c.ID)
	}
    if c.EN != "Portugal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Portugal", c.EN)
	}
	if c.TH != "โปรตุเกส" {
		t.Errorf("invalid value: expect: %s actual: %s", "โปรตุเกส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("qat")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 634 {
		t.Errorf("invalid value: expect: %d actual: %d", 634, c.ID)
	}
    if c.EN != "Qatar" {
		t.Errorf("invalid value: expect: %s actual: %s", "Qatar", c.EN)
	}
	if c.TH != "กาตาร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "กาตาร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("rou")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 642 {
		t.Errorf("invalid value: expect: %d actual: %d", 642, c.ID)
	}
    if c.EN != "Romania" {
		t.Errorf("invalid value: expect: %s actual: %s", "Romania", c.EN)
	}
	if c.TH != "โรมาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โรมาเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("rus")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 643 {
		t.Errorf("invalid value: expect: %d actual: %d", 643, c.ID)
	}
    if c.EN != "Russian Federation" {
		t.Errorf("invalid value: expect: %s actual: %s", "Russian Federation", c.EN)
	}
	if c.TH != "รัสเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "รัสเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("rwa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 646 {
		t.Errorf("invalid value: expect: %d actual: %d", 646, c.ID)
	}
    if c.EN != "Rwanda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Rwanda", c.EN)
	}
	if c.TH != "รวันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "รวันดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("kna")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 659 {
		t.Errorf("invalid value: expect: %d actual: %d", 659, c.ID)
	}
    if c.EN != "Saint Kitts and Nevis" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Kitts and Nevis", c.EN)
	}
	if c.TH != "เซนต์คิตส์และเนวิส" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์คิตส์และเนวิส", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lca")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 662 {
		t.Errorf("invalid value: expect: %d actual: %d", 662, c.ID)
	}
    if c.EN != "Saint Lucia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Lucia", c.EN)
	}
	if c.TH != "เซนต์ลูเชีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์ลูเชีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("vct")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 670 {
		t.Errorf("invalid value: expect: %d actual: %d", 670, c.ID)
	}
    if c.EN != "Saint Vincent and the Grenadines" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saint Vincent and the Grenadines", c.EN)
	}
	if c.TH != "เซนต์วินเซนต์และเกรนาดีนส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซนต์วินเซนต์และเกรนาดีนส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("wsm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 882 {
		t.Errorf("invalid value: expect: %d actual: %d", 882, c.ID)
	}
    if c.EN != "Samoa" {
		t.Errorf("invalid value: expect: %s actual: %s", "Samoa", c.EN)
	}
	if c.TH != "ซามัว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซามัว", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("smr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 674 {
		t.Errorf("invalid value: expect: %d actual: %d", 674, c.ID)
	}
    if c.EN != "San Marino" {
		t.Errorf("invalid value: expect: %s actual: %s", "San Marino", c.EN)
	}
	if c.TH != "ซานมารีโน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซานมารีโน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("stp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 678 {
		t.Errorf("invalid value: expect: %d actual: %d", 678, c.ID)
	}
    if c.EN != "Sao Tome and Principe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sao Tome and Principe", c.EN)
	}
	if c.TH != "เซาตูเมและปรินซีปี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาตูเมและปรินซีปี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("sau")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 682 {
		t.Errorf("invalid value: expect: %d actual: %d", 682, c.ID)
	}
    if c.EN != "Saudi Arabia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Saudi Arabia", c.EN)
	}
	if c.TH != "ซาอุดีอาระเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซาอุดีอาระเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("sen")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 686 {
		t.Errorf("invalid value: expect: %d actual: %d", 686, c.ID)
	}
    if c.EN != "Senegal" {
		t.Errorf("invalid value: expect: %s actual: %s", "Senegal", c.EN)
	}
	if c.TH != "เซเนกัล" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเนกัล", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("srb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 688 {
		t.Errorf("invalid value: expect: %d actual: %d", 688, c.ID)
	}
    if c.EN != "Serbia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Serbia", c.EN)
	}
	if c.TH != "เซอร์เบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซอร์เบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("syc")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 690 {
		t.Errorf("invalid value: expect: %d actual: %d", 690, c.ID)
	}
    if c.EN != "Seychelles" {
		t.Errorf("invalid value: expect: %s actual: %s", "Seychelles", c.EN)
	}
	if c.TH != "เซเชลส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซเชลส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("sle")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 694 {
		t.Errorf("invalid value: expect: %d actual: %d", 694, c.ID)
	}
    if c.EN != "Sierra Leone" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sierra Leone", c.EN)
	}
	if c.TH != "เซียร์ราลีโอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซียร์ราลีโอน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("sgp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 702 {
		t.Errorf("invalid value: expect: %d actual: %d", 702, c.ID)
	}
    if c.EN != "Singapore" {
		t.Errorf("invalid value: expect: %s actual: %s", "Singapore", c.EN)
	}
	if c.TH != "สิงคโปร์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สิงคโปร์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("svk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 703 {
		t.Errorf("invalid value: expect: %d actual: %d", 703, c.ID)
	}
    if c.EN != "Slovakia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovakia", c.EN)
	}
	if c.TH != "สโลวาเกีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวาเกีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("svn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 705 {
		t.Errorf("invalid value: expect: %d actual: %d", 705, c.ID)
	}
    if c.EN != "Slovenia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Slovenia", c.EN)
	}
	if c.TH != "สโลวีเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "สโลวีเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("slb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 90 {
		t.Errorf("invalid value: expect: %d actual: %d", 90, c.ID)
	}
    if c.EN != "Solomon Islands" {
		t.Errorf("invalid value: expect: %s actual: %s", "Solomon Islands", c.EN)
	}
	if c.TH != "หมู่เกาะโซโลมอน" {
		t.Errorf("invalid value: expect: %s actual: %s", "หมู่เกาะโซโลมอน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("som")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 706 {
		t.Errorf("invalid value: expect: %d actual: %d", 706, c.ID)
	}
    if c.EN != "Somalia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Somalia", c.EN)
	}
	if c.TH != "โซมาเลีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "โซมาเลีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("zaf")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 710 {
		t.Errorf("invalid value: expect: %d actual: %d", 710, c.ID)
	}
    if c.EN != "South Africa" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Africa", c.EN)
	}
	if c.TH != "แอฟริกาใต้" {
		t.Errorf("invalid value: expect: %s actual: %s", "แอฟริกาใต้", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ssd")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 728 {
		t.Errorf("invalid value: expect: %d actual: %d", 728, c.ID)
	}
    if c.EN != "South Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "South Sudan", c.EN)
	}
	if c.TH != "เซาท์ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เซาท์ซูดาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("esp")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 724 {
		t.Errorf("invalid value: expect: %d actual: %d", 724, c.ID)
	}
    if c.EN != "Spain" {
		t.Errorf("invalid value: expect: %s actual: %s", "Spain", c.EN)
	}
	if c.TH != "สเปน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สเปน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("lka")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 144 {
		t.Errorf("invalid value: expect: %d actual: %d", 144, c.ID)
	}
    if c.EN != "Sri Lanka" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sri Lanka", c.EN)
	}
	if c.TH != "ศรีลังกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ศรีลังกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("sdn")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 729 {
		t.Errorf("invalid value: expect: %d actual: %d", 729, c.ID)
	}
    if c.EN != "Sudan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sudan", c.EN)
	}
	if c.TH != "ซูดาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูดาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("sur")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 740 {
		t.Errorf("invalid value: expect: %d actual: %d", 740, c.ID)
	}
    if c.EN != "Suriname" {
		t.Errorf("invalid value: expect: %s actual: %s", "Suriname", c.EN)
	}
	if c.TH != "ซูรินาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซูรินาม", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("swz")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 748 {
		t.Errorf("invalid value: expect: %d actual: %d", 748, c.ID)
	}
    if c.EN != "Eswatini" {
		t.Errorf("invalid value: expect: %s actual: %s", "Eswatini", c.EN)
	}
	if c.TH != "เอสวาตีนี" {
		t.Errorf("invalid value: expect: %s actual: %s", "เอสวาตีนี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("swe")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 752 {
		t.Errorf("invalid value: expect: %d actual: %d", 752, c.ID)
	}
    if c.EN != "Sweden" {
		t.Errorf("invalid value: expect: %s actual: %s", "Sweden", c.EN)
	}
	if c.TH != "สวีเดน" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวีเดน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("che")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 756 {
		t.Errorf("invalid value: expect: %d actual: %d", 756, c.ID)
	}
    if c.EN != "Switzerland" {
		t.Errorf("invalid value: expect: %s actual: %s", "Switzerland", c.EN)
	}
	if c.TH != "สวิตเซอร์แลนด์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สวิตเซอร์แลนด์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("syr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 760 {
		t.Errorf("invalid value: expect: %d actual: %d", 760, c.ID)
	}
    if c.EN != "Syrian Arab Republic" {
		t.Errorf("invalid value: expect: %s actual: %s", "Syrian Arab Republic", c.EN)
	}
	if c.TH != "ซีเรีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซีเรีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tjk")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 762 {
		t.Errorf("invalid value: expect: %d actual: %d", 762, c.ID)
	}
    if c.EN != "Tajikistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tajikistan", c.EN)
	}
	if c.TH != "ทาจิกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ทาจิกิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tza")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 834 {
		t.Errorf("invalid value: expect: %d actual: %d", 834, c.ID)
	}
    if c.EN != "Tanzania, United Republic of" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tanzania, United Republic of", c.EN)
	}
	if c.TH != "แทนซาเนีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แทนซาเนีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tha")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 764 {
		t.Errorf("invalid value: expect: %d actual: %d", 764, c.ID)
	}
    if c.EN != "Thailand" {
		t.Errorf("invalid value: expect: %s actual: %s", "Thailand", c.EN)
	}
	if c.TH != "ไทย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ไทย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tls")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 626 {
		t.Errorf("invalid value: expect: %d actual: %d", 626, c.ID)
	}
    if c.EN != "Timor-Leste" {
		t.Errorf("invalid value: expect: %s actual: %s", "Timor-Leste", c.EN)
	}
	if c.TH != "ติมอร์-เลสเต" {
		t.Errorf("invalid value: expect: %s actual: %s", "ติมอร์-เลสเต", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tgo")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 768 {
		t.Errorf("invalid value: expect: %d actual: %d", 768, c.ID)
	}
    if c.EN != "Togo" {
		t.Errorf("invalid value: expect: %s actual: %s", "Togo", c.EN)
	}
	if c.TH != "โตโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "โตโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ton")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 776 {
		t.Errorf("invalid value: expect: %d actual: %d", 776, c.ID)
	}
    if c.EN != "Tonga" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tonga", c.EN)
	}
	if c.TH != "ตองงา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตองงา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tto")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 780 {
		t.Errorf("invalid value: expect: %d actual: %d", 780, c.ID)
	}
    if c.EN != "Trinidad and Tobago" {
		t.Errorf("invalid value: expect: %s actual: %s", "Trinidad and Tobago", c.EN)
	}
	if c.TH != "ตรินิแดดและโตเบโก" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตรินิแดดและโตเบโก", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tun")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 788 {
		t.Errorf("invalid value: expect: %d actual: %d", 788, c.ID)
	}
    if c.EN != "Tunisia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tunisia", c.EN)
	}
	if c.TH != "ตูนิเซีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูนิเซีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tur")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 792 {
		t.Errorf("invalid value: expect: %d actual: %d", 792, c.ID)
	}
    if c.EN != "Turkey" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkey", c.EN)
	}
	if c.TH != "ตุรกี" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตุรกี", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tkm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 795 {
		t.Errorf("invalid value: expect: %d actual: %d", 795, c.ID)
	}
    if c.EN != "Turkmenistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Turkmenistan", c.EN)
	}
	if c.TH != "เติร์กเมนิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เติร์กเมนิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("tuv")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 798 {
		t.Errorf("invalid value: expect: %d actual: %d", 798, c.ID)
	}
    if c.EN != "Tuvalu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Tuvalu", c.EN)
	}
	if c.TH != "ตูวาลู" {
		t.Errorf("invalid value: expect: %s actual: %s", "ตูวาลู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("uga")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 800 {
		t.Errorf("invalid value: expect: %d actual: %d", 800, c.ID)
	}
    if c.EN != "Uganda" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uganda", c.EN)
	}
	if c.TH != "ยูกันดา" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูกันดา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ukr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 804 {
		t.Errorf("invalid value: expect: %d actual: %d", 804, c.ID)
	}
    if c.EN != "Ukraine" {
		t.Errorf("invalid value: expect: %s actual: %s", "Ukraine", c.EN)
	}
	if c.TH != "ยูเครน" {
		t.Errorf("invalid value: expect: %s actual: %s", "ยูเครน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("are")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 784 {
		t.Errorf("invalid value: expect: %d actual: %d", 784, c.ID)
	}
    if c.EN != "United Arab Emirates" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Arab Emirates", c.EN)
	}
	if c.TH != "สหรัฐอาหรับเอมิเรตส์" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอาหรับเอมิเรตส์", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("gbr")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 826 {
		t.Errorf("invalid value: expect: %d actual: %d", 826, c.ID)
	}
    if c.EN != "United Kingdom of Great Britain and Northern Ireland" {
		t.Errorf("invalid value: expect: %s actual: %s", "United Kingdom of Great Britain and Northern Ireland", c.EN)
	}
	if c.TH != "สหราชอาณาจักร" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหราชอาณาจักร", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("usa")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 840 {
		t.Errorf("invalid value: expect: %d actual: %d", 840, c.ID)
	}
    if c.EN != "United States of America" {
		t.Errorf("invalid value: expect: %s actual: %s", "United States of America", c.EN)
	}
	if c.TH != "สหรัฐอเมริกา" {
		t.Errorf("invalid value: expect: %s actual: %s", "สหรัฐอเมริกา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ury")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 858 {
		t.Errorf("invalid value: expect: %d actual: %d", 858, c.ID)
	}
    if c.EN != "Uruguay" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uruguay", c.EN)
	}
	if c.TH != "อุรุกวัย" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุรุกวัย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("uzb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 860 {
		t.Errorf("invalid value: expect: %d actual: %d", 860, c.ID)
	}
    if c.EN != "Uzbekistan" {
		t.Errorf("invalid value: expect: %s actual: %s", "Uzbekistan", c.EN)
	}
	if c.TH != "อุซเบกิสถาน" {
		t.Errorf("invalid value: expect: %s actual: %s", "อุซเบกิสถาน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("vut")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 548 {
		t.Errorf("invalid value: expect: %d actual: %d", 548, c.ID)
	}
    if c.EN != "Vanuatu" {
		t.Errorf("invalid value: expect: %s actual: %s", "Vanuatu", c.EN)
	}
	if c.TH != "วานูอาตู" {
		t.Errorf("invalid value: expect: %s actual: %s", "วานูอาตู", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("ven")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 862 {
		t.Errorf("invalid value: expect: %d actual: %d", 862, c.ID)
	}
    if c.EN != "Venezuela (Bolivarian Republic of)" {
		t.Errorf("invalid value: expect: %s actual: %s", "Venezuela (Bolivarian Republic of)", c.EN)
	}
	if c.TH != "เวเนซุเอลา" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวเนซุเอลา", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("vnm")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 704 {
		t.Errorf("invalid value: expect: %d actual: %d", 704, c.ID)
	}
    if c.EN != "Viet Nam" {
		t.Errorf("invalid value: expect: %s actual: %s", "Viet Nam", c.EN)
	}
	if c.TH != "เวียดนาม" {
		t.Errorf("invalid value: expect: %s actual: %s", "เวียดนาม", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("yem")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 887 {
		t.Errorf("invalid value: expect: %d actual: %d", 887, c.ID)
	}
    if c.EN != "Yemen" {
		t.Errorf("invalid value: expect: %s actual: %s", "Yemen", c.EN)
	}
	if c.TH != "เยเมน" {
		t.Errorf("invalid value: expect: %s actual: %s", "เยเมน", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("zmb")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 894 {
		t.Errorf("invalid value: expect: %d actual: %d", 894, c.ID)
	}
    if c.EN != "Zambia" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zambia", c.EN)
	}
	if c.TH != "แซมเบีย" {
		t.Errorf("invalid value: expect: %s actual: %s", "แซมเบีย", c.TH)
	}
	c, ok = gocountries.GetCountryFromAlpha3Switch("zwe")
	if !ok {
		t.Errorf("invalid value: expect: %t actual: %t", true, ok)
	}
    if c.ID != 716 {
		t.Errorf("invalid value: expect: %d actual: %d", 716, c.ID)
	}
    if c.EN != "Zimbabwe" {
		t.Errorf("invalid value: expect: %s actual: %s", "Zimbabwe", c.EN)
	}
	if c.TH != "ซิมบับเว" {
		t.Errorf("invalid value: expect: %s actual: %s", "ซิมบับเว", c.TH)
	}

	c, ok = gocountries.GetCountryFromAlpha3Switch("aaa")
	if ok {
		t.Errorf("invalid value: expect: %t actual: %t", false, ok)
	}
    if c.ID != 0 {
		t.Errorf("invalid value: expect: %d actual: %d", 0, c.ID)
	}
    if c.EN != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.EN)
	}
	if c.TH != "" {
		t.Errorf("invalid value: expect: %s actual: %s", "", c.TH)
	}
}

func BenchmarkGetCountryFromID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gocountries.GetCountryFromID(input[n%len(input)].ID)
	}
}

func BenchmarkGetCountryFromAlpha2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gocountries.GetCountryFromAlpha2(input[n%len(input)].Alpha2)
	}
}

func BenchmarkGetCountryFromAlpha3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gocountries.GetCountryFromAlpha3(input[n%len(input)].Alpha3)
	}
}

func BenchmarkGetCountryFromAlpha3Switch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gocountries.GetCountryFromAlpha3Switch(input[n%len(input)].Alpha3)
	}
}
