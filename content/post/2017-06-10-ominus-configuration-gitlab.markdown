# Omnibus configuration for GitLab CE in kubernetes

## Introduction

I don't know if you tried to install gitlab-ce with docker, but It took me some time to figure out how to setup this bloody omnibus configuration. I'll give you my little config and hopefully that will make you win a lot of time.
the struggle appeared when I started to configure SSO with gitlab-ce.

## The fix

With kubernetes you would like to setup your gitlab-ce configuration with nice yaml files, well it's start to be annoying when your omni configuration is a bit fancy. I have to admit it's silly once you know it though.
So the trick is to use multiline configuration + ; at the end. Here's the sample of my DeploymentConfig. In my case it's specific to Openshift, but it would work exactly the same with kubernetes.

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


Hope that helped you, and if you have a better way to setup your gitlab-ce configuration, I am really excited to hear about it !
Sur ce, codez bien!
