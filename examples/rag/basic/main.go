// Inspired by https://learnbybuilding.ai/tutorials/rag-from-scratch

package main

import (
	"context"
	"fmt"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
	"github.com/TobiasGleiter/langchain-go/core/output"
	"github.com/TobiasGleiter/langchain-go/core/pipe"
	"github.com/TobiasGleiter/langchain-go/rag/similarity"
)

func main() {
	corpus_of_documents := []string{
		"Take a leisurely walk in the park and enjoy the fresh air.",
		"Visit a local museum and discover something new.",
		"Attend a live music concert and feel the rhythm.",
		"Go for a hike and admire the natural scenery.",
		"Have a picnic with friends and share some laughs.",
		"Explore a new cuisine by dining at an ethnic restaurant.",
		"Take a yoga class and stretch your body and mind.",
		"Join a local sports league and enjoy some friendly competition.",
		"Attend a workshop or lecture on a topic you're interested in.",
		"Visit an amusement park and ride the roller coasters.",
	}

	// 1. User input
	// 2. Similarity (e.g. Jaccard)
	// 3. Pass best result of similarity to LLM
	// 4. Pass Output to user

	userPrompt := "I don't like to hike"
	var jaccardSimilarity = &similarity.Jaccard{}
	relevantContext := selectOneSimilarDocument(jaccardSimilarity, userPrompt, corpus_of_documents) // Complexity of O(n)

	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
		{Role: "user", Content: `You are a bot that makes recommendations for activities. You answer in very short sentences and do not include extra information.
			This is the recommended activity: {{.relevant_document}}
			The user input is: {{.input}}
			Output a recommondation for the activity. Output full text recommondation.
		`},
	})

	data := map[string]interface{}{
		"relevant_document": relevantContext,
		"input":             userPrompt,
	}

	messages, err := chatPrompt.FormatMessages(data)
	if err != nil {
		panic(err)
	}

	llama3_8b_model := ollama.OllamaModel{
		Model:   "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream:  false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	var parser output.OutputParser[string] = &output.StringOutputParser{}

	ctx := context.TODO()
	pipe := pipe.NewPipe(messages, ollamaClient, parser)
	result, _ := pipe.Invoke(ctx)
	fmt.Println(result)
}

func selectOneSimilarDocument(similarity similarity.Similarity[string], query string, corpus []string) string {
	maxSimilarity := -1.0
	bestMatchIndex := -1
	for i, doc := range corpus {
		similarity, _ := similarity.Calculate(query, doc)
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			bestMatchIndex = i
		}
	}

	return corpus[bestMatchIndex]
}
