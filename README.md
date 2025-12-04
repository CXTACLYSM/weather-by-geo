## weather-by-geo

### How to run

1. Copy `.env.example` to `.env`.
2. Create tl bot via botfather.
2. Paste your tl bot token to `TELEGRAM_TOKEN` env variable.
3. Run `ngrok http 8080` (taken port from `compose.yml`, `app.ports` section).
4. Paste your ngrok host, port (443) to `APP_HOST`, `APP_PORT` corresponding env variables.
3. Run `docker compose up -d`.
4. Run `docker logs -f weather_app`: you might see the fololowing output:
```
2026/01/01 00:00:00 Webhook URL: https://abc123.ngrok-free.app:443
2026/01/01 00:00:00 Successfully posted webhook URL
2026/01/01 00:00:00 Starting server on port 8080
2026/01/01 00:00:00 Received:
{
  ...
}
```

#### That's It, you have tl bot with ngrok reverse proxy from tls www to your localhost.