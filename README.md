# Ansible Share

Share ansible roles between projects

## Purpose

There are a few ways to share roles used in [ansible]() such as Ansible Galaxy or using git submodules. Both of these have some issues in practice. Ansible Galaxy has only public roles so you can't store anything specific to you infrastructure where as git submodules can cause problems when upgrading roles with other playbooks not ready for the upgrade.

Ansible Share was created to try and solve this issue by providing a simple cli backed by [Amazon S3]() that will version ansible roles when uploaded and allow you to up date them when you wish.

## Current status

v0.1

## Pre-requisites

* aws account with an S3 bucket setup to store roles 
* golang

## Installation

go get -u github.com/pythonandchips/ansible-share

## Usage

Before using ansible share you will need to create an S3 bucket to store you roles and aws access key with access to the bucket  ether stored as environment variables or held in the `.aws/credentials` file

### Pushing roles to the store

`ansible-share push -t {bucket name}/{role name}:{version} {path to role}`

- bucket name: Then name of the bucket to push the roles to
- role name: Name of the role to be pushed to S3
- version (optional): a specifed version for that upload e.g v1.1. if this is not supplied a random generated number will be used instead  
- path to role: the relative path to the role to be uploaded

This will push a new version of a role to S3.

### Pulling roles from repo

`ansible-share pull {bucket name}/{role name}:{version}`

- bucket name: Then name of the bucket to push the roles to
- role name: Name of the role to be pushed to S3
- version (optional): a specifed version for that upload e.g v1.1. if this is not supplied the latest version will be pulled down

This will pull a role and install it in the `roles` directory. It will also write the role to the ansifile.

### Pulling all roles in the ansifile

`ansible-share pull`

This will pull all roles in the ansifile and install them in a `roles` directory

## Securing Roles

As roles are stored in S3 the easiest way to secure roles is using using amazon IAM polices. 

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:ListBucket",
                "s3:PutObject"
            ],
            "Resource": [
                "arn:aws:s3:::{bucket name}/*"
            ]
        }
    ]
}
```
Create a policy as above and attach it to any users you wish to manage ansible roles from ansible share. Each user should use their own aws creditials to access stored roles.

## Future features

In the future I'll be looking to add the following features

- better error handling and check
- allow configuration and specification of where roles are saved
- Update roles command
- list available roles
- Support multiple storage types (e.g. rackspace files, private server with http)

## Bugs/Features/Prase

It you find any bugs or have some feature requests please add an issue on the repository. Or if you just want to get in touch and tell me how awesome git presenter is you can get me on twitter @colin_gemmell or drop me an email at pythonandchips{at}gmail.com

## Contributing to ansible-share

* Check out the latest master to make sure the feature hasn't been implemented or the bug hasn't been fixed yet
* Check out the issue tracker to make sure someone already hasn't requested it and/or contributed it
* Fork the project
* Start a feature/bugfix branch
* Commit and push until you are happy with your contribution
* Make sure to add tests for it. This is important so I don't break it in a future version unintentionally.

## Copyright

Copyright (c) 2016 Colin Gemmell

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
