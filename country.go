package gocountries

type Country struct {
    Alpha3 string
    Alpha2 string
    ID int32
    TH string
    EN string
}

func GetCountryFromAlpha3(alpha3 string) (Country) {
    switch alpha3 {
        case "afg": 
            return Country{Alpha2: "af", Alpha3: "afg", ID: 4, EN:"Afghanistan", TH:"อัฟกานิสถาน"}
        case "alb": 
            return Country{Alpha2: "al", Alpha3: "alb", ID: 8, EN:"Albania", TH:"แอลเบเนีย"}
        case "dza": 
            return Country{Alpha2: "dz", Alpha3: "dza", ID: 12, EN:"Algeria", TH:"แอลจีเรีย"}
        case "and": 
            return Country{Alpha2: "ad", Alpha3: "and", ID: 20, EN:"Andorra", TH:"อันดอร์รา"}
        case "ago": 
            return Country{Alpha2: "ao", Alpha3: "ago", ID: 24, EN:"Angola", TH:"แองโกลา"}
        case "atg": 
            return Country{Alpha2: "ag", Alpha3: "atg", ID: 28, EN:"Antigua and Barbuda", TH:"แอนติกาและบาร์บูดา"}
        case "arg": 
            return Country{Alpha2: "ar", Alpha3: "arg", ID: 32, EN:"Argentina", TH:"อาร์เจนตินา"}
        case "arm": 
            return Country{Alpha2: "am", Alpha3: "arm", ID: 51, EN:"Armenia", TH:"อาร์มีเนีย"}
        case "aus": 
            return Country{Alpha2: "au", Alpha3: "aus", ID: 36, EN:"Australia", TH:"ออสเตรเลีย"}
        case "aut": 
            return Country{Alpha2: "at", Alpha3: "aut", ID: 40, EN:"Austria", TH:"ออสเตรีย"}
        case "aze": 
            return Country{Alpha2: "az", Alpha3: "aze", ID: 31, EN:"Azerbaijan", TH:"อาเซอร์ไบจาน"}
        case "bhs": 
            return Country{Alpha2: "bs", Alpha3: "bhs", ID: 44, EN:"Bahamas", TH:"บาฮามาส"}
        case "bhr": 
            return Country{Alpha2: "bh", Alpha3: "bhr", ID: 48, EN:"Bahrain", TH:"บาห์เรน"}
        case "bgd": 
            return Country{Alpha2: "bd", Alpha3: "bgd", ID: 50, EN:"Bangladesh", TH:"บังกลาเทศ"}
        case "brb": 
            return Country{Alpha2: "bb", Alpha3: "brb", ID: 52, EN:"Barbados", TH:"บาร์เบโดส"}
        case "blr": 
            return Country{Alpha2: "by", Alpha3: "blr", ID: 112, EN:"Belarus", TH:"เบลารุส"}
        case "bel": 
            return Country{Alpha2: "be", Alpha3: "bel", ID: 56, EN:"Belgium", TH:"เบลเยียม"}
        case "blz": 
            return Country{Alpha2: "bz", Alpha3: "blz", ID: 84, EN:"Belize", TH:"เบลีซ"}
        case "ben": 
            return Country{Alpha2: "bj", Alpha3: "ben", ID: 204, EN:"Benin", TH:"เบนิน"}
        case "btn": 
            return Country{Alpha2: "bt", Alpha3: "btn", ID: 64, EN:"Bhutan", TH:"ภูฏาน"}
        case "bol": 
            return Country{Alpha2: "bo", Alpha3: "bol", ID: 68, EN:"Bolivia (Plurinational State of)", TH:"โบลิเวีย"}
        case "bih": 
            return Country{Alpha2: "ba", Alpha3: "bih", ID: 70, EN:"Bosnia and Herzegovina", TH:"บอสเนียและเฮอร์เซโกวีนา"}
        case "bwa": 
            return Country{Alpha2: "bw", Alpha3: "bwa", ID: 72, EN:"Botswana", TH:"บอตสวานา"}
        case "bra": 
            return Country{Alpha2: "br", Alpha3: "bra", ID: 76, EN:"Brazil", TH:"บราซิล"}
        case "brn": 
            return Country{Alpha2: "bn", Alpha3: "brn", ID: 96, EN:"Brunei Darussalam", TH:"บรูไน"}
        case "bgr": 
            return Country{Alpha2: "bg", Alpha3: "bgr", ID: 100, EN:"Bulgaria", TH:"บัลแกเรีย"}
        case "bfa": 
            return Country{Alpha2: "bf", Alpha3: "bfa", ID: 854, EN:"Burkina Faso", TH:"บูร์กินาฟาโซ"}
        case "bdi": 
            return Country{Alpha2: "bi", Alpha3: "bdi", ID: 108, EN:"Burundi", TH:"บุรุนดี"}
        case "khm": 
            return Country{Alpha2: "kh", Alpha3: "khm", ID: 116, EN:"Cambodia", TH:"กัมพูชา"}
        case "cmr": 
            return Country{Alpha2: "cm", Alpha3: "cmr", ID: 120, EN:"Cameroon", TH:"แคเมอรูน"}
        case "can": 
            return Country{Alpha2: "ca", Alpha3: "can", ID: 124, EN:"Canada", TH:"แคนาดา"}
        case "cpv": 
            return Country{Alpha2: "cv", Alpha3: "cpv", ID: 132, EN:"Cabo Verde", TH:"กาบูเวร์ดี"}
        case "caf": 
            return Country{Alpha2: "cf", Alpha3: "caf", ID: 140, EN:"Central African Republic", TH:"สาธารณรัฐแอฟริกากลาง"}
        case "tcd": 
            return Country{Alpha2: "td", Alpha3: "tcd", ID: 148, EN:"Chad", TH:"ชาด"}
        case "chl": 
            return Country{Alpha2: "cl", Alpha3: "chl", ID: 152, EN:"Chile", TH:"ชิลี"}
        case "chn": 
            return Country{Alpha2: "cn", Alpha3: "chn", ID: 156, EN:"China", TH:"จีน"}
        case "col": 
            return Country{Alpha2: "co", Alpha3: "col", ID: 170, EN:"Colombia", TH:"โคลอมเบีย"}
        case "com": 
            return Country{Alpha2: "km", Alpha3: "com", ID: 174, EN:"Comoros", TH:"คอโมโรส"}
        case "cog": 
            return Country{Alpha2: "cg", Alpha3: "cog", ID: 178, EN:"Congo", TH:"สาธารณรัฐคองโก"}
        case "cod": 
            return Country{Alpha2: "cd", Alpha3: "cod", ID: 180, EN:"Congo, Democratic Republic of the", TH:"สาธารณรัฐประชาธิปไตยคองโก"}
        case "cri": 
            return Country{Alpha2: "cr", Alpha3: "cri", ID: 188, EN:"Costa Rica", TH:"คอสตาริกา"}
        case "civ": 
            return Country{Alpha2: "ci", Alpha3: "civ", ID: 384, EN:"Côte d'Ivoire", TH:"โกตดิวัวร์"}
        case "hrv": 
            return Country{Alpha2: "hr", Alpha3: "hrv", ID: 191, EN:"Croatia", TH:"โครเอเชีย"}
        case "cub": 
            return Country{Alpha2: "cu", Alpha3: "cub", ID: 192, EN:"Cuba", TH:"คิวบา"}
        case "cyp": 
            return Country{Alpha2: "cy", Alpha3: "cyp", ID: 196, EN:"Cyprus", TH:"ไซปรัส"}
        case "cze": 
            return Country{Alpha2: "cz", Alpha3: "cze", ID: 203, EN:"Czechia", TH:"เช็กเกีย"}
        case "dnk": 
            return Country{Alpha2: "dk", Alpha3: "dnk", ID: 208, EN:"Denmark", TH:"เดนมาร์ก"}
        case "dji": 
            return Country{Alpha2: "dj", Alpha3: "dji", ID: 262, EN:"Djibouti", TH:"จิบูตี"}
        case "dma": 
            return Country{Alpha2: "dm", Alpha3: "dma", ID: 212, EN:"Dominica", TH:"ดอมินีกา"}
        case "dom": 
            return Country{Alpha2: "do", Alpha3: "dom", ID: 214, EN:"Dominican Republic", TH:"สาธารณรัฐโดมินิกัน"}
        case "ecu": 
            return Country{Alpha2: "ec", Alpha3: "ecu", ID: 218, EN:"Ecuador", TH:"เอกวาดอร์"}
        case "egy": 
            return Country{Alpha2: "eg", Alpha3: "egy", ID: 818, EN:"Egypt", TH:"อียิปต์"}
        case "slv": 
            return Country{Alpha2: "sv", Alpha3: "slv", ID: 222, EN:"El Salvador", TH:"เอลซัลวาดอร์"}
        case "gnq": 
            return Country{Alpha2: "gq", Alpha3: "gnq", ID: 226, EN:"Equatorial Guinea", TH:"อิเควทอเรียลกินี"}
        case "eri": 
            return Country{Alpha2: "er", Alpha3: "eri", ID: 232, EN:"Eritrea", TH:"เอริเทรีย"}
        case "est": 
            return Country{Alpha2: "ee", Alpha3: "est", ID: 233, EN:"Estonia", TH:"เอสโตเนีย"}
        case "eth": 
            return Country{Alpha2: "et", Alpha3: "eth", ID: 231, EN:"Ethiopia", TH:"เอธิโอเปีย"}
        case "fji": 
            return Country{Alpha2: "fj", Alpha3: "fji", ID: 242, EN:"Fiji", TH:"ฟีจี"}
        case "fin": 
            return Country{Alpha2: "fi", Alpha3: "fin", ID: 246, EN:"Finland", TH:"ฟินแลนด์"}
        case "fra": 
            return Country{Alpha2: "fr", Alpha3: "fra", ID: 250, EN:"France", TH:"ฝรั่งเศส"}
        case "gab": 
            return Country{Alpha2: "ga", Alpha3: "gab", ID: 266, EN:"Gabon", TH:"กาบอง"}
        case "gmb": 
            return Country{Alpha2: "gm", Alpha3: "gmb", ID: 270, EN:"Gambia", TH:"แกมเบีย"}
        case "geo": 
            return Country{Alpha2: "ge", Alpha3: "geo", ID: 268, EN:"Georgia", TH:"จอร์เจีย"}
        case "deu": 
            return Country{Alpha2: "de", Alpha3: "deu", ID: 276, EN:"Germany", TH:"เยอรมนี"}
        case "gha": 
            return Country{Alpha2: "gh", Alpha3: "gha", ID: 288, EN:"Ghana", TH:"กานา"}
        case "grc": 
            return Country{Alpha2: "gr", Alpha3: "grc", ID: 300, EN:"Greece", TH:"กรีซ"}
        case "grd": 
            return Country{Alpha2: "gd", Alpha3: "grd", ID: 308, EN:"Grenada", TH:"เกรเนดา"}
        case "gtm": 
            return Country{Alpha2: "gt", Alpha3: "gtm", ID: 320, EN:"Guatemala", TH:"กัวเตมาลา"}
        case "gin": 
            return Country{Alpha2: "gn", Alpha3: "gin", ID: 324, EN:"Guinea", TH:"กินี"}
        case "gnb": 
            return Country{Alpha2: "gw", Alpha3: "gnb", ID: 624, EN:"Guinea-Bissau", TH:"กินี-บิสเซา"}
        case "guy": 
            return Country{Alpha2: "gy", Alpha3: "guy", ID: 328, EN:"Guyana", TH:"กายอานา"}
        case "hti": 
            return Country{Alpha2: "ht", Alpha3: "hti", ID: 332, EN:"Haiti", TH:"เฮติ"}
        case "hnd": 
            return Country{Alpha2: "hn", Alpha3: "hnd", ID: 340, EN:"Honduras", TH:"ฮอนดูรัส"}
        case "hun": 
            return Country{Alpha2: "hu", Alpha3: "hun", ID: 348, EN:"Hungary", TH:"ฮังการี"}
        case "isl": 
            return Country{Alpha2: "is", Alpha3: "isl", ID: 352, EN:"Iceland", TH:"ไอซ์แลนด์"}
        case "ind": 
            return Country{Alpha2: "in", Alpha3: "ind", ID: 356, EN:"India", TH:"อินเดีย"}
        case "idn": 
            return Country{Alpha2: "id", Alpha3: "idn", ID: 360, EN:"Indonesia", TH:"อินโดนีเซีย"}
        case "irn": 
            return Country{Alpha2: "ir", Alpha3: "irn", ID: 364, EN:"Iran (Islamic Republic of)", TH:"อิหร่าน"}
        case "irq": 
            return Country{Alpha2: "iq", Alpha3: "irq", ID: 368, EN:"Iraq", TH:"อิรัก"}
        case "irl": 
            return Country{Alpha2: "ie", Alpha3: "irl", ID: 372, EN:"Ireland", TH:"ไอร์แลนด์"}
        case "isr": 
            return Country{Alpha2: "il", Alpha3: "isr", ID: 376, EN:"Israel", TH:"อิสราเอล"}
        case "ita": 
            return Country{Alpha2: "it", Alpha3: "ita", ID: 380, EN:"Italy", TH:"อิตาลี"}
        case "jam": 
            return Country{Alpha2: "jm", Alpha3: "jam", ID: 388, EN:"Jamaica", TH:"จาเมกา"}
        case "jpn": 
            return Country{Alpha2: "jp", Alpha3: "jpn", ID: 392, EN:"Japan", TH:"ญี่ปุ่น"}
        case "jor": 
            return Country{Alpha2: "jo", Alpha3: "jor", ID: 400, EN:"Jordan", TH:"จอร์แดน"}
        case "kaz": 
            return Country{Alpha2: "kz", Alpha3: "kaz", ID: 398, EN:"Kazakhstan", TH:"คาซัคสถาน"}
        case "ken": 
            return Country{Alpha2: "ke", Alpha3: "ken", ID: 404, EN:"Kenya", TH:"เคนยา"}
        case "kir": 
            return Country{Alpha2: "ki", Alpha3: "kir", ID: 296, EN:"Kiribati", TH:"คิริบาส"}
        case "prk": 
            return Country{Alpha2: "kp", Alpha3: "prk", ID: 408, EN:"Korea (Democratic People's Republic of)", TH:"เกาหลีเหนือ"}
        case "kor": 
            return Country{Alpha2: "kr", Alpha3: "kor", ID: 410, EN:"Korea, Republic of", TH:"เกาหลีใต้"}
        case "kwt": 
            return Country{Alpha2: "kw", Alpha3: "kwt", ID: 414, EN:"Kuwait", TH:"คูเวต"}
        case "kgz": 
            return Country{Alpha2: "kg", Alpha3: "kgz", ID: 417, EN:"Kyrgyzstan", TH:"คีร์กีซสถาน"}
        case "lao": 
            return Country{Alpha2: "la", Alpha3: "lao", ID: 418, EN:"Lao People's Democratic Republic", TH:"ลาว"}
        case "lva": 
            return Country{Alpha2: "lv", Alpha3: "lva", ID: 428, EN:"Latvia", TH:"ลัตเวีย"}
        case "lbn": 
            return Country{Alpha2: "lb", Alpha3: "lbn", ID: 422, EN:"Lebanon", TH:"เลบานอน"}
        case "lso": 
            return Country{Alpha2: "ls", Alpha3: "lso", ID: 426, EN:"Lesotho", TH:"เลโซโท"}
        case "lbr": 
            return Country{Alpha2: "lr", Alpha3: "lbr", ID: 430, EN:"Liberia", TH:"ไลบีเรีย"}
        case "lby": 
            return Country{Alpha2: "ly", Alpha3: "lby", ID: 434, EN:"Libya", TH:"ลิเบีย"}
        case "lie": 
            return Country{Alpha2: "li", Alpha3: "lie", ID: 438, EN:"Liechtenstein", TH:"ลิกเตนสไตน์"}
        case "ltu": 
            return Country{Alpha2: "lt", Alpha3: "ltu", ID: 440, EN:"Lithuania", TH:"ลิทัวเนีย"}
        case "lux": 
            return Country{Alpha2: "lu", Alpha3: "lux", ID: 442, EN:"Luxembourg", TH:"ลักเซมเบิร์ก"}
        case "mkd": 
            return Country{Alpha2: "mk", Alpha3: "mkd", ID: 807, EN:"North Macedonia", TH:"นอร์ทมาซิโดเนีย"}
        case "mdg": 
            return Country{Alpha2: "mg", Alpha3: "mdg", ID: 450, EN:"Madagascar", TH:"มาดากัสการ์"}
        case "mwi": 
            return Country{Alpha2: "mw", Alpha3: "mwi", ID: 454, EN:"Malawi", TH:"มาลาวี"}
        case "mys": 
            return Country{Alpha2: "my", Alpha3: "mys", ID: 458, EN:"Malaysia", TH:"มาเลเซีย"}
        case "mdv": 
            return Country{Alpha2: "mv", Alpha3: "mdv", ID: 462, EN:"Maldives", TH:"มัลดีฟส์"}
        case "mli": 
            return Country{Alpha2: "ml", Alpha3: "mli", ID: 466, EN:"Mali", TH:"มาลี"}
        case "mlt": 
            return Country{Alpha2: "mt", Alpha3: "mlt", ID: 470, EN:"Malta", TH:"มอลตา"}
        case "mhl": 
            return Country{Alpha2: "mh", Alpha3: "mhl", ID: 584, EN:"Marshall Islands", TH:"หมู่เกาะมาร์แชลล์"}
        case "mrt": 
            return Country{Alpha2: "mr", Alpha3: "mrt", ID: 478, EN:"Mauritania", TH:"มอริเตเนีย"}
        case "mus": 
            return Country{Alpha2: "mu", Alpha3: "mus", ID: 480, EN:"Mauritius", TH:"มอริเชียส"}
        case "mex": 
            return Country{Alpha2: "mx", Alpha3: "mex", ID: 484, EN:"Mexico", TH:"เม็กซิโก"}
        case "fsm": 
            return Country{Alpha2: "fm", Alpha3: "fsm", ID: 583, EN:"Micronesia (Federated States of)", TH:"ไมโครนีเซีย"}
        case "mar": 
            return Country{Alpha2: "ma", Alpha3: "mar", ID: 504, EN:"Morocco", TH:"โมร็อกโก"}
        case "mda": 
            return Country{Alpha2: "md", Alpha3: "mda", ID: 498, EN:"Moldova, Republic of", TH:"มอลโดวา"}
        case "mco": 
            return Country{Alpha2: "mc", Alpha3: "mco", ID: 492, EN:"Monaco", TH:"โมนาโก"}
        case "mng": 
            return Country{Alpha2: "mn", Alpha3: "mng", ID: 496, EN:"Mongolia", TH:"มองโกเลีย"}
        case "mne": 
            return Country{Alpha2: "me", Alpha3: "mne", ID: 499, EN:"Montenegro", TH:"มอนเตเนโกร"}
        case "moz": 
            return Country{Alpha2: "mz", Alpha3: "moz", ID: 508, EN:"Mozambique", TH:"โมซัมบิก"}
        case "mmr": 
            return Country{Alpha2: "mm", Alpha3: "mmr", ID: 104, EN:"Myanmar", TH:"พม่า"}
        case "nam": 
            return Country{Alpha2: "na", Alpha3: "nam", ID: 516, EN:"Namibia", TH:"นามิเบีย"}
        case "nru": 
            return Country{Alpha2: "nr", Alpha3: "nru", ID: 520, EN:"Nauru", TH:"นาอูรู"}
        case "npl": 
            return Country{Alpha2: "np", Alpha3: "npl", ID: 524, EN:"Nepal", TH:"เนปาล"}
        case "nld": 
            return Country{Alpha2: "nl", Alpha3: "nld", ID: 528, EN:"Netherlands", TH:"เนเธอร์แลนด์"}
        case "nzl": 
            return Country{Alpha2: "nz", Alpha3: "nzl", ID: 554, EN:"New Zealand", TH:"นิวซีแลนด์"}
        case "nic": 
            return Country{Alpha2: "ni", Alpha3: "nic", ID: 558, EN:"Nicaragua", TH:"นิการากัว"}
        case "ner": 
            return Country{Alpha2: "ne", Alpha3: "ner", ID: 562, EN:"Niger", TH:"ไนเจอร์"}
        case "nga": 
            return Country{Alpha2: "ng", Alpha3: "nga", ID: 566, EN:"Nigeria", TH:"ไนจีเรีย"}
        case "nor": 
            return Country{Alpha2: "no", Alpha3: "nor", ID: 578, EN:"Norway", TH:"นอร์เวย์"}
        case "omn": 
            return Country{Alpha2: "om", Alpha3: "omn", ID: 512, EN:"Oman", TH:"โอมาน"}
        case "pak": 
            return Country{Alpha2: "pk", Alpha3: "pak", ID: 586, EN:"Pakistan", TH:"ปากีสถาน"}
        case "plw": 
            return Country{Alpha2: "pw", Alpha3: "plw", ID: 585, EN:"Palau", TH:"ปาเลา"}
        case "pan": 
            return Country{Alpha2: "pa", Alpha3: "pan", ID: 591, EN:"Panama", TH:"ปานามา"}
        case "png": 
            return Country{Alpha2: "pg", Alpha3: "png", ID: 598, EN:"Papua New Guinea", TH:"ปาปัวนิวกินี"}
        case "pry": 
            return Country{Alpha2: "py", Alpha3: "pry", ID: 600, EN:"Paraguay", TH:"ปารากวัย"}
        case "per": 
            return Country{Alpha2: "pe", Alpha3: "per", ID: 604, EN:"Peru", TH:"เปรู"}
        case "phl": 
            return Country{Alpha2: "ph", Alpha3: "phl", ID: 608, EN:"Philippines", TH:"ฟิลิปปินส์"}
        case "pol": 
            return Country{Alpha2: "pl", Alpha3: "pol", ID: 616, EN:"Poland", TH:"โปแลนด์"}
        case "prt": 
            return Country{Alpha2: "pt", Alpha3: "prt", ID: 620, EN:"Portugal", TH:"โปรตุเกส"}
        case "qat": 
            return Country{Alpha2: "qa", Alpha3: "qat", ID: 634, EN:"Qatar", TH:"กาตาร์"}
        case "rou": 
            return Country{Alpha2: "ro", Alpha3: "rou", ID: 642, EN:"Romania", TH:"โรมาเนีย"}
        case "rus": 
            return Country{Alpha2: "ru", Alpha3: "rus", ID: 643, EN:"Russian Federation", TH:"รัสเซีย"}
        case "rwa": 
            return Country{Alpha2: "rw", Alpha3: "rwa", ID: 646, EN:"Rwanda", TH:"รวันดา"}
        case "kna": 
            return Country{Alpha2: "kn", Alpha3: "kna", ID: 659, EN:"Saint Kitts and Nevis", TH:"เซนต์คิตส์และเนวิส"}
        case "lca": 
            return Country{Alpha2: "lc", Alpha3: "lca", ID: 662, EN:"Saint Lucia", TH:"เซนต์ลูเชีย"}
        case "vct": 
            return Country{Alpha2: "vc", Alpha3: "vct", ID: 670, EN:"Saint Vincent and the Grenadines", TH:"เซนต์วินเซนต์และเกรนาดีนส์"}
        case "wsm": 
            return Country{Alpha2: "ws", Alpha3: "wsm", ID: 882, EN:"Samoa", TH:"ซามัว"}
        case "smr": 
            return Country{Alpha2: "sm", Alpha3: "smr", ID: 674, EN:"San Marino", TH:"ซานมารีโน"}
        case "stp": 
            return Country{Alpha2: "st", Alpha3: "stp", ID: 678, EN:"Sao Tome and Principe", TH:"เซาตูเมและปรินซีปี"}
        case "sau": 
            return Country{Alpha2: "sa", Alpha3: "sau", ID: 682, EN:"Saudi Arabia", TH:"ซาอุดีอาระเบีย"}
        case "sen": 
            return Country{Alpha2: "sn", Alpha3: "sen", ID: 686, EN:"Senegal", TH:"เซเนกัล"}
        case "srb": 
            return Country{Alpha2: "rs", Alpha3: "srb", ID: 688, EN:"Serbia", TH:"เซอร์เบีย"}
        case "syc": 
            return Country{Alpha2: "sc", Alpha3: "syc", ID: 690, EN:"Seychelles", TH:"เซเชลส์"}
        case "sle": 
            return Country{Alpha2: "sl", Alpha3: "sle", ID: 694, EN:"Sierra Leone", TH:"เซียร์ราลีโอน"}
        case "sgp": 
            return Country{Alpha2: "sg", Alpha3: "sgp", ID: 702, EN:"Singapore", TH:"สิงคโปร์"}
        case "svk": 
            return Country{Alpha2: "sk", Alpha3: "svk", ID: 703, EN:"Slovakia", TH:"สโลวาเกีย"}
        case "svn": 
            return Country{Alpha2: "si", Alpha3: "svn", ID: 705, EN:"Slovenia", TH:"สโลวีเนีย"}
        case "slb": 
            return Country{Alpha2: "sb", Alpha3: "slb", ID: 90, EN:"Solomon Islands", TH:"หมู่เกาะโซโลมอน"}
        case "som": 
            return Country{Alpha2: "so", Alpha3: "som", ID: 706, EN:"Somalia", TH:"โซมาเลีย"}
        case "zaf": 
            return Country{Alpha2: "za", Alpha3: "zaf", ID: 710, EN:"South Africa", TH:"แอฟริกาใต้"}
        case "ssd": 
            return Country{Alpha2: "ss", Alpha3: "ssd", ID: 728, EN:"South Sudan", TH:"เซาท์ซูดาน"}
        case "esp": 
            return Country{Alpha2: "es", Alpha3: "esp", ID: 724, EN:"Spain", TH:"สเปน"}
        case "lka": 
            return Country{Alpha2: "lk", Alpha3: "lka", ID: 144, EN:"Sri Lanka", TH:"ศรีลังกา"}
        case "sdn": 
            return Country{Alpha2: "sd", Alpha3: "sdn", ID: 729, EN:"Sudan", TH:"ซูดาน"}
        case "sur": 
            return Country{Alpha2: "sr", Alpha3: "sur", ID: 740, EN:"Suriname", TH:"ซูรินาม"}
        case "swz": 
            return Country{Alpha2: "sz", Alpha3: "swz", ID: 748, EN:"Eswatini", TH:"เอสวาตีนี"}
        case "swe": 
            return Country{Alpha2: "se", Alpha3: "swe", ID: 752, EN:"Sweden", TH:"สวีเดน"}
        case "che": 
            return Country{Alpha2: "ch", Alpha3: "che", ID: 756, EN:"Switzerland", TH:"สวิตเซอร์แลนด์"}
        case "syr": 
            return Country{Alpha2: "sy", Alpha3: "syr", ID: 760, EN:"Syrian Arab Republic", TH:"ซีเรีย"}
        case "tjk": 
            return Country{Alpha2: "tj", Alpha3: "tjk", ID: 762, EN:"Tajikistan", TH:"ทาจิกิสถาน"}
        case "tza": 
            return Country{Alpha2: "tz", Alpha3: "tza", ID: 834, EN:"Tanzania, United Republic of", TH:"แทนซาเนีย"}
        case "tha": 
            return Country{Alpha2: "th", Alpha3: "tha", ID: 764, EN:"Thailand", TH:"ไทย"}
        case "tls": 
            return Country{Alpha2: "tl", Alpha3: "tls", ID: 626, EN:"Timor-Leste", TH:"ติมอร์-เลสเต"}
        case "tgo": 
            return Country{Alpha2: "tg", Alpha3: "tgo", ID: 768, EN:"Togo", TH:"โตโก"}
        case "ton": 
            return Country{Alpha2: "to", Alpha3: "ton", ID: 776, EN:"Tonga", TH:"ตองงา"}
        case "tto": 
            return Country{Alpha2: "tt", Alpha3: "tto", ID: 780, EN:"Trinidad and Tobago", TH:"ตรินิแดดและโตเบโก"}
        case "tun": 
            return Country{Alpha2: "tn", Alpha3: "tun", ID: 788, EN:"Tunisia", TH:"ตูนิเซีย"}
        case "tur": 
            return Country{Alpha2: "tr", Alpha3: "tur", ID: 792, EN:"Turkey", TH:"ตุรกี"}
        case "tkm": 
            return Country{Alpha2: "tm", Alpha3: "tkm", ID: 795, EN:"Turkmenistan", TH:"เติร์กเมนิสถาน"}
        case "tuv": 
            return Country{Alpha2: "tv", Alpha3: "tuv", ID: 798, EN:"Tuvalu", TH:"ตูวาลู"}
        case "uga": 
            return Country{Alpha2: "ug", Alpha3: "uga", ID: 800, EN:"Uganda", TH:"ยูกันดา"}
        case "ukr": 
            return Country{Alpha2: "ua", Alpha3: "ukr", ID: 804, EN:"Ukraine", TH:"ยูเครน"}
        case "are": 
            return Country{Alpha2: "ae", Alpha3: "are", ID: 784, EN:"United Arab Emirates", TH:"สหรัฐอาหรับเอมิเรตส์"}
        case "gbr": 
            return Country{Alpha2: "gb", Alpha3: "gbr", ID: 826, EN:"United Kingdom of Great Britain and Northern Ireland", TH:"สหราชอาณาจักร"}
        case "usa": 
            return Country{Alpha2: "us", Alpha3: "usa", ID: 840, EN:"United States of America", TH:"สหรัฐอเมริกา"}
        case "ury": 
            return Country{Alpha2: "uy", Alpha3: "ury", ID: 858, EN:"Uruguay", TH:"อุรุกวัย"}
        case "uzb": 
            return Country{Alpha2: "uz", Alpha3: "uzb", ID: 860, EN:"Uzbekistan", TH:"อุซเบกิสถาน"}
        case "vut": 
            return Country{Alpha2: "vu", Alpha3: "vut", ID: 548, EN:"Vanuatu", TH:"วานูอาตู"}
        case "ven": 
            return Country{Alpha2: "ve", Alpha3: "ven", ID: 862, EN:"Venezuela (Bolivarian Republic of)", TH:"เวเนซุเอลา"}
        case "vnm": 
            return Country{Alpha2: "vn", Alpha3: "vnm", ID: 704, EN:"Viet Nam", TH:"เวียดนาม"}
        case "yem": 
            return Country{Alpha2: "ye", Alpha3: "yem", ID: 887, EN:"Yemen", TH:"เยเมน"}
        case "zmb": 
            return Country{Alpha2: "zm", Alpha3: "zmb", ID: 894, EN:"Zambia", TH:"แซมเบีย"}
        case "zwe": 
            return Country{Alpha2: "zw", Alpha3: "zwe", ID: 716, EN:"Zimbabwe", TH:"ซิมบับเว"}
        default:
            return Country{}
    }
}

var alpha3CountryMap map[string]Country = map[string]Country{
    "afg": {Alpha2: "af", Alpha3: "afg", ID: 4, EN:"Afghanistan", TH:"อัฟกานิสถาน"},
    "alb": {Alpha2: "al", Alpha3: "alb", ID: 8, EN:"Albania", TH:"แอลเบเนีย"},
    "dza": {Alpha2: "dz", Alpha3: "dza", ID: 12, EN:"Algeria", TH:"แอลจีเรีย"},
    "and": {Alpha2: "ad", Alpha3: "and", ID: 20, EN:"Andorra", TH:"อันดอร์รา"},
    "ago": {Alpha2: "ao", Alpha3: "ago", ID: 24, EN:"Angola", TH:"แองโกลา"},
    "atg": {Alpha2: "ag", Alpha3: "atg", ID: 28, EN:"Antigua and Barbuda", TH:"แอนติกาและบาร์บูดา"},
    "arg": {Alpha2: "ar", Alpha3: "arg", ID: 32, EN:"Argentina", TH:"อาร์เจนตินา"},
    "arm": {Alpha2: "am", Alpha3: "arm", ID: 51, EN:"Armenia", TH:"อาร์มีเนีย"},
    "aus": {Alpha2: "au", Alpha3: "aus", ID: 36, EN:"Australia", TH:"ออสเตรเลีย"},
    "aut": {Alpha2: "at", Alpha3: "aut", ID: 40, EN:"Austria", TH:"ออสเตรีย"},
    "aze": {Alpha2: "az", Alpha3: "aze", ID: 31, EN:"Azerbaijan", TH:"อาเซอร์ไบจาน"},
    "bhs": {Alpha2: "bs", Alpha3: "bhs", ID: 44, EN:"Bahamas", TH:"บาฮามาส"},
    "bhr": {Alpha2: "bh", Alpha3: "bhr", ID: 48, EN:"Bahrain", TH:"บาห์เรน"},
    "bgd": {Alpha2: "bd", Alpha3: "bgd", ID: 50, EN:"Bangladesh", TH:"บังกลาเทศ"},
    "brb": {Alpha2: "bb", Alpha3: "brb", ID: 52, EN:"Barbados", TH:"บาร์เบโดส"},
    "blr": {Alpha2: "by", Alpha3: "blr", ID: 112, EN:"Belarus", TH:"เบลารุส"},
    "bel": {Alpha2: "be", Alpha3: "bel", ID: 56, EN:"Belgium", TH:"เบลเยียม"},
    "blz": {Alpha2: "bz", Alpha3: "blz", ID: 84, EN:"Belize", TH:"เบลีซ"},
    "ben": {Alpha2: "bj", Alpha3: "ben", ID: 204, EN:"Benin", TH:"เบนิน"},
    "btn": {Alpha2: "bt", Alpha3: "btn", ID: 64, EN:"Bhutan", TH:"ภูฏาน"},
    "bol": {Alpha2: "bo", Alpha3: "bol", ID: 68, EN:"Bolivia (Plurinational State of)", TH:"โบลิเวีย"},
    "bih": {Alpha2: "ba", Alpha3: "bih", ID: 70, EN:"Bosnia and Herzegovina", TH:"บอสเนียและเฮอร์เซโกวีนา"},
    "bwa": {Alpha2: "bw", Alpha3: "bwa", ID: 72, EN:"Botswana", TH:"บอตสวานา"},
    "bra": {Alpha2: "br", Alpha3: "bra", ID: 76, EN:"Brazil", TH:"บราซิล"},
    "brn": {Alpha2: "bn", Alpha3: "brn", ID: 96, EN:"Brunei Darussalam", TH:"บรูไน"},
    "bgr": {Alpha2: "bg", Alpha3: "bgr", ID: 100, EN:"Bulgaria", TH:"บัลแกเรีย"},
    "bfa": {Alpha2: "bf", Alpha3: "bfa", ID: 854, EN:"Burkina Faso", TH:"บูร์กินาฟาโซ"},
    "bdi": {Alpha2: "bi", Alpha3: "bdi", ID: 108, EN:"Burundi", TH:"บุรุนดี"},
    "khm": {Alpha2: "kh", Alpha3: "khm", ID: 116, EN:"Cambodia", TH:"กัมพูชา"},
    "cmr": {Alpha2: "cm", Alpha3: "cmr", ID: 120, EN:"Cameroon", TH:"แคเมอรูน"},
    "can": {Alpha2: "ca", Alpha3: "can", ID: 124, EN:"Canada", TH:"แคนาดา"},
    "cpv": {Alpha2: "cv", Alpha3: "cpv", ID: 132, EN:"Cabo Verde", TH:"กาบูเวร์ดี"},
    "caf": {Alpha2: "cf", Alpha3: "caf", ID: 140, EN:"Central African Republic", TH:"สาธารณรัฐแอฟริกากลาง"},
    "tcd": {Alpha2: "td", Alpha3: "tcd", ID: 148, EN:"Chad", TH:"ชาด"},
    "chl": {Alpha2: "cl", Alpha3: "chl", ID: 152, EN:"Chile", TH:"ชิลี"},
    "chn": {Alpha2: "cn", Alpha3: "chn", ID: 156, EN:"China", TH:"จีน"},
    "col": {Alpha2: "co", Alpha3: "col", ID: 170, EN:"Colombia", TH:"โคลอมเบีย"},
    "com": {Alpha2: "km", Alpha3: "com", ID: 174, EN:"Comoros", TH:"คอโมโรส"},
    "cog": {Alpha2: "cg", Alpha3: "cog", ID: 178, EN:"Congo", TH:"สาธารณรัฐคองโก"},
    "cod": {Alpha2: "cd", Alpha3: "cod", ID: 180, EN:"Congo, Democratic Republic of the", TH:"สาธารณรัฐประชาธิปไตยคองโก"},
    "cri": {Alpha2: "cr", Alpha3: "cri", ID: 188, EN:"Costa Rica", TH:"คอสตาริกา"},
    "civ": {Alpha2: "ci", Alpha3: "civ", ID: 384, EN:"Côte d'Ivoire", TH:"โกตดิวัวร์"},
    "hrv": {Alpha2: "hr", Alpha3: "hrv", ID: 191, EN:"Croatia", TH:"โครเอเชีย"},
    "cub": {Alpha2: "cu", Alpha3: "cub", ID: 192, EN:"Cuba", TH:"คิวบา"},
    "cyp": {Alpha2: "cy", Alpha3: "cyp", ID: 196, EN:"Cyprus", TH:"ไซปรัส"},
    "cze": {Alpha2: "cz", Alpha3: "cze", ID: 203, EN:"Czechia", TH:"เช็กเกีย"},
    "dnk": {Alpha2: "dk", Alpha3: "dnk", ID: 208, EN:"Denmark", TH:"เดนมาร์ก"},
    "dji": {Alpha2: "dj", Alpha3: "dji", ID: 262, EN:"Djibouti", TH:"จิบูตี"},
    "dma": {Alpha2: "dm", Alpha3: "dma", ID: 212, EN:"Dominica", TH:"ดอมินีกา"},
    "dom": {Alpha2: "do", Alpha3: "dom", ID: 214, EN:"Dominican Republic", TH:"สาธารณรัฐโดมินิกัน"},
    "ecu": {Alpha2: "ec", Alpha3: "ecu", ID: 218, EN:"Ecuador", TH:"เอกวาดอร์"},
    "egy": {Alpha2: "eg", Alpha3: "egy", ID: 818, EN:"Egypt", TH:"อียิปต์"},
    "slv": {Alpha2: "sv", Alpha3: "slv", ID: 222, EN:"El Salvador", TH:"เอลซัลวาดอร์"},
    "gnq": {Alpha2: "gq", Alpha3: "gnq", ID: 226, EN:"Equatorial Guinea", TH:"อิเควทอเรียลกินี"},
    "eri": {Alpha2: "er", Alpha3: "eri", ID: 232, EN:"Eritrea", TH:"เอริเทรีย"},
    "est": {Alpha2: "ee", Alpha3: "est", ID: 233, EN:"Estonia", TH:"เอสโตเนีย"},
    "eth": {Alpha2: "et", Alpha3: "eth", ID: 231, EN:"Ethiopia", TH:"เอธิโอเปีย"},
    "fji": {Alpha2: "fj", Alpha3: "fji", ID: 242, EN:"Fiji", TH:"ฟีจี"},
    "fin": {Alpha2: "fi", Alpha3: "fin", ID: 246, EN:"Finland", TH:"ฟินแลนด์"},
    "fra": {Alpha2: "fr", Alpha3: "fra", ID: 250, EN:"France", TH:"ฝรั่งเศส"},
    "gab": {Alpha2: "ga", Alpha3: "gab", ID: 266, EN:"Gabon", TH:"กาบอง"},
    "gmb": {Alpha2: "gm", Alpha3: "gmb", ID: 270, EN:"Gambia", TH:"แกมเบีย"},
    "geo": {Alpha2: "ge", Alpha3: "geo", ID: 268, EN:"Georgia", TH:"จอร์เจีย"},
    "deu": {Alpha2: "de", Alpha3: "deu", ID: 276, EN:"Germany", TH:"เยอรมนี"},
    "gha": {Alpha2: "gh", Alpha3: "gha", ID: 288, EN:"Ghana", TH:"กานา"},
    "grc": {Alpha2: "gr", Alpha3: "grc", ID: 300, EN:"Greece", TH:"กรีซ"},
    "grd": {Alpha2: "gd", Alpha3: "grd", ID: 308, EN:"Grenada", TH:"เกรเนดา"},
    "gtm": {Alpha2: "gt", Alpha3: "gtm", ID: 320, EN:"Guatemala", TH:"กัวเตมาลา"},
    "gin": {Alpha2: "gn", Alpha3: "gin", ID: 324, EN:"Guinea", TH:"กินี"},
    "gnb": {Alpha2: "gw", Alpha3: "gnb", ID: 624, EN:"Guinea-Bissau", TH:"กินี-บิสเซา"},
    "guy": {Alpha2: "gy", Alpha3: "guy", ID: 328, EN:"Guyana", TH:"กายอานา"},
    "hti": {Alpha2: "ht", Alpha3: "hti", ID: 332, EN:"Haiti", TH:"เฮติ"},
    "hnd": {Alpha2: "hn", Alpha3: "hnd", ID: 340, EN:"Honduras", TH:"ฮอนดูรัส"},
    "hun": {Alpha2: "hu", Alpha3: "hun", ID: 348, EN:"Hungary", TH:"ฮังการี"},
    "isl": {Alpha2: "is", Alpha3: "isl", ID: 352, EN:"Iceland", TH:"ไอซ์แลนด์"},
    "ind": {Alpha2: "in", Alpha3: "ind", ID: 356, EN:"India", TH:"อินเดีย"},
    "idn": {Alpha2: "id", Alpha3: "idn", ID: 360, EN:"Indonesia", TH:"อินโดนีเซีย"},
    "irn": {Alpha2: "ir", Alpha3: "irn", ID: 364, EN:"Iran (Islamic Republic of)", TH:"อิหร่าน"},
    "irq": {Alpha2: "iq", Alpha3: "irq", ID: 368, EN:"Iraq", TH:"อิรัก"},
    "irl": {Alpha2: "ie", Alpha3: "irl", ID: 372, EN:"Ireland", TH:"ไอร์แลนด์"},
    "isr": {Alpha2: "il", Alpha3: "isr", ID: 376, EN:"Israel", TH:"อิสราเอล"},
    "ita": {Alpha2: "it", Alpha3: "ita", ID: 380, EN:"Italy", TH:"อิตาลี"},
    "jam": {Alpha2: "jm", Alpha3: "jam", ID: 388, EN:"Jamaica", TH:"จาเมกา"},
    "jpn": {Alpha2: "jp", Alpha3: "jpn", ID: 392, EN:"Japan", TH:"ญี่ปุ่น"},
    "jor": {Alpha2: "jo", Alpha3: "jor", ID: 400, EN:"Jordan", TH:"จอร์แดน"},
    "kaz": {Alpha2: "kz", Alpha3: "kaz", ID: 398, EN:"Kazakhstan", TH:"คาซัคสถาน"},
    "ken": {Alpha2: "ke", Alpha3: "ken", ID: 404, EN:"Kenya", TH:"เคนยา"},
    "kir": {Alpha2: "ki", Alpha3: "kir", ID: 296, EN:"Kiribati", TH:"คิริบาส"},
    "prk": {Alpha2: "kp", Alpha3: "prk", ID: 408, EN:"Korea (Democratic People's Republic of)", TH:"เกาหลีเหนือ"},
    "kor": {Alpha2: "kr", Alpha3: "kor", ID: 410, EN:"Korea, Republic of", TH:"เกาหลีใต้"},
    "kwt": {Alpha2: "kw", Alpha3: "kwt", ID: 414, EN:"Kuwait", TH:"คูเวต"},
    "kgz": {Alpha2: "kg", Alpha3: "kgz", ID: 417, EN:"Kyrgyzstan", TH:"คีร์กีซสถาน"},
    "lao": {Alpha2: "la", Alpha3: "lao", ID: 418, EN:"Lao People's Democratic Republic", TH:"ลาว"},
    "lva": {Alpha2: "lv", Alpha3: "lva", ID: 428, EN:"Latvia", TH:"ลัตเวีย"},
    "lbn": {Alpha2: "lb", Alpha3: "lbn", ID: 422, EN:"Lebanon", TH:"เลบานอน"},
    "lso": {Alpha2: "ls", Alpha3: "lso", ID: 426, EN:"Lesotho", TH:"เลโซโท"},
    "lbr": {Alpha2: "lr", Alpha3: "lbr", ID: 430, EN:"Liberia", TH:"ไลบีเรีย"},
    "lby": {Alpha2: "ly", Alpha3: "lby", ID: 434, EN:"Libya", TH:"ลิเบีย"},
    "lie": {Alpha2: "li", Alpha3: "lie", ID: 438, EN:"Liechtenstein", TH:"ลิกเตนสไตน์"},
    "ltu": {Alpha2: "lt", Alpha3: "ltu", ID: 440, EN:"Lithuania", TH:"ลิทัวเนีย"},
    "lux": {Alpha2: "lu", Alpha3: "lux", ID: 442, EN:"Luxembourg", TH:"ลักเซมเบิร์ก"},
    "mkd": {Alpha2: "mk", Alpha3: "mkd", ID: 807, EN:"North Macedonia", TH:"นอร์ทมาซิโดเนีย"},
    "mdg": {Alpha2: "mg", Alpha3: "mdg", ID: 450, EN:"Madagascar", TH:"มาดากัสการ์"},
    "mwi": {Alpha2: "mw", Alpha3: "mwi", ID: 454, EN:"Malawi", TH:"มาลาวี"},
    "mys": {Alpha2: "my", Alpha3: "mys", ID: 458, EN:"Malaysia", TH:"มาเลเซีย"},
    "mdv": {Alpha2: "mv", Alpha3: "mdv", ID: 462, EN:"Maldives", TH:"มัลดีฟส์"},
    "mli": {Alpha2: "ml", Alpha3: "mli", ID: 466, EN:"Mali", TH:"มาลี"},
    "mlt": {Alpha2: "mt", Alpha3: "mlt", ID: 470, EN:"Malta", TH:"มอลตา"},
    "mhl": {Alpha2: "mh", Alpha3: "mhl", ID: 584, EN:"Marshall Islands", TH:"หมู่เกาะมาร์แชลล์"},
    "mrt": {Alpha2: "mr", Alpha3: "mrt", ID: 478, EN:"Mauritania", TH:"มอริเตเนีย"},
    "mus": {Alpha2: "mu", Alpha3: "mus", ID: 480, EN:"Mauritius", TH:"มอริเชียส"},
    "mex": {Alpha2: "mx", Alpha3: "mex", ID: 484, EN:"Mexico", TH:"เม็กซิโก"},
    "fsm": {Alpha2: "fm", Alpha3: "fsm", ID: 583, EN:"Micronesia (Federated States of)", TH:"ไมโครนีเซีย"},
    "mar": {Alpha2: "ma", Alpha3: "mar", ID: 504, EN:"Morocco", TH:"โมร็อกโก"},
    "mda": {Alpha2: "md", Alpha3: "mda", ID: 498, EN:"Moldova, Republic of", TH:"มอลโดวา"},
    "mco": {Alpha2: "mc", Alpha3: "mco", ID: 492, EN:"Monaco", TH:"โมนาโก"},
    "mng": {Alpha2: "mn", Alpha3: "mng", ID: 496, EN:"Mongolia", TH:"มองโกเลีย"},
    "mne": {Alpha2: "me", Alpha3: "mne", ID: 499, EN:"Montenegro", TH:"มอนเตเนโกร"},
    "moz": {Alpha2: "mz", Alpha3: "moz", ID: 508, EN:"Mozambique", TH:"โมซัมบิก"},
    "mmr": {Alpha2: "mm", Alpha3: "mmr", ID: 104, EN:"Myanmar", TH:"พม่า"},
    "nam": {Alpha2: "na", Alpha3: "nam", ID: 516, EN:"Namibia", TH:"นามิเบีย"},
    "nru": {Alpha2: "nr", Alpha3: "nru", ID: 520, EN:"Nauru", TH:"นาอูรู"},
    "npl": {Alpha2: "np", Alpha3: "npl", ID: 524, EN:"Nepal", TH:"เนปาล"},
    "nld": {Alpha2: "nl", Alpha3: "nld", ID: 528, EN:"Netherlands", TH:"เนเธอร์แลนด์"},
    "nzl": {Alpha2: "nz", Alpha3: "nzl", ID: 554, EN:"New Zealand", TH:"นิวซีแลนด์"},
    "nic": {Alpha2: "ni", Alpha3: "nic", ID: 558, EN:"Nicaragua", TH:"นิการากัว"},
    "ner": {Alpha2: "ne", Alpha3: "ner", ID: 562, EN:"Niger", TH:"ไนเจอร์"},
    "nga": {Alpha2: "ng", Alpha3: "nga", ID: 566, EN:"Nigeria", TH:"ไนจีเรีย"},
    "nor": {Alpha2: "no", Alpha3: "nor", ID: 578, EN:"Norway", TH:"นอร์เวย์"},
    "omn": {Alpha2: "om", Alpha3: "omn", ID: 512, EN:"Oman", TH:"โอมาน"},
    "pak": {Alpha2: "pk", Alpha3: "pak", ID: 586, EN:"Pakistan", TH:"ปากีสถาน"},
    "plw": {Alpha2: "pw", Alpha3: "plw", ID: 585, EN:"Palau", TH:"ปาเลา"},
    "pan": {Alpha2: "pa", Alpha3: "pan", ID: 591, EN:"Panama", TH:"ปานามา"},
    "png": {Alpha2: "pg", Alpha3: "png", ID: 598, EN:"Papua New Guinea", TH:"ปาปัวนิวกินี"},
    "pry": {Alpha2: "py", Alpha3: "pry", ID: 600, EN:"Paraguay", TH:"ปารากวัย"},
    "per": {Alpha2: "pe", Alpha3: "per", ID: 604, EN:"Peru", TH:"เปรู"},
    "phl": {Alpha2: "ph", Alpha3: "phl", ID: 608, EN:"Philippines", TH:"ฟิลิปปินส์"},
    "pol": {Alpha2: "pl", Alpha3: "pol", ID: 616, EN:"Poland", TH:"โปแลนด์"},
    "prt": {Alpha2: "pt", Alpha3: "prt", ID: 620, EN:"Portugal", TH:"โปรตุเกส"},
    "qat": {Alpha2: "qa", Alpha3: "qat", ID: 634, EN:"Qatar", TH:"กาตาร์"},
    "rou": {Alpha2: "ro", Alpha3: "rou", ID: 642, EN:"Romania", TH:"โรมาเนีย"},
    "rus": {Alpha2: "ru", Alpha3: "rus", ID: 643, EN:"Russian Federation", TH:"รัสเซีย"},
    "rwa": {Alpha2: "rw", Alpha3: "rwa", ID: 646, EN:"Rwanda", TH:"รวันดา"},
    "kna": {Alpha2: "kn", Alpha3: "kna", ID: 659, EN:"Saint Kitts and Nevis", TH:"เซนต์คิตส์และเนวิส"},
    "lca": {Alpha2: "lc", Alpha3: "lca", ID: 662, EN:"Saint Lucia", TH:"เซนต์ลูเชีย"},
    "vct": {Alpha2: "vc", Alpha3: "vct", ID: 670, EN:"Saint Vincent and the Grenadines", TH:"เซนต์วินเซนต์และเกรนาดีนส์"},
    "wsm": {Alpha2: "ws", Alpha3: "wsm", ID: 882, EN:"Samoa", TH:"ซามัว"},
    "smr": {Alpha2: "sm", Alpha3: "smr", ID: 674, EN:"San Marino", TH:"ซานมารีโน"},
    "stp": {Alpha2: "st", Alpha3: "stp", ID: 678, EN:"Sao Tome and Principe", TH:"เซาตูเมและปรินซีปี"},
    "sau": {Alpha2: "sa", Alpha3: "sau", ID: 682, EN:"Saudi Arabia", TH:"ซาอุดีอาระเบีย"},
    "sen": {Alpha2: "sn", Alpha3: "sen", ID: 686, EN:"Senegal", TH:"เซเนกัล"},
    "srb": {Alpha2: "rs", Alpha3: "srb", ID: 688, EN:"Serbia", TH:"เซอร์เบีย"},
    "syc": {Alpha2: "sc", Alpha3: "syc", ID: 690, EN:"Seychelles", TH:"เซเชลส์"},
    "sle": {Alpha2: "sl", Alpha3: "sle", ID: 694, EN:"Sierra Leone", TH:"เซียร์ราลีโอน"},
    "sgp": {Alpha2: "sg", Alpha3: "sgp", ID: 702, EN:"Singapore", TH:"สิงคโปร์"},
    "svk": {Alpha2: "sk", Alpha3: "svk", ID: 703, EN:"Slovakia", TH:"สโลวาเกีย"},
    "svn": {Alpha2: "si", Alpha3: "svn", ID: 705, EN:"Slovenia", TH:"สโลวีเนีย"},
    "slb": {Alpha2: "sb", Alpha3: "slb", ID: 90, EN:"Solomon Islands", TH:"หมู่เกาะโซโลมอน"},
    "som": {Alpha2: "so", Alpha3: "som", ID: 706, EN:"Somalia", TH:"โซมาเลีย"},
    "zaf": {Alpha2: "za", Alpha3: "zaf", ID: 710, EN:"South Africa", TH:"แอฟริกาใต้"},
    "ssd": {Alpha2: "ss", Alpha3: "ssd", ID: 728, EN:"South Sudan", TH:"เซาท์ซูดาน"},
    "esp": {Alpha2: "es", Alpha3: "esp", ID: 724, EN:"Spain", TH:"สเปน"},
    "lka": {Alpha2: "lk", Alpha3: "lka", ID: 144, EN:"Sri Lanka", TH:"ศรีลังกา"},
    "sdn": {Alpha2: "sd", Alpha3: "sdn", ID: 729, EN:"Sudan", TH:"ซูดาน"},
    "sur": {Alpha2: "sr", Alpha3: "sur", ID: 740, EN:"Suriname", TH:"ซูรินาม"},
    "swz": {Alpha2: "sz", Alpha3: "swz", ID: 748, EN:"Eswatini", TH:"เอสวาตีนี"},
    "swe": {Alpha2: "se", Alpha3: "swe", ID: 752, EN:"Sweden", TH:"สวีเดน"},
    "che": {Alpha2: "ch", Alpha3: "che", ID: 756, EN:"Switzerland", TH:"สวิตเซอร์แลนด์"},
    "syr": {Alpha2: "sy", Alpha3: "syr", ID: 760, EN:"Syrian Arab Republic", TH:"ซีเรีย"},
    "tjk": {Alpha2: "tj", Alpha3: "tjk", ID: 762, EN:"Tajikistan", TH:"ทาจิกิสถาน"},
    "tza": {Alpha2: "tz", Alpha3: "tza", ID: 834, EN:"Tanzania, United Republic of", TH:"แทนซาเนีย"},
    "tha": {Alpha2: "th", Alpha3: "tha", ID: 764, EN:"Thailand", TH:"ไทย"},
    "tls": {Alpha2: "tl", Alpha3: "tls", ID: 626, EN:"Timor-Leste", TH:"ติมอร์-เลสเต"},
    "tgo": {Alpha2: "tg", Alpha3: "tgo", ID: 768, EN:"Togo", TH:"โตโก"},
    "ton": {Alpha2: "to", Alpha3: "ton", ID: 776, EN:"Tonga", TH:"ตองงา"},
    "tto": {Alpha2: "tt", Alpha3: "tto", ID: 780, EN:"Trinidad and Tobago", TH:"ตรินิแดดและโตเบโก"},
    "tun": {Alpha2: "tn", Alpha3: "tun", ID: 788, EN:"Tunisia", TH:"ตูนิเซีย"},
    "tur": {Alpha2: "tr", Alpha3: "tur", ID: 792, EN:"Turkey", TH:"ตุรกี"},
    "tkm": {Alpha2: "tm", Alpha3: "tkm", ID: 795, EN:"Turkmenistan", TH:"เติร์กเมนิสถาน"},
    "tuv": {Alpha2: "tv", Alpha3: "tuv", ID: 798, EN:"Tuvalu", TH:"ตูวาลู"},
    "uga": {Alpha2: "ug", Alpha3: "uga", ID: 800, EN:"Uganda", TH:"ยูกันดา"},
    "ukr": {Alpha2: "ua", Alpha3: "ukr", ID: 804, EN:"Ukraine", TH:"ยูเครน"},
    "are": {Alpha2: "ae", Alpha3: "are", ID: 784, EN:"United Arab Emirates", TH:"สหรัฐอาหรับเอมิเรตส์"},
    "gbr": {Alpha2: "gb", Alpha3: "gbr", ID: 826, EN:"United Kingdom of Great Britain and Northern Ireland", TH:"สหราชอาณาจักร"},
    "usa": {Alpha2: "us", Alpha3: "usa", ID: 840, EN:"United States of America", TH:"สหรัฐอเมริกา"},
    "ury": {Alpha2: "uy", Alpha3: "ury", ID: 858, EN:"Uruguay", TH:"อุรุกวัย"},
    "uzb": {Alpha2: "uz", Alpha3: "uzb", ID: 860, EN:"Uzbekistan", TH:"อุซเบกิสถาน"},
    "vut": {Alpha2: "vu", Alpha3: "vut", ID: 548, EN:"Vanuatu", TH:"วานูอาตู"},
    "ven": {Alpha2: "ve", Alpha3: "ven", ID: 862, EN:"Venezuela (Bolivarian Republic of)", TH:"เวเนซุเอลา"},
    "vnm": {Alpha2: "vn", Alpha3: "vnm", ID: 704, EN:"Viet Nam", TH:"เวียดนาม"},
    "yem": {Alpha2: "ye", Alpha3: "yem", ID: 887, EN:"Yemen", TH:"เยเมน"},
    "zmb": {Alpha2: "zm", Alpha3: "zmb", ID: 894, EN:"Zambia", TH:"แซมเบีย"},
    "zwe": {Alpha2: "zw", Alpha3: "zwe", ID: 716, EN:"Zimbabwe", TH:"ซิมบับเว"},
}

func GetCountryFromAlpha3Map(alpha3 string) (Country, error) {
    return alpha3CountryMap[alpha3], nil
}