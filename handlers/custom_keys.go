// Code generated by go generate; DO NOT EDIT.
package handlers

import (
	"github.com/bgadrian/fastfaker"
	"strconv"
)

type fakerer func(*fastfaker.Faker) string
var keys map[string]fakerer

func init() {
	keys = make(map[string]fakerer)
	
	keys["beeralcohol"] = func(f *fastfaker.Faker) string { return f.BeerAlcohol() }
	keys["beerblg"] = func(f *fastfaker.Faker) string { return f.BeerBlg() }
	keys["beerhop"] = func(f *fastfaker.Faker) string { return f.BeerHop() }
	keys["beeribu"] = func(f *fastfaker.Faker) string { return f.BeerIbu() }
	keys["beermalt"] = func(f *fastfaker.Faker) string { return f.BeerMalt() }
	keys["beername"] = func(f *fastfaker.Faker) string { return f.BeerName() }
	keys["beerstyle"] = func(f *fastfaker.Faker) string { return f.BeerStyle() }
	keys["beeryeast"] = func(f *fastfaker.Faker) string { return f.BeerYeast() }
	keys["bs"] = func(f *fastfaker.Faker) string { return f.BS() }
	keys["buzzword"] = func(f *fastfaker.Faker) string { return f.BuzzWord() }
	keys["carmaker"] = func(f *fastfaker.Faker) string { return f.CarMaker() }
	keys["carmodel"] = func(f *fastfaker.Faker) string { return f.CarModel() }
	keys["chromeuseragent"] = func(f *fastfaker.Faker) string { return f.ChromeUserAgent() }
	keys["city"] = func(f *fastfaker.Faker) string { return f.City() }
	keys["color"] = func(f *fastfaker.Faker) string { return f.Color() }
	keys["company"] = func(f *fastfaker.Faker) string { return f.Company() }
	keys["companysuffix"] = func(f *fastfaker.Faker) string { return f.CompanySuffix() }
	keys["country"] = func(f *fastfaker.Faker) string { return f.Country() }
	keys["countryabr"] = func(f *fastfaker.Faker) string { return f.CountryAbr() }
	keys["creditcardcvv"] = func(f *fastfaker.Faker) string { return f.CreditCardCvv() }
	keys["creditcardexp"] = func(f *fastfaker.Faker) string { return f.CreditCardExp() }
	keys["creditcardnumber"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.CreditCardNumber())) }
	keys["creditcardnumberluhn"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.CreditCardNumberLuhn())) }
	keys["creditcardtype"] = func(f *fastfaker.Faker) string { return f.CreditCardType() }
	keys["currencylong"] = func(f *fastfaker.Faker) string { return f.CurrencyLong() }
	keys["currencyshort"] = func(f *fastfaker.Faker) string { return f.CurrencyShort() }
	keys["day"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Day())) }
	keys["digit"] = func(f *fastfaker.Faker) string { return f.Digit() }
	keys["domainname"] = func(f *fastfaker.Faker) string { return f.DomainName() }
	keys["domainsuffix"] = func(f *fastfaker.Faker) string { return f.DomainSuffix() }
	keys["email"] = func(f *fastfaker.Faker) string { return f.Email() }
	keys["extension"] = func(f *fastfaker.Faker) string { return f.Extension() }
	keys["firefoxuseragent"] = func(f *fastfaker.Faker) string { return f.FirefoxUserAgent() }
	keys["firstname"] = func(f *fastfaker.Faker) string { return f.FirstName() }
	keys["float32"] = func(f *fastfaker.Faker) string { return strconv.FormatFloat(float64(f.Float32()), 103, -1, 32) }
	keys["float64"] = func(f *fastfaker.Faker) string { return strconv.FormatFloat(float64(f.Float64()), 103, -1, 64) }
	keys["fueltype"] = func(f *fastfaker.Faker) string { return f.FuelType() }
	keys["gender"] = func(f *fastfaker.Faker) string { return f.Gender() }
	keys["hackerabbreviation"] = func(f *fastfaker.Faker) string { return f.HackerAbbreviation() }
	keys["hackeradjective"] = func(f *fastfaker.Faker) string { return f.HackerAdjective() }
	keys["hackeringverb"] = func(f *fastfaker.Faker) string { return f.HackerIngverb() }
	keys["hackernoun"] = func(f *fastfaker.Faker) string { return f.HackerNoun() }
	keys["hackerphrase"] = func(f *fastfaker.Faker) string { return f.HackerPhrase() }
	keys["hackerverb"] = func(f *fastfaker.Faker) string { return f.HackerVerb() }
	keys["hexcolor"] = func(f *fastfaker.Faker) string { return f.HexColor() }
	keys["hipsterword"] = func(f *fastfaker.Faker) string { return f.HipsterWord() }
	keys["hour"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Hour())) }
	keys["httpmethod"] = func(f *fastfaker.Faker) string { return f.HTTPMethod() }
	keys["int16"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Int16())) }
	keys["int32"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Int32())) }
	keys["int64"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Int64())) }
	keys["int8"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Int8())) }
	keys["ipv4address"] = func(f *fastfaker.Faker) string { return f.IPv4Address() }
	keys["ipv6address"] = func(f *fastfaker.Faker) string { return f.IPv6Address() }
	keys["jobdescriptor"] = func(f *fastfaker.Faker) string { return f.JobDescriptor() }
	keys["joblevel"] = func(f *fastfaker.Faker) string { return f.JobLevel() }
	keys["jobtitle"] = func(f *fastfaker.Faker) string { return f.JobTitle() }
	keys["lastname"] = func(f *fastfaker.Faker) string { return f.LastName() }
	keys["latitude"] = func(f *fastfaker.Faker) string { return strconv.FormatFloat(float64(f.Latitude()), 103, -1, 64) }
	keys["letter"] = func(f *fastfaker.Faker) string { return f.Letter() }
	keys["longitude"] = func(f *fastfaker.Faker) string { return strconv.FormatFloat(float64(f.Longitude()), 103, -1, 64) }
	keys["mimetype"] = func(f *fastfaker.Faker) string { return f.MimeType() }
	keys["minute"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Minute())) }
	keys["month"] = func(f *fastfaker.Faker) string { return f.Month() }
	keys["name"] = func(f *fastfaker.Faker) string { return f.Name() }
	keys["nameprefix"] = func(f *fastfaker.Faker) string { return f.NamePrefix() }
	keys["namesuffix"] = func(f *fastfaker.Faker) string { return f.NameSuffix() }
	keys["nanosecond"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.NanoSecond())) }
	keys["operauseragent"] = func(f *fastfaker.Faker) string { return f.OperaUserAgent() }
	keys["phone"] = func(f *fastfaker.Faker) string { return f.Phone() }
	keys["phoneformatted"] = func(f *fastfaker.Faker) string { return f.PhoneFormatted() }
	keys["safariuseragent"] = func(f *fastfaker.Faker) string { return f.SafariUserAgent() }
	keys["safecolor"] = func(f *fastfaker.Faker) string { return f.SafeColor() }
	keys["second"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Second())) }
	keys["simplestatuscode"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.SimpleStatusCode())) }
	keys["ssn"] = func(f *fastfaker.Faker) string { return f.SSN() }
	keys["state"] = func(f *fastfaker.Faker) string { return f.State() }
	keys["stateabr"] = func(f *fastfaker.Faker) string { return f.StateAbr() }
	keys["statuscode"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.StatusCode())) }
	keys["street"] = func(f *fastfaker.Faker) string { return f.Street() }
	keys["streetname"] = func(f *fastfaker.Faker) string { return f.StreetName() }
	keys["streetnumber"] = func(f *fastfaker.Faker) string { return f.StreetNumber() }
	keys["streetprefix"] = func(f *fastfaker.Faker) string { return f.StreetPrefix() }
	keys["streetsuffix"] = func(f *fastfaker.Faker) string { return f.StreetSuffix() }
	keys["timezone"] = func(f *fastfaker.Faker) string { return f.TimeZone() }
	keys["timezoneabv"] = func(f *fastfaker.Faker) string { return f.TimeZoneAbv() }
	keys["timezonefull"] = func(f *fastfaker.Faker) string { return f.TimeZoneFull() }
	keys["timezoneoffset"] = func(f *fastfaker.Faker) string { return strconv.FormatFloat(float64(f.TimeZoneOffset()), 103, -1, 32) }
	keys["transmissiongeartype"] = func(f *fastfaker.Faker) string { return f.TransmissionGearType() }
	keys["uint16"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Uint16())) }
	keys["uint32"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Uint32())) }
	keys["uint64"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Uint64())) }
	keys["uint8"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Uint8())) }
	keys["url"] = func(f *fastfaker.Faker) string { return f.URL() }
	keys["useragent"] = func(f *fastfaker.Faker) string { return f.UserAgent() }
	keys["username"] = func(f *fastfaker.Faker) string { return f.Username() }
	keys["uuid"] = func(f *fastfaker.Faker) string { return f.UUID() }
	keys["vehicletype"] = func(f *fastfaker.Faker) string { return f.VehicleType() }
	keys["weekday"] = func(f *fastfaker.Faker) string { return f.WeekDay() }
	keys["word"] = func(f *fastfaker.Faker) string { return f.Word() }
	keys["year"] = func(f *fastfaker.Faker) string { return strconv.Itoa(int(f.Year())) }
	keys["zip"] = func(f *fastfaker.Faker) string { return f.Zip() }
	
}
