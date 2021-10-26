<template>
	<view class="content">
		<view class="navbar">
			<view v-for="(item, index) in navList" :key="index" 
				class="nav-item" 
				:class="{current: tabCurrentIndex === index}"
				@click="tabClick(index)"
			>
				{{item.text}}
			</view>
		</view>

		<swiper :current="tabCurrentIndex" class="swiper-box" duration="300" @change="changeTab">
			<swiper-item class="tab-content" v-for="(tabItem,tabIndex) in navList" :key="tabIndex">
				<scroll-view 
					class="list-scroll-content" 
					scroll-y
					@scrolltolower="loadData"
				>
					<!-- 空白页 -->
					<empty v-if="tabItem.loaded === true && tabItem.orderList.length === 0">
					</empty>
					
					<!-- 订单列表 -->
					<view v-for="(item,index) in tabItem.orderList" :key="index" class="order-item" >
						<view class="i-top b-b">
							<text class="time">{{item.time}}</text>
							<text class="state" :style="{color: item.stateTipColor}">{{item.stateTip}}</text>
							<!-- 删除订单 -->
							<text v-if="item.status===-1 ||item.status===3 ||item.status===5 || item.status === 6 " 
							class="del-btn yticon icon-iconfontshanchu1" @click="deleteOrder(item.billId,index)" ></text>
						</view>
						
						<scroll-view v-if="item.items.length > 1" class="goods-box" scroll-x>
							<view v-for="(goodsItem, goodsIndex) in item.items" :key="goodsIndex" class="goods-item" >
								<image class="goods-img" :src="goodsItem.icon" mode="aspectFill" @click="naveToProduct(goodsItem.gpid)"></image>
							</view>
						</scroll-view>
						<view v-if="item.items.length === 1" >
							<view class="goods-box-single" v-for="(goodsItem, goodsIndex) in item.items" :key="goodsIndex">
								<image class="goods-img" :src="goodsItem.icon" mode="aspectFill" @click="naveToProduct(goodsItem.gpid)"></image>
								<view class="right">
									<text class="title clamp">{{goodsItem.name}}</text>
									<text class="attr-box">{{goodsItem.skuVal}}  x {{goodsItem.number}}</text>
									<text class="price">{{goodsItem.price*0.01}}</text>
								</view>
								
							</view> 
						</view>
						<text class="order-box clamp" selectable = "true"> 订 单 号 : {{item.billId}}</text>
						<text class="order-box clamp" selectable = "true"> 配送地址 : {{item.addr.address}}</text>

						<view class="price-box">
							共
							<text class="num">{{item.number}}</text>
							件商品 实付款
							<text class="price">{{(item.totalFee*0.01)}}</text>
						</view>
						<view class="action-box b-t" v-if="item.checkStatus == 1">
							<button class="action-btn" @click="cancelOrder(item.billId,index)">取消订单</button>
							<button class="action-btn recom" @click="payOrder(item.billId,index)">立即支付</button>
						</view>
						<view class="action-box b-t" v-if="item.checkStatus == 2">
							<button class="action-btn" @click="cancelOrder(item.billId,index)">取消订单</button>
							<button class="action-btn"  open-type="contact">联系客服</button>
						</view>
						<view class="action-box b-t" v-if="item.checkStatus == 4">
							<button class="action-btn" @click="toLogistics(item.billId)">查看物流</button>
							<button class="action-btn" @click="afterSale(item.billId,index)">申请售后</button>
							<button class="action-btn"  open-type="contact">联系客服</button>
						</view>
						<view class="action-box b-t" v-if="item.checkStatus == 6">
							<button class="action-btn" @click="toGetLogistics(item.billId)">催促售后</button>
							<button class="action-btn"  open-type="contact">联系客服</button>
						</view>
					</view>

					<uni-load-more :status="tabItem.loadingType"></uni-load-more>

				</scroll-view>
			</swiper-item>
		</swiper>
	</view>
</template> 

<script>
	import uniLoadMore from '@/components/uni-load-more/uni-load-more.vue';
	import empty from "@/components/empty";
	import logisticsVue from '../logistics/logistics.vue';
	export default {
		components: {
			uniLoadMore,
			empty
		},
		data() {
			return {
				tabCurrentIndex: 0,
				navList: [],
			};
		},
		onLoad(options){
			/**
			 * 修复app端点击除全部订单外的按钮进入时不加载数据的问题
			 * 替换onLoad下代码即可
			 */
			this.tabCurrentIndex = +options.state;
			this.onInit();
		},
		 
		methods: {
			async onInit(){
				 let resp = await this.$Order.GetMyOrdersConfig(false);
				// this.navList = resp.navList;
				resp.navList.forEach(item=>{
					this.$set(item, "orderList", []);
					this.navList.push(item);
				})

				 await this.loadData();
			},

			//获取订单列表
			async loadData(source){
				//这里是将订单挂载到tab列表下
				let index = this.tabCurrentIndex;
				let navItem = this.navList[index];
				let state = navItem.state;
				let pageNumber = navItem.pageNumber;
				if (typeof(pageNumber) == "undefined" || pageNumber == null){
					pageNumber = 0;
				}
				if (typeof(navItem.orderList)== "undefined" || navItem.orderList == null){
					navItem.orderList = [];
				}
				
				if(source === 'tabChange' && navItem.loaded === true){
					//tab切换只有第一次需要加载数据
					return;
				}
				if(navItem.loadingType === 'loading'){
					//防止重复加载
					return;
				}

				navItem.loadingType = 'loading';
				setTimeout(async ()=>{
					let resp = await  this.$Order.GetMyOrders(state,pageNumber);
					if (resp == null || typeof(resp.items) == "undefined" || resp.items == null){
						navItem.loadingType = 'noMore';
						// navItem.pageNumber = pageNumber;
					} else {
						navItem.loadingType = 'more';
						this.$set(navItem, 'pageNumber', pageNumber+1);
						let items = resp.items;
						items.forEach(item=>{
							navItem.orderList.push(item);
						})	
					}
					//loaded新字段用于表示数据加载完毕，如果为空可以显示空白页
					this.$set(navItem, 'loaded', true);
				}, 600);	
			}, 

			// 跳转到商品详情
			naveToProduct(gpid){
				uni.navigateTo({
					url: `/pages/product/product?gpid=` + gpid
				})
			},

			//swiper 切换
			changeTab(e){
				this.tabCurrentIndex = e.target.current;
				this.loadData('tabChange');
			},
			//顶部tab点击
			tabClick(index){
				this.tabCurrentIndex = index;
			},
			// 取消订单
			async cancelOrder1(billId,index){
				console.log(billId,index)
				uni.showLoading({
					title: '请稍后'
				})
				
				setTimeout(async ()=>{
					let resp = await this.$Order.DealOrder(billId,-1,"","",[])
					if (resp != null){
						console.log(resp)
						//取消订单后删除待付款中该项
						let list = this.navList[this.tabCurrentIndex].orderList;
						list[index].status = 3;
						list[index].stateTip = "已取消";
						list[index].checkStatus = -1;
						//this.navList[this.tabCurrentIndex].orderList[index] = list[index];
						uni.showToast({
							title: "成功取消",
							icon: 'success'
						})	
					}
					uni.hideLoading();
				}, 600)
			},
			// 取消订单
			async cancelOrder(billId,index){
				let cancelOrder1 = this.cancelOrder1;
				uni.showModal({
					title: '提示',
					content: '是否取消订单?',
					success: function (res) {
						if (res.confirm) {
							cancelOrder1(billId,index);
						} else if (res.cancel) {
							console.log('用户点击取消');
						}
					}
				});
			},
			// 查看物流
			async toLogistics(billId){
				uni.navigateTo({
					url: `/pages/logistics/logistics?billId=` + billId
				})
			},
			
			// 查看售后回复
			async toGetLogistics(billId){
					uni.showToast({
						title: "申请已达,请耐心等待",
						icon: 'success'
					})
			},
			// 支付订单
			async payOrder(billId,index){
					let resp = await this.$Order.Pay(billId);
					if (resp != null && resp.status == true){
						this.$Order.WxPay(resp.info)
					}
			},
			// 删除订单
			async deleteOrder(billId,index){
				uni.showLoading({
					title: '请稍后'
				})
				
				setTimeout(async ()=>{
					let resp = await this.$Order.DealOrder(billId,2,"","",[])
					if (resp != null){
						this.navList[this.tabCurrentIndex].orderList.splice(index, 1);
						uni.showToast({
							title: "已删除",
							icon: 'success'
						})	
					}
					uni.hideLoading();
				}, 600)
			},

			// 申请售后
			async afterSale(billId,index){
				uni.navigateTo({
					url: `/pages/feedback/feedback?billId=` + billId+"&index="+index
				})
			},

			feedback(index){
					let list = this.navList[this.tabCurrentIndex].orderList;
					list[index].status = 6;
					list[index].stateTip = "售后";
					list[index].checkStatus = 6;
			},


			//订单状态文字和颜色
			orderStateExp(state){
				let stateTip = '',
					stateTipColor = '#fa436a';
				switch(+state){
					case 1:
						stateTip = '待付款'; break;
					case 2:
						stateTip = '待发货'; break;
					case 9:
						stateTip = '订单已关闭'; 
						stateTipColor = '#909399';
						break;
						
					//更多自定义
				}
				return {stateTip, stateTipColor};
			}
		},
	}
</script>

<style lang="scss">
	page, .content{
		background: $page-color-base;
		height: 100%;
	}
	
	.swiper-box{
		height: calc(100% - 40px);
	}
	.list-scroll-content{
		height: 100%;
	}
	
	.navbar{
		display: flex;
		height: 40px;
		padding: 0 5px;
		background: #fff;
		box-shadow: 0 1px 5px rgba(0,0,0,.06);
		position: relative;
		z-index: 10;
		.nav-item{
			flex: 1;
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100%;
			font-size: 15px;
			color: $font-color-dark;
			position: relative;
			&.current{
				color: $base-color;
				&:after{
					content: '';
					position: absolute;
					left: 50%;
					bottom: 0;
					transform: translateX(-50%);
					width: 44px;
					height: 0;
					border-bottom: 2px solid $base-color;
				}
			}
		}
	}

	.uni-swiper-item{
		height: auto;
	}
	.order-item{
		display: flex;
		flex-direction: column;
		padding-left: 30upx;
		background: #fff;
		margin-top: 16upx;
		.i-top{
			display: flex;
			align-items: center;
			height: 80upx;
			padding-right:30upx;
			font-size: $font-base;
			color: $font-color-dark;
			position: relative;
			.time{
				flex: 1;
			}
			.state{
				color: $base-color;
			}
			.del-btn{
				padding: 10upx 0 10upx 36upx;
				font-size: $font-lg;
				color: $font-color-light;
				position: relative;
				&:after{
					content: '';
					width: 0;
					height: 30upx;
					border-left: 1px solid $border-color-dark;
					position: absolute;
					left: 20upx;
					top: 50%;
					transform: translateY(-50%);
				}
			}
		}
		/* 多条商品 */
		.goods-box{
			height: 160upx;
			padding: 20upx 0;
			white-space: nowrap;
			.goods-item{
				width: 120upx;
				height: 120upx;
				display: inline-block;
				margin-right: 24upx;
			}
			.goods-img{
				display: block;
				width: 100%;
				height: 100%;
			}
		}
						.title{
					font-size: $font-base + 2upx;
					color: $font-color-dark;
					line-height: 1;
				}
		/* 单条商品 */
		.goods-box-single{
			display: flex;
			padding: 20upx 0;
			.goods-img{
				display: block;
				width: 120upx;
				height: 120upx;
			}
			.right{
				flex: 1;
				display: flex;
				flex-direction: column;
				padding: 0 30upx 0 24upx;
				overflow: hidden;
				.title{
					font-size: $font-base + 2upx;
					color: $font-color-dark;
					line-height: 1;
				}
				.attr-box{
					font-size: $font-sm + 2upx;
					color: $font-color-light;
					padding: 10upx 12upx;
				}
				.price{
					font-size: $font-base + 2upx;
					color: $uni-color-primary;
					&:before{
						content: '￥';
						font-size: $font-sm;
						margin: 0 2upx 0 8upx;
					}
				}
			}
		}

		.order-box{
			font-size: $font-sm;
			color: $font-color-dark;
			padding: 4upx 0upx;
		}
		
		.price-box{
			display: flex;
			justify-content: flex-end;
			align-items: baseline;
			padding: 20upx 30upx;
			font-size: $font-sm + 2upx;
			color: $font-color-light;
			.num{
				margin: 0 8upx;
				color: $uni-color-primary;
			}
			.price{
				font-size: $font-lg;
				color: $uni-color-primary;
				&:before{
					content: '￥';
					font-size: $font-sm;
					margin: 0 2upx 0 8upx;
				}
			}
		}
		.action-box{
			display: flex;
			justify-content: flex-end;
			align-items: center;
			height: 100upx;
			position: relative;
			padding-right: 30upx;
		}
		.action-btn{
			width: 160upx;
			height: 60upx;
			margin: 0;
			margin-left: 24upx;
			padding: 0;
			text-align: center;
			line-height: 60upx;
			font-size: $font-sm + 2upx;
			color: $font-color-dark;
			background: #fff;
			border-radius: 100px;
			&:after{
				border-radius: 100px;
			}
			&.recom{
				background: #fff9f9;
				color: $base-color;
				&:after{
					border-color: #f7bcc8;
				}
			}
		}
	}
	
	
	/* load-more */
	.uni-load-more {
		display: flex;
		flex-direction: row;
		height: 80upx;
		align-items: center;
		justify-content: center
	}
	
	.uni-load-more__text {
		font-size: 28upx;
		color: #999
	}
	
	.uni-load-more__img {
		height: 24px;
		width: 24px;
		margin-right: 10px
	}
	
	.uni-load-more__img>view {
		position: absolute
	}
	
	.uni-load-more__img>view view {
		width: 6px;
		height: 2px;
		border-top-left-radius: 1px;
		border-bottom-left-radius: 1px;
		background: #999;
		position: absolute;
		opacity: .2;
		transform-origin: 50%;
		animation: load 1.56s ease infinite
	}
	
	.uni-load-more__img>view view:nth-child(1) {
		transform: rotate(90deg);
		top: 2px;
		left: 9px
	}
	
	.uni-load-more__img>view view:nth-child(2) {
		transform: rotate(180deg);
		top: 11px;
		right: 0
	}
	
	.uni-load-more__img>view view:nth-child(3) {
		transform: rotate(270deg);
		bottom: 2px;
		left: 9px
	}
	
	.uni-load-more__img>view view:nth-child(4) {
		top: 11px;
		left: 0
	}
	
	.load1,
	.load2,
	.load3 {
		height: 24px;
		width: 24px
	}
	
	.load2 {
		transform: rotate(30deg)
	}
	
	.load3 {
		transform: rotate(60deg)
	}
	
	.load1 view:nth-child(1) {
		animation-delay: 0s
	}
	
	.load2 view:nth-child(1) {
		animation-delay: .13s
	}
	
	.load3 view:nth-child(1) {
		animation-delay: .26s
	}
	
	.load1 view:nth-child(2) {
		animation-delay: .39s
	}
	
	.load2 view:nth-child(2) {
		animation-delay: .52s
	}
	
	.load3 view:nth-child(2) {
		animation-delay: .65s
	}
	
	.load1 view:nth-child(3) {
		animation-delay: .78s
	}
	
	.load2 view:nth-child(3) {
		animation-delay: .91s
	}
	
	.load3 view:nth-child(3) {
		animation-delay: 1.04s
	}
	
	.load1 view:nth-child(4) {
		animation-delay: 1.17s
	}
	
	.load2 view:nth-child(4) {
		animation-delay: 1.3s
	}
	
	.load3 view:nth-child(4) {
		animation-delay: 1.43s
	}
	
	@-webkit-keyframes load {
		0% {
			opacity: 1
		}
	
		100% {
			opacity: .2
		}
	}
</style>
