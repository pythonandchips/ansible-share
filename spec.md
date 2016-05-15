- push to storage

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
push to storage as tag v1.1, reject if duplicate tag

- pull/update from storage

-- client side
```
ansible-share pull
ansible-share pull ansible.1pcdev.com/nginx:latest
ansible-share pull ansible.1pcdev.com/nginx:v1.1
ansible-share pull ansible.1pcdev.com/postgres:203918ad
```
request from storage, untar and unzip file and place in roles folder
add entry to Ansifile if one does not exist for specific version

- Ansifile example
```
ansible.1pcdev.com/nginx:latest
ansible.1pcdev.com/postgres:latest
ansible.1pcdev.com/mongodb:v1.3
```
