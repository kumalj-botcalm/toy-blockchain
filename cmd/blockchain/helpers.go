func loadChain() (*blockchain.Blockchain, error)
chain, err := storage.Load(blockchain.DefaultChainFile)

func saveChain(chain *blockchain.Blockchain) error
return storage.Save(
    blockchain.DefaultChainFile,
    chain,
)