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
所有请求都是二阶段的,也就是至少要发2个请求. 一个是先创建一个session, 一个是在该session里做最后的事务执行.s

1. 首先创建一个payment handle, 这个是通过参数来区分deposit/withdraw的.
2. 对于withdraw而言: 一旦payment handle的状态是payable的，那就可以直接处理了, 所谓的直接处理就是: 调用process-standalone-credits来落地 (二阶段构造)
3. 但是注意: withdraw调用process-standalone-credits结束后, 其状态还是pending.  其最终完成：要等webhook通知.
4. 对于deposit而言：一旦payment handle创建了, 其会返回一个收银台的url,前端打开后让用户支付即可.  付款后会有一个webhook发出来.  
5. 我们收到webhook后：要看一下是不是payable状态,如果是的话: 那就可以调用调用process-payment来完成事务落地了(二阶段构造)
6. 只是depposit比withdraw还要复杂的地方在于: 在调用process之前，需要获取token, 这个webhook给不了.需要单独去发一个请求查询.