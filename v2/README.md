## weather-by-geo

### How to run

1. Copy `.env.example` to `.env`.
2. Set `TELEGRAM_IS_WEBHOOK` to `true` or `false`
3. Create tl bot via botfather.
4. Paste your tl bot token to `TELEGRAM_TOKEN` env variable.
5. Run `ngrok http 8080` (taken port from `compose.yml`, `app.ports` section).
6. Paste your ngrok host, port (443) to `APP_HOST`, `APP_PORT` corresponding env variables.
7. Run `docker compose up -d`.
8. Run `docker logs -f weather_app_v2`: you might see the fololowing output:
```
2026/01/01 00:00:00 Authorized on account yout_bot
2026/01/01 00:00:00 Forecast api url: https://api.open-meteo.com/v1/forecast?current_weather=true&latitude=33.4126&longitude=-88.4819&timezone=auto
```

#### That's It, you have tl bot with ngrok reverse proxy from tls www to your localhost.