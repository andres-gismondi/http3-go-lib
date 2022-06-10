package http3_go_lib

import "form3/model"

type AccountOption func(data *model.AccountData)

func BaseCurrency(code string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.BaseCurrency = code
	}
}

func BankID(id string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.BankID = id
	}
}

func BankIDCode(code string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.BankIDCode = code
	}
}

func AccountNumber(number string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.AccountNumber = number
	}
}

func BIC(number string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.Bic = number
	}
}

func Join(val bool) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.JointAccount = val
	}
}

func Classification(name string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.AccountClassification = name
	}
}

func IBAN(number string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.Iban = number
	}
}

func SecondaryID(number string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.SecondaryIdentification = number
	}
}

func AlternativeNames(names []string) AccountOption {
	return func(data *model.AccountData) {
		data.Attributes.AlternativeNames = names
	}
}
