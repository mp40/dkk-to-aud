### dkk-to-aud
[Deloyed on Heroku](https://dkk-to-aud.herokuapp.com/)

Application to track the exchange rate of DKK to AUD over the time I am in Denmark.
The purpose to help make decisions on best time to send money back to Australia.

#### Currently at MVP
Current Implementation Details
* Gets historical data from a google sheet as APIs with historical date range queries require paid subscription
* Displays data on static page (Highest, Lowest, Average rates)

#### Possible Areas Of Improvement
* Refactoring and improved tests as I get a better feel for Go
* Add Median and Mode to data shown
* Containerise with Docker
* Add Redis to cache results
* Receive updates on exchange rate via email
