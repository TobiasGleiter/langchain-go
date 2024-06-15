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

- [x] Input chat messages and prompt
- [x] Ollama and OpenAI Models
- [x] String and Json output parsers
- [x] Simple pipe to easily use input, models and output.

## Agents

- [x] Basic Agent using the ReAct Pattern
- [x] Static Tools use (Probably add input instruction)
- [x] Agent Executer
- [ ] Limits and Logging of Agents and Executor
- [ ] Static Agent validator

## Extensions

- [ ] Tools (.txt, pdf, ...)
- [ ] Vector Database and Embedding

## Examples

- Core (Input, Model, Output, Pipe)
- Temperature and time Agents
