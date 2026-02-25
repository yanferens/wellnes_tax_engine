<script setup lang="ts">

</script>

<template>
  <div class="logo_wrap">
    <img src="@/assets/logo.png" alt="logo" class="logo"/>
    <div class="logo_cont">
      назва
      <span>податки</span>
    </div>
  </div>
</template>

<style scoped>
  .logo_wrap{
    display: flex;
    flex-direction: row;
    align-items: center;
    column-gap: 50px;
  }
  .logo{
    width: 70px;
    height: 60px;
  }
  .logo_cont{
    display: flex;
    flex-direction: column;
    align-items: center;
    color: #000;
    text-align: center;
    font-size: 19px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
    text-shadow: 0 5px 3px rgba(0, 0, 0, 0.15);
  }
  .logo_cont span{
    color: rgba(99, 99, 99, 0.96);
    font-size: 19px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
</style>