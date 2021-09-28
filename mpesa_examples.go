package yuri

import (
	"log"
)

func B2CExample(ConsumerKey, ConsumerSecret, sec string) {
	mpesa := NewMpesa(ConsumerKey, ConsumerSecret, false)

	///security, err := mpesa.GetSecurityCredential(initiatorPassword)
	resp, err := mpesa.B2CRequest(B2CRequestBody{
		InitiatorName:      "",
		SecurityCredential: sec,
		CommandID:          "",
		Amount:             10,
		PartyA:             "",
		PartyB:             "",
		Remarks:            "testing",
		QueueTimeOutURL:    "",
		ResultURL:          "",
		Occassion:          "",
	})
	if err != nil {
		log.Fatal(err)
	}
	PrintStruct(resp)
}
