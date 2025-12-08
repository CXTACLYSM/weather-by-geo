## weather-by-geo

### How to run

1. Copy `.env.example` to `.env`.
2. Create tl bot via botfather.
2. Paste your tl bot token to `TELEGRAM_TOKEN` env variable.
3. Run `ngrok http 8080` (taken port from `compose.yml`, `app.ports` section).
4. Paste your ngrok host, port (443) to `APP_HOST`, `APP_PORT` corresponding env variables.
3. Run `docker compose up -d`.
4. Run `docker logs -f weather_app_v1`: you might see the fololowing output:
```
2026/01/01 00:00:00 Webhook URL: https://abc123.ngrok-free.app:443
2026/01/01 00:00:00 Successfully posted webhook URL
2026/01/01 00:00:00 Starting server on port 8080
2026/01/01 00:00:00 Received:
{
  ...
}
2026/01/01 00:00:00 Forecast api url: https://api.open-meteo.com/v1/forecast?current_weather=true&latitude=33.4126&longitude=-88.4819&timezone=auto
```

#### That's It, you have tl bot with ngrok reverse proxy from tls www to your localhost.