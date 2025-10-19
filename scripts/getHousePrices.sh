#!/bin/bash
set -euo pipefail

if [-z "${1:-}"]; then
  echo "Usage: $0 <postcode>"
  exit 1
fi

POSTCODE="${1}"
DB_USER=""
DB_PASS=""
DB_NAME=""
DB_HOST=""

# Dates: previous month
MIN_DATE=$(date -d "$(date +%Y-%m-01) -1 month" +%Y-%m-01)
MAX_DATE=$(date -d "$(date +%Y-%m-01) -1 day" +%Y-%m-%d)

# for example: https://landregistry.data.gov.uk/app/ppd/
URL="You can get this from UK gov land registry"

TMP_FILE="./ppd_last_month.csv"

echo "Downloading PPD data for $MIN_DATE to $MAX_DATE..."
curl -k -v -s -o "$TMP_FILE" "$URL"

echo "writing to database"
mysql --local-infile=1 -u "$DB_USER" -p"$DB_PASS" -h "$DB_HOST" "$DB_NAME" <<EOF
LOAD DATA LOCAL INFILE '$(pwd)/ppd_last_month.csv'
INTO TABLE house_sales
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
LINES TERMINATED BY '\n'
(transaction_id, price, transfer_date, postcode, property_type,
 new_build, tenure, paon, street, record_status);
EOF
