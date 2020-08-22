package router

import (
	"api"
	"github.com/gin-gonic/gin"
)

func BackendRouter(router *gin.Engine) {
	//会议配置
	backend_router := router.Group("backend/datacenter/admin")
	backend_router.GET("meeting_list", api.BackendGetMettingList) //获取大会列表
	backend_router.GET("meeting",api.BackendGetMeeting)   //获取当前大会信息
	backend_router.DELETE("meeting",api.BackendDeleteMeeting)
	backend_router.POST("meeting", api.BackendAddMeeting) //新建大会
	//backend_router.POST("modify_meeting",BackendModifyMeeting)  //修改大会
	//backend_router.POST("change_meeting_state",BackendChangeMeetingState) //修改大会状态
	//backend_router.GET("system_admin",BackendGetSystemAdmin) //系统管理
	//backend_router.GET("person_info",BackendGetPersonInfo)  //大会人员信息表

}
