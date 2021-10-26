<template>
	<view class="page">
		<br/>
		<text>-　订单号:　{{ billId }}  </text>
		<view class="feedback-title">
			<text>问题和意见</text>
			<text class="feedback-quick" @tap="chooseMsg">快速键入</text>
		</view>
		<view class="feedback-body"><textarea placeholder="请详细描述你的问题和意见..." v-model="sendDate.content" class="feedback-textare"></textarea></view>
		<view class="feedback-title"><text>图片(选填,提供问题截图,总大小10M以下)</text></view>
		<view class="feedback-body feedback-uploader">
			<view class="uni-uploader">
				<view class="uni-uploader-head">
					<view class="uni-uploader-title">点击预览图片</view>
					<view class="uni-uploader-info">{{ imageList.length }}/5</view>
				</view>
				<view class="uni-uploader-body">
					<view class="uni-uploader__files">
						<block v-for="(image, index) in imageList" :key="index">
							<view class="uni-uploader__file" style="position: relative;">
								<image class="uni-uploader__img" :src="image" @tap="previewImage(index)"></image>
								<view class="close-view" @click="close(index)">x</view>
							</view>
						</block>
						<view class="uni-uploader__input-box" v-show="imageList.length < 5"><view class="uni-uploader__input" @tap="chooseImg"></view></view>
					</view>
				</view>
			</view>
		</view>
		<view class="feedback-title"><text>QQ/邮箱</text></view>
		<view class="feedback-body"><input class="feedback-input" v-model="sendDate.contact" placeholder="(选填,手机 QQ或E-mail,方便我们联系您 )" /></view>
		<button type="default" class="feedback-submit" @tap="send">提交</button>
		<view class="feedback-title"><text>您反馈的信息我们将以最快速度处理</text></view>
	</view>
</template>

<script>
import uniRate from '@/components/uni-rate/uni-rate.vue'
export default {
	components:{uniRate},
	data() {
		return {
			billId:"",
			index:0,
			msgContents: [],
			imageList: [],
			sendDate: {
				content: '',// 内容
				contact: '' // 联系方式
			}
		};
	},
	onLoad(options) {
		this.billId = options.billId;
		this.index = parseInt(options.index);

		this.msgContents.push("订单号："+this.billId );
		this.msgContents.push("申请退款");
		this.msgContents.push("订单损坏");
		this.msgContents.push("损坏包赔");
	},
	methods: {
		/**
		 * 关闭图片
		 * @param {Object} e
		 */
		close(e) {
			this.imageList.splice(e, 1);
		},

		/**
		 * 快速输入
		 */
		chooseMsg() {
			uni.showActionSheet({
				itemList: this.msgContents,
				success: res => {
					if (this.sendDate.content.length >0){
						this.sendDate.content += "\n"+this.msgContents[res.tapIndex];
					}else{
						this.sendDate.content += this.msgContents[res.tapIndex];
					}
					
				}
			});
		},

		/**
		 * 选择图片
		 */
		chooseImg() {
			//选择图片
			uni.chooseImage({
				sourceType: ['camera', 'album'],
				sizeType: 'compressed',
				count: 5 - this.imageList.length,
				success: res => {
					this.imageList = this.imageList.concat(res.tempFilePaths);
				}
			});
		},

		/**
		 * 预览图片
		 * @param {Object} index
		 */
		previewImage(index) {
			uni.previewImage({
				urls: this.imageList,
				current: this.imageList[index]
			});
		},

		/**
		 * 提交
		 */
		async send() {
			//发送反馈
			if (this.sendDate.content.length === 0) {
				uni.showModal({
					content: '请输入问题和意见',
					showCancel: false
				});
				return;
			}
			uni.showLoading({
				title: '上传中...'
			});
			let imgs = this.imageList.map((value, index) => {
				return {
					name: 'images' + index,
					uri: value
				};
			});
			
			// TODO 服务端限制上传文件一次最大不超过 2M, 图片一次最多不超过5张
			let simgs = [];
			if (imgs.length > 0){
				// #ifdef MP-WEIXIN
				simgs = await this.$Server.UploadFileWx(imgs);
				//#endif
				
				// #ifndef MP-WEIXIN
				let res = await this.$Server.UploadFile(imgs)
				if (res.fail){ // 失败打印输出
					uni.showToast({
						title: res.fail.error,
						icon: 'none'
					})	
					return;
				}
				if (res.success){// 设置session
					simgs = res.success;
				}
				//#endif
				console.log(simgs);
			}
			let resp = await this.$Order.DealOrder(this.billId,1,this.sendDate.content,this.sendDate.contact,simgs)
			if (resp != null) {
				uni.showModal({
				content: '成功反馈',
				showCancel: false,
				success: (res) => {
					this.$api.prePage().feedback(this.index);
					uni.navigateBack();
				}
				});
			}
			uni.hideLoading();
		},
	}
};
</script>

<style>
	 @import '/static/uni.wxss';
</style>

<style>
page {
	background-color: #efeff4;
}

.input-view {
	font-size: 28rpx;
}

.close-view {
	text-align: center;
	line-height: 14px;
	height: 16px;
	width: 16px;
	border-radius: 50%;
	background: #ff5053;
	color: #ffffff;
	position: absolute;
	top: -6px;
	right: -4px;
	font-size: 12px;
}
</style>
