#!/bin/sh
# Generate config.json from environment variables

cat <<EOF > /usr/share/nginx/html/config.json
{
  "API_BASE_URL": "${API_BASE_URL}",
  "APP_TITLE": "${APP_TITLE}"
}
EOF

exec nginx -g 'daemon off;'
