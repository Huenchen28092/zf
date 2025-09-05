/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/16 11:11
 ** @Author : yuebin
 ** @File : account
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/16 11:11
 ** @Software: GoLand
****************************************************/
package models

import (
	"github.com/beego/beego/v2/client/orm"
	//"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
	"time"
)

type AccountInfo struct {
	Id           int
	Status       string
	AccountUid   string
	AccountName  string
	Balance      float64 //账户总余额
	SettleAmount float64 //已经结算的金额
	LoanAmount   float64 //账户押款金额
	FreezeAmount float64 //账户冻结金额
	WaitAmount   float64 //待结算资金
	PayforAmount float64 //代付在途金额
	//AbleBalance  float64 //账户可用金额
	UpdateTime string
	CreateTime string
}

const ACCOUNT_INFO = "account_info"

func InsetAcount(account AccountInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&account)
	if err != nil {
		logs.Error("insert account fail: ", err)
		return false
	}
	return true
}

func GetAccountByUid(accountUid string) AccountInfo {
	o := orm.NewOrm()
	var account AccountInfo
	_, err := o.QueryTable(ACCOUNT_INFO).Filter("account_uid", accountUid).Limit(1).All(&account)
	if err != nil {
		logs.Error("get account by uid fail: ", err)
	}

	return account
}

func GetAccountLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(ACCOUNT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Limit(-1).OrderBy("-update_time").Count()
	if err != nil {
		logs.Error("get account len by map fail: ", err)
	}
	return int(cnt)
}

func GetAccountByMap(params map[string]string, displayCount, offset int) []AccountInfo {
	o := orm.NewOrm()
	var accountList []AccountInfo
	qs := o.QueryTable(ACCOUNT_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}

	_, err := qs.Limit(displayCount, offset).OrderBy("-update_time").All(&accountList)
	if err != nil {
		logs.Error("get account by map fail: ", err)
	}
	return accountList
}
func Todayallmonert(user string)(todayall ,todaypress ,yesdatay ,yesdataypress float64) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, now.Location())

	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayStart := yesterday + " 00:00:00"
	yesterdayEnd := yesterday + " 23:59:59"

	ystart, _ := time.Parse("2006-01-02 15:04:05", yesterdayStart)
	yend, _ := time.Parse("2006-01-02 15:04:05", yesterdayEnd)


	o := orm.NewOrm()

	cond := "select  sum(fact_amount) as order_amount,sum(all_profit) as today_prtess, count(1) as order_count, " +
		"sum(platform_profit) as platform_profit, sum(agent_profit) as agent_profit from " + ORDER_PROFIT_INFO + " where status='success' and merchant_uid =? and create_time BETWEEN ? AND ?"
	ycond := "select  sum(fact_amount) as yorder_amount,sum(all_profit) as ytoday_prtess, count(1) as order_count, " +
		"sum(platform_profit) as platform_profit, sum(agent_profit) as agent_profit from " + ORDER_PROFIT_INFO + " where status='success' and merchant_uid =? and create_time BETWEEN ? AND ?"
	var maps []orm.Params
	var ymaps []orm.Params
	o.Raw(cond,user, start, end).Values(&maps)
	o.Raw(ycond,user, ystart, yend).Values(&ymaps)
	//logs.Info("sql语句是=="+cond)
	//_, _ = o.Raw(cond).Values(&maps)
	allAmount := 0.00
	today_prtess :=0.00
	yesAmount := 0.00
	yes_prtess :=0.00

	if maps[0]["order_amount"] == nil{
		allAmount = 0.00
	}else {
		allAmount, _ = strconv.ParseFloat(maps[0]["order_amount"].(string), 64)
	}
	if maps[0]["today_prtess"] == nil{
		today_prtess = 0.00
	}else {
		today_prtess, _ = strconv.ParseFloat(maps[0]["today_prtess"].(string), 64)
	}

	if ymaps[0]["yorder_amount"] == nil{
		yesAmount = 0.00
	}else {
		yesAmount, _ = strconv.ParseFloat(ymaps[0]["yorder_amount"].(string), 64)
	}
	if ymaps[0]["ytoday_prtess"] == nil{
		yes_prtess = 0.00
	}else {
		yes_prtess, _ = strconv.ParseFloat(ymaps[0]["ytoday_prtess"].(string), 64)
	}



	//allAmount, _ = strconv.ParseFloat(maps[0]["order_amount"].(string), 64)
	//today_prtess, _ = strconv.ParseFloat(maps[0]["today_prtess"].(string), 64)
	//

	return allAmount,today_prtess,yesAmount,yes_prtess
}

func GetAllAccount() []AccountInfo {
	o := orm.NewOrm()
	var accountList []AccountInfo

	_, err := o.QueryTable(ACCOUNT_INFO).Limit(-1).All(&accountList)

	if err != nil {
		logs.Error("get all account fail: ", err)
	}

	return accountList
}

func UpdateAccount(account AccountInfo) bool {
	o := orm.NewOrm()
	_, err := o.Update(&account)
	if err != nil {
		logs.Error("update account fail: ", err)
		return false
	}
	return true
}

func DeleteAccountByUid(accountUid string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(ACCOUNT_INFO).Filter("account_uid", accountUid).Delete()
	if err != nil {
		logs.Error("delete account fail: ", err)
		return false
	}
	return true
}
