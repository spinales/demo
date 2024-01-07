#!/bin/bash

# install
sudo apt update
sudo apt install -y nginx postgresql postgresql-contrib golang make

# postgres
sudo systemctl start postgresql.service
sudo systemctl status postgresql.service
sudo -u postgres createuser app --interactive
sudo -u postgres createdb northwind
psql -U postgres -d northwind -f ./northwind.sql

# app
sudo sh -c 'printf "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile'
go version
sudo touch /etc/systemd/system/app.service
sudo chmod 664 /etc/systemd/system/app.service
make compile
sudo sh -c 'printf "[Unit]\nDescription=Construccion Demo\n\n[Service]\nExecStart=/opt/demo/build/demo\n\n[Install]\nWantedBy=multi-user.target" > /etc/systemd/system/app.service'
sudo systemctl enable app
sudo systemctl start app
sudo systemctl status app

# nginx
sudo systemctl enable nginx
sudo systemctl start nginx
sudo ufw allow 'Nginx HTTP'
sudo ufw status
systemctl status nginx
sudo sh -c 'printf "server {\n\tlisten 80;\n\tlisten 443 ssl;\n\n\tlocation / {\n\t\tproxy_pass http://localhost:8080;\n\t\tproxy_set_header Host \$host;\n\t\tproxy_set_header X-Real-IP \$remote_addr;\n\t\tproxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;\n\t\tproxy_set_header X-Forwarded-Proto \$scheme;\n\t}\n\n}" > /etc/nginx/conf.d/default.conf'
nginx -t
nginx -s reload
