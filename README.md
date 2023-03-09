# LineGPT
---

LineGPT is a Line Bot that responds with OpenAI model

## Installation

Use Docker

```bash
    docker pull wpted/line-gpt
```

## Usage

```bash
    docker run --rm -d -p 8000:8000 wpted/line-gpt
```

Connect the local port to Line Messaging bot by creating a temporary connection using ngrok.

```bash
    ngrok http 8000
```

Go to Line Developer Console, add the ngrok url(with endpoint /callback) to the webhook field.
The url will look something like this

```text
    https://8d3a-211-23-199-105.jp.ngrok.io/callback
```

Verify the url to see the service is ready.


### Or
The project is deployed on Azure Container Registry and Azure App service, try by scanning the below QRCode

![](LineGPT.png)

Add the bot as friend, then, voilà!






