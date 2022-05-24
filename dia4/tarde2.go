package dia4

// Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
// produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
// mesmo endereço de memória no main do programa e nas funções.

// Estruturas necessárias:
// - Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
// - Produto: Nome, preço, quantidade.
// Algumas funções necessárias:
// - Novo produto: recebe nome e preço, e retorna um produto.
// - Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
// o produto ao usuário.
// - Deletar produtos: recebe um usuário, apaga os produtos do usuário.

type Usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []Produto
}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func NovoProduto(nome string, preco float64) Produto {
	return Produto{Nome: nome, Preco: preco}
}

func AdicionarProduto(u *Usuario, p Produto, qty int) {
	p.Quantidade = qty
	u.Produtos = append(u.Produtos, p)
}

func DeletarProdutos(u *Usuario) {
	u.Produtos = nil
}
