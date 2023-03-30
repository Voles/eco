// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package payment

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de_DE": &dictionary{index: de_DEIndex, data: de_DEData},
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
	}
	fallback := language.MustParse("de-DE")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"BIC (falls nötig)":  13,
	"Bank (falls nötig)": 14,
	"Bargeld":            2,
	"Bezahle den angegebenen Betrag in Monero (XMR) oder Bitcoin (BTC). Der Betrag muss innerhalb von 60 Minuten vollständig und als einzelne Transaktion auf der angegebenen Adresse eingehen. Falls deine Zahlung verspätet eintrifft, müssen wir sie manuell bestätigen. Im Zweifel kontaktiere uns bitte.": 4,
	"Bitte lege einen Zettel mit der Bestellnummer bei": 7,
	"Falls du TOR oder einen VPN benutzt: Die angezeigten Bezahlmöglichkeiten sind von der Länderzuordnung deiner IP-Adresse abhängig. Darüber hinaus blockiert PayPal manche TOR Exit Nodes. In dem Fall versuche es mit „New Circuit for this Site“.": 9,
	"Führe eine SEPA-Überweisung (einheitlicher Euro-Zahlungsverkehrsraum) auf unser deutsches Bankkonto aus. Wir prüfen es täglich manuell auf neue Zahlungseingänge. Wir werden deinen Namen und deine IBAN auf unserem Kontoauszug sehen.":           10,
	"IBAN":                 12,
	"Kontoinhaber":         11,
	"Monero oder Bitcoin":  0,
	"SEPA-Banküberweisung": 3,
	"Sende uns Bargeld in einem versichertem Brief oder Paket. Nachdem wir das Geld entnommen haben, schreddern wir den Brief. Bitte beachte die Höchstgrenzen deines Postunternehmens für den Bargeldversand (z. B. Deutsche Post „Einschreiben Wert“ bis 100 Euro innerhalb Deutschlands, DHL Paket bis 500 Euro). Sende es an:": 6,
	"Wir übermitteln nur die Bestellnummer an PayPal. Deine bestellten Artikel sowie die Details zu Lieferung oder Abholung werden nicht an PayPal gesendet.":                                                                                                                                                                      8,
	"Zur Bezahlung mit Monero oder Bitcoin": 5,
	"btcpay:de-DE":                          1,
	"Überweisungszweck":                     15,
}

var de_DEIndex = []uint32{ // 17 elements
	0x00000000, 0x00000014, 0x00000021, 0x00000029,
	0x0000003f, 0x0000016c, 0x00000192, 0x000002d5,
	0x00000307, 0x000003a0, 0x0000049a, 0x00000587,
	0x00000594, 0x00000599, 0x000005ac, 0x000005c0,
	0x000005d3,
} // Size: 92 bytes

const de_DEData string = "" + // Size: 1491 bytes
	"\x02Monero oder Bitcoin\x02btcpay:de-DE\x02Bargeld\x02SEPA-Banküberweisu" +
	"ng\x02Bezahle den angegebenen Betrag in Monero (XMR) oder Bitcoin (BTC)." +
	" Der Betrag muss innerhalb von 60 Minuten vollständig und als einzelne T" +
	"ransaktion auf der angegebenen Adresse eingehen. Falls deine Zahlung ver" +
	"spätet eintrifft, müssen wir sie manuell bestätigen. Im Zweifel kontakti" +
	"ere uns bitte.\x02Zur Bezahlung mit Monero oder Bitcoin\x02Sende uns Bar" +
	"geld in einem versichertem Brief oder Paket. Nachdem wir das Geld entnom" +
	"men haben, schreddern wir den Brief. Bitte beachte die Höchstgrenzen dei" +
	"nes Postunternehmens für den Bargeldversand (z. B. Deutsche Post „Einsch" +
	"reiben Wert“ bis 100 Euro innerhalb Deutschlands, DHL Paket bis 500 Euro" +
	"). Sende es an:\x02Bitte lege einen Zettel mit der Bestellnummer bei\x02" +
	"Wir übermitteln nur die Bestellnummer an PayPal. Deine bestellten Artike" +
	"l sowie die Details zu Lieferung oder Abholung werden nicht an PayPal ge" +
	"sendet.\x02Falls du TOR oder einen VPN benutzt: Die angezeigten Bezahlmö" +
	"glichkeiten sind von der Länderzuordnung deiner IP-Adresse abhängig. Dar" +
	"über hinaus blockiert PayPal manche TOR Exit Nodes. In dem Fall versuch" +
	"e es mit „New Circuit for this Site“.\x02Führe eine SEPA-Überweisung (ei" +
	"nheitlicher Euro-Zahlungsverkehrsraum) auf unser deutsches Bankkonto aus" +
	". Wir prüfen es täglich manuell auf neue Zahlungseingänge. Wir werden de" +
	"inen Namen und deine IBAN auf unserem Kontoauszug sehen.\x02Kontoinhaber" +
	"\x02IBAN\x02BIC (falls nötig)\x02Bank (falls nötig)\x02Überweisungszweck"

var en_USIndex = []uint32{ // 17 elements
	0x00000000, 0x00000012, 0x0000001c, 0x00000021,
	0x00000034, 0x0000011d, 0x00000139, 0x00000269,
	0x00000296, 0x0000030d, 0x000003d1, 0x000004a2,
	0x000004b1, 0x000004b6, 0x000004c8, 0x000004e0,
	0x000004e8,
} // Size: 92 bytes

const en_USData string = "" + // Size: 1256 bytes
	"\x02Monero or Bitcoin\x02btcpay:en\x02Cash\x02SEPA Bank Transfer\x02Pay " +
	"with Monero (XMR) or Bitcoin (BTC). The full amount must be paid with a " +
	"single transaction to the given address within 60 minutes. If your payme" +
	"nt arrives too late, we have to confirm it manually. If in doubt, please" +
	" contact us.\x02Pay using Monero or Bitcoin\x02Send cash in an insured l" +
	"etter or package to our store address in Germany. After we take out the " +
	"money, we shred the letter. Please check the cash shipment limits of you" +
	"r postal company (e. g. Deutsche Post „Einschreiben Wert“ up to 100 Euro" +
	"s within Germany, DHL Parcel up to 500 Euros). Send it to:\x02Please inc" +
	"lude a note with this order number\x02We only send the order number to P" +
	"ayPal. Your ordered items and delivery or pickup details will not be sen" +
	"t to PayPal.\x02If you use TOR or a VPN: The payment options displayed d" +
	"epend on the country of your IP address. In addition, PayPal blocks some" +
	" TOR exit nodes. In that case, try „New Circuit for this Site“.\x02Make " +
	"a SEPA (Single Euro Payments Area) bank transfer to our German bank acco" +
	"unt. We manually check for new incoming payments every day. We will see " +
	"your name and bank account number on our account statement.\x02Account h" +
	"older\x02IBAN\x02BIC (if required)\x02Bank name (if required)\x02Purpose"

	// Total table size 2931 bytes (2KiB); checksum: 3926848A