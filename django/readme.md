Создание пользователя
---------------------
adduser username
usermod -aG sudo username
group username
su username
---------------------------------------
Компиляции python 3.6
----------------------
sudo apt-get install -y make build-essential libssl-dev zlib1g-dev
sudo apt-get install -y libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm
sudo apt-get install -y libncurses5-dev  libncursesw5-dev xz-utils tk-dev

wget https://www.python.org/ftp/python/3.6.4/Python-3.6.4.tgz
tar xvf Python-3.6.4.tgz
cd Python-3.6.4
./configure --enable-optimizations
make -j8
sudo make altinstall
python3.6

-----------------------------------------
Создание базы данных
--------------------
sudo -u postgres psql
CREATE DATABASE banket;
CREATE USER b_user WITH PASSWORD 'password';
ALTER ROLE userdb SET client_encoding TO 'utf8';
ALTER ROLE userdb SET default_transaction_isolation TO 'read committed';
ALTER ROLE userdb SET timezone TO 'UTC';
GRANT ALL PRIVILEGES ON DATABASE movie TO userdb;
\q



-------------------------------------------
nano /etc/postgresql/15/main/pg_hba.conf

# Database administrative login by Unix domain socket
local  all            postgres                                peer
local  all            userdb                                    password


sudo -u postgres psql
postgres=# GRANT ALL ON SCHEMA public TO john;
---------------------------------------------


----------------------------------------
Установка Gunicorn
------------------
gunicorn project.wsgi:application --bind 111.222.333.44:8000
----------------------------------------
Настрока nginx
--------------
server {
    listen 80;
    server_name 111.222.333.44; # здесь прописать или IP-адрес или доменное имя сервера
    access_log  /var/log/nginx/example.log;
 
    location /static/ {
        root /home/user/myprojectenv/myproject/myproject/;
        expires 30d;
    }
 
    location / {
        proxy_pass http://127.0.0.1:8000; 
        proxy_set_header Host $server_name;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}

sudo service nginx restart

------------------------------------------
For SSL
-------
map $sent_http_content_type $expires {
    "text/html"                 epoch;
    "text/html; charset=utf-8"  epoch;
    default                     off;
}
server {
    listen 80;
    server_name www.django.com;
    return 301 https://django.com$request_uri;
}
server{
    listen 443 ssl;
    ssl on;                                      
    ssl_certificate /etc/ssl/django.crt;     
    ssl_certificate_key /etc/ssl/django.key; 
    server_name django.com;
    client_max_body_size 100M;

    gzip on;
    gzip_vary on;
    gzip_proxied any;
    gzip_http_version 1.1;
    gzip_types text/plain text/css application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript application/javascript;

    location /static/ {
        root /home/user/pj;
        expires 1d;
    }

    location /media/ {
        root /home/user/pj;
        expires 1d;
    }

    location / {
        proxy_pass http://127.0.0.1:8000;
        proxy_set_header Host $server_name;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}

---------------------------------------
Настройка supervisor
cd /etc/supervisor/conf.d/
sudo update-rc.d supervisor enable
sudo service supervisor start
sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl status project
sudo supervisorctl restart project
--------------------