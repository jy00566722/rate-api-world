package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func InitMapData() {
	fmt.Println("InitMapData 开始")
	s := `AFGHANISTAN	Afghani	AFN
ÅLAND ISLANDS	Euro	EUR
ALBANIA	Lek	ALL
ALGERIA	Algerian Dinar	DZD
AMERICAN SAMOA	US Dollar	USD
ANDORRA	Euro	EUR
ANGOLA	Kwanza	AOA
ANGUILLA	East Caribbean Dollar	XCD
ANTIGUA AND BARBUDA	East Caribbean Dollar	XCD
ARGENTINA	Argentine Peso	ARS
ARMENIA	Armenian Dram	AMD
ARUBA	Aruban Florin	AWG
AUSTRALIA	Australian Dollar	AUD
AUSTRIA	Euro	EUR
AZERBAIJAN	Azerbaijan Manat	AZN
BAHAMAS (THE)	Bahamian Dollar	BSD
BAHRAIN	Bahraini Dinar	BHD
BANGLADESH	Taka	BDT
BARBADOS	Barbados Dollar	BBD
BELARUS	Belarusian Ruble	BYN
BELGIUM	Euro	EUR
BELIZE	Belize Dollar	BZD
BENIN	CFA Franc BCEAO	XOF
BERMUDA	Bermudian Dollar	BMD
BHUTAN	Indian Rupee	INR
BHUTAN	Ngultrum	BTN
BOLIVIA (PLURINATIONAL STATE OF)	Boliviano	BOB
BOLIVIA (PLURINATIONAL STATE OF)	Mvdol	BOV
BONAIRE, SINT EUSTATIUS AND SABA	US Dollar	USD
BOSNIA AND HERZEGOVINA	Convertible Mark	BAM
BOTSWANA	Pula	BWP
BOUVET ISLAND	Norwegian Krone	NOK
BRAZIL	Brazilian Real	BRL
BRITISH INDIAN OCEAN TERRITORY (THE)	US Dollar	USD
BRUNEI DARUSSALAM	Brunei Dollar	BND
BULGARIA	Bulgarian Lev	BGN
BURKINA FASO	CFA Franc BCEAO	XOF
BURUNDI	Burundi Franc	BIF
CABO VERDE	Cabo Verde Escudo	CVE
CAMBODIA	Riel	KHR
CAMEROON	CFA Franc BEAC	XAF
CANADA	Canadian Dollar	CAD
CAYMAN ISLANDS (THE)	Cayman Islands Dollar	KYD
CENTRAL AFRICAN REPUBLIC (THE)	CFA Franc BEAC	XAF
CHAD	CFA Franc BEAC	XAF
CHILE	Chilean Peso	CLP
CHILE	Unidad de Fomento	CLF
CHINA	Yuan Renminbi	CNY
CHINA	Offshore Chinese Yuan	CNH
CHRISTMAS ISLAND	Australian Dollar	AUD
COCOS (KEELING) ISLANDS (THE)	Australian Dollar	AUD
COLOMBIA	Colombian Peso	COP
COLOMBIA	Unidad de Valor Real	COU
COMOROS (THE)	Comorian Franc 	KMF
CONGO (THE DEMOCRATIC REPUBLIC OF THE)	Congolese Franc	CDF
CONGO (THE)	CFA Franc BEAC	XAF
COOK ISLANDS (THE)	New Zealand Dollar	NZD
COSTA RICA	Costa Rican Colon	CRC
CÔTE D'IVOIRE	CFA Franc BCEAO	XOF
CROATIA	Euro	EUR
CUBA	Cuban Peso	CUP
CUBA	Peso Convertible	CUC
CURAÇAO	Netherlands Antillean Guilder	ANG
CYPRUS	Euro	EUR
CZECHIA	Czech Koruna	CZK
DENMARK	Danish Krone	DKK
DJIBOUTI	Djibouti Franc	DJF
DOMINICA	East Caribbean Dollar	XCD
DOMINICAN REPUBLIC (THE)	Dominican Peso	DOP
ECUADOR	US Dollar	USD
EGYPT	Egyptian Pound	EGP
EL SALVADOR	El Salvador Colon	SVC
EL SALVADOR	US Dollar	USD
EQUATORIAL GUINEA	CFA Franc BEAC	XAF
ERITREA	Nakfa	ERN
ESTONIA	Euro	EUR
ESWATINI	Lilangeni	SZL
ETHIOPIA	Ethiopian Birr	ETB
EUROPEAN UNION	Euro	EUR
FALKLAND ISLANDS (THE) [MALVINAS]	Falkland Islands Pound	FKP
FAROE ISLANDS (THE)	Danish Krone	DKK
FIJI	Fiji Dollar	FJD
FINLAND	Euro	EUR
FRANCE	Euro	EUR
FRENCH GUIANA	Euro	EUR
FRENCH POLYNESIA	CFP Franc	XPF
FRENCH SOUTHERN TERRITORIES (THE)	Euro	EUR
GABON	CFA Franc BEAC	XAF
GAMBIA (THE)	Dalasi	GMD
GEORGIA	Lari	GEL
GERMANY	Euro	EUR
GHANA	Ghana Cedi	GHS
GIBRALTAR	Gibraltar Pound	GIP
GREECE	Euro	EUR
GREENLAND	Danish Krone	DKK
GRENADA	East Caribbean Dollar	XCD
GUADELOUPE	Euro	EUR
GUAM	US Dollar	USD
GUATEMALA	Quetzal	GTQ
GUERNSEY	Pound Sterling	GBP
GUINEA	Guinean Franc	GNF
GUINEA-BISSAU	CFA Franc BCEAO	XOF
GUYANA	Guyana Dollar	GYD
HAITI	Gourde	HTG
HAITI	US Dollar	USD
HEARD ISLAND AND McDONALD ISLANDS	Australian Dollar	AUD
HOLY SEE (THE)	Euro	EUR
HONDURAS	Lempira	HNL
HONG KONG	Hong Kong Dollar	HKD
HUNGARY	Forint	HUF
ICELAND	Iceland Krona	ISK
INDIA	Indian Rupee	INR
INDONESIA	Rupiah	IDR
INTERNATIONAL MONETARY FUND (IMF) 	SDR (Special Drawing Right)	XDR
IRAN (ISLAMIC REPUBLIC OF)	Iranian Rial	IRR
IRAQ	Iraqi Dinar	IQD
IRELAND	Euro	EUR
ISLE OF MAN	Pound Sterling	GBP
ISRAEL	New Israeli Sheqel	ILS
ITALY	Euro	EUR
JAMAICA	Jamaican Dollar	JMD
JAPAN	Yen	JPY
JERSEY	Pound Sterling	GBP
JORDAN	Jordanian Dinar	JOD
KAZAKHSTAN	Tenge	KZT
KENYA	Kenyan Shilling	KES
KIRIBATI	Australian Dollar	AUD
KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)	North Korean Won	KPW
KOREA (THE REPUBLIC OF)	Won	KRW
KUWAIT	Kuwaiti Dinar	KWD
KYRGYZSTAN	Som	KGS
LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)	Lao Kip	LAK
LATVIA	Euro	EUR
LEBANON	Lebanese Pound	LBP
LESOTHO	Loti	LSL
LESOTHO	Rand	ZAR
LIBERIA	Liberian Dollar	LRD
LIBYA	Libyan Dinar	LYD
LIECHTENSTEIN	Swiss Franc	CHF
LITHUANIA	Euro	EUR
LUXEMBOURG	Euro	EUR
MACAO	Pataca	MOP
NORTH MACEDONIA	Denar	MKD
MADAGASCAR	Malagasy Ariary	MGA
MALAWI	Malawi Kwacha	MWK
MALAYSIA	Malaysian Ringgit	MYR
MALDIVES	Rufiyaa	MVR
MALI	CFA Franc BCEAO	XOF
MALTA	Euro	EUR
MARSHALL ISLANDS (THE)	US Dollar	USD
MARTINIQUE	Euro	EUR
MAURITANIA	Ouguiya	MRU
MAURITIUS	Mauritius Rupee	MUR
MAYOTTE	Euro	EUR
MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP	ADB Unit of Account	XUA
MEXICO	Mexican Peso	MXN
MEXICO	Mexican Unidad de Inversion (UDI)	MXV
MICRONESIA (FEDERATED STATES OF)	US Dollar	USD
MOLDOVA (THE REPUBLIC OF)	Moldovan Leu	MDL
MONACO	Euro	EUR
MONGOLIA	Tugrik	MNT
MONTENEGRO	Euro	EUR
MONTSERRAT	East Caribbean Dollar	XCD
MOROCCO	Moroccan Dirham	MAD
MOZAMBIQUE	Mozambique Metical	MZN
MYANMAR	Kyat	MMK
NAMIBIA	Namibia Dollar	NAD
NAMIBIA	Rand	ZAR
NAURU	Australian Dollar	AUD
NEPAL	Nepalese Rupee	NPR
NETHERLANDS (THE)	Euro	EUR
NEW CALEDONIA	CFP Franc	XPF
NEW ZEALAND	New Zealand Dollar	NZD
NICARAGUA	Cordoba Oro	NIO
NIGER (THE)	CFA Franc BCEAO	XOF
NIGERIA	Naira	NGN
NIUE	New Zealand Dollar	NZD
NORFOLK ISLAND	Australian Dollar	AUD
NORTHERN MARIANA ISLANDS (THE)	US Dollar	USD
NORWAY	Norwegian Krone	NOK
OMAN	Rial Omani	OMR
PAKISTAN	Pakistan Rupee	PKR
PALAU	US Dollar	USD
PANAMA	Balboa	PAB
PANAMA	US Dollar	USD
PAPUA NEW GUINEA	Kina	PGK
PARAGUAY	Guarani	PYG
PERU	Sol	PEN
PHILIPPINES (THE)	Philippine Peso	PHP
PITCAIRN	New Zealand Dollar	NZD
POLAND	Zloty	PLN
PORTUGAL	Euro	EUR
PUERTO RICO	US Dollar	USD
QATAR	Qatari Rial	QAR
RÉUNION	Euro	EUR
ROMANIA	Romanian Leu	RON
RUSSIAN FEDERATION (THE)	Russian Ruble	RUB
RWANDA	Rwanda Franc	RWF
SAINT BARTHÉLEMY	Euro	EUR
SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA	Saint Helena Pound	SHP
SAINT KITTS AND NEVIS	East Caribbean Dollar	XCD
SAINT LUCIA	East Caribbean Dollar	XCD
SAINT MARTIN (FRENCH PART)	Euro	EUR
SAINT PIERRE AND MIQUELON	Euro	EUR
SAINT VINCENT AND THE GRENADINES	East Caribbean Dollar	XCD
SAMOA	Tala	WST
SAN MARINO	Euro	EUR
SAO TOME AND PRINCIPE	Dobra	STN
SAUDI ARABIA	Saudi Riyal	SAR
SENEGAL	CFA Franc BCEAO	XOF
SERBIA	Serbian Dinar	RSD
SEYCHELLES	Seychelles Rupee	SCR
SIERRA LEONE	Leone	SLL
SIERRA LEONE	Leone	SLE
SINGAPORE	Singapore Dollar	SGD
SINT MAARTEN (DUTCH PART)	Netherlands Antillean Guilder	ANG
SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS "SUCRE"	Sucre	XSU
SLOVAKIA	Euro	EUR
SLOVENIA	Euro	EUR
SOLOMON ISLANDS	Solomon Islands Dollar	SBD
SOMALIA	Somali Shilling	SOS
SOUTH AFRICA	Rand	ZAR
SOUTH SUDAN	South Sudanese Pound	SSP
SPAIN	Euro	EUR
SRI LANKA	Sri Lanka Rupee	LKR
SUDAN (THE)	Sudanese Pound	SDG
SURINAME	Surinam Dollar	SRD
SVALBARD AND JAN MAYEN	Norwegian Krone	NOK
SWEDEN	Swedish Krona	SEK
SWITZERLAND	Swiss Franc	CHF
SWITZERLAND	WIR Euro	CHE
SWITZERLAND	WIR Franc	CHW
SYRIAN ARAB REPUBLIC	Syrian Pound	SYP
TAIWAN (PROVINCE OF CHINA)	New Taiwan Dollar	TWD
TAJIKISTAN	Somoni	TJS
TANZANIA, UNITED REPUBLIC OF	Tanzanian Shilling	TZS
THAILAND	Baht	THB
TIMOR-LESTE	US Dollar	USD
TOGO	CFA Franc BCEAO	XOF
TOKELAU	New Zealand Dollar	NZD
TONGA	Pa’anga	TOP
TRINIDAD AND TOBAGO	Trinidad and Tobago Dollar	TTD
TUNISIA	Tunisian Dinar	TND
TÜRKİYE	Turkish Lira	TRY
TURKMENISTAN	Turkmenistan New Manat	TMT
TURKS AND CAICOS ISLANDS (THE)	US Dollar	USD
TUVALU	Australian Dollar	AUD
UGANDA	Uganda Shilling	UGX
UKRAINE	Hryvnia	UAH
UNITED ARAB EMIRATES (THE)	UAE Dirham	AED
UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)	Pound Sterling	GBP
UNITED STATES MINOR OUTLYING ISLANDS (THE)	US Dollar	USD
UNITED STATES OF AMERICA (THE)	US Dollar	USD
UNITED STATES OF AMERICA (THE)	US Dollar (Next day)	USN
URUGUAY	Peso Uruguayo	UYU
URUGUAY	Uruguay Peso en Unidades Indexadas (UI)	UYI
URUGUAY	Unidad Previsional	UYW
UZBEKISTAN	Uzbekistan Sum	UZS
VANUATU	Vatu	VUV
VENEZUELA (BOLIVARIAN REPUBLIC OF)	Bolívar Soberano	VES
VENEZUELA (BOLIVARIAN REPUBLIC OF)	Bolívar Soberano	VED
VIET NAM	Dong	VND
VIRGIN ISLANDS (BRITISH)	US Dollar	USD
VIRGIN ISLANDS (U.S.)	US Dollar	USD
WALLIS AND FUTUNA	CFP Franc	XPF
WESTERN SAHARA	Moroccan Dirham	MAD
YEMEN	Yemeni Rial	YER
ZAMBIA	Zambian Kwacha	ZMW
ZIMBABWE	Zimbabwe Dollar	ZWL
ZZ01_Bond Markets Unit European_EURCO	Bond Markets Unit European Composite Unit (EURCO)	XBA
ZZ02_Bond Markets Unit European_EMU-6	Bond Markets Unit European Monetary Unit (E.M.U.-6)	XBB
ZZ03_Bond Markets Unit European_EUA-9	Bond Markets Unit European Unit of Account 9 (E.U.A.-9)	XBC
ZZ04_Bond Markets Unit European_EUA-17	Bond Markets Unit European Unit of Account 17 (E.U.A.-17)	XBD
ZZ06_Testing_Code	Codes specifically reserved for testing purposes	XTS
ZZ07_No_Currency	The codes assigned for transactions where no currency is involved	XXX
ZZ08_Gold	Gold	XAU
ZZ09_Palladium	Palladium	XPD
ZZ10_Platinum	Platinum	XPT
ZZ11_Silver	Silver	XAG`

	d := `
	[{"name":"美元","abbreviation":"USD","symbol":"$"},
        {"name":"欧元","abbreviation":"EUR","symbol":"€"},
        {"name":"日圆","abbreviation":"JPY","symbol":"¥"},
        {"name":"英镑","abbreviation":"GBP","symbol":"£"},
        {"name":"澳大利亚元","abbreviation":"AUD","symbol":"$"},
        {"name":"加拿大元","abbreviation":"CAD","symbol":"$"},
        {"name":"瑞士法郎","abbreviation":"CHF","symbol":"₣"},
        {"name":"人民币","abbreviation":"CNY","symbol":"¥"},
        {"name":"瑞典克朗","abbreviation":"SEK","symbol":"kr"},
        {"name":"新西兰元","abbreviation":"NZD","symbol":"$"},
        {"name":"墨西哥比索","abbreviation":"MXN","symbol":"$"},
        {"name":"新加坡元","abbreviation":"SGD","symbol":"$"},
        {"name":"港元","abbreviation":"HKD","symbol":"$"},
        {"name":"挪威克朗","abbreviation":"NOK","symbol":"kr"},
        {"name":"韩元","abbreviation":"KRW","symbol":"₩"},
        {"name":"土耳其里拉","abbreviation":"TRY","symbol":"₺"},
        {"name":"俄罗斯卢布","abbreviation":"RUB","symbol":"₽"},
        {"name":"印度卢比","abbreviation":"INR","symbol":"₹"},
        {"name":"巴西雷亚尔","abbreviation":"BRL","symbol":"R$"},
        {"name":"南非兰特","abbreviation":"ZAR","symbol":"R"},
        {"name":"菲律宾比索","abbreviation":"PHP","symbol":"₱"},
        {"name":"捷克克朗","abbreviation":"CZK","symbol":"Kč"},
        {"name":"印尼盾","abbreviation":"IDR","symbol":"Rp"},
        {"name":"马来西亚林吉特","abbreviation":"MYR","symbol":"RM"},
        {"name":"匈牙利福林","abbreviation":"HUF","symbol":"Ft"},
        {"name":"冰岛克朗","abbreviation":"ISK","symbol":"kr"},
        {"name":"克罗地亚库纳","abbreviation":"HRK","symbol":"kn"},
        {"name":"保加利亚列弗","abbreviation":"BGN","symbol":"лв"},
        {"name":"罗马尼亚列伊","abbreviation":"RON","symbol":"lei"},
        {"name":"丹麦克朗","abbreviation":"DKK","symbol":"kr"},
        {"name":"泰铢","abbreviation":"THB","symbol":"฿"},
        {"name":"波兰兹罗提","abbreviation":"PLN","symbol":"zł"},
        {"name":"以色列新谢克尔","abbreviation":"ILS","symbol":"₪"},
        {"name":"阿联酋迪拉姆","abbreviation":"AED","symbol":"د.إ"},
        {"name":"阿根廷比索","abbreviation":"ARS","symbol":"$"},
        {"name":"巴哈马元","abbreviation":"BSD","symbol":"$"},
        {"name":"智利比索","abbreviation":"CLP","symbol":"$"},
        {"name":"哥伦比亚比索","abbreviation":"COP","symbol":"$"},
        {"name":"多米尼加比索","abbreviation":"DOP","symbol":"$"},
        {"name":"埃及镑","abbreviation":"EGP","symbol":"£"},
        {"name":"斐济元","abbreviation":"FJD","symbol":"$"},
        {"name":"危地马拉格查尔","abbreviation":"GTQ","symbol":"Q"},
        {"name":"哈萨克斯坦坚戈","abbreviation":"KZT","symbol":"₸"},
        {"name":"巴拿马巴波亚","abbreviation":"PAB","symbol":"B/."},
        {"name":"秘鲁新索尔","abbreviation":"PEN","symbol":"S/."},
        {"name":"巴基斯坦卢比","abbreviation":"PKR","symbol":"Rs"},
        {"name":"巴拉圭瓜拉尼","abbreviation":"PYG","symbol":"Gs"},
        {"name":"沙特里亚尔","abbreviation":"SAR","symbol":"﷼"},
        {"name":"新台币","abbreviation":"TWD","symbol":"T$"},
        {"name":"乌克兰格里夫纳","abbreviation":"UAH","symbol":"₴"},
        {"name":"乌拉圭比索","abbreviation":"UYU","symbol":"$U"},
        {"name":"越南盾","abbreviation":"VND","symbol":"Đ"},
		{"name":"阿富汗尼","abbreviation":"AFN","symbol":"؋"},
		{"name":"亚美尼亚德拉姆","abbreviation":"AMD","symbol":"֏"},
		{"name":"荷属安的列斯盾","abbreviation":"ANG","symbol":"NAƒ"},
		{"name":"安哥拉宽扎","abbreviation":"AOA","symbol":"Kz"},
		{"name":"阿鲁巴弗罗林","abbreviation":"AWG","symbol":"Afl."},
		{"name":"阿塞拜疆马纳特","abbreviation":"AZN","symbol":"₼"},
		{"name":"可兑换马克","abbreviation":"BAM","symbol":"KM"},
		{"name":"巴巴多斯元","abbreviation":"BBD","symbol":"$"},
		{"name":"孟加拉塔卡","abbreviation":"BDT","symbol":"৳"},
		{"name":"布隆迪法郎","abbreviation":"BIF","symbol":"Fr"},
		{"name":"文莱令吉","abbreviation":"BND","symbol":"B$"},
		{"name":"玻利维亚诺","abbreviation":"BOB","symbol":"Bs."},
		{"name":"不丹努尔特鲁姆","abbreviation":"BTN","symbol":"Nu."},
		{"name":"博茨瓦纳普拉","abbreviation":"BWP","symbol":"P"},
		{"name":"白俄罗斯卢布","abbreviation":"BYN","symbol":"Br"},
		{"name":"伯利兹元","abbreviation":"BZD","symbol":"BZ$"},
		{"name":"刚果法郎","abbreviation":"CDF","symbol":"FC"},
		{"name":"哥斯达黎加科朗","abbreviation":"CRC","symbol":"₡"},
		{"name":"古巴比索","abbreviation":"CUP","symbol":"$"},
		{"name":"佛得角埃斯库多","abbreviation":"CVE","symbol":"$"},
		{"name":"吉布提法郎","abbreviation":"DJF","symbol":"Fdj"},
		{"name":"阿尔及利亚第纳尔","abbreviation":"DZD","symbol":"د.ج"},
		{"name":"厄立特里亚纳克法","abbreviation":"ERN","symbol":"Nkf"},
		{"name":"埃塞俄比亚比尔","abbreviation":"ETB","symbol":"Br"},
		{"name":"福克兰群岛镑","abbreviation":"FKP","symbol":"£"},
		{"name":"格鲁吉亚拉里","abbreviation":"GEL","symbol":"₾"},
		{"name":"加纳赛地","abbreviation":"GHS","symbol":"GH₵"},
		{"name":"直布罗陀镑","abbreviation":"GIP","symbol":"£"},
		{"name":"冈比亚达拉西","abbreviation":"GMD","symbol":"D"},
		{"name":"几内亚法郎","abbreviation":"GNF","symbol":"FG"},
		{"name":"圭亚那元","abbreviation":"GYD","symbol":"$"},
		{"name":"洪都拉斯伦皮拉","abbreviation":"HNL","symbol":"L"},
		{"name":"伊拉克第纳尔","abbreviation":"IQD","symbol":"ع.د"},
		{"name":"伊朗里亚尔","abbreviation":"IRR","symbol":"﷼"},
		{"name":"牙买加元","abbreviation":"JMD","symbol":"J$"},
		{"name":"约旦第纳尔","abbreviation":"JOD","symbol":"د.ا"},
		{"name":"肯尼亚先令","abbreviation":"KES","symbol":"KSh"},
		{"name":"吉尔吉斯斯坦索姆","abbreviation":"KGS","symbol":"лв"},
		{"name":"柬埔寨瑞尔","abbreviation":"KHR","symbol":"៛"},
		{"name":"科威特第纳尔","abbreviation":"KWD","symbol":"د.ك"},
		{"name":"开曼群岛元","abbreviation":"KYD","symbol":"$"},
		{"name":"老挝基普","abbreviation":"LAK","symbol":"₭"},
		{"name":"黎巴嫩镑","abbreviation":"LBP","symbol":"ل.ل"},
		{"name":"斯里兰卡卢比","abbreviation":"LKR","symbol":"₨"},
		{"name":"利比里亚元","abbreviation":"LRD","symbol":"L$"},
		{"name":"利比亚第纳尔","abbreviation":"LYD","symbol":"د.ل"},
		{"name":"摩洛哥迪拉姆","abbreviation":"MAD","symbol":"د.م."},
		{"name":"摩尔多瓦列伊","abbreviation":"MDL","symbol":"L"},
		{"name":"马达加斯加阿里亚里","abbreviation":"MGA","symbol":"Ar"},
		{"name":"北马其顿代纳尔","abbreviation":"MKD","symbol":"ден"},
		{"name":"缅甸元","abbreviation":"MMK","symbol":"K"},
		{"name":"蒙古图格里克","abbreviation":"MNT","symbol":"₮"},
		{"name":"澳门币","abbreviation":"MOP","symbol":"MOP$"},
		{"name":"毛里塔尼亚乌吉亚","abbreviation":"MRU","symbol":"UM"},
		{"name":"毛里求斯卢比","abbreviation":"MUR","symbol":"₨"},
		{"name":"马尔代夫卢菲亚","abbreviation":"MVR","symbol":"Rf"},
		{"name":"马拉维克瓦查","abbreviation":"MWK","symbol":"MK"},
		{"name":"莫桑比克梅蒂卡尔","abbreviation":"MZN","symbol":"MT"},
		{"name":"纳米比亚元","abbreviation":"NAD","symbol":"N$"},
		{"name":"尼日利亚奈拉","abbreviation":"NGN","symbol":"₦"},
		{"name":"尼加拉瓜科多巴","abbreviation":"NIO","symbol":"C$"},
		{"name":"尼泊尔卢比","abbreviation":"NPR","symbol":"रू"},
		{"name":"阿曼里亚尔","abbreviation":"OMR","symbol":"﷼"},
		{"name":"巴布亚新几内亚基那","abbreviation":"PGK","symbol":"K"},
		{"name":"卡塔尔里亚尔","abbreviation":"QAR","symbol":"﷼"},
		{"name":"塞尔维亚第纳尔","abbreviation":"RSD","symbol":"дин"},
		{"name":"卢旺达法郎","abbreviation":"RWF","symbol":"R₣"},
		{"name":"所罗门群岛元","abbreviation":"SBD","symbol":"SI$"},
		{"name":"塞舌尔卢比","abbreviation":"SCR","symbol":"₨"},
		{"name":"苏丹镑","abbreviation":"SDG","symbol":"£"},
		{"name":"圣赫勒拿镑","abbreviation":"SHP","symbol":"£"},
		{"name":"索马里先令","abbreviation":"SOS","symbol":"S"},
		{"name":"苏里南元","abbreviation":"SRD","symbol":"Sr$"},
		{"name":"南苏丹镑","abbreviation":"SSP","symbol":"£"},
		{"name":"圣多美和普林西比多布拉","abbreviation":"STN","symbol":"Db"},
		{"name":"叙利亚镑","abbreviation":"SYP","symbol":"LS"},
		{"name":"斯威士兰里兰吉尼","abbreviation":"SZL","symbol":"L"},
		{"name":"塔吉克斯坦索莫尼","abbreviation":"TJS","symbol":"смн"},
		{"name":"土库曼斯坦新马纳特","abbreviation":"TMT","symbol":"m"},
		{"name":"突尼斯第纳尔","abbreviation":"TND","symbol":"د.ت"},
		{"name":"汤加帕安加","abbreviation":"TOP","symbol":"T$"},
		{"name":"特立尼达和多巴哥元","abbreviation":"TTD","symbol":"TT$"},
		{"name":"坦桑尼亚先令","abbreviation":"TZS","symbol":"TSh"},
		{"name":"乌干达先令","abbreviation":"UGX","symbol":"USh"},
		{"name":"乌兹别克斯坦苏姆","abbreviation":"UZS","symbol":"so'm"},
		{"name":"委内瑞拉玻利瓦尔","abbreviation":"VES","symbol":"Bs."},
		{"name":"瓦努阿图瓦图","abbreviation":"VUV","symbol":"Vt"},
		{"name":"萨摩亚塔拉","abbreviation":"WST","symbol":"WS$"},
		{"name":"也门里亚尔","abbreviation":"YER","symbol":"ر.ي"},
		{"name":"赞比亚克瓦查","abbreviation":"ZMW","symbol":"K"},
		{"name":"津巴布韦元","abbreviation":"ZWL","symbol":"$"}]
	`
	// 调用函数解析  Symbol
	currencyInfos, err := parseCurrencyInfo(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rateList, err := parseRateInfo(d)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 把currencyInfos的内容转变为map[stinrg]CurrencyInfo
	// currencyInfoMap := make(map[string]CurrencyInfo)
	for _, info := range currencyInfos {
		CurrencyInfoMap[info.Code] = info
	}

	//把rateList的Abbreviation与currencyInfoMap的Code对应起来,把rateList的Symbol付值给currencyInfoMap的Symbol
	for _, rate := range rateList {
		info, ok := CurrencyInfoMap[rate.Abbreviation]
		if ok {
			info.Symbol = rate.Symbol
			CurrencyInfoMap[rate.Abbreviation] = info
		}
	}
}

func GetData() {
	fmt.Println("开始从汇率网站获取数据,更新数据")
	//从汇率网获取的汇率数据
	data, err := getExchangData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	Lock.Lock()         // 获取写锁
	defer Lock.Unlock() // 释放写锁，延迟到函数结束

	for key, value := range data.Rates {
		info, ok := CurrencyInfoMap[key]
		if ok {
			info.ExchangeRate = value
			info.UpdateTime = "e" + time.Now().Local().Format("2006-01-02 15:04:05")
			CurrencyInfoMap[key] = info
		} else {
			fmt.Println("没有找着的CODE:", key)
		}
	}

	CurrencyInfoList = make([]CurrencyInfo, 0, cap(CurrencyInfoList))
	for _, value := range CurrencyInfoMap {
		if value.ExchangeRate > 0 {
			CurrencyInfoList = append(CurrencyInfoList, value)

		}
	}
	fmt.Println("从汇率网站获取数据,更新数据--完毕")

}

// CurrencyInfo 定义货币信息的结构体
type CurrencyInfo struct {
	Country      string  // 国家或地区
	Currency     string  // 货币名称
	Code         string  // 货币代码
	Country_code string  // 国家代码 2位字符
	Symbol       string  // 货币符号
	ExchangeRate float64 // 汇率
	IsVirtual    bool    // 是否虚拟货币
	FlagURL      string  //加密货币的图标url
	UpdateTime   string  // 更新时间,同时记录是哪个渠道更新的 e代表汇率网站,g代表google sheets
}

// parseCurrencyInfo 解析多行字符串，并返回CurrencyInfo的切片
func parseCurrencyInfo(input string) ([]CurrencyInfo, error) {
	lines := strings.Split(input, "\n") // 按换行符分割字符串
	var result []CurrencyInfo           // 创建空切片，用于存放结果

	for _, line := range lines {
		// 使用制表符分割每行为三个部分
		columns := strings.Split(line, "\t")
		if len(columns) != 3 { // 检查是否有三列
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		// 将分割后的列作为CurrencyInfo属性
		info := CurrencyInfo{
			Country:      columns[0],
			Currency:     columns[1],
			Code:         columns[2],
			Country_code: strings.ToLower(columns[2][:2]),
		}
		result = append(result, info) // 追加到结果切片中
	}
	return result, nil // 返回结果切片和nil错误
}

type RateInfo struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Symbol       string `json:"symbol"`
	FlagURL      string `json:"flagURL"`
}

func parseRateInfo(input string) ([]RateInfo, error) {
	var result []RateInfo
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, fmt.Errorf("invalid json: %s", err)
	}
	return result, nil
}

// 定义一个汇率信息的结构体
type ExchangeInfo struct {
	BaseCode           string             `json:"base_oode"`
	Documentation      string             `json:"documentation"`
	Provider           string             `json:"provider"`
	Result             string             `json:"result"`
	TermsOfUse         string             `json:"terms_of_use"`
	TimeEolUnix        int64              `json:"time_eol_unix"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64              `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string             `json:"time_next_update_utc"`
	Rates              map[string]float64 `json:"rates"`
}

func getExchangData() (*ExchangeInfo, error) {
	url_ := "https://open.er-api.com/v6/latest/USD" // 假设要GET请求的URL
	// proxyURL, err := url.Parse("http://localhost:1080")
	// if err != nil {
	// 	panic(err)
	// }

	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }

	// client := &http.Client{
	// 	Transport: transport,
	// }
	// 发起HTTP GET请求
	resp, err := http.Get(url_)
	if err != nil {
		// 处理错误
		fmt.Println("Error sending request to API endpoint. ", err)
		return nil, err
	}
	defer resp.Body.Close() // 确保在函数结束时关闭响应体

	// 用io.ReadAll读取响应体中的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		fmt.Println("Error reading response body. ", err)
		return nil, err
	}

	// 如果响应的内容类型是JSON，则可以进行解析
	if resp.Header.Get("Content-Type") == "application/json" {
		var data ExchangeInfo
		// 解析JSON数据到结构体
		if err := json.Unmarshal(body, &data); err != nil {
			// 处理错误
			fmt.Println("Error unmarshaling JSON to struct. ", err)
			return nil, err
		}
		if data.Result != "success" {
			fmt.Println("获取汇率出错")
			return nil, fmt.Errorf("获取汇率出错")
		}
		fmt.Println("汇率更新时间:", data.TimeLastUpdateUtc)
		//把data.TimeLastUpdateUtc时间字符串，转变为本地时区字符串
		utcTimestamp := int64(data.TimeLastUpdateUnix)
		utcTime := time.Unix(utcTimestamp, 0)
		localTime := utcTime.In(time.Local)
		localTimeString := localTime.Format("2006-01-02 15:04:05")

		TimeLastUpdate = localTimeString
		TimeLastGet = time.Now().Format("2006-01-02 15:04:05")
		fmt.Println("汇率数据获取完成.")
		return &data, nil
	} else {
		// 如果响应不是JSON，则返回一个错误
		return nil, fmt.Errorf("response content-type is not application/json")
	}
}
