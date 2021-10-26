<template>
	<view>
		<view class="link-container" v-if="items.length > 0">
			<view class="link-item" v-for="(item, index) in items" :key="index" @tap="navigateTo(item.gpid)">
				<view class="logo">
					<image  :src="item.icon"></image>
				</view>
				<view class="link-middle">
					<view class="title">
						 {{item.name}}
					</view>
					<view class="desciption">
						￥{{(item.price * 0.01)}}
					</view>  
				</view>
				<view class="go-link">
					<image src="/static/to.png"></image>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				items: [],
			}
		},
		onLoad() {
			this.OnInit()
		},
		methods: {
			async OnInit(){
				let resp = await this.$Product.GetFavorite();
				if ((typeof(resp.items) != "undefined" && resp.items != null) ){
					this.items = resp.items
				}
			},
			// 查看物流
			async navigateTo(gpid){
				uni.navigateTo({
					url: `/pages/product/product?gpid=${gpid}`
				})
			},
		}
	}
</script>

<style>
	page{
		background-color: #f1f1f1 !important
	}
	.link-container{
		margin: 50rpx 30rpx; 
	}
	.link-item{
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		height: 150rpx;
		padding: 30rpx;
		margin:20rpx auto;
		background-color:#ffffff;
		border-radius: 10upx;
		box-shadow: 0 5upx 20upx 0upx rgba(0, 0, 150, .2);
	}	
	.logo image{
		border-radius: 5%;
		width: 100rpx;
		height: 100rpx;
		box-shadow: 0 0 2px 1px rgba(0, 140, 186, 0.5);
	}
	.link-middle{
		width: 70%;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		margin: 0 20rpx;
	}
	.title {
		vertical-align: text-top;
		font-size: 32rpx;
		font-weight: 500;
	}	
	.desciption{
		font-size: 26rpx;
		color: #DD524D;
	}
	.go-link{
		display: flex;
		flex-direction: column;
		justify-items:center;
		justify-content: center;
	}
	.go-link image {
		height: 50rpx;
		width: 50rpx;
	}

</style>
