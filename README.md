### dkk-to-aud
[MVP Deloyed on Heroku](https://dkk-to-aud.herokuapp.com/)

Application to track the exchange rate of DKK to AUD over the time I am in Denmark.
The purpose to help make decisions on best time to send money back to Australia.

#### Currently at Version Two
Current Implementation Details
* Gets historical data from a google sheet as APIs with historical date range queries require paid subscription
* Displays data on static page (Highest, Lowest, Latest, Average, Median rates)
* Builds app with two stage Docker build process.
* Uses Redis to cache calculations to reduce calls for data.

#### Possible Areas Of Improvement
* Refactoring and improved tests as I get a better feel for Go
* Add Mode to data shown
* Receive updates on exchange rate via email

#### Deployment
Current Heroku deployed version is the un-dockerized MVP without Redis caching, latest and median rates.
