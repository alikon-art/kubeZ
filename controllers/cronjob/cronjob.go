package cronjob

import (
	"context"
	"fmt"
	"kubez_project/config"
	"kubez_project/models"
	"kubez_project/utils/gins"
	goclient "kubez_project/utils/go_client"

	"github.com/gin-gonic/gin"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add 创建一个新的 CronJob
func Add(c *gin.Context) {
	var cronJob batchv1beta1.CronJob
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &cronJob)
	if err != nil {
		return
	}

	fmt.Println(cronJob)
	_, err = clientset.BatchV1beta1().CronJobs(cronJob.Namespace).Create(context.TODO(), &cronJob, metav1.CreateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "创建 CronJob 失败", err)
		return
	}

	gins.ReturnData(c, "200", "创建 CronJob 成功", nil)
}

// Delete 删除一个指定的 CronJob
func Delete(c *gin.Context) {
	var cronJob batchv1beta1.CronJob
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &cronJob)
	if err != nil {
		return
	}

	err = clientset.BatchV1beta1().CronJobs(requestData.Namespace).Delete(context.TODO(), requestData.Name, metav1.DeleteOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "删除 CronJob 失败", err)
		return
	}

	gins.ReturnData(c, "200", "删除 CronJob 成功", nil)
}

// Update 更新一个现有的 CronJob
func Update(c *gin.Context) {
	var cronJob batchv1beta1.CronJob
	clientset, _, err := gins.BoundJsonAndInitClientSet(c, &cronJob)
	if err != nil {
		return
	}

	_, err = clientset.BatchV1beta1().CronJobs(cronJob.Namespace).Update(context.TODO(), &cronJob, metav1.UpdateOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "更新 CronJob 失败", err)
		return
	}

	gins.ReturnData(c, "200", "更新 CronJob 成功", nil)
}

// List 列出所有或特定的 CronJobs
func List(c *gin.Context) {
	var cronJob batchv1beta1.CronJob
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &cronJob)
	if err != nil {
		return
	}

	listOptions := metav1.ListOptions{
		LabelSelector: goclient.Convert2LabelSelector(requestData.Labels),
	}

	cronJobList, err := clientset.BatchV1beta1().CronJobs(requestData.Namespace).List(context.TODO(), listOptions)
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取 CronJob 列表失败", err)
		return
	}

	var returnDataList models.BasicReturnList
	for _, job := range cronJobList.Items {
		returnData := models.BasicReturn{
			Name:       job.Name,
			Namespace:  job.Namespace,
			Labels:     job.Labels,
			CreateTime: job.CreationTimestamp.Format(config.TimestampFormat),
		}
		returnDataList.AddBasicReturn(returnData)
	}

	gins.ReturnData(c, "200", "获取 CronJob 列表成功", returnDataList)
}

// Get 获取一个指定的 CronJob 详细信息
func Get(c *gin.Context) {
	var cronJob batchv1beta1.CronJob
	clientset, requestData, err := gins.BoundJsonAndInitClientSet(c, &cronJob)
	if err != nil {
		return
	}

	job, err := clientset.BatchV1beta1().CronJobs(requestData.Namespace).Get(context.TODO(), requestData.Name, metav1.GetOptions{})
	if err != nil {
		gins.ReturnErrorData(c, "500", "获取 CronJob 详情失败", err)
		return
	}

	gins.ReturnData(c, "200", "获取 CronJob 详情成功", job)
}
