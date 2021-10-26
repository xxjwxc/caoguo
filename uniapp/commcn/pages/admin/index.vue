<template>
	<view class="content">
		<view class="titleNview-placing"></view>
		<view class="navbar">
			<view v-for="(item, index) in navList" :key="index" 
				class="nav-item" 
				:class="{current: tabCurrentIndex === index}"
				@click="tabClick(index)" >
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
						</view>
						
						<view v-if="true" >
							<view class="goods-box-single" v-for="(goodsItem, goodsIndex) in item.items" :key="goodsIndex">
								<image class="goods-img" :src="goodsItem.icon" mode="aspectFill" @click="naveToProduct(goodsItem.gpid)"></image>
								<view class="right">
									<text class="title clamp">{{goodsItem.name}}</text>
									<text class="attr-box" >{{goodsItem.skuVal}}  x {{goodsItem.number}}</text>
									<text class="price">{{goodsItem.price*0.01}}</text>
								</view>
							</view> 
							<text class="title clamp" selectable = "true"> 订 单 号 : {{item.billId}}</text>
							<text class="title clamp" selectable = "true"> 优惠券 : {{item.couponTitle}}</text>
						</view>
						<div style="border:1px solid #CCC"> </div>
						<div style="border:1px solid #CCC"> 
							</br><text class="title clamp" selectable = "true"> 收 货 人 : {{item.addr.name}}</text>
							</br><text class="title clamp" selectable = "true"> 手 机 号 : {{item.addr.mobile}}</text>
							</br><text selectable = "true" style="word-break:break-all;color:#F76260"> 收货地址 : {{item.addr.address}}</text>
							</br>
						</div>

						<view class="price-box">
							共 <text class="num">{{item.number}}</text>件商品
						</view>
						
						<text class="title" selectable = "true" v-for="(shipmentInfo, index) in item.shipmentInfo" :key="index">
							{{shipmentInfo.shipmentName}} : {{shipmentInfo.shipmentId}}
						</text>
						
						<!-- 待发货 -->
						<view class="action-box b-t" v-if="item.checkStatus == 2">
							<button class="action-btn" @click="addShipment(item.billId,index)">添加快递单</button>
							<button class="action-btn" @click="scanShipment(item.billId,index)">扫快递单</button>
							<button class="action-btn"  open-type="contact" @click="afterMsg(item.billId,index)">售后信息</button>
						</view>
						<view class="action-box b-t" v-if="item.checkStatus == 4">
							<button class="action-btn" @click="toLogistics(item.billId)">查看物流</button>
							<button class="action-btn" @click="addShipment(item.billId,index)">添加快递单</button>
							<button class="action-btn" @click="scanShipment(item.billId,index)">扫快递单</button>
							<button class="action-btn"  open-type="contact" @click="afterMsg(item.billId,index)">售后信息</button>
						</view>
						<view class="action-box b-t" v-if="item.checkStatus == 6">
							<button class="action-btn"  open-type="contact" @click="afterMsg(item.billId,index)">售后信息</button>
						</view>
					</view>

					<uni-load-more :status="tabItem.loadingType"></uni-load-more>

				</scroll-view>
			</swiper-item>
		</swiper>
		
		<prompt :isMutipleLine="false" :visible.sync="promptVisible" inputStyle="" title="添加运单号" placeholder="请输入运单号" :defaultValue="shipmentId"
		  mainColor="#e74a39" @confirm="clickPromptConfirm">
		  <!-- 这里放入slot内容-->
		  <view>
			<view @tap="toOpen">{{postName}}</view>
			<jpSelect ref="jpSelect" :list="list" @checked="checked" :item="item" select="radio" tite="请选择快递公司"></jpSelect>
		  </view>
		</prompt>
	</view>
	
</template> 

<script>
	import uniLoadMore from '@/components/uni-load-more/uni-load-more.vue';
	import empty from "@/components/empty";
	 import Prompt from '@/components/zz-prompt/index.vue'
	 import jpSelect from '@/components/jp-select/jp-select.vue';
	export default {
		components: {
			uniLoadMore,
			empty,
			Prompt,
			jpSelect
		},
		data() {
			return {
				search:"",
				index :0,
				shipmentId:"",
				postKey:"",
				postName:"请选择快递公司",
				tabCurrentIndex: 0,
				navList: [],
				shipmentBill:"",
				// 控制弹框输入框显示, 在需要的地方调用this.promptVisible = true
				promptVisible: false,
				list: []
			};
		},
		//点击导航栏 buttons 时触发
		onNavigationBarButtonTap(e) {
			const index = e.index;
			if (index === 0) {
				uni.removeStorageSync("access_token");
				uni.navigateTo({
					url: `/pages/public/login`
				})
			} 
		},
		onLoad(options){
			var that = this;
			// 判断是否登录
			uni.getStorage({
				key:"access_token",
				success(e){
					let resp = JSON.parse(e.data);
					var nowTime = ((new Date()).getTime())
					if (resp.expire_time < nowTime){
						that.RefreshToken(resp.refresh_token);
					}else{
						that.$Admin.SetAccessToken(resp.access_token);
						that.onInit();
					}
				},
				fail(e){ 
					uni.navigateTo({
						url: `/pages/public/login`
					})	
				}
			})
			return;
		},
		onShow(){
			
		},
		methods: {
			checked(el) {
				this.postKey = el.id;
				this.postName = el.name;
				console.log(el,name);
			},
			toOpen() {
			    this.$refs.jpSelect.toOpen()
			},
			// 刷新token
			async RefreshToken(refresh_token){
				let resp =await this.$Admin.RefreshToken(refresh_token);
				console.log(resp)
				if (typeof(resp) == "undefined" || resp == null){ // 失败
					uni.navigateTo({
						url: "/pages/public/login"
					})
					return
				}else{
					this.$Admin.SetAccessToken(resp.access_token);
					this.onInit();
				}
			},
			async onInit(){
				let resp = await this.$Order.GetMyOrdersConfig(true);
				// this.navList = resp.navList;
				resp.navList.forEach(item=>{
					this.$set(item, "orderList", []);
					this.navList.push(item);
				})

				let shipment = await this.$Admin.GetShipmentPost();
				// this.navList = resp.navList;
				shipment.items.forEach(item=>{
					this.list.push(item);
				})
				

				 this.loadData();
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
					pageNumber = 0;
					navItem.orderList = [];
					// return;
				}
				if(navItem.loadingType === 'loading'){
					//防止重复加载
					
					return;
				}

				navItem.loadingType = 'loading';
				setTimeout(async ()=>{
					let resp = await this.$Admin.GetBillInfoByAdmin(state,pageNumber,this.search);
					if (resp == null){ // 需要重新登录
						uni.navigateTo({
							url: "/pages/public/login"
						})
						return;
					}
					if (typeof(resp.items) == "undefined" || resp.items == null){
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
			// 手动输入快递单
			async addShipment(billId,index){
				console.log(billId);
				this.shipmentBill = billId;
				this.index = index;
				this.promptVisible = true;
			},
			/**
			 * 点击弹出输入框确定
			 */
			async clickPromptConfirm(val) {
			  this.promptVisible = false;
			  this.shipmentId = val;
			  if (this.postKey.length == 0){
				this.promptVisible = true
				uni.showToast({
					title: "请选择快递公司",
					icon: 'none'
				})
				return;
			  }
			  
			  // 添加运单号
			  let rest = await this.$Admin.AddBillShipment(this.shipmentBill,this.shipmentId,this.postKey);
			  if (rest){
				uni.showToast({
					title: "添加成功",
					icon: 'success'
				})
				
				let list = this.navList[this.tabCurrentIndex].orderList;
				list[this.index].shipmentInfo.push({"shipmentId":this.shipmentId,"shipmentName":this.postName});

				list[this.index].status = "4"
				list[index].stateTip = "待收货";
				list[index].checkStatus = -1;
			  }

			},
			// 扫码快递单
			async scanShipment(billId,index){
				let that = this;
				this.shipmentBill = billId;
				this.index = index;
				uni.scanCode({
				    success: function (res) {
						that.shipmentId = res.result;
						that.promptVisible = true;
				        // console.log('条码类型：' + res.scanType);
				        // console.log('条码内容：' + res.result);
				    }
				});
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

			// 查看物流
			async toLogistics(billId){
				uni.navigateTo({
					url: `/pages/logistics/logistics?billId=` + billId
				})
			},

			// 售后信息
			async afterMsg(billId,index){
				uni.navigateTo({
					url: `/pages/admin/afterMsg?billId=` + billId
				})
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
			},
		
			// 标题栏input搜索框点击
			onNavigationBarSearchInputConfirmed: async function(e) {
				this.search = e.text;
				this.loadData('tabChange');
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
	.titleNview-placing {
		height: var(--status-bar-height);			
		padding-top: 44px;
		box-sizing: content-box;
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
					font-size: $font-sm + 10upx;
					color: red;
					padding: 10upx 12upx;
				}
				.price{
					font-size: $font-base + 2upx;
					color: $font-color-dark;
					&:before{
						content: '￥';
						font-size: $font-sm;
						margin: 0 2upx 0 8upx;
					}
				}
			}
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
				color: $font-color-dark;
			}
			.price{
				font-size: $font-lg;
				color: $font-color-dark;
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
