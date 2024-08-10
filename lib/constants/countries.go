package constants

import "fmt"

type Country int

const (
	AFGHANISTAN Country = iota
	ALBANIA
	ALGERIA
	ANDORRA
	ANGOLA
	ANTIGUA_AND_BARBUDA
	ARGENTINA
	ARMENIA
	AUSTRALIA
	AUSTRIA
	AZERBAIJAN
	BAHAMAS
	BAHRAIN
	BANGLADESH
	BARBADOS
	BELARUS
	BELGIUM
	BELIZE
	BENIN
	BHUTAN
	BOLIVIA
	BOSNIA_AND_HERZEGOVINA
	BOTSWANA
	BRAZIL
	BRUNEI
	BULGARIA
	BURKINA_FASO
	BURUNDI
	CABO_VERDE
	CAMBODIA
	CAMEROON
	CANADA
	CENTRAL_AFRICAN_REPUBLIC
	CHAD
	CHILE
	CHINA
	COLOMBIA
	COMOROS
	CONGO_DEMOCRATIC_REPUBLIC
	CONGO_REPUBLIC
	COSTA_RICA
	COTE_DIVOIRE
	CROATIA
	CUBA
	CYPRUS
	CZECH_REPUBLIC
	DENMARK
	DJIBOUTI
	DOMINICA
	DOMINICAN_REPUBLIC
	ECUADOR
	EGYPT
	EL_SALVADOR
	EQUATORIAL_GUINEA
	ERITREA
	ESTONIA
	ESWATINI
	ETHIOPIA
	FIJI
	FINLAND
	FRANCE
	GABON
	GAMBIA
	GEORGIA
	GERMANY
	GHANA
	GREECE
	GRENADA
	GUATEMALA
	GUINEA
	GUINEA_BISSAU
	GUYANA
	HAITI
	HONDURAS
	HUNGARY
	ICELAND
	INDIA
	INDONESIA
	IRAN
	IRAQ
	IRELAND
	ISRAEL
	ITALY
	JAMAICA
	JAPAN
	JORDAN
	KAZAKHSTAN
	KENYA
	KIRIBATI
	KOREA_NORTH
	KOREA_SOUTH
	KOSOVO
	KUWAIT
	KYRGYZSTAN
	LAOS
	LATVIA
	LEBANON
	LESOTHO
	LIBERIA
	LIBYA
	LIECHTENSTEIN
	LITHUANIA
	LUXEMBOURG
	MADAGASCAR
	MALAWI
	MALAYSIA
	MALDIVES
	MALI
	MALTA
	MARSHALL_ISLANDS
	MAURITANIA
	MAURITIUS
	MEXICO
	MICRONESIA
	MOLDOVA
	MONACO
	MONGOLIA
	MONTENEGRO
	MOROCCO
	MOZAMBIQUE
	MYANMAR
	NAMIBIA
	NAURU
	NEPAL
	NETHERLANDS
	NEW_ZEALAND
	NICARAGUA
	NIGER
	NIGERIA
	NORTH_MACEDONIA
	NORWAY
	OMAN
	PAKISTAN
	PALAU
	PALESTINE
	PANAMA
	PAPUA_NEW_GUINEA
	PARAGUAY
	PERU
	PHILIPPINES
	POLAND
	PORTUGAL
	QATAR
	ROMANIA
	RUSSIA
	RWANDA
	SAINT_KITTS_AND_NEVIS
	SAINT_LUCIA
	SAINT_VINCENT_AND_GRENADINES
	SAMOA
	SAN_MARINO
	SAO_TOME_AND_PRINCIPE
	SAUDI_ARABIA
	SENEGAL
	SERBIA
	SEYCHELLES
	SIERRA_LEONE
	SINGAPORE
	SLOVAKIA
	SLOVENIA
	SOLOMON_ISLANDS
	SOMALIA
	SOUTH_AFRICA
	SOUTH_SUDAN
	SPAIN
	SRI_LANKA
	SUDAN
	SURINAME
	SWEDEN
	SWITZERLAND
	SYRIA
	TAIWAN
	TAJIKISTAN
	TANZANIA
	THAILAND
	TIMOR_LESTE
	TOGO
	TONGA
	TRINIDAD_AND_TOBAGO
	TUNISIA
	TURKEY
	TURKMENISTAN
	TUVALU
	UGANDA
	UKRAINE
	UNITED_ARAB_EMIRATES
	UNITED_KINGDOM
	UNITED_STATES
	URUGUAY
	UZBEKISTAN
	VANUATU
	VATICAN_CITY
	VENEZUELA
	VIETNAM
	YEMEN
	ZAMBIA
	ZIMBABWE
)

var countries = map[Country][]string{
	AFGHANISTAN:                  {"Afghanistan", "AF"},
	ALBANIA:                      {"Albania", "AL"},
	ALGERIA:                      {"Algeria", "DZ"},
	ANDORRA:                      {"Andorra", "AD"},
	ANGOLA:                       {"Angola", "AO"},
	ANTIGUA_AND_BARBUDA:          {"Antigua and Barbuda", "AG"},
	ARGENTINA:                    {"Argentina", "AR"},
	ARMENIA:                      {"Armenia", "AM"},
	AUSTRALIA:                    {"Australia", "AU"},
	AUSTRIA:                      {"Austria", "AT"},
	AZERBAIJAN:                   {"Azerbaijan", "AZ"},
	BAHAMAS:                      {"Bahamas", "BS"},
	BAHRAIN:                      {"Bahrain", "BH"},
	BANGLADESH:                   {"Bangladesh", "BD"},
	BARBADOS:                     {"Barbados", "BB"},
	BELARUS:                      {"Belarus", "BY"},
	BELGIUM:                      {"Belgium", "BE"},
	BELIZE:                       {"Belize", "BZ"},
	BENIN:                        {"Benin", "BJ"},
	BHUTAN:                       {"Bhutan", "BT"},
	BOLIVIA:                      {"Bolivia", "BO"},
	BOSNIA_AND_HERZEGOVINA:       {"Bosnia and Herzegovina", "BA"},
	BOTSWANA:                     {"Botswana", "BW"},
	BRAZIL:                       {"Brazil", "BR"},
	BRUNEI:                       {"Brunei", "BN"},
	BULGARIA:                     {"Bulgaria", "BG"},
	BURKINA_FASO:                 {"Burkina Faso", "BF"},
	BURUNDI:                      {"Burundi", "BI"},
	CABO_VERDE:                   {"Cabo Verde", "CV"},
	CAMBODIA:                     {"Cambodia", "KH"},
	CAMEROON:                     {"Cameroon", "CM"},
	CANADA:                       {"Canada", "CA"},
	CENTRAL_AFRICAN_REPUBLIC:     {"Central African Republic", "CF"},
	CHAD:                         {"Chad", "TD"},
	CHILE:                        {"Chile", "CL"},
	CHINA:                        {"China", "CN"},
	COLOMBIA:                     {"Colombia", "CO"},
	COMOROS:                      {"Comoros", "KM"},
	CONGO_DEMOCRATIC_REPUBLIC:    {"Congo (Democratic Republic)", "CD"},
	CONGO_REPUBLIC:               {"Congo (Republic)", "CG"},
	COSTA_RICA:                   {"Costa Rica", "C"},
	COTE_DIVOIRE:                 {"Côte d'Ivoire", "CI"},
	CROATIA:                      {"Croatia", "HR"},
	CUBA:                         {"Cuba", "CU"},
	CYPRUS:                       {"Cyprus", "CY"},
	CZECH_REPUBLIC:               {"Czech Republic", "CZ"},
	DENMARK:                      {"Denmark", "DK"},
	DJIBOUTI:                     {"Djibouti", "DJ"},
	DOMINICA:                     {"Dominica", "DM"},
	DOMINICAN_REPUBLIC:           {"Dominican Republic", "DO"},
	ECUADOR:                      {"Ecuador", "EC"},
	EGYPT:                        {"Egypt", "EG"},
	EL_SALVADOR:                  {"El Salvador", "SV"},
	EQUATORIAL_GUINEA:            {"Equatorial Guinea", "GQ"},
	ERITREA:                      {"Eritrea", "ER"},
	ESTONIA:                      {"Estonia", "EE"},
	ESWATINI:                     {"Eswatini", "SZ"},
	ETHIOPIA:                     {"Ethiopia", "ET"},
	FIJI:                         {"Fiji", "FJ"},
	FINLAND:                      {"Finland", "FI"},
	FRANCE:                       {"France", "FR"},
	GABON:                        {"Gabon", "GA"},
	GAMBIA:                       {"Gambia", "GM"},
	GEORGIA:                      {"Georgia", "GE"},
	GERMANY:                      {"Germany", "DE"},
	GHANA:                        {"Ghana", "GH"},
	GREECE:                       {"Greece", "GR"},
	GRENADA:                      {"Grenada", "GD"},
	GUATEMALA:                    {"Guatemala", "GT"},
	GUINEA:                       {"Guinea", "GN"},
	GUINEA_BISSAU:                {"Guinea-Bissau", "GW"},
	GUYANA:                       {"Guyana", "GY"},
	HAITI:                        {"Haiti", "HT"},
	HONDURAS:                     {"Honduras", "HN"},
	HUNGARY:                      {"Hungary", "HU"},
	ICELAND:                      {"Iceland", "IS"},
	INDIA:                        {"India", "IN"},
	INDONESIA:                    {"Indonesia", "ID"},
	IRAN:                         {"Iran", "IR"},
	IRAQ:                         {"Iraq", "IQ"},
	IRELAND:                      {"Ireland", "IE"},
	ISRAEL:                       {"Israel", "IL"},
	ITALY:                        {"Italy", "IT"},
	JAMAICA:                      {"Jamaica", "JM"},
	JAPAN:                        {"Japan", "JP"},
	JORDAN:                       {"Jordan", "JO"},
	KAZAKHSTAN:                   {"Kazakhstan", "KZ"},
	KENYA:                        {"Kenya", "KE"},
	KIRIBATI:                     {"Kiribati", "KI"},
	KOREA_NORTH:                  {"Korea (North)", "KP"},
	KOREA_SOUTH:                  {"Korea (South)", "KR"},
	KOSOVO:                       {"Kosovo", "XK"},
	KUWAIT:                       {"Kuwait", "KW"},
	KYRGYZSTAN:                   {"Kyrgyzstan", "KG"},
	LAOS:                         {"Laos", "LA"},
	LATVIA:                       {"Latvia", "LV"},
	LEBANON:                      {"Lebanon", "LB"},
	LESOTHO:                      {"Lesotho", "LS"},
	LIBERIA:                      {"Liberia", "LR"},
	LIBYA:                        {"Libya", "LY"},
	LIECHTENSTEIN:                {"Liechtenstein", "LI"},
	LITHUANIA:                    {"Lithuania", "LT"},
	LUXEMBOURG:                   {"Luxembourg", "LU"},
	MADAGASCAR:                   {"Madagascar", "MG"},
	MALAWI:                       {"Malawi", "MW"},
	MALAYSIA:                     {"Malaysia", "MY"},
	MALDIVES:                     {"Maldives", "MV"},
	MALI:                         {"Mali", "ML"},
	MALTA:                        {"Malta", "MT"},
	MARSHALL_ISLANDS:             {"Marshall Islands", "MH"},
	MAURITANIA:                   {"Mauritania", "MR"},
	MAURITIUS:                    {"Mauritius", "MU"},
	MEXICO:                       {"Mexico", "MX"},
	MICRONESIA:                   {"Micronesia", "FM"},
	MOLDOVA:                      {"Moldova", "MD"},
	MONACO:                       {"Monaco", "MC"},
	MONGOLIA:                     {"Mongolia", "MN"},
	MONTENEGRO:                   {"Montenegro", "ME"},
	MOROCCO:                      {"Morocco", "MA"},
	MOZAMBIQUE:                   {"Mozambique", "MZ"},
	MYANMAR:                      {"Myanmar", "MM"},
	NAMIBIA:                      {"Namibia", "NA"},
	NAURU:                        {"Nauru", "NR"},
	NEPAL:                        {"Nepal", "NP"},
	NETHERLANDS:                  {"Netherlands", "NL"},
	NEW_ZEALAND:                  {"New Zealand", "NZ"},
	NICARAGUA:                    {"Nicaragua", "NI"},
	NIGER:                        {"Niger", "NE"},
	NIGERIA:                      {"Nigeria", "NG"},
	NORTH_MACEDONIA:              {"North Macedonia", "MK"},
	NORWAY:                       {"Norway", "NO"},
	OMAN:                         {"Oman", "OM"},
	PAKISTAN:                     {"Pakistan", "PK"},
	PALAU:                        {"Palau", "PW"},
	PALESTINE:                    {"Palestine", "PS"},
	PANAMA:                       {"Panama", "PA"},
	PAPUA_NEW_GUINEA:             {"Papua New Guinea", "PG"},
	PARAGUAY:                     {"Paraguay", "PY"},
	PERU:                         {"Peru", "PE"},
	PHILIPPINES:                  {"Philippines", "PH"},
	POLAND:                       {"Poland", "PL"},
	PORTUGAL:                     {"Portugal", "PT"},
	QATAR:                        {"Qatar", "QA"},
	ROMANIA:                      {"Romania", "RO"},
	RUSSIA:                       {"Russia", "RU"},
	RWANDA:                       {"Rwanda", "RW"},
	SAINT_KITTS_AND_NEVIS:        {"Saint Kitts and Nevis", "KN"},
	SAINT_LUCIA:                  {"Saint Lucia", "LC"},
	SAINT_VINCENT_AND_GRENADINES: {"Saint Vincent and the Grenadines", "VC"},
	SAMOA:                        {"Samoa", "WS"},
	SAN_MARINO:                   {"San Marino", "SM"},
	SAO_TOME_AND_PRINCIPE:        {"Sao Tome and Principe", "ST"},
	SAUDI_ARABIA:                 {"Saudi Arabia", "SA"},
	SENEGAL:                      {"Senegal", "SN"},
	SERBIA:                       {"Serbia", "RS"},
	SEYCHELLES:                   {"Seychelles", "SC"},
	SIERRA_LEONE:                 {"Sierra Leone", "SL"},
	SINGAPORE:                    {"Singapore", "SG"},
	SLOVAKIA:                     {"Slovakia", "SK"},
	SLOVENIA:                     {"Slovenia", "SI"},
	SOLOMON_ISLANDS:              {"Solomon Islands", "SB"},
	SOMALIA:                      {"Somalia", "SO"},
	SOUTH_AFRICA:                 {"South Africa", "ZA"},
	SOUTH_SUDAN:                  {"South Sudan", "SS"},
	SPAIN:                        {"Spain", "ES"},
	SRI_LANKA:                    {"Sri Lanka", "LK"},
	SUDAN:                        {"Sudan", "SD"},
	SURINAME:                     {"Suriname", "SR"},
	SWEDEN:                       {"Sweden", "SE"},
	SWITZERLAND:                  {"Switzerland", "CH"},
	SYRIA:                        {"Syria", "SY"},
	TAIWAN:                       {"Taiwan", "TW"},
	TAJIKISTAN:                   {"Tajikistan", "TJ"},
	TANZANIA:                     {"Tanzania", "TZ"},
	THAILAND:                     {"Thailand", "TH"},
	TIMOR_LESTE:                  {"Timor-Leste", "TL"},
	TOGO:                         {"Togo", "TG"},
	TONGA:                        {"Tonga", "TO"},
	TRINIDAD_AND_TOBAGO:          {"Trinidad and Tobago", "TT"},
	TUNISIA:                      {"Tunisia", "TN"},
	TURKEY:                       {"Turkey", "TR"},
	TURKMENISTAN:                 {"Turkmenistan", "TM"},
	TUVALU:                       {"Tuvalu", "TV"},
	UGANDA:                       {"Uganda", "UG"},
	UKRAINE:                      {"Ukraine", "UA"},
	UNITED_ARAB_EMIRATES:         {"United Arab Emirates", "AE"},
	UNITED_KINGDOM:               {"United Kingdom", "GB"},
	UNITED_STATES:                {"United States", "US"},
	URUGUAY:                      {"Uruguay", "UY"},
	UZBEKISTAN:                   {"Uzbekistan", "UZ"},
	VANUATU:                      {"Vanuatu", "VU"},
	VATICAN_CITY:                 {"Vatican City", "VA"},
	VENEZUELA:                    {"Venezuela", "VE"},
	VIETNAM:                      {"Vietnam", "VN"},
	YEMEN:                        {"Yemen", "YE"},
	ZAMBIA:                       {"Zambia", "ZM"},
	ZIMBABWE:                     {"Zimbabwe", "ZW"},
}

func (c Country) String() ([]string, error) {
	if value, ok := countries[c]; ok {
		return value, nil
	}
	return []string{"", ""}, fmt.Errorf("invalid country")
}
