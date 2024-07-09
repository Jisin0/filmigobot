# filmigobot
[![telegram badge](https://img.shields.io/badge/Telegram-Channel-30302f?style=flat&logo=telegram)](https://telegram.dog/FractalProjects)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jisin0/filmigobot)](https://goreportcard.com/report/github.com/Jisin0/filmigobot)
[![Go Build](https://github.com/Jisin0/filmigobot/workflows/Go/badge.svg)](https://github.com/Jisin0/filmigobot/actions?query=workflow%3AGo+event%3Apush+branch%3Amain)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[**filmigobot** ](https://filmigobot.vercel.app) is a fully serverless high-performace inline telegram bot to search different movie databases using the [filmigo library](https://github.com/Jisin0/filmigo) written in [GO](https://go.dev). It is designed to be easily deployed to Vercel but has support almost any other servers. It currently supports IMDb, JustWatch and OMDb.

Connect a new bot to the [Public App](https://filmigobot.vercel.app) now or deploy a new instance following the instructions below.
[Sample Bot](https://telegram.dog/SurfOTTBot)

## Commands
```
/start : Check if the bot is alive.
/about: Basic Information About the bot.
/help: Short Guide on How to Use the Bot.
/imdb: Search or get a movie from IMDb.
/jw: Search or get a movie from JustWatch
```

## Variables

- `BOT_TOKEN`  : Optional. On vercel, a list of bot tokens allowed to connect to the app or leave empty allow anyone to connect. On servers, a single bot token.
- `DEFAULT_SEARCH_METHOD` : The default method to use for inline search. Possible values are jw, imdb & omdb.

## Deploy
Deploy your own **filmigobot** app to vercel

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/project?template=https://github.com/Jisin0/filmigobot/tree/main&env=BOT_TOKEN&envDescription=List%20of%20of%20allowed%20bot%20tokens%20or%20leave%20empty%20to%20allow%20all)

<details><summary>Deploy To Heroku</summary>
<p>
<br>
<a href="https://heroku.com/deploy?template=https://github.com/Jisin0/filmigobot/tree/main">
  <img src="https://www.herokucdn.com/deploy/button.svg" alt="Deploy">
</a>
</p>
</details>

<details><summary>Deploy To Scalingo</summary>
<p>
<br>
<a href="https://dashboard.scalingo.com/create/app?source=https://github.com/Jisin0/filmigobot#main">
   <img src="https://cdn.scalingo.com/deploy/button.svg" alt="Deploy on Scalingo" data-canonical-src="https://cdn.scalingo.com/deploy/button.svg" style="max-width:100%;">
</a>
</p>
</details>


<details><summary>Deploy To Render</summary>
<p>
<br>
<a href="https://dashboard.render.com/select-repo?type=web">
  <img src="https://render.com/images/deploy-to-render-button.svg" alt="deploy-to-render">
</a>
</p>
<p>
Make sure to have the following options set :

<b>Environment</b>
<pre>Go</pre>

<b>Build Command</b>
<pre>go build .</pre>

<b>Start Command</b>
<pre>./filmigobot</pre>

<b>Advanced >> Health Check Path</b>
<pre>/</pre>
</p>
</details>


<details><summary>Deploy To Koyeb</summary>
<p>
<br>
<a href="https://app.koyeb.com/deploy?type=git&repository=github.com/Jisin0/filmigobot&branch=main">
  <img src="https://www.koyeb.com/static/images/deploy/button.svg" alt="deploy-to-koyeb">
</a>
</p>
<p>
You must set the Run command to :
<pre>./bin/filmigobot</pre>
</p>
</details>

<details><summary>Deploy To Okteto</summary>
<p>
<br>
<a href="https://cloud.okteto.com/deploy?repository=https://github.com/Jisin0/filmigobot">
  <img src="https://okteto.com/develop-okteto.svg" alt="deploy-to-okteto">
</a>
</p>
</details>

<details><summary>Deploy To Railway</summary>
<p>
<br>
<a href="https://railway.app/new/template?template=https%3A%2F%2Fgithub.com%2FJisin0%2Ffilmigobot">
  <img src="https://railway.app/button.svg" alt="deploy-to-railway">
</a>
</p>
</details>

<details><summary>Run Locally/VPS</summary>
<p>
You must have the latest version of <a href="https://go.dev/dl">GO</a> installed first
<pre>
git clone https://github.com/Jisin0/filmigobot
cd filmigobot
go build .
./filmigobot
</pre>
</p>
</details>

## Thanks

 - Thanks to Paul for his awesome [Library](https://github.com/PaulSonOfLars/gotgbot)
 - Thanks To [SpEcHIDe](https://github.com/SpEcHIDe) for his awesome [IMDbOT](https://github.com/TelegramPlayGround/IMDbOT)

## Disclaimer
Any data obtained using the bot is not allowed to be used for commercial use. Please read the Privacy Policy of the respective movie database before using, sharing or modifying data from them.

[![GNU General Public License 3.0](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html#header)    
Licensed under [GNU GPL 3.0.](https://github.com/Jisin0/filmigobot/blob/main/LICENSE).