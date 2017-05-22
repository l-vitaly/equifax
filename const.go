package equifax

// цель согласия
type Reason uint32

const (
    ReasonType0 Reason = 0 // заключение и исполнение договора
    ReasonType1        = 1 // проверка благонадежности
    ReasonType2        = 2 // прием на работу
    ReasonType9        = 9 // иная цель согласия
)

// цель финансирования
type Purpose string

const (
    PurposeType01 Purpose = "01" // новый автомобиль
    PurposeType02         = "02" // подержанный автомобиль
    PurposeType03         = "03" // другое транспортное средство
    PurposeType04         = "04" // мебель
    PurposeType05         = "05" // ремонт дома (в том числе электрификация, газификация, топливное обеспечение)
    PurposeType06         = "06" // бытовая техника
    PurposeType07         = "07" // одежда
    PurposeType08         = "08" // путешествия
    PurposeType09         = "09" // земля
    PurposeType10         = "10" // дом
    PurposeType11         = "11" // возврат долга
    PurposeType12         = "12" // свадьба
    PurposeType13         = "13" // образование
    PurposeType14         = "14" // компьютерная техника
    PurposeType15         = "15" // услуги
    PurposeType16         = "16" // кооперативные платежи, рента, депозит
    PurposeType17         = "17" // инвестиции (Ценные бумаги, облигации…)
    PurposeType18         = "18" // здоровье / Затраты на лечение
    PurposeType19         = "19" // хобби
    PurposeType20         = "20" // коммерческие, деловые (Фонды)
    PurposeType21         = "21" // телекоммуникационное оборудование (не мобильное)
    PurposeType22         = "22" // мобильный телефон
    PurposeType23         = "23" // оборотные средства
    PurposeType24         = "24" // вложения в основной капитал
    PurposeType25         = "25" // сельскохозяйственный заём
    PurposeType98         = "98" // другое
    PurposeType99         = "99" // неизвестно
)

type Cred string

const (
    CredType00 Cred = "00" // неизвестный тип кредита
    CredType01      = "01" // кредит на автомобиль
    CredType02      = "02" // лизинг
    CredType03      = "03" // ипотека
    CredType04      = "04" // кредитная карта
    CredType05      = "05" // потребительский кредит
    CredType06      = "06" // кредит на развитие бизнеса
    CredType07      = "07" // кредит на пополнение оборотных средств
    CredType08      = "08" // кредит на покупку оборудования
    CredType09      = "09" // кредит на строительство недвижимости
    CredType10      = "10" // кредит на покупку акций (маржинальное кредитование)
    CredType11      = "11" // межбанковский кредит
    CredType12      = "12" // кредит мобильного оператора
    CredType13      = "13" // кредит на обучение
    CredType14      = "14" // дебетовая карта с овердрафтом
    CredType15      = "15" // ипотека (первичный рынок)
    CredType16      = "16" // ипотека (вторичный рынок)
    CredType17      = "17" // ипотека (ломбардный кредит)
    CredType18      = "18" // кредит наличными (нецелевой)
    CredType19      = "19" // микрозайм
    CredType90      = "90" // договор поручительства
    CredType99      = "99" // другой тип кредита
)

// тип обеспечения
type CredSecurity uint32

const (
    CredSecurityType0 CredSecurity = 0 // недвижимость
    CredSecurityType1              = 1 // валюта или ценные бумаги
    CredSecurityType2              = 2 // залог товаров
    CredSecurityType3              = 3 // частный гарант или корпоративный гарант
    CredSecurityType4              = 4 // автомобиль
    CredSecurityType8              = 8 // другое
    CredSecurityType9              = 9 // нет
)

// период получения дохода
type IncomeFrequency uint32

const (
    IncomeFrequencyType0 IncomeFrequency = 0
    IncomeFrequencyType1                 = 1
    IncomeFrequencyType2                 = 2
    IncomeFrequencyType3                 = 3
    IncomeFrequencyType4                 = 4
    IncomeFrequencyType5                 = 5
)

// тип согласия для запроса
type Consent uint32

const (
    ConsentType0 Consent = 0 // согласие не дано
    ConsentType1         = 1 // согласие дано
)

// пол
type Gender uint32

const (
    GenderType1 Gender = 1 // мужской
    GenderType2        = 2 // женский
    GenderType9        = 9 // неизвестно
)

// допустимые значений полей responsecode и responsestring
type ResponseCode int32

const (
    ResponseCodeType0  ResponseCode = 0  // без ошибок
    ResponseCodeType1               = 1  // заёмщик найден
    ResponseCodeType3               = 3  // заёмщик с такими данными не найден
    ResponseCodeType4               = 4  // указанного типа отчета не существует
    ResponseCodeType5               = 5  // нет такого Партнера
    ResponseCodeType11              = 11 // подпись запроса не соответствует Партнеру
    ResponseCodeType12              = 12 // структура XML запроса не корректна
    ResponseCodeType15              = 15 // неверная версия XML-запроса
    ResponseCodeType19              = 19 // запрос не подписан
    ResponseCodeType24              = 24 // на запрашиваемую дату кредитной истории не существовало
    ResponseCodeType30              = 30 // не дано согласие CКИ на получение его КО (consent = 0) и/или Партнер не проинформирован об ответственности по ст.5.53 и 14.29 КоАП РФ (admcode_inform = 0)
    ResponseCodeType31              = 31 // некорректно указана дата выдачи согласия СКИ на получение его КО
    ResponseCodeType32              = 32 // отсутствует блок информации заявления (блок application)
    ResponseCodeType33              = 33 // в запросе некорректно указано поле reason {reason} и/или идентификатор отчета {type}
    ResponseCodeType34              = 34 // в запросе по юр. лицу – резиденту отсутствуют (отсутствует) ИНН и/или ОГРН
    ResponseCodeType36              = 36 // не указан СНИЛС (pfno)
    ResponseCodeType37              = 37 // не указана иная цель согласия
    ResponseCodeType99              = 99 // сервис недоступен
)

// текущее / предыдущее место работы
type EmploymentCurrent uint32

const (
    EmploymentCurrentType0 EmploymentCurrent = 0 // предыдущее
    EmploymentCurrentType1                   = 1 // текущее
)

type EmploymentType uint32

const (
    EmploymentTypeType0 EmploymentType = 0
    EmploymentTypeType1                = 1
    EmploymentTypeType2                = 2
    EmploymentTypeType3                = 3
    EmploymentTypeType4                = 4
    EmploymentTypeType5                = 5
    EmploymentTypeType6                = 6
    EmploymentTypeType7                = 7
    EmploymentTypeType9                = 9
)

// профессия
type Profession string

const (
    ProfessionType00 Profession = "00" // топ-менеджер
    ProfessionType01            = "01" // государственный служащий
    ProfessionType02            = "02" // собственник бизнеса
    ProfessionType03            = "03" // высококвалифицированный персонал
    ProfessionType04            = "04" // офисный служащий
    ProfessionType05            = "05" // специалист
    ProfessionType06            = "06" // работник сферы обслуживания
    ProfessionType07            = "07" // работник с/х
    ProfessionType08            = "08" // рабочий
    ProfessionType09            = "09" // неквалифицированный рабочий
    ProfessionType10            = "10" // военнослужащий
    ProfessionType98            = "98" // другое
    ProfessionType99            = "99" // неизвестно
)

// государственность предприятия
type CompanyState uint32

const (
    CompanyStateType0 CompanyState = 0 // коммерческое предприятие
    CompanyStateType1              = 1 // государственное предприятие
    CompanyStateType9              = 9 // неизвестно
)

// размер предприятия
type CompanySize uint32

const (
    CompanySizeType0 CompanySize = 0 // < 50 человек
    CompanySizeType1             = 1 // 50-100 человек
    CompanySizeType2             = 2 // 101-249 человек
    CompanySizeType3             = 3 // 250-499 человек
    CompanySizeType4             = 4 // >500 человек
    CompanySizeType9             = 9 // неизвестно
)

// вид деятельности предприятия
type CompanyArea string

const (
    CompanyAreaType00 CompanyArea = "00" // промышленность и машиностроение
    CompanyAreaType01             = "01" // сельское, хозяйство
    CompanyAreaType02             = "02" // строительство
    CompanyAreaType03             = "03" // горное дело
    CompanyAreaType04             = "04" // энергетика
    CompanyAreaType05             = "05" // оптовая торговля
    CompanyAreaType06             = "06" // финансовое дело и страхование
    CompanyAreaType07             = "07" // здравоохранение
    CompanyAreaType08             = "08" // социальная помощь
    CompanyAreaType09             = "09" // денежные переводы и обмен валют
    CompanyAreaType10             = "10" // искусство, развлечение, отдых
    CompanyAreaType11             = "11" // казино и игорный бизнес
    CompanyAreaType12             = "12" // торговля предметами искусства и антиквариатом
    CompanyAreaType13             = "13" // авиа-, авто- и железнодорожные перевозки, складское хранение
    CompanyAreaType14             = "14" // административно-хозяйственные службы
    CompanyAreaType15             = "15" // розничная торговля
    CompanyAreaType16             = "16" // водо-, тепло- и энергоснабжение
    CompanyAreaType17             = "17" // управление компанией
    CompanyAreaType18             = "18" // предпринимательская деятельность
    CompanyAreaType19             = "19" // научная, техническая и профессиональная деятельность
    CompanyAreaType20             = "20" // образование
    CompanyAreaType21             = "21" // торговля ювелирными украшениями и драгоценными металлами
    CompanyAreaType22             = "22" // посреднические услуги по продаже и аренде недвижимости
    CompanyAreaType23             = "23" // снабжение
    CompanyAreaType24             = "24" // пресса, телевидение и радио
    CompanyAreaType25             = "25" // государственное управление
    CompanyAreaType98             = "98" // другое
    CompanyAreaType99             = "99" // неизвестно
)

type AddressOwner uint32

const (
    AddressOwnerType0 AddressOwner = 0 // семейное владение
    AddressOwnerType1              = 1 // собственность
    AddressOwnerType2              = 2 // ипотека
    AddressOwnerType3              = 3 // аренда
    AddressOwnerType4              = 4 // живет с родителями
    AddressOwnerType5              = 5 // живет с кем-либо
    AddressOwnerType6              = 6 // жилье обеспечивается работодателем
    AddressOwnerType7              = 7 // общежитие / воинская часть
    AddressOwnerType8              = 8 // муниципальное жилье
    AddressOwnerType9              = 9 // неизвестно
)

// регион
type Region string

const (
    RegionType01 Region = "01" // алтайский край
    RegionType03        = "03" // краснодарский край
    RegionType04        = "04" // красноярский край, таймырский, эвенкийский район
    RegionType05        = "05" // приморский край
    RegionType07        = "07" // ставропольский край
    RegionType08        = "08" // хабаровский край
    RegionType10        = "10" // амурская область
    RegionType11        = "11" // архангельская область, ненецкий ао
    RegionType12        = "12" // астраханская область
    RegionType14        = "14" // белгородская область
    RegionType15        = "15" // брянская область
    RegionType17        = "17" // владимирская область
    RegionType18        = "18" // волгоградская область
    RegionType19        = "19" // вологодская область
    RegionType20        = "20" // воронежская область
    RegionType22        = "22" // нижегородская область
    RegionType24        = "24" // ивановская область
    RegionType25        = "25" // иркутская область, усть-ордынский бурятский округ
    RegionType26        = "26" // республика ингушетия
    RegionType27        = "27" // калининградская область
    RegionType28        = "28" // тверская область
    RegionType29        = "29" // калужская область
    RegionType30        = "30" // камчатский край, корякский округ
    RegionType32        = "32" // кемеровская область
    RegionType33        = "33" // кировская область
    RegionType34        = "34" // костромская область
    RegionType36        = "36" // самарская область
    RegionType35        = "35" // республика крым
    RegionType37        = "37" // курганская область
    RegionType38        = "38" // курская область
    RegionType40        = "40" // санкт-петербург
    RegionType41        = "41" // ленинградская область
    RegionType42        = "42" // липецкая область
    RegionType44        = "44" // магаданская область
    RegionType45        = "45" // москва
    RegionType46        = "46" // московская область
    RegionType47        = "47" // мурманская область
    RegionType49        = "49" // новгородская область
    RegionType50        = "50" // новосибирская область
    RegionType52        = "52" // омская область
    RegionType53        = "53" // оренбургская область
    RegionType54        = "54" // орловская область
    RegionType55        = "55" // байконур
    RegionType56        = "56" // пензенская область
    RegionType57        = "57" // пермский край, коми-пермяцкий округ
    RegionType58        = "58" // псковская область
    RegionType60        = "60" // ростовская область
    RegionType61        = "61" // рязанская область
    RegionType63        = "63" // саратовская область
    RegionType64        = "64" // сахалинская область
    RegionType65        = "65" // свердловская область
    RegionType66        = "66" // смоленская область
    RegionType67        = "67" // севастополь
    RegionType68        = "68" // тамбовская область
    RegionType69        = "69" // томская область
    RegionType70        = "70" // тульская область
    RegionType71        = "71" // тюменская область, ханты-мансийский ао – югра, ямало-ненецкий ао
    RegionType73        = "73" // ульяновская область
    RegionType75        = "75" // челябинская область
    RegionType76        = "76" // забайкальский край, агинский бурятский округ
    RegionType77        = "77" // чукотский автономный округ
    RegionType78        = "78" // ярославская область
    RegionType79        = "79" // республика адыгея (адыгея)
    RegionType80        = "80" // республика башкортостан
    RegionType81        = "81" // республика бурятия
    RegionType82        = "82" // республика дагестан
    RegionType83        = "83" // кабардино-балкарская республика
    RegionType84        = "84" // республика алтай
    RegionType85        = "85" // республика калмыкия
    RegionType86        = "86" // республика карелия
    RegionType87        = "87" // республика коми
    RegionType88        = "88" // республика марий эл
    RegionType89        = "89" // республика мордовия
    RegionType90        = "90" // республика северная осетия-алания
    RegionType91        = "91" // карачаево-черкесская республика
    RegionType92        = "92" // республика татарстан (татарстан)
    RegionType93        = "93" // республика тыва
    RegionType94        = "94" // удмуртская республика
    RegionType95        = "95" // республика хакасия
    RegionType96        = "96" // чеченская республика
    RegionType97        = "97" // чувашская республика - чувашия
    RegionType98        = "98" // республика саха (якутия)
    RegionType99        = "99" // еврейская автономная область
    RegionType00        = "00" // неизвестно
)

type AdmCodeInForm uint32

const (
    AdmCodeInFormType0 AdmCodeInForm = 0 // пользователь КИ не проинформирован об административной ответственности
    AdmCodeInFormType1               = 1 // пользователь КИ проинформирован об административной ответственности
)

type Resident uint32

const (
    ResidentType0 Resident = 0 // нерезидент
    ResidentType1          = 1 // резидент
)

// документ удостоверяющий личность
type DocType uint32

const (
    DocType1  DocType = 1  // паспорт гражданина Российской Федерации – для гражданина Российской Федерации, достигшего 14 лет
    DocType2          = 2  // свидетельство органов ЗАГСа, органа исполнительной власти или органа местного самоуправления о рождении гражданина – для гражданина Российской Федерации, не достигшего 14 лет
    DocType3          = 3  // удостоверение личности – для офицеров, прапорщиков и мичманов
    DocType4          = 4  // военный билет – для сержантов, старшин, солдат и матросов, а также курсантов военных образовательных учреждений профессионального образования
    DocType5          = 5  // паспорт моряка – для граждан Российской Федерации, работающих на судах заграничного плавания или на иностранных судах, курсантов учебных заведений
    DocType6          = 6  // паспорт иностранного гражданина либо иной документ, установленный федеральным законом или признаваемый в соответствии с международным договором Российской Федерации в качестве документа, удостоверяющего личность иностранного гражданина
    DocType7          = 7  // документ, выданный иностранным государством и признаваемый в соответствии с международным договором Российской Федерации в качестве документа, удостоверяющего личность лица без гражданства
    DocType8          = 8  // разрешение на временное проживание лица без гражданства
    DocType9          = 9  // вид на жительство лица без гражданства
    DocType10         = 10 // иные документы, предусмотренные федеральным законом или признаваемые в соответствии с международным договором Российской Федерации в качестве документов, удостоверяющих личность лица без гражданства
    DocType11         = 11 // свидетельство о регистрации ходатайства о признании иммигранта беженцем
    DocType12         = 12 // удостоверение беженца
    DocType13         = 13 // временное удостоверение личности гражданина
    DocType14         = 14 // иные документы, выдаваемые уполномоченными органами
    DocType15         = 15 // паспорт гражданина СССР
    DocType98         = 98 // нет (Значение может передаваться только для ранее выданного документа)
    DocType99         = 99 // значение не передается (Значение может передаваться только для ранее выданного документа)
)

type Sex uint32

const (
    SexType1  Sex = 1  // мужской
    SexType2      = 2  // женский
    SexType99     = 99 // значение не передается
)

type Country string

func (c Country) ToString() string {
    return string(c)
}

const (
    CountryTypeAU Country = "AU" // Австралия
    CountryTypeAT         = "AT" // Австрия
    CountryTypeAZ         = "AZ" // Азербайджан
    CountryTypeAX         = "AX" // Аландские острова
    CountryTypeAL         = "AL" // Албания
    CountryTypeDZ         = "DZ" // Алжир
    CountryTypeAS         = "AS" // Американское Самоа
    CountryTypeAI         = "AI" // Ангилья
    CountryTypeAO         = "AO" // Ангола
    CountryTypeAD         = "AD" // Андорра
    CountryTypeAG         = "AG" // Антигуа и Барбуда
    CountryTypeAR         = "AR" // Аргентина
    CountryTypeAM         = "AM" // Армения
    CountryTypeAW         = "AW" // Аруба
    CountryTypeAF         = "AF" // Афганистан
    CountryTypeBS         = "BS" // Багамы
    CountryTypeBD         = "BD" // Бангладеш
    CountryTypeBB         = "BB" // Барбадос
    CountryTypeBH         = "BH" // Бахрейн
    CountryTypeBY         = "BY" // Беларусь
    CountryTypeBZ         = "BZ" // Белиз
    CountryTypeBE         = "BE" // Бельгия
    CountryTypeBJ         = "BJ" // Бенин
    CountryTypeBM         = "BM" // Бермуды
    CountryTypeBG         = "BG" // Болгария
    CountryTypeBO         = "BO" // Боливия
    CountryTypeBA         = "BA" // Босния и Герцеговина
    CountryTypeBW         = "BW" // Ботсвана
    CountryTypeBR         = "BR" // Бразилия
    CountryTypeIO         = "IO" // Британская территория в Индийском океане
    CountryTypeBN         = "BN" // Бруней-Даруссалам
    CountryTypeBV         = "BV" // Буве Остров
    CountryTypeBF         = "BF" // Буркина Фасо
    CountryTypeBI         = "BI" // Бурунди
    CountryTypeBT         = "BT" // Бутан
    CountryTypeVU         = "VU" // Вануату
    CountryTypeVA         = "VA" // Ватикан
    CountryTypeGB         = "GB" // Великобритания
    CountryTypeHU         = "HU" // Венгрия
    CountryTypeVE         = "VE" // Венесуэла
    CountryTypeVG         = "VG" // Виргинские острова Британские
    CountryTypeVI         = "VI" // Виргинские острова США
    CountryTypeUM         = "UM" // Внешние малые острова США
    CountryTypeTL         = "TL" // Восточный Тимор
    CountryTypeVN         = "VN" // Вьетнам
    CountryTypeGA         = "GA" // Габон
    CountryTypeGY         = "GY" // Гайана
    CountryTypeHT         = "HT" // Гаити
    CountryTypeGM         = "GM" // Гамбия
    CountryTypeGH         = "GH" // Гана
    CountryTypeGP         = "GP" // Гваделупа
    CountryTypeGT         = "GT" // Гватемала
    CountryTypeGF         = "GF" // Гвиана
    CountryTypeGN         = "GN" // Гвинея
    CountryTypeGW         = "GW" // Гвинея-Бисау
    CountryTypeDE         = "DE" // Германия
    CountryTypeGG         = "GG" // Гернси
    CountryTypeGI         = "GI" // Гибралтар
    CountryTypeHN         = "HN" // Гондурас
    CountryTypeHK         = "HK" // Гонконг
    CountryTypeGD         = "GD" // Гренада
    CountryTypeGL         = "GL" // Гренландия
    CountryTypeGR         = "GR" // Греция
    CountryTypeGE         = "GE" // Грузия
    CountryTypeGU         = "GU" // Гуам
    CountryTypeDK         = "DK" // Дания
    CountryTypeJE         = "JE" // Джерси
    CountryTypeDJ         = "DJ" // Джибути
    CountryTypeDM         = "DM" // Доминика
    CountryTypeDO         = "DO" // Доминиканская Республика
    CountryTypeEG         = "EG" // Египет
    CountryTypeZM         = "ZM" // Замбия
    CountryTypeEH         = "EH" // Западная Сахара
    CountryTypeZW         = "ZW" // Зимбабве
    CountryTypeYE         = "YE" // Йемен
    CountryTypeIL         = "IL" // Израиль
    CountryTypeIN         = "IN" // Индия
    CountryTypeID         = "ID" // Индонезия
    CountryTypeJO         = "JO" // Иордания
    CountryTypeIQ         = "IQ" // Ирак
    CountryTypeIR         = "IR" // Иран
    CountryTypeIE         = "IE" // Ирландия
    CountryTypeIS         = "IS" // Исландия
    CountryTypeES         = "ES" // Испания
    CountryTypeIT         = "IT" // Италия
    CountryTypeCV         = "CV" // Кабо-Верде
    CountryTypeKZ         = "KZ" // Казахстан
    CountryTypeKY         = "KY" // Каймановы острова
    CountryTypeKH         = "KH" // Камбоджа
    CountryTypeCM         = "CM" // Камерун
    CountryTypeCA         = "CA" // Канада
    CountryTypeQA         = "QA" // Катар
    CountryTypeKE         = "KE" // Кения
    CountryTypeCY         = "CY" // Кипр
    CountryTypeKG         = "KG" // Киргизия
    CountryTypeKI         = "KI" // Кирибати
    CountryTypeCN         = "CN" // Китай
    CountryTypeKP         = "KP" // КНДР
    CountryTypeCC         = "CC" // Кокосовые (Килинг) острова
    CountryTypeCO         = "CO" // Колумбия
    CountryTypeKM         = "KM" // Коморы
    CountryTypeCD         = "CD" // Конго (Демократическая Республика Конго)
    CountryTypeCG         = "CG" // Конго (Республика Конго)
    CountryTypeCR         = "CR" // Коста-Рика
    CountryTypeCI         = "CI" // Кот-д'Ивуар
    CountryTypeCU         = "CU" // Куба
    CountryTypeKW         = "KW" // Кувейт
    CountryTypeCK         = "CK" // Кука Острова
    CountryTypeLA         = "LA" // Лаос
    CountryTypeLV         = "LV" // Латвия
    CountryTypeLS         = "LS" // Лесото
    CountryTypeLR         = "LR" // Либерия
    CountryTypeLB         = "LB" // Ливан
    CountryTypeLY         = "LY" // Ливия
    CountryTypeLT         = "LT" // Литва
    CountryTypeLI         = "LI" // Лихтенштейн
    CountryTypeLU         = "LU" // Люксембург
    CountryTypeMU         = "MU" // Маврикий
    CountryTypeMR         = "MR" // Мавритания
    CountryTypeMG         = "MG" // Мадагаскар
    CountryTypeYT         = "YT" // Майотта
    CountryTypeMO         = "MO" // Макао
    CountryTypeMW         = "MW" // Малави
    CountryTypeMY         = "MY" // Малайзия
    CountryTypeML         = "ML" // Мали
    CountryTypeMV         = "MV" // Мальдивы
    CountryTypeMT         = "MT" // Мальта
    CountryTypeMA         = "MA" // Марокко
    CountryTypeMQ         = "MQ" // Мартиника
    CountryTypeMH         = "MH" // Маршалловы острова
    CountryTypeMX         = "MX" // Мексика
    CountryTypeMZ         = "MZ" // Мозамбик
    CountryTypeMD         = "MD" // Молдавия
    CountryTypeMC         = "MC" // Монако
    CountryTypeMN         = "MN" // Монголия
    CountryTypeMS         = "MS" // Монтсеррат
    CountryTypeMM         = "MM" // Мьянма
    CountryTypeNA         = "NA" // Намибия
    CountryTypeNR         = "NR" // Науру
    CountryTypeNP         = "NP" // Непал
    CountryTypeNE         = "NE" // Нигер
    CountryTypeNG         = "NG" // Нигерия
    CountryTypeAN         = "AN" // Нидерландские Антильские о-ва
    CountryTypeNL         = "NL" // Нидерланды
    CountryTypeNI         = "NI" // Никарагуа
    CountryTypeNU         = "NU" // Ниуэ
    CountryTypeNZ         = "NZ" // Новая Зеландия
    CountryTypeNC         = "NC" // Новая Каледония
    CountryTypeNO         = "NO" // Норвегия
    CountryTypeNF         = "NF" // Норфолк Остров
    CountryTypeAE         = "AE" // Объединенные Арабские Эмираты
    CountryTypeOM         = "OM" // Оман
    CountryTypeIM         = "IM" // Остров Мэн
    CountryTypeHM         = "HM" // Остров Херд и острова Макдональд
    CountryTypePK         = "PK" // Пакистан
    CountryTypePW         = "PW" // Палау
    CountryTypePS         = "PS" // Палестина
    CountryTypePA         = "PA" // Панама
    CountryTypePG         = "PG" // Папуа - Новая Гвинея
    CountryTypePY         = "PY" // Парагвай
    CountryTypePE         = "PE" // Перу
    CountryTypePN         = "PN" // Питкэрн Острова
    CountryTypePL         = "PL" // Польша
    CountryTypePT         = "PT" // Португалия
    CountryTypePR         = "PR" // Пуэрто-Рико
    CountryTypeKR         = "KR" // Республика Корея
    CountryTypeMK         = "MK" // Республика Македония
    CountryTypeRE         = "RE" // Реюньон Остров
    CountryTypeCX         = "CX" // Рождества Остров
    CountryTypeRU         = "RU" // Россия
    CountryTypeRW         = "RW" // Руанда
    CountryTypeRO         = "RO" // Румыния
    CountryTypeSV         = "SV" // Сальвадор
    CountryTypeWS         = "WS" // Самоа
    CountryTypeSM         = "SM" // Сан-Марино
    CountryTypeST         = "ST" // Сан-Томе и Принсипи
    CountryTypeSA         = "SA" // Саудовская Аравия
    CountryTypeSZ         = "SZ" // Свазиленд
    CountryTypeMF         = "MF" // Святого Мартина Остров
    CountryTypeSH         = "SH" // Святой Елены Остров
    CountryTypeMP         = "MP" // Северные Марианские острова
    CountryTypeSC         = "SC" // Сейшелы
    CountryTypeBL         = "BL" // Сен-Бартельми
    CountryTypeSN         = "SN" // Сенегал
    CountryTypePM         = "PM" // Сен-Пьер и Микелон
    CountryTypeVC         = "VC" // Сент-Винсент и Гренадины
    CountryTypeKN         = "KN" // Сент-Китс и Невис
    CountryTypeLC         = "LC" // Сент-Люсия
    CountryTypeRS         = "RS" // Сербия
    CountryTypeSG         = "SG" // Сингапур
    CountryTypeSY         = "SY" // Сирийская Арабская Республика
    CountryTypeSK         = "SK" // Словакия
    CountryTypeSI         = "SI" // Словения
    CountryTypeSB         = "SB" // Соломоновы острова
    CountryTypeSO         = "SO" // Сомали
    CountryTypeSD         = "SD" // Судан
    CountryTypeSR         = "SR" // Суринам
    CountryTypeUS         = "US" // США
    CountryTypeSL         = "SL" // Сьерра-Леоне
    CountryTypeTJ         = "TJ" // Таджикистан
    CountryTypeTW         = "TW" // Тайвань (Китай)
    CountryTypeTH         = "TH" // Таиланд
    CountryTypeTZ         = "TZ" // Танзания
    CountryTypeTC         = "TC" // Теркс и Кайкос
    CountryTypeTG         = "TG" // Того
    CountryTypeTK         = "TK" // Токелау
    CountryTypeTO         = "TO" // Тонга
    CountryTypeTT         = "TT" // Тринидад и Тобаго
    CountryTypeTV         = "TV" // Тувалу
    CountryTypeTN         = "TN" // Тунис
    CountryTypeTM         = "TM" // Туркмения
    CountryTypeTR         = "TR" // Турция
    CountryTypeUG         = "UG" // Уганда
    CountryTypeUZ         = "UZ" // Узбекистан
    CountryTypeUA         = "UA" // Украина
    CountryTypeWF         = "WF" // Уоллис и Футуна
    CountryTypeUY         = "UY" // Уругвай
    CountryTypeFO         = "FO" // Фарерские острова
    CountryTypeFM         = "FM" // Федеративные Штаты Микронезии
    CountryTypeFJ         = "FJ" // Фиджи
    CountryTypePH         = "PH" // Филиппины
    CountryTypeFI         = "FI" // Финляндия
    CountryTypeFK         = "FK" // Фолклендские острова (Мальвинские)
    CountryTypeFR         = "FR" // Франция
    CountryTypePF         = "PF" // Французская Полинезия
    CountryTypeTF         = "TF" // Французские Южные и Антарктические территории
    CountryTypeHR         = "HR" // Хорватия
    CountryTypeCF         = "CF" // Центральноафриканская Республика
    CountryTypeTD         = "TD" // Чад
    CountryTypeME         = "ME" // Черногория
    CountryTypeCZ         = "CZ" // Чехия
    CountryTypeCL         = "CL" // Чили
    CountryTypeCH         = "CH" // Швейцария
    CountryTypeSE         = "SE" // Швеция
    CountryTypeLK         = "LK" // Шри-Ланка
    CountryTypeEC         = "EC" // Эквадор
    CountryTypeGQ         = "GQ" // Экваториальная Гвинея
    CountryTypeER         = "ER" // Эритрея
    CountryTypeEE         = "EE" // Эстония
    CountryTypeET         = "ET" // Эфиопия
    CountryTypeGS         = "GS" // Южная Георгия и Южные Сандвичевы острова
    CountryTypeZA         = "ZA" // Южно-Африканская Республика
    CountryTypeJM         = "JM" // Ямайка
    CountryTypeJP         = "JP" // Япония
    CountryType98         = "98" // Другая
    CountryType99         = "99" // Неизвестно
)

// образование
type Education uint32

const (
    EducationType0  Education = 0  // начальная школа
    EducationType1            = 1  // средняя школа
    EducationType2            = 2  // специализированная средняя школа
    EducationType3            = 3  // незаконченное высшее образование
    EducationType4            = 4  // высшее образование
    EducationType5            = 5  // два и более высших образования
    EducationType6            = 6  // ученая степень
    EducationType8            = 8  // другое
    EducationType9            = 9  // не известно
    EducationType99           = 99 // значение не передается
)

// семейное положение
type Marital uint32

const (
    MaritalType0  Marital = 0  // холост / не замужем
    MaritalType1          = 1  // женат / замужем
    MaritalType2          = 2  // разведен / разведена
    MaritalType3          = 3  // вдовец / вдова
    MaritalType4          = 4  // гражданский брак / совместное проживание
    MaritalType9          = 9  // не известно
    MaritalType99         = 99 // значение не передается
)

type EmployerSize uint32

const (
    EmployerSizeType0  EmployerSize = 0  // < 50 человек
    EmployerSizeType1               = 1  // 50-100 человек
    EmployerSizeType2               = 2  // 101-249 человек
    EmployerSizeType3               = 3  // 250-499 человек
    EmployerSizeType4               = 4  // > 500 человек
    EmployerSizeType9               = 9  // не известно (EMPTY)
    EmployerSizeType99              = 99 // значение не передается
)

type BusinessIndustry uint32

const (
    BusinessIndustryType0  BusinessIndustry = 0  // промышленность и машиностроение
    BusinessIndustryType1                   = 1  // сельское хозяйство
    BusinessIndustryType2                   = 2  // строительство
    BusinessIndustryType3                   = 3  // горное дело
    BusinessIndustryType4                   = 4  // энергетика
    BusinessIndustryType5                   = 5  // оптовая торговля
    BusinessIndustryType6                   = 6  // финансовое дело и страхование
    BusinessIndustryType7                   = 7  // здравоохранение
    BusinessIndustryType8                   = 8  // социальная помощь
    BusinessIndustryType9                   = 9  // денежные переводы и обмен валют
    BusinessIndustryType10                  = 10 // искусство, развлечение, отдых
    BusinessIndustryType11                  = 11 // казино и игорный бизнес
    BusinessIndustryType12                  = 12 // торговля предметами искусства и антиквариатом
    BusinessIndustryType13                  = 13 // aвиа-, авто- и железнодорожные перевозки, складское хранение
    BusinessIndustryType14                  = 14 // административно-хозяйственные службы
    BusinessIndustryType15                  = 15 // розничная торговля
    BusinessIndustryType16                  = 16 // водо-, тепло- и энергоснабжение
    BusinessIndustryType17                  = 17 // управление компанией
    BusinessIndustryType18                  = 18 // предпринимательская деятельность
    BusinessIndustryType19                  = 19 // научная, техническая и профессиональная деятельность
    BusinessIndustryType20                  = 20 // образование
    BusinessIndustryType21                  = 21 // торговля ювелирными украшениями и драгоценными металлами
    BusinessIndustryType22                  = 22 // посреднические услуги по продаже и аренде недвижимости
    BusinessIndustryType23                  = 23 // снабжение
    BusinessIndustryType24                  = 24 // пресса, телевидение и радио
    BusinessIndustryType25                  = 25 // государственное управление
    BusinessIndustryType26                  = 26 // ресторан / кафе
    BusinessIndustryType27                  = 27 // бытовые услуги
    BusinessIndustryType28                  = 28 // аудиторские услуги / консалтинг
    BusinessIndustryType29                  = 29 // туризм
    BusinessIndustryType30                  = 30 // юридические и нотариальные услуги
    BusinessIndustryType31                  = 31 // ТЭК, добывающая промышленность
    BusinessIndustryType32                  = 32 // детективное и / или охранное предприятие
    BusinessIndustryType33                  = 33 // гостиницы
    BusinessIndustryType34                  = 34 // посредническая деятельность
    BusinessIndustryType35                  = 35 // издательская деятельность / рекламная деятельность
    BusinessIndustryType36                  = 36 // информатика, телекоммуникации
    BusinessIndustryType37                  = 37 // легкая и пищевая промышленность
    BusinessIndustryType38                  = 38 // государственные органы (в том числе правоохранительные)
    BusinessIndustryType39                  = 39 // вооруженные силы
    BusinessIndustryType40                  = 40 // салоны красоты, фитнес центры
    BusinessIndustryType41                  = 41 // транспорт
    BusinessIndustryType42                  = 42 // сборочное производство
    BusinessIndustryType97                  = 97 // другое
    BusinessIndustryType98                  = 98 // не известно
    BusinessIndustryType99                  = 99 // значение не передается
)

type IncomeProof uint32

const (
    IncomeProofType1  IncomeProof = 1  // справка 2-НДФЛ
    IncomeProofType2              = 2  // автомобиль
    IncomeProofType3              = 3  // недвижимость
    IncomeProofType4              = 4  // выписка по счету зарплатной карты
    IncomeProofType5              = 5  // справка из ПФР / выписка с пенсионного счета
    IncomeProofType97             = 97 // справка в свободной форме
    IncomeProofType98             = 98 // нет подтверждения
    IncomeProofType99             = 99 // значение не передается
)

type ProductType uint32

const (
    ProductTypeType0  ProductType = 0  // неизвестный тип кредита
    ProductTypeType1              = 1  // кредит на автомобиль
    ProductTypeType2              = 2  // лизинг
    ProductTypeType3              = 3  // ипотека
    ProductTypeType4              = 4  // кредитная карта
    ProductTypeType5              = 5  // POS кредит (потребительский кредит, кредит на товар)
    ProductTypeType6              = 6  // кредит на развитие бизнеса
    ProductTypeType7              = 7  // кредит на пополнение оборотных средств
    ProductTypeType8              = 8  // кредит на покупку оборудования
    ProductTypeType9              = 9  // кредит на строительство недвижимости
    ProductTypeType10             = 10 // кредит на покупку акций (маржинальное кредитование)
    ProductTypeType11             = 11 // межбанковский кредит
    ProductTypeType12             = 12 // кредит мобильного оператора
    ProductTypeType13             = 13 // кредит на обучение
    ProductTypeType14             = 14 // дебетовая карта с овердрафтом
    ProductTypeType15             = 15 // ипотека (первичный рынок)
    ProductTypeType16             = 16 // ипотека (вторичный рынок)
    ProductTypeType17             = 17 // ипотека (ломбардный кредит)
    ProductTypeType18             = 18 // кредит наличными (нецелевой)
    ProductTypeType19             = 19 // микрозайм
    ProductTypeType20             = 20 // нецелевой кредит под залог автомобиля или недвижимости
    ProductTypeType21             = 21 // депозит
    ProductTypeType99             = 99 // другой тип кредита
)

type OriginalChannel uint32

const (
    OriginalChannelType1  OriginalChannel = 1  // отделение
    OriginalChannelType2                  = 2  // колл-центр
    OriginalChannelType3                  = 3  // брокер
    OriginalChannelType4                  = 4  // интернет
    OriginalChannelType5                  = 5  // кросс-селл
    OriginalChannelType6                  = 6  // точка продаж
    OriginalChannelType7                  = 7  // корпоративные продажи
    OriginalChannelType98                 = 98 // другое
    OriginalChannelType99                 = 99 // значение не передается
)

type SumCurrency string

func (p SumCurrency) ToString() string {
    return string(p)
}

const (
    SumCurrencyType840 SumCurrency = "840" // американский доллар
    SumCurrencyTypeUSD             = "USD" // американский доллар
    SumCurrencyType810             = "810" // рубль
    SumCurrencyTypeRUR             = "RUR" // рубль
    SumCurrencyTypeRUB             = "RUB" // рубль
    SumCurrencyType978             = "978" // евро
    SumCurrencyTypeEUR             = "EUR" // евро
    SumCurrencyType756             = "756" // швейцарский франк
    SumCurrencyTypeCHF             = "CHF" // швейцарский франк
    SumCurrencyType392             = "392" // японская йена
    SumCurrencyTypeJPY             = "JPY" // японская йена
)

type CollateralExistence uint32

const (
    CollateralExistenceType0  CollateralExistence = 0  // нет
    CollateralExistenceType1                      = 1  // да
    CollateralExistenceType99                     = 99 // значение не передает
)

type PurchaseExistence uint32

const (
    PurchaseExistenceType0  PurchaseExistence = 0  // нет
    PurchaseExistenceType1                    = 1  // да
    PurchaseExistenceType99                   = 99 // значение не передается
)

type NewApplicant uint32

const (
    NewApplicantType0  NewApplicant = 0  // нет
    NewApplicantType1               = 1  // да
    NewApplicantType9               = 9  // не определен
    NewApplicantType99              = 99 // значение не передается
)

type ApplicantType uint32

const (
    ApplicantTypeType1 ApplicantType = 1 // заемщик
    ApplicantTypeType2               = 2 // созаемщик
    ApplicantTypeType3               = 3 // поручитель
)

type ResponseIsNeeded uint32

const (
    ResponseIsNeededType0 = 0 // выходной вектор не рассчитывается в ответ на запрос. В данном случае не происходит расчета Предикторов, Правил и Баллов подозрительности и Выходной вектор не предоставляется.
    ResponseIsNeededType1 = 1 // выходной вектор рассчитывается по кредитным заявкам РБД в зависимости от наличия фотографий по ним – с учетом фотографий Аппликантов при их наличии или без учета фотографий Аппликантов при их отсутствии;
    ResponseIsNeededType3 = 3 // выходной вектор рассчитывается по кредитным заявкам РБД только с учетом фотографий Аппликантов;
)

type ApplicationStatus uint32

const (
    ApplicationStatusType1 ApplicationStatus = 1 // одобрена
    ApplicationStatusType2                   = 2 // отказана
    ApplicationStatusType3                   = 3 // одобрена, но кредит не выдан
    ApplicationStatusType4                   = 4 // одобрена. Кредит выдан
    ApplicationStatusType5                   = 5 // аннулирована
    ApplicationStatusType8                   = 8 // закрыт автоматически (Проставляется в системе автоматически. Для внутреннего использования. Партнер не должен выгружать данное значение)
    ApplicationStatusType9                   = 9 // не определен
)

type ApplicationFraudStatus uint32

const (
    ApplicationFraudStatusType1 ApplicationFraudStatus = 1 // мошенничество
    ApplicationFraudStatusType2                        = 2 // фрод
    ApplicationFraudStatusType3                        = 3 // нет признаков фрода
    ApplicationFraudStatusType4                        = 4 // подозрение в мошенничестве
    ApplicationFraudStatusType8                        = 8 // закрыт автоматически (Проставляется в системе автоматически. Для внутреннего использования. Партнер не должен выгружать данное значение)
    ApplicationFraudStatusType9                        = 9 // не определен
)

type DefaultStatus uint32

const (
    DefaultStatusType1 DefaultStatus = 1 // дефолт
    DefaultStatusType2               = 2 // технический дефолт
    DefaultStatusType3               = 3 // нет дефолта
    DefaultStatusType8               = 8 // закрыт автоматически (Проставляется в системе автоматически. Для внутреннего использования. Партнер не должен выгружать данное значение)
    DefaultStatusType9               = 9 // не определен
)

type Status uint32

const (
    StatusType0  Status = 0  // без ошибок
    StatusType1         = 1  // заявка отправлена в карантин
    StatusType2         = 2  // не заданы обязательные поля
    StatusType3         = 3  // не найдена заявка с переданным Application ID
    StatusType4         = 4  // зарезервирован
    StatusType5         = 5  // данные не соответствуют формату
    StatusType6         = 6  // ошибка сервиса нормализации
    StatusType7         = 7  // не задан статус кредитной заявки
    StatusType8         = 8  // не задан фрод-статус кредитной заявки
    StatusType9         = 9  // не задан дефолт-статус кредитной заявки
    StatusType10        = 10 // выходной вектор еще не рассчитан
    StatusType11        = 11 // флаг предоставления выходного вектора из Системы не передавался. Выходной вектор не рассчитывался.
    StatusType12        = 12 // кредитная заявка с данным ID обрабатывается другим пользователем
    StatusType13        = 13 // кредитная заявка для данного пользователя не доступна
    StatusType14        = 14 // кредитная заявка уже была отправлена на обработку ранее
    StatusType15        = 15 // кредитная заявка с данным ID уже есть в Системе
    StatusType20        = 20 // зарезервирован
    StatusType21        = 21 // зарезервирован
    StatusType22        = 22 // зарезервирован
    StatusType23        = 23 // невозможно установить данный статус кредитной заявки. Данный статус заявки уже был установлен ранее
    StatusType24        = 24 // невозможно установить данный фрод-статус кредитной заявки. Данный фрод-статус заявки уже был установлен ранее
    StatusType25        = 25 // невозможно установить данный дефолт-статус кредитной заявки. Данный дефолт-статус заявки уже был установлен ранее
    StatusType26        = 26 // неверно задан статус заявки. Статус «Закрыт автоматически» устанавливается в системе автоматически и только для внутреннего использования
    StatusType27        = 27 // неверно задан фрод-статус заявки. Фрод-статус «Закрыт автоматически» устанавливается в системе автоматически и только для внутреннего использования
    StatusType28        = 28 // неверно задан дефолт-статус заявки. Дефолт-статус «Закрыт автоматически» устанавливается в системе автоматически и только для внутреннего использования
    StatusType30        = 30 // ошибка при обращении к ПО Оператора
    StatusType31        = 31 // не задан признак запроса к базе данных Участников
    StatusType39        = 39 // при сравнении контрольных сумм фотографии произошла ошибка. Принятая от Партнера контрольная сумма фотографии (checksumphoto) не совпадает с контрольной суммой, вычисленной Оператором от принятого файла фотографии (photo)
    StatusType40        = 40 // при загрузке фотографии произошла ошибка. Файл фотографии не читается
    StatusType41        = 41 // фотография не пригодна к обработке. На фотографии отсутствуют глаза
    StatusType42        = 42 // фотография не пригодна к обработке. Слишком низкое качество фотографии
    StatusType43        = 43 // фотография не пригодна к обработке. Лицо на фотографии обрезано
    StatusType44        = 44 // файл с фотографией не загружен. Размер файла фотографии превышает допустимый предел
    StatusType45        = 45 // для данной кредитной заявки уже имеется фотография Аппликанта. Повторный прием фотографии невозможен
    StatusType46        = 46 // предупреждение! На фотографии присутствует несколько лиц. Автоматически Системой было выбрано наиболее крупное! Загрузка фотографии прошла успешно
    StatusType49        = 49 // фотография с данным уникальным идентификатором уже имеется в Системе
    StatusType50        = 50 // фотография с данным уникальным идентификатором не найдена в Системе
    StatusType47        = 47 // предупреждение! Фотография Аппликанта не найдена для данной кредитной заявки. При расчете выходного вектора не были использованы биометрические Правила
    StatusType48        = 48 // фотография Аппликанта не найдена для данной кредитной заявки. Выходной вектор не рассчитан
    StatusType60        = 60 // партнер заблокирован в системе. Загрузка файла {file} невозможна
    StatusType61        = 61 // файл {file} не принимается. Неправильное имя файла
    StatusType62        = 62 // подпись и/или шифрование выполнены не корректно
    StatusType63        = 63 // файл {file} подписан неизвестным сертификатом
    StatusType64        = 64 // файл {file} не принимается, т.к. архив не распаковывается или является пустым
    StatusType65        = 65 // файл {file} не принимается, т.к. в архиве содержатся посторонние файлы
    StatusType66        = 66 // передано некорректное количество полей
    StatusType67        = 67 // архив с файлами фотографий {photoarchive} не найден
    StatusType68        = 68 // файл с фотографией {photofile} не найден в архиве {photoarchive}
    StatusType90        = 90 // ошибка справочника Oracle
    StatusType98        = 98 // нет соединения с ЛБД
    StatusType99        = 99 // другая ошибка. См. таблицу логов
)

type PhoneType string

const (
    PhoneType1 PhoneType = "1" // мобильный
    PhoneType2           = "2" // домашний
    PhoneType3           = "3" // рабочий
    PhoneType4           = "4" // мобильный рабочий
    PhoneType5           = "5" // факс
    PhoneType9           = "9" // другой
)

type ApplicationWay string

const (
    ApplicationWayType1 ApplicationWay = "1" // посреднический (через агента, через брокера)
    ApplicationWayType2                = "2" // дистанционный (с использование средств телекоммуникации)
    ApplicationWayType3                = "3" // прямой (прямое обращение в отделение или офис Партнера)
)

type LoanCredit string

const (
    LoanCreditType1 LoanCredit = "1" // договор займа
    LoanCreditType2            = "2" // договор кредита
    LoanCreditType9            = "9" // неизвестно
)

type FullPay string

const (
    FullPayType0 FullPay = "0" // обязательства не исполнены в полном объеме
    FullPayType1         = "1" // обязательства исполнены в полном объеме
)

type CreditActive string

const (
    CreditActiveType0  CreditActive = "0"  // договор закрыт
    CreditActiveType1  CreditActive = "1"  // договор активен нет просроченных платежей
    CreditActiveType2  CreditActive = "2"  // договор продан (переуступка прав требований)
    CreditActiveType3  CreditActive = "3"  // безнадежный долг (списан с баланса)
    CreditActiveType4  CreditActive = "4"  // договор рефинансирован
    CreditActiveType5  CreditActive = "5"  // договор передан коллекторам
    CreditActiveType7  CreditActive = "7"  // договор отменен
    CreditActiveType9  CreditActive = "9"  // субъект ки освобожден от дальнейшего исполнения требований кредиторов так как судом или арбитражным судом принято решение о признании субъекта ки банкротом
    CreditActiveType10 CreditActive = "10" // договор активен просрочка от 1 до 5 дней
    CreditActiveType11 CreditActive = "11" // договор активен просрочка от 6 до 29 договор
    CreditActiveType12 CreditActive = "12" // договор активен просрочка от 30 до 59 дней
    CreditActiveType13 CreditActive = "13" // договор активен просрочка от 60 до 89 дней
    CreditActiveType14 CreditActive = "14" // договор активен просрочка от 90 до 119 дней
    CreditActiveType15 CreditActive = "15" // договор активен просрочка от 120 до 149 дней
    CreditActiveType16 CreditActive = "16" // договор активен просрочка от 150 до 179 дней
    CreditActiveType17 CreditActive = "17" // договор активен просрочка от 180 до 209 дней
    CreditActiveType18 CreditActive = "18" // договор активен просрочка от 210 до 239 дней
    CreditActiveType19 CreditActive = "19" // вся задолженность переведена в просроченную задолженность / договор активен просрочка 240 и более дней
    CreditActiveType20 CreditActive = "20" // договор расторгнут
)

type Collateral string

const (
    CollateralType0 Collateral = "0" // не было погашения за счет обеспечения
    CollateralType1            = "1" // было погашение за счет обеспечения
    CollateralType2            = "2" // было погашение за счет поручителя
    CollateralType3            = "3" // погашен/частично погашен за счет средств правообладателя требований по договору
)
