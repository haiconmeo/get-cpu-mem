echo $(pwd)
chmod +x ./bin/main_ubuntu
sudo mv ./bin/main_ubuntu /usr/local/bin
sudo mv ./config/metric.service /etc/systemd/system
cd /etc/systemd/system
sudo chmod 664 metric.service 
sudo systemctl daemon-reload   
sudo systemct enable --now metric.service 
sudo systemctl status metric.service