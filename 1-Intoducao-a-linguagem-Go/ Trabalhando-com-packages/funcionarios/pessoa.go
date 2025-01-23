package funcionarios

type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func (p *Pessoa) addSalaryPessoa(bonus int) {
	p.salario += bonus
}
