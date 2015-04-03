- push to server

-- client side
```
ansible-share push -t ansible.1pcdev.com/nginx .
ansible-share push -t ansible.1pcdev.com/nginx .
ansible-share push -t ansible.1pcdev.com/nginx:latest .
ansible-share push -t ansible.1pcdev.com/postgres ~/ansible/postgres
```
push tar and gzipped file to server with a unique id. possibly encrypt as well

```
ansible-share push -t ansible.1pcdev.com/nginx:v1.1 .
```
push to server as tag v1.1, reject if duplicate tag

-- server side
```
POST ansible.1pcdev.com/role/nginx/29334143
POST ansible.1pcdev.com/role/nginx/v1.1
```
Store file in folder on server eg. /data/ansible_share/nginx/v1.1.tar.gz


- pull/update from server
---Tracks can changes on server

-- client side
```
ansible-share pull
ansible-share pull ansible.1pcdev.com/nginx:latest
ansible-share pull ansible.1pcdev.com/nginx:v1.1
ansible-share pull ansible.1pcdev.com/postgres:203918ad
```
request from servers, untar and unzip file and place in roles folder
add entry to Ansifile if one does not exist for specific version

-- server side
```
GET ansible.1pcdev.com/role/nginx/v1.1
```
send file to client

- clone
--- do not track server version
```
ansible-share clone ansible.1pcdev.com/nginx
ansible-share clone ansible.1pcdev.com/nginx:latest
ansible-share clone ansible.1pcdev.com/nginx:v1.3
```
request from servers, untar and unzip file and place in roles folder
do not add entry to Ansifile if one does not exist for specific version

- Ansifile example
```
source 'ansible.1pcdev.com'

nginx:latest
postgres:latest
mongodb:v1.3
```





