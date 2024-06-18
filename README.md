# filmigobot

**filmigobot** is a fully serverless high-performace inline telegram bot to search different movie databases using the [filmigo library](https://github.com/Jisin0/filmigo) written in [GO](https://go.dev). It is designed to be easily deplyed to Vercel but has support almost any other servers. It currently supports IMDb, JustWatch and OMDb.

## Variables

- [ ] `BOT_TOKEN`  : Optional. On vercel, a list of bot tokens allowed to connect to the app or leave empty allow anyone to connect. On servers, a single bot token.

## Deploy
<details><summary>Deploy To Heroku</summary>
<p>
<br>
<a href="https://heroku.com/deploy?template=https://github.com/Jisin0/filmigobot/tree/main">
  <img src="https://www.herokucdn.com/deploy/button.svg" alt="Deploy">
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
You must have the latest version of <a href="https://go.dev/dl">go</a> installed first
<pre>
git clone https://github.com/Jisin0/filmigobot
cd filmigobot
go build .
./filmigobot
</pre>
</p>
</details>

## Support

Ask any doubts or help in our support chat.
[![telegram badge](https://img.shields.io/badge/Telegram-Group-30302f?style=flat&logo=telegram)](https://telegram.dog/jisin_hub)

Join our telegram channel for more latest news and cool projects
[![telegram badge](https://img.shields.io/badge/Telegram-Channel-30302f?style=flat&logo=telegram)](https://telegram.dog/jisin_0)

## Thanks

 - Thanks to Paul for his awesome [Library](https://github.com/PaulSonOfLars/gotgbot)
 - Thanks To [S](https://github.com/trojanzhex) for Their Awesome [Unlimited Filter Bot](https://github.com/TroJanzHEX/Unlimited-Filter-Bot)

## Disclaimer
[![GNU Affero General Public License 2.0](https://www.gnu.org/graphics/agplv3-155x51.png)](https://www.gnu.org/licenses/agpl-3.0.en.html#header)    
Licensed under [GNU AGPL 2.0.](https://github.com/Jisin0/evamaria/blob/master/LICENSE)
Selling The Codes To Other People For Money Is *Strictly Prohibited*.