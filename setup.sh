#!/bin/bash

# install
sudo apt update
sudo apt install -y nginx golang make

# todo: download repo
# todo: create unit file environment, vim /etc/systemd/system/app.service.d/env.conf
# app
sudo sh -c 'printf "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile'
go version
sudo touch /etc/systemd/system/app.service
sudo chmod 664 /etc/systemd/system/app.service
make compile
sudo sh -c 'printf "[Unit]\nDescription=Construccion Demo\n\n[Service]\nExecStart=/root/demo/build/demo\n\n[Install]\nWantedBy=multi-user.target\n" > /etc/systemd/system/app.service'
mkdir -p /etc/systemd/system/app.service.d/
systemctl daemon-reload
sudo systemctl enable app
sudo systemctl start app
sudo systemctl status app

# nginx
sudo systemctl enable nginx
sudo systemctl start nginx
sudo ufw allow 'Nginx HTTP'
sudo ufw status
systemctl status nginx
# this a template for reverse proxy
sudo sh -c 'printf "server {\n\t\tlisten 80;\n\t\tlisten [::]:80;\n\t\tlisten 443;\n\t\tlisten [::]:443;\n\n\t\tserver_name ${IP};\n\n\n\t\tproxy_set_header Host $host;\n\t\tproxy_set_header X-Real-IP $remote_addr;\n\t\tproxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;\n\t\tproxy_set_header X-Forwarded-Proto $scheme;\n\n\t\tlocation / {\n\t\t\t\tproxy_pass http://127.0.0.1:8080/;\n\t\t}\n\n}" > /etc/nginx/conf.d/default.conf'
nginx -t
nginx -s reload

# todo: edit nginx reverse proxy config