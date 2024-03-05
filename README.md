# langchaingo-llm-cloudflare
Langchaingo extension to use Cloudflare Workers AI as an LLM

## Usage

### Initialize LLM:
```go
    // Create a new LLM client
    llm, err := cloudflare.New(
        cloudflare.WithAccountID("<account-id>"),
        cloudflare.WithServerURL("https://api.cloudflare.com/client/v4"),
        cloudflare.WithToken("<token>"),
        cloudflare.WithModel("<model>"),
        cloudflare.WithEmbeddingModel("<embedding-model>"),
        cloudflare.WithSystemPrompt("<system-prompt>"),
    )
    if err != nil {
        panic(err)
    }
	
```

### Generate Vector Embeddings
```go
    // Create a new text splitter
    splitter := textsplitter.NewRecursiveCharacter()
    
    splitter.ChunkOverlap = 200
    splitter.ChunkSize = 1000
    
    // Split the prompt into chunks
    promptSplit, err := splitter.SplitText("<prompt>")
    if err != nil {
        panic(err)
    }

    // Create documents from the prompt chunks
    docs, err := textsplitter.CreateDocuments(splitter, promptSplit, nil)
    if err != nil {
        panic(err)
    }
    
    d := []string{}
    for i := range docs {
        d = append(d, doc[i].PageContent)
    }

    // Initialize embedder
    embedder, err := embeddings.NewEmbedder(llm)
    if err != nil {
        panic(err)
    }
	
    // Fetch embeddings for the documents
    embs, err := embedder.EmbedDocuments(context.Background(), d)
    if err != nil {
        panic(err)
    }
	
```

### Generate text completions:
```go
    mc := []llms.MessageContent{
        llms.TextParts(schema.ChatMessageTypeAI, "<assistant-prompt>"),
    }
    
    for i := range promptSplit {
        mc = append(mc, llms.TextParts(schema.ChatMessageTypeHuman, promptSplit[i]))
    }

    // Create a new completion request
    res, err := llm.GenerateContent(context.Background(), mc)
    if err != nil {
        panic(err)
    }
```