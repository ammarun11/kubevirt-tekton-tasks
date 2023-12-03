## All namespace
#!/usr/bin/env bash
for TASK_NAME in cleanup-vm create-vm-from-manifest disk-virt-customize disk-virt-sysprep execute-in-vm generate-ssh-keys modify-data-object modify-windows-iso-file wait-for-vmi-status; do
    kubectl apply -f - << EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: ${TASK_NAME}-task
roleRef:
  kind: ClusterRole
  name: ${TASK_NAME}-task
  apiGroup: rbac.authorization.k8s.io
subjects:
  - kind: ServiceAccount
    name:  ${TASK_NAME}-task
    namespace: default
EOF
done

## S=pesific namespace
#!/usr/bin/env bash
for NAMESPACE in default; do
    for TASK_NAME in cleanup-vm create-vm-from-manifest disk-virt-customize disk-virt-sysprep execute-in-vm generate-ssh-keys modify-data-object modify-windows-iso-file wait-for-vmi-status; do
        kubectl apply -f - << EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: ${TASK_NAME}-task
  namespace: ${NAMESPACE}
roleRef:
  kind: ClusterRole
  name: ${TASK_NAME}-task
  apiGroup: rbac.authorization.k8s.io
subjects:
  - kind: ServiceAccount
    name:  ${TASK_NAME}-task
    namespace: default
EOF
    done
done


cleanup-vm create-vm-from-manifest disk-virt-customize disk-virt-sysprep execute-in-vm generate-ssh-keys modify-data-object modify-windows-iso-file wait-for-vmi-status    