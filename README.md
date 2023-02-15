# serverless-gpt

A simple API endpoint that anwsers your prompt using OpenAI's GPT-3

## Prerequisites

- You need to create an API key from OpenAI and set it as the value of OPENAI_API_TOKEN environment variable.
- You need to set up credentials for AWS. Follow [this tutorial](https://www.serverless.com/framework/docs/providers/aws/guide/credentials/) to know how to configure it.
  
## Deployment

```
make deploy
```

## Using

```bash
curl -X POST <your-api-gateway>/ask \
  -H "Content-Type: application/json" \
  -d '{"prompt": "say this is a test"}'

# Response
{"answer": "This is a test."}
```
