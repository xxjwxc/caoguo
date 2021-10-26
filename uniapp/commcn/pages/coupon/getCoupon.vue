<template>
	<view>
		<view class="container">
			<view v-if="empty===true" class=" empty">
				<view v-if="true" class="empty-tips">
					暂无更多优惠券
					<navigator class="navigator" url="myCoupon" >我的优惠券></navigator>
				</view>
			</view>
		</view>
		
		<view class="mask-content"  v-for="(list,index) in couponList" :key="index">
			<view class="yt-list-cell b-b">
				<view class="cell-icon">券</view>
				<text class="cell-tit clamp">{{list.key}}</text>
			</view>
			<!-- 优惠券页面，仿mt -->
			<view class="coupon-item" v-for="(item,index) in list.items" :key="index">
				<view class="con">
					<view class="left">
						<text class="t\itle">{{item.title}}</text>
						<text class="time">{{item.validity}}</text>
					</view>
					<view class="right">
						<text class="price">{{item.denom*0.01}}</text>
						<text>{{item.greatMoney}}</text>
					</view>

					<view class="circle l"></view>
					<view class="circle r"></view>
				</view>
				<text class="tips">{{item.describle}}</text>
			</view>
		</view>

		<!-- 底部 -->
		<view class="footer" v-if="empty===false">
			<view class="price-content">
				<text>优惠面值</text>
				<text class="price-tip">￥</text>
				<text class="price">{{(totalFee*0.01).toFixed(2)}}</text>
			</view>
			<text class="submit" @click="submit">一键领取</text>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				empty: false, //空白页现实  true|false
				totalFee:100,
				couponList: [
				],
				ids:[],
			}
		},
		onLoad(option) {
			this.OnInit();
		},
		watch: {
			//显示空白页
			couponList(e) {
				let empty = e.length === 0 ? true : false;
				if (this.empty !== empty) {
					this.empty = empty;
				}
			}
		},
		methods: {
			async OnInit() {
				let resp = await this.$Order.GetPromotionCoupon();
				if (typeof(resp.items) != "undefined" && resp.items != null) {
					this.couponList = resp.items;
					this.totalFee = resp.totalFee;
					this.ids = resp.ids;
				}
				this.empty = this.couponList.length === 0 ? true : false;
			},
			async submit() {
				let resp = await this.$Order.GoGetCoupon(this.ids);
				if (resp == true){
					this.$api.msg("添加成功，请到我的优惠券中查看");
					setTimeout(()=>{
						uni.navigateBack()
					}, 800)
				}
			}
		}
	}
</script>

<style lang="scss">
	page {
		background: $page-color-base;
		padding-bottom: 100upx;
	}

	.empty-tips {
		left: 0;
		top: 0;
		width: 100%;
		height: 100vh;
		padding-bottom: 100upx;
		justify-content: center;
		flex-direction: column;
		align-items: center;
		background: #fff;

		display: flex;
		font-size: $font-sm+2upx;
		color: $font-color-disabled;

		.navigator {
			color: $uni-color-primary;
			margin-left: 16upx;
		}
	}

	.yt-list {
		margin-top: 16upx;
		background: #fff;
	}

	.yt-list-cell {
		display: flex;
		align-items: center;
		padding: 10upx 30upx 10upx 40upx;
		line-height: 70upx;
		position: relative;

		&.cell-hover {
			background: #fafafa;
		}

		&.b-b:after {
			left: 30upx;
		}

		.cell-icon {
			height: 32upx;
			width: 32upx;
			font-size: 22upx;
			color: #fff;
			text-align: center;
			line-height: 32upx;
			background: #f85e52;
			border-radius: 4upx;
			margin-right: 12upx;

			&.hb {
				background: #ffaa0e;
			}

			&.lpk {
				background: #3ab54a;
			}
		}

		.cell-more {
			align-self: center;
			font-size: 24upx;
			color: $font-color-light;
			margin-left: 8upx;
			margin-right: -10upx;
		}

		.cell-tit {
			flex: 1;
			font-size: 26upx;
			color: $font-color-light;
			margin-right: 10upx;
		}

		.cell-tip {
			font-size: 26upx;
			color: $font-color-dark;

			&.disabled {
				color: $font-color-light;
			}

			&.active {
				color: $base-color;
			}

			&.red {
				color: $base-color;
			}
		}

		&.desc-cell {
			.cell-tit {
				max-width: 90upx;
			}
		}

		.desc {
			flex: 1;
			font-size: $font-base;
			color: $font-color-dark;
		}
	}

	.footer {
		position: fixed;
		left: 0;
		bottom: 0;
		z-index: 995;
		display: flex;
		align-items: center;
		width: 100%;
		height: 90upx;
		justify-content: space-between;
		font-size: 30upx;
		background-color: #fff;
		z-index: 998;
		color: $font-color-base;
		box-shadow: 0 -1px 5px rgba(0, 0, 0, .1);

		.price-content {
			padding-left: 30upx;
		}

		.price-tip {
			color: $base-color;
			margin-left: 8upx;
		}

		.price {
			font-size: 36upx;
			color: $base-color;
		}

		.submit {
			display: flex;
			align-items: center;
			justify-content: center;
			width: 280upx;
			height: 100%;
			color: #fff;
			font-size: 32upx;
			background-color: $base-color;
		}
	}

	/* 优惠券列表 */
	.coupon-item {
		display: flex;
		flex-direction: column;
		margin: 30upx 50upx;
		background: #f2aa3c;

		.con {
			display: flex;
			align-items: center;
			position: relative;
			height: 120upx;
			padding: 0 30upx;

			&:after {
				position: absolute;
				left: 0;
				bottom: 0;
				content: '';
				width: 100%;
				height: 0;
				border-bottom: 1px dashed #f5c76f;
				transform: scaleY(50%);
			}
		}

		.left {
			display: flex;
			flex-direction: column;
			justify-content: center;
			flex: 1;
			overflow: hidden;
			height: 100upx;
		}

		.title {
			font-size: 32upx;
			color: #ffffff;
			margin-bottom: 10upx;
		}

		.time {
			font-size: 24upx;
			color: $font-color-light;
		}

		.right {
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			font-size: 26upx;
			color: $font-color-base;
			height: 100upx;
		}

		.price {
			font-size: 44upx;
			color: $base-color;

			&:before {
				content: '￥';
				font-size: 34upx;
			}
		}

		.tips {
			font-size: 24upx;
			color: $font-color-light;
			line-height: 60upx;
			padding-left: 30upx;
		}

		.circle {
			position: absolute;
			left: -12upx;
			bottom: -20upx;
			z-index: 10;
			width: 40upx;
			height: 40upx;
			background: #f3f3f3;
			border-radius: 100px;

			&.r {
				left: auto;
				right: -12upx;
			}
		}
	}
</style>
