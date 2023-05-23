https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-ansible-on-ubuntu

sudo apt install -y python3-pip
sudo pip3 install ansible
sudo ln -s /usr/local/bin/ansible /usr/bin/ansible
sudo apt install sshpass
ansible-galaxy collection install community.general

### Подключение шары через smb в linux
mount -t cifs //10.5.50.231/student/Devops\ практика target_folder
