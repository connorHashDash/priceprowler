# Price Prowler
See if your area is gaining or losing property value.

There's a server side element to this, getHousePrices.sh is on a VPS, runs 
monthly via a CRON job, writes to MariaDB.

You'll need to go to the UK gov website and generate your own URL for your local area and get the params you want.

The go code runs locally, speaks to the db, and generates a report on local 
house prices going up or down, perhaps some other info, we'll see.
