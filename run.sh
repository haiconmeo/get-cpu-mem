echo $(pwd)
chmod +x ./bin/main_ubuntu
sudo cp ./bin/main_ubuntu /usr/local/bin
sudo cp ./config/metric.service /etc/systemd/system
cd /etc/systemd/system
sudo chmod 664 metric.service 
sudo systemctl daemon-reload   
sudo systemctl enable --now metric.service 
sudo systemctl status metric.service