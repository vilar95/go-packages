package accounts

import "cursos/gobank/clients"

type CurrentAccount struct {
	Holder        clients.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

// Funcão para sacar um valor na conta
func (c *CurrentAccount) AmountToWithdraw(withdrawalAmount float64) string {
	canWithdraw := withdrawalAmount <= c.balance && withdrawalAmount > 0
	if canWithdraw {
		c.balance -= withdrawalAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}

}

// Função depositar um valor na conta
func (c *CurrentAccount) DepositAmount(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Depósito realizado com sucesso", c.balance
	} else {
		return "Valor de depósito inválido", c.balance
	}
}

// Função de transferir valor entre contas
func (c *CurrentAccount) TransferAmount(transferAmount float64, d *CurrentAccount) string {
	if transferAmount > 0 && transferAmount <= c.balance {
		c.balance -= transferAmount
		d.balance += transferAmount
		return "Transferência realizada com sucesso"
	} else {
		return "Saldo insuficiente para transferência"
	}
}

// Função para consultar o saldo da conta
func (c *CurrentAccount) GetBalance() (string, float64) {
	return "O saldo do titular " + c.Holder.Name + " é de:", c.balance
}

// Função de pagar conta
func (c *CurrentAccount) Pay(billAmount float64) string {
	canPay := billAmount <= c.balance && billAmount > 0
	if canPay {
		c.balance -= billAmount
		return "Conta paga com sucesso"
	} else {
		return "Saldo insuficiente para pagar a conta"
	}
}
