### Установка и настройка Kubespray
pip3 install ansible==5.10.0
pip3 install ruamel-yaml netaddr

Настройка inventory:
declare -a IPS=(192.168.1.101 192.168.1.102 192.168.1.103)
CONFIG_FILE=inventory/mycluster/hosts.yaml python3 contrib/inventory_builder/inventory.py ${IPS[@]}
Добавить в
containerd_registries:
  "gcr.io": "https://mirror.gcr.io"
Настроить необходимые параметры переменных в hosts.yaml, all.yml, containerd.yml, addons.yml, k8s-cluster.yml, etcd.yml.
[Параметры](https://github.com/kubernetes-sigs/kubespray/blob/master/docs/large-deployments.md) _kube_ переменных для пересоздания подов.

Запуск установка кластера:
ansible-playbook -i inventory/mycluster/hosts.yaml -u root -k cluster.yml


### Установка и настройка Kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
mv kubectl /usr/local/bin/kubectl && chmod 755 /usr/local/bin/kubectl
kubectl completion bash | sudo tee /etc/bash_completion.d/kubectl > /dev/null
Копируем содержимое /etc/kubernetes/admin.conf с master node себе в ~/.kube/config

### Установка Helm
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
