<template>
	<view class="page">
		<!-- 圆通快递 -->
		<view class="express">
			<view class="top ali-c">
				<view class="picture"> 
					<image class="img" mode="aspectFill" :src="wlInfo.logo"></image>
				</view>
				<view class="text">
					<view class="text1">{{wlInfo.postName}}</view>
					<view class="text2">官方电话 {{wlInfo.expPhone}}</view>
				</view>
			</view>
			<view class="right">
				<text class="text" selectable = "true"> {{wlInfo.postName}} {{wlInfo.postNo}}</text>
				<view class="text">订单号:　{{ wlInfo.billId }} </view>
			</view>
		</view>

		<!-- 收货地址 -->
		<view class="content bgf">
			<view>
				<view class="flex list">
					<view class="time"></view>
					<view class="info flex1">
						<view class="title address">[收货地址]{{wlInfo.addr}}</view>
					</view>
				</view>
				<view class="flex list" :class="{one: index == 0 && wlInfo.deliveryStatus == 3}" v-for="(item, index) in wlInfo.list"
				 :key="index">
					<view class="time">
						<view class="day">{{item.timeArr[0]}}</view>
						<view>{{item.timeArr[1]}}</view>
					</view>
					<view class="info flex1">
						<view class="title">{{index == 0 && wlInfo.deliveryStatus == 3 ? '已签收' : '配送中'}}</view>
						<view class="text">{{item.context}}</view>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		props: {
			wlInfo: {
				type: Object,
				default: () => ({})
			}
		},
		data() {
			return {}
		},

		onLoad() {

		},

		methods: {

		}
	}
</script>

<style lang="scss" scoped>
	.page {
		height: 100vh;
		/* #ifdef H5 */
		height: calc(100vh - 44px);
		/* #endif */
		font-size: 24rpx;
		background-color: #F6F7F8;
		overflow-y: auto;
	}

	/*flex 转换成flex容器*/
	.flex {
		display: flex;
		flex-direction: row;
	}

	/*flex1 自动填充*/
	.flex1 {
		flex: 1;
	}

	/*ali-c 竖直居中*/
	.ali-c {
		display: flex;
		flex-direction: row;
		align-items: center;
	}

	.bgf {
		background-color: #fff;
	}

	.express {
		//圆通快递
		background-color: #ffffff;
		border-radius: 20rpx;
		margin: 30rpx 20rpx 30rpx 20rpx;

		.top {
			padding: 28rpx 29rpx 25rpx 29rpx;

			.img {
				display: block;
				width: 88rpx;
				height: 88rpx;
				border-radius: 50%;
			}

			.text {
				margin-left: 20rpx;

				.text1 {
					margin-bottom: 6rpx;
					font-size: 28rpx;
					color: #000000;
				}

				.text2 {
					font-size: 20rpx;
					color: #000000;
				}
				selectable:true;
			}
		}

		.right {
			background-color: rgba(8, 175, 254, 0.02);
			border-radius: 0rpx 0rpx 20rpx 20rpx;
			padding: 15rpx 30rpx 20rpx 30rpx;

			.text {
				color: #666666;
				font-size: 22rpx;
			}
		}
	}

	//收货地址
	.content {
		margin: 20rpx;
		padding: 56rpx 46rpx 56rpx 5rpx;
		border-radius: 20rpx;

		.list {
			&:first-child {
				.info::before {
					bottom: -20rpx;
					margin-top: 40rpx;
					border-left: 1px dashed #e5e5e5;
				}

				.title {
					margin-bottom: 56rpx;

					&::before {
						width: 46rpx;
						height: 46rpx;
						left: -23rpx;
						background: url('data:image/png;base64,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') no-repeat center;
						background-size: 46rpx 46rpx;
					}
				}
			}

			&:last-child {
				.info::before {
					height: 32rpx;
				}
			}

			&.one {
				.info::before {
					margin-top: 20rpx;
				}

				.title {
					color: #1a1a1a;

					&::before {
						width: 46rpx;
						height: 46rpx;
						left: -23rpx;
						background: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAC4AAAAuCAYAAABXuSs3AAABS2lUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4KPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iQWRvYmUgWE1QIENvcmUgNS42LWMxNDIgNzkuMTYwOTI0LCAyMDE3LzA3LzEzLTAxOjA2OjM5ICAgICAgICAiPgogPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIi8+CiA8L3JkZjpSREY+CjwveDp4bXBtZXRhPgo8P3hwYWNrZXQgZW5kPSJyIj8+nhxg7wAABdFJREFUaIHFmX1MVXUYxz+/Cwj5AogQaGrE1SXKVWmVpqJhmqm45ebcSpvrZRUOm72tP1rq3Nya2uyFpJc159T+0JpNTS2VRdoL5WsQYsJtCoEkIoooIJfTH79zL+ec+3bOuRf7bmzc3+95fs/3d87zPL/n+R2RsLuHCJEMTFf/xgMjgbuBAep8O/AvcBH4AzgKlAGtkRgVNonHA4uApcBsIMaivgc4DGwDvgI6rRKwSvwuoBB4DbjHqrEgaADeA0qAW2aVHBYMzAUqVCPRIg0wTF2zQrVhCmaIJwCbgf2A0xY1c3CqNjarNkMiHPF0ZDAVRs7LNAqBn1TbQRGKeBZwDHgwiqTM4gHV9n3BBIIRHw4cAUb1ASmzGAWUqlz8EIh4f2APkNl3nEwjE8mlv3EiEPH3gdxIrC0cBl9PFiwcFskqPuQCHxgHjXl8HvCtXQuuRFg9VjA/o3ds3yVYW6VQcd3uqj7MR2YdQE88AajERspLjoO3xwiKQmgW18K6aoXW21ZX98ENjAM6QO8qhdggXZgFlbP1pG/3wJ5G6NK8zCKnlCvMsscameV8adn7xOOROzLtlXMzYHW2YEKSfnxXPaw5q+Buh6wBUmaxIS+cuQZrqhQONlkm36huoMNLfAmw3YzmmEGSzJOGLZa3wNqzCqWX/XUeTYN3xgimDNGP726QOtVtlsgvBXZ4iR8E5oSSTukHb4wWvDpaP153S/ru1gsmLI6UsZBpSG6bzsPG8wotXaaIfwc8IRJ29yQDzYQpTUvzBI9onljNDdhUo7DtovRps4gVsNwJS0YIxmvc7Ocr8NhRxcwSHiA1FpgRjjRAhqHsaemCnfXWSAN0K/BhDaTEoSNuXD8EYoAZDmCaGenr3frfD6fA5QLBepdgUKxpo0wbAj9MF7x1v2F9a2kyzwFkW9EwvswVTqiYJXj23tB66QnwSa7gUJ5gUoolkoGQ7QBGhxXTwKPA53+Du11PanOu4IghDrzwbm6ZZnNt3fJU9Zhyaz+McgAZYcU0iBWw7pzCxCMKWwyZZMoQGcQfTxSkxUt3+jXf350ONoHrsMLKMwoxwhbxVAeQaFVrQpIMyuWnFOYcUzh+VT//XCZUPy4om64/oM61wVO/KSz8RaGpQ54JNjHQSs/pQ7cmk/zYDHllCm9WKLRpAri/Jk/d8sDqKvmWvmnoHVfsuQkgaxXLdVsgvyyula9/R11vACvAl3VyfP1f/jpd9q90bjiAS1a1rgQ54Zo64IUTCp0e+bvTA8+fUPgnyKVDp33izQ7gvBnJtPje/9e79KeeFqnx+ieeGh9YbsFQKMntjcxgckFQ4wCqzUie0ARgfhqU5ws2uARJcXq5GAFeOgKZhbQYlwi7Jgt2ThKM06SFU9Yu5M46kNcPYbG4XKavZs1lWZETKmcJXtT04t09+FJcrAPa1YBNjIV3cwTHZwoKNAm4uRNWnlFYXG4pUo+aLrK8SFK7nRWGluP3q5LAyVZomCcY3E+SGnlAYdFw2JAjSDfUIx+pXdE1a8e9B0gzXdYakZMIq7IFC4bqx/c2wuQUGRMtXbLsNTYbextlHV5prw/9HphjuZEwoiADVo0VuEwcYxXXZeO8z3Ie0+EZYLuXeAKydRsaUiUEipzShZLj/Odab0uXKK61u7oPvtbNe3J2ABsiWbG4FnIOKZS49eMlbjkeBdIAG1G7fOP1xJ/IHUWEhwbD0yMEO+r865gIoLueiOqFUB+jAA03Y5G1H/jijtIxhy0YHmig6vAV4PQdoWMOp4Ei42Ag4jeBBYCJC4c+xwUkl5vGiWD1eD0wExkQ/xfcKof6QJOhGgk3MBU42QekwuGkajvogwvXAV1SF/gsiqTC4VPVZsjz1Uzr1gG8hLyfjs4xEhi1yHT8smozJKz0nPsBF/A68uiNFhrVNV3AAbNKkX4SX4YMIDufxEuBrdyhT+KBMBh5/zgVGIssGdKBger8DaAJGWhVyG+YZUBExcB/VHe+rpeXNoEAAAAASUVORK5CYII=') no-repeat center;
						background-size: 46rpx 46rpx;
					}
				}

				.text {
					color: #666;
				}

				.time {
					color: #333;

					.day {
						font-size: 24rpx;
					}
				}
			}
		}

		.time {
			width: 200rpx;
			padding-right: 30rpx;
			font-size: 20rpx;
			text-align: right;
			color: #999;

			.day {
				margin-bottom: 4rpx;
			}
		}

		.info {
			position: relative;
			padding-top: 12rpx;
			color: #999;

			&::before {
				content: "";
				position: absolute;
				left: 0;
				top: 0;
				bottom: 0;
				z-index: 1;
				width: 0;
				border-left: 1px solid #e5e5e5;
			}

			.title {
				position: relative;
				margin-bottom: 10rpx;
				padding-left: 32rpx;
				font-size: 28rpx;

				&::before {
					content: "";
					position: absolute;
					left: -3px;
					top: 0;
					bottom: 0;
					z-index: 1;
					width: 7px;
					height: 7px;
					margin: auto 0;
					border-radius: 50%;
					background-color: #cecece;
				}

				&.address {
					font-size: 24rpx;
					color: #333;
				}
			}

			.text {
				padding: 0 0 44rpx 32rpx;
			}
		}
	}
</style>
