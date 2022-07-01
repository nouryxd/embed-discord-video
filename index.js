const express = require('express')
const app = express()

// Change to whatever port you want to run on.
const PORT = 7532;

// The URL it runs on 
const INSTANCE_URL = "https://dc.noury.ee"

// Homepage with a short introduction of the tool and how to use it.
app.get('/', function (_req, res) {
    res.send(`
    <html>
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <title>Embed Discord Video</title>
        <meta charset="UTF-8" />
        <meta name="description" content="" />
    </head>
    <body style="background-color:#181a1b;">
    <h1 style="color:#d8d4cf">Discord Video Embedder</h1>
    <p style="color:#d8d4cf">Lets you watch a discord video in the browser instead of downloading it by default.</p>
    <p style="color:#d8d4cf">Simply paste the link to a discord video at the end of this url(after the /, if there is none add one).</p>
    <p style="color:#d8d4cf">Like this: <a style="color:#3391ff" href="${INSTANCE_URL}/https://cdn.discordapp.com/attachments/381520882608373761/989666371178754068/denkcats_1639474686233272.mp4">${INSTANCE_URL}/https://cdn.discordapp.com/attachments/381520882608373761/989666371178754068/denkcats_1639474686233272.mp4</a></p>
    <p style="color:#d8d4cf">Source: <a style="color:#3391ff" href="https://github.com/noury-ee/embed-discord-video">https://github.com/noury-ee/embed-discord-video</a></p>
    <p style="color:#d8d4cf"><sup>this is beautiful web design shut up</sup></p>
    </body>
    </html>
    `)
})

// Listen on anything after the initial slash.
app.get('/*', function (req, res) {
    
    // url here contains everything after the initial slash and will be
    // embedded as a video on this page.
    const url = req.params[0];
  
    res.send(`
    <html>
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <title>Discord Video Embedder</title>
        <meta charset="UTF-8" />
        <meta name="description" content="" />
    </head>
    <body style="background-color:#181a1b;">
    <p>
    <video
        controls
        autoplay
        width="500px"
        src="${url}"
    />
    </p>
    <p style="color:#d8d4cf">Original url: <a style="color:#3391ff" href="${url}">${url}</a></p>
    </body>
    </html>
    `)
})

console.log(`Listening on port: ${PORT}`)
app.listen(PORT)