<template>
	<view class="container">
		<view class="carousel">
			<swiper indicator-dots circular=true duration="400">
				<swiper-item class="swiper-item" v-for="(item,index) in info.img" :key="index">
					<view class="image-wrapper">
						<image :src="item" class="loaded" mode="aspectFill"></image>
					</view>
				</swiper-item>
			</swiper>
		</view>

		<view class="introduce-section">
			<text class="title">{{info.name}}</text>
			<view class="price-box">
				<text class="price-tip">¥</text>
				<text class="price">{{(price * 0.01).toFixed(2)}}</text>
				<text class="m-price" v-if="info.originalPrice > 0">¥{{(info.originalPrice * 0.01).toFixed(2)}}</text>
				<text class="coupon-tip" v-if="info.originalPrice > 0">{{((price/info.originalPrice)*10).toFixed(1)}}折</text>
			</view>
			<view class="bot-row">
				<text>销量: {{info.sales}}</text>
				<text>库存: {{info.stock}}</text>
				<text>浏览量: {{info.views}}</text>
			</view>
		</view>

		<!--  分享 -->
		<view class="share-section" @click="share">
			<view class="share-icon">
				<text class="yticon icon-xingxing"></text>
				返
			</view>
			<text class="tit">{{info.shareTit}}</text>
			<text class="yticon icon-bangzhu1"></text>
			<view class="share-btn" open-type="share" catchtap="btn">
				<button class="share-btn" open-type="share" style="height: 52upx;margin-top: -25upx;">立即分享
					<text class="yticon icon-you"></text>
				</button>
			</view>
		</view>

		<view class="c-list">
			<view v-if="(skuList !== undefined && skuList !== null) && skuList.length > 0">
				<view class="c-row b-b" @click="buy_before(0)">
					<text class="tit">购买类型</text>
					<view class="con">
						<text class="selected-text" v-for="(sItem, sIndex) in specSelected" :key="sIndex">
							{{sItem.tagName}}
						</text>
					</view>
					<text class="yticon icon-you"></text>
				</view>
			</view>

			<view v-if="info.couponName != ''">
				<view class="c-row b-b" @click="naveToCouponList">
					<text class="tit">优惠券</text>
					<text class="con t-r red">{{info.couponName}}</text>
					<text class="yticon icon-you"></text>
				</view>
			</view>

			<view v-if="(info.conList !== undefined && info.conList !== null) && info.conList.length > 0">
				<view class="c-row b-b">
					<text class="tit">促销活动</text>
					<view class="con-list">
						<text v-for="(sItem, sIndex) in info.conList" :key="sIndex">
							{{sItem}}
						</text>
					</view>
				</view>
			</view>

			<view v-if="(info.bzList !== undefined && info.bzList !== null) && info.bzList.length > 0">
				<view class="c-row b-b">
					<text class="tit">服务</text>
					<view class="bz-list con">
						<text v-for="(sItem, sIndex) in info.bzList" :key="sIndex">
							{{sItem}}
						</text>
					</view>
				</view>
			</view>
		</view>

		<!-- 评价 -->
		<!-- 	<view class="eva-section">
			<view class="e-header">
				<text class="tit">评价</text>
				<text>(86)</text>
				<text class="tip">好评率 100%</text>
				<text class="yticon icon-you"></text>
			</view> 
			<view class="eva-box">
				<image class="portrait" src="http://img3.imgtn.bdimg.com/it/u=1150341365,1327279810&fm=26&gp=0.jpg" mode="aspectFill"></image>
				<view class="right">
					<text class="name">Leo yo</text>
					<text class="con">商品收到了，79元两件，质量不错，试了一下有点瘦，但是加个外罩很漂亮，我很喜欢</text>
					<view class="bot">
						<text class="attr">购买类型：XL 红色</text>
						<text class="time">2019-04-01 19:21</text>
					</view>
				</view>
			</view>
		</view>
		 -->
		<view class="detail-desc">
			<view class="d-header">
				<text>图文详情</text>
			</view>

			<div style="width:100%">
				<video style="width:100%;display:block;" v-for="(item, index) in info.vdDetail" :key="index" :src="item" controls="controls"
				 mode="widthFix"></video>
				<img style="width: 100%;display:block;" v-for="(item, index) in info.imgDetail" :key="index" :src="item" mode="widthFix"
				 @tap="_previewImage(item)" />
			</div>


			<rich-text :nodes="info.richText"></rich-text>
		</view>

		<!-- 底部操作菜单 -->
		<view class="page-bottom">
			<navigator url="/pages/index/index" open-type="switchTab" class="p-b-btn">
				<text class="yticon icon-xiatubiao--copy"></text>
				<text>首页</text>
			</navigator>
			<navigator url="/pages/cart/cart" open-type="switchTab" class="p-b-btn">
				<text class="yticon icon-gouwuche"></text>
				<text>购物车</text>
			</navigator>
			<view class="p-b-btn" :class="{active:favorite}" @click="toFavorite">
				<text class="yticon icon-shoucang"></text>
				<text>收藏</text>
			</view>

			<view class="action-btn-group">
				<button type="primary" class=" action-btn no-border add-cart-btn" @click="buy_before(2)">加入购物车</button>
				<button type="primary" class=" action-btn no-border buy-now-btn" @click="buy_before(1)">立即购买</button>
			</view>
		</view>


		<!-- 规格-模态层弹窗 -->
		<view class="popup spec" :class="specClass" @touchmove.stop.prevent="stopPrevent" @click="buy_before(0)">
			<!-- 遮罩层 -->
			<view class="mask"></view>
			<view class="layer attr-content" @click.stop="stopPrevent">
				<view class="a-t">
					<image :src="info.icon"></image>
					<view class="right">
						<text class="price">¥{{(price * 0.01).toFixed(2)}}</text>
						<text class="stock">库存：{{info.stock}}件</text>
						<view class="selected">
							已选：
							<text class="selected-text" v-for="(sItem, sIndex) in specSelected" :key="sIndex">
								{{sItem.tagName}}
							</text>
						</view>
					</view>
				</view>
				<view v-for="(item,index) in skuList" :key="index" class="attr-list">
					<text>{{item.typeName}}</text>
					<view class="item-list">
						<text v-for="(childItem, childIndex) in item.Items" :key="childIndex" class="tit" :class="{selected: childItem.selected}"
						 @click="selectSpec(index,childIndex, childItem.id)">
							{{childItem.tagName}}
						</text>
					</view>
				</view>

				<view class="item-list-number">
					<text>购买数量:</text>
					<uni-number-box :min="1" :max="100" :value="number" :isMax="false" :isMin="number===1" @eventChange="numberChange"></uni-number-box>
				</view>

				<button class="btn" @click="buy_after">{{buyTag}}</button>
			</view>
		</view>
		<!-- 分享 -->
		<share ref="share" :contentHeight="580" :shareList="shareList"></share>
	</view>
</template>

<script>
	import uniNumberBox from '@/components/uni-number-box-new.vue'
	import share from '@/components/share';
	export default {
		components: {
			share,
			uniNumberBox
		},
		data() {
			return {
				buyType: 0,
				buyTag: "完成",
				number: 1,
				specClass: 'none',
				specSelected: [],

				favorite: false,
				shareList: [],

				price: 0,

				info: {},
				skuList: [],
				skuPriceList: [],
				sku: []
			};
		},
		async onLoad(options) {
			// #ifdef MP-WEIXIN
			uni.showShareMenu({
				withShareTicket: true,
				menus: ['shareAppMessage', 'shareTimeline']
			})
			//#endif

			await this.$User.WxLogin(); // 登录
			//接收传值,id里面放的是标题，因为测试数据并没写id 
			let gpid = options.gpid;
			let openid = options.openid;

			if (!gpid) {
				let tmp = decodeURIComponent(options.scene).split("@")
				gpid = tmp[0]
				openid = tmp[1]
			}
			console.log(openid)
			if (typeof(openid) != "undefined" && openid != "undefined" && openid != null) { // 建立关系
				this.$User.LinkUser(openid);
			}


			if (!gpid || gpid == "undefined") {
				uni.showLoading({
					title: '请稍后重试..',
					mask: true,
				});

				setTimeout(() => {
					uni.hideLoading();
					uni.navigateTo({
						url: `/pages/index/index`
					})
				}, 800)
			}

			let info = await this.$Product.GetProductDetails(gpid);
			if ((typeof(info.info.sku.groups) != "undefined" && info.info.sku.groups != null) && info.info.sku.groups != null) {
				this.skuList = info.info.sku.groups;
				this.skuPriceList = info.info.skuPrice.list;
			}
			if ((typeof(info.info.skuPrice.list) != "undefined" && info.info.skuPrice.list != null) && info.info.skuPrice.list !=
				null) {
				this.skuPriceList = info.info.skuPrice.list;
			}

			this.info = info.info;
			this.price = 0; //this.info.price;
			this.favorite = this.info.isFavorite;


			//规格 默认选中第一条
			for (let item of this.skuList) {
				for (let cItem of item.Items) {
					this.$set(cItem, 'selected', true);
					this.specSelected.push(cItem);
					this.sku.push(cItem.id);
					break; //forEach不能使用break
				}
			}

			this.sku.sort(function(a, b) { // 升序排序
				return a - b;
			});
			var key = this.sku.join(",")
			for (let item of this.skuPriceList) {
				if (item.skuList === key) {
					this.price += item.premium;
					break;
				}
			};
		},
		onShareAppMessage() {
			return {
				title: "潞潮公社-" + this.info.name,
				path: "/pages/product/product?gpid=" + this.info.gpid + "&openid=" + this.$Server.OpenID,
			}
		},

		// 分享到朋友圈
		onShareTimeline() {
			return {
				title: "潞潮公社-" + this.info.name,
				imageUrl: this.info.icon
			}
		},

		methods: {
			//数量
			numberChange(data) {
				this.number = data.number;
				console.log(this.number);
			},
			//规格弹窗开关
			toggleSpec() {
				if (this.specClass === 'show') {
					this.specClass = 'hide';
					setTimeout(() => {
						this.specClass = 'none';
					}, 250);
				} else if (this.specClass === 'none') {
					this.specClass = 'show';
				}
			},
			//选择规格
			selectSpec(index, childIndex, id) {
				let list = this.skuList;
				for (let cItem of list[index].Items) {
					this.$set(cItem, 'selected', false);
				}
				this.$set(list[index].Items[childIndex], 'selected', true); // 设置选中


				let _price = 0; // this.info.price;
				//存储已选择
				/**
				 * 修复选择规格存储错误
				 * 将这几行代码替换即可
				 * 选择的规格存放在specSelected中
				 */
				this.specSelected = [];
				this.sku = [];
				list.forEach(item => { // reset
					for (let cItem of item.Items) {
						if (cItem.selected === true) {
							this.specSelected.push(cItem);
							this.sku.push(cItem.id);
						}
					}
				});

				this.sku.sort(function(a, b) { // 升序排序
					return a - b;
				});
				let key = this.sku.join(",")
				for (let item of this.skuPriceList) {
					if (item.skuList === key) {
						_price += item.premium;
						break;
					}
				}

				this.price = _price;
			},
			//收藏
			toFavorite() {
				this.favorite = !this.favorite;
				this.$Product.ToFavorite(this.info.gpid, this.favorite);
			},
			// 加入购物车
			async addCart() {
				let isSuccess = await this.$Product.AddToCart(this.info.gpid, this.number, this.sku);
				if (isSuccess) {
					uni.showToast({
						title: "成功加入购物车",
						icon: 'success'
					});
				} else {
					uni.showToast({
						title: "加入购物车失败",
						icon: 'none'
					});
				}
			},
			//分享
			share() {
				// #ifndef MP-WEIXIN
				this.$refs.share.toggleMask();
				// #endif
			},
			// 跳转到优惠券领取平台
			naveToCouponList() {
				uni.navigateTo({
					url: `/pages/coupon/getCoupon`
				})
			},
			buy_before(tag) {
				if (tag == 0) {
					this.buyType = 0
					this.buyTag = "完成";
					this.toggleSpec();
				}
				if (tag == 1) { // 立即购买
					this.buyType = 1
					this.buyTag = "立即购买";
					this.toggleSpec();
				}
				if (tag == 2) { // 加入购物车
					this.buyType = 2
					this.buyTag = "加入购物车";
					this.toggleSpec();
				}
			},
			async buy_after() {
				this.toggleSpec();
				if (this.buyType == 0) { // 关闭遮罩
				}
				if (this.buyType == 1) { // 立即购买
					this.buy();
				}
				if (this.buyType == 2) { // 加入购物车
					this.addCart();
				}
			},
			async buy() {
				// 添加到虚拟购物车
				let cartTmp = await this.$Product.AddToBuyTmpCart(this.info.gpid, this.number, this.sku);
				let id = cartTmp.id;
				if (id > 0) { // 加入购物车成功
					uni.navigateTo({
						url: `/pages/order/createOrder?type=1&ids=` + id
					})
				} else {
					uni.showToast({
						title: "商品内部错误，请稍后重试",
						icon: 'none'
					});
				}

			},
			stopPrevent() {},
			_previewImage(image) {// 预览图片
				var imgArr = [];
				imgArr.push(image);
				//预览图片
				uni.previewImage({
					urls: imgArr,
					current: imgArr[0]
				});
			}
		},
	}
</script>

<style lang='scss'>
	page {
		background: $page-color-base;
		padding-bottom: 160upx;
	}

	.icon-you {
		font-size: $font-base + 2upx;
		color: #888;
	}

	.carousel {
		height: 722upx;
		position: relative;

		swiper {
			height: 100%;
		}

		.image-wrapper {
			width: 100%;
			height: 100%;
		}

		.swiper-item {
			display: flex;
			justify-content: center;
			align-content: center;
			height: 750upx;
			overflow: hidden;

			image {
				width: 100%;
				height: 100%;
			}
		}

	}

	button::after {
		border: none;
	}

	/* 标题简介 */
	.introduce-section {
		background: #fff;
		padding: 20upx 30upx;

		.title {
			font-size: 32upx;
			color: $font-color-dark;
			height: 50upx;
			line-height: 50upx;
		}

		.price-box {
			display: flex;
			align-items: baseline;
			height: 64upx;
			padding: 10upx 0;
			font-size: 26upx;
			color: $uni-color-primary;
		}

		.price {
			font-size: $font-lg + 2upx;
		}

		.m-price {
			margin: 0 12upx;
			color: $font-color-light;
			text-decoration: line-through;
		}

		.coupon-tip {
			align-items: center;
			padding: 4upx 10upx;
			background: $uni-color-primary;
			font-size: $font-sm;
			color: #fff;
			border-radius: 6upx;
			line-height: 1;
			transform: translateY(-4upx);
		}

		.bot-row {
			display: flex;
			align-items: center;
			height: 50upx;
			font-size: $font-sm;
			color: $font-color-light;

			text {
				flex: 1;
			}
		}
	}

	/* 分享 */
	.share-section {
		display: flex;
		align-items: center;
		color: $font-color-base;
		background: linear-gradient(left, #fdf5f6, #fbebf6);
		padding: 12upx 30upx;

		.share-icon {
			display: flex;
			align-items: center;
			width: 70upx;
			height: 30upx;
			line-height: 1;
			border: 1px solid $uni-color-primary;
			border-radius: 4upx;
			position: relative;
			overflow: hidden;
			font-size: 22upx;
			color: $uni-color-primary;

			&:after {
				content: '';
				width: 50upx;
				height: 50upx;
				border-radius: 50%;
				left: -20upx;
				top: -12upx;
				position: absolute;
				background: $uni-color-primary;
			}
		}

		.icon-xingxing {
			position: relative;
			z-index: 1;
			font-size: 24upx;
			margin-left: 2upx;
			margin-right: 10upx;
			color: #fff;
			line-height: 1;
		}

		.tit {
			font-size: $font-base;
			margin-left: 10upx;
		}

		.icon-bangzhu1 {
			padding: 10upx;
			font-size: 30upx;
			line-height: 1;
		}

		.share-btn {
			flex: 1;
			text-align: right;
			font-size: $font-sm;
			background: rgba(0, 0, 0, 0);
			color: $uni-color-primary;
		}

		.icon-you {
			font-size: $font-sm;
			margin-left: 4upx;
			color: $uni-color-primary;
		}
	}

	.c-list {
		font-size: $font-sm + 2upx;
		color: $font-color-base;
		background: #fff;

		.c-row {
			display: flex;
			align-items: center;
			padding: 20upx 30upx;
			position: relative;
		}

		.tit {
			width: 140upx;
		}

		.con {
			flex: 1;
			color: $font-color-dark;

			.selected-text {
				margin-right: 10upx;
			}
		}

		.bz-list {
			height: 40upx;
			font-size: $font-sm+2upx;
			color: $font-color-dark;

			text {
				display: inline-block;
				margin-right: 30upx;
			}
		}

		.con-list {
			flex: 1;
			display: flex;
			flex-direction: column;
			color: $font-color-dark;
			line-height: 40upx;
		}

		.red {
			color: $uni-color-primary;
		}
	}

	/* 评价 */
	.eva-section {
		display: flex;
		flex-direction: column;
		padding: 20upx 30upx;
		background: #fff;
		margin-top: 16upx;

		.e-header {
			display: flex;
			align-items: center;
			height: 70upx;
			font-size: $font-sm + 2upx;
			color: $font-color-light;

			.tit {
				font-size: $font-base + 2upx;
				color: $font-color-dark;
				margin-right: 4upx;
			}

			.tip {
				flex: 1;
				text-align: right;
			}

			.icon-you {
				margin-left: 10upx;
			}
		}
	}

	.eva-box {
		display: flex;
		padding: 20upx 0;

		.portrait {
			flex-shrink: 0;
			width: 80upx;
			height: 80upx;
			border-radius: 100px;
		}

		.right {
			flex: 1;
			display: flex;
			flex-direction: column;
			font-size: $font-base;
			color: $font-color-base;
			padding-left: 26upx;

			.con {
				font-size: $font-base;
				color: $font-color-dark;
				padding: 20upx 0;
			}

			.bot {
				display: flex;
				justify-content: space-between;
				font-size: $font-sm;
				color: $font-color-light;
			}
		}
	}

	/*  详情 */
	.detail-desc {
		background: #fff;
		margin-top: 16upx;

		.d-header {
			display: flex;
			justify-content: center;
			align-items: center;
			height: 80upx;
			font-size: $font-base + 2upx;
			color: $font-color-dark;
			position: relative;

			text {
				padding: 0 20upx;
				background: #fff;
				position: relative;
				z-index: 1;
			}

			&:after {
				position: absolute;
				left: 50%;
				top: 50%;
				transform: translateX(-50%);
				width: 300upx;
				height: 0;
				content: '';
				border-bottom: 1px solid #ccc;
			}
		}
	}

	/* 规格选择弹窗 */
	.attr-content {
		padding: 10upx 30upx;

		.a-t {
			display: flex;

			image {
				width: 170upx;
				height: 170upx;
				flex-shrink: 0;
				margin-top: -40upx;
				border-radius: 8upx;
				;
			}

			.right {
				display: flex;
				flex-direction: column;
				padding-left: 24upx;
				font-size: $font-sm + 2upx;
				color: $font-color-base;
				line-height: 42upx;

				.price {
					font-size: $font-lg;
					color: $uni-color-primary;
					margin-bottom: 10upx;
				}

				.selected-text {
					margin-right: 10upx;
				}
			}
		}

		.attr-list {
			display: flex;
			flex-direction: column;
			font-size: $font-base + 2upx;
			color: $font-color-base;
			padding-top: 30upx;
			padding-left: 10upx;
		}

		.item-list {
			padding: 20upx 0 0;
			display: flex;
			flex-wrap: wrap;

			text {
				display: flex;
				align-items: center;
				justify-content: center;
				background: #eee;
				margin-right: 20upx;
				margin-bottom: 20upx;
				border-radius: 100upx;
				min-width: 60upx;
				height: 60upx;
				padding: 0 20upx;
				font-size: $font-base;
				color: $font-color-dark;
			}

			.selected {
				background: #fbebee;
				color: $uni-color-primary;
			}
		}

		.item-list-number {
			padding: 20upx 0 0;
			display: flex;
			flex-wrap: wrap;

			text {
				display: flex;
				align-items: center;
				justify-content: center;
				margin-right: 20upx;
				margin-bottom: 20upx;
				border-radius: 100upx;
				min-width: 60upx;
				height: 60upx;
				padding: 0 20upx;
				font-size: $font-base;
				color: $font-color-dark;
			}

			.selected {
				background: #fbebee;
				color: $uni-color-primary;
			}
		}
	}

	/*  弹出层 */
	.popup {
		position: fixed;
		left: 0;
		top: 0;
		right: 0;
		bottom: 0;
		z-index: 99;

		&.show {
			display: block;

			.mask {
				animation: showPopup 0.2s linear both;
			}

			.layer {
				animation: showLayer 0.2s linear both;
			}
		}

		&.hide {
			.mask {
				animation: hidePopup 0.2s linear both;
			}

			.layer {
				animation: hideLayer 0.2s linear both;
			}
		}

		&.none {
			display: none;
		}

		.mask {
			position: fixed;
			top: 0;
			width: 100%;
			height: 100%;
			z-index: 1;
			background-color: rgba(0, 0, 0, 0.4);
		}

		.layer {
			position: fixed;
			z-index: 99;
			bottom: 0;
			width: 100%;
			min-height: 40vh;
			border-radius: 10upx 10upx 0 0;
			background-color: #fff;

			.btn {
				height: 66upx;
				line-height: 66upx;
				border-radius: 100upx;
				background: $uni-color-primary;
				font-size: $font-base + 2upx;
				color: #fff;
				margin: 30upx auto 20upx;
			}
		}

		@keyframes showPopup {
			0% {
				opacity: 0;
			}

			100% {
				opacity: 1;
			}
		}

		@keyframes hidePopup {
			0% {
				opacity: 1;
			}

			100% {
				opacity: 0;
			}
		}

		@keyframes showLayer {
			0% {
				transform: translateY(120%);
			}

			100% {
				transform: translateY(0%);
			}
		}

		@keyframes hideLayer {
			0% {
				transform: translateY(0);
			}

			100% {
				transform: translateY(120%);
			}
		}
	}

	/* 底部操作菜单 */
	.page-bottom {
		position: fixed;
		left: 30upx;
		bottom: 30upx;
		z-index: 95;
		display: flex;
		justify-content: center;
		align-items: center;
		width: 690upx;
		height: 100upx;
		background: rgba(255, 255, 255, .9);
		box-shadow: 0 0 20upx 0 rgba(0, 0, 0, .5);
		border-radius: 16upx;

		.p-b-btn {
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			font-size: $font-sm;
			color: $font-color-base;
			width: 96upx;
			height: 80upx;

			.yticon {
				font-size: 40upx;
				line-height: 48upx;
				color: $font-color-light;
			}

			&.active,
			&.active .yticon {
				color: $uni-color-primary;
			}

			.icon-fenxiang2 {
				font-size: 42upx;
				transform: translateY(-2upx);
			}

			.icon-shoucang {
				font-size: 46upx;
			}
		}

		.action-btn-group {
			display: flex;
			height: 76upx;
			border-radius: 100px;
			overflow: hidden;
			box-shadow: 0 20upx 40upx -16upx #fa436a;
			box-shadow: 1px 2px 5px rgba(219, 63, 96, 0.4);
			background: linear-gradient(to right, #ffac30,#ffac30 50%, #fa436a 50%, #fa436a);
			margin-left: 20upx;
			position: relative;

			// &:after {
			// 	content: '';
			// 	position: absolute;
			// 	top: 50%;
			// 	right: 50%;
			// 	transform: translateY(-50%);
			// 	height: 28upx;
			// 	width: 0;
			// 	border-right: 1px solid rgba(255, 255, 255, .5);
			// }

			.action-btn {
				display: flex;
				align-items: center;
				justify-content: center;
				width: 180upx;
				height: 100%;
				font-size: $font-base;
				padding: 0;
				border-radius: 0;
				background: transparent;
			}
		}
	}
</style>
