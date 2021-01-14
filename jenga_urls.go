package yuri

const jengaTokenUrl = "identity/v2/token"
const jengaKycUrl = "customer/v2/identity/verify"
const JengaLiveUrl = "https://uat.jengahq.io/"
const JengaSandboxUrl = "https://uat.jengahq.io/"
func (J *Jenga) getBaseUrl() string {
	if !J.Live {

		return JengaSandboxUrl
	}

	return JengaLiveUrl

}

func (J *Jenga) getAccessTokenUrl() string {

	return J.getBaseUrl() +jengaTokenUrl
}

func (J *Jenga)getKycUrl()string{
	return  J.getBaseUrl()+jengaKycUrl

}