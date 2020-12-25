package yuri



const MpesaLiveUrl = "https://api.safaricom.co.ke/"
const MpesaSandboxUrl = "https://sandbox.safaricom.co.ke/"

const tokenUrl="oauth/v1/generate?grant_type=client_credentials"
const b2cUrl="mpesa/b2c/v1/paymentrequest"
const b2bUrl="mpesa/b2b/v1/paymentrequest"
const balance="mpesa/accountbalance/v1/query"


const SandBox string = "sandbox"
const BusinessPayBill string ="BusinessPayBill"
const TransactionReversal string ="TransactionReversal"
const SalaryPayment string ="SalaryPayment"
const BusinessPayment string ="BusinessPayment"
const PromotionPayment string ="PromotionPayment"
const AccountBalance string ="AccountBalance"
const CustomerPayBillOnline string ="CustomerPayBillOnline"
const TransactionStatusQuery string ="TransactionStatusQuery"
const BusinessBuyGoods string ="BusinessBuyGoods"


//Production environment
const Production string = "production"
func (m *Mpesa) getAccessTokenUrl() string {
	if !m.Live {

		return MpesaSandboxUrl+tokenUrl
	}

	return MpesaLiveUrl+tokenUrl
}
func (m *Mpesa)getB2CUrl() string {
	if !m.Live {

		return MpesaSandboxUrl+b2cUrl
	}

	return MpesaLiveUrl+b2cUrl

}

func (m *Mpesa)getB2BUrl() string {
	if !m.Live {

		return MpesaSandboxUrl+b2bUrl
	}

	return MpesaLiveUrl+b2bUrl

}

func (m *Mpesa)getBalanceUrl() string {
	if !m.Live {

		return MpesaSandboxUrl+balance
	}

	return MpesaLiveUrl+balance

}