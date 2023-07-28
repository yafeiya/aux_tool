<template>
    <Button @click="anim2 = !anim2">显示动画</Button>
    <transition name="ease">
        <div v-show="anim2" class="transition-box">缩放</div>
    </transition>
</template>
<script>
export default {
    data () {
        return {
            anim1: true,
            anim2: false,
            anim3: true,
            anim4: true,
            anim5: true,
            anim6: true,
            anim7: true,
            anim8: true,
            anim9: true,
            anim10: true,
            anim11: true
        }
    },
    beforeMount() {
        anim2 = true;
    }
}
</script>
<style>
.transition-box {
    margin-top: 32px;
    width: 100%;
    border-radius: 4px;
    background-color: #409EFF;
    text-align: center;
    color: #fff;
    padding: 24px 0;
}
</style>