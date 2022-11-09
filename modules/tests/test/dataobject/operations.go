package dataobject

import (
	"context"
	"time"

	"github.com/kubevirt/kubevirt-tekton-tasks/modules/tests/test/constants"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	cdiv1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
	cdicliv1beta1 "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/typed/core/v1beta1"
)

func WaitForSuccessfulDataVolume(cdiClientSet cdicliv1beta1.CdiV1beta1Interface, namespace, name string, timeout time.Duration) error {
	return wait.PollImmediate(constants.PollInterval, timeout, func() (bool, error) {
		dataVolume, err := cdiClientSet.DataVolumes(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return true, err
		}
		return isDataVolumeImportStatusSuccessful(dataVolume), nil
	})
}

func WaitForSuccessfulDataSource(cdiClientSet cdicliv1beta1.CdiV1beta1Interface, namespace, name string, timeout time.Duration) error {
	return wait.PollImmediate(constants.PollInterval, timeout, func() (bool, error) {
		dataSource, err := cdiClientSet.DataSources(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return true, err
		}
		return IsDataSourceReady(dataSource), nil
	})
}

func WaitForBoundPVC(k8sClient *kubernetes.Clientset, namespace, name string, timeout time.Duration) error {
	return wait.PollImmediate(constants.PollInterval, timeout, func() (bool, error) {
		pvc, err := k8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return true, err
		}
		return isPVCStatusBound(pvc), nil
	})
}

func IsDataVolumeImportSuccessful(cdiClientSet cdicliv1beta1.CdiV1beta1Interface, namespace string, name string) bool {
	dataVolume, err := cdiClientSet.DataVolumes(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return false
	}
	return isDataVolumeImportStatusSuccessful(dataVolume)
}

func IsDataSourceReady(dataSource *cdiv1beta1.DataSource) bool {
	return getConditionMapDs(dataSource)[cdiv1beta1.DataSourceReady].Status == v1.ConditionTrue
}

func HasDataVolumeFailedToImport(dataVolume *cdiv1beta1.DataVolume) bool {
	conditions := getConditionMapDv(dataVolume)
	return dataVolume.Status.Phase == cdiv1beta1.ImportInProgress &&
		dataVolume.Status.RestartCount > constants.UnusualRestartCountThreshold &&
		conditions[cdiv1beta1.DataVolumeBound].Status == v1.ConditionTrue &&
		conditions[cdiv1beta1.DataVolumeRunning].Status == v1.ConditionFalse &&
		conditions[cdiv1beta1.DataVolumeRunning].Reason == constants.ReasonError
}

func isDataVolumeImportStatusSuccessful(dataVolume *cdiv1beta1.DataVolume) bool {
	return getConditionMapDv(dataVolume)[cdiv1beta1.DataVolumeBound].Status == v1.ConditionTrue &&
		dataVolume.Status.Phase == cdiv1beta1.Succeeded
}

func isPVCStatusBound(pvc *v1.PersistentVolumeClaim) bool {
	return pvc.Status.Phase == v1.ClaimBound
}
