# 🦜⛓️‍💥 Simplified LangChain for Go

Pipe (Chain) uses Input, Model and Output to make the interaction with the model easier.

## Concepts

### Core

Input -> Model -> Output

### Agents

LLM -> Agent
Tools -> Agent
Agent -> Executor (Iterator)

(Agents: Static Validator/Limits/Logger)

## Core

- [x] Input: Chat messages and prompt
- [x] Models: Ollama and OpenAI
- [x] Parser: String, Json, Markdown, Symbol-Seperator and Map
- [x] Pipe: Simple pipe to easily use input, models and output.
- [x] Embedder: For vector embeddings
- [x] Evaluator: Whitelist or Blacklist output before parsing (Soon in pipe and agent available)

## Agents

- [x] Basic Agent using the ReAct Pattern
- [x] Static Tools use (Probably add input instruction)
- [x] Agent Executer
- [ ] Limits and Logging of Agents and Executor
- [ ] Static Agent validator

## RAG Retrival Argument Generation

- [x] Similarity Search (Jaccard and Cosine)
- [x] Vector Database (Qdrant)

## Extensions

- [ ] Tools (.txt, pdf, ...)

## Examples

- Core (Input, Model, Output, Pipe, Embedder)
- Temperature and time Agents
- RAG: Simple RAG with Jaccard
- Vectorstore: Qdrant
