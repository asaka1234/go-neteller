doc
========================
https://developer.paysafe.com/en/neteller-api-1/

http://paysafegroup.github.io/neteller_rest_api_v1/#/introduction/release-notes
https://www.neteller.com/fileadmin/content/pdfs/NETELLER_Mass_Payments_Guide.pdf


流程
========================
1. create a Payment Handle //创建pre-order
2. Use Webhooks  //callback


Comment
===============
1预充值/提现是同一个接口,通过参数来区分
2. 入金：只要创建handle即可, 它会返回一个收银台url,打开付款
3. 出金: 