package yuri

const MpesaLiveUrl = "https://api.safaricom.co.ke/"
const MpesaSandboxUrl = "https://sandbox.safaricom.co.ke/"

const tokenUrl = "oauth/v1/generate?grant_type=client_credentials"
const b2cUrl = "mpesa/b2c/v1/paymentrequest"
const b2bUrl = "mpesa/b2b/v1/paymentrequest"
const balance = "mpesa/accountbalance/v1/query"
const transactionStatus = "mpesa/transactionstatus/v1/query"
const registerUrl="mpesa/c2b/v1/registerurl"
const SandBox string = "sandbox"
const BusinessPayBill string = "BusinessPayBill"
const TransactionReversal string = "TransactionReversal"
const SalaryPayment string = "SalaryPayment"
const BusinessPayment string = "BusinessPayment"
const PromotionPayment string = "PromotionPayment"
const AccountBalance string = "AccountBalance"
const CustomerPayBillOnline string = "CustomerPayBillOnline"
const TransactionStatusQuery string = "TransactionStatusQuery"
const BusinessBuyGoods string = "BusinessBuyGoods"

//Production environment
const Production string = "production"

func (m *Mpesa) getAccessTokenUrl() string {

	return m.getBaseUrl() + tokenUrl
}
func (m *Mpesa) getB2CUrl() string {

	return m.getBaseUrl() + b2cUrl

}

func (m *Mpesa) getB2BUrl() string {

	return m.getBaseUrl() + b2bUrl

}

func (m *Mpesa) getBalanceUrl() string {

	return m.getBaseUrl() + balance

}

func (m *Mpesa) getTransactionStatusUrl() string {

	return m.getBaseUrl() + transactionStatus

}

func (m *Mpesa)getC2BRegisterUrl() string {


	return m.getBaseUrl() + registerUrl


}

func (m *Mpesa) getBaseUrl() string {
	if !m.Live {

		return MpesaSandboxUrl
	}

	return MpesaLiveUrl

}
