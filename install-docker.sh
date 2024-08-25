#!/bin/bash
# The "-x" here allows to print each Bash command before it gets executed

echo -e "\n--------------------------------------------------------------------------"
echo "Removing any older installations of Docker..."
echo -e "--------------------------------------------------------------------------\n"
for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do apt-get remove $pkg; done
apt purge -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
rm -rf /var/lib/docker
rm -rf /var/lib/containerd

echo -e "\n--------------------------------------------------------------------------"
echo "Adding Docker repository to APT's sources..."
echo -e "--------------------------------------------------------------------------\n"
apt update
apt install ca-certificates curl
install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
chmod a+r /etc/apt/keyrings/docker.asc

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  tee /etc/apt/sources.list.d/docker.list > /dev/null

echo -e "\n--------------------------------------------------------------------------"
echo "Installing Docker and Docker Compose..."
echo -e "--------------------------------------------------------------------------\n"
apt update && apt upgrade -y
apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

echo -e "\n--------------------------------------------------------------------------"
echo "Running a test container on Docker to verify successful installation..."
echo -e "--------------------------------------------------------------------------\n"
docker run hello-world

echo -e "\n--------------------------------------------------------------------------"
echo "Looks like we are up and running, here are the versions installed:"
docker -v
docker compose version
echo -e "--------------------------------------------------------------------------\n"

echo -e "\n--------------------------------------------------------------------------"
echo "Setting up some common aliases..."
echo 'alias "myip"="curl ipnr.dk"' >> /root/.bashrc
echo 'alias "ll"="ls -lhta"' >> /root/.bashrc
echo 'alias "llog"="less -N +F /var/log/syslog"' >> /root/.bashrc
echo 'alias "dc"="docker container ls -a"' >> /root/.bashrc
echo 'alias "di"="docker image ls -a"' >> /root/.bashrc
echo 'alias "dcr"="docker compose rm -fsv"' >> /root/.bashrc
echo 'alias "dir"="docker image rm"' >> /root/.bashrc
echo 'alias "dup"="docker compose up -d"' >> /root/.bashrc
echo 'alias "dl"="docker logs"' >> /root/.bashrc
echo 'alias "dl100"="docker logs -n 100"' >> /root/.bashrc
echo 'alias "myaliases"="cat /root/.bashrc | grep alias"' >> /root/.bashrc
echo -e "--------------------------------------------------------------------------\n"
