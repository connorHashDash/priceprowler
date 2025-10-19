# Price Prowler

Thinking about buying in a new area? Want to have more info on your investment?

Price prowler might be what you're looking for.

There's a server side element to this, scripts/getHousePrices.sh is on a VPS, with MariaDB.
If you pass ```--postcode=<your postcode>``` flag to the function it will pass it to the server, downloading the CSV and populating the DB table.

The go code runs locally, speaks to the db, and generates a report on local 
house prices going up or down, perhaps some other info, we'll see.
