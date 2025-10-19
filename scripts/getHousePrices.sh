#!/bin/bash
set -euo pipefail

POSTCODE=${1}
DB_USER="con"
DB_PASS="Scoibert212."
DB_NAME="priceProwler"
DB_HOST="localHost"

# Dates: previous month
MIN_DATE=$(date -d "$(date +%Y-%m-01) -24 month" +%Y-%m-01)
MAX_DATE=$(date -d "$(date +%Y-%m-01) -1 day" +%Y-%m-%d)

URL="https://landregistry.data.gov.uk/app/ppd/ppd_data.csv?et%5B%5D=lrcommon%3Afreehold&et%5B%5D=lrcommon%3Aleasehold&limit=all&max_date=${MAX_DATE}&min_date=${MIN_DATE}&nb%5B%5D=true&nb%5B%5D=false&postcode=${POSTCODE}&ptype%5B%5D=lrcommon%3Adetached&ptype%5B%5D=lrcommon%3Asemi-detached&ptype%5B%5D=lrcommon%3Aterraced&ptype%5B%5D=lrcommon%3Aflat-maisonette&ptype%5B%5D=lrcommon%3AotherPropertyType&relative_url_root=%2Fapp%2Fppd&tc%5B%5D=ppd%3AstandardPricePaidTransaction&tc%5B%5D=ppd%3AadditionalPricePaidTransaction"

TMP_FILE="./ppd_last_month.csv"

echo "Downloading PPD data for $MIN_DATE to $MAX_DATE..."
curl -k -v -s -o "$TMP_FILE" "$URL"

echo "writing to database"
mysql --local-infile=1 -u "$DB_USER" -p"$DB_PASS" -h "$DB_HOST" "$DB_NAME" <<EOF
TRUNCATE TABLE house_sales;
EOF

echo "writing to database"
mysql --local-infile=1 -u "$DB_USER" -p"$DB_PASS" -h "$DB_HOST" "$DB_NAME" <<EOF
LOAD DATA LOCAL INFILE '$(pwd)/ppd_last_month.csv'
INTO TABLE house_sales
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(transaction_id, price, transfer_date, postcode, property_type,
 @old_new, tenure, paon, @saon, street, @locality, @town, @district, @county, @category, record_status, @uri)
SET new_build = IF(@old_new = 'Y', 1, 0);
EOF
