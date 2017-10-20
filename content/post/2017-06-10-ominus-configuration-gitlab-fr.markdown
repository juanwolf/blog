---
title:  Configuration omnibus pour GitLab CE avec kubernetes
date:   2017-05-26
tags: ["kubernetes"]
---

# Configuration Omnibus pour GitLab CE avec kubernetes

## Introduction

Si vous avez une configuration un peu funkie de gitlab, vous avez du surement galere avec la variable GITLAB_OMINBUS_CONFIG. Pour etre franc, j'ai un peu galere a la definir dans la configuration de mon deployment pour kubernetes.

## Le fix

Avec Kubernetes, vous voudrez surement definir la configuration de votre deploiment en yaml. L'astuce est d'utiliser le symbole de ligne multiple ">" avec les ";" a la fin de chaque expression. Voici un morceau de mon "DeploymentConfig":

```
apiVersion: v1
kind: DeploymentConfig
metadata:
...
spec:
  replicas: 1
  template:
    spec:
      containers:
      - env:
        - name: GITLAB_OMNIBUS_CONFIG
          value: >
            hostname='your-hostname.com';
            external_url "http://#{hostname}/" unless hostname.to_s == ''; root_pass='';
            gitlab_rails['initial_root_password']=password unless root_pass.to_s == '';
            postgresql['enable']=false;
            gitlab_rails['db_host'] = 'gitlab-ce-postgresql';
            gitlab_rails['db_password']='the_password_of_your_psql';
            gitlab_rails['db_username']='your_db_username';
            gitlab_rails['db_database']='gitlab_db';
            redis['enable'] = false;
            gitlab_rails['redis_host']='gitlab-ce-redis';
            unicorn['worker_processes'] = 2 ; manage_accounts['enable'] =
            true; manage_storage_directories['manage_etc'] = false;
            gitlab_shell['auth_file'] = '/gitlab-data/ssh/authorized_keys';
            git_data_dir '/gitlab-data/git-data';
            gitlab_rails['shared_path'] = '/gitlab-data/shared' ;
            gitlab_rails['uploads_directory'] = '/gitlab-data/uploads';
            gitlab_ci['builds_directory'] = '/gitlab-data/builds';
            prometheus_monitoring['enable'] = false;
            gitlab_rails['omniauth_enabled'] = true;
            gitlab_rails['omniauth_allow_single_sign_on'] = ['saml'];
            gitlab_rails['omniauth_block_auto_created_users'] = false;
            gitlab_rails['omniauth_auto_link_ldap_user'] = false;
            gitlab_rails['omniauth_auto_link_saml_user'] = false;
            gitlab_rails['omniauth_providers'] = [{
                name: 'saml',
                label: 'Your Nice Label',
                args: {
                    issuer: 'https://your-hostname',
                    assertion_consumer_service_url: 'https://your_hostname.com/users/auth/saml/callback',
                    idp_cert_fingerprint: 'the_nice_fingerprint',
                    idp_sso_target_url: 'https://your_sso_provider/adfs/ls/',
                    allowed_clock_drift: 1,
                    name_identifier_format: 'identifier_formate',
                    attribute_statements: {
                        email: ['email'],
                        name: ['fullname']
                    }
                }
            }];
        ...
```


J'espere que ca vous aura aide. Si jamais vous avez une meilleure facon de faire, hesitez pas a la poster dans les commentaires !
Sur ce, codez bien!
