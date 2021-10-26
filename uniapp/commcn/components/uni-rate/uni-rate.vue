<template>
	<view>
		<view
		    ref="uni-rate"
		    class="uni-rate"
		>
			<view
			    class="uni-rate__icon"
			    :style="{ 'margin-right': margin + 'px' }"
			    v-for="(star, index) in stars"
			    :key="index"
			    @touchstart.stop="touchstart"
			    @touchmove.stop="touchmove"
			>
				<uni-icons
				    :color="color"
				    :size="size"
				    :type="isFill ? 'star-filled' : 'star'"
				/>
				<!-- #ifdef APP-NVUE -->
				<view
				    :style="{ width: star.activeWitch.replace('%','')*size/100+'px'}"
				    class="uni-rate__icon-on"
				>
					<uni-icons
					    style="text-align: left;"
					    :color="disabled?'#ccc':activeColor"
					    :size="size"
					    type="star-filled"
					/>
				</view>
				<!-- #endif -->
				<!-- #ifndef APP-NVUE -->
				<view
				    :style="{ width: star.activeWitch}"
				    class="uni-rate__icon-on"
				>
					<uni-icons
					    :color="disabled?disabledColor:activeColor"
					    :size="size"
					    type="star-filled"
					/>
				</view>
				<!-- #endif -->
			</view>
		</view>
	</view>
</template>

<script>
	// #ifdef APP-NVUE
	const dom = uni.requireNativePlugin('dom');
	// #endif
</script>

<style
    lang="scss"
    scoped
>
	.uni-rate {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		line-height: 1;
		font-size: 0;
		flex-direction: row;
	}

	.uni-rate__icon {
		position: relative;
		line-height: 1;
		font-size: 0;
	}

	.uni-rate__icon-on {
		overflow: hidden;
		position: absolute;
		top: 0;
		left: 0;
		line-height: 1;
		text-align: left;
	}
</style>
