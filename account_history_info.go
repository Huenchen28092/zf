/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/19 14:56
 ** @Author : yuebin
 ** @File : account_history_info
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/19 14:56
 ** @Software: GoLand
****************************************************/
package models

import (
	//"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
)

type AccountHistoryInfo struct {
	Id          int
	AccountUid  string
	AccountName string
	Type        string
	Amount      float64
	Balance     float64
	MerchantOrderId  string
	BankOrderId      string
	UpdateTime  string
	CreateTime  string
	Remark      string
	Notify      string
}

const ACCOUNT_HISTORY_INFO = "account_history_info"

func InsertAccountHistory(accountHistory AccountHistoryInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(accountHistory)
	if err != nil {
		logs.Error("insert account history fail: ", err)
		return false
	}
	return true
}

func GetAccountHistoryLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ACCOUNT_HISTORY_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Limit(-1).Count()
	if err != nil {
		logs.Error("get account history len by map fail: ", err)
	}
	return int(cnt)
}


func GetAccountHistoryLenAndSumByMapmount(params map[string]string) (float64,int) {
	o := orm.NewOrm()
	condition := "select sum(amount) as allAmount,COUNT(*) AS count from account_history_info "

	//successcondition := "select sum(amount) as allsuccessAmount,COUNT(*) AS count from account_history_info  "
	for _, v := range params {
		if len(v) > 0 {
			condition = condition + "where "
			//successcondition = successcondition + "where status = 'success' and "
			break
		}
	}
	flag := false
	if params["create_time__gte"] != "" {
		flag = true
		condition = condition + " update_time >= '" + params["create_time__gte"] + "'"
		//successcondition= successcondition + " update_time >= '" + params["create_time__gte"] + "'"
	}
	if params["create_time__lte"] != "" {
		if flag {
			condition = condition + " and "
			//successcondition= successcondition + " and "
		}
		condition = condition + " update_time <= '" + params["create_time__lte"] + "'"
		//successcondition= successcondition  + " update_time <= '" + params["create_time__lte"] + "'"
	}
	if params["account_name__icontains"] != "" {
		if flag {
			condition = condition + " and "
			//successcondition= successcondition + " and "
		}
		condition = condition + "account_name like '%" + params["account_name__icontains"] + "%' "
		//successcondition= successcondition + "merchant_name like '%" + params["merchant_name__icontains"] + "%' "
	}
	if params["account_uid"] != "" {
		if flag {
			condition = condition + " and "
			//successcondition= successcondition + " and "
		}
		condition = condition + " account_uid = '" + params["account_uid"] + "'"
		//successcondition= successcondition + " merchant_uid = '" + params["merchant_uid"] + "'"
	}
	if params["type"] != "" {
		if flag {
			condition = condition + " and "
			//successcondition= successcondition + " and "
		}
		condition = condition + " type = '" + params["type"] + "'"
		//successcondition= successcondition + " merchant_order_id = '" + params["merchant_order_id"] + "'"
	}
	if params["notify"] != "" {
		if flag {
			condition = condition + " and "
			//successcondition= successcondition + " and "
		}
		condition = condition + " notify = '" + params["notify"] + "'"
		//successcondition= successcondition + " bank_order_id = '" + params["bank_order_id"] + "'"
	}

	logs.Info("get order amount str = ", condition)
	//logs.Info("get order successamount str = ", successcondition)
	var maps []orm.Params
	//var mapssuces []orm.Params
	allAmount := 0.00
	//allsuccessAmount := 0.00
	count:=0
	//sxfallmount:=0.00
	//sxfpress:=0.00
	num, err := o.Raw(condition).Values(&maps)
	//numsuccess, errc := o.Raw(successcondition).Values(&mapssuces)
	if err == nil && num > 0 {
		if maps[0]["allAmount"] == nil{

			allAmount = 0.00
			count =0
		}else {
			count, _ = strconv.Atoi(maps[0]["count"].(string))
			allAmount, _ = strconv.ParseFloat(maps[0]["allAmount"].(string), 64)
		}
	}



	return allAmount,count
}


func GetAccountHistoryByMap(params map[string]string, displayCount, offset int) []AccountHistoryInfo {
	o := orm.NewOrm()
	qs := o.QueryTable(ACCOUNT_HISTORY_INFO)
	var accountHistoryList []AccountHistoryInfo
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&accountHistoryList)
	if err != nil {
		logs.Error("get account history by map fail: ", err)
	}
	return accountHistoryList
}
