package dia4

// Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
// que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
// que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
// para as funções:
// - A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
// senha
// E devem implementar as seguintes funcionalidades:
// - mudar o nome: me permite mudar o nome e sobrenome
// - mudar a idade: me permite mudar a idade
// - mudar e-mail: me permite mudar o e-mail
// - mudar a senha: me permite mudar a senha

type User struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     string
}

func (u *User) MudarNome(novoNome string) {
	u.Nome = novoNome
}

func (u *User) MudarSobrenome(novoSobrenome string) {
	u.Sobrenome = novoSobrenome
}

func (u *User) MudarIdade(novaIdade int) {
	u.Idade = novaIdade
}

func (u *User) MudarEmail(novoEmail string) {
	u.Email = novoEmail
}

func (u *User) MudarSenha(novaSenha string) {
	u.Senha = novaSenha
}
