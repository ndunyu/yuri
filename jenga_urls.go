package yuri

import "fmt"

const jengaTokenUrl = "identity/v2/token"
const jengaKycUrl = "customer/v2/identity/verify"
const jengaAirTimeUrl = "transaction/v2/airtime"
const jengaMerchantsUrl = "transaction/v2/merchants"
const jengaBankToMobileWalletUrl = "transaction/v2/remittance#sendmobile"
const pesaLinkToBankUrl = "transaction/v2/remittance"
const equityToequity = "transaction/v2/remittance#sendeqtybank"
const accountBalance = "account/v2/accounts/balances/%s/%s"

const JengaLiveUrl = "https://api.jengahq.io/"
const JengaSandboxUrl = "https://uat.jengahq.io/"

func (J *Jenga) getBaseUrl() string {
	if !J.Live {

		return JengaSandboxUrl
	}

	return JengaLiveUrl

}
func (J *Jenga)getAccountBalanceUrl(countryCode, accountId string) string {
	url:=fmt.Sprintf(accountBalance,countryCode,accountId)
	return J.getBaseUrl() + url

}

func (J *Jenga) getJengaMerchantsUrl() string {
	return J.getBaseUrl() + jengaMerchantsUrl
}
func (J *Jenga) getBankToMobileWalletUrl() string {

	return J.getBaseUrl() + jengaBankToMobileWalletUrl
}
func (J *Jenga) getPesaLinkToBankUrl() string {

	return J.getBaseUrl() + pesaLinkToBankUrl
}

func (J *Jenga) getEquityToEquityUrl()string {

	return J.getBaseUrl() + equityToequity

}

func (J *Jenga) getAirTimeUrl() string {
	return J.getBaseUrl() + jengaAirTimeUrl
}
func (J *Jenga) getAccessTokenUrl() string {

	return J.getBaseUrl() + jengaTokenUrl
}

func (J *Jenga) getKycUrl() string {
	return J.getBaseUrl() + jengaKycUrl

}
