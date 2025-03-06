# cloud-importer

This is a small tool to import and manage private images to cloud providers. It basically automate (and optimize) the commands 
you would need to run otherwise to import an image as a image on a cloud provider.

In addition to the `import` command it offers a `share` command to allow share images accross  accounts. Initially the image imported is private 
to the account which run the tool. In order to allow to use the image from a different account we would need to run the share command.

## RHEL AI

In order to test RHEL AI on AWS we need to import the image according to [RHEL AI installation guide](https://docs.redhat.com/en/documentation/red_hat_enterprise_linux_ai/1.4/html/installing/installing_on_aws) this tool will run those steps for us. Alhough previously the raw image should be donwloaded by an authenticated user to agree with EULA License.

To run the tool we can use the OCI container:

```bash
podman run --rm --name import-rhelai -d \
    -v ${PWD}:/workspace:z \
    -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
    -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
    -e AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} \
    quay.io/devtools-qe-incubator/cloud-importer:v0.0.1 rhelai aws \
        --backed-url "file:///workspace" \
        --raw-image-path "/workspace/rhel-ai-nvidia-aws-1.3.2-1736778584-x86_64.raw" \
        --ami-name "rhel-ai-nvidia-aws-1.3.2" \
        --debug \
        --debug-level 9

podman logs -f import-rhelai
```


