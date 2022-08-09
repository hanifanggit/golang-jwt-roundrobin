# golang-jwt-roundrobin

How to config Round Robin with Nginx !


Go to Nginx/conf/sites-available/localhost.conf


Copy this code

upstream localhost {
  server localhost:8081;
  server localhost:8082;
  server localhost:8083;
}

server {
  listen          8080;
  server_name     localhost;
  access_log      "C:/Program Files/nginx-1.23.0/logs/localhost.access.log" combined;

  location / {
    proxy_pass      http://localhost;
  }
}
