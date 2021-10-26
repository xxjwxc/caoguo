<template>
	<view v-if="selectType" class="wallet_class" :class="type?'wallet_class_kq':'wallet_class_gb'">
		<view style="width: 750rpx;height:100vh;" @tap="toCancel"></view>
		<view class="wallet_content" :style="'height:'+ height" :class="type?'wallet_kq':'wallet_gb'">
			<!-- 头部 -->
			<view class="wallet_content_tb" v-if="select=='radio'">
				<view class="cancel0"></view>
				<view class="tite">{{tite}}</view>
				<view class="operate0" @tap="toCancel">
					<uniIcons color="#333" size="32" type="closeempty"></uniIcons>
				</view>
			</view>
			<view class="wallet_content_tb" v-else>
				<view class="cancel1" @tap="toCancel">
					<view class="cencel1_q">
						<uniIcons color="#333" size="32" type="closeempty"></uniIcons>
					</view>
				</view>
				<view class="tite">{{tite}}</view>
				<view class="operate1">
					<view v-if="select=='more'&&checkAll&&list" class="operate1_q" hover-class="click_hover" :style="'color:'+color"
					 @tap="operate">
						{{list.length == item1.length?'清空':'全选'}}
					</view>
					<view v-if="list" class="operate1_p" hover-class="click_hover" @tap="toSelect" :style="'color:'+color">
						确定
					</view>
				</view>
			</view>
			<view>
				<scroll-view v-if="list" class="scroll-Y" scroll-y="true" :style="[style]">
					<view v-for="(items,index) in list" :key="index" class="wallet_list" @tap="toChecked(items)" hover-class="click_hover">
						<view class="text" :style="[textStyle]"> {{items[name]||items}}</view>
						<view v-if="select=='radio'" style="width: 50rpx;">
							<radio style="transform: scale(0.7)" value="r2" :color="color" :checked="items[id]?items[id]==item1[id]:items==item1" />
						</view>
						<view v-else tyle="width: 50rpx;">
							<radio style="transform: scale(0.7)" value="r2" :color="color" :checked="getCheckedS(items)" />
						</view>
					</view>
				</scroll-view>
				<view v-else style="text-align: center;color: #888;line-height: 100rpx;font-size: 26rpx;">暂无数据</view>
			</view>
		</view>
	</view>
</template>

<script>
	import uniIcons from '@/components/uni-icons/uni-icons.vue';
	export default {
		name: 'jp-select',
		components: {
			uniIcons
		},
		props: {
			list: Array, //数据
			tite: { //标题
				type: String,
				default: '请选择'
			},
			name: { //需要显示的内容
				type: String,
				default: 'name'
			},
			id: { //键
				type: String,
				default: 'id'
			},
			select: { //radio单选  more多选
				type: String,
				default: 'radio'
			},
			checkAll: { //是否显示清空
				type: Boolean,
				default: true
			},
			item: { //初始值
				type: [String, Object, Array],
				default: ''
			},
			amount: { //多选时选取最少数量   
				type: Number,
				default: 1
			},
			height: {
				type: String,
				default: '50vh'
			},
			color: { //主题颜色
				type: String,
				default: '#007AFF'
			},
			textColor: { //内容颜色
				type: String,
				default: '#333'
			},
			fontSize: { //内容大小
				type: String,
				default: '30rpx'
			}
		},
		data() {
			return {
				selectType: false,
				type: true,
				item1: '' || {} || [],
				style: {},
				textStyle: {}
			}
		},
		mounted() {
			this.item1 = JSON.parse(JSON.stringify(this.item))
			this.style = {
				'height': `calc( ${this.height} - 80rpx)`
			}
			this.textStyle = {
				'color': `${this.textColor}`,
				'font-size': `${this.fontSize}`
			}
		},
		watch: {
			item() {
				this.item1 = JSON.parse(JSON.stringify(this.item))
			}
		},
		methods: {
			operate() {
				if (this.list.length != this.item1.length) {
					this.item1 = []
					this.list.forEach((el, index) => {
						this.item1.push(el)
					});
				} else {
					this.item1 = []
				}
			},
			toOpen() {
				this.item1 = JSON.parse(JSON.stringify(this.item))
				this.type = true
				this.selectType = true
			},
			toCancel() {
				let taht = this
				this.type = false
				this.$emit('toCancel');
				setTimeout(function clock() {
					taht.selectType = false
					taht.item1 = JSON.parse(JSON.stringify(taht.item))
				}, 200)
			},
			toChecked(item) {
				if (this.select == 'radio') {
					this.oneChecked(item)
				} else {
					this.moreChecked(item)
				}
			},
			oneChecked(item) {
				let taht = this
				this.item1 = item
				this.$emit('checked', item);
				this.type = false
				setTimeout(function clock() {
					taht.selectType = false
				}, 200)
			},
			toSelect() {
				if (this.item1.length >= this.amount) {
					let taht = this
					this.$emit('checked', this.item1);
					this.type = false
					setTimeout(function clock() {
						taht.selectType = false
					}, 200)
				} else {
					uni.showToast({
						title: '至少选择一项!',
						icon: "none"
					})
				}
			},
			moreChecked(item) {
				if (!this.getChecked(item)) {
					this.item1.push(item)
				} else {
					this.item1.forEach((el, index) => {
						if (item[this.id] ? item[this.id] == el[this.id] : el == item) {
							this.item1.splice(index, 1);
						}
					});
				}
			},
			getChecked(item) {
				let c = false
				if (this.item1 && this.item1.length > 0) {
					this.item1.forEach((el, index) => {
						if (item[this.id] ? item[this.id] == el[this.id] : el == item) {
							c = true
						}
					})
				} else {
					this.item1 = []
				}
				return c
			},
			getCheckedS(itemS) {
				let c = false
				if (this.item1 && this.item1.length > 0) {
					this.list.forEach((el, index) => {
						this.item1.forEach((item, index) => {
							if (item[this.id] ? item[this.id] == el[this.id] : el == item) {
								if (item[this.id] ? item[this.id] == itemS[this.id] : itemS == item) {
									c = true
								}
							}
						})
					})
				}
				return c
			}
		}
	}
</script>

<style lang="scss" scoped>
	@keyframes downIn {
		from {
			transform: translateY(100%);
			opacity: 1;
		}

		to {
			transform: translateY(0vh);
			opacity: 1;
		}
	}

	@keyframes downInClose {
		from {
			transform: translateY(0vh);
			opacity: 1;
		}

		to {
			transform: translateY(100%);
			opacity: 1;
		}
	}

	.wallet_class {
		position: fixed;
		z-index: 20;
		height: 100vh;
		width: 750rpx;
		top: 0;
		left: 0;
		background-color: rgba(0, 0, 0, .5);

		.wallet_kq {
			animation: downIn 0.3s;
			-webkit-animation: downIn 0.3s;
		}

		.wallet_gb {
			animation: downInClose 0.3s;
			-webkit-animation: downInClose 0.3s;
		}

		.wallet_content {
			position: fixed;
			bottom: 0;
			left: 0;
			border-top-left-radius: 20rpx;
			border-top-right-radius: 20rpx;
			width: 750rpx;
			background-color: #fff;

			.wallet_content_tb {
				display: flex;
				justify-content: space-between;
				height: 80rpx;
				font-size: 32rpx;
				border-bottom: #ccc solid 1px;
				width: 750rpx;
				line-height: 80rpx;

				.cancel0 {
					width: 100rpx;
					text-align: center;
					font-size: 30rpx;
					color: #333;
				}

				.cancel1 {
					width: 180rpx;
					font-size: 30rpx;
					display: flex;
					justify-content: flex-start;

					.cencel1_q {
						width: 100rpx;
						text-align: center;
						font-size: 30rpx;
						color: #333;
					}
				}

				.tite {
					color: #333;
				}

				.operate0 {
					width: 100rpx;
					text-align: center;
					font-size: 30rpx;
				}

				.operate1 {
					width: 180rpx;
					display: flex;
					justify-content: flex-end;

					.operate1_q {
						width: 80rpx;
						text-align: center;
						font-size: 26rpx;
					}

					.operate1_p {
						width: 100rpx;
						text-align: center;
						font-size: 26rpx;
					}
				}

			}

			.wallet_list {
				height: 60rpx;
				display: flex;
				line-height: 60rpx;
				padding: 5rpx 20rpx;
				justify-content: space-between;

				.text {
					text-align: left;
					width: 660rpx;
					overflow: hidden; //超出的文本隐藏
					text-overflow: ellipsis; //溢出用省略号显示
					white-space: nowrap; //溢出不换行
				}

			}
		}
	}

	.click_hover {
		animation: showPopup 0.2s linear both;
	}

	@keyframes showPopup {
		0% {
			opacity: 0;
			background-color: #d8d8d8;
		}

		100% {
			opacity: 1;
		}
	}
</style>
