# Ansible Jenkins creation

Plugins defined inside role geerlingguy.jenkins/defaults/main
For ubuntu18:
```
apt-get update
apt install ansible -y
git clone https://github.com/jportasa/nr-cicd-basic.git
export ANSIBLE_ROLES_PATH=./nr-cicd-basic/ansible/roles/
ansible-playbook ./nr-cicd-basic/ansible/playbook/create_jenkins.yml
```
