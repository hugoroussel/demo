package main

func main() {
	Blockchain := NewBlockchain()

	Blockchain.AddBlock("Send 1 coin to Patrick")
	Blockchain.AddBlock("Send 2 coins to Martin")
}
