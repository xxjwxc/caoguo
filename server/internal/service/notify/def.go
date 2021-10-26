package notify

const (
	body = `<html>
	<body>
	<div style="background-color:#ECECEC; padding: 35px;">
	<table cellpadding="0" align="center" style="width: 600px; margin: 0px; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:Courier New,Bold; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;">
	<tbody>
	<tr>
	<th valign="middle" style="height: 25px; line-height: 25px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: #000000; background-color: #62c30b; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;">
	<font face="Courier New" size="5" style="color: rgb(255, 255, 255); ">潞潮公社</font>
	</th>
	</tr>
	<tr>
	<td>
	<div style="padding:25px 35px 40px; background-color:#fff;">
	<h2 style="margin: 5px 0px; "><font color="#333333" style="line-height: 20px; ">
	<font style="line-height: 22px; " size="4">新下单提醒</font></font></h2>
	<p>&nbsp;</p>
	<p>订单号 : %v </p>
	<p>商品列表：</p>
	%v
	<p>请不要回复此邮件,谢谢!</p>
	<p align="right">LUCHAO development center &nbsp;</p>
	<p align="right">%v</p>
	</div>
	</td>
	</tr>
	</tbody>
	</table>
	</div>
	</body>
	</html>`
	subject = "LUCHAO place order notify. %v"
)
