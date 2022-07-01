# Embed-Disord-Video

Lets you watch a discord video in the browser instead of downloading it by default.

Simply paste the link to a discord video at the end of this url(after the /, if there is none add one).

Make sure you configure the `PORT` and `INSTANCE_URL` in the `index.js` file for your use.

Install:
```sh
$ npm install
```

Run:
```sh
$ npm start
```

Deploy with pm2:
```sh
$ pm2 start "npm start" --name=dc
```