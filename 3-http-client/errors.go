package cloudpayments

type CPError struct {
	Code           int
	VerboseReason  string
	VerboseMessage string
}

func NewError(code int) *CPError {
	return getErrorByCode(code)
}

func (e *CPError) Error() string {
	return e.VerboseMessage
}

var errorMessages = []CPError{
	{5001, "ReferToCardIssuer", "ReferToCardIssuer"},
	{5003, "InvalidMerchant", "InvalidMerchant"},
	{5004, "PickUpCard", "PickUpCard"},
	{5005, "DoNotHonor", "DoNotHonor"},
	{5006, "Error", "Error"},
	{5007, "PickUpCardSpecialConditions", "PickUpCardSpecialConditions"},
	{5012, "InvalidTransaction", "InvalidTransaction"},
	{5013, "AmountError", "AmountError"},
	{5014, "InvalidCardNumber", "InvalidCardNumber"},
	{5015, "NoSuchIssuer", "NoSuchIssuer"},
	{5019, "TransactionError", "TransactionError"},
	{5030, "FormatError", "FormatError"},
	{5031, "BankNotSupportedBySwitch", "BankNotSupportedBySwitch"},
	{5033, "ExpiredCardPickup", "ExpiredCardPickup"},
	{5034, "SuspectedFraud", "SuspectedFraud"},
	{5041, "LostCard", "LostCard"},
	{5043, "StolenCard", "StolenCard"},
	{5051, "InsufficientFunds", "InsufficientFunds"},
	{5054, "ExpiredCard", "ExpiredCard"},
	{5057, "TransactionNotPermitted", "TransactionNotPermitted"},
	{5059, "SuspectedFraudDecline", "SuspectedFraudDecline"},
	{5062, "RestrictedCard2", "RestrictedCard2"},
	{5063, "SecurityViolation", "SecurityViolation"},
	{5065, "ExceedWithdrawalFrequency", "ExceedWithdrawalFrequency"},
	{5082, "IncorrectCVV", "IncorrectCVV"},
	{5091, "Timeout", "Timeout"},
	{5092, "CannotReachNetwork", "CannotReachNetwork"},
	{5096, "SystemError", "SystemError"},
	{5204, "UnableToProcess", "UnableToProcess"},
	{5206, "AuthenticationFailed", "AuthenticationFailed"},
	{5207, "AuthenticationUnavailable", "AuthenticationUnavailable"},
	{5300, "AntiFraud", "AntiFraud"},
}

func getErrorByCode(code int) *CPError {
	for _, e := range errorMessages {
		if e.Code == code {
			return &e
		}
	}
	return nil
}
