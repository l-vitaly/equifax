package equifax

type DocType uint32

const (
	// Паспорт гражданина Российской Федерации – для гражданина Российской Федерации, достигшего 14 лет
	DocType1 DocType = 1
	// Свидетельство органов ЗАГСа, органа исполнительной власти или органа местного самоуправления о рождении гражданина – для гражданина Российской Федерации, не достигшего 14 лет
	DocType2 DocType = 2
	// Удостоверение личности – для офицеров, прапорщиков и мичманов
	DocType3 DocType = 3
	// Военный билет – для сержантов, старшин, солдат и матросов, а также курсантов военных образовательных учреждений профессионального образования
	DocType4 DocType = 4
	// Паспорт моряка – для граждан Российской Федерации, работающих на судах заграничного плавания или на иностранных судах, курсантов учебных заведений
	DocType5 DocType = 5
	// Паспорт иностранного гражданина либо иной документ, установленный федеральным законом или признаваемый в соответствии с международным договором Российской Федерации в качестве документа, удостоверяющего личность иностранного гражданина
	DocType6 DocType = 6
	// Документ, выданный иностранным государством и признаваемый в соответствии с международным договором Российской Федерации в качестве документа, удостоверяющего личность лица без гражданства
	DocType7 DocType = 7
	// Разрешение на временное проживание лица без гражданства
	DocType8 DocType = 8
	// Вид на жительство лица без гражданства
	DocType9 DocType = 9
	// Иные документы, предусмотренные федеральным законом или признаваемые в соответствии с международным договором Российской Федерации в качестве документов, удостоверяющих личность лица без гражданства
	DocType10 DocType = 10
	// Свидетельство о регистрации ходатайства о признании иммигранта беженцем
	DocType11 DocType = 11
	// Удостоверение беженца
	DocType12 DocType = 12
	// Временное удостоверение личности гражданина
	DocType13 DocType = 13
	// Иные документы, выдаваемые уполномоченными органами
	DocType14 DocType = 14
	// Паспорт гражданина СССР
	DocType15 DocType = 15
	// Нет (Значение может передаваться только для ранее выданного документа)
	DocType98 DocType = 98
	// Значение не передается (Значение может передаваться только для ранее выданного документа)
	DocType99 DocType = 99
)

type Sex uint32

const (
	// Мужской
	SexType1 Sex = 1
	// Женский
	SexType2 Sex = 2
	// Значение не передается
	SexType99 Sex = 99
)

type Citizenship string

func (c Citizenship) ToString() string {
	return string(c)
}

const (
	// Австралия
	CitizenshipTypeAU Citizenship = "AU"
	// Австрия
	CitizenshipTypeAT Citizenship = "AT"
	// Азербайджан
	CitizenshipTypeAZ Citizenship = "AZ"
	// Аландские острова
	CitizenshipTypeAX Citizenship = "AX"
	// Албания
	CitizenshipTypeAL Citizenship = "AL"
	// Алжир
	CitizenshipTypeDZ Citizenship = "DZ"
	// Американское Самоа
	CitizenshipTypeAS Citizenship = "AS"
	// Ангилья
	CitizenshipTypeAI Citizenship = "AI"
	// Ангола
	CitizenshipTypeAO Citizenship = "AO"
	// Андорра
	CitizenshipTypeAD Citizenship = "AD"
	// Антигуа и Барбуда
	CitizenshipTypeAG Citizenship = "AG"
	// Аргентина
	CitizenshipTypeAR Citizenship = "AR"
	// Армения
	CitizenshipTypeAM Citizenship = "AM"
	// Аруба
	CitizenshipTypeAW Citizenship = "AW"
	// Афганистан
	CitizenshipTypeAF Citizenship = "AF"
	// Багамы
	CitizenshipTypeBS Citizenship = "BS"
	// Бангладеш
	CitizenshipTypeBD Citizenship = "BD"
	// Барбадос
	CitizenshipTypeBB Citizenship = "BB"
	// Бахрейн
	CitizenshipTypeBH Citizenship = "BH"
	// Беларусь
	CitizenshipTypeBY Citizenship = "BY"
	// Белиз
	CitizenshipTypeBZ Citizenship = "BZ"
	// Бельгия
	CitizenshipTypeBE Citizenship = "BE"
	// Бенин
	CitizenshipTypeBJ Citizenship = "BJ"
	// Бермуды
	CitizenshipTypeBM Citizenship = "BM"
	// Болгария
	CitizenshipTypeBG Citizenship = "BG"
	// Боливия
	CitizenshipTypeBO Citizenship = "BO"
	// Босния и Герцеговина
	CitizenshipTypeBA Citizenship = "BA"
	// Ботсвана
	CitizenshipTypeBW Citizenship = "BW"
	// Бразилия
	CitizenshipTypeBR Citizenship = "BR"
	// Британская территория в Индийском океане
	CitizenshipTypeIO Citizenship = "IO"
	// Бруней-Даруссалам
	CitizenshipTypeBN Citizenship = "BN"
	// Буве Остров
	CitizenshipTypeBV Citizenship = "BV"
	// Буркина Фасо
	CitizenshipTypeBF Citizenship = "BF"
	// Бурунди
	CitizenshipTypeBI Citizenship = "BI"
	// Бутан
	CitizenshipTypeBT Citizenship = "BT"
	// Вануату
	CitizenshipTypeVU Citizenship = "VU"
	// Ватикан
	CitizenshipTypeVA Citizenship = "VA"
	// Великобритания
	CitizenshipTypeGB Citizenship = "GB"
	// Венгрия
	CitizenshipTypeHU Citizenship = "HU"
	// Венесуэла
	CitizenshipTypeVE Citizenship = "VE"
	// Виргинские острова Британские
	CitizenshipTypeVG Citizenship = "VG"
	// Виргинские острова США
	CitizenshipTypeVI Citizenship = "VI"
	// Внешние малые острова США
	CitizenshipTypeUM Citizenship = "UM"
	// Восточный Тимор
	CitizenshipTypeTL Citizenship = "TL"
	// Вьетнам
	CitizenshipTypeVN Citizenship = "VN"
	// Габон
	CitizenshipTypeGA Citizenship = "GA"
	// Гайана
	CitizenshipTypeGY Citizenship = "GY"
	// Гаити
	CitizenshipTypeHT Citizenship = "HT"
	// Гамбия
	CitizenshipTypeGM Citizenship = "GM"
	// Гана
	CitizenshipTypeGH Citizenship = "GH"
	// Гваделупа
	CitizenshipTypeGP Citizenship = "GP"
	// Гватемала
	CitizenshipTypeGT Citizenship = "GT"
	// Гвиана
	CitizenshipTypeGF Citizenship = "GF"
	// Гвинея
	CitizenshipTypeGN Citizenship = "GN"
	// Гвинея-Бисау
	CitizenshipTypeGW Citizenship = "GW"
	// Германия
	CitizenshipTypeDE Citizenship = "DE"
	// Гернси
	CitizenshipTypeGG Citizenship = "GG"
	// Гибралтар
	CitizenshipTypeGI Citizenship = "GI"
	// Гондурас
	CitizenshipTypeHN Citizenship = "HN"
	// Гонконг
	CitizenshipTypeHK Citizenship = "HK"
	// Гренада
	CitizenshipTypeGD Citizenship = "GD"
	// Гренландия
	CitizenshipTypeGL Citizenship = "GL"
	// Греция
	CitizenshipTypeGR Citizenship = "GR"
	// Грузия
	CitizenshipTypeGE Citizenship = "GE"
	// Гуам
	CitizenshipTypeGU Citizenship = "GU"
	// Дания
	CitizenshipTypeDK Citizenship = "DK"
	// Джерси
	CitizenshipTypeJE Citizenship = "JE"
	// Джибути
	CitizenshipTypeDJ Citizenship = "DJ"
	// Доминика
	CitizenshipTypeDM Citizenship = "DM"
	// Доминиканская Республика
	CitizenshipTypeDO Citizenship = "DO"
	// Египет
	CitizenshipTypeEG Citizenship = "EG"
	// Замбия
	CitizenshipTypeZM Citizenship = "ZM"
	// Западная Сахара
	CitizenshipTypeEH Citizenship = "EH"
	// Зимбабве
	CitizenshipTypeZW Citizenship = "ZW"
	// Йемен
	CitizenshipTypeYE Citizenship = "YE"
	// Израиль
	CitizenshipTypeIL Citizenship = "IL"
	// Индия
	CitizenshipTypeIN Citizenship = "IN"
	// Индонезия
	CitizenshipTypeID Citizenship = "ID"
	// Иордания
	CitizenshipTypeJO Citizenship = "JO"
	// Ирак
	CitizenshipTypeIQ Citizenship = "IQ"
	// Иран
	CitizenshipTypeIR Citizenship = "IR"
	// Ирландия
	CitizenshipTypeIE Citizenship = "IE"
	// Исландия
	CitizenshipTypeIS Citizenship = "IS"
	// Испания
	CitizenshipTypeES Citizenship = "ES"
	// Италия
	CitizenshipTypeIT Citizenship = "IT"
	// Кабо-Верде
	CitizenshipTypeCV Citizenship = "CV"
	// Казахстан
	CitizenshipTypeKZ Citizenship = "KZ"
	// Каймановы острова
	CitizenshipTypeKY Citizenship = "KY"
	// Камбоджа
	CitizenshipTypeKH Citizenship = "KH"
	// Камерун
	CitizenshipTypeCM Citizenship = "CM"
	// Канада
	CitizenshipTypeCA Citizenship = "CA"
	// Катар
	CitizenshipTypeQA Citizenship = "QA"
	// Кения
	CitizenshipTypeKE Citizenship = "KE"
	// Кипр
	CitizenshipTypeCY Citizenship = "CY"
	// Киргизия
	CitizenshipTypeKG Citizenship = "KG"
	// Кирибати
	CitizenshipTypeKI Citizenship = "KI"
	// Китай
	CitizenshipTypeCN Citizenship = "CN"
	// КНДР
	CitizenshipTypeKP Citizenship = "KP"
	// Кокосовые (Килинг) острова
	CitizenshipTypeCC Citizenship = "CC"
	// Колумбия
	CitizenshipTypeCO Citizenship = "CO"
	// Коморы
	CitizenshipTypeKM Citizenship = "KM"
	// Конго (Демократическая Республика Конго)
	CitizenshipTypeCD Citizenship = "CD"
	// Конго (Республика Конго)
	CitizenshipTypeCG Citizenship = "CG"
	// Коста-Рика
	CitizenshipTypeCR Citizenship = "CR"
	// Кот-д'Ивуар
	CitizenshipTypeCI Citizenship = "CI"
	// Куба
	CitizenshipTypeCU Citizenship = "CU"
	// Кувейт
	CitizenshipTypeKW Citizenship = "KW"
	// Кука Острова
	CitizenshipTypeCK Citizenship = "CK"
	// Лаос
	CitizenshipTypeLA Citizenship = "LA"
	// Латвия
	CitizenshipTypeLV Citizenship = "LV"
	// Лесото
	CitizenshipTypeLS Citizenship = "LS"
	// Либерия
	CitizenshipTypeLR Citizenship = "LR"
	// Ливан
	CitizenshipTypeLB Citizenship = "LB"
	// Ливия
	CitizenshipTypeLY Citizenship = "LY"
	// Литва
	CitizenshipTypeLT Citizenship = "LT"
	// Лихтенштейн
	CitizenshipTypeLI Citizenship = "LI"
	// Люксембург
	CitizenshipTypeLU Citizenship = "LU"
	// Маврикий
	CitizenshipTypeMU Citizenship = "MU"
	// Мавритания
	CitizenshipTypeMR Citizenship = "MR"
	// Мадагаскар
	CitizenshipTypeMG Citizenship = "MG"
	// Майотта
	CitizenshipTypeYT Citizenship = "YT"
	// Макао
	CitizenshipTypeMO Citizenship = "MO"
	// Малави
	CitizenshipTypeMW Citizenship = "MW"
	// Малайзия
	CitizenshipTypeMY Citizenship = "MY"
	// Мали
	CitizenshipTypeML Citizenship = "ML"
	// Мальдивы
	CitizenshipTypeMV Citizenship = "MV"
	// Мальта
	CitizenshipTypeMT Citizenship = "MT"
	// Марокко
	CitizenshipTypeMA Citizenship = "MA"
	// Мартиника
	CitizenshipTypeMQ Citizenship = "MQ"
	// Маршалловы острова
	CitizenshipTypeMH Citizenship = "MH"
	// Мексика
	CitizenshipTypeMX Citizenship = "MX"
	// Мозамбик
	CitizenshipTypeMZ Citizenship = "MZ"
	// Молдавия
	CitizenshipTypeMD Citizenship = "MD"
	// Монако
	CitizenshipTypeMC Citizenship = "MC"
	// Монголия
	CitizenshipTypeMN Citizenship = "MN"
	// Монтсеррат
	CitizenshipTypeMS Citizenship = "MS"
	// Мьянма
	CitizenshipTypeMM Citizenship = "MM"
	// Намибия
	CitizenshipTypeNA Citizenship = "NA"
	// Науру
	CitizenshipTypeNR Citizenship = "NR"
	// Непал
	CitizenshipTypeNP Citizenship = "NP"
	// Нигер
	CitizenshipTypeNE Citizenship = "NE"
	// Нигерия
	CitizenshipTypeNG Citizenship = "NG"
	// Нидерландские Антильские о-ва
	CitizenshipTypeAN Citizenship = "AN"
	// Нидерланды
	CitizenshipTypeNL Citizenship = "NL"
	// Никарагуа
	CitizenshipTypeNI Citizenship = "NI"
	// Ниуэ
	CitizenshipTypeNU Citizenship = "NU"
	// Новая Зеландия
	CitizenshipTypeNZ Citizenship = "NZ"
	// Новая Каледония
	CitizenshipTypeNC Citizenship = "NC"
	// Норвегия
	CitizenshipTypeNO Citizenship = "NO"
	// Норфолк Остров
	CitizenshipTypeNF Citizenship = "NF"
	// Объединенные Арабские Эмираты
	CitizenshipTypeAE Citizenship = "AE"
	// Оман
	CitizenshipTypeOM Citizenship = "OM"
	// Остров Мэн
	CitizenshipTypeIM Citizenship = "IM"
	// Остров Херд и острова Макдональд
	CitizenshipTypeHM Citizenship = "HM"
	// Пакистан
	CitizenshipTypePK Citizenship = "PK"
	// Палау
	CitizenshipTypePW Citizenship = "PW"
	// Палестина
	CitizenshipTypePS Citizenship = "PS"
	// Панама
	CitizenshipTypePA Citizenship = "PA"
	// Папуа - Новая Гвинея
	CitizenshipTypePG Citizenship = "PG"
	// Парагвай
	CitizenshipTypePY Citizenship = "PY"
	// Перу
	CitizenshipTypePE Citizenship = "PE"
	// Питкэрн Острова
	CitizenshipTypePN Citizenship = "PN"
	// Польша
	CitizenshipTypePL Citizenship = "PL"
	// Португалия
	CitizenshipTypePT Citizenship = "PT"
	// Пуэрто-Рико
	CitizenshipTypePR Citizenship = "PR"
	// Республика Корея
	CitizenshipTypeKR Citizenship = "KR"
	// Республика Македония
	CitizenshipTypeMK Citizenship = "MK"
	// Реюньон Остров
	CitizenshipTypeRE Citizenship = "RE"
	// Рождества Остров
	CitizenshipTypeCX Citizenship = "CX"
	// Россия
	CitizenshipTypeRU Citizenship = "RU"
	// Руанда
	CitizenshipTypeRW Citizenship = "RW"
	// Румыния
	CitizenshipTypeRO Citizenship = "RO"
	// Сальвадор
	CitizenshipTypeSV Citizenship = "SV"
	// Самоа
	CitizenshipTypeWS Citizenship = "WS"
	// Сан-Марино
	CitizenshipTypeSM Citizenship = "SM"
	// Сан-Томе и Принсипи
	CitizenshipTypeST Citizenship = "ST"
	// Саудовская Аравия
	CitizenshipTypeSA Citizenship = "SA"
	// Свазиленд
	CitizenshipTypeSZ Citizenship = "SZ"
	// Святого Мартина Остров
	CitizenshipTypeMF Citizenship = "MF"
	// Святой Елены Остров
	CitizenshipTypeSH Citizenship = "SH"
	// Северные Марианские острова
	CitizenshipTypeMP Citizenship = "MP"
	// Сейшелы
	CitizenshipTypeSC Citizenship = "SC"
	// Сен-Бартельми
	CitizenshipTypeBL Citizenship = "BL"
	// Сенегал
	CitizenshipTypeSN Citizenship = "SN"
	// Сен-Пьер и Микелон
	CitizenshipTypePM Citizenship = "PM"
	// Сент-Винсент и Гренадины
	CitizenshipTypeVC Citizenship = "VC"
	// Сент-Китс и Невис
	CitizenshipTypeKN Citizenship = "KN"
	// Сент-Люсия
	CitizenshipTypeLC Citizenship = "LC"
	// Сербия
	CitizenshipTypeRS Citizenship = "RS"
	// Сингапур
	CitizenshipTypeSG Citizenship = "SG"
	// Сирийская Арабская Республика
	CitizenshipTypeSY Citizenship = "SY"
	// Словакия
	CitizenshipTypeSK Citizenship = "SK"
	// Словения
	CitizenshipTypeSI Citizenship = "SI"
	// Соломоновы острова
	CitizenshipTypeSB Citizenship = "SB"
	// Сомали
	CitizenshipTypeSO Citizenship = "SO"
	// Судан
	CitizenshipTypeSD Citizenship = "SD"
	// Суринам
	CitizenshipTypeSR Citizenship = "SR"
	// США
	CitizenshipTypeUS Citizenship = "US"
	// Сьерра-Леоне
	CitizenshipTypeSL Citizenship = "SL"
	// Таджикистан
	CitizenshipTypeTJ Citizenship = "TJ"
	// Тайвань (Китай)
	CitizenshipTypeTW Citizenship = "TW"
	// Таиланд
	CitizenshipTypeTH Citizenship = "TH"
	// Танзания
	CitizenshipTypeTZ Citizenship = "TZ"
	// Теркс и Кайкос
	CitizenshipTypeTC Citizenship = "TC"
	// Того
	CitizenshipTypeTG Citizenship = "TG"
	// Токелау
	CitizenshipTypeTK Citizenship = "TK"
	// Тонга
	CitizenshipTypeTO Citizenship = "TO"
	// Тринидад и Тобаго
	CitizenshipTypeTT Citizenship = "TT"
	// Тувалу
	CitizenshipTypeTV Citizenship = "TV"
	// Тунис
	CitizenshipTypeTN Citizenship = "TN"
	// Туркмения
	CitizenshipTypeTM Citizenship = "TM"
	// Турция
	CitizenshipTypeTR Citizenship = "TR"
	// Уганда
	CitizenshipTypeUG Citizenship = "UG"
	// Узбекистан
	CitizenshipTypeUZ Citizenship = "UZ"
	// Украина
	CitizenshipTypeUA Citizenship = "UA"
	// Уоллис и Футуна
	CitizenshipTypeWF Citizenship = "WF"
	// Уругвай
	CitizenshipTypeUY Citizenship = "UY"
	// Фарерские острова
	CitizenshipTypeFO Citizenship = "FO"
	// Федеративные Штаты Микронезии
	CitizenshipTypeFM Citizenship = "FM"
	// Фиджи
	CitizenshipTypeFJ Citizenship = "FJ"
	// Филиппины
	CitizenshipTypePH Citizenship = "PH"
	// Финляндия
	CitizenshipTypeFI Citizenship = "FI"
	// Фолклендские острова (Мальвинские)
	CitizenshipTypeFK Citizenship = "FK"
	// Франция
	CitizenshipTypeFR Citizenship = "FR"
	// Французская Полинезия
	CitizenshipTypePF Citizenship = "PF"
	// Французские Южные и Антарктические территории
	CitizenshipTypeTF Citizenship = "TF"
	// Хорватия
	CitizenshipTypeHR Citizenship = "HR"
	// Центральноафриканская Республика
	CitizenshipTypeCF Citizenship = "CF"
	// Чад
	CitizenshipTypeTD Citizenship = "TD"
	// Черногория
	CitizenshipTypeME Citizenship = "ME"
	// Чехия
	CitizenshipTypeCZ Citizenship = "CZ"
	// Чили
	CitizenshipTypeCL Citizenship = "CL"
	// Швейцария
	CitizenshipTypeCH Citizenship = "CH"
	// Швеция
	CitizenshipTypeSE Citizenship = "SE"
	// Шри-Ланка
	CitizenshipTypeLK Citizenship = "LK"
	// Эквадор
	CitizenshipTypeEC Citizenship = "EC"
	// Экваториальная Гвинея
	CitizenshipTypeGQ Citizenship = "GQ"
	// Эритрея
	CitizenshipTypeER Citizenship = "ER"
	// Эстония
	CitizenshipTypeEE Citizenship = "EE"
	// Эфиопия
	CitizenshipTypeET Citizenship = "ET"
	// Южная Георгия и Южные Сандвичевы острова
	CitizenshipTypeGS Citizenship = "GS"
	// Южно-Африканская Республика
	CitizenshipTypeZA Citizenship = "ZA"
	// Ямайка
	CitizenshipTypeJM Citizenship = "JM"
	// Япония
	CitizenshipTypeJP Citizenship = "JP"
	// Другая
	CitizenshipType98 Citizenship = "98"
	// Неизвестно
	CitizenshipType99 Citizenship = "99"
)

type Education uint32

const (
	// Начальная школа
	EducationType0 Education = 0
	// Средняя школа
	EducationType1 Education = 1
	// Специализированная средняя школа
	EducationType2 Education = 2
	// Незаконченное высшее образование
	EducationType3 Education = 3
	// Высшее образование
	EducationType4 Education = 4
	// Два и более высших образования
	EducationType5 Education = 5
	// Ученая степень
	EducationType6 Education = 6
	// Другое
	EducationType8 Education = 8
	// Не известно
	EducationType9 Education = 9
	// Значение не передается
	EducationType99 Education = 99
)

type Marital uint32

const (
	// Холост / Не замужем
	MaritalType0 Marital = 0
	// Женат / Замужем
	MaritalType1 Marital = 1
	// Разведен / Разведена
	MaritalType2 Marital = 2
	// Вдовец / Вдова
	MaritalType3 Marital = 3
	// Гражданский брак / Совместное проживание
	MaritalType4 Marital = 4
	// Не известно
	MaritalType9 Marital = 9
	// Значение не передается
	MaritalType99 Marital = 99
)

type EmployerSize uint32

const (
	// < 50 человек
	EmployerSizeType0 EmployerSize = 0
	// 50-100 человек
	EmployerSizeType1 EmployerSize = 1
	// 101-249 человек
	EmployerSizeType2 EmployerSize = 2
	// 250-499 человек
	EmployerSizeType3 EmployerSize = 3
	// > 500 человек
	EmployerSizeType4 EmployerSize = 4
	// Не известно (EMPTY)
	EmployerSizeType9 EmployerSize = 9
	// Значение не передается
	EmployerSizeType99 EmployerSize = 99
)

type BusinessIndustry uint32

const (
	// Промышленность и машиностроение
	BusinessIndustryType0 BusinessIndustry = 0
	// Сельское хозяйство
	BusinessIndustryType1 BusinessIndustry = 1
	// Строительство
	BusinessIndustryType2 BusinessIndustry = 2
	// Горное дело
	BusinessIndustryType3 BusinessIndustry = 3
	// Энергетика
	BusinessIndustryType4 BusinessIndustry = 4
	// Оптовая торговля
	BusinessIndustryType5 BusinessIndustry = 5
	// Финансовое дело и страхование
	BusinessIndustryType6 BusinessIndustry = 6
	// Здравоохранение
	BusinessIndustryType7 BusinessIndustry = 7
	// Социальная помощь
	BusinessIndustryType8 BusinessIndustry = 8
	// Денежные переводы и обмен валют
	BusinessIndustryType9 BusinessIndustry = 9
	// Искусство, развлечение, отдых
	BusinessIndustryType10 BusinessIndustry = 10
	// Казино и игорный бизнес
	BusinessIndustryType11 BusinessIndustry = 11
	// Торговля предметами искусства и антиквариатом
	BusinessIndustryType12 BusinessIndustry = 12
	// Aвиа-, авто- и железнодорожные перевозки, складское хранение
	BusinessIndustryType13 BusinessIndustry = 13
	// Административно-хозяйственные службы
	BusinessIndustryType14 BusinessIndustry = 14
	// Розничная торговля
	BusinessIndustryType15 BusinessIndustry = 15
	// Водо-, тепло- и энергоснабжение
	BusinessIndustryType16 BusinessIndustry = 16
	// Управление компанией
	BusinessIndustryType17 BusinessIndustry = 17
	// Предпринимательская деятельность
	BusinessIndustryType18 BusinessIndustry = 18
	// Научная, техническая и профессиональная деятельность
	BusinessIndustryType19 BusinessIndustry = 19
	// Образование
	BusinessIndustryType20 BusinessIndustry = 20
	// Торговля ювелирными украшениями и драгоценными металлами
	BusinessIndustryType21 BusinessIndustry = 21
	// Посреднические услуги по продаже и аренде недвижимости
	BusinessIndustryType22 BusinessIndustry = 22
	// Снабжение
	BusinessIndustryType23 BusinessIndustry = 23
	// Пресса, телевидение и радио
	BusinessIndustryType24 BusinessIndustry = 24
	// Государственное управление
	BusinessIndustryType25 BusinessIndustry = 25
	// Ресторан / Кафе
	BusinessIndustryType26 BusinessIndustry = 26
	// Бытовые услуги
	BusinessIndustryType27 BusinessIndustry = 27
	// Аудиторские услуги / консалтинг
	BusinessIndustryType28 BusinessIndustry = 28
	// Туризм
	BusinessIndustryType29 BusinessIndustry = 29
	// Юридические и нотариальные услуги
	BusinessIndustryType30 BusinessIndustry = 30
	// ТЭК, добывающая промышленность
	BusinessIndustryType31 BusinessIndustry = 31
	// Детективное и / или охранное предприятие
	BusinessIndustryType32 BusinessIndustry = 32
	// Гостиницы
	BusinessIndustryType33 BusinessIndustry = 33
	// Посредническая деятельность
	BusinessIndustryType34 BusinessIndustry = 34
	// Издательская деятельность / рекламная деятельность
	BusinessIndustryType35 BusinessIndustry = 35
	// Информатика, телекоммуникации
	BusinessIndustryType36 BusinessIndustry = 36
	// Легкая и пищевая промышленность
	BusinessIndustryType37 BusinessIndustry = 37
	// Государственные органы (в том числе правоохранительные)
	BusinessIndustryType38 BusinessIndustry = 38
	// Вооруженные силы
	BusinessIndustryType39 BusinessIndustry = 39
	// Салоны красоты, фитнес центры
	BusinessIndustryType40 BusinessIndustry = 40
	// Транспорт
	BusinessIndustryType41 BusinessIndustry = 41
	// Сборочное производство
	BusinessIndustryType42 BusinessIndustry = 42
	// Другое
	BusinessIndustryType97 BusinessIndustry = 97
	// Не известно
	BusinessIndustryType98 BusinessIndustry = 98
	// Значение не передается
	BusinessIndustryType99 BusinessIndustry = 99
)

type IncomeProof uint32

const (
	// Справка 2-НДФЛ
	IncomeProofType1 IncomeProof = 1
	// Автомобиль
	IncomeProofType2 IncomeProof = 2
	// Недвижимость
	IncomeProofType3 IncomeProof = 3
	// Выписка по счету зарплатной карты
	IncomeProofType4 IncomeProof = 4
	// Справка из ПФР / Выписка с пенсионного счета
	IncomeProofType5 IncomeProof = 5
	// Справка в свободной форме
	IncomeProofType97 IncomeProof = 97
	// Нет подтверждения
	IncomeProofType98 IncomeProof = 98
	// Значение не передается
	IncomeProofType99 IncomeProof = 99
)

type ProductType uint32

const (
	// Неизвестный тип кредита
	ProductTypeType0 ProductType = 0
	// Кредит на автомобиль
	ProductTypeType1 ProductType = 1
	// Лизинг
	ProductTypeType2 ProductType = 2
	// Ипотека
	ProductTypeType3 ProductType = 3
	// Кредитная карта
	ProductTypeType4 ProductType = 4
	// POS кредит (потребительский кредит, кредит на товар)
	ProductTypeType5 ProductType = 5
	// Кредит на развитие бизнеса
	ProductTypeType6 ProductType = 6
	// Кредит на пополнение оборотных средств
	ProductTypeType7 ProductType = 7
	// Кредит на покупку оборудования
	ProductTypeType8 ProductType = 8
	// Кредит на строительство недвижимости
	ProductTypeType9 ProductType = 9
	// Кредит на покупку акций (маржинальное кредитование)
	ProductTypeType10 ProductType = 10
	// Межбанковский кредит
	ProductTypeType11 ProductType = 11
	// Кредит мобильного оператора
	ProductTypeType12 ProductType = 12
	// Кредит на обучение
	ProductTypeType13 ProductType = 13
	// Дебетовая карта с овердрафтом
	ProductTypeType14 ProductType = 14
	// Ипотека (первичный рынок)
	ProductTypeType15 ProductType = 15
	// Ипотека (вторичный рынок)
	ProductTypeType16 ProductType = 16
	// Ипотека (ломбардный кредит)
	ProductTypeType17 ProductType = 17
	// Кредит наличными (нецелевой)
	ProductTypeType18 ProductType = 18
	// Микрозайм
	ProductTypeType19 ProductType = 19
	// Нецелевой кредит под залог автомобиля или недвижимости
	ProductTypeType20 ProductType = 20
	// Депозит
	ProductTypeType21 ProductType = 21
	// Другой тип кредита
	ProductTypeType99 ProductType = 99
)

type OriginalChannel uint32

const (
	// Отделение
	OriginalChannelType1 OriginalChannel = 1
	// Колл-центр
	OriginalChannelType2 OriginalChannel = 2
	// Брокер
	OriginalChannelType3 OriginalChannel = 3
	// Интернет
	OriginalChannelType4 OriginalChannel = 4
	// Кросс-селл
	OriginalChannelType5 OriginalChannel = 5
	// Точка продаж
	OriginalChannelType6 OriginalChannel = 6
	// Корпоративные продажи
	OriginalChannelType7 OriginalChannel = 7
	// Другое
	OriginalChannelType98 OriginalChannel = 98
	// Значение не передается
	OriginalChannelType99 OriginalChannel = 99
)

type SumCurrency string

func (p SumCurrency) ToString() string {
	return string(p)
}

const (
	// Американский доллар
	SumCurrencyType840 SumCurrency = "840"
	// Американский доллар
	SumCurrencyTypeUSD SumCurrency = "USD"
	// Рубль
	SumCurrencyType810 SumCurrency = "810"
	// Рубль
	SumCurrencyTypeRUR SumCurrency = "RUR"
	// Рубль
	SumCurrencyTypeRUB SumCurrency = "RUB"
	// Евро
	SumCurrencyType978 SumCurrency = "978"
	// Евро
	SumCurrencyTypeEUR SumCurrency = "EUR"
	// Швейцарский франк
	SumCurrencyType756 SumCurrency = "756"
	// Швейцарский франк
	SumCurrencyTypeCHF SumCurrency = "CHF"
	// Японская йена
	SumCurrencyType392 SumCurrency = "392"
	// Японская йена
	SumCurrencyTypeJPY SumCurrency = "JPY"
)

type CollateralExistence uint32

const (
	// Нет
	CollateralExistenceType0 CollateralExistence = 0
	// Да
	CollateralExistenceType1 CollateralExistence = 1
	// Значение не передает
	CollateralExistenceType99 CollateralExistence = 99
)

type PurchaseExistence uint32

const (
	// Нет
	PurchaseExistenceType0 PurchaseExistence = 0
	// Да
	PurchaseExistenceType1 PurchaseExistence = 1
	// Значение не передается
	PurchaseExistenceType99 PurchaseExistence = 99
)

type NewApplicant uint32

const (
	// Нет
	NewApplicantType0 NewApplicant = 0
	// Да
	NewApplicantType1 NewApplicant = 1
	// Не определен
	NewApplicantType9 NewApplicant = 9
	// Значение не передается
	NewApplicantType99 NewApplicant = 99
)

type ApplicantType uint32

const (
	// Заемщик
	ApplicantTypeType1 ApplicantType = 1
	// Созаемщик
	ApplicantTypeType2 ApplicantType = 2
	// Поручитель
	ApplicantTypeType3 ApplicantType = 3
)

type ResponseIsNeeded uint32

const (
	// выходной вектор рассчитывается по кредитным заявкам РБД в зависимости от наличия фотографий по ним – с учетом фотографий Аппликантов при их наличии или без учета фотографий Аппликантов при их отсутствии;
	ResponseIsNeededType1 ResponseIsNeeded = 1
	// выходной вектор рассчитывается по кредитным заявкам РБД только с учетом фотографий Аппликантов;
	ResponseIsNeededType3 ResponseIsNeeded = 3
	// выходной вектор не рассчитывается в ответ на запрос. В данном случае не происходит расчета Предикторов, Правил и Баллов подозрительности и Выходной вектор не предоставляется.
	ResponseIsNeededType0 ResponseIsNeeded = 0
)

type ApplicationStatus uint32

const (
	// Одобрена
	ApplicationStatusType1 ApplicationStatus = 1
	// Отказана
	ApplicationStatusType2 ApplicationStatus = 2
	// Одобрена, но кредит не выдан
	ApplicationStatusType3 ApplicationStatus = 3
	// Одобрена. Кредит выдан
	ApplicationStatusType4 ApplicationStatus = 4
	// Аннулирована
	ApplicationStatusType5 ApplicationStatus = 5
	// Закрыт автоматически (Проставляется в системе автоматически. Для внутреннего использования. Партнер не должен выгружать данное значение)
	ApplicationStatusType8 ApplicationStatus = 8
	// Не определен
	ApplicationStatusType9 ApplicationStatus = 9
)

type ApplicationFraudStatus uint32

const (
	// Мошенничество
	ApplicationFraudStatusType1 ApplicationFraudStatus = 1
	// Фрод
	ApplicationFraudStatusType2 ApplicationFraudStatus = 2
	// Нет признаков фрода
	ApplicationFraudStatusType3 ApplicationFraudStatus = 3
	// Подозрение в мошенничестве
	ApplicationFraudStatusType4 ApplicationFraudStatus = 4
	// Закрыт автоматически (Проставляется в системе автоматически. Для внутреннего использования. Партнер не должен выгружать данное значение)
	ApplicationFraudStatusType8 ApplicationFraudStatus = 8
	// Не определен
	ApplicationFraudStatusType9 ApplicationFraudStatus = 9
)

type DefaultStatus uint32

const (
	// Дефолт
	DefaultStatusType1 DefaultStatus = 1
	// Технический дефолт
	DefaultStatusType2 DefaultStatus = 2
	// Нет дефолта
	DefaultStatusType3 DefaultStatus = 3
	// Закрыт автоматически (Проставляется в системе автоматически. Для внутреннего использования. Партнер не должен выгружать данное значение)
	DefaultStatusType8 DefaultStatus = 8
	// Не определен
	DefaultStatusType9 DefaultStatus = 9
)

type Status uint32

const (
	// Без ошибок
	StatusType0 Status = 0
	// Заявка отправлена в карантин
	StatusType1 Status = 1
	// Не заданы обязательные поля
	StatusType2 Status = 2
	// Не найдена заявка с переданным Application ID
	StatusType3 Status = 3
	// Зарезервирован
	StatusType4 Status = 4
	// Данные не соответствуют формату
	StatusType5 Status = 5
	// Ошибка сервиса нормализации
	StatusType6 Status = 6
	// Не задан статус кредитной заявки
	StatusType7 Status = 7
	// Не задан фрод-статус кредитной заявки
	StatusType8 Status = 8
	// Не задан дефолт-статус кредитной заявки
	StatusType9 Status = 9
	// Выходной вектор еще не рассчитан
	StatusType10 Status = 10
	// Флаг предоставления выходного вектора из Системы не передавался. Выходной вектор не рассчитывался.
	StatusType11 Status = 11
	// Кредитная заявка с данным ID обрабатывается другим пользователем
	StatusType12 Status = 12
	// Кредитная заявка для данного пользователя не доступна
	StatusType13 Status = 13
	// Кредитная заявка уже была отправлена на обработку ранее
	StatusType14 Status = 14
	// Кредитная заявка с данным ID уже есть в Системе
	StatusType15 Status = 15
	// Зарезервирован
	StatusType20 Status = 20
	// Зарезервирован
	StatusType21 Status = 21
	// Зарезервирован
	StatusType22 Status = 22
	// Невозможно установить данный статус кредитной заявки. Данный статус заявки уже был установлен ранее
	StatusType23 Status = 23
	// Невозможно установить данный фрод-статус кредитной заявки. Данный фрод-статус заявки уже был установлен ранее
	StatusType24 Status = 24
	// Невозможно установить данный дефолт-статус кредитной заявки. Данный дефолт-статус заявки уже был установлен ранее
	StatusType25 Status = 25
	// Неверно задан статус заявки. Статус «Закрыт автоматически» устанавливается в системе автоматически и только для внутреннего использования
	StatusType26 Status = 26
	// Неверно задан фрод-статус заявки. Фрод-статус «Закрыт автоматически» устанавливается в системе автоматически и только для внутреннего использования
	StatusType27 Status = 27
	// Неверно задан дефолт-статус заявки. Дефолт-статус «Закрыт автоматически» устанавливается в системе автоматически и только для внутреннего использования
	StatusType28 Status = 28
	// Ошибка при обращении к ПО Оператора
	StatusType30 Status = 30
	// Не задан признак запроса к базе данных Участников
	StatusType31 Status = 31
	// При сравнении контрольных сумм фотографии произошла ошибка. Принятая от Партнера контрольная сумма фотографии (checksumphoto) не совпадает с контрольной суммой, вычисленной Оператором от принятого файла фотографии (photo)
	StatusType39 Status = 39
	// При загрузке фотографии произошла ошибка. Файл фотографии не читается
	StatusType40 Status = 40
	// Фотография не пригодна к обработке. На фотографии отсутствуют глаза
	StatusType41 Status = 41
	// Фотография не пригодна к обработке. Слишком низкое качество фотографии
	StatusType42 Status = 42
	// Фотография не пригодна к обработке. Лицо на фотографии обрезано
	StatusType43 Status = 43
	// Файл с фотографией не загружен. Размер файла фотографии превышает допустимый предел
	StatusType44 Status = 44
	// Для данной кредитной заявки уже имеется фотография Аппликанта. Повторный прием фотографии невозможен
	StatusType45 Status = 45
	// Предупреждение! На фотографии присутствует несколько лиц. Автоматически Системой было выбрано наиболее крупное! Загрузка фотографии прошла успешно
	StatusType46 Status = 46
	// Фотография с данным уникальным идентификатором уже имеется в Системе
	StatusType49 Status = 49
	// Фотография с данным уникальным идентификатором не найдена в Системе
	StatusType50 Status = 50
	// Предупреждение! Фотография Аппликанта не найдена для данной кредитной заявки. При расчете выходного вектора не были использованы биометрические Правила
	StatusType47 Status = 47
	// Фотография Аппликанта не найдена для данной кредитной заявки. Выходной вектор не рассчитан
	StatusType48 Status = 48
	// Партнер заблокирован в системе. Загрузка файла {file} невозможна
	StatusType60 Status = 60
	// Файл {file} не принимается. Неправильное имя файла
	StatusType61 Status = 61
	// Подпись и/или шифрование выполнены не корректно
	StatusType62 Status = 62
	// Файл {file} подписан неизвестным сертификатом
	StatusType63 Status = 63
	// Файл {file} не принимается, т.к. архив не распаковывается или является пустым
	StatusType64 Status = 64
	// Файл {file} не принимается, т.к. в архиве содержатся посторонние файлы
	StatusType65 Status = 65
	// Передано некорректное количество полей
	StatusType66 Status = 66
	// Архив с файлами фотографий {photoarchive} не найден
	StatusType67 Status = 67
	// Файл с фотографией {photofile} не найден в архиве {photoarchive}
	StatusType68 Status = 68
	// Ошибка справочника Oracle
	StatusType90 Status = 90
	// Нет соединения с ЛБД
	StatusType98 Status = 98
	// Другая ошибка. См. таблицу логов
	StatusType99 Status = 99
)
