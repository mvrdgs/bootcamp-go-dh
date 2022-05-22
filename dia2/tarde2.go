package dia2

// - Criar uma interface “Produto” que possua o método “CalcularCusto”
type Produto interface {
	CalcularCusto() float64
}

// - Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
type Ecommerce interface {
	Total() float64
	Adicionar(p produto)
}

// - Criar uma estrutura “loja” que guarde uma lista de produtos.
type loja struct {
	produtos []produto
}

// - Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
// base no custo total dos produtos + o adicional citado anteriormente (caso a categoria tenha)
func (l loja) Total() float64 {
	total := 0.0

	for _, produto := range l.produtos {
		adicional := produto.CalcularCusto()

		total += produto.preco + adicional
	}

	return total
}

// - Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
// e adicioná-lo a lista da loja
func (l loja) Adicionar(p produto) {
	l.produtos = append(l.produtos, p)
}

// - Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
type produto struct {
	nome  string
	tipo  string
	preco float64
}

// - Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
// custo adicional segundo o tipo de produto.
func (p produto) CalcularCusto() float64 {
	// - Pequeno: O custo do produto (sem custo adicional)
	// - Médio: O custo do produto + 3% pela disponibilidade no estoque
	// - Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
	// adicional pelo envio de $2500.
	switch p.tipo {
	case "Medio":
		return p.preco * .03
	case "Grande":
		return p.preco*.06 + 2500
	default:
		return 0.0
	}
}

// - Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome e preço, e devolva um Produto.
func NovoProduto(nome, tipo string, preco float64) produto {
	return produto{
		nome,
		tipo,
		preco,
	}
}

// - Será necessário uma função “novaLoja” que retorne um Ecommerce.
func NovaLoja() Ecommerce {
	return loja{make([]produto, 0)}
}
