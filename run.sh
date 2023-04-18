echo $(pwd)
wget https://cloud-bot.s3.ap-southeast-1.amazonaws.com/tools/main_ubuntu
chmod +x ./main_ubuntu
sudo cp ./main_ubuntu /usr/local/bin
wget https://cloud-bot.s3.ap-southeast-1.amazonaws.com/tools/metric.service
sudo cp ./metric.service /etc/systemd/system
cd /etc/systemd/system
sudo chmod 664 metric.service 
sudo systemctl daemon-reload   
sudo systemctl enable --now metric.service 