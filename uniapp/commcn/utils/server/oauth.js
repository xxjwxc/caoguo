/* 用户信息相关 */
import Server from './def';

const user = {}

user.IsLogin = islogin=>{
//	console.log(Server.SessionID);

	let sessonid = Server.SessionID
	if (sessonid== "undefined" || sessonid== null || sessonid == null || sessonid == ""){
		return false;
	}

	return true;
};

// 获取用户信息
user.GetWxUserInfo = async function(){
	if (Server.IsLogin){
		return;
	}
	
	// #ifdef MP-WEIXIN
	await uni.getProvider({
		service: 'oauth',
		success: function(res) {
			if (~res.provider.indexOf('weixin')) {
				uni.login({
					provider: 'weixin',
					success: (res) => {
						uni.getUserInfo({
							provider: 'weixin',
							success: (info) => { //这里请求接口
								Server.Post("/weixin.update_user_info",info.rawData);// 更新用户信息
								Server.UserInfo = info.rawData;
								Server.IsLogin = true;
							},
							fail: () => {
								uni.showToast({
									title: "微信登录授权失败",
									icon: "none"
								});
							}
						})
					},
					fail: () => {
						uni.showToast({
							title: "微信登录授权失败",
							icon: "none"
						});
					}
				})

			} else {
				uni.showToast({
					title: '请先安装微信或升级版本',
					icon: "none"
				});
			}
		}
	});
	//#endif
};

async function wxLogin(jscode) {
	console.log(jscode)
	let resp = await Server.Post("/weixin.oauth",{"jscode":jscode});
	
	if (resp.success){// 设置session
		Server.SessionID = resp.success.session_id;
		Server.OpnID  = resp.success.open_id;
		console.log("session:"+Server.SessionID);
	}
	if (resp.fail){ // 失败打印输出
		uni.showToast({
			title: resp.fail.error,
			icon: 'none'
		})
	}
	
};

user.WxLogin = async function(){
	if (!this.IsLogin()){
		// #ifdef MP-WEIXIN
		await uni.login({
		  success: res => { 
			  wxLogin(res.code);
		  }
		});
		// #endif
	};
};

// 获取用户信息
user.GetUserInfo = async function(){
	let resp = await Server.Post("/user.get_user_info",{});
	return Server.GetReturn(resp);
};

// 请求vip
user.RequestVip = async function(){
	let resp = await Server.Post("/user.request_vip",{});
	return Server.GetReturn(resp);
};

// 关联用户
user.LinkUser = async function(openId){
	let resp = await Server.Post("/user.link_user",{openId:openId});
	return Server.GetReturn(resp);
};

export default user