<template>
	<view>
		<view v-for="(it, index) in list" :key="index" class="notice-item">
			<text class="time">{{it.time}}</text>
			<view class="content">
				<text class="title">联系方式:{{it.contact}}</text>
				<view class="img-wrapper">
					<image class="pic" v-for="(img, index1) in it.imgs" :key="index1" :src="img"></image>
				</view>
				<text class="introduce">
					{{it.remak}}
				</text>
				<view class="bot b-t">
					<text>{{billId}}</text>
					<text class="more-icon yticon icon-you"></text>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				billId:"",
				list: [],
			}
		},
		onLoad(options){
			this.billId = options.billId;
			console.log(options.billId)
			this.init(options.billId);
		},
		methods: {
			async init(billId){
				console.log(billId)
				let resp = await this.$Admin.GetAfterMsg(billId);
				this.list = resp.items;
			}
		}
	}
</script>

<style lang='scss'>
	page {
		background-color: #f7f7f7;
		padding-bottom: 30upx;
	}

	.notice-item {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.time {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 80upx;
		padding-top: 10upx;
		font-size: 26upx;
		color: #7d7d7d;
	}

	.content {
		width: 710upx;
		padding: 0 24upx;
		background-color: #fff;
		border-radius: 4upx;
	}

	.title {
		display: flex;
		align-items: center;
		height: 90upx;
		font-size: 32upx;
		color: #303133;
	}

	.img-wrapper {
		position: relative;
	}

	.pic {
		display: block;
		border-radius: 10upx;
	}

	.cover {
		display: flex;
		justify-content: center;
		align-items: center;
		position: absolute;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		background-color: rgba(0, 0, 0, .5);
		font-size: 36upx;
		color: #fff;
	}

	.introduce {
		display: inline-block;
		padding: 16upx 0;
		font-size: 28upx;
		color: #606266;
		line-height: 38upx;
	}

	.bot {
		display: flex;
		align-items: center;
		justify-content: space-between;
		height: 80upx;
		font-size: 24upx;
		color: #707070;
		position: relative;
	}

	.more-icon {
		font-size: 32upx;
	}
</style>
