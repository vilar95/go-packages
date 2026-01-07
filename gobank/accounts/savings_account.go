package accounts

import "cursos/gobank/clients"

type SavingsAccount struct {
	Holder        clients.Holder
	AgencyNumber  int
	AccountNumber int
	operation     int
	balance       float64
}

// Funcão para sacar um valor na conta
func (s *SavingsAccount) AmountToWithdraw(withdrawalAmount float64) string {
	canWithdraw := withdrawalAmount <= s.balance && withdrawalAmount > 0
	if canWithdraw {
		s.balance -= withdrawalAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}

}

// Função depositar um valor na conta
func (s *SavingsAccount) DepositAmount(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		s.balance += depositAmount
		return "Depósito realizado com sucesso", s.balance
	} else {
		return "Valor de depósito inválido", s.balance
	}
}

// Função de transferir valor entre contas
func (s *SavingsAccount) TransferAmount(transferAmount float64, d *SavingsAccount) string {
	if transferAmount > 0 && transferAmount <= s.balance {
		s.balance -= transferAmount
		d.balance += transferAmount
		return "Transferência realizada com sucesso"
	} else {
		return "Saldo insuficiente para transferência"
	}
}

// Função para consultar o saldo da conta
func (s *SavingsAccount) GetBalance() (string, float64) {
	return "O saldo do titular " + s.Holder.Name + " é de:", s.balance
}

// Função de pagar conta
func (s *SavingsAccount) Pay(billAmount float64) string {
	canPay := billAmount <= s.balance && billAmount > 0
	if canPay {
		s.balance -= billAmount
		return "Conta paga com sucesso"
	} else {
		return "Saldo insuficiente para pagar a conta"
	}
}
