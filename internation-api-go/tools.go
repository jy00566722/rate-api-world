package main

import (
	"encoding/json"
	"fmt"
)

func Json2stust() {
	fmt.Println("哈哈哈")
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
		{"name":"哥斯达黎加科朗","abbreviation":"CUP","symbol":"$MN"},
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
	rateList, err := parseRateInfo(d)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	type CurrencyInfoType struct {
		// FlagURL string `json:"flagURL,omitempty"`
		Name string `json:"name,omitempty"`
		Rate string `json:"rate"`
		// RateNm string `json:"ratenm"`
		Scur   string `json:"scur"`
		Status string `json:"status"`
		Symbol string `json:"symbol,omitempty"`
		Tcur   string `json:"tcur"`
		Update string `json:"update"`
		Hot    int    `json:"hot,omitempty"`
	}
	type CurrencyInfoMapType map[string]CurrencyInfoType
	//把rateList转换成CurrencyInfoMapType
	currencyInfoMap := make(CurrencyInfoMapType)
	for _, v := range rateList {
		currencyInfoMap[v.Abbreviation] = CurrencyInfoType{
			Name:   v.Abbreviation + " - " + v.Name,
			Rate:   "1",
			Scur:   "CNY",
			Status: "ALREADY",
			Symbol: v.Symbol,
			Tcur:   v.Abbreviation,
			Update: "2023-12-26 12:58:01",
		}
	}
	// fmt.Println(currencyInfoMap)
	//把currencyInfoMap转换成json
	jsonStr, err := json.Marshal(currencyInfoMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonStr))

}
