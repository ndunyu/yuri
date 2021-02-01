package yuri

const jengaTokenUrl = "identity/v2/token"
const jengaKycUrl = "customer/v2/identity/verify"
const jengaAirTimeUrl = "transaction/v2/airtime"
const jengaMerchantsUrl = "transaction/v2/merchants"
const jengaBankToMobileWalletUrl = "transaction/v2/remittance#sendmobile"
const pesaLinkToBankUrl = "transaction/v2/remittance'"

const JengaLiveUrl = "https://api.jengahq.io/"
const JengaSandboxUrl = "https://uat.jengahq.io/"

func (J *Jenga) getBaseUrl() string {
	if !J.Live {

		return JengaSandboxUrl
	}

	return JengaLiveUrl

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

func (J *Jenga) getAirTimeUrl() string {
	return J.getBaseUrl() + jengaAirTimeUrl
}
func (J *Jenga) getAccessTokenUrl() string {

	return J.getBaseUrl() + jengaTokenUrl
}

func (J *Jenga) getKycUrl() string {
	return J.getBaseUrl() + jengaKycUrl

}
