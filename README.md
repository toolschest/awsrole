# AWS Assume Role with MFA token

##### Features:
- Easily switch roles with yaml configuration using mfa token.

## Installation

Requires AWS CLI tools to be pre-installed.

Add alias in your .zshrc or .bashrc
```sh
alias awsrole='function(){output=$(command awsrole $@);if [ $? -eq 0 ]; then eval $output; else echo $output; fi}'
```

## Configuration

Checks for `$HOME/.aws/assumerole.yml`
OR
ENV `AWS_ASSUMEROLE_CONFIG` can be set

Environment variable has precedence over home directory path.

```sh
---
master_account_id: {master_account_id}
roles:
  production:
    environment: production
    username: developer
    role_arn: arn:aws:iam::{account_id}:role/Developer
    region: us-west-2
  sandbox:
    environment: sandbox
    username: developer
    role_arn: arn:aws:iam::{account_id}:role/Developer
    region: us-west-2
```

## Usage

```sh
awsrole --help
awsrole --env sandbox --mfa 123456
OR
awsrole -e production -mfa 123456
```

## License

MIT

**Free Software, Hell Yeah!**
