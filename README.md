# slack-weather-bot
A simple Slack Weather Bot written in golang

##### overview
it's been really hot in Los Angeles recently, and i realised i was switching to my web browser from Slack to find out the current temperature downstairs a lot before leaving the office. i realised that a `/weather` command would far more efficient. 

##### example run
assuming slash command `/weather` has been setup in your Slack and you have a running instance of the app available in Google Compute Engine, you should see something like this:

![ScreenShot](http://i.imgur.com/7etzvIx.png)

##### what's happening here
1. Slack slash command `/weather` is configured to GET a specific URL. in my case i use http://myweatherboturl.appspot.com/weather?zip=90071,US because i'm interested in downtown Los Angeles weather
2. slackweatherbot takes the zip parameter and assuming a valid openweather.org API key and a valid zip code, gets the current conditions for that zip, formats it and sends it to you on Slack 

##### notes
1. i put all my Slack bots up in Google App Engine, so this code is slightly different than if you want to self-host. 
2. this is just a v0.1/skeleton.
